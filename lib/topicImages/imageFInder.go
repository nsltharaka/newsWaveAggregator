package topicImages

type ImageFinder func(topic string) (imageUrls []string, err error)
