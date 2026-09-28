package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_opfamily_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	F_DeconstructQualifiedName(m, l1, v12+int32(28), v12+int32(24))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg_internal(m, int32(_a_F_get_opfamily_oid_0), v12)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L42
	}
L4:
	;
	if l2|v109 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L5:
	;
	v24 = F_LookupExplicitNamespace(m, v22, l2)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L11
	}
L8:
	;
	if v24 == int32(0) {
		v109 = int32(0)
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v30 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+24)))
	v32 = F_SearchSysCache3(m, int32(41), base.I64_extend_i32_u(l0), v30, base.I64_extend_i32_u(v24))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v109 = v32
	goto L4
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_get_opfamily_oid[0]))
	if v39 == int32(0) {
		v86 = int32(0)
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v93 = int32(0)
	if v86 == v93 {
		v109 = v93
		goto L4
	} else {
		goto L25
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if int32(0) < v42 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v52 = int32(0)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v86 = int32(0)
	goto L12
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v52<<(uint(int32(2))%32))))
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_get_opfamily_oid[1]))
	if v60 != v62 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L16
L19:
	;
	v67 = F_GetSysCacheOid(m, int32(41), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(v34), base.I64_extend_i32_u(v60), int64(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v71 = v52 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v71 < v72 {
		v52 = v71
		goto L17
	} else {
		goto L24
	}
L22:
	;
	if v67 != 0 {
		v86 = v67
		goto L12
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L18
L25:
	;
	v98 = F_SearchSysCache1(m, int32(42), base.I64_extend_i32_u(v86))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v109 = v98
	goto L4
L27:
	;
	v115 = F_SearchSysCache1(m, int32(2), base.I64_extend_i32_u(l0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v109 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v115 == int32(0) {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v126 = F_NameListToString(m, l1)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v128 + v129 + int32(4)
	F_errmsg(m, int32(_a_F_get_opfamily_oid_1), v12+int32(16))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_get_opfamily_oid_2), int32(126), int32(_a_F_get_opfamily_oid_3))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	m.G0 = v12 + int32(32)
	return v155
L38:
	;
	v155 = int32(0)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+22)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v148+v149)))
	F_ReleaseCatCache(m, v109)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v155 = v151
	goto L37
L42:
	;
	F_errfinish(m, int32(_a_F_get_opfamily_oid_2), int32(121), int32(_a_F_get_opfamily_oid_3))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
