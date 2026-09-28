package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SpGistInitPage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v2 = l1
	v3 = int32(_a_F_SpGistInitPage_0)
	v5 = int32(0)
	if v5|(l0&int32(3)|int32(1)) == v5 {
		v21 = l0 + v3
		v23 = l0 + int32(4)
		if base.Ui32(v23) < base.Ui32(v21) {
			v25 = v21
		} else {
			v25 = v23
		}
		v30 = (l0^int32(-1)+v25)&int32(-4) + int32(4)
		if v30 == int32(0) {
		} else {
			base.MemoryFill(m, l0, int32(0), v30)
		}
	} else {
		base.MemoryFill(m, l0, int32(0), v3)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(_a_F_SpGistInitPage_1)
	v44 = int32(_a_F_SpGistInitPage_2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v44)
	v50 = int32(_a_F_SpGistInitPage_3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v50)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v50)
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v54 = l0 + v53
	v55 = int32(_a_F_SpGistInitPage_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+6)) = uint16(v55)
	*(*uint16)(unsafe.Add(mBase, uint32(v54))) = uint16(v2)
	return
}
func F_initSpGistState(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v6 = F_spgGetCache(m, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v6)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v16
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v6)+28))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v20
	v22 = *(*int64)(unsafe.Add(mBase, uint32(v6)+40))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v24
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v6)+52))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30+v31<<(uint(int32(3))%32))+96))
	if v28 != v35 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v37 = F_CreateTupleDescCopy(m, v30)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v149 = v30
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v149
	v154 = F_palloc0(m, int32(16))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L27
	}
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v43 = v37 + v40<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+104)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+96)) = v39
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+100)) = uint16(v47)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+110)) = uint8(v49)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+111)) = uint8(v51)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+124)) = v54
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+113)) = uint8(v54)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+112)) = uint8(v53)
	F_populate_compact_attribute(m, v37, v54)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v62 = int32(0)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v62 < v71 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v149 = v37
	goto L5
L9:
	;
	v75 = v37 + int32(28)
	v82 = v62
	v83 = v71
	v85 = v62
	goto L13
L10:
	;
	v139 = v62
	v146 = v71
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v139
	goto L8
L12:
	;
	v139 = v133
	v146 = v112
	goto L11
L13:
	;
	v91 = v75 + v71<<(uint(int32(3))%32) + v82*int32(100)
	v94 = v75 + v82<<(uint(int32(3))%32)
	if v71 != v83 {
		v112 = v83
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v133 = v71
	goto L12
L15:
	;
	v113 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+2)))
	if v113 <= int32(0) {
		v133 = v82
		goto L12
	} else {
		goto L23
	}
L16:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+7)))
	if v96 != int32(118) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v112 = v82
	goto L15
L18:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)))
	if v99 != int32(1) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+6)))
	if v102&int32(6) != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v105 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+2)))
	if v105 <= int32(0) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+90)))
	if v108 != int32(118) {
		v112 = v71
		goto L15
	} else {
		goto L22
	}
L22:
	;
	goto L17
L23:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+90)))
	if v116 == int32(118) {
		v133 = v82
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+5)))
	v125 = (v85 + v119 - int32(1)) & (int32(0) - v119)
	if int32(_a_F_initSpGistState_0) < v125 {
		v133 = v82
		goto L12
	} else {
		goto L25
	}
L25:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v94))) = uint16(v125)
	v131 = v82 + int32(1)
	if v131 != v71 {
		v82 = v131
		v83 = v112
		v85 = v125 + v113
		goto L13
	} else {
		goto L26
	}
L26:
	;
	goto L14
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v154
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_initSpGistState[0]))
	v159 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)) = uint8(v159)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v158
	return
}
