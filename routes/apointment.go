package routes

import (
	h "../handlers"
	m "../models"
)

//ApointmentRoutes apointments routes.
var ApointmentRoutes = m.Routes{
	m.Route{
		"ApointmentIndex",
		"GET",
		"/apointment",
		h.ApointmentIndex,
	},
	m.Route{
		"ApointmentSave",
		"POST",
		"/apointment",
		h.ApointmentSave,
	},
	m.Route{
		"ApointmentFindByID",
		"GET",
		"/apointment/{apointmentID}",
		h.ApointmentFindByID,
	},
	m.Route{
		"ApointmentUpdate",
		"PUT",
		"/apointment/{apointmentID}",
		h.ApointmentUpdate,
	},
	m.Route{
		"ApointmentDelete",
		"DELETE",
		"/apointment/{apointmentID}",
		h.ApointmentDelete,
	},
}
