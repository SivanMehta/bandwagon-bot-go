This is a bot that hops on bandwagons and pretends to be a real fan.

- [x] Completed Feature
- [ ] Incomplete Feature

# Basic Architecture

- [x] On a set interval do the following:
- [x] Get trending topics on twitter in the US
- [x] Get the top tweets for that topic
- [x] Generate a markov chain for each trending topic using the tweets as a corpus
- [x] Expose a function on `/api/:topic` that allows you to play with each chain
- [ ] Post a tweet for each trending topic and post it to Twitter

# Deployment

- [ ] Deploy to [Elastic Beanstalk](https://aws.amazon.com/elasticbeanstalk/) with
[`eb deploy`]

```
$ eb deploy
2019-06-18 21:05:22    INFO: Environment update is starting.
2019-06-18 21:05:27    INFO: Deploying new version to instance(s).
2019-06-18 21:05:53    INFO: New application version was deployed to running EC2 instances.
2019-06-18 21:05:53    INFO: Environment update completed successfully.
```

- [ ] Set environment variables with [`eb setenv`]

```
$ eb setenv KEY=TWITTER_CONSUMER_KEY SECRET=TWITTER_CONSUMER_SECRET_KEY
2019-06-19 21:05:25    INFO: Environment update is starting.
2019-06-19 21:05:29    INFO: Updating environment tmp-dev's configuration settings.
2019-06-19 21:06:50    INFO: Successfully deployed new configuration to environment.
2019-06-19 21:06:51    INFO: Environment update completed successfully.
```


[`eb setenv`]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/eb3-setenv.html
[`eb deploy`]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/eb3-deploy.html
