import configparser
import datetime
import json
import logging
#import lxml
import os
from pprint import pprint
import sys
import time

import boto3
from bs4 import BeautifulSoup
import requests
from slackclient import SlackClient
import wget

#logging.basicConfig()
logger = logging.getLogger(__name__)
#logger.setLevel(logging.DEBUG)

#config_path = '../../config/config.ini'


class FlyOnTheWall:
    ## TO DO: ##
    # - Create independent directory for thread archive storage [DONE]
    # - Create independent directory for downloads [DONE]
    # - Add "last updated" timestamp after check complete [DONE]

    url_base = 'https://boards.4chan.org/'
    url_board = 'biz/'
    url_thread_base = 'thread/'

    loop_time = 60

    keyword_file = 'keywords.txt'

    thread_limit = 99   # If set to below number of threads on front page, will only analyze that many number of threads

    excluded_threads = ['4884770', '904256']    # Pinned FAQ and general info threads

    thread_archive_file = 'thread_archive.json'


    def __init__(self, exchange, market, config_path,
                 board=url_board,
                 thread_limit=thread_limit,
                 keyword_data=None,
                 keyword_file=keyword_file,
                 excluded_threads=excluded_threads,
                 thread_archive_file=thread_archive_file,
                 slack_thread=None,
                 pages=1,
                 persistent=False, persistent_loop_time=1800, persistent_loops=5,
                 analyze_sentiment=False, sentiment_results_max=None):
        self.url_board_base = FlyOnTheWall.url_base + board

        self.url_prefix = FlyOnTheWall.url_base + board + FlyOnTheWall.url_thread_base

        self.pages = pages

        self.persistent = persistent

        self.keyword_data = keyword_data

        self.keyword_file = keyword_file

        if self.keyword_data == None and not os.path.exists(self.keyword_file):
            logger.error('Could not find keyword file at path \'' + self.keyword_file + '\'. Exiting.')

            sys.exit(1)

        self.slack_thread = slack_thread

        if not os.path.exists('json/'):
            os.mkdir('json/')

        self.persistent_loop_time = persistent_loop_time

        self.persistent_loops = persistent_loops

        self.thread_limit = thread_limit

        self.excluded_threads = excluded_threads

        market_raw = market.lower()

        if '-' in market_raw:
            market_compact = market_raw.split('-')[0] + market_raw.split('-')[1]

        elif '/' in market_raw:
            market_compact = market_raw.split('/')[0] + market_raw.split('/')[1]

        else:
            market_compact = market_raw

        logger.debug('market_compact: ' + market_compact)

        market_directory = 'json/' + exchange.lower() + '_' + market_compact + '/'

        self.product_directory = market_directory + datetime.datetime.now().strftime('%m%d%Y_%H%M%S') + '/'

        self.download_directory = self.product_directory + 'downloads/'

        if not os.path.exists(market_directory):
            logger.debug('Creating market directory: ' + market_directory)

            os.mkdir(market_directory)

        if not os.path.exists(self.product_directory):
            logger.debug('Creating product directory: ' + self.product_directory)

            os.mkdir(self.product_directory)

        if not os.path.exists(self.download_directory):
            logger.debug('Creating download directory: ' + self.download_directory)

            os.mkdir(self.download_directory)

        self.thread_archive_file = self.product_directory + thread_archive_file

        self.analyze_sentiment = analyze_sentiment

        if self.analyze_sentiment == True:
            self.comprehend_client = boto3.client(service_name='comprehend', region_name='us-east-1')

            self.comprehend_results_file = self.product_directory + 'comprehend_results.json'

        self.sentiment_results_max = sentiment_results_max

        self.keyword_list = []
        self.excluded_list = []

        self.thread_archive = {}

        self.last_updated = None

        if self.slack_thread != None:
            config = configparser.ConfigParser()
            config.read(config_path)

            slack_token = config['slack']['slack_token_test']

            self.slack_client = SlackClient(slack_token)

            self.slack_alert_user = config['settings']['slack_alert_user']

            self.slack_user_icon = config['settings']['slack_user_icon']

        else:
            self.slack_client = None


    def purge_save_data(self):
        logger.info('Purging save data.')

        if os.path.exists(self.thread_archive_file):
            logger.debug('Removing thread archive json file.')

            os.remove(self.thread_archive_file)

        downloads = os.listdir(self.download_directory)

        for dl in downloads:
            file_path = self.download_directory + dl

            logger.debug('Removing download: ' + file_path)

            os.remove(file_path)


    def load_keywords(self):
        logger.debug('Loading keyword file.')

        if self.keyword_data == None:
            keyword_list_read = self.keyword_data

        else:
            with open(self.keyword_file, 'r', encoding='utf-8') as kw:
                keyword_list_read = kw.read().split()

        #keyword_list = []
        #excluded_list = []
        for keyword in keyword_list_read:
            if '|' in keyword:
                keyword_split = keyword.split('|')

                self.keyword_list.append(keyword_split[0])

                if ',' in keyword_split[1]:
                    excluded_split = keyword_split[1].split(',')

                    for word in excluded_split:
                        if word != '' and word != '\n' and word != '\r':
                            self.excluded_list.append(word)

                else:
                    self.excluded_list.append(keyword_split[1])

            #else:
            elif keyword != '' and keyword != '\n' and keyword != '\r':
                    self.keyword_list.append(keyword)

        logger.debug('self.keyword_list: ' + str(self.keyword_list))
        logger.debug('self.excluded_list: ' + str(self.excluded_list))

        logger.info('----- KEYWORDS -----')
        for word in self.keyword_list:
            logger.info(word.upper())

        logger.info('-- EXCLUDED WORDS --')
        for word in self.excluded_list:
            logger.info(word.upper())

        #time.sleep(1)


    def run_search(self):
        def send_slack_alert(channel_id, message, thread_id=None):
            alert_result = True

            try:
                #alert_message = 'TEST MESSAGE'

                """
                attachment_array =  [{"fallback": "New exchange, " + alert_data['exchange'] + ", added to Coinigy.",
                                      "color": "#FFA500",
                                      "pretext": "Exchange Name: *" + alert_data['exchange'] + "*",
                                      "title": "Click here to visit " + alert_data['exchange'] + ".",
                                      "title_link": alert_data['url']}]

                attachments = json.dumps(attachment_array)
                """

                ############################################

                slack_client.api_call(
                    'chat.postMessage',
                    channel=channel_id,
                    text=message,
                    username=self.slack_alert_user,
                    #icon_emoji=slack_alert_user_icon,
                    icon_url=self.slack_user_icon,
                    thread_ts=thread_id
                    #attachments=attachments
                )

                ###########################################

            except Exception as e:
                logger.exception('Exception while sending Slack alert.')
                logger.exception(e)

                alert_result = False

            finally:
                return alert_result


        def get_threads():#self):
            try:
                logger.debug('Retrieving thread list.')

                url = self.url_board_base

                r = requests.get(url)

                soup = BeautifulSoup(r.text, 'html.parser')# 'lxml')

                threads = soup.find_all('a', attrs={'class': 'replylink'})

                thread_list = []
                for thread in threads:
                    thread_num = thread['href'].split('/')[1]

                    thread_list.append(thread_num)

                return thread_list

            except Exception as e:
                logger.exception('Exception while getting threads.')
                logger.exception(e)

                #raise


        def get_posts(thread_num):#self, thread_num):    #(thread_num)
            try:
                logger.debug('Retrieving posts.')

                url = self.url_prefix + thread_num

                r = requests.get(url)

                soup = BeautifulSoup(r.text, 'html.parser')#'lxml')

                #posts = soup.find_all('blockquote', attrs={'class': 'postMessage'})

                data = soup.find_all('div', attrs={'class': ['post op', 'post reply']})

                post_list = []
                for post in data:
                    post_data = {}

                    post_data['post'] = post.text

                    attachments = post.find_all('a', attrs={'class': 'fileThumb'})

                    # ONLY GETS LAST FILE (MULTIPLE POSSIBLE?)
                    for file in attachments:
                        file_url = 'https:' + file['href']
                        logger.debug('file_url: ' + file_url)

                        post_data['file'] = file_url

                    post_list.append(post_data)
                    logger.debug('post_data: ' + str(post_data))

                logger.debug('post_list: ' + str(post_list))

                return post_list


            except Exception as e:
                logger.exception('Exception while getting posts.')
                logger.exception(e)

                #raise


        def filter_threads():#self):   #(threads, keywords)
            try:
                logger.debug('Filtering threads.')

                threads = self.thread_archive

                threads_filtered = {}

                thread_count = len(threads)

                thread_loops = 0
                for key in threads:
                    thread_loops += 1

                    logger.info('Thread #' + str(thread_loops) + ' of ' + str(thread_count))

                    posts = threads[key]
                    logger.debug('posts: ' + str(posts))

                    found_list = []

                    post_count = len(posts)

                    post_loops = 0
                    for post in posts:
                        post_loops += 1

                        logger.info('Post #' + str(post_loops) + ' of ' + str(post_count))

                        word_count = len(self.keyword_list)

                        word_loops = 0
                        for word in self.keyword_list:
                            word_loops += 1

                            logger.debug('Word #' + str(word_loops) + ' of ' + str(word_count))

                            if word in post['post'].lower():
                                #logger.debug('FOUND: ' + word)
                                passed_excluded = True

                                for excluded in self.excluded_list:
                                    if excluded in post['post'].lower():
                                        passed_excluded = False

                                        logger.debug('Found excluded word: ' + excluded)

                                        logger.debug('Excluding post: ' + str(post))

                                if passed_excluded == True:
                                    entry = word + '|' + post['post'].lower()

                                    if 'file' in post:
                                        entry = entry + '|' + post['file']

                                        logger.info('Downloading attachment.')

                                        file_name = wget.detect_filename(url=post['file'])
                                        logger.debug('file_name: ' + file_name)

                                        if not os.path.isfile(self.download_directory + file_name):
                                            dl_file = wget.download(post['file'], out=self.download_directory.rstrip('/'))

                                            logger.debug('Successful download: ' + dl_file)

                                    found_list.append(entry)

                    if len(found_list) > 0:
                        threads_filtered[key] = found_list

                return threads_filtered


            except Exception as e:
                logger.exception('Exception while filtering threads.')
                logger.exception(e)

                #raise


        def create_trimmed_archive():
            try:
                thread_archive_trimmed = {'Exception': False}

                for thread in self.thread_archive:
                    thread_archive_trimmed[thread] = []

                    for post in self.thread_archive[thread]:
                        post_truncated = post['post']

                        logger.debug('[PRE] post_truncated: ' + post_truncated)

                        while (True):
                            post_num_index = post_truncated.lower().find('no.')

                            if post_num_index == -1:
                                while (True):
                                    post_num_index = post_truncated.lower().find('>>')

                                    if post_num_index == -1:
                                        break

                                    else:
                                        post_truncated = post_truncated[(post_num_index + 2):]

                                break

                            else:
                                post_truncated = post_truncated[(post_num_index + 3):]

                        first_letter_index = 0
                        for x in range(0, len(post_truncated)):
                            if post_truncated[x].isalpha():
                                first_letter_index = x

                                break

                        post_truncated = post_truncated[first_letter_index:]
                        logger.debug('[POST] post_truncated: ' + post_truncated)

                        thread_archive_trimmed[thread].append(post_truncated)

            except Exception as e:
                logger.exception('Exception while creating trimmed archive.')
                logger.exception(e)

                thread_archive_trimmed['Exception'] = True

            finally:
                return thread_archive_trimmed


        def get_entities(input_text):
            entities = None
            metadata = None

            try:
                comprehend_results = self.comprehend_client.detect_entities(Text=input_text, LanguageCode='en')

                entities = comprehend_results['Entities']
                metadata = comprehend_results['ResponseMetadata']

            except Exception as e:
                logger.exception('Exception while getting entities.')
                logger.exception(e)

            finally:
                return entities, metadata


        def get_key_phrases(input_text):
            key_phrases = None
            metadata = None

            try:
                comprehend_results = self.comprehend_client.detect_key_phrases(Text=input_text, LanguageCode='en')

                key_phrases = comprehend_results['KeyPhrases']
                metadata = comprehend_results['ResponseMetadata']

            except Exception as e:
                logger.exception('Exception while getting entities.')
                logger.exception(e)

            finally:
                return key_phrases, metadata


        def get_sentiment(input_text):
            sentiment = None
            metadata = None

            try:
                comprehend_results = self.comprehend_client.detect_sentiment(Text=input_text, LanguageCode='en')

                sentiment = {'sentiment': comprehend_results['Sentiment'], 'score': comprehend_results['SentimentScore']}
                metadata = comprehend_results['ResponseMetadata']

            except Exception as e:
                logger.exception('Exception while getting entities.')
                logger.exception(e)

            finally:
                return sentiment, metadata

        def threshold_sentiment_results(sentiment_results):
            sentiment_results_thresholded = {'positive': None, 'negative': None}

            positive_sentiment = []
            negative_sentiment = []
            for result in comprehend_results:
                if result['sentiment']['sentiment'] == 'POSITIVE':
                    positive_sentiment.append(result)

                elif result['sentiment']['sentiment'] == 'NEGATIVE':
                    negative_sentiment.append(result)

            positive_sentiment_sorted = sorted(positive_sentiment,
                                               key=lambda sent: sent['sentiment']['score']['Positive'],
                                               reverse=True)

            logger.debug('Removing positive results with score < ' + str(positive_score_threshold) + '.')

            for result in positive_sentiment_sorted:
                if result['sentiment']['score']['Positive'] < positive_score_threshold:
                    positive_sentiment_sorted.remove(result)

            sentiment_results_thresholded['positive'] = positive_results_sorted

            negative_sentiment_sorted = sorted(negative_sentiment,
                                               key=lambda sent: sent['sentiment']['score']['Negative'],
                                               reverse=True)

            logger.debug('Removing negative results with score < ' + str(negative_score_threshold) + '.')

            for result in negative_sentiment_sorted:
                if result['sentiment']['score']['Negative'] < negative_score_threshold:
                    negative_sentiment_sorted.remove(result)

            sentiment_results_thresholded['negative'] = negative_results_sorted

            if self.sentiment_results_max != None:
                for result in sentiment_results_thresholded:
                    if len(sentiment_results_thresholded[result]) > self.sentiment_results_max:
                        sentiment_results_thresholded[result] = sentiment_results_thresholded[result][:self.sentiment_results_max]

            return sentiment_results_thresholded


        try:
            loop_count = 0
            while (True):
                loop_count += 1
                logger.debug('loop_count: ' + str(loop_count))

                logger.info('Retrieving threads.')

                thread_list = get_threads()

                for thread in self.excluded_threads:
                    if thread in thread_list:
                        logger.debug('Removing excluded thread: ' + thread)

                        thread_list.remove(thread)

                logger.info('Gathering posts from threads.')

                thread_count = len(thread_list)

                thread_loops = 0
                for thread in thread_list:
                    thread_loops += 1

                    logger.info('Thread #' + str(thread_loops) + ' of ' + str(thread_count))

                    if thread not in self.thread_archive:
                        self.thread_archive[thread] = []

                    post_list = get_posts(thread_num=thread)
                    logger.debug('post_list: ' + str(post_list))

                    #thread_archive[thread] = post_list
                    for post in post_list:
                        if post not in self.thread_archive[thread]:
                            logger.debug('Appending post: ' + str(post))

                            self.thread_archive[thread].append(post)

                        else:
                            logger.debug('Skipping post. Already in archive.')

                    if thread_loops == self.thread_limit:
                        logger.debug('Thread limit reached. Breaking early.')

                        break

                logger.info('Filtering threads for relevant content.')

                #time.sleep(3)

                relevant_threads = filter_threads()#self.thread_archive, self.keyword_list)

                if os.path.isfile(self.thread_archive_file):
                    with open(self.thread_archive_file, 'r', encoding='utf-8') as file:
                        thread_data = json.load(file)

                else:
                    thread_data = {}

                for thread in relevant_threads:
                    logger.info('Relevant thread: ' + str(thread))

                    if thread not in thread_data:
                        thread_data[thread] = []

                    for post in relevant_threads[thread]:
                        logger.info('Relevant post: ' + post)

                        if post not in thread_data[thread]:
                            thread_data[thread].append(post)

                logger.info('Writing relevant posts to json archive.')

                with open(self.thread_archive_file, 'w', encoding='utf-8') as file:
                    json.dump(thread_data, file, indent=4, sort_keys=True, ensure_ascii=False)

                self.last_updated = datetime.datetime.now()

                if self.analyze_sentiment == True:
                    thread_archive_trimmed = create_trimmed_archive()

                    if thread_archive_trimmed['Exception'] == True:
                        logger.warning('Failed to create trimmed thread archive. Bypassing sentiment analysis.')

                    else:
                        del thread_archive_trimmed['Exception']

                        sentiment_results_list = []

                        sentiment_total = 0
                        for thread in thread_archive_trimmed:
                            if thread != 'Exception':
                                sentiment_total += len(thread_archive_trimmed[thread])

                        logger.debug('sentiment_total: ' + str(sentiment_total))

                        sentiment_count = 0
                        for thread in thread_archive_trimmed:
                            logger.debug('thread: ' + thread)

                            if thread != 'Exception':
                                for post in thread_archive_trimmed[thread]:
                                    logger.debug('post: ' + post)

                                    sentiment_count += 1

                                    logger.info('Analyzing sentiment for post #' + str(sentiment_count) + ' of ' + str(sentiment_total) + '.')

                                    sentiment_results = {'entities': [], 'entities_metadata': None,
                                                         'key_phrases': [], 'key_phrases_metadata': None,
                                                         'sentiment': [], 'sentiment_metadata': None,
                                                         'post': ''}

                                    sentiment_results['post'] = post

                                    sentiment_results['entities'], sentiment_results['entities_metadata'] = get_entities(post)

                                    #print(sentiment_results['entities'], sentiment_results['entities_metadata'])

                                    time.sleep(0.05)    # To prevent hitting API throttling limit (20/sec)

                                    sentiment_results['key_phrases'], sentiment_results['key_phrases_metadata'] = get_key_phrases(post)

                                    #print(sentiment_results['key_phrases'], sentiment_results['key_phrases_metadata'])

                                    time.sleep(0.05)    # To prevent hitting API throttling limit (20/sec)

                                    sentiment_results['sentiment'], sentiment_results['sentiment_metadata'] = get_sentiment(post)

                                    #print(sentiment_results['sentiment'], sentiment_results['sentiment_metadata'])

                                    time.sleep(0.05)    # To prevent hitting API throttling limit (20/sec)

                                    sentiment_results_list.append(sentiment_results)

                        #print('Sentiment Analysis Results:')
                        #pprint(sentiment_results_list)

                        logger.info('Sorting sentiment results and thresholding by score.')

                        sentiment_thresholded = threshold_sentiment_results(sentiment_results_list)

                        logger.info('Dumping sentiment analysis results to json file.')

                        with open(self.comprehend_results_file, 'w', encoding='utf-8') as file:
                            #json.dump(sentiment_results_list, file, indent=4, sort_keys=True, ensure_ascii=False)
                            json.dump(sentiment_thresholded, file, indent=4, sort_keys=True, ensure_ascii=False)

                        if self.slack_client != None:
                            #### SEND ALL NECESSECARY ALERT MESSAGES ####
                            # Use sentiment_thresholded.pop()?
                            # Don't proceed to next alert until success confirmation
                            # Make sure alerts sent from highest to lowest score

                            pass

                if self.persistent == True and loop_count < self.persistent_loops:
                    logger.info('Sleeping ' + str(self.loop_time) + ' seconds.')

                    time.sleep(self.persistent_loop_time)

                elif self.persistent == True:
                    logger.info('Completed all ' + str(self.persistent_loops) + ' persistence loops.')

                    break

                else:
                    logger.info('Persistent mode disabled. Search complete.')

                    break

            logger.debug('Completed successfully. Exiting.')

        except botocore.exceptions.NoCredentialsError as e:
            logger.exception('botocore.exceptions.NoCredentialsError raised in run_search().')
            logger.exception(e)

        except Exception as e:
            logger.exception('Exception raised in run_search().')
            logger.exception(e)

        except KeyboardInterrupt:
            logger.info('Exit signal received.')

            sys.exit()


if __name__ == '__main__':
    # 1 - Get all threads on main page
    # 2 - Save thread numbers to list
    # 3 - For each:
    #       a) get_posts()
    #       b) Search posts for keywords
    #           i. If found, save text (and media) to directory

    sample_keyword_data = ['stellar', 'lumens', 'hyperledger', 'fairx', 'xlm']

    config_path = '../../config/config.ini'

    flyonthewall = FlyOnTheWall(exchange='TestExchange', market='TEST-MARKET',keyword_data=sample_keyword_data, config_path=config_path,
                                slack_thread=None, persistent=False, analyze_sentiment=True, sentiment_results_max=20)
                                #keyword_file='keywords.txt', persistent=False, analyze_sentiment=True)

    flyonthewall.purge_save_data()

    if flyonthewall.keyword_data == None:
        flyonthewall.load_keywords(flyonthewall.keyword_file)

    else:
        flyonthewall.load_keywords()

    flyonthewall.run_search()

    print('Last Updated: ', flyonthewall.last_updated)

    #print('Sleeping for 10 seconds.')

    #time.sleep(10)

    #flyonthewall.run_search()

    #print('Last Updated: ', flyonthewall.last_updated)

    print('Done.')
