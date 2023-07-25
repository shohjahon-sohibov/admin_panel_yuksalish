package handlers

import (
	"freelance/admin_panel/api/http"
	"freelance/admin_panel/models"

	"github.com/gin-gonic/gin"
)

// CreateGroup godoc
// @ID create_group
// @Router /group [POST]
// @Summary Create Group
// @Description Create Group	
// @Tags group
// @Accept json
// @Produce json
// @Param event body models.CreateGroupRequest true "Group"
// @Success 201 {object} http.Response{data=models.Group} "Module data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateGroup(c *gin.Context) {
	var group models.CreateGroupRequest

	err := c.ShouldBindJSON(&group)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	groupID, err := h.db.Group().Create(
		c.Request.Context(),
		&group,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, groupID)
}


// GetSingleGroup godoc
// @ID get_group
// @Router /group/{group_id} [GET]
// @Summary Get Single Group
// @Description Get Single Group
// @Tags group
// @Accept json
// @Produce json
// @Param group_id path string true "group_id"
// @Success 200 {object} http.Response{data=models.Group} "Group"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetSingleGroup(c *gin.Context) {
	groupId := c.Param("group_id")
	resp, err := h.db.Group().Single(
		c.Request.Context(),
		groupId,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetGroupList godoc
// @ID get_group_list
// @Router /group [GET]
// @Summary Get group List
// @Description Get group List
// @Tags group
// @Accept json
// @Produce json
// @Param GroupListReq query models.GetGroupListRequest false "GroupListReq"
// @Success 200 {object} http.Response{data=models.GetGroupListResponse} "GetGroupListResponse"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetGroupList(c *gin.Context) {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}
	resp, err := h.db.Group().List(
		c.Request.Context(),
		&models.GetGroupListRequest{
			Offset:   int64(offset),
			Limit:    int64(limit),
			BranchID: c.Query("branch_id"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateGroup godoc
// @ID update_group
// @Router /group/{group_id} [PUT]
// @Summary Update Group
// @Description Update Group
// @Tags group
// @Accept json
// @Produce json
// @Param group_id path string true "group_id"
// @Param event body models.GroupUpdate true "Group"
// @Success 200 {object} http.Response{data=models.Response} "Response"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateGroup(c *gin.Context) {
	var group models.GroupUpdate

	err := c.ShouldBindJSON(&group)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	group.ID = c.Param("group_id")

	resp, err := h.db.Group().Update(
		c.Request.Context(),
		&group,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteGroup godoc
// @ID delete_group
// @Router /group/{group_id} [DELETE]
// @Summary Delete Group
// @Description Delete Group
// @Tags group
// @Accept json
// @Produce json
// @Param group_id path string true "group_id"
// @Success 204 {object} http.Response{data=models.Response} "Response"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteGroup(c *gin.Context) {
	groupID := c.Param("group_id")

	resp, err := h.db.Group().Delete(
		c.Request.Context(),
		groupID,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
