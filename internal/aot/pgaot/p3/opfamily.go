package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_opfamily_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	F_DeconstructQualifiedName(m, l1, v11+int32(28), v11+int32(24))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if l2 != 0 {
		goto L27
	} else {
		goto L28
	}
L4:
	;
	v23 = F_LookupExplicitNamespace(m, v21, l2)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L10
	}
L7:
	;
	if v23 == int32(0) {
		v98 = int32(0)
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v29 = F_SearchSysCache3(m, int32(41), l0, v28, v23)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v98 = v29
	goto L3
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[437]))
	if v36 == int32(0) {
		v79 = int32(0)
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v84 = int32(0)
	if v79 == v84 {
		v98 = v84
		goto L3
	} else {
		goto L24
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if int32(0) < v39 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v47 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	v79 = int32(0)
	goto L11
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v47<<(uint(int32(2))%32))))
	v56 = *(*int32)(unsafe.Add(mBase, _consts[135]))
	if v54 != v56 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	v60 = F_GetSysCacheOid(m, int32(41), l0, v31, v54, int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v64 = v47 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v64 < v65 {
		v47 = v64
		goto L16
	} else {
		goto L23
	}
L21:
	;
	if v60 != 0 {
		v79 = v60
		goto L11
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	goto L17
L24:
	;
	v88 = F_SearchSysCache1(m, int32(42), v79)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v98 = v88
	goto L3
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(53202), v11)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L42
	}
L27:
	;
	if v98 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L28:
	;
	if v98 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v100 = F_SearchSysCache1(m, int32(2), l0)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v100 == int32(0) {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v111 = F_NameListToString(m, l1)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v100)+16))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v114 + v113 + int32(4)
	F_errmsg(m, int32(722303), v11+int32(16))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(494550), int32(126), int32(233177))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
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
	m.G0 = v11 + int32(32)
	return v139
L38:
	;
	v139 = int32(0)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v98)+16))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+22)))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v133+v134)))
	F_ReleaseCatCache(m, v98)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v139 = v136
	goto L37
L42:
	;
	F_errfinish(m, int32(494550), int32(121), int32(233177))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
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
