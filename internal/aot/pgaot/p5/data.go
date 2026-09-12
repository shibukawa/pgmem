package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyErrorData(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1144]))
	if int32(0) <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = F_palloc(m, int32(100))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1144])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L78
	}
L4:
	;
	return int32(0)
L5:
	;
	v12 = int32(100)
	goto L7
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v19 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v17 = F__emscripten_memcpy_bulkmem(m, v8, v4*v12+int32(4541696), v12)
	mBase = m.M
	goto L9
L9:
	;
	goto L6
L10:
	;
	v20 = F_pstrdup(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v23 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v20
	goto L12
L14:
	;
	v24 = F_pstrdup(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if v27 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v24
	goto L16
L18:
	;
	v28 = F_pstrdup(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	if v31 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v28
	goto L20
L22:
	;
	v32 = F_pstrdup(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if v35 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v32
	goto L24
L26:
	;
	v36 = F_pstrdup(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	if v39 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v36
	goto L28
L30:
	;
	v40 = F_pstrdup(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	if v43 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v40
	goto L32
L34:
	;
	v44 = F_pstrdup(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	if v47 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v44
	goto L36
L38:
	;
	v48 = F_pstrdup(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	if v51 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v48
	goto L40
L42:
	;
	v52 = F_pstrdup(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	if v55 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v52
	goto L44
L46:
	;
	v56 = F_pstrdup(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	if v59 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v56
	goto L48
L50:
	;
	v60 = F_pstrdup(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	if v63 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = v60
	goto L52
L54:
	;
	v64 = F_pstrdup(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	if v67 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v64
	goto L56
L58:
	;
	v68 = F_pstrdup(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v17)+68))
	if v71 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v68
	goto L60
L62:
	;
	v72 = F_pstrdup(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	if v75 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v72
	goto L64
L66:
	;
	v76 = F_pstrdup(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	if v79 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v76
	goto L68
L70:
	;
	v80 = F_pstrdup(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v17)+88))
	if v83 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v80
	goto L72
L74:
	;
	v84 = F_pstrdup(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = v88
	return v17
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+88)) = v84
	goto L76
L78:
	;
	F_errmsg_internal(m, int32(471275), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(518217), int32(1760), int32(526150))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CopyGetData(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v13 {
	case 0:
		goto L3
	case 1:
		goto L5
	case 2:
		goto L4
	default:
		v197 = v4
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L17
	} else {
		goto L66
	}
L2:
	;
	m.G0 = v11 + int32(32)
	return v197
L3:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v177 = F_fread(m, l1, int32(1), l2, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L17
	} else {
		goto L58
	}
L4:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v173 = m.T0[v172].(func(*base.Module, int32, int32, int32) int32)(m, l1, int32(1), l2)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L17
	} else {
		goto L57
	}
L5:
	;
	if l2 <= int32(0) {
		v197 = v4
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v17 = l1
	v18 = l2
	v20 = v4
	goto L7
L7:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v24 != 0 {
		v197 = v20
		goto L2
	} else {
		goto L9
	}
L8:
	;
	v197 = v164
	goto L2
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v26 < v27 {
		v154 = v26
		v156 = v25
		v158 = v27
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v159 = v158 - v154
	if v159 < v18 {
		goto L51
	} else {
		goto L52
	}
L11:
	;
	goto L15
L12:
	;
	if v47 == int32(-1) {
		goto L40
	} else {
		goto L41
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L17
	} else {
		goto L36
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L17
	} else {
		goto L31
	}
L15:
	;
	v37 = int32(4543424)
	v39 = *(*int32)(unsafe.Add(mBase, _consts[283]))
	*(*int32)(unsafe.Add(mBase, _consts[283])) = v39 + int32(1)
	F_pq_startmsgread(m)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v77 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)) = uint8(v77)
	v197 = v20
	goto L2
L17:
	;
	return int32(0)
L18:
	;
	v47 = F_pq_getbyte(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v50 = v47 - int32(72)
	if base.Ui32(int32(30)) < base.Ui32(v50) {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	if int32(1)<<(uint(v50)%32)&int32(1207961601) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v65 = F_pq_getmessage(m, v64, v63)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L17
	} else {
		goto L26
	}
L22:
	;
	if v50 != int32(28) {
		goto L12
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v63 = int32(10000)
	goto L21
L25:
	;
	v63 = int32(1073741822)
	goto L21
L26:
	;
	if v65 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	v67 = int32(4543424)
	v69 = *(*int32)(unsafe.Add(mBase, _consts[283]))
	*(*int32)(unsafe.Add(mBase, _consts[283])) = v69 - int32(1)
	switch v50 {
	case 0, 11:
		goto L15
	default:
		goto L29
	case 27:
		goto L28
	case 30:
		goto L14
	}
L28:
	;
	goto L16
L29:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v74 < v75 {
		v154 = v74
		v156 = v73
		v158 = v75
		goto L10
	} else {
		goto L30
	}
L30:
	;
	goto L15
L31:
	;
	F_errcode(m, int32(67371461))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L17
	} else {
		goto L32
	}
L32:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v87 = F_pq_getmsgstring(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L17
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v87
	F_errmsg(m, int32(212927), v11+int32(16))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L17
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(518774), int32(318), int32(526097))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L17
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L17
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(267500), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L17
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(518774), int32(303), int32(526097))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L17
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L17
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L17
	} else {
		goto L47
	}
L43:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L17
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(267500), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L17
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(518774), int32(278), int32(526097))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L17
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L17
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v47
	F_errmsg(m, int32(287983), v11)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L17
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(518774), int32(295), int32(526097))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L17
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	v161 = v159
	goto L53
L52:
	;
	v161 = v18
	goto L53
L53:
	;
	F_pq_copymsgbytes(m, v156, v17, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L17
	} else {
		goto L54
	}
L54:
	;
	v164 = v161 + v20
	v165 = v18 - v161
	if v165 <= int32(0) {
		v197 = v164
		goto L2
	} else {
		goto L55
	}
L55:
	;
	if v164 <= int32(0) {
		v17 = v17 + v161
		v18 = v165
		v20 = v164
		goto L7
	} else {
		goto L56
	}
L56:
	;
	goto L8
L57:
	;
	v197 = v173
	goto L2
L58:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+76))
	if v180 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	if int32(base.Ui32(v185)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L1
	} else {
		goto L64
	}
L60:
	;
	goto L59
L61:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v185 = v183
	goto L60
L62:
	;
	goto L63
L63:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v185 = v184
	goto L60
L64:
	;
	if v177 != 0 {
		v197 = v177
		goto L2
	} else {
		goto L65
	}
L65:
	;
	v190 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)) = uint8(v190)
	v197 = int32(0)
	goto L2
L66:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L17
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(306606), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L17
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(518774), int32(256), int32(526097))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L17
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_checkDataDir(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v4 = m.G0
	v6 = v4 - int32(144)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v14 = F___fstatat(m, int32(-100), v9, v6+int32(48), int32(0))
	mBase = m.M
	if v14 != 0 {
		v16 = *(*int32)(unsafe.Add(mBase, _consts[87]))
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			F_errcode_for_file_access(m)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _consts[136]))
				if v16 == int32(44) {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v24
					F_errmsg(m, int32(75747), v6+int32(16))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						F_errfinish(m, int32(512281), int32(359), int32(222912))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v24
					F_errmsg(m, int32(308174), v6+int32(32))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						F_errfinish(m, int32(512281), int32(364), int32(222912))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
		if v38&int32(61440) != int32(16384) {
			F_errstart_cold(m, int32(22), int32(0))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					v69 = *(*int32)(unsafe.Add(mBase, _consts[136]))
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v69
					F_errmsg(m, int32(13576), v6)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						F_errfinish(m, int32(512281), int32(372), int32(222912))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, _consts[136]))
			F_ValidatePgVersion(m, v44)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				m.G0 = v6 + int32(144)
				return
			}
		}
	}
}
func F_dataPrepareDownlink(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	v5 = F_palloc(m, int32(10))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if l1 < int32(0) {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[12]))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v12+(l1^int32(-1))<<(uint(int32(2))%32))))
			v26 = v18
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _consts[13]))
			v26 = v20 + l1<<(uint(int32(13))%32) + int32(-8192)
		}
		if l1 < int32(0) {
			v30 = *(*int32)(unsafe.Add(mBase, _consts[16]))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v30+(l1^int32(-1))<<(uint(int32(6))%32))+16))
			v45 = v36
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, _consts[17]))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v38+l1<<(uint(int32(6))%32)+int32(-64))+16))
			v45 = v44
		}
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = base.I32_rotr(v45, int32(16))
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v49
		v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+28)))
		*(*uint16)(unsafe.Add(mBase, uint32(v5)+8)) = uint16(v51)
		return v5
	}
}
