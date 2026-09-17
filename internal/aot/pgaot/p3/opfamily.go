package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_opfamily_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	F_DeconstructQualifiedName(m, l1, v10+int32(28), v10+int32(24))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	F_errmsg_internal(m, int32(_a_F_get_opfamily_oid_0), v10)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L42
	}
L4:
	;
	if l2|v93 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L5:
	;
	v22 = F_LookupExplicitNamespace(m, v20, l2)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L11
	}
L8:
	;
	if v22 == int32(0) {
		v93 = int32(0)
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v28 = F_SearchSysCache3(m, int32(41), l0, v27, v22)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v93 = v28
	goto L4
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_get_opfamily_oid[0]))
	if v35 == int32(0) {
		v75 = int32(0)
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v80 = int32(0)
	if v75 == v80 {
		v93 = v80
		goto L4
	} else {
		goto L25
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if int32(0) < v38 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v46 = int32(0)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v75 = int32(0)
	goto L12
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v46<<(uint(int32(2))%32))))
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_get_opfamily_oid[1]))
	if v52 != v54 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L16
L19:
	;
	v58 = F_GetSysCacheOid(m, int32(41), l0, v30, v52, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v62 = v46 + int32(1)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v62 < v63 {
		v46 = v62
		goto L17
	} else {
		goto L24
	}
L22:
	;
	if v58 != 0 {
		v75 = v58
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
	v84 = F_SearchSysCache1(m, int32(42), v75)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v93 = v84
	goto L4
L27:
	;
	v98 = F_SearchSysCache1(m, int32(2), l0)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v93 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v98 == int32(0) {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v109 = F_NameListToString(m, l1)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v98)+16))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v111 + v112 + int32(4)
	F_errmsg(m, int32(_a_F_get_opfamily_oid_1), v10+int32(16))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_get_opfamily_oid_2), int32(126), int32(_a_F_get_opfamily_oid_3))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
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
	m.G0 = v10 + int32(32)
	return v138
L38:
	;
	v138 = int32(0)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+22)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v131+v132)))
	F_ReleaseCatCache(m, v93)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v138 = v134
	goto L37
L42:
	;
	F_errfinish(m, int32(_a_F_get_opfamily_oid_2), int32(121), int32(_a_F_get_opfamily_oid_3))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
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
