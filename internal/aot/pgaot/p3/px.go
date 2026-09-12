package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_px_find_combo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	v9 = F_palloc0(m, int32(32))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = F_pstrdup(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L6
	}
L3:
	;
	F_pfree(m, v13)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L75
	}
L4:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	if v199 != 0 {
		goto L69
	} else {
		goto L70
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(7433)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(7434)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(7435)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(7436)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(7437)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(7438)
	F_pfree(m, v13)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L68
	}
L6:
	;
	v15 = int32(47)
	v16 = F___strchrnul(m, v13, v15)
	mBase = m.M
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v18 == v15 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v22 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v22 = v16
	goto L10
L9:
	;
	v22 = int32(0)
	goto L10
L10:
	;
	goto L7
L11:
	;
	v23 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v23)
	v27 = v22 + int32(1)
	v32 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v156 = v9 + int32(24)
	v157 = F_px_find_cipher(m, v13, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L66
	}
L14:
	;
	v34 = int32(0)
	v35 = int32(47)
	v36 = F___strchrnul(m, v27, v35)
	mBase = m.M
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v38 == v35 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v94 = v9 + int32(24)
	v95 = F_px_find_cipher(m, v13, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L45
	}
L16:
	;
	if v42 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v42 = v36
	goto L19
L18:
	;
	v42 = v34
	goto L19
L19:
	;
	goto L16
L20:
	;
	v43 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v43)
	v47 = v42 + int32(1)
	goto L22
L21:
	;
	v47 = v34
	goto L22
L22:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v48 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v49 = int32(58)
	v50 = F___strchrnul(m, v27, v49)
	mBase = m.M
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v52 == v49 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v91 = v32
	goto L25
L25:
	;
	if v47 != 0 {
		v27 = v47
		v32 = v91
		goto L14
	} else {
		goto L44
	}
L26:
	;
	if v56 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v56 = v50
	goto L29
L28:
	;
	v56 = int32(0)
	goto L29
L29:
	;
	goto L26
L30:
	;
	v209 = int32(-6)
	goto L3
L31:
	;
	goto L32
L32:
	;
	v60 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v60)
	v62 = int32(485262)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1325])))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v66 == v60 {
		v85 = v65
		v86 = v66
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v86-v85 != 0 {
		goto L41
	} else {
		goto L42
	}
L34:
	;
	goto L33
L35:
	;
	if v65 != v66 {
		v85 = v65
		v86 = v66
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v70 = v27
	v71 = v62
	goto L37
L37:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v75 == int32(0) {
		v85 = v74
		v86 = v75
		goto L34
	} else {
		goto L39
	}
L38:
	;
	v85 = v74
	v86 = v75
	goto L34
L39:
	;
	v78 = int32(1)
	if v74 == v75 {
		v70 = v70 + v78
		v71 = v71 + v78
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v209 = int32(-5)
	goto L3
L42:
	;
	goto L43
L43:
	;
	v91 = v56 + int32(1)
	goto L25
L44:
	;
	goto L15
L45:
	;
	if v95 != 0 {
		v192 = v94
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v97 = int32(1)
	if v91 == int32(0) {
		v163 = v97
		goto L5
	} else {
		goto L47
	}
L47:
	;
	v100 = int32(184506)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1326])))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v104 == int32(0) {
		v123 = v103
		v124 = v104
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v124-v123 == int32(0) {
		v163 = v97
		goto L5
	} else {
		goto L56
	}
L49:
	;
	goto L48
L50:
	;
	if v103 != v104 {
		v123 = v103
		v124 = v104
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v108 = v91
	v109 = v100
	goto L52
L52:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+1)))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	if v113 == int32(0) {
		v123 = v112
		v124 = v113
		goto L49
	} else {
		goto L54
	}
L53:
	;
	v123 = v112
	v124 = v113
	goto L49
L54:
	;
	v116 = int32(1)
	if v112 == v113 {
		v108 = v108 + v116
		v109 = v109 + v116
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v128 = int32(391048)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1327])))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v132 == int32(0) {
		v151 = v131
		v152 = v132
		goto L58
	} else {
		goto L59
	}
L57:
	;
	if v152-v151 != 0 {
		v192 = v94
		goto L4
	} else {
		goto L65
	}
L58:
	;
	goto L57
L59:
	;
	if v131 != v132 {
		v151 = v131
		v152 = v132
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v136 = v91
	v137 = v128
	goto L61
L61:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+1)))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	if v141 == int32(0) {
		v151 = v140
		v152 = v141
		goto L58
	} else {
		goto L63
	}
L62:
	;
	v151 = v140
	v152 = v141
	goto L58
L63:
	;
	v144 = int32(1)
	if v140 == v141 {
		v136 = v136 + v144
		v137 = v137 + v144
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v163 = int32(0)
	goto L5
L66:
	;
	if v157 != 0 {
		v192 = v156
		goto L4
	} else {
		goto L67
	}
L67:
	;
	v163 = int32(1)
	goto L5
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
	return int32(0)
L69:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	m.T0[v200].(func(*base.Module, int32))(m, v199)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	F_pfree(m, v9)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	F_pfree(m, v13)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	return int32(-3)
L75:
	;
	F_pfree(m, v9)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	return v209
}
