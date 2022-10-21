package gorm

import (
	"final_project/server/models"
	"final_project/server/repositories"

	"gorm.io/gorm"
)

type socialMediaRepo struct {
	db *gorm.DB
}

func NewSocialMediaRepo(db *gorm.DB) repositories.SocialMediaRepo {
	return &socialMediaRepo{db: db}
}

func (s *socialMediaRepo) CreateSocialMedia(socialMedia *models.SocialMedia) error {
	return s.db.Create(socialMedia).Error
}

func (s *socialMediaRepo) GetSocialMedia() (*[]models.SocialMedias, error) {
	var socMeds []models.SocialMedias
	err := s.db.Model(&models.SocialMedia{}).Select("social_media.id, social_media.name, social_media.social_media_url, social_media.user_id, social_media.created_at, social_media.updated_at, users.username, users.email").Joins("left join users on users.id = social_media.user_id").Find(&socMeds).Error
	if err != nil {
		return nil, err
	}

	return &socMeds, nil
}

func (s *socialMediaRepo) UpdateSocialMedia(id int, req *models.SocialMedia) error {
	socmed := models.SocialMedia{}
	err := s.db.Model(&socmed).Where("id = ?", id).Updates(models.SocialMedia{Name: req.Name, Social_media_url: req.Social_media_url}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *socialMediaRepo) DeleteSocMed(id int) error {
	socmed := models.SocialMedia{}
	err := s.db.Where("id=?", id).Delete(&socmed).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *socialMediaRepo) GetSocMedById(id int) (*models.SocialMedia, error) {
	socmed := models.SocialMedia{}
	err := s.db.First(&socmed, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &socmed, nil
}

func (s *socialMediaRepo) DeleteSocmedByUserId(user_id int) error {
	socmed := models.SocialMedia{}
	err := s.db.Where("user_id=?", user_id).Delete(&socmed).Error
	if err != nil {
		return err
	}
	return nil
}
