package userRequest

type AddFriendRequest struct {
	DstId int `json:"dst_id" binding:"required"`
}

type LoadFriendRequest struct {
}

type CreateCommunityRequest struct {
}

type LoadCommunityRequest struct {
}

type JoinCommunityRequest struct {
}
