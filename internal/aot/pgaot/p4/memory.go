package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetMemoryChunkSpace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4&int32(15)*int32(36))+uint32(_c_F_GetMemoryChunkSpace[0])))
	v10 = m.T0[v9].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_MemoryContextDeleteChildren(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = v8
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v20 = v12
	goto L6
L5:
	;
	goto L3
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	if v23 != 0 {
		v20 = v23
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v67 != 0 {
		v12 = v67
		goto L4
	} else {
		goto L29
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v25 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v28 = v25
	goto L12
L10:
	;
	v46 = v24
	goto L11
L11:
	;
	if v46 != 0 {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	m.T0[v36].(func(*base.Module, int32))(m, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v46 = v40
	goto L11
L14:
	;
	return
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v39 != 0 {
		v28 = v39
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v49 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = int32(0)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	m.T0[v63].(func(*base.Module, int32))(m, v20)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L14
	} else {
		goto L27
	}
L20:
	;
	if v48 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+28)) = v48
	goto L20
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v48
	goto L20
L24:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v52
	goto L26
L25:
	;
	goto L26
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = int32(0)
	goto L19
L27:
	;
	if v20 != v12 {
		v20 = v24
		goto L6
	} else {
		goto L28
	}
L28:
	;
	goto L7
L29:
	;
	goto L5
}
func F_MemoryContextStatsPrint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	v14 = m.G0
	v16 = v14 - int32(144)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if l3 != 0 {
		goto L49
	} else {
		goto L50
	}
L2:
	;
	v51 = F_strlen(m, v20)
	mBase = m.M
	v53 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MemoryContextStatsPrint[0])))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+34)) = uint8(v53)
	v56 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_MemoryContextStatsPrint[1])))
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+32)) = uint16(v56)
	v60 = F_strlen(m, v16+int32(32))
	mBase = m.M
	if int32(101) <= v51 {
		goto L14
	} else {
		goto L15
	}
L3:
	;
	v21 = int32(_a_F_MemoryContextStatsPrint_0)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MemoryContextStatsPrint[2])))
	if base.B2i32(v24 == int32(0))|base.B2i32(v24 != v27) != 0 {
		v45 = v24
		v46 = v27
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v48 = v18
	goto L5
L5:
	;
	v49 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+32)) = uint8(v49)
	v199 = v48
	goto L1
L6:
	;
	if v45-v46 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	goto L6
L8:
	;
	v30 = v18
	v31 = v21
	goto L9
L9:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v35 == int32(0) {
		v45 = v35
		v46 = v34
		goto L7
	} else {
		goto L11
	}
L10:
	;
	v45 = v35
	v46 = v34
	goto L7
L11:
	;
	v38 = int32(1)
	if v35 == v34 {
		v30 = v30 + v38
		v31 = v31 + v38
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v48 = v20
	goto L5
L14:
	;
	v64 = F_pg_mbcliplen(m, v20, v51, int32(100))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v66 = v51
	goto L16
L16:
	;
	if v66 <= int32(0) {
		v168 = v60
		goto L19
	} else {
		goto L20
	}
L17:
	;
	return
L18:
	;
	v66 = v64
	goto L16
L19:
	;
	v181 = v16 + int32(32)
	v183 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v181+v168))) = uint8(v183)
	if v51 < int32(101) {
		v199 = v18
		goto L1
	} else {
		goto L47
	}
L20:
	;
	v70 = v66 & int32(3)
	if v70 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if base.Ui32(v66) < base.Ui32(int32(4)) {
		v168 = v105
		goto L19
	} else {
		goto L31
	}
L22:
	;
	v104 = v20
	v105 = v60
	v111 = v66
	goto L21
L23:
	;
	goto L24
L24:
	;
	v73 = v20
	v74 = v60
	v80 = v66
	v85 = int32(0)
	goto L25
L25:
	;
	v86 = int32(32)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if base.Ui32(v90) <= base.Ui32(v86) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v104 = v98
	v105 = v96
	v111 = v100
	goto L21
L27:
	;
	v93 = v86
	goto L29
L28:
	;
	v93 = v90
	goto L29
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16+v86+v74))) = uint8(v93)
	v95 = int32(1)
	v96 = v74 + v95
	v98 = v73 + v95
	v100 = v80 - v95
	v102 = v85 + v95
	if v102 != v70 {
		v73 = v98
		v74 = v96
		v80 = v100
		v85 = v102
		goto L25
	} else {
		goto L30
	}
L30:
	;
	goto L26
L31:
	;
	v119 = v104
	v120 = v105
	v126 = v111
	goto L32
L32:
	;
	v132 = int32(32)
	v134 = v16 + v132 + v120
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if base.Ui32(v136) <= base.Ui32(v132) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v168 = v160
	goto L19
L34:
	;
	v139 = v132
	goto L36
L35:
	;
	v139 = v136
	goto L36
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v139)
	v141 = int32(32)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
	if base.Ui32(v142) <= base.Ui32(v141) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v145 = v141
	goto L39
L38:
	;
	v145 = v142
	goto L39
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v134)+1)) = uint8(v145)
	v147 = int32(32)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+2)))
	if base.Ui32(v148) <= base.Ui32(v147) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v151 = v147
	goto L42
L41:
	;
	v151 = v148
	goto L42
L42:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v134)+2)) = uint8(v151)
	v153 = int32(32)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+3)))
	if base.Ui32(v154) <= base.Ui32(v153) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v157 = v153
	goto L45
L44:
	;
	v157 = v154
	goto L45
L45:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v134)+3)) = uint8(v157)
	v159 = int32(4)
	v160 = v120 + v159
	if v159 < v126 {
		v119 = v119 + v159
		v120 = v160
		v126 = v126 - v159
		goto L32
	} else {
		goto L46
	}
L46:
	;
	goto L33
L47:
	;
	v187 = F_strlen(m, v181)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v187+v181))) = int32(_a_F_MemoryContextStatsPrint_1)
	v199 = v18
	goto L1
L48:
	;
	m.G0 = v16 + int32(144)
	return
L49:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_MemoryContextStatsPrint[3]))
	if int32(2) <= v19 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v252 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L17
	} else {
		goto L60
	}
L52:
	;
	v209 = int32(1)
	goto L55
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v16 + int32(32)
	v248 = F_pg_fprintf(m, v205, int32(_a_F_MemoryContextStatsPrint_2), v16)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L17
	} else {
		goto L59
	}
L55:
	;
	v224 = F_pg_fprintf(m, v205, int32(_a_F_MemoryContextStatsPrint_3), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L17
	} else {
		goto L57
	}
L56:
	;
	goto L54
L57:
	;
	v227 = v209 + int32(1)
	if v227 != v19 {
		v209 = v227
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	goto L48
L60:
	;
	if v252 == int32(0) {
		goto L48
	} else {
		goto L61
	}
L61:
	;
	F_errhidestmt(m)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L17
	} else {
		goto L62
	}
L62:
	;
	F_errhidecontext(m)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L17
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v16 + int32(32)
	F_errmsg_internal(m, int32(_a_F_MemoryContextStatsPrint_4), v16+int32(16))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L17
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_MemoryContextStatsPrint_5), int32(1044), int32(_a_F_MemoryContextStatsPrint_6))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L17
	} else {
		goto L65
	}
L65:
	;
	goto L48
}
