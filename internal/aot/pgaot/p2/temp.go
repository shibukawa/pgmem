package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetTempNamespaceProcNumber(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	v4 = F_get_namespace_name(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v4 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(-1)
L4:
	;
	goto L5
L5:
	;
	v12 = int32(8)
	v13 = int32(_a_F_GetTempNamespaceProcNumber_0)
	goto L9
L6:
	;
	F_pfree(m, v4)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L53
	}
L7:
	;
	if v51-v52 != 0 {
		goto L20
	} else {
		goto L21
	}
L9:
	;
	goto L10
L10:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v20 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v21 = v4
	v22 = v13
	v23 = v12
	v24 = v20
	goto L15
L12:
	;
	v47 = v13
	v51 = int32(0)
	goto L13
L13:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	goto L7
L14:
	;
	v47 = v42
	v51 = v44
	goto L13
L15:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if base.B2i32(v24 != v26)|base.B2i32(v26 == int32(0)) != 0 {
		v42 = v22
		v44 = v24
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v42 = v36
	v44 = int32(0)
	goto L14
L17:
	;
	v32 = v23 - int32(1)
	if v32 == int32(0) {
		v42 = v22
		v44 = v24
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v35 = int32(1)
	v36 = v22 + v35
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v37 != 0 {
		v21 = v21 + v35
		v22 = v36
		v23 = v32
		v24 = v37
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v61 = int32(14)
	v62 = int32(_a_F_GetTempNamespaceProcNumber_1)
	goto L25
L21:
	;
	v109 = v12
	goto L22
L22:
	;
	v115 = v4 + v109
	goto L38
L23:
	;
	if v100-v101 != 0 {
		v161 = int32(-1)
		goto L6
	} else {
		goto L36
	}
L25:
	;
	goto L26
L26:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v69 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v70 = v4
	v71 = v62
	v72 = v61
	v73 = v69
	goto L31
L28:
	;
	v96 = v62
	v100 = int32(0)
	goto L29
L29:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	goto L23
L30:
	;
	v96 = v91
	v100 = v93
	goto L29
L31:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if base.B2i32(v73 != v75)|base.B2i32(v75 == int32(0)) != 0 {
		v91 = v71
		v93 = v73
		goto L30
	} else {
		goto L33
	}
L32:
	;
	v91 = v85
	v93 = int32(0)
	goto L30
L33:
	;
	v81 = v72 - int32(1)
	if v81 == int32(0) {
		v91 = v71
		v93 = v73
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v84 = int32(1)
	v85 = v71 + v84
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v86 != 0 {
		v70 = v70 + v84
		v71 = v85
		v72 = v81
		v73 = v86
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v109 = v61
	goto L22
L37:
	;
	v161 = v159
	goto L6
L38:
	;
	v120 = v115 + int32(1)
	v121 = int32(*(*int8)(unsafe.Add(mBase, uint32(v115))))
	v122 = F___isspace(m, v121)
	mBase = m.M
	if v122 != 0 {
		v115 = v120
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v123 = int32(1)
	switch v121&int32(255) - int32(43) {
	case 0:
		v129 = v123
		goto L42
	default:
		v131 = v121
		v132 = v115
		v133 = v123
		goto L41
	case 2:
		goto L43
	}
L40:
	;
	goto L39
L41:
	;
	v134 = int32(0)
	v136 = v131 - int32(48)
	if base.Ui32(v136) <= base.Ui32(int32(9)) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v130 = int32(*(*int8)(unsafe.Add(mBase, uint32(v120))))
	v131 = v130
	v132 = v120
	v133 = v129
	goto L41
L43:
	;
	v129 = int32(0)
	goto L42
L44:
	;
	v139 = v134
	v140 = v136
	v141 = v132
	goto L47
L45:
	;
	v153 = v134
	goto L46
L46:
	;
	if v133 != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v143 = int32(10)
	v145 = v139*v143 - v140
	v146 = int32(*(*int8)(unsafe.Add(mBase, uint32(v141)+1)))
	v150 = v146 - int32(48)
	if base.Ui32(v150) < base.Ui32(v143) {
		v139 = v145
		v140 = v150
		v141 = v141 + int32(1)
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v153 = v145
	goto L46
L49:
	;
	goto L48
L50:
	;
	v159 = int32(0) - v153
	goto L52
L51:
	;
	v159 = v153
	goto L52
L52:
	;
	goto L37
L53:
	;
	return v161
}
func F_assign_temp_tablespaces(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	if l1 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, _c_F_assign_temp_tablespaces[0])) = v5
		*(*int32)(unsafe.Add(mBase, _c_F_assign_temp_tablespaces[1])) = l1 + int32(4)
		if int32(2) <= v5 {
			v18 = F_pg_prng_uint64_range(m, int32(_a_F_assign_temp_tablespaces_0), int64(0), base.I64_extend_i32_u(v5-int32(1)))
			mBase = m.M
			v21 = base.I32_wrap_i64(v18)
		} else {
			v21 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, _c_F_assign_temp_tablespaces[2])) = v21
		return
	} else {
		v23 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_assign_temp_tablespaces[0])) = v23
		*(*int32)(unsafe.Add(mBase, _c_F_assign_temp_tablespaces[1])) = v23
		*(*int32)(unsafe.Add(mBase, _c_F_assign_temp_tablespaces[2])) = int32(0)
		return
	}
}
func F_createTempGistContext(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_createTempGistContext[0]))
	v7 = F_AllocSetContextCreateInternal(m, v2, int32(_a_F_createTempGistContext_0), int32(0), int32(_a_F_createTempGistContext_1), int32(_a_F_createTempGistContext_2))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
