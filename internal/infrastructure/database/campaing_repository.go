package database

import "emailn/internal/domain/campaing"

type CampaingRepository struct {
	campains []campaing.Campaing
}

func (c *CampaingRepository) Save(campaing *campaing.Campaing) error {
	c.campains = append(c.campains, *campaing)
	return nil
}
