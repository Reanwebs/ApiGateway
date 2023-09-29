package models

type RegisterRequestBody struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	CPassword   string `json:"cPassword"`
	PhoneNumber string `json:"phoneNumber"`
	UserName    string `json:"userName"`
	Referral    string `json:"referral"`
	Otp         string `json:"otp"`
}

type OtpValidation struct {
	PhoneNumber string `json:"phoneNumber"`
	UserName    string `json:"userName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	CPassword   string `json:"cPassword"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ValidName struct {
	UserName string `json:"userName"`
}

type ResendOtp struct {
	PhoneNumber string `json:"phoneNumber"`
}

type ForgotPasswordValidateOtpRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Otp         string `json:"otp"`
}

type Interests struct {
	Interest   string
	InterestID string
}

type Connection struct {
	UserID          uint
	ConnectedUserID uint
	ConnectionDate  string
}

type ValidateUserRequest struct {
	Email string `json:"email"`
}

type GoogleLoginRequest struct {
	Token string `json:"token"`
}

type ChangeUserNameRequest struct {
	UserName string `json:"userName"`
}

type ChangeEmailRequest struct {
	Email string `json:"email"`
}

type ChangeEmailVerifyOtpRequest struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}

type ChangePasswordRequest struct {
	Password string `json:"password"`
}

type ChangePhoneNumberOtpRequest struct {
	PhoneNumber string `json:"phoneNumber"`
}

type ChangePhoneNumberRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Otp         string `json:"otp"`
}

type ChangeAvatarRequest struct {
	AvatarId string `json:"avatarId"`
}

type Moderator struct {
	moderatorId string
}

type Members struct {
	memberId string
}

type CreateCommunityRequest struct {
	AdminId        string   `json:"adminId"`
	CommunityName  string   `json:"communityName"`
	Members        []string `json:"members"`
	Description    string   `json:"description"`
	JoinedType     string   `json:"joinedType"`
	CommunityImage string   `json:"communityImage"`
}

type JoinCommunityRequest struct {
	CommunityId string `json:"communityId"`
	Message     string `json:"message"`
}

type LeaveCommunityRequest struct {
	CommunityId string `json:"communityId"`
}

type AcceptJoinCommunityRequest struct {
	CommunityId    string `json:"communityId"`
	RequstedUserId string `json:"userId"`
}

type RemoveMemberRequest struct {
	CommunityId  string `json:"communityId"`
	MemberUserId string `json:"userId"`
}

type AddModeratorRequest struct {
	ModeratorUserId string `json:"userId"`
	CommunityId     string `json:""`
}

type AddMemberRequest struct {
	MemberUserId string `json:"userId"`
	CommunityId  string `json:"communityId"`
}

type ChangeCommunityJoinTypeRequest struct {
	CommunityId string `json:"communityId"`
}

type DeleteCommunityRequest struct {
	CommunityId string `json:"communityId"`
}

type ManageCommunityRequest struct {
	CommunityId string `json:"communityId"`
}

type GetUserByNameRequest struct {
	UserName string `json:"userName"`
}

type GetCommunityByIdRequest struct {
	CommunityId string `json:"communityId"`
}

type ValidateCommunityNameRequest struct {
	CommunityName string `json:"communityName"`
}

type SearchCommunityRequest struct {
	CommunityName string `json:"communityName"`
}
