package typereq

// reference: https://www.mongodb.com/docs/manual/reference/privilege-actions/#std-label-security-user-actions

type TypeAction string

// Query and Write Actions
const (
	ActionFind                     TypeAction = "find"
	ActionInsert                   TypeAction = "insert"
	ActionRemove                   TypeAction = "remove"
	ActionUpdate                   TypeAction = "update"
	ActionBypassDocumentValidation TypeAction = "bypassDocumentValidation"
	ActionUseUUID                  TypeAction = "useUUID"
)

// Database Management Actions
const (
	ActionChangeCustomData               TypeAction = "changeCustomData"
	ActionChangeOwnCustomData            TypeAction = "changeOwnCustomData"
	ActionChangeOwnPassword              TypeAction = "changeOwnPassword"
	ActionChangePassword                 TypeAction = "changePassword"
	ActionCreateCollection               TypeAction = "createCollection"
	ActionCreateIndex                    TypeAction = "createIndex"
	ActionCreateRole                     TypeAction = "createRole"
	ActionCreateUser                     TypeAction = "createUser"
	ActionDropCollection                 TypeAction = "dropCollection"
	ActionDropRole                       TypeAction = "dropRole"
	ActionDropUser                       TypeAction = "dropUser"
	ActionEnableProfiler                 TypeAction = "enableProfiler"
	ActionGrantRole                      TypeAction = "grantRole"
	ActionKillCursors                    TypeAction = "killCursors"
	ActionKillAnyCursor                  TypeAction = "killAnyCursor"
	ActionPlanCacheIndexFilter           TypeAction = "planCacheIndexFilter"
	ActionRevokeRole                     TypeAction = "revokeRole"
	ActionSetAuthenticationRestriction   TypeAction = "setAuthenticationRestriction"
	ActionSetFeatureCompatibilityVersion TypeAction = "setFeatureCompatibilityVersion"
	ActionUnlock                         TypeAction = "unlock"
	ActionViewRole                       TypeAction = "viewRole"
	ActionViewUser                       TypeAction = "viewUser"
)

// TODO
