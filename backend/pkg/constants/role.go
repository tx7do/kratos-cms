package constants

const (
	// RoleCodeSpilt 角色代码分隔符
	RoleCodeSpilt = ":"

	// TemplateRoleCodePrefix 模板角色代码前缀
	TemplateRoleCodePrefix = "template" + RoleCodeSpilt
	// PlatformRoleCodePrefix 平台角色代码前缀
	PlatformRoleCodePrefix = "platform" + RoleCodeSpilt
	// TenantRoleCodePrefix 租户角色代码前缀
	TenantRoleCodePrefix = "tenant" + RoleCodeSpilt

	// PlatformAdminRoleCode 平台管理员角色代码
	PlatformAdminRoleCode = PlatformRoleCodePrefix + "admin"
	// TenantAdminRoleCode 租户管理员角色代码
	TenantAdminRoleCode = TenantRoleCodePrefix + "manager"
	// TenantUserRoleCode 租户普通用户角色代码(C 端注册默认角色)
	TenantUserRoleCode = TenantRoleCodePrefix + "user"
	// TenantAdminTemplateRoleCode 租户管理员模板角色代码
	TenantAdminTemplateRoleCode = TemplateRoleCodePrefix + TenantAdminRoleCode

	// DefaultPlatformAdminRoleName 平台管理员角色默认名称
	DefaultPlatformAdminRoleName = "平台管理员"
	// DefaultTenantManagerRoleName 租户管理员角色默认名称
	DefaultTenantManagerRoleName = "租户管理员"
	// DefaultTenantUserRoleName 租户普通用户角色默认名称
	DefaultTenantUserRoleName = "普通用户"
)

func HasRoleCodePrefix(roleCode string, prefix string) bool {
	return len(roleCode) >= len(prefix) && roleCode[0:len(prefix)] == prefix
}

func IsPlatformRoleCode(roleCode string) bool {
	return HasRoleCodePrefix(roleCode, PlatformRoleCodePrefix)
}

func IsTenantRoleCode(roleCode string) bool {
	return HasRoleCodePrefix(roleCode, TenantRoleCodePrefix)
}

func IsTemplateRoleCode(roleCode string) bool {
	return HasRoleCodePrefix(roleCode, TemplateRoleCodePrefix)
}

func ExtractRoleCodeFromTemplate(roleCode string) string {
	if IsTemplateRoleCode(roleCode) {
		return roleCode[len(TemplateRoleCodePrefix):]
	}
	return roleCode
}
