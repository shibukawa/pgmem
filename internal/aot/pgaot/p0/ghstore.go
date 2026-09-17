package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ghstore_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
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
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v176 int32
	_ = v176
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v222 int32
	_ = v222
	var v236 int32
	_ = v236
	v2 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v22 == v2 {
		v39 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v39&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v26 == int32(0) {
		v39 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v29 != int32(7) {
		v39 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v32 != int32(17) {
		v39 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)))
	v39 = v35 ^ int32(1)
	goto L2
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = F_get_fn_opclass_options(m, v42)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v48 = int32(16)
	goto L9
L9:
	;
	v50 = v48 + int32(8)
	v51 = F_palloc(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return int32(0)
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v48 = v47
	goto L9
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = int32(0)
	v56 = v50 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v56
	v59 = v51 + int32(8)
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	base.MemoryFill(m, v59, int32(0), v48)
	goto L15
L14:
	;
	goto L15
L15:
	;
	if v20 <= int32(0) {
		v236 = v56
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(base.Ui32(v236) >> (uint(int32(2)) % 32))
	return v51
L17:
	;
	v67 = v48 & int32(3)
	v83 = v2
	goto L18
L18:
	;
	v89 = int32(4)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(4)+v83<<(uint(v89)%32))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+4)))
	if v93&v89 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = int64(17179869216)
	v236 = int32(32)
	goto L16
L20:
	;
	if v48 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	goto L19
L23:
	;
	v222 = v83 + int32(1)
	if v222 != v20 {
		v83 = v222
		goto L18
	} else {
		goto L35
	}
L24:
	;
	v99 = v92 + int32(8)
	v100 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v48) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v105 = v100
	v112 = v100
	goto L28
L26:
	;
	v159 = v100
	goto L27
L27:
	;
	v176 = v159
	v188 = v100
	goto L32
L28:
	;
	v122 = v105 + v59
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105+v99))))
	v126 = v123 | v125
	*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v126)
	v129 = v105 | int32(1)
	v130 = v59 + v129
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99+v129))))
	v134 = v131 | v133
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v134)
	v137 = v105 | int32(2)
	v138 = v59 + v137
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99+v137))))
	v142 = v139 | v141
	*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v142)
	v145 = v105 | int32(3)
	v146 = v59 + v145
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99+v145))))
	v150 = v147 | v149
	*(*uint8)(unsafe.Add(mBase, uint32(v146))) = uint8(v150)
	v152 = int32(4)
	v153 = v105 + v152
	v155 = v112 + v152
	if v155 != v48&int32(2147483644) {
		v105 = v153
		v112 = v155
		goto L28
	} else {
		goto L30
	}
L29:
	;
	if v67 == int32(0) {
		goto L23
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v159 = v153
	goto L27
L32:
	;
	v193 = v176 + v59
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176+v99))))
	v197 = v194 | v196
	*(*uint8)(unsafe.Add(mBase, uint32(v193))) = uint8(v197)
	v199 = int32(1)
	v202 = v188 + v199
	if v202 != v67 {
		v176 = v176 + v199
		v188 = v202
		goto L32
	} else {
		goto L34
	}
L33:
	;
	goto L23
L34:
	;
	goto L33
L35:
	;
	v236 = v56
	goto L16
}
