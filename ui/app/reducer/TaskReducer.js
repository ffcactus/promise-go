import * as types from '../actions/types';

const defaultTaskState = {
  current: null,
  taskList: [],
  taskLoading: false,
  taskListLoading: false,
};

const task = (state = defaultTaskState, action) => {
  let newState;
  switch (action.type) {
    case types.TASK_GET_BEGIN:
      return Object.assign({}, state, {
        current: null,
        taskLoading: true,
      });
    case types.TASK_GET_SUCCESS:
      return Object.assign({}, state, {
        current: action.info,
        taskLoading: false,
      });
    case types.TASK_GET_FAILURE:
      return Object.assign({}, state, {
        current: null,
        taskLoading: false,
      });
    case types.TASK_LIST_GET_BEGIN:
      return Object.assign({}, state, {
        taskList: [],
        current: null,
        taskListLoading: true
      });
    case types.TASK_LIST_GET_SUCCESS:
      return Object.assign({}, state, {
        taskList: action.info.Members.map((each) => {
          return {
            Uri: each.Uri,
            Name: each.Name,
            Description: each.Description,
            ExecutionState: each.ExecutionState,
            Percentage: each.Percentage,
            ExecutionResult: each.ExecutionResult,
          };
        }),
        taskListLoading: false,
      });
    case types.SERVER_LIST_LOAD_FAILURE:
      return Object.assign({}, state, {
        taskList: [],
        current: null,
        taskListLoading: false,
      });
    case types.TASK_CREATE:
      newState = Object.assign({}, state, {
        current: state.current,
        taskListLoading: state.taskListLoading
      });
      newState.taskList.push({
        Uri: action.info.Uri,
        Name: action.info.Name,
        Description: action.info.Description,
        CreatedAt: action.info.CreatedAt,
        UpdatedAt: action.info.UpdatedAt,
        CreatedByName: action.info.CreatedByName,
        CreatedByUri: action.info.CreatedByUri,
        TargetName: action.info.TargetName,
        TargetUri: action.info.TargetUri,
        CurrentStep: action.info.CurrentStep,
        ExecutionState: action.info.ExecutionState,
        Percentage: action.info.Percentage,
        ExecutionResult: action.info.ExecutionResult,
      });
      return newState;
    case types.TASK_UPDATE:
      newState = Object.assign({}, state, {});
      newState.taskList.map((each) => {
        if (action.info.Uri === each.Uri) {
          each.Name = action.info.Name;
          each.Uri = action.info.Uri;
          each.Description = action.info.Description;
          each.CreatedAt = action.info.CreatedAt;
          each.UpdatedAt = action.info.UpdatedAt;
          each.CreatedByName = action.info.CreatedByName;
          each.CreatedByUri = action.info.CreatedByUri;
          each.TargetName = action.info.TargetName;
          each.TargetUri = action.info.TargetUri;
          each.CurrentStep = action.info.CurrentStep;
          each.ExecutionState = action.info.ExecutionState;
          each.Percentage = action.info.Percentage;
          each.ExecutionResult = action.info.ExecutionResult;
        }
      });
      if (newState.current !== null && newState.current.Id === action.info.Id) {
        newState.current = action.info;
      }
      return newState;
    default:
      return state;
  }
};

export default task;
