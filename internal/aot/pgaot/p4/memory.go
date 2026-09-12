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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v4&int32(15)*int32(36))+uint32(_consts[1443])))
	v12 = m.T0[v11].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
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
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
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
		goto L50
	} else {
		goto L51
	}
L2:
	;
	v50 = F_strlen(m, v20)
	mBase = m.M
	v52 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1444])))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+34)) = uint8(v52)
	v55 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1445])))
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+32)) = uint16(v55)
	v59 = F_strlen(m, v16+int32(32))
	mBase = m.M
	if int32(101) <= v50 {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	v21 = int32(319809)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1296])))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v25 == int32(0) {
		v44 = v24
		v45 = v25
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v47 = v18
	goto L5
L5:
	;
	v48 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+32)) = uint8(v48)
	v205 = v47
	goto L1
L6:
	;
	if v45-v44 != 0 {
		goto L2
	} else {
		goto L14
	}
L7:
	;
	goto L6
L8:
	;
	if v24 != v25 {
		v44 = v24
		v45 = v25
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v29 = v18
	v30 = v21
	goto L10
L10:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	if v34 == int32(0) {
		v44 = v33
		v45 = v34
		goto L7
	} else {
		goto L12
	}
L11:
	;
	v44 = v33
	v45 = v34
	goto L7
L12:
	;
	v37 = int32(1)
	if v33 == v34 {
		v29 = v29 + v37
		v30 = v30 + v37
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v47 = v20
	goto L5
L15:
	;
	v63 = F_pg_mbcliplen(m, v20, v50, int32(100))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v65 = v50
	goto L17
L17:
	;
	if v65 <= int32(0) {
		v169 = v59
		goto L20
	} else {
		goto L21
	}
L18:
	;
	return
L19:
	;
	v65 = v63
	goto L17
L20:
	;
	v184 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(32)+v169))) = uint8(v184)
	if v50 < int32(101) {
		v205 = v18
		goto L1
	} else {
		goto L48
	}
L21:
	;
	v69 = v65 & int32(3)
	if v69 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if base.Ui32(v65) < base.Ui32(int32(4)) {
		v169 = v104
		goto L20
	} else {
		goto L32
	}
L23:
	;
	v103 = v20
	v104 = v59
	v111 = v65
	goto L22
L24:
	;
	goto L25
L25:
	;
	v72 = v20
	v73 = v59
	v77 = int32(0)
	v80 = v65
	goto L26
L26:
	;
	v85 = int32(32)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if base.Ui32(v89) <= base.Ui32(v85) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v103 = v97
	v104 = v95
	v111 = v99
	goto L22
L28:
	;
	v92 = v85
	goto L30
L29:
	;
	v92 = v89
	goto L30
L30:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16+v85+v73))) = uint8(v92)
	v94 = int32(1)
	v95 = v73 + v94
	v97 = v72 + v94
	v99 = v80 - v94
	v101 = v77 + v94
	if v101 != v69 {
		v72 = v97
		v73 = v95
		v77 = v101
		v80 = v99
		goto L26
	} else {
		goto L31
	}
L31:
	;
	goto L27
L32:
	;
	v118 = v103
	v119 = v104
	v126 = v111
	goto L33
L33:
	;
	v131 = int32(32)
	v133 = v16 + v131 + v119
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if base.Ui32(v135) <= base.Ui32(v131) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v169 = v159
	goto L20
L35:
	;
	v138 = v131
	goto L37
L36:
	;
	v138 = v135
	goto L37
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v133))) = uint8(v138)
	v140 = int32(32)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	if base.Ui32(v141) <= base.Ui32(v140) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v144 = v140
	goto L40
L39:
	;
	v144 = v141
	goto L40
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+1)) = uint8(v144)
	v146 = int32(32)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+2)))
	if base.Ui32(v147) <= base.Ui32(v146) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v150 = v146
	goto L43
L42:
	;
	v150 = v147
	goto L43
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+2)) = uint8(v150)
	v152 = int32(32)
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+3)))
	if base.Ui32(v153) <= base.Ui32(v152) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v156 = v152
	goto L46
L45:
	;
	v156 = v153
	goto L46
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+3)) = uint8(v156)
	v158 = int32(4)
	v159 = v119 + v158
	if base.Ui32(v126-int32(5)) < base.Ui32(int32(-2)) {
		v118 = v118 + v158
		v119 = v159
		v126 = v126 - v158
		goto L33
	} else {
		goto L47
	}
L47:
	;
	goto L34
L48:
	;
	v189 = v16 + int32(32)
	v190 = F_strlen(m, v189)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v190+v189))) = int32(3026478)
	v205 = v18
	goto L1
L49:
	;
	m.G0 = v16 + int32(144)
	return
L50:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _consts[1411]))
	if int32(2) <= v19 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v257 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L18
	} else {
		goto L61
	}
L53:
	;
	v214 = int32(1)
	goto L56
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v16 + int32(32)
	v253 = F_pg_fprintf(m, v210, int32(726998), v16)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L18
	} else {
		goto L60
	}
L56:
	;
	v229 = F_pg_fprintf(m, v210, int32(725444), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L18
	} else {
		goto L58
	}
L57:
	;
	goto L55
L58:
	;
	v232 = v214 + int32(1)
	if v232 != v19 {
		v214 = v232
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	goto L49
L61:
	;
	if v257 == int32(0) {
		goto L49
	} else {
		goto L62
	}
L62:
	;
	F_errhidestmt(m)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L18
	} else {
		goto L63
	}
L63:
	;
	F_errhidecontext(m)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L18
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v16 + int32(32)
	F_errmsg_internal(m, int32(174348), v16+int32(16))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L18
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(487623), int32(1044), int32(87363))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L18
	} else {
		goto L66
	}
L66:
	;
	goto L49
}
