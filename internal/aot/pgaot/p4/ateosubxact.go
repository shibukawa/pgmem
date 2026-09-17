package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AtEOSubXact_SPI(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[0]))
	if v8 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l0 != 0 {
		goto L42
	} else {
		goto L43
	}
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[1]))
	v15 = v12 + v8<<(uint(int32(6))%32)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	if v16 != l1 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+41)))
	if v18 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_MemoryContextDelete(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v24 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	return
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(0)
	goto L7
L10:
	;
	F_MemoryContextDelete(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v15)+48))
	*(*int64)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[2])) = v30
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[3])) = v33
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[4])) = v36
	v38 = int32(_a_F_AtEOSubXact_SPI_0)
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[0]))
	v42 = v40 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[0])) = v42
	v45 = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[1]))
	v52 = base.B2i32(v42 < v45)
	if v42 < v45 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(0)
	goto L12
L14:
	;
	v53 = v45
	goto L16
L15:
	;
	v53 = v47 + v42<<(uint(int32(6))%32)
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[5])) = v53
	if v42 < v45 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if l0 == int32(0) {
		goto L1
	} else {
		goto L35
	}
L18:
	;
	v57 = v42
	v58 = v47
	goto L19
L19:
	;
	v63 = v58 + v57<<(uint(int32(6))%32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+32))
	if v64 != l1 {
		goto L17
	} else {
		goto L21
	}
L20:
	;
	goto L17
L21:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+41)))
	if v66 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+24))
	if v67 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_MemoryContextDelete(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L8
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	if v72 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+24)) = int32(0)
	goto L25
L27:
	;
	F_MemoryContextDelete(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L8
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v63)+48))
	*(*int64)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[2])) = v78
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v63)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[3])) = v81
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v63)+60))
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[4])) = v84
	v86 = int32(_a_F_AtEOSubXact_SPI_0)
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[0]))
	v90 = v88 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[0])) = v90
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[1]))
	v98 = int32(0)
	v100 = base.B2i32(v98 <= v90)
	if v98 <= v90 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+20)) = int32(0)
	goto L29
L31:
	;
	v101 = v94 + v90<<(uint(int32(6))%32)
	goto L33
L32:
	;
	v101 = v98
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[5])) = v101
	if v98 <= v90 {
		v57 = v90
		v58 = v94
		goto L19
	} else {
		goto L34
	}
L34:
	;
	goto L20
L35:
	;
	v113 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	if v113 == int32(0) {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(64))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(_a_F_AtEOSubXact_SPI_1), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	F_errhint(m, int32(_a_F_AtEOSubXact_SPI_2), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_AtEOSubXact_SPI_3), int32(532), int32(_a_F_AtEOSubXact_SPI_4))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L8
	} else {
		goto L41
	}
L41:
	;
	goto L1
L42:
	;
	return
L43:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[5]))
	if v140 == int32(0) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	if base.Ui32(l1) <= base.Ui32(v143) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = int32(0)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v140)+24))
	F_MemoryContextReset(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L8
	} else {
		goto L48
	}
L46:
	;
	v152 = v140
	goto L47
L47:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+16))
	if v153 == int32(0) {
		goto L42
	} else {
		goto L49
	}
L48:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[5]))
	v152 = v151
	goto L47
L49:
	;
	v158 = v152 + int32(16)
	v160 = v153
	goto L50
L50:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if base.Ui32(v165) < base.Ui32(l1) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L42
L52:
	;
	if v164 != 0 {
		v158 = v187
		v160 = v164
		goto L50
	} else {
		goto L63
	}
L53:
	;
	v187 = v160
	goto L52
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v164
	v169 = v160 - int32(28)
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[5]))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+8))
	if v169 == v172 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171)+8)) = int32(0)
	goto L58
L57:
	;
	goto L58
L58:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[3]))
	if v177 == v169 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_SPI[3])) = int32(0)
	goto L61
L60:
	;
	goto L61
L61:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v160-int32(4))))
	F_MemoryContextDelete(m, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L8
	} else {
		goto L62
	}
L62:
	;
	v187 = v158
	goto L52
L63:
	;
	goto L51
}
