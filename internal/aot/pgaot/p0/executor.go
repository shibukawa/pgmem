package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateExecutorState(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v31 int32
	_ = v31
	v5 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v10 = F_AllocSetContextCreateInternal(m, v5, int32(354525), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = int32(4520272)
		v15 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v10
		v19 = F_palloc0(m, int32(200))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v19)+188)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19))) = int64(4294967685)
			*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+60)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = v21
			v31 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v31
			*(*int32)(unsafe.Add(mBase, uint32(v19)+100)) = v10
			*(*int64)(unsafe.Add(mBase, uint32(v19)+176)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+164)) = v21
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+160)) = uint8(v31)
			*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v31
			*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v31
			*(*int64)(unsafe.Add(mBase, uint32(v19)+68)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+88)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v31
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+136)) = uint8(v31)
			*(*int64)(unsafe.Add(mBase, uint32(v19)+128)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+120)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+112)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+140)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+148)) = v21
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v15
			return v19
		}
	}
}
func F_ExecutorRun(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v93 int64
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
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
	var v115 int32
	_ = v115
	var v120 int64
	_ = v120
	var v125 int64
	_ = v125
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v173 int32
	_ = v173
	v15 = *(*int32)(unsafe.Add(mBase, _consts[325]))
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.T0[v15].(func(*base.Module, int32, int32, int64))(m, l0, l1, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v18 = int32(4520272)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+100))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	return
L6:
	;
	F_InstrStartNode(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+112)) = int64(0)
	if v28 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L8
L10:
	;
	if l1 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+24)))
	if v35 != int32(1) {
		v43 = int32(0)
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	m.T0[v39].(func(*base.Module, int32, int32, int32))(m, v27, v28, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v43 = int32(1)
	goto L10
L16:
	;
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v21)+120))
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v21)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+120)) = v162 + v163
	if v43 != 0 {
		goto L60
	} else {
		goto L61
	}
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = l1
	if l2 == int64(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v93 = int64(0)
	goto L26
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+29)))
	v63 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v63)
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+160)) = uint8(v62)
	if v62 != v63 {
		v78 = int32(0)
		goto L18
	} else {
		goto L24
	}
L20:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v51&int32(1) == int32(0) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v56 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v56)
	v58 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+160)) = uint8(v58)
	v78 = v58
	goto L18
L23:
	;
	goto L22
L24:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+72)) = v72 + int32(1)
	goto L25
L25:
	;
	v78 = int32(1)
	goto L18
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v47)+152))
	if v94 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+128)))
	if v132&int32(8) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	F_MemoryContextReset(m, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
	if v98 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	F_ExecReScan(m, v46)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v102 = m.T0[v101].(func(*base.Module, int32) int32)(m, v46)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	goto L27
L37:
	;
	if v102 == int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+4)))
	if v106&int32(2) != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v109 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v110 = F_ExecFilterJunk(m, v109, v102)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	v112 = v102
	goto L42
L42:
	;
	if v43 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v112 = v110
	goto L42
L44:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v114 = m.T0[v113].(func(*base.Module, int32, int32) int32)(m, v112, v27)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if base.B2i32(v28 != int32(1)) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	if v114 == int32(0) {
		goto L36
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v47)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v47)+112)) = v120 + int64(1)
	goto L51
L50:
	;
	goto L51
L51:
	;
	v125 = v93 + int64(1)
	if l2 == int64(0) {
		v93 = v125
		goto L26
	} else {
		goto L52
	}
L52:
	;
	if l2 != v125 {
		v93 = v125
		goto L26
	} else {
		goto L53
	}
L53:
	;
	goto L36
L54:
	;
	v138 = F_ExecShutdownNode_walker(m, v46, int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v78 == int32(0) {
		goto L16
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v144)+72)) = v145 - int32(1)
	goto L59
L59:
	;
	goto L16
L60:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	m.T0[v166].(func(*base.Module, int32))(m, v27)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v169 != 0 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	goto L62
L64:
	;
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v21)+112))
	F_InstrStopNode(m, v169, base.F64_convert_i64_u(v170))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v19
	return
L67:
	;
	goto L66
}
func F_FreeExecutorState(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v52 != 0 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	if v12 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v13 = int32(4520272)
	v14 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
	v19 = v12
	goto L9
L7:
	;
	goto L8
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	F_MemoryContextDelete(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L15
	}
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	m.T0[v25].(func(*base.Module, int32))(m, v24)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v14
	goto L8
L11:
	;
	return
L12:
	;
	F_pfree(m, v19)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	if v30 != 0 {
		v19 = v30
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	if v40 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+140))
	v42 = F_list_delete_ptr(m, v41, v11)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_pfree(m, v11)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L11
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+140)) = v42
	goto L18
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v47 != 0 {
		v7 = v47
		goto L4
	} else {
		goto L21
	}
L21:
	;
	goto L5
L22:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, _consts[327])))
	if v54 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v63 != 0 {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[328]))
	m.T0[v56].(func(*base.Module, int32))(m, v52)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L11
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	F_pfree(m, v52)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L11
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = int32(0)
	goto L24
L30:
	;
	F_DestroyPartitionDirectory(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L11
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	F_MemoryContextDelete(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L11
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = int32(0)
	goto L32
L34:
	;
	return
}
