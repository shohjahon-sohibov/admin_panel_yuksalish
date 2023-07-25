package handlers

import (
	"freelance/admin_panel/api/http"
	"freelance/admin_panel/models"

	"github.com/gin-gonic/gin"
)

// CreateStudent godoc
// @ID create_student
// @Router /student [POST]
// @Summary Create Student
// @Description Create Student	
// @Tags student
// @Accept json
// @Produce json
// @Param event body models.CreateStudentRequest true "Student"
// @Success 201 {object} http.Response{data=models.Branch} "Module data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateStudent(c *gin.Context) {
	var student models.CreateStudentRequest

	err := c.ShouldBindJSON(&student)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.db.Student().Create(
		c.Request.Context(),
		&student,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetSingleStudent godoc
// @ID get_student
// @Router /student/{student_id} [GET]
// @Summary Get Single Student
// @Description Get Single Student
// @Tags student
// @Accept json
// @Produce json
// @Param student_id path string true "student_id"
// @Success 200 {object} http.Response{data=models.Student} "Student"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetSingleStudent(c *gin.Context) {
	studentId := c.Param("student_id")
	resp, err := h.db.Student().Single(
		c.Request.Context(),
		studentId,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetStudentList godoc
// @ID get_student_list
// @Router /student [GET]
// @Summary Get student List
// @Description Get student List
// @Tags student
// @Accept json
// @Produce json
// @Param StudentListReq query models.GetStudentListRequest false "GetStudentListReq"
// @Success 200 {object} http.Response{data=models.GetStudentListResponse} "GetStudentListResponse"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetStudentList(c *gin.Context) {
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
	resp, err := h.db.Student().List(
		c.Request.Context(),
		&models.GetStudentListRequest{
			Offset:   int64(offset),
			Limit:    int64(limit),
			BranchID: c.Query("branch_id"),
			GroupID: c.Query("group_id"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateStudent godoc
// @ID update_student
// @Router /student/{student_id} [PUT]
// @Summary Update Student
// @Description Update Student
// @Tags student
// @Accept json
// @Produce json
// @Param student_id path string true "student_id"
// @Param event body models.StudentUpdate true "Student"
// @Success 200 {object} http.Response{data=models.Response} "Response"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateStudent(c *gin.Context) {
	var student models.StudentUpdate

	err := c.ShouldBindJSON(&student)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	student.ID = c.Param("student_id")

	resp, err := h.db.Student().Update(
		c.Request.Context(),
		&student,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteStudent godoc
// @ID delete_student
// @Router /student/{student_id} [DELETE]
// @Summary Delete Student
// @Description Delete Student
// @Tags student
// @Accept json
// @Produce json
// @Param student_id path string true "student_id"
// @Success 204 {object} http.Response{data=models.Response} "Response"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteStudent(c *gin.Context) {
	studentID := c.Param("student_id")

	resp, err := h.db.Student().Delete(
		c.Request.Context(),
		studentID,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
