package handlers

import (
	"freelance/admin_panel/api/http"
	"freelance/admin_panel/models"

	"github.com/gin-gonic/gin"
)

// CreateBranch godoc
// @ID create_branch
// @Router /branch [POST]
// @Summary Create Branch
// @Description Create Branch	
// @Tags branch
// @Accept json
// @Produce json
// @Param event body models.CreateBranchRequest true "Branch"
// @Success 201 {object} http.Response{data=models.Branch} "Module data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateBranch(c *gin.Context) {
	var branch models.CreateBranchRequest

	err := c.ShouldBindJSON(&branch)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.db.Branch().Create(
		c.Request.Context(),
		&branch,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetSingleBranch godoc
// @ID get_branch
// @Router /branch/{branch_id} [GET]
// @Summary Get Single Branch
// @Description Get Single Branch
// @Tags branch
// @Accept json
// @Produce json
// @Param branch_id path string true "branch_id"
// @Success 200 {object} http.Response{data=models.Branch} "Branch"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetSingleBranch(c *gin.Context) {
	branchId := c.Param("branch_id")
	resp, err := h.db.Branch().Single(
		c.Request.Context(),
		branchId,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetBranchList godoc
// @ID get_branch_list
// @Router /branch [GET]
// @Summary Get branch List
// @Description Get branch List
// @Tags branch
// @Accept json
// @Produce json
// @Param BranchListReq query models.GetBranchListRequest false "BranchListReq"
// @Success 200 {object} http.Response{data=models.GetBranchListResponse} "GetBranchListResponse"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetBranchList(c *gin.Context) {
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
	resp, err := h.db.Branch().List(
		c.Request.Context(),
		&models.GetBranchListRequest{
			Offset:   int64(offset),
			Limit:    int64(limit),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateBranch godoc
// @ID update_branch
// @Router /branch/{branch_id} [PUT]
// @Summary Update Branch
// @Description Update Branch
// @Tags branch
// @Accept json
// @Produce json
// @Param branch_id path string true "branch_id"
// @Param event body models.BranchUpdate true "Branch"
// @Success 200 {object} http.Response{data=models.Response} "Response"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateBranch(c *gin.Context) {
	var branch models.BranchUpdate

	err := c.ShouldBindJSON(&branch)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	branch.ID = c.Param("branch_id")

	resp, err := h.db.Branch().Update(
		c.Request.Context(),
		&branch,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteBranch godoc
// @ID delete_branch
// @Router /branch/{branch_id} [DELETE]
// @Summary Delete Branch
// @Description Delete Branch
// @Tags branch
// @Accept json
// @Produce json
// @Param branch_id path string true "branch_id"
// @Success 204 {object} http.Response{data=models.Response} "Response"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteBranch(c *gin.Context) {
	branchID := c.Param("branch_id")

	resp, err := h.db.Branch().Delete(
		c.Request.Context(),
		branchID,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
