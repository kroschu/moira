package controller

import (
	"fmt"
	"github.com/moira-alert/moira-alert"
	"github.com/moira-alert/moira-alert/api"
	"github.com/moira-alert/moira-alert/api/dto"
)

// GetAllTagsAndSubscriptions get tags subscriptions and triggerIDs
func GetAllTagsAndSubscriptions(database moira.Database, logger moira.Logger) (*dto.TagsStatistics, *api.ErrorResponse) {
	tagsNames, err := database.GetTagNames()
	if err != nil {
		return nil, api.ErrorInternalServer(err)
	}

	tagsStatistics := dto.TagsStatistics{
		List: make([]dto.TagStatistics, 0, len(tagsNames)),
	}
	rch := make(chan *dto.TagStatistics, len(tagsNames))

	for _, tagName := range tagsNames {
		go func(tagName string) {
			tagStat := &dto.TagStatistics{}
			tagStat.TagName = tagName
			tagStat.Subscriptions, err = database.GetTagsSubscriptions([]string{tagName})
			if err != nil {
				logger.Error(err.Error())
				rch <- nil
			}
			tagStat.Triggers, err = database.GetTagTriggerIDs(tagName)
			if err != nil {
				logger.Error(err.Error())
				rch <- nil
			}
			rch <- tagStat
		}(tagName)
	}

	for i := 0; i < len(tagsNames); i++ {
		if r := <-rch; r != nil {
			tagsStatistics.List = append(tagsStatistics.List, *r)
		}
	}
	return &tagsStatistics, nil
}

// GetAllTags gets all tag names
func GetAllTags(database moira.Database) (*dto.TagsData, *api.ErrorResponse) {
	tagsNames, err := database.GetTagNames()
	if err != nil {
		return nil, api.ErrorInternalServer(err)
	}

	tagsData := &dto.TagsData{
		TagNames: tagsNames,
	}

	return tagsData, nil
}

// RemoveTag deletes tag by name
func RemoveTag(database moira.Database, tagName string) (*dto.MessageResponse, *api.ErrorResponse) {
	triggerIDs, err := database.GetTagTriggerIDs(tagName)
	if err != nil {
		return nil, api.ErrorInternalServer(err)
	}

	if len(triggerIDs) > 0 {
		return nil, api.ErrorInvalidRequest(fmt.Errorf("This tag is assigned to %v triggers. Remove tag from triggers first", len(triggerIDs)))
	}
	if err = database.RemoveTag(tagName); err != nil {
		return nil, api.ErrorInternalServer(err)
	}
	return &dto.MessageResponse{Message: "tag deleted"}, nil
}
