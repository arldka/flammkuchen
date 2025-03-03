// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/arldka/flammkuchen/components/objects"
	"github.com/arldka/flammkuchen/internal/types"
	"time"
)

func HelmRelease(helmrelease types.HelmRelease, serverVersion string, objList *types.Objects) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"text-center\"><div class=\"card bg-gray-800 shadow-lg p-2 m-2 border-1 border-blue-400\"><table class=\"table w-full table-fixed\"><thead><tr class=\"grid grid-cols-5 gap-4 text-gray-400\"><th class=\"col-span-1\">Kind</th><th class=\"col-span-1\">Namespace</th><th class=\"col-span-1\">Name</th><th class=\"col-span-1\">Status</th><th class=\"col-span-1\">Last Refreshed</th></tr></thead> <tbody><tr class=\"grid grid-cols-5\"><td class=\"col-span-1\"><span class=\"px-2 py-1 inline-flex text-xs leading-5 font-semibold rounded-full bg-gray-100 text-gray-800\">HelmRelease</span></td><td class=\"col-span-1\"><div class=\"font-semibold\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(helmrelease.Namespace)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/hr.templ`, Line: 31, Col: 68}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</div></td><td class=\"col-span-1\"><div class=\"font-semibold\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(helmrelease.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/hr.templ`, Line: 34, Col: 63}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</div></td><td class=\"col-span-1\"><div class=\"text-sm\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if helmrelease.Status == "Ready" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<span class=\"px-2 py-1 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(helmrelease.Status)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/hr.templ`, Line: 40, Col: 44}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if helmrelease.Status == "Failed" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<span class=\"px-2 py-1 inline-flex text-xs leading-5 font-semibold rounded-full bg-red-100 text-red-800\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(helmrelease.Status)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/hr.templ`, Line: 44, Col: 44}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "</span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "<span class=\"px-2 py-1 inline-flex text-xs leading-5 font-semibold rounded-full bg-gray-100 text-gray-800\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(helmrelease.Status)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/hr.templ`, Line: 48, Col: 44}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "</span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "</div></td><td class=\"col-span-1\"><div class=\"font-semibold\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(time.Now().Format(time.RFC850))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/hr.templ`, Line: 54, Col: 77}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "</div></td></tr></tbody></table></div></div><div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if len(objList.Generics) > 0 {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "<span class=\"font-semibold text-2xl\">Generics</span> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				for _, generic := range objList.Generics {
					templ_7745c5c3_Err = objects.Generic(generic).Render(ctx, templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "</div><div class=\"mb-5\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if len(objList.RBACs) > 0 {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "<span class=\"font-semibold text-2xl\">RBAC</span> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				for _, rbac := range objList.RBACs {
					templ_7745c5c3_Err = objects.Generic(rbac).Render(ctx, templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "</div><div class=\"mb-5\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if len(objList.Workloads) > 0 {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "<span class=\"font-semibold text-2xl\">Workloads</span> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				for _, workload := range objList.Workloads {
					templ_7745c5c3_Err = objects.Workload(workload).Render(ctx, templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "</div><div class=\"mb-5\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if len(objList.Networkings) > 0 {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "<span class=\"font-semibold text-2xl\">Network</span> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				for _, networking := range objList.Networkings {
					templ_7745c5c3_Err = objects.Generic(networking).Render(ctx, templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "</div><div class=\"mb-5\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if len(objList.Fluxes) > 0 {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "<span class=\"font-semibold text-2xl\">FluxCD</span> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				for _, flux := range objList.Fluxes {
					templ_7745c5c3_Err = objects.Generic(flux).Render(ctx, templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "</div><div class=\"mt-4 mb-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if len(objList.CRDs) > 0 {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, "<span class=\"font-semibold text-2xl\">CRDs</span> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				for _, crd := range objList.CRDs {
					templ_7745c5c3_Err = objects.Generic(crd).Render(ctx, templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = Layout(serverVersion).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
