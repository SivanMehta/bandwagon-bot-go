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

- [x] Deploy to [Elastic Beanstalk](https://aws.amazon.com/elasticbeanstalk/) with
[`eb deploy`]

```
$ eb deploy
2019-06-18 21:05:22    INFO: Environment update is starting.
2019-06-18 21:05:27    INFO: Deploying new version to instance(s).
2019-06-18 21:05:53    INFO: New application version was deployed to running EC2 instances.
2019-06-18 21:05:53    INFO: Environment update completed successfully.
```

- [x] Set environment variables with [`eb setenv`]

You can generate values for these variable by creating an app on Twitter [here].

```
$ eb setenv KEY=TWITTER_CONSUMER_KEY SECRET=TWITTER_CONSUMER_SECRET_KEY
2019-06-19 21:05:25    INFO: Environment update is starting.
2019-06-19 21:05:29    INFO: Updating environment tmp-dev's configuration settings.
2019-06-19 21:06:50    INFO: Successfully deployed new configuration to environment.
2019-06-19 21:06:51    INFO: Environment update completed successfully.
```

# Running Locally

Similar to above you can run the application locally if you manually set the
`KEY` and `SECRET` environment variables. The `Makefile` already has a command
ready for you to go:

```
$ KEY=TWITTER_CONSUMER_KEY SECRET=TWITTER_CONSUMER_SECRET_KEY make serve
go run application.go
Generated Markov chain for Taylor Swift
Generated Markov chain for Whoopi
Generated Markov chain for Hump Day
Generated Markov chain for Congress
Generated Markov chain for #KickOffTheSummerBy
Listening on port 5000
```

[`eb setenv`]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/eb3-setenv.html
[`eb deploy`]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/eb3-deploy.html
[here]: https://developer.twitter.com/en/apps
