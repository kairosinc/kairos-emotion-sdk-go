/*
 * kairosemotionapi
 *
 * This file was automatically generated for kairos by APIMATIC BETA v2.0 on 01/15/2016
 */
package emotionanalysis

import(
    "encoding/json"
    "github.com/apimatic/unirest-go"
    "kairosemotionapi"
    "kairosemotionapi/apihelper"
    "kairosemotionapi/models"
)
/*
 * Client structure as interface implementation
 */
type EMOTIONANALYSIS_IMPL struct { }

/**
 * Create a new media object to be processed.
 * @param    *string        source     parameter: Optional
 * @return	Returns the *models.MediaResponse response from the API call
 */
func (me *EMOTIONANALYSIS_IMPL) CreateMedia (
            source *string) (*models.MediaResponse, error) {
    //the base uri for api requests
    queryBuilder := kairosemotionapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/media"

    //variable to hold errors
    var err error = nil
    
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithQueryParameters(queryBuilder, map[string]interface{} {
        "source" : source,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
    }

    //prepare API request
    request := unirest.Post(queryBuilder, headers)

    //Custom Authentication to be added by the developers for authorization
    apihelper.AppendCustomAuthParams(request)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper.NewAPIError("HTTP Response Not OK" , response.Code, response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal *models.MediaResponse = &models.MediaResponse{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Returns the results of a specific uploaded piece of media.
 * @param    string        id     parameter: Required
 * @return	Returns the interface{} response from the API call
 */
func (me *EMOTIONANALYSIS_IMPL) GetMediaById (
            id string) (interface{}, error) {
    //the base uri for api requests
    queryBuilder := kairosemotionapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/media/{id}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithTemplateParameters(queryBuilder, map[string]interface{} {
        "id" : id,
    }) 
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
    }

    //prepare API request
    request := unirest.Get(queryBuilder, headers)

    //Custom Authentication to be added by the developers for authorization
    apihelper.AppendCustomAuthParams(request)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper.NewAPIError("HTTP Response Not OK" , response.Code, response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal interface{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Delete media results. It returns the status of the operation.
 * @param    string        id     parameter: Required
 * @return	Returns the *models.MediaByIdResponse response from the API call
 */
func (me *EMOTIONANALYSIS_IMPL) DeleteMediaById (
            id string) (*models.MediaByIdResponse, error) {
    //the base uri for api requests
    queryBuilder := kairosemotionapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/media/{id}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithTemplateParameters(queryBuilder, map[string]interface{} {
        "id" : id,
    }) 
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
    }

    //prepare API request
    request := unirest.Delete(queryBuilder, headers)

    //Custom Authentication to be added by the developers for authorization
    apihelper.AppendCustomAuthParams(request)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper.NewAPIError("HTTP Response Not OK" , response.Code, response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal *models.MediaByIdResponse = &models.MediaByIdResponse{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

