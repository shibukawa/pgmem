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
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v177 int32
	_ = v177
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v223 int32
	_ = v223
	var v238 int32
	_ = v238
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
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v53
	v56 = v50 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v56
	v62 = F__emscripten_memset_bulkmem(m, v51+int32(8), base.I32_extend8_s(v53), v48)
	mBase = m.M
	goto L13
L13:
	;
	if v20 <= int32(0) {
		v238 = v56
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(base.Ui32(v238) >> (uint(int32(2)) % 32))
	return v51
L15:
	;
	v68 = v48 & int32(3)
	v81 = v2
	goto L16
L16:
	;
	v90 = int32(4)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(4)+v81<<(uint(v90)%32))))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+4)))
	if v94&v90 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = int64(17179869216)
	v238 = int32(32)
	goto L14
L18:
	;
	if v48 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	v223 = v81 + int32(1)
	if v223 != v20 {
		v81 = v223
		goto L16
	} else {
		goto L33
	}
L22:
	;
	v100 = v93 + int32(8)
	v101 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v48) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v106 = v101
	v113 = v101
	goto L26
L24:
	;
	v158 = v101
	goto L25
L25:
	;
	if v68 == int32(0) {
		goto L21
	} else {
		goto L29
	}
L26:
	;
	v123 = v106 + v62
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106+v100))))
	v127 = v124 | v126
	*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v127)
	v130 = v106 | int32(1)
	v131 = v62 + v130
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+v130))))
	v135 = v132 | v134
	*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v135)
	v138 = v106 | int32(2)
	v139 = v62 + v138
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+v138))))
	v143 = v140 | v142
	*(*uint8)(unsafe.Add(mBase, uint32(v139))) = uint8(v143)
	v146 = v106 | int32(3)
	v147 = v62 + v146
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+v146))))
	v151 = v148 | v150
	*(*uint8)(unsafe.Add(mBase, uint32(v147))) = uint8(v151)
	v153 = int32(4)
	v154 = v106 + v153
	v156 = v113 + v153
	if v156 != v48&int32(2147483644) {
		v106 = v154
		v113 = v156
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v158 = v154
	goto L25
L28:
	;
	goto L27
L29:
	;
	v177 = v158
	v189 = v101
	goto L30
L30:
	;
	v194 = v177 + v62
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177+v100))))
	v198 = v195 | v197
	*(*uint8)(unsafe.Add(mBase, uint32(v194))) = uint8(v198)
	v200 = int32(1)
	v203 = v189 + v200
	if v203 != v68 {
		v177 = v177 + v200
		v189 = v203
		goto L30
	} else {
		goto L32
	}
L31:
	;
	goto L21
L32:
	;
	goto L31
L33:
	;
	v238 = v56
	goto L14
}
