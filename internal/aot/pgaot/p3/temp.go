package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SetTempTablespaces(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v44 int64
	_ = v44
	var v49 int64
	_ = v49
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	*(*int32)(unsafe.Add(mBase, _consts[815])) = l1
	*(*int32)(unsafe.Add(mBase, _consts[817])) = l0
	if int32(2) <= l1 {
		v11 = int32(4572944)
		v12 = int64(0)
		v15 = base.I64_extend_i32_u(l1 - int32(1))
		if base.Ui64(v15) <= base.Ui64(v12) {
			v62 = v12
		} else {
			v22 = v15 - v12
			v24 = *(*int64)(unsafe.Add(mBase, _consts[818]))
			v25 = *(*int64)(unsafe.Add(mBase, _consts[819]))
			v28 = v25
			v30 = v24
			for {
				v34 = v28 ^ v30
				v36 = base.I64_rotl(v34, int64(37))
				v44 = v34 ^ (v34<<(uint(int64(16))%64) ^ base.I64_rotl(v28, int64(24)))
				v49 = int64(base.Ui64(base.I64_rotl(v28*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v22)) % 64))
				if base.Ui64(v22) < base.Ui64(v49) {
					v28 = v44
					v30 = v36
					continue
				} else {
					break
				}
				break
			}
			*(*int64)(unsafe.Add(mBase, _consts[818])) = v36
			*(*int64)(unsafe.Add(mBase, _consts[819])) = v44
			v62 = v12 + v49
		}
		v64 = base.I32_wrap_i64(v62)
	} else {
		v64 = int32(0)
	}
	*(*int32)(unsafe.Add(mBase, _consts[816])) = v64
	return
}
func F_checkTempNamespaceStatus(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	v2 = int32(0)
	v5 = F_get_namespace_name(m, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v200
L2:
	;
	return int32(0)
L3:
	;
	if v5 == int32(0) {
		v200 = v2
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v11 = int32(8)
	v12 = int32(506846)
	goto L8
L5:
	;
	v118 = v5 + v113
	goto L38
L6:
	;
	if v49-v50 == int32(0) {
		v113 = v11
		goto L5
	} else {
		goto L20
	}
L8:
	;
	goto L9
L9:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v19 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v20 = v5
	v21 = v12
	v22 = v11
	v23 = v19
	goto L14
L11:
	;
	v45 = v12
	v49 = int32(0)
	goto L12
L12:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	goto L6
L13:
	;
	v45 = v40
	v49 = v42
	goto L12
L14:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v23 != v25 {
		v40 = v21
		v42 = v23
		goto L13
	} else {
		goto L16
	}
L15:
	;
	v40 = v34
	v42 = int32(0)
	goto L13
L16:
	;
	if v25 == int32(0) {
		v40 = v21
		v42 = v23
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v30 = v22 - int32(1)
	if v30 == int32(0) {
		v40 = v21
		v42 = v23
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v33 = int32(1)
	v34 = v21 + v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v35 != 0 {
		v20 = v20 + v33
		v21 = v34
		v22 = v30
		v23 = v35
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	v60 = int32(14)
	v61 = int32(506831)
	goto L23
L21:
	;
	if v98-v99 == int32(0) {
		v113 = v60
		goto L5
	} else {
		goto L35
	}
L23:
	;
	goto L24
L24:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v68 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v69 = v5
	v70 = v61
	v71 = v60
	v72 = v68
	goto L29
L26:
	;
	v94 = v61
	v98 = int32(0)
	goto L27
L27:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	goto L21
L28:
	;
	v94 = v89
	v98 = v91
	goto L27
L29:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v72 != v74 {
		v89 = v70
		v91 = v72
		goto L28
	} else {
		goto L31
	}
L30:
	;
	v89 = v83
	v91 = int32(0)
	goto L28
L31:
	;
	if v74 == int32(0) {
		v89 = v70
		v91 = v72
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v79 = v71 - int32(1)
	if v79 == int32(0) {
		v89 = v70
		v91 = v72
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v82 = int32(1)
	v83 = v70 + v82
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)))
	if v84 != 0 {
		v69 = v69 + v82
		v70 = v83
		v71 = v79
		v72 = v84
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	F_pfree(m, v5)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	return int32(0)
L37:
	;
	F_pfree(m, v5)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L53
	}
L38:
	;
	v123 = v118 + int32(1)
	v124 = int32(*(*int8)(unsafe.Add(mBase, uint32(v118))))
	v125 = F___isspace(m, v124)
	mBase = m.M
	if v125 != 0 {
		v118 = v123
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v126 = int32(1)
	switch v124&int32(255) - int32(43) {
	case 0:
		v132 = v126
		goto L42
	default:
		v134 = v124
		v135 = v118
		v136 = v126
		goto L41
	case 2:
		goto L43
	}
L40:
	;
	goto L39
L41:
	;
	v137 = int32(0)
	v139 = v134 - int32(48)
	if base.Ui32(v139) <= base.Ui32(int32(9)) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v133 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123))))
	v134 = v133
	v135 = v123
	v136 = v132
	goto L41
L43:
	;
	v132 = int32(0)
	goto L42
L44:
	;
	v142 = v137
	v143 = v139
	v144 = v135
	goto L47
L45:
	;
	v156 = v137
	goto L46
L46:
	;
	if v136 != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v146 = int32(10)
	v148 = v142*v146 - v143
	v149 = int32(*(*int8)(unsafe.Add(mBase, uint32(v144)+1)))
	v153 = v149 - int32(48)
	if base.Ui32(v153) < base.Ui32(v146) {
		v142 = v148
		v143 = v153
		v144 = v144 + int32(1)
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v156 = v148
	goto L46
L49:
	;
	goto L48
L50:
	;
	v162 = int32(0) - v156
	goto L52
L51:
	;
	v162 = v156
	goto L52
L52:
	;
	goto L37
L53:
	;
	if v162 == int32(-1) {
		v200 = v2
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v167 = int32(1)
	v168 = int32(0)
	if v162 < v168 {
		v186 = v168
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v186 == int32(0) {
		v200 = v167
		goto L1
	} else {
		goto L62
	}
L56:
	;
	goto L55
L57:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+16))
	if base.Ui32(v175) <= base.Ui32(v162) {
		v186 = int32(0)
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v180 = v177 + v162*int32(640)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v180)+44))
	if v182 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v183 = v180
	goto L61
L60:
	;
	v183 = int32(0)
	goto L61
L61:
	;
	v186 = v183
	goto L56
L62:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v186)+60))
	v191 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	if v189 != v191 {
		v200 = v167
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v186)+68))
	if v195 == l0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v197 = int32(2)
	goto L66
L65:
	;
	v197 = int32(1)
	goto L66
L66:
	;
	v200 = v197
	goto L1
}
