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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
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
	v13 = int32(508713)
	goto L9
L6:
	;
	F_pfree(m, v4)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L55
	}
L7:
	;
	if v50-v51 != 0 {
		goto L21
	} else {
		goto L22
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
	v46 = v13
	v50 = int32(0)
	goto L13
L13:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	goto L7
L14:
	;
	v46 = v41
	v50 = v43
	goto L13
L15:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v24 != v26 {
		v41 = v22
		v43 = v24
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v41 = v35
	v43 = int32(0)
	goto L14
L17:
	;
	if v26 == int32(0) {
		v41 = v22
		v43 = v24
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v31 = v23 - int32(1)
	if v31 == int32(0) {
		v41 = v22
		v43 = v24
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v34 = int32(1)
	v35 = v22 + v34
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v36 != 0 {
		v21 = v21 + v34
		v22 = v35
		v23 = v31
		v24 = v36
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	v60 = int32(14)
	v61 = int32(508698)
	goto L26
L22:
	;
	v107 = v12
	goto L23
L23:
	;
	v113 = v4 + v107
	goto L40
L24:
	;
	if v98-v99 != 0 {
		v159 = int32(-1)
		goto L6
	} else {
		goto L38
	}
L26:
	;
	goto L27
L27:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v68 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v69 = v4
	v70 = v61
	v71 = v60
	v72 = v68
	goto L32
L29:
	;
	v94 = v61
	v98 = int32(0)
	goto L30
L30:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	goto L24
L31:
	;
	v94 = v89
	v98 = v91
	goto L30
L32:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v72 != v74 {
		v89 = v70
		v91 = v72
		goto L31
	} else {
		goto L34
	}
L33:
	;
	v89 = v83
	v91 = int32(0)
	goto L31
L34:
	;
	if v74 == int32(0) {
		v89 = v70
		v91 = v72
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v79 = v71 - int32(1)
	if v79 == int32(0) {
		v89 = v70
		v91 = v72
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v82 = int32(1)
	v83 = v70 + v82
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)))
	if v84 != 0 {
		v69 = v69 + v82
		v70 = v83
		v71 = v79
		v72 = v84
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v107 = v60
	goto L23
L39:
	;
	v159 = v157
	goto L6
L40:
	;
	v118 = v113 + int32(1)
	v119 = int32(*(*int8)(unsafe.Add(mBase, uint32(v113))))
	v120 = F___isspace(m, v119)
	mBase = m.M
	if v120 != 0 {
		v113 = v118
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v121 = int32(1)
	switch v119&int32(255) - int32(43) {
	case 0:
		v127 = v121
		goto L44
	default:
		v129 = v119
		v130 = v113
		v131 = v121
		goto L43
	case 2:
		goto L45
	}
L42:
	;
	goto L41
L43:
	;
	v132 = int32(0)
	v134 = v129 - int32(48)
	if base.Ui32(v134) <= base.Ui32(int32(9)) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v128 = int32(*(*int8)(unsafe.Add(mBase, uint32(v118))))
	v129 = v128
	v130 = v118
	v131 = v127
	goto L43
L45:
	;
	v127 = int32(0)
	goto L44
L46:
	;
	v137 = v132
	v138 = v134
	v139 = v130
	goto L49
L47:
	;
	v151 = v132
	goto L48
L48:
	;
	if v131 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v141 = int32(10)
	v143 = v137*v141 - v138
	v144 = int32(*(*int8)(unsafe.Add(mBase, uint32(v139)+1)))
	v148 = v144 - int32(48)
	if base.Ui32(v148) < base.Ui32(v141) {
		v137 = v143
		v138 = v148
		v139 = v139 + int32(1)
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v151 = v143
	goto L48
L51:
	;
	goto L50
L52:
	;
	v157 = int32(0) - v151
	goto L54
L53:
	;
	v157 = v151
	goto L54
L54:
	;
	goto L39
L55:
	;
	return v159
}
func F_assign_temp_tablespaces(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	if l1 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, _consts[408])) = v5
		*(*int32)(unsafe.Add(mBase, _consts[409])) = l1 + int32(4)
		if int32(2) <= v5 {
			v19 = F_pg_prng_uint64_range(m, int32(4603776), int64(0), base.I64_extend_i32_u(v5-int32(1)))
			mBase = m.M
			v21 = base.I32_wrap_i64(v19)
		} else {
			v21 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, _consts[410])) = v21
		return
	} else {
		v23 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[408])) = v23
		*(*int32)(unsafe.Add(mBase, _consts[409])) = v23
		*(*int32)(unsafe.Add(mBase, _consts[410])) = v23
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
	v2 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v7 = F_AllocSetContextCreateInternal(m, v2, int32(60299), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
