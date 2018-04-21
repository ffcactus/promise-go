package base

// TaskController is the controller that handle the action which is asychronously.
// We create context, strategy for these kinds of actions.
type TaskController struct {
	beego.Controller
}

func (c *TaskController) Post() {
	var (
		messages   []Message
		action     = c.Ctx.Input.Param(":action")
		id         = c.Ctx.Input.Param(":id")
		actionInfo = c.TemplateImpl.GetActionInfo()
		service    ActionServiceInterface
		request    UpdateActionRequestInterface		
	)

	// Find the matching ActionInfo.s
	for _, v := range actionInfo {
		if strings.ToLower(action) == strings.ToLower(v.Name) {
			service = v.Service
			request = v.Request()
		}
	}
	if service == nil {
		messages = append(messages, NewMessageInvalidRequest())
		log.WithFields(log.Fields{
			"resource": c.TemplateImpl.GetResourceName(),
			"action":   action,
			"id":       id,
			"message":  messages[0].ID,
		}).Warn("Perform action failed, unknown action.")
		c.Data["json"] = &messages
		c.Ctx.Output.SetStatus(messages[0].StatusCode)
		c.ServeJSON()
		return
	}
}