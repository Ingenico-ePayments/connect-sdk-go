// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package disputes

import "github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"

// UploadFileRequest represents multipart/form-data parameters for Upload File
// Documentation can be found at https://epayments-api.developer-ingenico.com/fileserviceapi/v1/en_US/go/disputes/uploadFile.html
type UploadFileRequest struct {
	File *communication.UploadableFile
}

// ToMultipartFormDataObject converts the multipart/form-data request to communication.MultipartFormDataObject
func (request *UploadFileRequest) ToMultipartFormDataObject() *communication.MultipartFormDataObject {
	multipart, _ := communication.NewMultipartFormDataObject()

	if request.File != nil {
		multipart.AddFile("file", *request.File)
	}

	return multipart
}

// NewUploadFileRequest constructs an instance of UploadFileRequest
func NewUploadFileRequest() *UploadFileRequest {
	return &UploadFileRequest{}
}
