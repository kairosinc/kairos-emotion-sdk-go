/*
 * kairosemotionapi
 *
 * This file was automatically generated for kairos by APIMATIC BETA v2.0 on 01/15/2016
 */

package emotionanalysis


import "kairosemotionapi/models"

/*
 * Interface for the EMOTIONANALYSIS_IMPL
 */
type EMOTIONANALYSIS interface {
    CreateMedia (*string) (*models.MediaResponse, error)

    GetMediaById (string) (interface{}, error)

    DeleteMediaById (string) (*models.MediaByIdResponse, error)
}

/*
 * Factory for the EMOTIONANALYSIS interaface returning EMOTIONANALYSIS_IMPL
 */
func NewEMOTIONANALYSIS() EMOTIONANALYSIS {
    return &EMOTIONANALYSIS_IMPL{}
}