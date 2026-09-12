package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetXLogInsertRecPtr(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	v6 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(1)
	if v7 != 0 {
		F_s_lock(m, v6, int32(518546), int32(9492), int32(215718))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(0)
			v19 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
			v21 = int64(*(*int32)(unsafe.Add(mBase, _consts[316])))
			v22 = base.I64_div_u_s(v19, v21)
			v24 = v19 - v22*v21
			if base.Ui64(v24) <= base.Ui64(int64(8151)) {
				v42 = v24 + int64(40)
			} else {
				v30 = v24 - int64(8152)
				v31 = int64(8168)
				v32 = base.I64_div_u_s(v30, v31)
				v42 = v30 - v32*v31 + v32<<(uint(int64(13))%64) + int64(8216)
			}
			v44 = int64(*(*int32)(unsafe.Add(mBase, _consts[191])))
			return v22*v44 + v42&int64(4294967295)
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(0)
		v19 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
		v21 = int64(*(*int32)(unsafe.Add(mBase, _consts[316])))
		v22 = base.I64_div_u_s(v19, v21)
		v24 = v19 - v22*v21
		if base.Ui64(v24) <= base.Ui64(int64(8151)) {
			v42 = v24 + int64(40)
		} else {
			v30 = v24 - int64(8152)
			v31 = int64(8168)
			v32 = base.I64_div_u_s(v30, v31)
			v42 = v30 - v32*v31 + v32<<(uint(int64(13))%64) + int64(8216)
		}
		v44 = int64(*(*int32)(unsafe.Add(mBase, _consts[191])))
		return v22*v44 + v42&int64(4294967295)
	}
}
func F_GetXLogWriteRecPtr(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	v3 = int32(4444592)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v4)+280)) = v5
	*(*int64)(unsafe.Add(mBase, _consts[189])) = v5
	v10 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+272)) = v11
	*(*int64)(unsafe.Add(mBase, _consts[190])) = v11
	return v11
}
func F_XLogArchiveNotify(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	v4 = m.G0
	v6 = v4 - int32(1072)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = int32(23166)
	v17 = F_pg_snprintf(m, v6+int32(48), int32(1024), int32(184631), v6+int32(32))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = F_AllocateFile(m, v6+int32(48), int32(33726))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v6 + int32(1072)
	return
L4:
	;
	if v22 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v28 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v45 = F_FreeFile(m, v22)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	if v28 == int32(0) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v6 + int32(48)
	F_errmsg(m, int32(309654), v6)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(518854), int32(457), int32(20981))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L3
L13:
	;
	if v45 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v49 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v68 = F_strlen(m, l0)
	mBase = m.M
	if v68 != int32(16) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	if v49 == int32(0) {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v6 + int32(48)
	F_errmsg(m, int32(309609), v6+int32(16))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(518854), int32(465), int32(20981))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L3
L22:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, _consts[104])))
	if v190 != int32(1) {
		goto L3
	} else {
		goto L55
	}
L23:
	;
	v71 = int32(558417)
	v75 = m.G0
	v77 = v75 - int32(32)
	v78 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v77)+24)) = v78
	*(*int64)(unsafe.Add(mBase, uint32(v77)+16)) = v78
	*(*int64)(unsafe.Add(mBase, uint32(v77)+8)) = v78
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v78
	v86 = int32(*(*uint8)(unsafe.Add(mBase, _consts[273])))
	if v86 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v154 != int32(8) {
		goto L22
	} else {
		goto L45
	}
L25:
	;
	v154 = int32(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, _consts[274])))
	if v90 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v94 = l0
	goto L31
L29:
	;
	goto L30
L30:
	;
	v104 = v71
	v105 = v86
	goto L34
L31:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	if v100 == v86 {
		v94 = v94 + int32(1)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v154 = v94 - l0
	goto L24
L33:
	;
	goto L32
L34:
	;
	v112 = v77 + int32(base.Ui32(v105)>>(uint(int32(3))%32))&int32(28)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v114 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = v113 | v114<<(uint(v105)%32)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
	if v118 != 0 {
		v104 = v104 + v114
		v105 = v118
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v121 == int32(0) {
		v146 = l0
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	v154 = v146 - l0
	goto L24
L38:
	;
	v125 = l0
	v126 = v121
	goto L39
L39:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v77+int32(base.Ui32(v126)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v134)>>(uint(v126)%32))&int32(1) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v146 = v142
	goto L37
L41:
	;
	v146 = v125
	goto L37
L42:
	;
	goto L43
L43:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+1)))
	v142 = v125 + int32(1)
	if v140 != 0 {
		v125 = v142
		v126 = v140
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	v158 = l0 + int32(8)
	v159 = int32(12961)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, _consts[364])))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	if v163 == int32(0) {
		v182 = v162
		v183 = v163
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v183-v182 != 0 {
		goto L22
	} else {
		goto L54
	}
L47:
	;
	goto L46
L48:
	;
	if v162 != v163 {
		v182 = v162
		v183 = v163
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v167 = v158
	v168 = v159
	goto L50
L50:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
	if v172 == int32(0) {
		v182 = v171
		v183 = v172
		goto L47
	} else {
		goto L52
	}
L51:
	;
	v182 = v171
	v183 = v172
	goto L47
L52:
	;
	v175 = int32(1)
	if v171 == v172 {
		v167 = v167 + v175
		v168 = v168 + v175
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _consts[365]))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+4)) = int32(1)
	goto L22
L55:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _consts[365]))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	if v195 != int32(-1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _consts[152]))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	F_SetLatch(m, v200+v195*int32(640)+int32(20))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	goto L3
L59:
	;
	goto L58
}
func F_XLogDropRelation(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v14 int32
	_ = v14
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v10
	F_forget_invalid_pages(m, v6, l1, int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		m.G0 = v6 + int32(16)
		return
	}
}
func F_XLogFileInitInternal(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v135 int64
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int64
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v275 int32
	_ = v275
	var v290 int32
	_ = v290
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int64
	_ = v332
	var v335 int32
	_ = v335
	var v336 int64
	_ = v336
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v348 int64
	_ = v348
	var v349 int64
	_ = v349
	var v353 int64
	_ = v353
	var v405 int32
	_ = v405
	var v406 int64
	_ = v406
	var v410 int32
	_ = v410
	var v428 int32
	_ = v428
	var v429 int64
	_ = v429
	var v433 int32
	_ = v433
	var v450 int32
	_ = v450
	var v451 int64
	_ = v451
	var v456 int32
	_ = v456
	var v457 int64
	_ = v457
	var v462 int32
	_ = v462
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v480 int64
	_ = v480
	var v481 int64
	_ = v481
	var v485 int64
	_ = v485
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int64
	_ = v518
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v530 int64
	_ = v530
	var v531 int64
	_ = v531
	var v535 int64
	_ = v535
	var v587 int32
	_ = v587
	var v588 int64
	_ = v588
	var v592 int32
	_ = v592
	var v610 int32
	_ = v610
	var v611 int64
	_ = v611
	var v615 int32
	_ = v615
	var v632 int32
	_ = v632
	var v633 int64
	_ = v633
	var v638 int32
	_ = v638
	var v639 int64
	_ = v639
	var v644 int32
	_ = v644
	var v652 int32
	_ = v652
	var v660 int64
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	v10 = m.G0
	v12 = v10 - int32(1168)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = l1
	v17 = int64(*(*int32)(unsafe.Add(mBase, _consts[191])))
	v18 = base.I64_div_u_s(int64(4294967296), v17)
	v19 = base.I64_div_u_s(l0, v18)
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+116)) = uint32(v19)
	v22 = l0 - v18*v19
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+120)) = uint32(v22)
	v28 = F_pg_snprintf(m, l3, int32(1024), int32(531342), v12+int32(112))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v32)
	v35 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	v39 = *(*int32)(unsafe.Add(mBase, _consts[200]))
	v40 = int32(14)
	v44 = v35 << (uint(int32(13)) % 32) & (base.B2i32(v39 != v40) << (uint(v40) % 32))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, _consts[44])))
	if v46 != int32(1) {
		v68 = v44
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v72 = F_BasicOpenFile(m, l3, v68|int32(524290))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L16
	}
L4:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[205]))
	switch v50 {
	case 0, 1, 3:
		v68 = v44
		goto L3
	case 2:
		goto L5
	case 4:
		goto L7
	default:
		goto L6
	}
L5:
	;
	v68 = v44 | int32(1052672)
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v68 = v44 | int32(4096)
	goto L3
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v50
	F_errmsg_internal(m, int32(506686), v12)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(518546), int32(8695), int32(109472))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L1
	} else {
		goto L164
	}
L12:
	;
	v765 = int32(4713716)
	v766 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	v767 = F_close(m, v115)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[43])) = v766
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L1
	} else {
		goto L160
	}
L13:
	;
	v742 = F_unlink(m, v12+int32(144))
	mBase = m.M
	v743 = F_close(m, v115)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[43])) = v322
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L156
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L152
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L148
	}
L16:
	;
	if v72 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v77 != int32(44) {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	v698 = v72
	goto L19
L19:
	;
	m.G0 = v12 + int32(1168)
	return v698
L20:
	;
	v82 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v82 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_errmsg_internal(m, int32(405100), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = int32(42)
	v101 = F_pg_snprintf(m, v12+int32(144), int32(1024), int32(484607), v12+int32(80))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	F_errfinish(m, int32(518546), int32(3227), int32(325130))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v104 = v12 + int32(144)
	v105 = F_unlink(m, v104)
	mBase = m.M
	v111 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	if v111&int32(4) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v114 = int32(16578)
	goto L30
L29:
	;
	v114 = int32(194)
	goto L30
L30:
	;
	v115 = F_BasicOpenFile(m, v104, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v115 < int32(0) {
		goto L14
	} else {
		goto L32
	}
L32:
	;
	v119 = int32(0)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, _consts[195])))
	v124 = m.G0
	v126 = v124 - int32(16)
	m.G0 = v126
	if v121 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = int32(167772234)
	v144 = int32(*(*uint8)(unsafe.Add(mBase, _consts[208])))
	if v144 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	F___clock_gettime(m, int32(1), v126)
	mBase = m.M
	v130 = int64(*(*int32)(unsafe.Add(mBase, uint32(v126)+8)))
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v126)))
	v135 = v130 + v131*int64(1000000000)
	goto L36
L35:
	;
	v135 = int64(0)
	goto L36
L36:
	;
	m.G0 = v126 + int32(16)
	goto L33
L37:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v324))) = int32(0)
	v327 = int32(2)
	v330 = int32(1)
	v332 = int64(*(*int32)(unsafe.Add(mBase, _consts[191])))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, _consts[208])))
	if v335 != 0 {
		goto L85
	} else {
		goto L86
	}
L38:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v158 = m.G0
	v160 = v158 - int32(1024)
	m.G0 = v160
	v163 = v148
	v164 = int64(0)
	v166 = int32(0)
	goto L42
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v309 = int32(1)
	v311 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v315 = F_pwrite(m, v115, int32(787992), v309, base.I64_extend_i32_s(v311-v309))
	mBase = m.M
	if v315 == v309 {
		v322 = v119
		goto L37
	} else {
		goto L81
	}
L41:
	;
	if int32(0) <= v290 {
		v322 = v119
		goto L37
	} else {
		goto L80
	}
L42:
	;
	v173 = int32(0)
	if v163 == v173 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	m.G0 = v160 + int32(1024)
	goto L41
L44:
	;
	goto L43
L45:
	;
	v290 = v166
	goto L44
L46:
	;
	goto L47
L47:
	;
	v177 = v163
	v179 = v173
	goto L48
L48:
	;
	v189 = v160 + v179<<(uint(int32(3))%32)
	v190 = int32(8192)
	if base.Ui32(v190) <= base.Ui32(v177) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v204 = m.G0
	v206 = v204 - int32(1024)
	m.G0 = v206
	if v198 <= int32(128) {
		goto L57
	} else {
		goto L58
	}
L50:
	;
	goto L49
L51:
	;
	v193 = v190
	goto L53
L52:
	;
	v193 = v177
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189)+4)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = int32(1875968)
	v198 = v179 + int32(1)
	v199 = v177 - v193
	if base.Ui32(int32(126)) < base.Ui32(v179) {
		goto L50
	} else {
		goto L54
	}
L54:
	;
	if v199 != 0 {
		v177 = v199
		v179 = v198
		goto L48
	} else {
		goto L55
	}
L55:
	;
	goto L50
L56:
	;
	m.G0 = v206 + int32(1024)
	if int32(0) <= v275 {
		v163 = v199
		v164 = v164 + base.I64_extend_i32_u(v275)
		v166 = v166 + v275
		goto L42
	} else {
		goto L79
	}
L57:
	;
	v213 = v160
	v215 = v198
	v217 = int32(0)
	v220 = v164
	goto L60
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(28)
	v275 = int32(-1)
	goto L56
L60:
	;
	if v215 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v275 = v231
	goto L56
L62:
	;
	if v227 < int32(0) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v225 = F_pwrite(m, v115, v223, v224, v220)
	mBase = m.M
	v227 = v225
	goto L62
L64:
	;
	goto L65
L65:
	;
	v226 = F_pwritev(m, v115, v213, v215, v220)
	mBase = m.M
	v227 = v226
	goto L62
L66:
	;
	v275 = int32(-1)
	goto L56
L67:
	;
	goto L68
L68:
	;
	v231 = v227 + v217
	v237 = v213
	v239 = v215
	v240 = v227
	goto L69
L69:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if base.Ui32(v245) <= base.Ui32(v240) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if v237 != v206 {
		goto L75
	} else {
		goto L76
	}
L71:
	;
	v251 = v239 - int32(1)
	if v251 != 0 {
		v237 = v237 + int32(8)
		v239 = v251
		v240 = v240 - v245
		goto L69
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	goto L70
L74:
	;
	v275 = v231
	goto L56
L75:
	;
	v255 = F_memmove(m, v206, v237, v239<<(uint(int32(3))%32))
	mBase = m.M
	goto L77
L76:
	;
	goto L77
L77:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v256 + v240
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v206)+4)) = v259 - v240
	if int32(0) < v239 {
		v213 = v206
		v215 = v239
		v217 = v231
		v220 = v220 + base.I64_extend_i32_u(v227)
		goto L60
	} else {
		goto L78
	}
L78:
	;
	goto L61
L79:
	;
	v290 = v275
	goto L44
L80:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	v322 = v304
	goto L37
L81:
	;
	v319 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v319 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v321 = v319
	goto L84
L83:
	;
	v321 = int32(51)
	goto L84
L84:
	;
	v322 = v321
	goto L37
L85:
	;
	v336 = v332
	goto L87
L86:
	;
	v336 = int64(1)
	goto L87
L87:
	;
	v340 = m.G0
	v342 = v340 - int32(16)
	m.G0 = v342
	if v135 != int64(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	if v322 != 0 {
		goto L13
	} else {
		goto L105
	}
L89:
	;
	F___clock_gettime(m, int32(1), v342)
	mBase = m.M
	v348 = int64(*(*int32)(unsafe.Add(mBase, uint32(v342)+8)))
	v349 = *(*int64)(unsafe.Add(mBase, uint32(v342)))
	v353 = v348 + (v349*int64(1000000000) - v135)
	goto L92
L90:
	;
	goto L91
L91:
	;
	v450 = int32(4530592)
	v451 = *(*int64)(unsafe.Add(mBase, _consts[209]))
	*(*int64)(unsafe.Add(mBase, _consts[209])) = v451 + base.I64_extend_i32_u(v330)
	v456 = int32(4529632)
	v457 = *(*int64)(unsafe.Add(mBase, _consts[210]))
	*(*int64)(unsafe.Add(mBase, _consts[210])) = v457 + v336
	F_pgstat_count_backend_io_op(m, v327, v327, int32(7), v330, v336)
	mBase = m.M
	v462 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[154])) = uint8(v462)
	*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v462)
	m.G0 = v342 + int32(16)
	goto L88
L92:
	;
	v405 = int32(4531552)
	v406 = *(*int64)(unsafe.Add(mBase, _consts[211]))
	*(*int64)(unsafe.Add(mBase, _consts[211])) = v406 + v353
	v410 = *(*int32)(unsafe.Add(mBase, _consts[200]))
	if base.Ui32(int32(16)) < base.Ui32(v410) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	goto L91
L103:
	;
	if int32(1)<<(uint(v410)%32)&int32(115186) == int32(0) {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v428 = int32(4528448)
	v429 = *(*int64)(unsafe.Add(mBase, _consts[212]))
	*(*int64)(unsafe.Add(mBase, _consts[212])) = v429 + v353
	v433 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[154])) = uint8(v433)
	*(*uint8)(unsafe.Add(mBase, _consts[202])) = uint8(v433)
	goto L102
L105:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, _consts[195])))
	v474 = m.G0
	v476 = v474 - int32(16)
	m.G0 = v476
	if v471 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v490 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v490))) = int32(167772233)
	v495 = int32(*(*uint8)(unsafe.Add(mBase, _consts[44])))
	if v495 != int32(1) {
		v509 = int32(0)
		goto L111
	} else {
		goto L112
	}
L107:
	;
	F___clock_gettime(m, int32(1), v476)
	mBase = m.M
	v480 = int64(*(*int32)(unsafe.Add(mBase, uint32(v476)+8)))
	v481 = *(*int64)(unsafe.Add(mBase, uint32(v476)))
	v485 = v480 + v481*int64(1000000000)
	goto L109
L108:
	;
	v485 = int64(0)
	goto L109
L109:
	;
	m.G0 = v476 + int32(16)
	goto L106
L110:
	;
	if v509 != 0 {
		goto L12
	} else {
		goto L117
	}
L111:
	;
	goto L110
L112:
	;
	goto L113
L113:
	;
	v500 = F_fsync(m, v115)
	mBase = m.M
	if v500 != int32(-1) {
		v509 = v500
		goto L111
	} else {
		goto L115
	}
L114:
	;
	v509 = int32(-1)
	goto L111
L115:
	;
	v504 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v504 == int32(27) {
		goto L113
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	v511 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v511))) = int32(0)
	v514 = int32(2)
	v516 = int32(1)
	v518 = int64(0)
	v522 = m.G0
	v524 = v522 - int32(16)
	m.G0 = v524
	if v485 != v518 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v652 = F_close(m, v115)
	mBase = m.M
	if v652 != 0 {
		goto L11
	} else {
		goto L135
	}
L119:
	;
	F___clock_gettime(m, int32(1), v524)
	mBase = m.M
	v530 = int64(*(*int32)(unsafe.Add(mBase, uint32(v524)+8)))
	v531 = *(*int64)(unsafe.Add(mBase, uint32(v524)))
	v535 = v530 + (v531*int64(1000000000) - v485)
	goto L122
L120:
	;
	goto L121
L121:
	;
	v632 = int32(4530544)
	v633 = *(*int64)(unsafe.Add(mBase, _consts[213]))
	*(*int64)(unsafe.Add(mBase, _consts[213])) = v633 + base.I64_extend_i32_u(v516)
	v638 = int32(4529584)
	v639 = *(*int64)(unsafe.Add(mBase, _consts[214]))
	*(*int64)(unsafe.Add(mBase, _consts[214])) = v639 + v518
	F_pgstat_count_backend_io_op(m, v514, v514, v516, v516, v518)
	mBase = m.M
	v644 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[154])) = uint8(v644)
	*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v644)
	m.G0 = v524 + int32(16)
	goto L118
L122:
	;
	v587 = int32(4531504)
	v588 = *(*int64)(unsafe.Add(mBase, _consts[215]))
	*(*int64)(unsafe.Add(mBase, _consts[215])) = v588 + v535
	v592 = *(*int32)(unsafe.Add(mBase, _consts[200]))
	if base.Ui32(int32(16)) < base.Ui32(v592) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	goto L121
L133:
	;
	if int32(1)<<(uint(v592)%32)&int32(115186) == int32(0) {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v610 = int32(4528400)
	v611 = *(*int64)(unsafe.Add(mBase, _consts[216]))
	*(*int64)(unsafe.Add(mBase, _consts[216])) = v611 + v535
	v615 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[154])) = uint8(v615)
	*(*uint8)(unsafe.Add(mBase, _consts[202])) = uint8(v615)
	goto L132
L135:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+136)) = l0
	v660 = int64(*(*int32)(unsafe.Add(mBase, _consts[217])))
	v662 = F_InstallXLogFileSegment(m, v12+int32(136), v12+int32(144), int32(1), l0+v660, l1)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L138
	}
L136:
	;
	v698 = int32(-1)
	goto L19
L137:
	;
	F_errmsg_internal(m, v685, int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L146
	}
L138:
	;
	if v662 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v664 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v664)
	v668 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v676 = F_unlink(m, v12+int32(144))
	mBase = m.M
	v679 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	if v668 == int32(0) {
		goto L136
	} else {
		goto L143
	}
L143:
	;
	v685 = int32(405095)
	v686 = int32(3349)
	goto L137
L144:
	;
	if v679 == int32(0) {
		goto L136
	} else {
		goto L145
	}
L145:
	;
	v685 = int32(405134)
	v686 = int32(3359)
	goto L137
L146:
	;
	F_errfinish(m, int32(518546), v686, int32(325130))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	goto L136
L148:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = l3
	F_errmsg(m, int32(310478), v12+int32(96))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(518546), int32(3216), int32(325130))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(144)
	F_errmsg(m, int32(311157), v12+int32(16))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(518546), int32(3241), int32(325130))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v12 + int32(144)
	F_errmsg(m, int32(310161), v12-int32(-64))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(518546), int32(3302), int32(325130))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v12 + int32(144)
	F_errmsg(m, int32(311426), v12+int32(48))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(518546), int32(3316), int32(325130))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L164:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v12 + int32(144)
	F_errmsg(m, int32(311221), v12+int32(32))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(518546), int32(3326), int32(325130))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_XLogInsert(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v14 int32
	_ = v14
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v100 int64
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int64
	_ = v124
	var v127 int64
	_ = v127
	var v131 int32
	_ = v131
	var v133 int64
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int64
	_ = v195
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v422 int64
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v448 int64
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v488 int32
	_ = v488
	var v491 int64
	_ = v491
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int64
	_ = v506
	var v507 int64
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v531 int64
	_ = v531
	var v532 int64
	_ = v532
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v614 int64
	_ = v614
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v629 int64
	_ = v629
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v638 int64
	_ = v638
	var v641 int32
	_ = v641
	var v642 int64
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v653 int64
	_ = v653
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v720 int32
	_ = v720
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v846 int64
	_ = v846
	var v847 int64
	_ = v847
	var v851 int64
	_ = v851
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v893 int32
	_ = v893
	var v896 int64
	_ = v896
	var v897 int64
	_ = v897
	var v904 int64
	_ = v904
	var v908 int64
	_ = v908
	var v909 int64
	_ = v909
	var v911 int64
	_ = v911
	var v917 int64
	_ = v917
	var v918 int64
	_ = v918
	var v919 int64
	_ = v919
	var v929 int64
	_ = v929
	var v931 int64
	_ = v931
	var v935 int64
	_ = v935
	var v938 int64
	_ = v938
	var v939 int64
	_ = v939
	var v941 int64
	_ = v941
	var v944 int64
	_ = v944
	var v949 int64
	_ = v949
	var v951 int64
	_ = v951
	var v952 int64
	_ = v952
	var v953 int64
	_ = v953
	var v955 int64
	_ = v955
	var v960 int64
	_ = v960
	var v969 int64
	_ = v969
	var v971 int64
	_ = v971
	var v975 int64
	_ = v975
	var v979 int64
	_ = v979
	var v980 int64
	_ = v980
	var v982 int64
	_ = v982
	var v988 int64
	_ = v988
	var v989 int64
	_ = v989
	var v990 int64
	_ = v990
	var v1000 int64
	_ = v1000
	var v1002 int64
	_ = v1002
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1025 int32
	_ = v1025
	var v1026 int64
	_ = v1026
	var v1028 int64
	_ = v1028
	var v1029 int64
	_ = v1029
	var v1031 int64
	_ = v1031
	var v1034 int64
	_ = v1034
	var v1039 int64
	_ = v1039
	var v1041 int64
	_ = v1041
	var v1042 int64
	_ = v1042
	var v1043 int64
	_ = v1043
	var v1045 int64
	_ = v1045
	var v1050 int64
	_ = v1050
	var v1059 int64
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1062 int64
	_ = v1062
	var v1063 int64
	_ = v1063
	var v1066 int64
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1069 int64
	_ = v1069
	var v1070 int64
	_ = v1070
	var v1078 int64
	_ = v1078
	var v1079 int64
	_ = v1079
	var v1085 int64
	_ = v1085
	var v1086 int64
	_ = v1086
	var v1087 int64
	_ = v1087
	var v1097 int64
	_ = v1097
	var v1102 int64
	_ = v1102
	var v1104 int64
	_ = v1104
	var v1107 int64
	_ = v1107
	var v1112 int64
	_ = v1112
	var v1114 int64
	_ = v1114
	var v1115 int64
	_ = v1115
	var v1116 int64
	_ = v1116
	var v1118 int64
	_ = v1118
	var v1123 int64
	_ = v1123
	var v1132 int64
	_ = v1132
	var v1136 int64
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1144 int64
	_ = v1144
	var v1146 int64
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1150 int64
	_ = v1150
	var v1155 int64
	_ = v1155
	var v1173 int64
	_ = v1173
	var v1180 int64
	_ = v1180
	var v1187 int64
	_ = v1187
	var v1189 int64
	_ = v1189
	var v1195 int64
	_ = v1195
	var v1196 int64
	_ = v1196
	var v1197 int64
	_ = v1197
	var v1207 int64
	_ = v1207
	var v1209 int64
	_ = v1209
	var v1226 int64
	_ = v1226
	var v1227 int64
	_ = v1227
	var v1232 int32
	_ = v1232
	var v1236 int32
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1253 int32
	_ = v1253
	var v1256 int64
	_ = v1256
	var v1257 int64
	_ = v1257
	var v1264 int64
	_ = v1264
	var v1267 int64
	_ = v1267
	var v1268 int64
	_ = v1268
	var v1270 int64
	_ = v1270
	var v1276 int64
	_ = v1276
	var v1277 int64
	_ = v1277
	var v1278 int64
	_ = v1278
	var v1288 int64
	_ = v1288
	var v1290 int64
	_ = v1290
	var v1294 int64
	_ = v1294
	var v1296 int64
	_ = v1296
	var v1298 int64
	_ = v1298
	var v1301 int64
	_ = v1301
	var v1306 int64
	_ = v1306
	var v1308 int64
	_ = v1308
	var v1309 int64
	_ = v1309
	var v1310 int64
	_ = v1310
	var v1312 int64
	_ = v1312
	var v1317 int64
	_ = v1317
	var v1326 int64
	_ = v1326
	var v1330 int64
	_ = v1330
	var v1332 int64
	_ = v1332
	var v1334 int64
	_ = v1334
	var v1340 int64
	_ = v1340
	var v1341 int64
	_ = v1341
	var v1342 int64
	_ = v1342
	var v1352 int64
	_ = v1352
	var v1359 int64
	_ = v1359
	var v1361 int64
	_ = v1361
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1385 int64
	_ = v1385
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1396 int64
	_ = v1396
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1438 int64
	_ = v1438
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1457 int32
	_ = v1457
	var v1461 int32
	_ = v1461
	var v1475 int32
	_ = v1475
	var v1478 int64
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1492 int32
	_ = v1492
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1506 int64
	_ = v1506
	var v1507 int64
	_ = v1507
	var v1509 int64
	_ = v1509
	var v1514 int32
	_ = v1514
	var v1518 int64
	_ = v1518
	var v1529 int32
	_ = v1529
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1537 int32
	_ = v1537
	var v1541 int32
	_ = v1541
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1561 int64
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1574 int64
	_ = v1574
	var v1578 int64
	_ = v1578
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1622 int32
	_ = v1622
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1637 int32
	_ = v1637
	var v1638 int64
	_ = v1638
	var v1647 int64
	_ = v1647
	var v1655 int64
	_ = v1655
	var v1693 int32
	_ = v1693
	var v1697 int32
	_ = v1697
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1724 int32
	_ = v1724
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1764 int64
	_ = v1764
	var v1765 int64
	_ = v1765
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1775 int32
	_ = v1775
	var v1782 int32
	_ = v1782
	var v1783 int64
	_ = v1783
	var v1785 int32
	_ = v1785
	var v1786 int64
	_ = v1786
	var v1791 int64
	_ = v1791
	var v1796 int32
	_ = v1796
	var v1797 int64
	_ = v1797
	var v1804 int64
	_ = v1804
	var v1806 int32
	_ = v1806
	var v1807 int64
	_ = v1807
	var v1811 int64
	_ = v1811
	var v1818 int32
	_ = v1818
	var v1827 int64
	_ = v1827
	var v1829 int64
	_ = v1829
	var v1835 int64
	_ = v1835
	var v1838 int64
	_ = v1838
	var v1842 int64
	_ = v1842
	var v1844 int64
	_ = v1844
	var v1846 int32
	_ = v1846
	var v1848 int32
	_ = v1848
	var v1850 int64
	_ = v1850
	var v1853 int32
	_ = v1853
	var v1855 int64
	_ = v1855
	var v1859 int32
	_ = v1859
	var v1861 int64
	_ = v1861
	var v1871 int64
	_ = v1871
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1921 int32
	_ = v1921
	var v1926 int32
	_ = v1926
	var v1929 int32
	_ = v1929
	var v1931 int32
	_ = v1931
	var v1935 int32
	_ = v1935
	var v1937 int32
	_ = v1937
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1982 int32
	_ = v1982
	var v1985 int32
	_ = v1985
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2039 int32
	_ = v2039
	var v2055 int32
	_ = v2055
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2125 int32
	_ = v2125
	var v2127 int32
	_ = v2127
	var v2130 int32
	_ = v2130
	var v2135 int32
	_ = v2135
	var v2139 int32
	_ = v2139
	var v2144 int32
	_ = v2144
	var v2148 int32
	_ = v2148
	var v2154 int32
	_ = v2154
	var v2159 int32
	_ = v2159
	var v2163 int32
	_ = v2163
	var v2167 int32
	_ = v2167
	var v2171 int64
	_ = v2171
	var v2177 int32
	_ = v2177
	var v2182 int32
	_ = v2182
	var v2186 int32
	_ = v2186
	var v2190 int32
	_ = v2190
	var v2198 int32
	_ = v2198
	var v2203 int32
	_ = v2203
	var v2204 int64
	_ = v2204
	var v2206 int32
	_ = v2206
	var v2210 int32
	_ = v2210
	var v2212 int32
	_ = v2212
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2257 int32
	_ = v2257
	var v2260 int32
	_ = v2260
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2314 int32
	_ = v2314
	var v2330 int32
	_ = v2330
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2400 int32
	_ = v2400
	var v2402 int32
	_ = v2402
	var v2405 int32
	_ = v2405
	var v2409 int64
	_ = v2409
	var v2453 int32
	_ = v2453
	v1 = l0
	v14 = int32(0)
	v40 = m.G0
	v42 = v40 - int32(8256)
	m.G0 = v42
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _consts[369])))
	if v45 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, _consts[370])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[371])) = int32(4444864)
	v2453 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[372])) = v2453
	*(*int32)(unsafe.Add(mBase, _consts[373])) = v2453
	*(*uint8)(unsafe.Add(mBase, _consts[360])) = uint8(v2453)
	*(*uint8)(unsafe.Add(mBase, _consts[369])) = uint8(v2453)
	m.G0 = v42 + int32(8256)
	return v2409
L2:
	;
	v2204 = int64(40)
	v2206 = *(*int32)(unsafe.Add(mBase, _consts[372]))
	if v2206 <= int32(0) {
		v2409 = v2204
		goto L1
	} else {
		goto L398
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2186 = m.ExcPending
	if v2186 != 0 {
		goto L78
	} else {
		goto L394
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2163 = m.ExcPending
	if v2163 != 0 {
		goto L78
	} else {
		goto L390
	}
L5:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2148 = m.ExcPending
	if v2148 != 0 {
		goto L78
	} else {
		goto L387
	}
L6:
	;
	if l1&int32(12) != 0 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2135 = m.ExcPending
	if v2135 != 0 {
		goto L78
	} else {
		goto L384
	}
L9:
	;
	if v1 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	if v49 == int32(0) {
		goto L2
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[374]))
	v83 = v14
	v85 = v14
	v86 = v14
	goto L14
L13:
	;
	goto L12
L14:
	;
	v100 = *(*int64)(unsafe.Add(mBase, _consts[277]))
	*(*int64)(unsafe.Add(mBase, uint32(v42+int32(56)))) = v100
	v103 = int32(*(*uint8)(unsafe.Add(mBase, _consts[276])))
	*(*uint8)(unsafe.Add(mBase, uint32(v42+int32(55)))) = uint8(v103)
	goto L16
L15:
	;
	v1929 = int32(0)
	v1931 = *(*int32)(unsafe.Add(mBase, _consts[372]))
	if v1931 <= v1929 {
		v2409 = v1871
		goto L1
	} else {
		goto L373
	}
L16:
	;
	v105 = int32(0)
	v108 = *(*int32)(unsafe.Add(mBase, _consts[375]))
	*(*int32)(unsafe.Add(mBase, _consts[376])) = v108
	*(*int32)(unsafe.Add(mBase, _consts[377])) = v105
	v114 = *(*int32)(unsafe.Add(mBase, _consts[378]))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114+v1))))
	if v116 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v117 = l1 | int32(2)
	goto L19
L18:
	;
	v117 = l1
	goto L19
L19:
	;
	v119 = v108 + int32(24)
	v121 = *(*int32)(unsafe.Add(mBase, _consts[372]))
	if v121 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, _consts[360])))
	if v569&int32(1) == int32(0) {
		v585 = v544
		goto L126
	} else {
		goto L127
	}
L21:
	;
	v124 = int64(0)
	v531 = v124
	v532 = v124
	v544 = v119
	v545 = int32(4444880)
	v556 = v83
	v558 = v85
	v559 = v86
	v561 = v105
	goto L20
L22:
	;
	goto L23
L23:
	;
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v42)+56))
	v131 = *(*int32)(unsafe.Add(mBase, _consts[379]))
	v133 = int64(0)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+55)))
	v138 = int32(0)
	v142 = v133
	v143 = v133
	v153 = v138
	v155 = v119
	v156 = int32(4444880)
	v158 = v131
	v160 = v121
	v161 = v138
	v167 = v83
	v169 = v85
	v170 = v86
	v172 = v105
	goto L24
L24:
	;
	v181 = v158 + v153*int32(8260)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	if v182 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v531 = v506
	v532 = v507
	v544 = v509
	v545 = v510
	v556 = v519
	v558 = v521
	v559 = v522
	v561 = v523
	goto L20
L26:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+1)))
	if v185&int32(1) != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v506 = v142
	v507 = v143
	v509 = v155
	v510 = v156
	v512 = v158
	v514 = v160
	v515 = v161
	v519 = v167
	v521 = v169
	v522 = v170
	v523 = v172
	goto L28
L28:
	;
	v527 = v153 + int32(1)
	if v527 < v514 {
		v142 = v506
		v143 = v507
		v153 = v527
		v155 = v509
		v156 = v510
		v158 = v512
		v160 = v514
		v161 = v515
		v167 = v519
		v169 = v521
		v170 = v522
		v172 = v523
		goto L24
	} else {
		goto L125
	}
L29:
	;
	v210 = int32(0)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v181)+28))
	if v212 != 0 {
		goto L41
	} else {
		goto L42
	}
L30:
	;
	v206 = v143
	v209 = int32(1)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v189 = int32(0)
	if v185&int32(2) != 0 {
		v206 = v143
		v209 = v189
		goto L29
	} else {
		goto L33
	}
L33:
	;
	if v134&int32(1) == int32(0) {
		v206 = v143
		v209 = v189
		goto L29
	} else {
		goto L34
	}
L34:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v181)+24))
	v195 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v194))))
	v198 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v194)+4)))
	v199 = v195<<(uint(int64(32))%64) | v198
	if base.Ui64(v199) <= base.Ui64(v127) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v206 = v143
	v209 = int32(1)
	goto L29
L36:
	;
	goto L37
L37:
	;
	if base.Ui64(v143-int64(1)) < base.Ui64(v199) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v205 = v143
	goto L40
L39:
	;
	v205 = v199
	goto L40
L40:
	;
	v206 = v205
	v209 = v189
	goto L29
L41:
	;
	v220 = v209 ^ int32(1) | int32(base.Ui32(v185&int32(16))>>(uint(int32(4))%32))
	goto L43
L42:
	;
	v220 = v210
	goto L43
L43:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+16)))
	v224 = int32(6)
	if v185&v224 == v224 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v228 = v221 | int32(64)
	goto L46
L45:
	;
	v228 = v221
	goto L46
L46:
	;
	v229 = int32(1)
	v232 = base.B2i32(v117&int32(2) != int32(0)) | v209
	if v232 != v229 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v436 = int32(0)
	if v220 == v436 {
		goto L109
	} else {
		goto L110
	}
L48:
	;
	v422 = v142
	v424 = v156
	v426 = v229
	v427 = v228
	v430 = v210
	v431 = v167
	v433 = v169
	v434 = v170
	v435 = v172
	goto L47
L49:
	;
	goto L50
L50:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v181)+24))
	v236 = int32(0)
	if v185&int32(8) == v236 {
		v255 = v210
		v257 = v236
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v259 = *(*int32)(unsafe.Add(mBase, _consts[380]))
	if v259 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L52:
	;
	v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v235)+12)))
	if base.Ui32(v241) < base.Ui32(int32(24)) {
		v255 = v210
		v257 = v236
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v235)+14)))
	v250 = base.B2i32(base.Ui32(v241) < base.Ui32(v244)) & base.B2i32(base.Ui32(v244) < base.Ui32(int32(8193)))
	if v250 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v251 = v244 - v241
	goto L56
L55:
	;
	v251 = int32(0)
	goto L56
L56:
	;
	if v250 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v253 = v241
	goto L59
L58:
	;
	v253 = int32(0)
	goto L59
L59:
	;
	v255 = v251
	v257 = v253
	goto L51
L60:
	;
	v338 = v181 + int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v338
	v341 = v255 & int32(65535)
	v343 = base.B2i32(v341 != int32(0))
	if v209 != 0 {
		goto L88
	} else {
		goto L89
	}
L61:
	;
	v264 = int32(0)
	v332 = v264
	v334 = v264
	v335 = v255 & int32(65535)
	goto L60
L62:
	;
	goto L63
L63:
	;
	v266 = int32(0)
	v269 = v255 & int32(65535)
	if v269 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if v257 != 0 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v287 = v235
	v288 = v266
	goto L66
L66:
	;
	switch v259 - int32(1) {
	case 0:
		goto L75
	case 1:
		goto L77
	case 2:
		goto L76
	default:
		v332 = int32(0)
		v334 = v266
		v335 = v269
		goto L60
	}
L67:
	;
	v277 = v269 + v257
	v280 = int32(8192) - v277
	if v280 != 0 {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	v272 = F__emscripten_memcpy_bulkmem(m, v42-int32(-64), v235, v257)
	mBase = m.M
	goto L70
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	v287 = v42 - int32(-64)
	v288 = int32(2)
	goto L66
L72:
	;
	v281 = F__emscripten_memcpy_bulkmem(m, v42-int32(-64)+v257, v235+v277, v280)
	mBase = m.M
	goto L74
L73:
	;
	goto L74
L74:
	;
	goto L71
L75:
	;
	v321 = int32(8192) - v269
	v324 = F_pglz_compress(m, v287, v321, v181-int32(-64), v55)
	mBase = m.M
	v325 = int32(0)
	v330 = base.B2i32(v324+v288 < v321) & base.B2i32(v325 <= v324)
	if v330 != 0 {
		goto L85
	} else {
		goto L86
	}
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L78
	} else {
		goto L82
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	return int64(0)
L79:
	;
	F_errmsg_internal(m, int32(448005), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(512304), int32(984), int32(329792))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errmsg_internal(m, int32(447821), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L78
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(512304), int32(995), int32(329792))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L78
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v331 = v324
	goto L87
L86:
	;
	v331 = v325
	goto L87
L87:
	;
	v332 = v330
	v334 = v331
	v335 = v269
	goto L60
L88:
	;
	v346 = v343 | int32(2)
	goto L90
L89:
	;
	v346 = v343
	goto L90
L90:
	;
	if v332 != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v410 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v181+v409))) = v406
	v422 = v142 + base.I64_extend_i32_u(v407)&int64(65535)
	v424 = v405
	v426 = v332 ^ v410
	v427 = v228 | int32(16)
	v430 = v255
	v431 = v407
	v433 = v257
	v434 = v408
	v435 = v172 + v410
	goto L47
L92:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _consts[380]))
	switch v348 - int32(1) {
	case 0:
		goto L96
	case 1:
		goto L98
	case 2:
		goto L97
	default:
		v379 = v346
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	if v341 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181)+44)) = v181 - int32(-64)
	v405 = v338
	v406 = v334 & int32(65535)
	v407 = v334
	v408 = v379
	v409 = int32(48)
	goto L91
L96:
	;
	v379 = v346 | int32(4)
	goto L95
L97:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L78
	} else {
		goto L102
	}
L98:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L78
	} else {
		goto L99
	}
L99:
	;
	F_errmsg_internal(m, int32(448005), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L78
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(512304), int32(739), int32(406488))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L78
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	F_errmsg_internal(m, int32(447821), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L78
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(512304), int32(747), int32(406488))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L78
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181)+44)) = v235
	v389 = int32(8192)
	v405 = v338
	v406 = v389
	v407 = v389
	v408 = v346
	v409 = int32(48)
	goto L91
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181)+48)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v181)+44)) = v235
	v395 = v181 + int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v181)+40)) = v395
	v397 = v335 + v257
	*(*int32)(unsafe.Add(mBase, uint32(v181)+56)) = v235 + v397
	v400 = int32(8192)
	v405 = v395
	v406 = v400 - v397
	v407 = v400 - v255
	v408 = v346
	v409 = int32(60)
	goto L91
L108:
	;
	if v161 == int32(0) {
		v466 = v450
		v467 = v436
		goto L112
	} else {
		goto L113
	}
L109:
	;
	v448 = v422
	v449 = int32(0)
	v450 = v427
	v451 = v424
	goto L108
L110:
	;
	goto L111
L111:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v181)+28))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v181)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v424))) = v441
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v181)+36))
	v448 = v422 + base.I64_extend_i32_u(v440)
	v449 = v440
	v450 = v427 | int32(32)
	v451 = v447
	goto L108
L112:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v155)+2)) = uint16(v449)
	*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)) = uint8(v466)
	*(*uint8)(unsafe.Add(mBase, uint32(v155))) = uint8(v153)
	if v232 == int32(0) {
		v488 = v155 + int32(4)
		goto L119
	} else {
		goto L120
	}
L113:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
	if v454 != v455 {
		v466 = v450
		v467 = v436
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v181)+8))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v161)+8))
	if v457 != v458 {
		v466 = v450
		v467 = v436
		goto L112
	} else {
		goto L115
	}
L115:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	v464 = base.B2i32(v462 == v463)
	if v462 == v463 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v465 = v450 | int32(-128)
	goto L118
L117:
	;
	v465 = v450
	goto L118
L118:
	;
	v466 = v465
	v467 = v464
	goto L112
L119:
	;
	if v467 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v155)+8)) = uint8(v434)
	*(*uint16)(unsafe.Add(mBase, uint32(v155)+6)) = uint16(v433)
	*(*uint16)(unsafe.Add(mBase, uint32(v155)+4)) = uint16(v431)
	if base.B2i32(v430&int32(65535) == int32(0))|v426 != 0 {
		v488 = v155 + int32(9)
		goto L119
	} else {
		goto L121
	}
L121:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v155)+9)) = uint16(v430)
	v488 = v155 + int32(11)
	goto L119
L122:
	;
	v491 = *(*int64)(unsafe.Add(mBase, uint32(v181)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v488))) = v491
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v488)+8)) = v493
	v497 = v488 + int32(12)
	goto L124
L123:
	;
	v497 = v488
	goto L124
L124:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v497))) = v498
	v503 = *(*int32)(unsafe.Add(mBase, _consts[372]))
	v505 = *(*int32)(unsafe.Add(mBase, _consts[379]))
	v506 = v448
	v507 = v206
	v509 = v497 + int32(4)
	v510 = v451
	v512 = v505
	v514 = v503
	v515 = v181
	v519 = v431
	v521 = v433
	v522 = v434
	v523 = v435
	goto L28
L125:
	;
	goto L25
L126:
	;
	v586 = int32(0)
	v588 = *(*int32)(unsafe.Add(mBase, _consts[75]))
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v588)+78)))
	if v589 != 0 {
		v603 = v586
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v575 = int32(*(*uint16)(unsafe.Add(mBase, _consts[381])))
	if v575 == int32(0) {
		v585 = v544
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v578 = int32(253)
	*(*uint8)(unsafe.Add(mBase, uint32(v544))) = uint8(v578)
	v581 = int32(*(*uint16)(unsafe.Add(mBase, _consts[381])))
	*(*uint16)(unsafe.Add(mBase, uint32(v544)+1)) = uint16(v581)
	v585 = v544 + int32(3)
	goto L126
L129:
	;
	if v603 != 0 {
		goto L134
	} else {
		goto L135
	}
L130:
	;
	v591 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	if v591 < int32(2) {
		v603 = v586
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v588)+20))
	if v594 != int32(2) {
		v603 = v586
		goto L129
	} else {
		goto L132
	}
L132:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v588)+28))
	if v597 < int32(2) {
		v603 = v586
		goto L129
	} else {
		goto L133
	}
L133:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v588)))
	v603 = base.B2i32(v600 != int32(0))
	goto L129
L134:
	;
	v605 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	v606 = int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v585))) = uint8(v606)
	*(*int32)(unsafe.Add(mBase, uint32(v585)+1)) = v605
	v612 = v585 + int32(5)
	goto L136
L135:
	;
	v612 = v585
	goto L136
L136:
	;
	v614 = *(*int64)(unsafe.Add(mBase, _consts[370]))
	if v614 != int64(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	if base.Ui64(int64(256)) <= base.Ui64(v614) {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v642 = v531
	v643 = v612
	v644 = v545
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v644))) = int32(0)
	v649 = *(*int32)(unsafe.Add(mBase, _consts[375]))
	v650 = v643 - v649
	*(*int32)(unsafe.Add(mBase, _consts[382])) = v650
	v653 = v642 + base.I64_extend_i32_u(v650)
	v655 = int32(24)
	v659 = m.Env.Pgmem_crc32c(m, int32(-1), v649+v655, v650-v655)
	mBase = m.M
	v661 = *(*int32)(unsafe.Add(mBase, _consts[377]))
	if v661 != 0 {
		goto L145
	} else {
		goto L146
	}
L140:
	;
	v635 = *(*int32)(unsafe.Add(mBase, _consts[383]))
	*(*int32)(unsafe.Add(mBase, uint32(v545))) = v635
	v638 = *(*int64)(unsafe.Add(mBase, _consts[370]))
	v641 = *(*int32)(unsafe.Add(mBase, _consts[371]))
	v642 = v638 + v531
	v643 = v633
	v644 = v641
	goto L139
L141:
	;
	if base.Ui64(int64(4294967296)) <= base.Ui64(v614) {
		goto L4
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v626 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v612))) = uint8(v626)
	v629 = *(*int64)(unsafe.Add(mBase, _consts[370]))
	*(*uint8)(unsafe.Add(mBase, uint32(v612)+1)) = uint8(v629)
	v633 = v612 + int32(2)
	goto L140
L144:
	;
	v621 = int32(254)
	*(*uint8)(unsafe.Add(mBase, uint32(v612))) = uint8(v621)
	*(*uint32)(unsafe.Add(mBase, uint32(v612)+1)) = uint32(v614)
	v633 = v612 + int32(5)
	goto L140
L145:
	;
	v676 = v661
	v677 = v659
	goto L148
L146:
	;
	v720 = v659
	goto L147
L147:
	;
	if base.Ui64(int64(1069547521)) <= base.Ui64(v653) {
		goto L3
	} else {
		goto L151
	}
L148:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v676)+4))
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v676)+8))
	v703 = m.Env.Pgmem_crc32c(m, v677, v701, v702)
	mBase = m.M
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v676)))
	if v704 != 0 {
		v676 = v704
		v677 = v703
		goto L148
	} else {
		goto L150
	}
L149:
	;
	v720 = v703
	goto L147
L150:
	;
	goto L149
L151:
	;
	v747 = *(*int32)(unsafe.Add(mBase, _consts[75]))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v747)))
	goto L152
L152:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v108)+17)) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, uint32(v108)+16)) = uint8(v117)
	*(*uint32)(unsafe.Add(mBase, uint32(v108))) = uint32(v653)
	*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v108)+20)) = v720
	*(*int64)(unsafe.Add(mBase, uint32(v108)+8)) = int64(0)
	v757 = int32(*(*uint8)(unsafe.Add(mBase, _consts[360])))
	v758 = int32(0)
	v761 = m.G0
	v763 = v761 - int32(16)
	m.G0 = v763
	v767 = *(*int32)(unsafe.Add(mBase, _consts[376]))
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767)+17)))
	if v768 != 0 {
		v781 = v758
		v782 = int32(1)
		goto L154
	} else {
		goto L155
	}
L153:
	;
	if v1871 == int64(0) {
		v83 = v556
		v85 = v558
		v86 = v559
		goto L14
	} else {
		goto L372
	}
L154:
	;
	v784 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v786 = int32(*(*uint8)(unsafe.Add(mBase, _consts[276])))
	v788 = *(*int32)(unsafe.Add(mBase, _consts[319]))
	if v788 < int32(0) {
		goto L167
	} else {
		goto L168
	}
L155:
	;
	v769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767)+16)))
	v771 = v769 & int32(240)
	if v771 == int32(64) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v781 = int32(1)
	v782 = int32(0)
	goto L154
L157:
	;
	goto L158
L158:
	;
	if v771 == int32(224) {
		v781 = v758
		v782 = int32(0)
		goto L154
	} else {
		goto L159
	}
L159:
	;
	v781 = v758
	v782 = int32(1)
	goto L154
L160:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1914 = m.ExcPending
	if v1914 != 0 {
		goto L78
	} else {
		goto L368
	}
L161:
	;
	m.G0 = v763 + int32(16)
	goto L153
L162:
	;
	F_WALInsertLockRelease(m)
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L78
	} else {
		goto L337
	}
L163:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v767)+20))
	v1378 = m.Env.Pgmem_crc32c(m, v1376, v767, int32(20))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v767)+20)) = v1378 ^ int32(-1)
	v1385 = v1361 & int64(8191)
	if v1385 == int64(0) {
		goto L283
	} else {
		goto L284
	}
L164:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v767)))
	v1245 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1245)))
	*(*int32)(unsafe.Add(mBase, uint32(v1245))) = int32(1)
	if v1246 != 0 {
		goto L263
	} else {
		goto L264
	}
L165:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L78
	} else {
		goto L260
	}
L166:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v784)+308))
	v808 = int32(4543684)
	v810 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	*(*int32)(unsafe.Add(mBase, _consts[17])) = v810 + int32(1)
	if v782 != 0 {
		goto L175
	} else {
		goto L176
	}
L167:
	;
	v792 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
	if v792 == int32(1) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	goto L169
L169:
	;
	if v788 == int32(0) {
		goto L165
	} else {
		goto L174
	}
L170:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v784)+316))
	v798 = base.B2i32(v796 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v798)
	if v796 != int32(2) {
		goto L165
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	*(*int32)(unsafe.Add(mBase, _consts[319])) = int32(1)
	goto L166
L173:
	;
	goto L172
L174:
	;
	goto L166
L175:
	;
	v815 = *(*int32)(unsafe.Add(mBase, _consts[320]))
	if v815 == int32(-1) {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	goto L177
L177:
	;
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L78
	} else {
		goto L195
	}
L178:
	;
	v820 = *(*int32)(unsafe.Add(mBase, _consts[112]))
	v822 = base.I32_rem_s(v820, int32(8))
	*(*int32)(unsafe.Add(mBase, _consts[320])) = v822
	v824 = v822
	goto L180
L179:
	;
	v824 = v815
	goto L180
L180:
	;
	*(*int32)(unsafe.Add(mBase, _consts[321])) = v824
	v828 = *(*int32)(unsafe.Add(mBase, _consts[322]))
	v833 = F_LWLockAcquire(m, v828+v824<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L78
	} else {
		goto L181
	}
L181:
	;
	if v833 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v837 = int32(4154408)
	v839 = *(*int32)(unsafe.Add(mBase, _consts[320]))
	v843 = base.I32_rem_s(v839+int32(1), int32(8))
	*(*int32)(unsafe.Add(mBase, _consts[320])) = v843
	goto L184
L183:
	;
	goto L184
L184:
	;
	v846 = *(*int64)(unsafe.Add(mBase, _consts[277]))
	v847 = *(*int64)(unsafe.Add(mBase, uint32(v784)+152))
	if v846 != v847 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	*(*int64)(unsafe.Add(mBase, _consts[277])) = v847
	v851 = v847
	goto L187
L186:
	;
	v851 = v846
	goto L187
L187:
	;
	v852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784)+160)))
	if v852 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	if v786&int32(1)&base.B2i32(base.Ui64(v851) <= base.Ui64(v532-int64(1))) != 0 {
		goto L164
	} else {
		goto L193
	}
L189:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v784)+164))
	v858 = base.B2i32(int32(0) < v856)
	*(*uint8)(unsafe.Add(mBase, _consts[276])) = uint8(v858)
	if int32(0) < v856 {
		goto L188
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v861 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[276])) = uint8(v861)
	goto L188
L192:
	;
	goto L164
L193:
	;
	F_WALInsertLockRelease(m)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L78
	} else {
		goto L194
	}
L194:
	;
	v872 = int32(4543684)
	v874 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	*(*int32)(unsafe.Add(mBase, _consts[17])) = v874 - int32(1)
	v1871 = int64(0)
	goto L161
L195:
	;
	if v781 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v767)))
	v885 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v885)))
	*(*int32)(unsafe.Add(mBase, uint32(v885))) = int32(1)
	if v886 != 0 {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	goto L198
L198:
	;
	v1012 = int32(8)
	v1013 = v763 + v1012
	v1017 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v1017)))
	*(*int32)(unsafe.Add(mBase, uint32(v1017))) = int32(1)
	if v1018 != 0 {
		goto L221
	} else {
		goto L222
	}
L199:
	;
	F_s_lock(m, v885, int32(518546), int32(1134), int32(276002))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L78
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v885))) = int32(0)
	v896 = *(*int64)(unsafe.Add(mBase, uint32(v885)+16))
	v897 = *(*int64)(unsafe.Add(mBase, uint32(v885)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v885)+16)) = v897
	v904 = v897 + base.I64_extend_i32_s((v883+int32(7))&int32(-8))
	*(*int64)(unsafe.Add(mBase, uint32(v885)+8)) = v904
	v908 = int64(*(*int32)(unsafe.Add(mBase, _consts[316])))
	v909 = base.I64_div_u_s(v897, v908)
	v911 = v897 - v909*v908
	if base.Ui64(v911) <= base.Ui64(int64(8151)) {
		goto L205
	} else {
		goto L206
	}
L202:
	;
	goto L201
L203:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v763)+8)) = v935
	v938 = int64(*(*int32)(unsafe.Add(mBase, _consts[316])))
	v939 = base.I64_div_u_s(v904, v938)
	v941 = v904 - v939*v938
	if base.Ui64(v941) <= base.Ui64(int64(8151)) {
		goto L209
	} else {
		goto L210
	}
L204:
	;
	v931 = int64(*(*int32)(unsafe.Add(mBase, _consts[191])))
	v935 = v909*v931 + v929&int64(4294967295)
	goto L203
L205:
	;
	v929 = v911 + int64(40)
	goto L204
L206:
	;
	goto L207
L207:
	;
	v917 = v911 - int64(8152)
	v918 = int64(8168)
	v919 = base.I64_div_u_s(v917, v918)
	v929 = v917 - v919*v918 + v919<<(uint(int64(13))%64) + int64(8216)
	goto L204
L208:
	;
	v971 = int64(*(*int32)(unsafe.Add(mBase, _consts[191])))
	v975 = v939*v971 + v969&int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v763))) = v975
	v979 = int64(*(*int32)(unsafe.Add(mBase, _consts[316])))
	v980 = base.I64_div_u_s(v896, v979)
	v982 = v896 - v980*v979
	if base.Ui64(v982) <= base.Ui64(int64(8151)) {
		goto L218
	} else {
		goto L219
	}
L209:
	;
	v944 = int64(0)
	if v941 == v944 {
		goto L212
	} else {
		goto L213
	}
L210:
	;
	goto L211
L211:
	;
	v951 = v941 - int64(8152)
	v952 = int64(8168)
	v953 = base.I64_div_u_s(v951, v952)
	v955 = v953 << (uint(int64(13)) % 64)
	v960 = v951 - v953*v952
	if v960 == int64(0) {
		v969 = v955 - int64(-8192)
		goto L208
	} else {
		goto L215
	}
L212:
	;
	v949 = v944
	goto L214
L213:
	;
	v949 = v941 + int64(40)
	goto L214
L214:
	;
	v969 = v949
	goto L208
L215:
	;
	v969 = v960 + v955 + int64(8216)
	goto L208
L216:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v767)+8)) = v980*v1002 + v1000&int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v784)+152)) = v935
	*(*int64)(unsafe.Add(mBase, _consts[277])) = v935
	v1359 = v975
	v1361 = v935
	goto L163
L217:
	;
	v1002 = int64(*(*int32)(unsafe.Add(mBase, _consts[191])))
	goto L216
L218:
	;
	v1000 = v982 + int64(40)
	goto L217
L219:
	;
	goto L220
L220:
	;
	v988 = v982 - int64(8152)
	v989 = int64(8168)
	v990 = base.I64_div_u_s(v988, v989)
	v1000 = v988 - v990*v989 + v990<<(uint(int64(13))%64) + int64(8216)
	goto L217
L221:
	;
	F_s_lock(m, v1017, int32(518546), int32(1183), int32(337131))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L78
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v1026 = *(*int64)(unsafe.Add(mBase, uint32(v1017)+8))
	v1028 = int64(*(*int32)(unsafe.Add(mBase, _consts[316])))
	v1029 = base.I64_div_u_s(v1026, v1028)
	v1031 = v1026 - v1029*v1028
	if base.Ui64(v1031) <= base.Ui64(int64(8151)) {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	goto L223
L225:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v1062 = base.I64_extend_i32_s(v1061)
	v1063 = v1029 * v1062
	v1066 = v1063 + v1059&int64(4294967295)
	v1068 = v1061 - int32(1)
	v1069 = base.I64_extend_i32_s(v1068)
	v1070 = v1066 & v1069
	if v1070 == int64(0) {
		goto L234
	} else {
		goto L235
	}
L226:
	;
	v1034 = int64(0)
	if v1031 == v1034 {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	goto L228
L228:
	;
	v1041 = v1031 - int64(8152)
	v1042 = int64(8168)
	v1043 = base.I64_div_u_s(v1041, v1042)
	v1045 = v1043 << (uint(int64(13)) % 64)
	v1050 = v1041 - v1043*v1042
	if v1050 == int64(0) {
		v1059 = v1045 - int64(-8192)
		goto L225
	} else {
		goto L232
	}
L229:
	;
	v1039 = v1034
	goto L231
L230:
	;
	v1039 = v1031 + int64(40)
	goto L231
L231:
	;
	v1059 = v1039
	goto L225
L232:
	;
	v1059 = v1050 + v1045 + int64(8216)
	goto L225
L233:
	;
	if v1070 == int64(0) {
		v1724 = int32(0)
		goto L162
	} else {
		goto L259
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1017))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1013))) = v1066
	*(*int64)(unsafe.Add(mBase, uint32(v763))) = v1066
	goto L233
L235:
	;
	goto L236
L236:
	;
	v1078 = v1026 + int64(24)
	v1079 = *(*int64)(unsafe.Add(mBase, uint32(v1017)+16))
	if base.Ui64(v1031) <= base.Ui64(int64(8151)) {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v1097 = v1031 + int64(40)
	goto L239
L238:
	;
	v1085 = v1031 - int64(8152)
	v1086 = int64(8168)
	v1087 = base.I64_div_u_s(v1085, v1086)
	v1097 = v1085 - v1087*v1086 + v1087<<(uint(int64(13))%64) + int64(8216)
	goto L239
L239:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1013))) = v1097&int64(4294967295) + v1063
	v1102 = base.I64_div_u_s(v1078, v1028)
	v1104 = v1078 - v1102*v1028
	if base.Ui64(v1104) <= base.Ui64(int64(8151)) {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	v1136 = v1062*v1102 + v1132&int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v763))) = v1136
	v1139 = v1068 & base.I32_wrap_i64(v1136)
	if v1139 == int32(0) {
		v1180 = v1078
		goto L248
	} else {
		goto L249
	}
L241:
	;
	v1107 = int64(0)
	if v1104 == v1107 {
		goto L244
	} else {
		goto L245
	}
L242:
	;
	goto L243
L243:
	;
	v1114 = v1104 - int64(8152)
	v1115 = int64(8168)
	v1116 = base.I64_div_u_s(v1114, v1115)
	v1118 = v1116 << (uint(int64(13)) % 64)
	v1123 = v1114 - v1116*v1115
	if v1123 == int64(0) {
		v1132 = v1118 - int64(-8192)
		goto L240
	} else {
		goto L247
	}
L244:
	;
	v1112 = v1107
	goto L246
L245:
	;
	v1112 = v1104 + int64(40)
	goto L246
L246:
	;
	v1132 = v1112
	goto L240
L247:
	;
	v1132 = v1123 + v1118 + int64(8216)
	goto L240
L248:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1017)+16)) = v1026
	*(*int64)(unsafe.Add(mBase, uint32(v1017)+8)) = v1180
	*(*int32)(unsafe.Add(mBase, uint32(v1017))) = int32(0)
	v1187 = base.I64_div_u_s(v1079, v1028)
	v1189 = v1079 - v1187*v1028
	if base.Ui64(v1189) <= base.Ui64(int64(8151)) {
		goto L256
	} else {
		goto L257
	}
L249:
	;
	v1144 = v1136 + base.I64_extend_i32_u(v1061-v1139)
	*(*int64)(unsafe.Add(mBase, uint32(v763))) = v1144
	v1146 = base.I64_div_u_s(v1144, v1062)
	v1149 = base.I32_wrap_i64(v1144) & int32(8191)
	v1150 = v1144 & v1069
	if v1150&int64(35184372080640) == int64(0) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v1155 = v1146 * v1028
	if v1149 == int32(0) {
		v1180 = v1155
		goto L248
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	v1173 = v1146*v1028 + (int64(base.Ui64(v1150)>>(uint(int64(13))%64))*int64(8168)+int64(4294959128))&int64(4294967288) + int64(8152)
	if v1149 == int32(0) {
		v1180 = v1173
		goto L248
	} else {
		goto L254
	}
L253:
	;
	v1180 = v1155 + base.I64_extend_i32_u(v1149-int32(40))
	goto L248
L254:
	;
	v1180 = v1173 + base.I64_extend_i32_u(v1149-int32(24))
	goto L248
L255:
	;
	v1209 = int64(*(*int32)(unsafe.Add(mBase, _consts[191])))
	*(*int64)(unsafe.Add(mBase, uint32(v767+v1012))) = v1187*v1209 + v1207&int64(4294967295)
	goto L233
L256:
	;
	v1207 = v1189 + int64(40)
	goto L255
L257:
	;
	goto L258
L258:
	;
	v1195 = v1189 - int64(8152)
	v1196 = int64(8168)
	v1197 = base.I64_div_u_s(v1195, v1196)
	v1207 = v1195 - v1197*v1196 + v1197<<(uint(int64(13))%64) + int64(8216)
	goto L255
L259:
	;
	v1226 = *(*int64)(unsafe.Add(mBase, uint32(v763)))
	v1227 = *(*int64)(unsafe.Add(mBase, uint32(v763)+8))
	v1359 = v1226
	v1361 = v1227
	goto L163
L260:
	;
	F_errmsg_internal(m, int32(14649), int32(0))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L78
	} else {
		goto L261
	}
L261:
	;
	F_errfinish(m, int32(518546), int32(779), int32(438559))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L78
	} else {
		goto L262
	}
L262:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L263:
	;
	F_s_lock(m, v1245, int32(518546), int32(1134), int32(276002))
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L78
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1245))) = int32(0)
	v1256 = *(*int64)(unsafe.Add(mBase, uint32(v1245)+16))
	v1257 = *(*int64)(unsafe.Add(mBase, uint32(v1245)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1245)+16)) = v1257
	v1264 = v1257 + base.I64_extend_i32_s((v1243+int32(7))&int32(-8))
	*(*int64)(unsafe.Add(mBase, uint32(v1245)+8)) = v1264
	v1267 = int64(*(*int32)(unsafe.Add(mBase, _consts[316])))
	v1268 = base.I64_div_u_s(v1257, v1267)
	v1270 = v1257 - v1268*v1267
	if base.Ui64(v1270) <= base.Ui64(int64(8151)) {
		goto L268
	} else {
		goto L269
	}
L266:
	;
	goto L265
L267:
	;
	v1290 = int64(*(*int32)(unsafe.Add(mBase, _consts[191])))
	v1294 = v1268*v1290 + v1288&int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v763)+8)) = v1294
	v1296 = base.I64_div_u_s(v1264, v1267)
	v1298 = v1264 - v1296*v1267
	if base.Ui64(v1298) <= base.Ui64(int64(8151)) {
		goto L272
	} else {
		goto L273
	}
L268:
	;
	v1288 = v1270 + int64(40)
	goto L267
L269:
	;
	goto L270
L270:
	;
	v1276 = v1270 - int64(8152)
	v1277 = int64(8168)
	v1278 = base.I64_div_u_s(v1276, v1277)
	v1288 = v1276 - v1278*v1277 + v1278<<(uint(int64(13))%64) + int64(8216)
	goto L267
L271:
	;
	v1330 = v1290*v1296 + v1326&int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v763))) = v1330
	v1332 = base.I64_div_u_s(v1256, v1267)
	v1334 = v1256 - v1332*v1267
	if base.Ui64(v1334) <= base.Ui64(int64(8151)) {
		goto L280
	} else {
		goto L281
	}
L272:
	;
	v1301 = int64(0)
	if v1298 == v1301 {
		goto L275
	} else {
		goto L276
	}
L273:
	;
	goto L274
L274:
	;
	v1308 = v1298 - int64(8152)
	v1309 = int64(8168)
	v1310 = base.I64_div_u_s(v1308, v1309)
	v1312 = v1310 << (uint(int64(13)) % 64)
	v1317 = v1308 - v1310*v1309
	if v1317 == int64(0) {
		v1326 = v1312 - int64(-8192)
		goto L271
	} else {
		goto L278
	}
L275:
	;
	v1306 = v1301
	goto L277
L276:
	;
	v1306 = v1298 + int64(40)
	goto L277
L277:
	;
	v1326 = v1306
	goto L271
L278:
	;
	v1326 = v1317 + v1312 + int64(8216)
	goto L271
L279:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v767)+8)) = v1290*v1332 + v1352&int64(4294967295)
	v1359 = v1330
	v1361 = v1294
	goto L163
L280:
	;
	v1352 = v1334 + int64(40)
	goto L279
L281:
	;
	goto L282
L282:
	;
	v1340 = v1334 - int64(8152)
	v1341 = int64(8168)
	v1342 = base.I64_div_u_s(v1340, v1341)
	v1352 = v1340 - v1342*v1341 + v1342<<(uint(int64(13))%64) + int64(8216)
	goto L279
L283:
	;
	v1390 = int32(0)
	goto L285
L284:
	;
	v1390 = int32(8192) - base.I32_wrap_i64(v1385)
	goto L285
L285:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v767)))
	v1392 = F_GetXLogBuffer(m, v1361, v807)
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L78
	} else {
		goto L286
	}
L286:
	;
	v1396 = v1361
	v1407 = v1390
	v1410 = v1392
	v1419 = v758
	v1420 = int32(4444880)
	goto L287
L287:
	;
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+4))
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+8))
	if v1407 < v1434 {
		goto L289
	} else {
		goto L290
	}
L288:
	;
	if v781 == int32(0) {
		goto L315
	} else {
		goto L316
	}
L289:
	;
	v1438 = v1396
	v1449 = v1407
	v1452 = v1410
	v1453 = v1434
	v1457 = v1433
	v1461 = v1419
	goto L292
L290:
	;
	v1518 = v1396
	v1529 = v1407
	v1532 = v1410
	v1533 = v1434
	v1537 = v1433
	v1541 = v1419
	goto L291
L291:
	;
	v1556 = v1529 - v1533
	if v1533 != 0 {
		goto L310
	} else {
		goto L311
	}
L292:
	;
	if v1449 != 0 {
		goto L295
	} else {
		goto L296
	}
L293:
	;
	v1518 = v1507
	v1529 = v1514
	v1532 = v1500
	v1533 = v1488
	v1537 = v1501
	v1541 = v1481
	goto L291
L294:
	;
	v1478 = v1438 + base.I64_extend_i32_s(v1449)
	v1479 = F_GetXLogBuffer(m, v1478, v807)
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L78
	} else {
		goto L298
	}
L295:
	;
	v1475 = F__emscripten_memcpy_bulkmem(m, v1452, v1457, v1449)
	mBase = m.M
	goto L297
L296:
	;
	goto L297
L297:
	;
	goto L294
L298:
	;
	v1481 = v1449 + v1461
	*(*int32)(unsafe.Add(mBase, uint32(v1479)+16)) = v1391 - v1481
	v1484 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1479)+2)))
	v1485 = int32(1)
	v1486 = v1484 | v1485
	*(*uint16)(unsafe.Add(mBase, uint32(v1479)+2)) = uint16(v1486)
	v1488 = v1453 - v1449
	v1492 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v1498 = base.B2i32(v1478&base.I64_extend_i32_s(v1492-v1485) == int64(0))
	if v1478&base.I64_extend_i32_s(v1492-v1485) == int64(0) {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1499 = int32(40)
	goto L301
L300:
	;
	v1499 = int32(24)
	goto L301
L301:
	;
	v1500 = v1479 + v1499
	v1501 = v1449 + v1457
	if v1478&base.I64_extend_i32_s(v1492-v1485) == int64(0) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1506 = int64(40)
	goto L304
L303:
	;
	v1506 = int64(24)
	goto L304
L304:
	;
	v1507 = v1506 + v1478
	v1509 = v1507 & int64(8191)
	if v1509 == int64(0) {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1514 = int32(0)
	goto L307
L306:
	;
	v1514 = int32(8192) - base.I32_wrap_i64(v1509)
	goto L307
L307:
	;
	if v1514 < v1488 {
		v1438 = v1507
		v1449 = v1514
		v1452 = v1500
		v1453 = v1488
		v1457 = v1501
		v1461 = v1481
		goto L292
	} else {
		goto L308
	}
L308:
	;
	goto L293
L309:
	;
	v1561 = v1518 + base.I64_extend_i32_s(v1533)
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1420)))
	if v1562 != 0 {
		v1396 = v1561
		v1407 = v1556
		v1410 = v1558 + v1533
		v1419 = v1533 + v1541
		v1420 = v1562
		goto L287
	} else {
		goto L313
	}
L310:
	;
	v1557 = F__emscripten_memcpy_bulkmem(m, v1532, v1537, v1533)
	mBase = m.M
	v1558 = v1557
	goto L312
L311:
	;
	v1558 = v1532
	goto L312
L312:
	;
	goto L309
L313:
	;
	goto L288
L314:
	;
	if v1655 != v1359 {
		goto L160
	} else {
		goto L332
	}
L315:
	;
	v1655 = (v1561 + int64(7)) & int64(-8)
	goto L314
L316:
	;
	v1566 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	if v1561&base.I64_extend_i32_s(v1566-int32(1)) == int64(0) {
		goto L315
	} else {
		goto L317
	}
L317:
	;
	v1574 = v1561 + base.I64_extend_i32_s(v1556)
	if base.Ui64(v1359) <= base.Ui64(v1574) {
		v1655 = v1574
		goto L314
	} else {
		goto L318
	}
L318:
	;
	v1578 = v1574
	goto L319
L319:
	;
	v1615 = F_GetXLogBuffer(m, v1578, v807)
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L78
	} else {
		goto L322
	}
L320:
	;
	v1655 = v1647
	goto L314
L321:
	;
	v1647 = v1578 - int64(-8192)
	if base.Ui64(v1647) < base.Ui64(v1359) {
		v1578 = v1647
		goto L319
	} else {
		goto L331
	}
L322:
	;
	if v1615&int32(3) == int32(0) {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v1622 = v1615 + int32(24)
	if base.Ui32(v1622) <= base.Ui32(v1615) {
		goto L321
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	v1638 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1615))) = v1638
	*(*int64)(unsafe.Add(mBase, uint32(v1615)+16)) = v1638
	*(*int64)(unsafe.Add(mBase, uint32(v1615)+8)) = v1638
	goto L321
L326:
	;
	v1628 = v1615 + int32(4)
	if base.Ui32(v1628) < base.Ui32(v1622) {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v1630 = v1622
	goto L329
L328:
	;
	v1630 = v1628
	goto L329
L329:
	;
	v1637 = F__emscripten_memset_bulkmem(m, v1615, base.I32_extend8_s(int32(0)), (v1615^int32(-1)+v1630)&int32(-4)+int32(4))
	mBase = m.M
	goto L330
L330:
	;
	goto L321
L331:
	;
	goto L320
L332:
	;
	v1693 = int32(1)
	if v757&int32(2) != 0 {
		v1724 = v1693
		goto L162
	} else {
		goto L333
	}
L333:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, _consts[322]))
	v1700 = *(*int32)(unsafe.Add(mBase, _consts[321]))
	v1702 = int32(*(*uint8)(unsafe.Add(mBase, _consts[384])))
	if v1702 != 0 {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v1703 = int32(0)
	goto L336
L335:
	;
	v1703 = v1700
	goto L336
L336:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1697+v1703<<(uint(int32(7))%32))+24)) = v1361
	v1724 = v1693
	goto L162
L337:
	;
	v1749 = int32(4543684)
	v1751 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	*(*int32)(unsafe.Add(mBase, _consts[17])) = v1751 - int32(1)
	v1756 = *(*int32)(unsafe.Add(mBase, _consts[75]))
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1756)))
	if v1757 != 0 {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v1758 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1756)+70)) = uint8(v1758)
	goto L340
L339:
	;
	goto L340
L340:
	;
	if v603 != 0 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v1761 = *(*int32)(unsafe.Add(mBase, _consts[75]))
	v1762 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1761)+78)) = uint8(v1762)
	goto L343
L342:
	;
	goto L343
L343:
	;
	v1764 = *(*int64)(unsafe.Add(mBase, uint32(v763)))
	v1765 = *(*int64)(unsafe.Add(mBase, uint32(v763)+8))
	if base.Ui64(int64(8192)) <= base.Ui64(v1764^v1765) {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v1770 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1770)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v1770)+440)) = int32(1)
	if v1771 != 0 {
		goto L347
	} else {
		goto L348
	}
L345:
	;
	goto L346
L346:
	;
	if v781 != 0 {
		goto L356
	} else {
		goto L357
	}
L347:
	;
	v1775 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_s_lock(m, v1775+int32(440), int32(518546), int32(968), int32(438559))
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L78
	} else {
		goto L350
	}
L348:
	;
	goto L349
L349:
	;
	v1783 = *(*int64)(unsafe.Add(mBase, uint32(v763)))
	v1785 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v1786 = *(*int64)(unsafe.Add(mBase, uint32(v1785)+184))
	if base.Ui64(v1786) < base.Ui64(v1783) {
		goto L351
	} else {
		goto L352
	}
L350:
	;
	goto L349
L351:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1785)+184)) = v1783
	goto L353
L352:
	;
	goto L353
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1785)+440)) = int32(0)
	v1791 = *(*int64)(unsafe.Add(mBase, uint32(v1785)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v1785)+280)) = v1791
	*(*int64)(unsafe.Add(mBase, _consts[189])) = v1791
	v1796 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v1797 = *(*int64)(unsafe.Add(mBase, uint32(v1796)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v1796)+272)) = v1797
	*(*int64)(unsafe.Add(mBase, _consts[190])) = v1797
	goto L346
L354:
	;
	*(*int64)(unsafe.Add(mBase, _consts[385])) = v1804
	*(*int64)(unsafe.Add(mBase, _consts[324])) = v1807
	v1871 = v1804
	goto L161
L355:
	;
	v1844 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v767))))
	v1846 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[154])) = uint8(v1846)
	v1848 = int32(4447432)
	v1850 = *(*int64)(unsafe.Add(mBase, _consts[71]))
	*(*int64)(unsafe.Add(mBase, _consts[71])) = v1844 + v1850
	v1853 = int32(4447416)
	v1855 = *(*int64)(unsafe.Add(mBase, _consts[72]))
	*(*int64)(unsafe.Add(mBase, _consts[72])) = v1855 + int64(1)
	v1859 = int32(4447424)
	v1861 = *(*int64)(unsafe.Add(mBase, _consts[73]))
	*(*int64)(unsafe.Add(mBase, _consts[73])) = v1861 + base.I64_extend_i32_s(v561)
	v1871 = v1842
	goto L161
L356:
	;
	v1804 = *(*int64)(unsafe.Add(mBase, uint32(v763)))
	F_XLogFlush(m, v1804)
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L78
	} else {
		goto L359
	}
L357:
	;
	goto L358
L358:
	;
	v1835 = *(*int64)(unsafe.Add(mBase, uint32(v763)+8))
	*(*int64)(unsafe.Add(mBase, _consts[324])) = v1835
	v1838 = *(*int64)(unsafe.Add(mBase, uint32(v763)))
	*(*int64)(unsafe.Add(mBase, _consts[385])) = v1838
	if v1724 == int32(0) {
		v1871 = v1838
		goto L161
	} else {
		goto L367
	}
L359:
	;
	v1807 = *(*int64)(unsafe.Add(mBase, uint32(v763)+8))
	if v1724 == int32(0) {
		goto L354
	} else {
		goto L360
	}
L360:
	;
	v1811 = v1807 + int64(24)
	if base.Ui64(int64(8192)) <= base.Ui64(v1811^v1807) {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	if v1811&base.I64_extend_i32_s(v1818-int32(1)^int32(8191)) == int64(0) {
		goto L364
	} else {
		goto L365
	}
L362:
	;
	v1829 = v1811
	goto L363
L363:
	;
	*(*int64)(unsafe.Add(mBase, _consts[385])) = v1829
	*(*int64)(unsafe.Add(mBase, _consts[324])) = v1807
	v1842 = v1829
	goto L355
L364:
	;
	v1827 = int64(64)
	goto L366
L365:
	;
	v1827 = int64(48)
	goto L366
L366:
	;
	v1829 = v1827 + v1807
	goto L363
L367:
	;
	v1842 = v1838
	goto L355
L368:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v1917 = m.ExcPending
	if v1917 != 0 {
		goto L78
	} else {
		goto L369
	}
L369:
	;
	F_errmsg_internal(m, int32(292548), int32(0))
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L78
	} else {
		goto L370
	}
L370:
	;
	F_errfinish(m, int32(518546), int32(1367), int32(555049))
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L78
	} else {
		goto L371
	}
L371:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L372:
	;
	goto L15
L373:
	;
	v1935 = v1931 & int32(7)
	v1937 = *(*int32)(unsafe.Add(mBase, _consts[379]))
	if base.Ui32(int32(8)) <= base.Ui32(v1931) {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v1956 = int32(0)
	v1957 = v1929
	goto L377
L375:
	;
	v2055 = v1929
	goto L376
L376:
	;
	if v1935 == int32(0) {
		v2409 = v1871
		goto L1
	} else {
		goto L380
	}
L377:
	;
	v1982 = int32(8260)
	v1985 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1937+v1957*v1982))) = uint8(v1985)
	*(*uint8)(unsafe.Add(mBase, uint32(v1937+(v1957|int32(1))*v1982))) = uint8(v1985)
	*(*uint8)(unsafe.Add(mBase, uint32(v1937+(v1957|int32(2))*v1982))) = uint8(v1985)
	*(*uint8)(unsafe.Add(mBase, uint32(v1937+(v1957|int32(3))*v1982))) = uint8(v1985)
	*(*uint8)(unsafe.Add(mBase, uint32(v1937+(v1957|int32(4))*v1982))) = uint8(v1985)
	*(*uint8)(unsafe.Add(mBase, uint32(v1937+(v1957|int32(5))*v1982))) = uint8(v1985)
	*(*uint8)(unsafe.Add(mBase, uint32(v1937+(v1957|int32(6))*v1982))) = uint8(v1985)
	*(*uint8)(unsafe.Add(mBase, uint32(v1937+(v1957|int32(7))*v1982))) = uint8(v1985)
	v2036 = int32(8)
	v2037 = v1957 + v2036
	v2039 = v1956 + v2036
	if v2039 != v1931&int32(2147483640) {
		v1956 = v2039
		v1957 = v2037
		goto L377
	} else {
		goto L379
	}
L378:
	;
	v2055 = v2037
	goto L376
L379:
	;
	goto L378
L380:
	;
	v2096 = int32(0)
	v2097 = v2055
	goto L381
L381:
	;
	v2125 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1937+v2097*int32(8260)))) = uint8(v2125)
	v2127 = int32(1)
	v2130 = v2096 + v2127
	if v2130 != v1935 {
		v2096 = v2130
		v2097 = v2097 + v2127
		goto L381
	} else {
		goto L383
	}
L382:
	;
	v2409 = v1871
	goto L1
L383:
	;
	goto L382
L384:
	;
	F_errmsg_internal(m, int32(471578), int32(0))
	mBase = m.M
	v2139 = m.ExcPending
	if v2139 != 0 {
		goto L78
	} else {
		goto L385
	}
L385:
	;
	F_errfinish(m, int32(512304), int32(480), int32(87609))
	mBase = m.M
	v2144 = m.ExcPending
	if v2144 != 0 {
		goto L78
	} else {
		goto L386
	}
L386:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L387:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+48)) = l1
	F_errmsg_internal(m, int32(531700), v42+int32(48))
	mBase = m.M
	v2154 = m.ExcPending
	if v2154 != 0 {
		goto L78
	} else {
		goto L388
	}
L388:
	;
	F_errfinish(m, int32(512304), int32(489), int32(87609))
	mBase = m.M
	v2159 = m.ExcPending
	if v2159 != 0 {
		goto L78
	} else {
		goto L389
	}
L389:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L390:
	;
	F_errmsg_internal(m, int32(526413), int32(0))
	mBase = m.M
	v2167 = m.ExcPending
	if v2167 != 0 {
		goto L78
	} else {
		goto L391
	}
L391:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+40)) = int32(-1)
	v2171 = *(*int64)(unsafe.Add(mBase, _consts[370]))
	*(*int64)(unsafe.Add(mBase, uint32(v42)+32)) = v2171
	F_errdetail_internal(m, int32(619738), v42+int32(32))
	mBase = m.M
	v2177 = m.ExcPending
	if v2177 != 0 {
		goto L78
	} else {
		goto L392
	}
L392:
	;
	F_errfinish(m, int32(512304), int32(874), int32(406488))
	mBase = m.M
	v2182 = m.ExcPending
	if v2182 != 0 {
		goto L78
	} else {
		goto L393
	}
L393:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L394:
	;
	F_errmsg_internal(m, int32(438512), int32(0))
	mBase = m.M
	v2190 = m.ExcPending
	if v2190 != 0 {
		goto L78
	} else {
		goto L395
	}
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v1
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = int32(1069547520)
	*(*int64)(unsafe.Add(mBase, uint32(v42))) = v653
	F_errdetail_internal(m, int32(600132), v42)
	mBase = m.M
	v2198 = m.ExcPending
	if v2198 != 0 {
		goto L78
	} else {
		goto L396
	}
L396:
	;
	F_errfinish(m, int32(512304), int32(919), int32(406488))
	mBase = m.M
	v2203 = m.ExcPending
	if v2203 != 0 {
		goto L78
	} else {
		goto L397
	}
L397:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L398:
	;
	v2210 = v2206 & int32(7)
	v2212 = *(*int32)(unsafe.Add(mBase, _consts[379]))
	if base.Ui32(int32(8)) <= base.Ui32(v2206) {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	v2231 = int32(0)
	v2232 = v14
	goto L402
L400:
	;
	v2330 = v14
	goto L401
L401:
	;
	if v2210 == int32(0) {
		v2409 = v2204
		goto L1
	} else {
		goto L405
	}
L402:
	;
	v2257 = int32(8260)
	v2260 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2212+v2232*v2257))) = uint8(v2260)
	*(*uint8)(unsafe.Add(mBase, uint32(v2212+(v2232|int32(1))*v2257))) = uint8(v2260)
	*(*uint8)(unsafe.Add(mBase, uint32(v2212+(v2232|int32(2))*v2257))) = uint8(v2260)
	*(*uint8)(unsafe.Add(mBase, uint32(v2212+(v2232|int32(3))*v2257))) = uint8(v2260)
	*(*uint8)(unsafe.Add(mBase, uint32(v2212+(v2232|int32(4))*v2257))) = uint8(v2260)
	*(*uint8)(unsafe.Add(mBase, uint32(v2212+(v2232|int32(5))*v2257))) = uint8(v2260)
	*(*uint8)(unsafe.Add(mBase, uint32(v2212+(v2232|int32(6))*v2257))) = uint8(v2260)
	*(*uint8)(unsafe.Add(mBase, uint32(v2212+(v2232|int32(7))*v2257))) = uint8(v2260)
	v2311 = int32(8)
	v2312 = v2232 + v2311
	v2314 = v2231 + v2311
	if v2314 != v2206&int32(2147483640) {
		v2231 = v2314
		v2232 = v2312
		goto L402
	} else {
		goto L404
	}
L403:
	;
	v2330 = v2312
	goto L401
L404:
	;
	goto L403
L405:
	;
	v2371 = int32(0)
	v2372 = v2330
	goto L406
L406:
	;
	v2400 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2212+v2372*int32(8260)))) = uint8(v2400)
	v2402 = int32(1)
	v2405 = v2371 + v2402
	if v2405 != v2210 {
		v2371 = v2405
		v2372 = v2372 + v2402
		goto L406
	} else {
		goto L408
	}
L407:
	;
	v2409 = v2204
	goto L1
L408:
	;
	goto L407
}
func F_XLogPageRead(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v75 int64
	_ = v75
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v116 int64
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int64
	_ = v124
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v173 int64
	_ = v173
	var v174 int32
	_ = v174
	var v175 int64
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v216 int32
	_ = v216
	var v231 int32
	_ = v231
	var v242 int32
	_ = v242
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v327 int64
	_ = v327
	var v329 int64
	_ = v329
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v342 int64
	_ = v342
	var v345 int32
	_ = v345
	var v348 int64
	_ = v348
	var v356 int64
	_ = v356
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int64
	_ = v401
	var v402 int64
	_ = v402
	var v412 int32
	_ = v412
	var v414 int64
	_ = v414
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v489 int32
	_ = v489
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v510 int32
	_ = v510
	var v512 int64
	_ = v512
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v539 int32
	_ = v539
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v573 int64
	_ = v573
	var v577 int64
	_ = v577
	var v578 int64
	_ = v578
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v659 int32
	_ = v659
	var v663 int64
	_ = v663
	var v664 int64
	_ = v664
	var v665 int64
	_ = v665
	var v668 int64
	_ = v668
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v733 int64
	_ = v733
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int64
	_ = v745
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int64
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v796 int32
	_ = v796
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v921 int32
	_ = v921
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1034 int64
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1042 int64
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1052 int32
	_ = v1052
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1071 int32
	_ = v1071
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1094 int64
	_ = v1094
	var v1100 int64
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1109 int64
	_ = v1109
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1120 int64
	_ = v1120
	var v1121 int64
	_ = v1121
	var v1129 int64
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1169 int64
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1181 int32
	_ = v1181
	var v1198 int32
	_ = v1198
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int64
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int64
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1251 int32
	_ = v1251
	var v1252 int64
	_ = v1252
	var v1257 int32
	_ = v1257
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1292 int32
	_ = v1292
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1325 int32
	_ = v1325
	var v1333 int32
	_ = v1333
	var v1336 int32
	_ = v1336
	var v1340 int64
	_ = v1340
	var v1346 int32
	_ = v1346
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1363 int32
	_ = v1363
	var v1368 int32
	_ = v1368
	var v1371 int64
	_ = v1371
	var v1377 int32
	_ = v1377
	var v1394 int32
	_ = v1394
	var v1418 int32
	_ = v1418
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1427 int64
	_ = v1427
	var v1428 int64
	_ = v1428
	var v1432 int64
	_ = v1432
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1444 int64
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1449 int32
	_ = v1449
	var v1451 int32
	_ = v1451
	var v1457 int32
	_ = v1457
	var v1458 int64
	_ = v1458
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1470 int64
	_ = v1470
	var v1471 int64
	_ = v1471
	var v1475 int64
	_ = v1475
	var v1527 int32
	_ = v1527
	var v1528 int64
	_ = v1528
	var v1532 int32
	_ = v1532
	var v1550 int32
	_ = v1550
	var v1551 int64
	_ = v1551
	var v1555 int32
	_ = v1555
	var v1572 int32
	_ = v1572
	var v1573 int64
	_ = v1573
	var v1578 int32
	_ = v1578
	var v1579 int64
	_ = v1579
	var v1584 int32
	_ = v1584
	var v1593 int32
	_ = v1593
	var v1596 int64
	_ = v1596
	var v1599 int64
	_ = v1599
	var v1600 int64
	_ = v1600
	var v1601 int64
	_ = v1601
	var v1604 int64
	_ = v1604
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1622 int32
	_ = v1622
	var v1627 int64
	_ = v1627
	var v1629 int64
	_ = v1629
	var v1634 int32
	_ = v1634
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1642 int32
	_ = v1642
	var v1645 int64
	_ = v1645
	var v1648 int32
	_ = v1648
	var v1657 int32
	_ = v1657
	var v1662 int32
	_ = v1662
	var v1667 int64
	_ = v1667
	var v1669 int64
	_ = v1669
	var v1674 int32
	_ = v1674
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1683 int32
	_ = v1683
	var v1688 int64
	_ = v1688
	var v1692 int32
	_ = v1692
	var v1701 int32
	_ = v1701
	var v1706 int32
	_ = v1706
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1717 int32
	_ = v1717
	var v1718 int64
	_ = v1718
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1730 int64
	_ = v1730
	var v1731 int64
	_ = v1731
	var v1735 int64
	_ = v1735
	var v1787 int32
	_ = v1787
	var v1788 int64
	_ = v1788
	var v1792 int32
	_ = v1792
	var v1810 int32
	_ = v1810
	var v1811 int64
	_ = v1811
	var v1815 int32
	_ = v1815
	var v1832 int32
	_ = v1832
	var v1833 int64
	_ = v1833
	var v1838 int32
	_ = v1838
	var v1839 int64
	_ = v1839
	var v1844 int32
	_ = v1844
	var v1853 int32
	_ = v1853
	var v1856 int32
	_ = v1856
	var v1860 int64
	_ = v1860
	var v1861 int64
	_ = v1861
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1873 int32
	_ = v1873
	var v1877 int64
	_ = v1877
	var v1879 int64
	_ = v1879
	var v1884 int32
	_ = v1884
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1891 int32
	_ = v1891
	var v1897 int32
	_ = v1897
	var v1902 int32
	_ = v1902
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1916 int32
	_ = v1916
	var v1919 int32
	_ = v1919
	var v1922 int32
	_ = v1922
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1931 int32
	_ = v1931
	var v1937 int32
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1946 int32
	_ = v1946
	v4 = l3
	v6 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(1216)
	m.G0 = v33
	v35 = base.I32_wrap_i64(l1)
	v37 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v40 = v35 & (v37 - int32(1))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v45 = *(*int32)(unsafe.Add(mBase, _consts[314]))
	if v6 <= v45 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v33 + int32(1216)
	return v1946
L2:
	;
	v147 = v141
	v154 = v138
	goto L25
L3:
	;
	v138 = v134
	v141 = int32(0)
	goto L2
L4:
	;
	v116 = base.I64_div_u_s(l1, v50)
	*(*int64)(unsafe.Add(mBase, _consts[386])) = v116
	v118 = int32(8192)
	v120 = *(*int32)(unsafe.Add(mBase, _consts[387]))
	if v120 == int32(3) {
		goto L20
	} else {
		goto L21
	}
L5:
	;
	v49 = *(*int64)(unsafe.Add(mBase, _consts[386]))
	v50 = base.I64_extend_i32_s(v37)
	v51 = base.I64_div_u_s(l1, v50)
	if v49 == v51 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	v111 = v37
	goto L7
L7:
	;
	v113 = base.I64_div_u_s(l1, base.I64_extend_i32_s(v111))
	*(*int64)(unsafe.Add(mBase, _consts[386])) = v113
	v134 = v6
	goto L3
L8:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, _consts[227])))
	if v54 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v98 = int32(4154428)
	v99 = *(*int32)(unsafe.Add(mBase, _consts[314]))
	v100 = F_close(m, v99)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[314])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[387])) = int32(0)
	v108 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v111 = v108
	goto L7
L10:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, _consts[104])))
	if v58 != int32(1) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[217]))
	v67 = *(*int64)(unsafe.Add(mBase, _consts[277]))
	v69 = int64(*(*int32)(unsafe.Add(mBase, _consts[191])))
	v70 = base.I64_div_u_s(v67, v69)
	goto L12
L12:
	;
	if base.B2i32(base.Ui64(base.I64_extend_i32_s(v62-int32(1))+v70) <= base.Ui64(v49)) == int32(0) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v75 = F_GetRedoRecPtr(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v80 = *(*int64)(unsafe.Add(mBase, _consts[386]))
	v82 = *(*int32)(unsafe.Add(mBase, _consts[217]))
	v87 = *(*int64)(unsafe.Add(mBase, _consts[277]))
	v89 = int64(*(*int32)(unsafe.Add(mBase, _consts[191])))
	v90 = base.I64_div_u_s(v87, v89)
	goto L16
L16:
	;
	if base.B2i32(base.Ui64(base.I64_extend_i32_s(v82-int32(1))+v90) <= base.Ui64(v80)) == int32(0) {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	F_RequestCheckpoint(m, int32(128))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L9
L19:
	;
	v138 = v118
	v141 = int32(2)
	goto L2
L20:
	;
	v124 = *(*int64)(unsafe.Add(mBase, _consts[388]))
	if base.Ui64(l1+base.I64_extend_i32_s(l2)) <= base.Ui64(v124) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v138 = v118
	v141 = int32(1)
	goto L2
L23:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1257)))
	if v128 == int32(0) {
		v134 = v118
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v1946 = int32(-2)
	goto L1
L25:
	;
	switch v147 {
	case 0:
		goto L30
	case 1:
		goto L27
	default:
		goto L29
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, _consts[389])) = v40
	*(*int32)(unsafe.Add(mBase, _consts[390])) = v154
	v1418 = int32(*(*uint8)(unsafe.Add(mBase, _consts[195])))
	v1421 = m.G0
	v1423 = v1421 - int32(16)
	m.G0 = v1423
	if v1418 != 0 {
		goto L340
	} else {
		goto L341
	}
L28:
	;
	v147 = int32(1)
	v154 = v1394
	goto L25
L29:
	;
	v1371 = *(*int64)(unsafe.Add(mBase, _consts[388]))
	if base.Ui64(int64(8191)) < base.Ui64(v1371^l1) {
		v1394 = int32(8192)
		goto L28
	} else {
		goto L338
	}
L30:
	;
	v173 = l1 + base.I64_extend_i32_s(l2)
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1257)))
	v175 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+5)))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, _consts[253])))
	if v181 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v204 = int32(1)
	v205 = v174 & v204
	v216 = v199
	v231 = int32(0)
	goto L43
L32:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _consts[391]))
	if v185 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v196 = int32(2)
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, _consts[391])) = v196
	v199 = v196
	goto L31
L35:
	;
	if v185 != int32(3) {
		v199 = v185
		goto L31
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v193 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[392])) = uint8(v193)
	v196 = int32(1)
	goto L34
L38:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, _consts[254])))
	if v189&int32(1) != 0 {
		v199 = v185
		goto L31
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L14
	} else {
		goto L335
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L14
	} else {
		goto L332
	}
L42:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, _consts[314]))
	if int32(0) <= v1316 {
		goto L329
	} else {
		goto L330
	}
L43:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, _consts[392])))
	if v242 != 0 {
		goto L49
	} else {
		goto L50
	}
L44:
	;
	F_XLogShutdownWalRcv(m)
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L14
	} else {
		goto L328
	}
L45:
	;
	goto L44
L46:
	;
	v480 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[392])) = uint8(v480)
	if base.Ui32(int32(2)) <= base.Ui32(v474-int32(1)) {
		goto L117
	} else {
		goto L118
	}
L47:
	;
	if v216 == v436 {
		v474 = v216
		v476 = v434
		goto L46
	} else {
		goto L102
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, _consts[391])) = v428
	v434 = v429
	v436 = v428
	goto L47
L49:
	;
	if v205 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v419 = int32(0)
	if v216 != int32(2) {
		v474 = v216
		v476 = v419
		goto L46
	} else {
		goto L100
	}
L52:
	;
	v1946 = int32(-2)
	goto L1
L53:
	;
	goto L54
L54:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v216-int32(1)) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v279 = F_WalRcvStreaming(m)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L14
	} else {
		goto L68
	}
L56:
	;
	if v216 == int32(3) {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, _consts[254])))
	if v266 != int32(1) {
		goto L42
	} else {
		goto L63
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L14
	} else {
		goto L60
	}
L60:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _consts[391]))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v255
	F_errmsg_internal(m, int32(496190), v33)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L14
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(511477), int32(3765), int32(411928))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L14
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	v269 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L14
	} else {
		goto L64
	}
L64:
	;
	if v269 != 0 {
		goto L45
	} else {
		goto L65
	}
L65:
	;
	v271 = int32(1)
	v273 = int32(*(*uint8)(unsafe.Add(mBase, _consts[254])))
	if v273&v271 == int32(0) {
		goto L42
	} else {
		goto L66
	}
L66:
	;
	v428 = int32(3)
	v429 = v271
	goto L48
L67:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _consts[237]))
	if v301 != int32(1) {
		goto L75
	} else {
		goto L76
	}
L68:
	;
	if v279 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	F_XLogShutdownWalRcv(m)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L14
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v288 = F_LWLockAcquire(m, v284+int32(1152), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L14
	} else {
		goto L73
	}
L72:
	;
	goto L67
L73:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v292 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v291)+320)) = uint8(v292)
	v295 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v295+int32(1152))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L14
	} else {
		goto L74
	}
L74:
	;
	goto L67
L75:
	;
	v313 = m.G0
	v314 = int32(16)
	v315 = v313 - v314
	m.G0 = v315
	F___gettimeofday(m, v315)
	mBase = m.M
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v315)))
	v319 = int64(*(*int32)(unsafe.Add(mBase, uint32(v315)+8)))
	m.G0 = v315 + v314
	v327 = v319 + v318*int64(1000000) - int64(946684800000000)
	goto L79
L76:
	;
	v304 = F_rescanLatestTimeLine(m, v176, v175)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L14
	} else {
		goto L77
	}
L77:
	;
	if v304 == int32(0) {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v428 = int32(1)
	v429 = int32(0)
	goto L48
L79:
	;
	v329 = *(*int64)(unsafe.Add(mBase, _consts[393]))
	v331 = *(*int32)(unsafe.Add(mBase, _consts[394]))
	goto L80
L80:
	;
	if base.B2i32(base.I64_extend_i32_s(v331)*int64(1000) <= v327-v329) == int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _consts[394]))
	v342 = *(*int64)(unsafe.Add(mBase, _consts[393]))
	if v327 <= v342 {
		v359 = int32(0)
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v414 = v327
	goto L83
L83:
	;
	*(*int64)(unsafe.Add(mBase, _consts[393])) = v414
	v428 = int32(1)
	v429 = int32(0)
	goto L48
L84:
	;
	v363 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L14
	} else {
		goto L89
	}
L85:
	;
	goto L84
L86:
	;
	v345 = int32(2147483647)
	v348 = v327 - v342
	if base.B2i32(int64(0) < v342)^base.B2i32(v348 < v327) != 0 {
		v359 = v345
		goto L85
	} else {
		goto L87
	}
L87:
	;
	if int64(2147483646000) < v348 {
		v359 = v345
		goto L85
	} else {
		goto L88
	}
L88:
	;
	v356 = base.I64_div_s(v348+int64(999), int64(1000))
	v359 = base.I32_wrap_i64(v356)
	goto L85
L89:
	;
	if v363 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = base.I32_wrap_i64(v173)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = base.I32_wrap_i64(int64(base.Ui64(v173) >> (uint(int64(32)) % 64)))
	F_errmsg_internal(m, int32(534135), v33+int32(176))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L14
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	F_KnownAssignedTransactionIdsIdleMaintenance(m)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L14
	} else {
		goto L95
	}
L93:
	;
	F_errfinish(m, int32(511477), int32(3744), int32(411928))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L14
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _consts[240]))
	v385 = F_WaitLatch(m, v380+int32(4), int32(41), v340-v359, int32(150994948))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L14
	} else {
		goto L96
	}
L96:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _consts[240]))
	*(*int32)(unsafe.Add(mBase, uint32(v388+int32(4)))) = int32(0)
	goto L97
L97:
	;
	v396 = m.G0
	v397 = int32(16)
	v398 = v396 - v397
	m.G0 = v398
	F___gettimeofday(m, v398)
	mBase = m.M
	v401 = *(*int64)(unsafe.Add(mBase, uint32(v398)))
	v402 = int64(*(*int32)(unsafe.Add(mBase, uint32(v398)+8)))
	m.G0 = v398 + v397
	goto L98
L98:
	;
	F_ProcessStartupProcInterrupts(m)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L14
	} else {
		goto L99
	}
L99:
	;
	v414 = v402 + v401*int64(1000000) - int64(946684800000000)
	goto L83
L100:
	;
	v422 = int32(1)
	v425 = int32(*(*uint8)(unsafe.Add(mBase, _consts[253])))
	if v425 != v422 {
		v434 = v419
		v436 = int32(2)
		goto L47
	} else {
		goto L101
	}
L101:
	;
	v428 = v422
	v429 = v419
	goto L48
L102:
	;
	v440 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L14
	} else {
		goto L103
	}
L103:
	;
	if v440 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v445 = int32(*(*uint8)(unsafe.Add(mBase, _consts[392])))
	if v445 != 0 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	v473 = *(*int32)(unsafe.Add(mBase, _consts[391]))
	v474 = v473
	v476 = v434
	goto L46
L107:
	;
	v446 = int32(378643)
	goto L109
L108:
	;
	v446 = int32(137542)
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v446
	v448 = int32(2)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v216<<(uint(v448)%32))+uint32(_consts[395])))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v452
	v455 = *(*int32)(unsafe.Add(mBase, _consts[391]))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v455<<(uint(v448)%32))+uint32(_consts[395])))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v460
	F_errmsg_internal(m, int32(190839), v33+int32(160))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L14
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(511477), int32(3782), int32(411928))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L14
	} else {
		goto L111
	}
L111:
	;
	goto L106
L112:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, _consts[240]))
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1303)+80))
	if v1304 != 0 {
		goto L323
	} else {
		goto L324
	}
L113:
	;
	v1217 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L14
	} else {
		goto L308
	}
L114:
	;
	v147 = int32(2)
	v154 = v1198
	goto L25
L115:
	;
	v717 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[396])) = uint8(v717)
	if v715&int32(1) == v717 {
		goto L176
	} else {
		goto L177
	}
L116:
	;
	F_XLogShutdownWalRcv(m)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L14
	} else {
		goto L173
	}
L117:
	;
	if v474 != int32(3) {
		goto L40
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v499 = *(*int32)(unsafe.Add(mBase, _consts[314]))
	if int32(0) <= v499 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v489 = int32(*(*uint8)(unsafe.Add(mBase, _consts[396])))
	if (v476|(v489^int32(-1)))&int32(1) == int32(0) {
		goto L116
	} else {
		goto L121
	}
L121:
	;
	v715 = v489 | v476
	goto L115
L122:
	;
	v502 = F_close(m, v499)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[314])) = int32(-1)
	goto L124
L123:
	;
	goto L124
L124:
	;
	if v178&v204 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, _consts[397])) = int32(0)
	goto L127
L126:
	;
	goto L127
L127:
	;
	v510 = *(*int32)(unsafe.Add(mBase, _consts[391]))
	v512 = *(*int64)(unsafe.Add(mBase, _consts[386]))
	v514 = *(*int32)(unsafe.Add(mBase, _consts[256]))
	if v514 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v659 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v659
	v663 = int64(*(*int32)(unsafe.Add(mBase, _consts[191])))
	v664 = base.I64_div_u_s(int64(4294967296), v663)
	v665 = base.I64_div_u_s(v512, v664)
	*(*uint32)(unsafe.Add(mBase, uint32(v33)+52)) = uint32(v665)
	v668 = v512 - v664*v665
	*(*uint32)(unsafe.Add(mBase, uint32(v33)+56)) = uint32(v668)
	v676 = F_pg_snprintf(m, v33+int32(192), int32(1024), int32(531342), v33+int32(48))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L14
	} else {
		goto L165
	}
L129:
	;
	v518 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	v519 = F_readTimeLineHistory(m, v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L14
	} else {
		goto L132
	}
L130:
	;
	v523 = v514
	goto L131
L131:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)+4))
	if v524 <= int32(0) {
		goto L128
	} else {
		goto L134
	}
L132:
	;
	if v519 == int32(0) {
		goto L128
	} else {
		goto L133
	}
L133:
	;
	v523 = v519
	goto L131
L134:
	;
	if v510 != int32(1) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v530 = v510
	goto L137
L136:
	;
	v530 = int32(0)
	goto L137
L137:
	;
	v539 = int32(0)
	goto L138
L138:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v523)+12))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v564+v539<<(uint(int32(2))%32))))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	v571 = *(*int32)(unsafe.Add(mBase, _consts[397]))
	if base.Ui32(v569) < base.Ui32(v571) {
		goto L128
	} else {
		goto L140
	}
L139:
	;
	goto L128
L140:
	;
	v573 = *(*int64)(unsafe.Add(mBase, uint32(v568)+8))
	if v573 != int64(0) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v625 = v539 + int32(1)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v523)+4))
	if v625 < v626 {
		v539 = v625
		goto L138
	} else {
		goto L164
	}
L142:
	;
	v577 = int64(*(*int32)(unsafe.Add(mBase, _consts[191])))
	v578 = base.I64_div_u_s(v573, v577)
	if base.Ui64(v512) < base.Ui64(v578) {
		goto L141
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	if base.Ui32(int32(1)) < base.Ui32(v530) {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	goto L144
L146:
	;
	*(*int32)(unsafe.Add(mBase, _consts[314])) = v615
	v618 = int32(8192)
	v620 = *(*int32)(unsafe.Add(mBase, _consts[387]))
	if v620 == int32(3) {
		v1198 = v618
		goto L114
	} else {
		goto L163
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, _consts[256])) = v523
	v615 = v612
	goto L146
L148:
	;
	if v530&int32(1) != 0 {
		goto L141
	} else {
		goto L159
	}
L149:
	;
	v582 = int32(1)
	v584 = F_XLogFileRead(m, v512, v569, v582, v582)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L14
	} else {
		goto L150
	}
L150:
	;
	if v584 == int32(-1) {
		goto L148
	} else {
		goto L151
	}
L151:
	;
	v590 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L14
	} else {
		goto L152
	}
L152:
	;
	if v590 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	F_errmsg_internal(m, int32(358526), int32(0))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L14
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v602 = *(*int32)(unsafe.Add(mBase, _consts[256]))
	if v602 != 0 {
		v615 = v584
		goto L146
	} else {
		goto L158
	}
L156:
	;
	F_errfinish(m, int32(511477), int32(4384), int32(555655))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L14
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	v612 = v584
	goto L147
L159:
	;
	v606 = F_XLogFileRead(m, v512, v569, int32(2), int32(1))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L14
	} else {
		goto L160
	}
L160:
	;
	if v606 == int32(-1) {
		goto L141
	} else {
		goto L161
	}
L161:
	;
	v611 = *(*int32)(unsafe.Add(mBase, _consts[256]))
	if v611 != 0 {
		v615 = v606
		goto L146
	} else {
		goto L162
	}
L162:
	;
	v612 = v606
	goto L147
L163:
	;
	v1394 = v618
	goto L28
L164:
	;
	goto L139
L165:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(44)
	v683 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L14
	} else {
		goto L166
	}
L166:
	;
	if v683 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L14
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, _consts[314])) = int32(-1)
	v704 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[392])) = uint8(v704)
	v1292 = v231
	goto L112
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v33 + int32(192)
	F_errmsg(m, int32(310478), v33+int32(32))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L14
	} else {
		goto L171
	}
L171:
	;
	F_errfinish(m, int32(511477), int32(4408), int32(555655))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L14
	} else {
		goto L172
	}
L172:
	;
	goto L169
L173:
	;
	v708 = int32(1)
	v710 = *(*int32)(unsafe.Add(mBase, _consts[237]))
	if v710 != v708 {
		v715 = v708
		goto L115
	} else {
		goto L174
	}
L174:
	;
	v713 = F_rescanLatestTimeLine(m, v176, v175)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L14
	} else {
		goto L175
	}
L175:
	;
	v715 = v708
	goto L115
L176:
	;
	v1086 = F_WalRcvStreaming(m)
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L14
	} else {
		goto L284
	}
L177:
	;
	v724 = *(*int32)(unsafe.Add(mBase, _consts[232]))
	if v724 == int32(0) {
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724))))
	if v727 == int32(0) {
		goto L176
	} else {
		goto L179
	}
L179:
	;
	if v177&v204 != 0 {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, _consts[397])) = v743
	v749 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v753 = F_LWLockAcquire(m, v749+int32(1152), int32(0))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L14
	} else {
		goto L189
	}
L181:
	;
	v731 = *(*int32)(unsafe.Add(mBase, _consts[251]))
	v733 = *(*int64)(unsafe.Add(mBase, _consts[252]))
	v743 = v731
	v745 = v733
	goto L180
L182:
	;
	goto L183
L183:
	;
	v735 = *(*int32)(unsafe.Add(mBase, _consts[256]))
	v736 = F_tliOfPointInHistory(m, v4, v735)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L14
	} else {
		goto L184
	}
L184:
	;
	v739 = *(*int32)(unsafe.Add(mBase, _consts[397]))
	if base.Ui32(v736) < base.Ui32(v739) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v742 = v739
	goto L187
L186:
	;
	v742 = int32(0)
	goto L187
L187:
	;
	if v742 != 0 {
		goto L41
	} else {
		goto L188
	}
L188:
	;
	v743 = v736
	v745 = v173
	goto L180
L189:
	;
	v756 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v757 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v756)+320)) = uint8(v757)
	v760 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v760+int32(1152))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L14
	} else {
		goto L190
	}
L190:
	;
	v766 = *(*int32)(unsafe.Add(mBase, _consts[232]))
	v768 = *(*int32)(unsafe.Add(mBase, _consts[398]))
	v770 = int32(*(*uint8)(unsafe.Add(mBase, _consts[399])))
	v772 = *(*int32)(unsafe.Add(mBase, _consts[400]))
	v773 = F___time(m)
	mBase = m.M
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v772)+1456))
	v776 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	*(*int32)(unsafe.Add(mBase, uint32(v772)+1456)) = int32(1)
	if v774 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	F_s_lock(m, v772+int32(1456), int32(514473), int32(263), int32(349610))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L14
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v789 = v772 + int32(104)
	if v766 != 0 {
		goto L196
	} else {
		goto L197
	}
L194:
	;
	goto L193
L195:
	;
	if v768 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L196:
	;
	goto L202
L197:
	;
	goto L198
L198:
	;
	v905 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v789))) = uint8(v905)
	goto L195
L199:
	;
	goto L195
L200:
	;
	v902 = F_strlen(m, v891)
	mBase = m.M
	goto L199
L202:
	;
	goto L203
L203:
	;
	v796 = int32(1023)
	if (v789^v766)&int32(3) != 0 {
		goto L207
	} else {
		goto L208
	}
L204:
	;
	v895 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v892))) = uint8(v895)
	goto L200
L205:
	;
	v876 = v871
	v877 = v872
	v878 = v873
	goto L227
L206:
	;
	if v866 == int32(0) {
		v891 = v864
		v892 = v865
		goto L204
	} else {
		goto L226
	}
L207:
	;
	v864 = v766
	v865 = v789
	v866 = v796
	goto L206
L208:
	;
	goto L209
L209:
	;
	if v766&int32(3) == int32(0) {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	if v833 == int32(0) {
		v891 = v830
		v892 = v831
		goto L204
	} else {
		goto L219
	}
L211:
	;
	v830 = v766
	v831 = v789
	v832 = v796
	v833 = int32(1)
	goto L210
L212:
	;
	goto L213
L213:
	;
	v809 = v766
	v810 = v789
	v811 = v796
	goto L214
L214:
	;
	v813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v809))))
	*(*uint8)(unsafe.Add(mBase, uint32(v810))) = uint8(v813)
	if v813 == int32(0) {
		v871 = v809
		v872 = v810
		v873 = v811
		goto L205
	} else {
		goto L216
	}
L215:
	;
	v830 = v824
	v831 = v818
	v832 = v820
	v833 = v822
	goto L210
L216:
	;
	v817 = int32(1)
	v818 = v810 + v817
	v820 = v811 - v817
	v821 = int32(0)
	v822 = base.B2i32(v820 != v821)
	v824 = v809 + v817
	if v824&int32(3) == v821 {
		v830 = v824
		v831 = v818
		v832 = v820
		v833 = v822
		goto L210
	} else {
		goto L217
	}
L217:
	;
	if v820 != 0 {
		v809 = v824
		v810 = v818
		v811 = v820
		goto L214
	} else {
		goto L218
	}
L218:
	;
	goto L215
L219:
	;
	v836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v830))))
	if v836 == int32(0) {
		v864 = v830
		v865 = v831
		v866 = v832
		goto L206
	} else {
		goto L220
	}
L220:
	;
	if base.Ui32(v832) < base.Ui32(int32(4)) {
		v864 = v830
		v865 = v831
		v866 = v832
		goto L206
	} else {
		goto L221
	}
L221:
	;
	v842 = v830
	v843 = v831
	v844 = v832
	goto L222
L222:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v842)))
	v850 = int32(-2139062144)
	if (int32(16843008)-v847|v847)&v850 != v850 {
		v871 = v842
		v872 = v843
		v873 = v844
		goto L205
	} else {
		goto L224
	}
L223:
	;
	v864 = v858
	v865 = v856
	v866 = v860
	goto L206
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v843))) = v847
	v855 = int32(4)
	v856 = v843 + v855
	v858 = v842 + v855
	v860 = v844 - v855
	if base.Ui32(int32(3)) < base.Ui32(v860) {
		v842 = v858
		v843 = v856
		v844 = v860
		goto L222
	} else {
		goto L225
	}
L225:
	;
	goto L223
L226:
	;
	v871 = v864
	v872 = v865
	v873 = v866
	goto L205
L227:
	;
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v876))))
	*(*uint8)(unsafe.Add(mBase, uint32(v877))) = uint8(v880)
	if v880 == int32(0) {
		v891 = v876
		v892 = v877
		goto L204
	} else {
		goto L229
	}
L228:
	;
	v891 = v887
	v892 = v885
	goto L204
L229:
	;
	v884 = int32(1)
	v885 = v877 + v884
	v887 = v876 + v884
	v889 = v878 - v884
	if v889 != 0 {
		v876 = v887
		v877 = v885
		v878 = v889
		goto L227
	} else {
		goto L230
	}
L230:
	;
	goto L228
L231:
	;
	v1034 = v745 & base.I64_extend_i32_s(int32(0)-v776)
	*(*uint8)(unsafe.Add(mBase, uint32(v772)+1452)) = uint8(v1033)
	*(*int64)(unsafe.Add(mBase, uint32(v772)+24)) = v773
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v772)+8))
	if v1039 != 0 {
		goto L267
	} else {
		goto L268
	}
L232:
	;
	v1031 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v772)+1388)) = uint8(v1031)
	v1033 = v770
	goto L231
L233:
	;
	v910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v768))))
	if v910 == int32(0) {
		goto L232
	} else {
		goto L234
	}
L234:
	;
	v914 = v772 + int32(1388)
	goto L238
L235:
	;
	v1033 = int32(0)
	goto L231
L236:
	;
	v1027 = F_strlen(m, v1016)
	mBase = m.M
	goto L235
L238:
	;
	goto L239
L239:
	;
	v921 = int32(63)
	if (v914^v768)&int32(3) != 0 {
		goto L243
	} else {
		goto L244
	}
L240:
	;
	v1020 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1017))) = uint8(v1020)
	goto L236
L241:
	;
	v1001 = v996
	v1002 = v997
	v1003 = v998
	goto L263
L242:
	;
	if v991 == int32(0) {
		v1016 = v989
		v1017 = v990
		goto L240
	} else {
		goto L262
	}
L243:
	;
	v989 = v768
	v990 = v914
	v991 = v921
	goto L242
L244:
	;
	goto L245
L245:
	;
	if v768&int32(3) == int32(0) {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	if v958 == int32(0) {
		v1016 = v955
		v1017 = v956
		goto L240
	} else {
		goto L255
	}
L247:
	;
	v955 = v768
	v956 = v914
	v957 = v921
	v958 = int32(1)
	goto L246
L248:
	;
	goto L249
L249:
	;
	v934 = v768
	v935 = v914
	v936 = v921
	goto L250
L250:
	;
	v938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v934))))
	*(*uint8)(unsafe.Add(mBase, uint32(v935))) = uint8(v938)
	if v938 == int32(0) {
		v996 = v934
		v997 = v935
		v998 = v936
		goto L241
	} else {
		goto L252
	}
L251:
	;
	v955 = v949
	v956 = v943
	v957 = v945
	v958 = v947
	goto L246
L252:
	;
	v942 = int32(1)
	v943 = v935 + v942
	v945 = v936 - v942
	v946 = int32(0)
	v947 = base.B2i32(v945 != v946)
	v949 = v934 + v942
	if v949&int32(3) == v946 {
		v955 = v949
		v956 = v943
		v957 = v945
		v958 = v947
		goto L246
	} else {
		goto L253
	}
L253:
	;
	if v945 != 0 {
		v934 = v949
		v935 = v943
		v936 = v945
		goto L250
	} else {
		goto L254
	}
L254:
	;
	goto L251
L255:
	;
	v961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955))))
	if v961 == int32(0) {
		v989 = v955
		v990 = v956
		v991 = v957
		goto L242
	} else {
		goto L256
	}
L256:
	;
	if base.Ui32(v957) < base.Ui32(int32(4)) {
		v989 = v955
		v990 = v956
		v991 = v957
		goto L242
	} else {
		goto L257
	}
L257:
	;
	v967 = v955
	v968 = v956
	v969 = v957
	goto L258
L258:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v967)))
	v975 = int32(-2139062144)
	if (int32(16843008)-v972|v972)&v975 != v975 {
		v996 = v967
		v997 = v968
		v998 = v969
		goto L241
	} else {
		goto L260
	}
L259:
	;
	v989 = v983
	v990 = v981
	v991 = v985
	goto L242
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v968))) = v972
	v980 = int32(4)
	v981 = v968 + v980
	v983 = v967 + v980
	v985 = v969 - v980
	if base.Ui32(int32(3)) < base.Ui32(v985) {
		v967 = v983
		v968 = v981
		v969 = v985
		goto L258
	} else {
		goto L261
	}
L261:
	;
	goto L259
L262:
	;
	v996 = v989
	v997 = v990
	v998 = v991
	goto L241
L263:
	;
	v1005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1001))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1002))) = uint8(v1005)
	if v1005 == int32(0) {
		v1016 = v1001
		v1017 = v1002
		goto L240
	} else {
		goto L265
	}
L264:
	;
	v1016 = v1012
	v1017 = v1010
	goto L240
L265:
	;
	v1009 = int32(1)
	v1010 = v1002 + v1009
	v1012 = v1001 + v1009
	v1014 = v1003 - v1009
	if v1014 != 0 {
		v1001 = v1012
		v1002 = v1010
		v1003 = v1014
		goto L263
	} else {
		goto L266
	}
L266:
	;
	goto L264
L267:
	;
	v1040 = int32(4)
	goto L269
L268:
	;
	v1040 = int32(1)
	goto L269
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v772)+8)) = v1040
	v1042 = *(*int64)(unsafe.Add(mBase, uint32(v772)+32))
	if v1042 != int64(0) {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v772)+40)) = v743
	*(*int64)(unsafe.Add(mBase, uint32(v772)+32)) = v1034
	v1052 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v772)+1456)) = v1052
	if v1039 == v1052 {
		goto L276
	} else {
		goto L277
	}
L271:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v772)+56))
	if v1045 == v743 {
		goto L270
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v772)+64)) = v1034
	*(*int32)(unsafe.Add(mBase, uint32(v772)+56)) = v743
	*(*int64)(unsafe.Add(mBase, uint32(v772)+48)) = v1034
	goto L270
L274:
	;
	goto L273
L275:
	;
	*(*int64)(unsafe.Add(mBase, _consts[388])) = int64(0)
	goto L176
L276:
	;
	F_SendPostmasterSignal(m, int32(7))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L14
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v772)))
	if v1059 != int32(-1) {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	goto L275
L280:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, _consts[152]))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1063)))
	F_SetLatch(m, v1064+v1059*int32(640)+int32(20))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L14
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	goto L275
L283:
	;
	goto L282
L284:
	;
	if v1086 == int32(0) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1091 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[392])) = uint8(v1091)
	v1292 = v231
	goto L112
L286:
	;
	goto L287
L287:
	;
	v1094 = *(*int64)(unsafe.Add(mBase, _consts[388]))
	if base.Ui64(v173) < base.Ui64(v1094) {
		goto L290
	} else {
		goto L291
	}
L288:
	;
	v1181 = int32(3)
	*(*int32)(unsafe.Add(mBase, _consts[401])) = v1181
	*(*int32)(unsafe.Add(mBase, _consts[387])) = v1181
	v1198 = v154
	goto L114
L289:
	;
	if v205 == int32(0) {
		goto L113
	} else {
		goto L307
	}
L290:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, _consts[314]))
	if int32(0) <= v1154 {
		goto L288
	} else {
		goto L301
	}
L291:
	;
	v1100 = F_GetWalRcvFlushRecPtr(m, v33+int32(192), int32(4445280))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L14
	} else {
		goto L292
	}
L292:
	;
	*(*int64)(unsafe.Add(mBase, _consts[388])) = v1100
	if base.Ui64(v1100) <= base.Ui64(v173) {
		goto L289
	} else {
		goto L293
	}
L293:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, _consts[402]))
	v1107 = *(*int32)(unsafe.Add(mBase, _consts[397]))
	if v1105 != v1107 {
		goto L289
	} else {
		goto L294
	}
L294:
	;
	v1109 = *(*int64)(unsafe.Add(mBase, uint32(v33)+192))
	if base.Ui64(v173) < base.Ui64(v1109) {
		goto L290
	} else {
		goto L295
	}
L295:
	;
	v1115 = m.G0
	v1116 = int32(16)
	v1117 = v1115 - v1116
	m.G0 = v1117
	F___gettimeofday(m, v1117)
	mBase = m.M
	v1120 = *(*int64)(unsafe.Add(mBase, uint32(v1117)))
	v1121 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1117)+8)))
	m.G0 = v1117 + v1116
	v1129 = v1121 + v1120*int64(1000000) - int64(946684800000000)
	goto L296
L296:
	;
	*(*int64)(unsafe.Add(mBase, _consts[294])) = v1129
	v1132 = *(*int32)(unsafe.Add(mBase, _consts[240]))
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+96)) = int32(1)
	if v1133 != 0 {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, _consts[240]))
	F_s_lock(m, v1137+int32(96), int32(511477), int32(4658), int32(391419))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L14
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, _consts[240]))
	*(*int32)(unsafe.Add(mBase, uint32(v1146)+96)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1146)+72)) = v1129
	goto L290
L300:
	;
	goto L299
L301:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, _consts[256]))
	if v1158 == int32(0) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	v1164 = F_readTimeLineHistory(m, v1163)
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L14
	} else {
		goto L305
	}
L303:
	;
	goto L304
L304:
	;
	v1169 = *(*int64)(unsafe.Add(mBase, _consts[386]))
	v1171 = *(*int32)(unsafe.Add(mBase, _consts[402]))
	v1174 = F_XLogFileRead(m, v1169, v1171, int32(3), int32(0))
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L14
	} else {
		goto L306
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, _consts[256])) = v1164
	goto L304
L306:
	;
	*(*int32)(unsafe.Add(mBase, _consts[314])) = v1174
	v1292 = v231
	goto L112
L307:
	;
	v1946 = int32(-2)
	goto L1
L308:
	;
	if v1217 != 0 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v1220 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[392])) = uint8(v1220)
	v1292 = v231
	goto L112
L310:
	;
	goto L311
L311:
	;
	if v231 == int32(0) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	F_WalRcvForceReply(m)
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L14
	} else {
		goto L315
	}
L313:
	;
	goto L314
L314:
	;
	F_KnownAssignedTransactionIdsIdleMaintenance(m)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L14
	} else {
		goto L316
	}
L315:
	;
	goto L314
L316:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, _consts[245]))
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1229)))
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+124))
	if v1235 != 0 {
		goto L318
	} else {
		goto L319
	}
L317:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, _consts[240]))
	v1263 = F_WaitLatch(m, v1257+int32(4), int32(33), int32(-1), int32(83886090))
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L14
	} else {
		goto L321
	}
L318:
	;
	v1236 = *(*int64)(unsafe.Add(mBase, uint32(v1235)+16))
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+120))
	v1238 = *(*int64)(unsafe.Add(mBase, uint32(v1237)+16))
	v1241 = base.I32_wrap_i64(v1236 - v1238)
	goto L320
L319:
	;
	v1241 = int32(0)
	goto L320
L320:
	;
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1229)+112))
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v1242)+16))
	v1245 = *(*int32)(unsafe.Add(mBase, _consts[243]))
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1242)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1245)+64)) = v1246
	*(*int32)(unsafe.Add(mBase, uint32(v1245)+56)) = v1241
	*(*int32)(unsafe.Add(mBase, uint32(v1245)+60)) = v1246 + v1243
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1229)))
	v1252 = *(*int64)(unsafe.Add(mBase, uint32(v1251)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1229)+16)) = v1252 - int64(-8192)
	goto L317
L321:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, _consts[240]))
	*(*int32)(unsafe.Add(mBase, uint32(v1266+int32(4)))) = int32(0)
	goto L322
L322:
	;
	v1292 = int32(1)
	goto L112
L323:
	;
	F_recoveryPausesHere(m, int32(0))
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L14
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	F_ProcessStartupProcInterrupts(m)
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L14
	} else {
		goto L327
	}
L326:
	;
	goto L325
L327:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, _consts[391]))
	v216 = v1311
	v231 = v1292
	goto L43
L328:
	;
	goto L42
L329:
	;
	v1319 = F_close(m, v1316)
	mBase = m.M
	goto L331
L330:
	;
	goto L331
L331:
	;
	v1320 = int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[314])) = v1320
	v1325 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[390])) = v1325
	*(*int32)(unsafe.Add(mBase, _consts[387])) = v1325
	v1946 = v1320
	goto L1
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v736
	v1336 = *(*int32)(unsafe.Add(mBase, _consts[397]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1336
	*(*uint32)(unsafe.Add(mBase, uint32(v33)+148)) = uint32(v4)
	v1340 = int64(base.Ui64(v4) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v33)+144)) = uint32(v1340)
	F_errmsg_internal(m, int32(55488), v33+int32(144))
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L14
	} else {
		goto L333
	}
L333:
	;
	F_errfinish(m, int32(511477), int32(3891), int32(411928))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L14
	} else {
		goto L334
	}
L334:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L335:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, _consts[391]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v1357
	F_errmsg_internal(m, int32(496190), v33+int32(16))
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L14
	} else {
		goto L336
	}
L336:
	;
	F_errfinish(m, int32(511477), int32(4033), int32(411928))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L14
	} else {
		goto L337
	}
L337:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L338:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v1394 = base.I32_wrap_i64(v1371)&(v1377-int32(1)) - v40
	goto L28
L339:
	;
	v1437 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v1437))) = int32(167772235)
	v1441 = *(*int32)(unsafe.Add(mBase, _consts[314]))
	v1442 = int32(8192)
	v1444 = int64(*(*uint32)(unsafe.Add(mBase, _consts[389])))
	v1445 = F_pread(m, v1441, l4, v1442, v1444)
	mBase = m.M
	if v1445 != v1442 {
		goto L346
	} else {
		goto L347
	}
L340:
	;
	F___clock_gettime(m, int32(1), v1423)
	mBase = m.M
	v1427 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1423)+8)))
	v1428 = *(*int64)(unsafe.Add(mBase, uint32(v1423)))
	v1432 = v1427 + v1428*int64(1000000000)
	goto L342
L341:
	;
	v1432 = int64(0)
	goto L342
L342:
	;
	m.G0 = v1423 + int32(16)
	goto L339
L343:
	;
	v147 = int32(0)
	goto L25
L344:
	;
	v1939 = *(*int32)(unsafe.Add(mBase, _consts[390]))
	v1946 = v1939
	goto L1
L345:
	;
	v1916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1257)))
	if v1916 != 0 {
		goto L419
	} else {
		goto L420
	}
L346:
	;
	v1449 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	v1451 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v1451))) = int32(0)
	v1457 = int32(1)
	v1458 = base.I64_extend_i32_s(v1445)
	v1462 = m.G0
	v1464 = v1462 - int32(16)
	m.G0 = v1464
	if v1432 != int64(0) {
		goto L350
	} else {
		goto L351
	}
L347:
	;
	goto L348
L348:
	;
	v1711 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v1711))) = int32(0)
	v1717 = int32(1)
	v1718 = int64(8192)
	v1722 = m.G0
	v1724 = v1722 - int32(16)
	m.G0 = v1724
	if v1432 != int64(0) {
		goto L389
	} else {
		goto L390
	}
L349:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, _consts[397]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1593
	v1596 = *(*int64)(unsafe.Add(mBase, _consts[386]))
	v1599 = int64(*(*int32)(unsafe.Add(mBase, _consts[191])))
	v1600 = base.I64_div_u_s(int64(4294967296), v1599)
	v1601 = base.I64_div_u_s(v1596, v1600)
	*(*uint32)(unsafe.Add(mBase, uint32(v33)+132)) = uint32(v1601)
	v1604 = v1596 - v1600*v1601
	*(*uint32)(unsafe.Add(mBase, uint32(v33)+136)) = uint32(v1604)
	v1612 = F_pg_snprintf(m, v33+int32(192), int32(64), int32(531349), v33+int32(128))
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L14
	} else {
		goto L366
	}
L350:
	;
	F___clock_gettime(m, int32(1), v1464)
	mBase = m.M
	v1470 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1464)+8)))
	v1471 = *(*int64)(unsafe.Add(mBase, uint32(v1464)))
	v1475 = v1470 + (v1471*int64(1000000000) - v1432)
	goto L353
L351:
	;
	goto L352
L352:
	;
	v1572 = int32(4530648)
	v1573 = *(*int64)(unsafe.Add(mBase, _consts[403]))
	*(*int64)(unsafe.Add(mBase, _consts[403])) = v1573 + base.I64_extend_i32_u(v1457)
	v1578 = int32(4529688)
	v1579 = *(*int64)(unsafe.Add(mBase, _consts[404]))
	*(*int64)(unsafe.Add(mBase, _consts[404])) = v1579 + v1458
	F_pgstat_count_backend_io_op(m, int32(2), int32(3), int32(6), v1457, v1458)
	mBase = m.M
	v1584 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[154])) = uint8(v1584)
	*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v1584)
	m.G0 = v1464 + int32(16)
	goto L349
L353:
	;
	v1527 = int32(4531608)
	v1528 = *(*int64)(unsafe.Add(mBase, _consts[405]))
	*(*int64)(unsafe.Add(mBase, _consts[405])) = v1528 + v1475
	v1532 = *(*int32)(unsafe.Add(mBase, _consts[200]))
	if base.Ui32(int32(16)) < base.Ui32(v1532) {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	goto L352
L364:
	;
	if int32(1)<<(uint(v1532)%32)&int32(115186) == int32(0) {
		goto L363
	} else {
		goto L365
	}
L365:
	;
	v1550 = int32(4528504)
	v1551 = *(*int64)(unsafe.Add(mBase, _consts[406]))
	*(*int64)(unsafe.Add(mBase, _consts[406])) = v1551 + v1475
	v1555 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[154])) = uint8(v1555)
	*(*uint8)(unsafe.Add(mBase, _consts[202])) = uint8(v1555)
	goto L363
L366:
	;
	if v1445 < int32(0) {
		goto L368
	} else {
		goto L369
	}
L367:
	;
	F_errfinish(m, int32(511477), v1706, int32(482736))
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		goto L14
	} else {
		goto L387
	}
L368:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = v1449
	if v42 != int32(15) {
		v1634 = v42
		goto L371
	} else {
		goto L372
	}
L369:
	;
	goto L370
L370:
	;
	if v42 != int32(15) {
		v1674 = v42
		goto L379
	} else {
		goto L380
	}
L371:
	;
	v1637 = F_errstart(m, v1634, int32(0))
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L14
	} else {
		goto L375
	}
L372:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, _consts[387]))
	if v1622 != int32(2) {
		v1634 = v42
		goto L371
	} else {
		goto L373
	}
L373:
	;
	v1627 = l1 + base.I64_extend_i32_s(l2)
	v1629 = *(*int64)(unsafe.Add(mBase, _consts[407]))
	if v1627 == v1629 {
		v1634 = int32(14)
		goto L371
	} else {
		goto L374
	}
L374:
	;
	*(*int64)(unsafe.Add(mBase, _consts[407])) = v1627
	v1634 = int32(15)
	goto L371
L375:
	;
	if v1637 == int32(0) {
		goto L345
	} else {
		goto L376
	}
L376:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L14
	} else {
		goto L377
	}
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v35
	v1645 = int64(base.Ui64(l1) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v33)+84)) = uint32(v1645)
	v1648 = *(*int32)(unsafe.Add(mBase, _consts[389]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1648
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v33 + int32(192)
	F_errmsg(m, int32(304173), v33+int32(80))
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L14
	} else {
		goto L378
	}
L378:
	;
	v1706 = int32(3445)
	goto L367
L379:
	;
	v1677 = F_errstart(m, v1674, int32(0))
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L14
	} else {
		goto L383
	}
L380:
	;
	v1662 = *(*int32)(unsafe.Add(mBase, _consts[387]))
	if v1662 != int32(2) {
		v1674 = v42
		goto L379
	} else {
		goto L381
	}
L381:
	;
	v1667 = l1 + base.I64_extend_i32_s(l2)
	v1669 = *(*int64)(unsafe.Add(mBase, _consts[407]))
	if v1667 == v1669 {
		v1674 = int32(14)
		goto L379
	} else {
		goto L382
	}
L382:
	;
	*(*int64)(unsafe.Add(mBase, _consts[407])) = v1667
	v1674 = int32(15)
	goto L379
L383:
	;
	if v1677 == int32(0) {
		goto L345
	} else {
		goto L384
	}
L384:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L14
	} else {
		goto L385
	}
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1445
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = int32(8192)
	v1688 = int64(base.Ui64(l1) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v33)+100)) = uint32(v1688)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v35
	v1692 = *(*int32)(unsafe.Add(mBase, _consts[389]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1692
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v33 + int32(192)
	F_errmsg(m, int32(38369), v33+int32(96))
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L14
	} else {
		goto L386
	}
L386:
	;
	v1706 = int32(3452)
	goto L367
L387:
	;
	goto L345
L388:
	;
	v1853 = *(*int32)(unsafe.Add(mBase, _consts[397]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1184)) = v1853
	v1856 = int32(*(*uint8)(unsafe.Add(mBase, _consts[254])))
	if v1856 != int32(1) {
		goto L344
	} else {
		goto L405
	}
L389:
	;
	F___clock_gettime(m, int32(1), v1724)
	mBase = m.M
	v1730 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1724)+8)))
	v1731 = *(*int64)(unsafe.Add(mBase, uint32(v1724)))
	v1735 = v1730 + (v1731*int64(1000000000) - v1432)
	goto L392
L390:
	;
	goto L391
L391:
	;
	v1832 = int32(4530648)
	v1833 = *(*int64)(unsafe.Add(mBase, _consts[403]))
	*(*int64)(unsafe.Add(mBase, _consts[403])) = v1833 + base.I64_extend_i32_u(v1717)
	v1838 = int32(4529688)
	v1839 = *(*int64)(unsafe.Add(mBase, _consts[404]))
	*(*int64)(unsafe.Add(mBase, _consts[404])) = v1839 + v1718
	F_pgstat_count_backend_io_op(m, int32(2), int32(3), int32(6), v1717, v1718)
	mBase = m.M
	v1844 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[154])) = uint8(v1844)
	*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v1844)
	m.G0 = v1724 + int32(16)
	goto L388
L392:
	;
	v1787 = int32(4531608)
	v1788 = *(*int64)(unsafe.Add(mBase, _consts[405]))
	*(*int64)(unsafe.Add(mBase, _consts[405])) = v1788 + v1735
	v1792 = *(*int32)(unsafe.Add(mBase, _consts[200]))
	if base.Ui32(int32(16)) < base.Ui32(v1792) {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	goto L391
L403:
	;
	if int32(1)<<(uint(v1792)%32)&int32(115186) == int32(0) {
		goto L402
	} else {
		goto L404
	}
L404:
	;
	v1810 = int32(4528504)
	v1811 = *(*int64)(unsafe.Add(mBase, _consts[406]))
	*(*int64)(unsafe.Add(mBase, _consts[406])) = v1811 + v1735
	v1815 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[154])) = uint8(v1815)
	*(*uint8)(unsafe.Add(mBase, _consts[202])) = uint8(v1815)
	goto L402
L405:
	;
	v1860 = int64(*(*int32)(unsafe.Add(mBase, _consts[191])))
	v1861 = base.I64_rem_u_s(l1, v1860)
	if v1861 != int64(0) {
		goto L344
	} else {
		goto L406
	}
L406:
	;
	v1864 = F_XLogReaderValidatePageHeader(m, l0, l1, l4)
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L14
	} else {
		goto L407
	}
L407:
	;
	if v1864 != 0 {
		goto L344
	} else {
		goto L408
	}
L408:
	;
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	v1867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1866))))
	if v1867 == int32(0) {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	v1906 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1905))) = uint8(v1906)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1256)) = uint8(v1906)
	goto L345
L410:
	;
	if v42 != int32(15) {
		v1884 = v42
		goto L411
	} else {
		goto L412
	}
L411:
	;
	v1887 = F_errstart(m, v1884, int32(0))
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L14
	} else {
		goto L415
	}
L412:
	;
	v1873 = *(*int32)(unsafe.Add(mBase, _consts[387]))
	if v1873 != int32(2) {
		v1884 = v42
		goto L411
	} else {
		goto L413
	}
L413:
	;
	v1877 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1879 = *(*int64)(unsafe.Add(mBase, _consts[407]))
	if v1877 == v1879 {
		v1884 = int32(14)
		goto L411
	} else {
		goto L414
	}
L414:
	;
	*(*int64)(unsafe.Add(mBase, _consts[407])) = v1877
	v1884 = int32(15)
	goto L411
L415:
	;
	if v1887 == int32(0) {
		goto L409
	} else {
		goto L416
	}
L416:
	;
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+64)) = v1891
	F_errmsg_internal(m, int32(215163), v33-int32(-64))
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L14
	} else {
		goto L417
	}
L417:
	;
	F_errfinish(m, int32(511477), int32(3508), int32(482736))
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L14
	} else {
		goto L418
	}
L418:
	;
	goto L409
L419:
	;
	v1946 = int32(-2)
	goto L1
L420:
	;
	goto L421
L421:
	;
	v1919 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[392])) = uint8(v1919)
	v1922 = *(*int32)(unsafe.Add(mBase, _consts[314]))
	if int32(0) <= v1922 {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	v1925 = F_close(m, v1922)
	mBase = m.M
	goto L424
L423:
	;
	goto L424
L424:
	;
	v1926 = int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[314])) = v1926
	v1931 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[390])) = v1931
	*(*int32)(unsafe.Add(mBase, _consts[387])) = v1931
	v1937 = int32(*(*uint8)(unsafe.Add(mBase, _consts[254])))
	if v1937 != 0 {
		goto L343
	} else {
		goto L425
	}
L425:
	;
	v1946 = v1926
	goto L1
}
func F_XLogPrefetchResetStats(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v2 = *(*int32)(unsafe.Add(mBase, _consts[243]))
	v6 = m.G0
	v7 = int32(16)
	v8 = v6 - v7
	m.G0 = v8
	F___gettimeofday(m, v8)
	mBase = m.M
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v12 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+8)))
	m.G0 = v8 + v7
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = v12 + v11*int64(1000000) - int64(946684800000000)
	v22 = int32(4444892)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[243]))
	v24 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v24
	v27 = *(*int32)(unsafe.Add(mBase, _consts[243]))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+16)) = v24
	v31 = *(*int32)(unsafe.Add(mBase, _consts[243]))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+24)) = v24
	v35 = *(*int32)(unsafe.Add(mBase, _consts[243]))
	*(*int64)(unsafe.Add(mBase, uint32(v35)+32)) = v24
	v39 = *(*int32)(unsafe.Add(mBase, _consts[243]))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+40)) = v24
	v43 = *(*int32)(unsafe.Add(mBase, _consts[243]))
	*(*int64)(unsafe.Add(mBase, uint32(v43)+48)) = v24
	return
}
func F_XLogPrefetcherBeginRead(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(l0)+104)) = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v6 - int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_XLogBeginRead(m, v10, l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		return
	}
}
func F_XLogPrefetcherGetReader(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return v2
}
func F_XLogReadAhead(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v102 int64
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v118 int64
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v144 int64
	_ = v144
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v246 int64
	_ = v246
	var v247 int32
	_ = v247
	var v249 int64
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v267 int64
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v285 int64
	_ = v285
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int64
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v384 int64
	_ = v384
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v404 int64
	_ = v404
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v450 int64
	_ = v450
	var v456 int32
	_ = v456
	var v470 int32
	_ = v470
	var v488 int64
	_ = v488
	var v489 int64
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v564 int64
	_ = v564
	var v566 int64
	_ = v566
	var v568 int64
	_ = v568
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v607 int32
	_ = v607
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v656 int32
	_ = v656
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v708 int32
	_ = v708
	var v735 int32
	_ = v735
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v774 int32
	_ = v774
	var v776 int64
	_ = v776
	var v780 int64
	_ = v780
	var v786 int32
	_ = v786
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v839 int64
	_ = v839
	var v843 int64
	_ = v843
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v883 int64
	_ = v883
	var v886 int64
	_ = v886
	var v892 int32
	_ = v892
	var v895 int64
	_ = v895
	var v899 int64
	_ = v899
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v966 int64
	_ = v966
	var v974 int64
	_ = v974
	var v980 int32
	_ = v980
	var v984 int64
	_ = v984
	var v991 int64
	_ = v991
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1004 int64
	_ = v1004
	var v1009 int64
	_ = v1009
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1020 int64
	_ = v1020
	var v1024 int64
	_ = v1024
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1045 int64
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1057 int64
	_ = v1057
	var v1060 int64
	_ = v1060
	var v1066 int32
	_ = v1066
	var v1067 int64
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1085 int32
	_ = v1085
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1100 int32
	_ = v1100
	var v1107 int32
	_ = v1107
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1219 int32
	_ = v1219
	var v1224 int32
	_ = v1224
	var v1239 int32
	_ = v1239
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1283 int64
	_ = v1283
	var v1286 int64
	_ = v1286
	var v1290 int32
	_ = v1290
	var v1316 int32
	_ = v1316
	var v1344 int32
	_ = v1344
	var v1348 int64
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1390 int64
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1395 int32
	_ = v1395
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1437 int32
	_ = v1437
	var v1442 int32
	_ = v1442
	v2 = l1
	v3 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(16512)
	m.G0 = v28
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1256)))
	if v30 != 0 {
		v1437 = v3
		v1442 = v28
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v1442 + int32(16512)
	return v1437
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v32)
	v34 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v34
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1257)) = uint8(v2)
	v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+1216)) = v39
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v43 = v39 & int64(-8192)
	v44 = int32(8168)
	v45 = base.I32_wrap_i64(v39)
	v47 = v45 & int32(8191)
	if base.Ui32(v44) <= base.Ui32(v47) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v50 = v44
	goto L5
L4:
	;
	v50 = v47
	goto L5
L5:
	;
	v53 = F_ReadPageInternal(m, l0, v43, v50+int32(24))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v53 == int32(-2) {
		v1437 = v3
		v1442 = v28
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v61 = v45
	v63 = v53
	v64 = v47
	v66 = v3
	v82 = v39
	v83 = v43
	goto L12
L9:
	;
	if v1402 == int32(0) {
		goto L309
	} else {
		goto L310
	}
L10:
	;
	v1391 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1366)+1256)) = uint8(v1391)
	*(*int64)(unsafe.Add(mBase, uint32(v1366)+56)) = v1390
	*(*int64)(unsafe.Add(mBase, uint32(v1366)+48)) = v110
	v1395 = v1366
	v1402 = v1373
	v1403 = v1374
	goto L9
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = v489
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470)+17)))
	if v491 != 0 {
		goto L149
	} else {
		goto L150
	}
L12:
	;
	if v63 < int32(0) {
		v1395 = l0
		v1402 = v66
		v1403 = v28
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v424 = int32(8192)
	v425 = v127 + v109
	if base.Ui32(v424) <= base.Ui32(v425) {
		goto L139
	} else {
		goto L140
	}
L14:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+2)))
	if v89&int32(2) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v92 = int32(40)
	goto L17
L16:
	;
	v92 = int32(24)
	goto L17
L17:
	;
	if v64 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v89&int32(1) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	v109 = v92
	v110 = v82 + base.I64_extend_i32_u(v92)
	goto L18
L20:
	;
	goto L21
L21:
	;
	if base.Ui32(v92) <= base.Ui32(v64) {
		v109 = v64
		v110 = v82
		goto L18
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v28)+120)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v28)+116)) = v61
	v102 = int64(base.Ui64(v82) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v28)+112)) = uint32(v102)
	F_report_invalid_record(m, l0, int32(42846), v28+int32(112))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v1395 = l0
	v1402 = v66
	v1403 = v28
	goto L9
L24:
	;
	v123 = base.I32_wrap_i64(v110)
	v125 = v123 & int32(8191)
	v126 = v88 + v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if base.Ui32(v109) <= base.Ui32(int32(8168)) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	if v92 != v109 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v28)+4)) = uint32(v110)
	v118 = int64(base.Ui64(v110) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v28))) = uint32(v118)
	F_report_invalid_record(m, l0, int32(532390), v28)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v1395 = l0
	v1402 = v66
	v1403 = v28
	goto L9
L28:
	;
	v152 = v127 + int32(2037)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v153 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L29:
	;
	v130 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v133 = F_ValidXLogRecordHeader(m, l0, v110, v130, v126, base.B2i32(v41 == int64(0)))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L6
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if base.Ui32(int32(23)) < base.Ui32(v127) {
		goto L28
	} else {
		goto L34
	}
L32:
	;
	if v133 == int32(0) {
		v1395 = l0
		v1402 = v66
		v1403 = v28
		goto L9
	} else {
		goto L33
	}
L33:
	;
	goto L28
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+108)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v28)+104)) = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+100)) = v123
	v144 = int64(base.Ui64(v110) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v28)+96)) = uint32(v144)
	F_report_invalid_record(m, l0, int32(42907), v28+int32(96))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v1395 = l0
	v1402 = v66
	v1403 = v28
	goto L9
L36:
	;
	v199 = int32(8192) - v125
	v200 = base.B2i32(base.Ui32(v127) <= base.Ui32(v199))
	if v200 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L37:
	;
	v193 = int32(0)
	if v2 != 0 {
		v1437 = v193
		v1442 = v28
		goto L1
	} else {
		goto L52
	}
L38:
	;
	v185 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+4)) = uint8(v185)
	v194 = v182
	v197 = base.B2i32(v182 == v185)
	goto L36
L39:
	;
	if base.Ui32(v169-v168) <= base.Ui32(v152) {
		goto L37
	} else {
		goto L51
	}
L40:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if base.Ui32(v152) <= base.Ui32(v174-v171+v172) {
		v182 = v171
		goto L38
	} else {
		goto L49
	}
L41:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v156 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if base.Ui32(v168) < base.Ui32(v169) {
		goto L39
	} else {
		goto L48
	}
L44:
	;
	v160 = v156
	goto L46
L45:
	;
	v157 = int32(65536)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v157
	v160 = v157
	goto L46
L46:
	;
	v161 = F_palloc(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v161
	v166 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v166)
	v171 = v161
	v172 = v161
	v173 = v161
	goto L40
L48:
	;
	v171 = v168
	v172 = v153
	v173 = v169
	goto L40
L49:
	;
	if base.Ui32(v152) < base.Ui32(v173-v172) {
		v182 = v172
		goto L38
	} else {
		goto L50
	}
L50:
	;
	goto L37
L51:
	;
	v182 = v168
	goto L38
L52:
	;
	v194 = v193
	v197 = int32(1)
	goto L36
L53:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v199 != 0 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	goto L55
L55:
	;
	goto L13
L56:
	;
	v210 = int32(40960)
	v211 = int32(-8192)
	v214 = v127&v211 - v211
	if base.Ui32(v214) <= base.Ui32(v210) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v208 = F__emscripten_memcpy_bulkmem(m, v205, v206+v125, v199)
	mBase = m.M
	goto L59
L58:
	;
	goto L59
L59:
	;
	goto L56
L60:
	;
	v217 = v210
	goto L62
L61:
	;
	v217 = v214
	goto L62
L62:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
	v226 = v199
	v231 = base.B2i32(base.Ui32(v109) < base.Ui32(int32(8169)))
	v237 = v220 + v199
	v246 = v83
	goto L64
L63:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1257)) = uint8(v2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v110
	*(*int64)(unsafe.Add(mBase, uint32(l0)+1216)) = v249
	v410 = int32(8168)
	v411 = base.I32_wrap_i64(v249)
	v413 = v411 & int32(8191)
	if base.Ui32(v410) <= base.Ui32(v413) {
		goto L134
	} else {
		goto L135
	}
L64:
	;
	v247 = int32(0)
	v249 = v246 - int64(-8192)
	v251 = F_ReadPageInternal(m, l0, v249, int32(24))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L6
	} else {
		goto L66
	}
L65:
	;
	v368 = int32(-1)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
	v370 = int32(24)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	v375 = m.Env.Pgmem_crc32c(m, v368, v369+v370, v372-v370)
	mBase = m.M
	v377 = m.Env.Pgmem_crc32c(m, v375, v369, int32(20))
	mBase = m.M
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v369)+20))
	if v377^v378 != v368 {
		goto L127
	} else {
		goto L128
	}
L66:
	;
	if v251 == int32(-2) {
		v1437 = v247
		v1442 = v28
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if v251 < int32(0) {
		v1366 = l0
		v1373 = v194
		v1374 = v28
		v1390 = v249
		goto L10
	} else {
		goto L68
	}
L68:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v257)+2)))
	if v258&int32(8) != 0 {
		goto L63
	} else {
		goto L69
	}
L69:
	;
	if v258&int32(1) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v123
	v267 = int64(base.Ui64(v110) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v28)+16)) = uint32(v267)
	F_report_invalid_record(m, l0, int32(534098), v28+int32(16))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L6
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v257)+16))
	if v127 == v226+v274 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v1366 = l0
	v1373 = v194
	v1374 = v28
	v1390 = v249
	goto L10
L74:
	;
	v278 = v274
	goto L76
L75:
	;
	v278 = int32(0)
	goto L76
L76:
	;
	if v278 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+68)) = v123
	v285 = int64(base.Ui64(v110) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v28-int32(-64)))) = uint32(v285)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v274
	*(*int64)(unsafe.Add(mBase, uint32(v28)+56)) = base.I64_extend_i32_u(v127) - base.I64_extend_i32_u(v226)
	F_report_invalid_record(m, l0, int32(535292), v28+int32(48))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L6
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v297 = int32(8192)
	v298 = v127 + int32(24) - v226
	if base.Ui32(v297) <= base.Ui32(v298) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v1366 = l0
	v1373 = v194
	v1374 = v28
	v1390 = v249
	goto L10
L81:
	;
	v301 = v297
	goto L83
L82:
	;
	v301 = v298
	goto L83
L83:
	;
	v302 = F_ReadPageInternal(m, l0, v249, v301)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	if v302 == int32(-2) {
		v1437 = v247
		v1442 = v28
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v302 < int32(0) {
		v1366 = l0
		v1373 = v194
		v1374 = v28
		v1390 = v249
		goto L10
	} else {
		goto L86
	}
L86:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+2)))
	if v310&int32(2) != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v313 = int32(40)
	goto L89
L88:
	;
	v313 = int32(24)
	goto L89
L89:
	;
	if base.Ui32(v302) < base.Ui32(v313) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v315 = F_ReadPageInternal(m, l0, v249, v313)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L6
	} else {
		goto L93
	}
L91:
	;
	v317 = v302
	goto L92
L92:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v257)+16))
	v322 = int32(8192) - v313
	if base.Ui32(v320) < base.Ui32(v322) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v317 = v315
	goto L92
L94:
	;
	v324 = v320
	goto L96
L95:
	;
	v324 = v322
	goto L96
L96:
	;
	v325 = v313 + v324
	if base.Ui32(v317) < base.Ui32(v325) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v327 = F_ReadPageInternal(m, l0, v249, v325)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L6
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	if v324 != 0 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L99
L101:
	;
	if v231&int32(1) == int32(0) {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v329 = F__emscripten_memcpy_bulkmem(m, v237, v318+v313, v324)
	mBase = m.M
	v330 = v329
	goto L104
L103:
	;
	v330 = v237
	goto L104
L104:
	;
	goto L101
L105:
	;
	v335 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
	v339 = F_ValidXLogRecordHeader(m, l0, v110, v335, v336, base.B2i32(v41 == int64(0)))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L6
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v343 = v226 + v324
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1248))
	if base.Ui32(v127) <= base.Ui32(v344) {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	if v339 == int32(0) {
		v1366 = l0
		v1373 = v194
		v1374 = v28
		v1390 = v249
		goto L10
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v365 = v330 + v324
	goto L112
L111:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
	if v343 != 0 {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	if base.Ui32(v343) < base.Ui32(v127) {
		v226 = v343
		v231 = int32(1)
		v237 = v365
		v246 = v249
		goto L64
	} else {
		goto L126
	}
L113:
	;
	if v349 != 0 {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	v350 = F__emscripten_memcpy_bulkmem(m, v28+int32(128), v349, v343)
	mBase = m.M
	goto L116
L115:
	;
	goto L116
L116:
	;
	goto L113
L117:
	;
	F_pfree(m, v349)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L6
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v354 = F_palloc(m, v217)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L6
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1248)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1244)) = v354
	if v343 != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
	v365 = v362 + v343
	goto L112
L123:
	;
	v360 = F__emscripten_memcpy_bulkmem(m, v354, v28+int32(128), v343)
	mBase = m.M
	goto L125
L124:
	;
	goto L125
L125:
	;
	goto L122
L126:
	;
	goto L65
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v123
	v384 = int64(base.Ui64(v110) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v28)+32)) = uint32(v384)
	F_report_invalid_record(m, l0, int32(534361), v28+int32(32))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L6
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v392 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v391)+2)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v110
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v257)+16))
	if v392&int32(2) != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v1366 = l0
	v1373 = v194
	v1374 = v28
	v1390 = v249
	goto L10
L131:
	;
	v404 = int64(40)
	goto L133
L132:
	;
	v404 = int64(24)
	goto L133
L133:
	;
	v470 = v369
	v488 = v249
	v489 = base.I64_extend_i32_u((v394+int32(7))&int32(-8)) + (v404 + v249)
	goto L11
L134:
	;
	v416 = v410
	goto L136
L135:
	;
	v416 = v413
	goto L136
L136:
	;
	v419 = F_ReadPageInternal(m, l0, v249, v416+int32(24))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	if v419 != int32(-2) {
		v61 = v411
		v63 = v419
		v64 = v413
		v66 = v194
		v82 = v249
		v83 = v249
		goto L12
	} else {
		goto L138
	}
L138:
	;
	v1437 = v247
	v1442 = v28
	goto L1
L139:
	;
	v428 = v424
	goto L141
L140:
	;
	v428 = v425
	goto L141
L141:
	;
	v429 = F_ReadPageInternal(m, l0, v83, v428)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L6
	} else {
		goto L142
	}
L142:
	;
	if v429 == int32(-2) {
		v1437 = int32(0)
		v1442 = v28
		goto L1
	} else {
		goto L143
	}
L143:
	;
	if v429 < int32(0) {
		v1395 = l0
		v1402 = v194
		v1403 = v28
		goto L9
	} else {
		goto L144
	}
L144:
	;
	v435 = int32(-1)
	v436 = int32(24)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v441 = m.Env.Pgmem_crc32c(m, v435, v126+v436, v438-v436)
	mBase = m.M
	v443 = m.Env.Pgmem_crc32c(m, v441, v126, int32(20))
	mBase = m.M
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
	if v443^v444 != v435 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v123
	v450 = int64(base.Ui64(v110) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v28)+80)) = uint32(v450)
	F_report_invalid_record(m, l0, int32(534361), v28+int32(80))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L6
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v110
	v470 = v126
	v488 = v83
	v489 = v110 + base.I64_extend_i32_u((v127+int32(7))&int32(-8))
	goto L11
L148:
	;
	v1395 = l0
	v1402 = v194
	v1403 = v28
	goto L9
L149:
	;
	if v197 != 0 {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470)+16)))
	if v492&int32(240) != int32(64) {
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1160))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = (v489 + base.I64_extend_i32_s(v497-int32(1))) & base.I64_extend_i32_s(int32(0)-v497)
	goto L149
L152:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v508 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L153:
	;
	v556 = v194
	goto L154
L154:
	;
	v559 = int32(0)
	v560 = m.G0
	v562 = v560 - int32(176)
	m.G0 = v562
	v564 = *(*int64)(unsafe.Add(mBase, uint32(v470)))
	*(*int64)(unsafe.Add(mBase, uint32(v556)+32)) = v564
	v566 = *(*int64)(unsafe.Add(mBase, uint32(v470)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v556)+48)) = v566
	v568 = *(*int64)(unsafe.Add(mBase, uint32(v470)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v556)+40)) = v568
	*(*int64)(unsafe.Add(mBase, uint32(v556)+16)) = v110
	*(*int64)(unsafe.Add(mBase, uint32(v556)+68)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(v556)+60)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v556)+56)) = uint16(v559)
	*(*int32)(unsafe.Add(mBase, uint32(v556)+8)) = v559
	v580 = v556 + int32(76)
	v581 = int32(24)
	v582 = v470 + v581
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	v585 = v583 - v581
	if v585 == v559 {
		v1219 = v582
		v1224 = v580
		goto L176
	} else {
		goto L177
	}
L155:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v551)+4)) = uint8(v550)
	v556 = v551
	goto L154
L156:
	;
	v546 = F_palloc(m, v152)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L6
	} else {
		goto L172
	}
L157:
	;
	if base.Ui32(v152) < base.Ui32(v524-v523) {
		v550 = int32(0)
		v551 = v523
		goto L155
	} else {
		goto L171
	}
L158:
	;
	v531 = int32(0)
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if base.Ui32(v152) <= base.Ui32(v532-v526+v530) {
		goto L167
	} else {
		goto L168
	}
L159:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v511 != 0 {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	goto L161
L161:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if base.Ui32(v523) < base.Ui32(v524) {
		goto L157
	} else {
		goto L166
	}
L162:
	;
	v515 = v511
	goto L164
L163:
	;
	v512 = int32(65536)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v512
	v515 = v512
	goto L164
L164:
	;
	v516 = F_palloc(m, v515)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L6
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v516
	v521 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v521)
	v526 = v516
	v527 = v516
	v530 = v516
	goto L158
L166:
	;
	v526 = v523
	v527 = v524
	v530 = v508
	goto L158
L167:
	;
	v550 = v531
	v551 = v526
	goto L155
L168:
	;
	goto L169
L169:
	;
	if base.Ui32(v527-v530) <= base.Ui32(v152) {
		goto L156
	} else {
		goto L170
	}
L170:
	;
	v550 = v531
	v551 = v530
	goto L155
L171:
	;
	goto L156
L172:
	;
	v550 = int32(1)
	v551 = v546
	goto L155
L173:
	;
	m.G0 = v562 + int32(176)
	if v1344 != 0 {
		goto L293
	} else {
		goto L294
	}
L174:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	*(*int32)(unsafe.Add(mBase, uint32(v28+int32(128)))) = v1316
	v1344 = int32(0)
	goto L173
L175:
	;
	v1283 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+4)) = uint32(v1283)
	v1286 = int64(base.Ui64(v1283) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v562))) = uint32(v1286)
	F_report_invalid_record(m, l0, int32(533977), v562)
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L6
	} else {
		goto L292
	}
L176:
	;
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v556)+68))
	if v1239 != 0 {
		goto L285
	} else {
		goto L286
	}
L177:
	;
	v592 = int32(-1)
	v594 = v582
	v598 = v585
	v600 = v559
	v607 = v3
	goto L178
L178:
	;
	v615 = v598 - int32(1)
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594))))
	switch v616 - int32(252) {
	case 0:
		goto L183
	case 1:
		goto L184
	case 2:
		goto L185
	case 3:
		goto L186
	default:
		goto L182
	}
L179:
	;
	if v1118 != v1120 {
		goto L175
	} else {
		goto L266
	}
L180:
	;
	goto L179
L181:
	;
	if base.Ui32(v1093) < base.Ui32(v1091) {
		v592 = v1085
		v594 = v1107
		v598 = v1091
		v600 = v1093
		v607 = v1100
		goto L178
	} else {
		goto L265
	}
L182:
	;
	if base.Ui32(v616) <= base.Ui32(int32(32)) {
		goto L193
	} else {
		goto L194
	}
L183:
	;
	if base.Ui32(v598) < base.Ui32(int32(5)) {
		goto L175
	} else {
		goto L190
	}
L184:
	;
	if base.Ui32(v598) < base.Ui32(int32(3)) {
		goto L175
	} else {
		goto L189
	}
L185:
	;
	if base.Ui32(v598) < base.Ui32(int32(5)) {
		goto L175
	} else {
		goto L188
	}
L186:
	;
	if v615 == int32(0) {
		goto L175
	} else {
		goto L187
	}
L187:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v556)+68)) = v621
	v623 = int32(2)
	v1112 = v592
	v1114 = v594 + v623
	v1118 = v598 - v623
	v1120 = v621 + v600
	goto L180
L188:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v594)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v556)+68)) = v630
	v632 = int32(5)
	v1112 = v592
	v1114 = v594 + v632
	v1118 = v598 - v632
	v1120 = v630 + v600
	goto L180
L189:
	;
	v639 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v594)+1)))
	*(*uint16)(unsafe.Add(mBase, uint32(v556)+56)) = uint16(v639)
	v641 = int32(3)
	v1085 = v592
	v1091 = v598 - v641
	v1093 = v600
	v1100 = v607
	v1107 = v594 + v641
	goto L181
L190:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v594)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v556)+60)) = v647
	v649 = int32(5)
	v1085 = v592
	v1091 = v598 - v649
	v1093 = v600
	v1100 = v607
	v1107 = v594 + v649
	goto L181
L191:
	;
	if v616 <= v816 {
		goto L208
	} else {
		goto L209
	}
L192:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v556)+72))
	v816 = v812
	goto L191
L193:
	;
	v656 = v592 + int32(1)
	if v616 <= v656 {
		v816 = v592
		goto L191
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v776 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+168)) = uint32(v776)
	*(*int32)(unsafe.Add(mBase, uint32(v562)+160)) = v616
	v780 = int64(base.Ui64(v776) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+164)) = uint32(v780)
	F_report_invalid_record(m, l0, int32(533157), v562+int32(160))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L6
	} else {
		goto L207
	}
L196:
	;
	v666 = (v592 ^ int32(-1) + v616) & int32(7)
	if v666 != 0 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v668 = int32(0)
	v673 = v656
	goto L200
L198:
	;
	v708 = v656
	goto L199
L199:
	;
	if base.Ui32(v616-v592-int32(2)) <= base.Ui32(int32(6)) {
		goto L192
	} else {
		goto L203
	}
L200:
	;
	v695 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v580+v673*int32(52)))) = uint8(v695)
	v697 = int32(1)
	v698 = v673 + v697
	v700 = v668 + v697
	if v700 != v666 {
		v668 = v700
		v673 = v698
		goto L200
	} else {
		goto L202
	}
L201:
	;
	v708 = v698
	goto L199
L202:
	;
	goto L201
L203:
	;
	v735 = v708
	goto L204
L204:
	;
	v756 = v580 + v735*int32(52)
	v757 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v756))) = uint8(v757)
	*(*uint8)(unsafe.Add(mBase, uint32(v756)+52)) = uint8(v757)
	*(*uint8)(unsafe.Add(mBase, uint32(v756)+104)) = uint8(v757)
	*(*uint8)(unsafe.Add(mBase, uint32(v756)+156)) = uint8(v757)
	*(*uint8)(unsafe.Add(mBase, uint32(v756)+208)) = uint8(v757)
	*(*uint8)(unsafe.Add(mBase, uint32(v756)+260)) = uint8(v757)
	*(*uint8)(unsafe.Add(mBase, uint32(v756)+312)) = uint8(v757)
	*(*uint8)(unsafe.Add(mBase, uint32(v756)+364)) = uint8(v757)
	v774 = v735 + int32(8)
	if v774 != v616 {
		v735 = v774
		goto L204
	} else {
		goto L206
	}
L205:
	;
	goto L192
L206:
	;
	goto L205
L207:
	;
	goto L174
L208:
	;
	v839 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+152)) = uint32(v839)
	*(*int32)(unsafe.Add(mBase, uint32(v562)+144)) = v616
	v843 = int64(base.Ui64(v839) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+148)) = uint32(v843)
	F_report_invalid_record(m, l0, int32(533123), v562+int32(144))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L6
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v556)+72)) = v616
	v853 = v580 + v616*int32(52)
	v854 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v853)+30)) = uint8(v854)
	v856 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v853))) = uint8(v856)
	if v615 == v854 {
		goto L175
	} else {
		goto L212
	}
L211:
	;
	goto L174
L212:
	;
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v853)+28)) = uint8(v860)
	*(*int32)(unsafe.Add(mBase, uint32(v853)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v853)+16)) = v860 & int32(15)
	v869 = int32(1)
	v870 = int32(base.Ui32(v860)>>(uint(int32(5))%32)) & v869
	*(*uint8)(unsafe.Add(mBase, uint32(v853)+43)) = uint8(v870)
	v875 = int32(base.Ui32(v860)>>(uint(int32(4))%32)) & v869
	*(*uint8)(unsafe.Add(mBase, uint32(v853)+29)) = uint8(v875)
	v878 = v598 & int32(-2)
	if v878 == int32(2) {
		goto L175
	} else {
		goto L213
	}
L213:
	;
	v881 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v594)+2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v853)+48)) = uint16(v881)
	if v870 != 0 {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	v906 = int32(4)
	v907 = v598 - v906
	v909 = v594 + v906
	v910 = v600 + v881
	if v875 == int32(0) {
		v1031 = v907
		v1033 = v909
		v1034 = v910
		goto L222
	} else {
		goto L223
	}
L215:
	;
	if v881 != 0 {
		goto L214
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	if v881 == int32(0) {
		goto L214
	} else {
		goto L220
	}
L218:
	;
	v883 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+20)) = uint32(v883)
	v886 = int64(base.Ui64(v883) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+16)) = uint32(v886)
	F_report_invalid_record(m, l0, int32(534981), v562+int32(16))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L6
	} else {
		goto L219
	}
L219:
	;
	goto L174
L220:
	;
	v895 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+136)) = uint32(v895)
	*(*int32)(unsafe.Add(mBase, uint32(v562)+128)) = v881
	v899 = int64(base.Ui64(v895) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+132)) = uint32(v899)
	F_report_invalid_record(m, l0, int32(532724), v562+int32(128))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L6
	} else {
		goto L221
	}
L221:
	;
	goto L174
L222:
	;
	if int32(0) <= base.I32_extend8_s(v860) {
		goto L256
	} else {
		goto L257
	}
L223:
	;
	if base.Ui32(v907) < base.Ui32(int32(2)) {
		goto L175
	} else {
		goto L224
	}
L224:
	;
	v915 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v909))))
	*(*uint16)(unsafe.Add(mBase, uint32(v853)+40)) = uint16(v915)
	if v878 == int32(6) {
		goto L175
	} else {
		goto L225
	}
L225:
	;
	v919 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v594)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v853)+36)) = uint16(v919)
	if v598 == int32(8) {
		goto L175
	} else {
		goto L226
	}
L226:
	;
	v923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v853)+42)) = uint8(v923)
	v925 = int32(1)
	v928 = int32(base.Ui32(v923)>>(uint(v925)%32)) & v925
	*(*uint8)(unsafe.Add(mBase, uint32(v853)+30)) = uint8(v928)
	v930 = int32(9)
	v931 = v598 - v930
	v935 = v923 & int32(28)
	if v935 != 0 {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	if v923&int32(1) != 0 {
		goto L237
	} else {
		goto L238
	}
L228:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v853)+38)) = uint16(v949)
	v951 = v931
	v952 = v594 + v930
	v953 = v949
	goto L227
L229:
	;
	if v923&int32(1) != 0 {
		goto L232
	} else {
		goto L233
	}
L230:
	;
	goto L231
L231:
	;
	v949 = int32(8192) - v915
	goto L228
L232:
	;
	if base.Ui32(v931) < base.Ui32(int32(2)) {
		goto L175
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	v949 = int32(0)
	goto L228
L235:
	;
	v940 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v594)+9)))
	*(*uint16)(unsafe.Add(mBase, uint32(v853)+38)) = uint16(v940)
	v942 = int32(11)
	v951 = v598 - v942
	v952 = v594 + v942
	v953 = v940
	goto L227
L236:
	;
	v1017 = v910 + v915
	if v923&int32(29) != 0 {
		v1031 = v951
		v1033 = v952
		v1034 = v1017
		goto L222
	} else {
		goto L252
	}
L237:
	;
	if v919 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L238:
	;
	goto L239
L239:
	;
	if (v953|v919)&int32(65535) != 0 {
		goto L245
	} else {
		goto L246
	}
L240:
	;
	v966 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+112)) = uint32(v966)
	*(*int32)(unsafe.Add(mBase, uint32(v562)+104)) = v915
	*(*int32)(unsafe.Add(mBase, uint32(v562)+100)) = v953 & int32(65535)
	*(*int32)(unsafe.Add(mBase, uint32(v562)+96)) = v919
	v974 = int64(base.Ui64(v966) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+108)) = uint32(v974)
	F_report_invalid_record(m, l0, int32(532935), v562+int32(96))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L6
	} else {
		goto L244
	}
L241:
	;
	if v953&int32(65535) == int32(0) {
		goto L240
	} else {
		goto L242
	}
L242:
	;
	if v915 != int32(8192) {
		v1016 = int32(0)
		goto L236
	} else {
		goto L243
	}
L243:
	;
	goto L240
L244:
	;
	goto L174
L245:
	;
	v984 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+92)) = uint32(v984)
	*(*int32)(unsafe.Add(mBase, uint32(v562)+84)) = v953 & int32(65535)
	*(*int32)(unsafe.Add(mBase, uint32(v562)+80)) = v919
	v991 = int64(base.Ui64(v984) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+88)) = uint32(v991)
	F_report_invalid_record(m, l0, int32(532870), v562+int32(80))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L6
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	v999 = base.B2i32(v915 == int32(8192))
	if v935 == int32(0) {
		v1016 = v999
		goto L236
	} else {
		goto L249
	}
L248:
	;
	goto L174
L249:
	;
	if v915 != int32(8192) {
		v1016 = v999
		goto L236
	} else {
		goto L250
	}
L250:
	;
	v1004 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+40)) = uint32(v1004)
	*(*int32)(unsafe.Add(mBase, uint32(v562)+32)) = int32(8192)
	v1009 = int64(base.Ui64(v1004) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+36)) = uint32(v1009)
	F_report_invalid_record(m, l0, int32(533018), v562+int32(32))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L6
	} else {
		goto L251
	}
L251:
	;
	goto L174
L252:
	;
	if v1016 != 0 {
		v1031 = v951
		v1033 = v952
		v1034 = v1017
		goto L222
	} else {
		goto L253
	}
L253:
	;
	v1020 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+72)) = uint32(v1020)
	*(*int32)(unsafe.Add(mBase, uint32(v562)+64)) = v881
	v1024 = int64(base.Ui64(v1020) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+68)) = uint32(v1024)
	F_report_invalid_record(m, l0, int32(532631), v562-int32(-64))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L6
	} else {
		goto L254
	}
L254:
	;
	goto L174
L255:
	;
	if base.Ui32(v1071) < base.Ui32(int32(4)) {
		goto L175
	} else {
		goto L264
	}
L256:
	;
	if base.Ui32(v1031) < base.Ui32(int32(12)) {
		goto L175
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	if v607 == int32(0) {
		goto L260
	} else {
		goto L261
	}
L259:
	;
	v1045 = *(*int64)(unsafe.Add(mBase, uint32(v1033)))
	*(*int64)(unsafe.Add(mBase, uint32(v853)+4)) = v1045
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1033)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v853)+12)) = v1047
	v1049 = int32(12)
	v1071 = v1031 - v1049
	v1072 = v1033 + v1049
	v1073 = v853 + int32(4)
	goto L255
L260:
	;
	v1057 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+52)) = uint32(v1057)
	v1060 = int64(base.Ui64(v1057) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v562)+48)) = uint32(v1060)
	F_report_invalid_record(m, l0, int32(533926), v562+int32(48))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L6
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	v1067 = *(*int64)(unsafe.Add(mBase, uint32(v607)))
	*(*int64)(unsafe.Add(mBase, uint32(v853)+4)) = v1067
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v607)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v853)+12)) = v1069
	v1071 = v1031
	v1072 = v1033
	v1073 = v607
	goto L255
L263:
	;
	goto L174
L264:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1072)))
	*(*int32)(unsafe.Add(mBase, uint32(v853)+20)) = v1076
	v1078 = int32(4)
	v1085 = v616
	v1091 = v1071 - v1078
	v1093 = v1034
	v1100 = v1073
	v1107 = v1072 + v1078
	goto L181
L265:
	;
	v1112 = v1085
	v1114 = v1107
	v1118 = v1091
	v1120 = v1093
	goto L180
L266:
	;
	v1136 = v556 + int32(76)
	v1137 = int32(52)
	v1141 = v1136 + v1112*v1137 + v1137
	v1142 = int32(0)
	if v1112 < v1142 {
		v1219 = v1114
		v1224 = v1141
		goto L176
	} else {
		goto L267
	}
L267:
	;
	v1147 = v1142
	v1151 = v1114
	v1152 = int32(0)
	v1156 = v1141
	goto L268
L268:
	;
	v1173 = v1136 + v1147*int32(52)
	v1174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1173))))
	if v1174 != int32(1) {
		v1206 = v1151
		v1207 = v1156
		goto L270
	} else {
		goto L271
	}
L269:
	;
	v1219 = v1206
	v1224 = v1207
	goto L176
L270:
	;
	v1209 = v1152 + int32(1)
	v1211 = v1209 & int32(255)
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v556)+72))
	if v1211 <= v1212 {
		v1147 = v1211
		v1151 = v1206
		v1152 = v1209
		v1156 = v1207
		goto L268
	} else {
		goto L284
	}
L271:
	;
	v1177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1173)+29)))
	if v1177 == int32(1) {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1173)+32)) = v1156
	v1181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1173)+40)))
	if v1181 != 0 {
		goto L276
	} else {
		goto L277
	}
L273:
	;
	v1188 = v1151
	v1189 = v1156
	goto L274
L274:
	;
	v1190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1173)+43)))
	if v1190 != int32(1) {
		v1206 = v1188
		v1207 = v1189
		goto L270
	} else {
		goto L279
	}
L275:
	;
	v1184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1173)+40)))
	v1188 = v1184 + v1151
	v1189 = v1183 + v1184
	goto L274
L276:
	;
	v1182 = F__emscripten_memcpy_bulkmem(m, v1156, v1151, v1181)
	mBase = m.M
	v1183 = v1182
	goto L278
L277:
	;
	v1183 = v1156
	goto L278
L278:
	;
	goto L275
L279:
	;
	v1196 = (v1189 + int32(7)) & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(v1173)+44)) = v1196
	v1198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1173)+48)))
	if v1198 != 0 {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	v1201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1173)+48)))
	v1206 = v1201 + v1188
	v1207 = v1200 + v1201
	goto L270
L281:
	;
	v1199 = F__emscripten_memcpy_bulkmem(m, v1196, v1188, v1198)
	mBase = m.M
	v1200 = v1199
	goto L283
L282:
	;
	v1200 = v1196
	goto L283
L283:
	;
	goto L280
L284:
	;
	goto L269
L285:
	;
	v1243 = (v1224 + int32(7)) & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(v556)+64)) = v1243
	if v1239 != 0 {
		goto L289
	} else {
		goto L290
	}
L286:
	;
	v1250 = v1224
	goto L287
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v556))) = (v1250 - v556 + int32(7)) & int32(-8)
	v1344 = int32(1)
	goto L173
L288:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v556)+68))
	v1250 = v1246 + v1247
	goto L287
L289:
	;
	v1245 = F__emscripten_memcpy_bulkmem(m, v1243, v1219, v1239)
	mBase = m.M
	v1246 = v1245
	goto L291
L290:
	;
	v1246 = v1243
	goto L291
L291:
	;
	goto L288
L292:
	;
	goto L174
L293:
	;
	v1348 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v556)+24)) = v1348
	v1350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+4)))
	if v1350 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L294:
	;
	goto L295
L295:
	;
	if base.Ui32(v127) <= base.Ui32(v199) {
		v1395 = l0
		v1402 = v556
		v1403 = v28
		goto L9
	} else {
		goto L308
	}
L296:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v1353 != v556 {
		goto L299
	} else {
		goto L300
	}
L297:
	;
	goto L298
L298:
	;
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v1361 != 0 {
		goto L302
	} else {
		goto L303
	}
L299:
	;
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v1356 = v1355
	goto L301
L300:
	;
	v1356 = v1353
	goto L301
L301:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v1356 + v1357
	goto L298
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1361)+8)) = v556
	goto L304
L303:
	;
	goto L304
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v556
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1364 != 0 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1437 = v556
	v1442 = v28
	goto L1
L306:
	;
	goto L307
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v556
	v1437 = v556
	v1442 = v28
	goto L1
L308:
	;
	v1366 = l0
	v1373 = v556
	v1374 = v28
	v1390 = v488
	goto L10
L309:
	;
	v1427 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1395)+1192)) = v1427
	*(*int64)(unsafe.Add(mBase, uint32(v1395)+1176)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1395)+132)) = v1427
	v1437 = v1427
	v1442 = v1403
	goto L1
L310:
	;
	v1422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1402)+4)))
	if v1422 != int32(1) {
		goto L309
	} else {
		goto L311
	}
L311:
	;
	F_pfree(m, v1402)
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L6
	} else {
		goto L312
	}
L312:
	;
	goto L309
}
func F_XLogReadBufferForRedo(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = int32(0)
	v6 = F_XLogReadBufferForRedoExtended(m, l0, l1, v4, v4, l2)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_XLogReaderFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1168))
	if v3 != int32(-1) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		m.T0[v6].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
			if v9 == int32(0) {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
				F_pfree(m, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
					if v20 != 0 {
						F_pfree(m, v20)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							F_pfree(m, v23)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								F_pfree(m, l0)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						F_pfree(m, v23)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							F_pfree(m, l0)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
				if v12 != int32(1) {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
					F_pfree(m, v17)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
						if v20 != 0 {
							F_pfree(m, v20)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return
							} else {
								v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								F_pfree(m, v23)
								mBase = m.M
								v25 = m.ExcPending
								if v25 != 0 {
									return
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v27 = m.ExcPending
									if v27 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							F_pfree(m, v23)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								F_pfree(m, l0)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					F_pfree(m, v9)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
						F_pfree(m, v17)
						mBase = m.M
						v19 = m.ExcPending
						if v19 != 0 {
							return
						} else {
							v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
							if v20 != 0 {
								F_pfree(m, v20)
								mBase = m.M
								v22 = m.ExcPending
								if v22 != 0 {
									return
								} else {
									v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									F_pfree(m, v23)
									mBase = m.M
									v25 = m.ExcPending
									if v25 != 0 {
										return
									} else {
										F_pfree(m, l0)
										mBase = m.M
										v27 = m.ExcPending
										if v27 != 0 {
											return
										} else {
											return
										}
									}
								}
							} else {
								v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								F_pfree(m, v23)
								mBase = m.M
								v25 = m.ExcPending
								if v25 != 0 {
									return
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v27 = m.ExcPending
									if v27 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		if v9 == int32(0) {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
			F_pfree(m, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
				if v20 != 0 {
					F_pfree(m, v20)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						F_pfree(m, v23)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							F_pfree(m, l0)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					F_pfree(m, v23)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						F_pfree(m, l0)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
			if v12 != int32(1) {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
				F_pfree(m, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
					if v20 != 0 {
						F_pfree(m, v20)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							F_pfree(m, v23)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								F_pfree(m, l0)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						F_pfree(m, v23)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							F_pfree(m, l0)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				F_pfree(m, v9)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
					F_pfree(m, v17)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
						if v20 != 0 {
							F_pfree(m, v20)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return
							} else {
								v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								F_pfree(m, v23)
								mBase = m.M
								v25 = m.ExcPending
								if v25 != 0 {
									return
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v27 = m.ExcPending
									if v27 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							F_pfree(m, v23)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								F_pfree(m, l0)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_XLogReaderValidatePageHeader(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v34 int64
	_ = v34
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int64
	_ = v50
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v72 int64
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int64
	_ = v86
	var v96 int32
	_ = v96
	var v100 int64
	_ = v100
	var v103 int64
	_ = v103
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v136 int64
	_ = v136
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int64
	_ = v152
	var v160 int32
	_ = v160
	var v163 int64
	_ = v163
	var v165 int32
	_ = v165
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v172 int64
	_ = v172
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v185 int64
	_ = v185
	var v186 int64
	_ = v186
	var v190 int64
	_ = v190
	var v199 int32
	_ = v199
	var v201 int64
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int64
	_ = v209
	var v210 int64
	_ = v210
	var v213 int64
	_ = v213
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v230 int64
	_ = v230
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	v13 = m.G0
	v15 = v13 - int32(320)
	m.G0 = v15
	v17 = base.I32_wrap_i64(l1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1160))
	v21 = v17 & (v18 - int32(1))
	v22 = base.I64_extend_i32_s(v18)
	v23 = base.I64_div_u_s(l1, v22)
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	if v24 != int32(53528) {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
		*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v27
		v30 = base.I64_div_u_s(int64(4294967296), v22)
		v31 = base.I64_div_u_s(v23, v30)
		*(*uint32)(unsafe.Add(mBase, uint32(v15)+244)) = uint32(v31)
		v34 = v23 - v30*v31
		*(*uint32)(unsafe.Add(mBase, uint32(v15)+248)) = uint32(v34)
		v42 = F_pg_snprintf(m, v15+int32(256), int32(64), int32(531349), v15+int32(240))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+224)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v15)+220)) = v17
			v50 = int64(base.Ui64(l1) >> (uint(int64(32)) % 64))
			*(*uint32)(unsafe.Add(mBase, uint32(v15)+216)) = uint32(v50)
			*(*int32)(unsafe.Add(mBase, uint32(v15)+208)) = v46
			*(*int32)(unsafe.Add(mBase, uint32(v15)+212)) = v15 + int32(256)
			F_report_invalid_record(m, l0, int32(44746), v15+int32(208))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				v247 = int32(0)
				m.G0 = v15 + int32(320)
				return v247
			}
		}
	} else {
		v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
		if base.Ui32(int32(16)) <= base.Ui32(v62) {
			v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+192)) = v65
			v68 = base.I64_div_u_s(int64(4294967296), v22)
			v69 = base.I64_div_u_s(v23, v68)
			*(*uint32)(unsafe.Add(mBase, uint32(v15)+196)) = uint32(v69)
			v72 = v23 - v68*v69
			*(*uint32)(unsafe.Add(mBase, uint32(v15)+200)) = uint32(v72)
			v80 = F_pg_snprintf(m, v15+int32(256), int32(64), int32(531349), v15+int32(192))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v21
				*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v17
				v86 = int64(base.Ui64(l1) >> (uint(int64(32)) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v15)+168)) = uint32(v86)
				*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v82
				*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v15 + int32(256)
				F_report_invalid_record(m, l0, int32(44683), v15+int32(160))
				mBase = m.M
				v96 = m.ExcPending
				if v96 != 0 {
					return int32(0)
				} else {
					v247 = int32(0)
					m.G0 = v15 + int32(320)
					return v247
				}
			}
		} else {
			if v62&int32(2) != 0 {
				v100 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
				if v100 == int64(0) {
					v114 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
					if v18 != v114 {
						v116 = int32(0)
						F_report_invalid_record(m, l0, int32(237498), v116)
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							v247 = v116
							m.G0 = v15 + int32(320)
							return v247
						}
					} else {
						v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
						if v121 == int32(8192) {
							v163 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
							if l1 != v163 {
								v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
								*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v165
								v168 = base.I64_div_u_s(int64(4294967296), v22)
								v169 = base.I64_div_u_s(v23, v168)
								*(*uint32)(unsafe.Add(mBase, uint32(v15)+132)) = uint32(v169)
								v172 = v23 - v168*v169
								*(*uint32)(unsafe.Add(mBase, uint32(v15)+136)) = uint32(v172)
								v180 = F_pg_snprintf(m, v15+int32(256), int32(64), int32(531349), v15+int32(128))
								mBase = m.M
								v181 = m.ExcPending
								if v181 != 0 {
									return int32(0)
								} else {
									v182 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v15)+116)) = v21
									*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v17
									v185 = int64(32)
									v186 = int64(base.Ui64(l1) >> (uint(v185) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v15)+108)) = uint32(v186)
									*(*uint32)(unsafe.Add(mBase, uint32(v15)+100)) = uint32(v182)
									v190 = int64(base.Ui64(v182) >> (uint(v185) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v15)+96)) = uint32(v190)
									*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = v15 + int32(256)
									F_report_invalid_record(m, l0, int32(44812), v15+int32(96))
									mBase = m.M
									v199 = m.ExcPending
									if v199 != 0 {
										return int32(0)
									} else {
										v247 = int32(0)
										m.G0 = v15 + int32(320)
										return v247
									}
								}
							} else {
								v201 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1200))
								if base.Ui64(l1) <= base.Ui64(v201) {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+1200)) = l1
									v244 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+1208)) = v244
									v247 = int32(1)
									m.G0 = v15 + int32(320)
									return v247
								} else {
									v203 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
									v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1208))
									if base.Ui32(v204) <= base.Ui32(v203) {
										*(*int64)(unsafe.Add(mBase, uint32(l0)+1200)) = l1
										v244 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+1208)) = v244
										v247 = int32(1)
										m.G0 = v15 + int32(320)
										return v247
									} else {
										v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
										*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v206
										v209 = base.I64_div_u_s(int64(4294967296), v22)
										v210 = base.I64_div_u_s(v23, v209)
										*(*uint32)(unsafe.Add(mBase, uint32(v15)+84)) = uint32(v210)
										v213 = v23 - v209*v210
										*(*uint32)(unsafe.Add(mBase, uint32(v15)+88)) = uint32(v213)
										v221 = F_pg_snprintf(m, v15+int32(256), int32(64), int32(531349), v15+int32(80))
										mBase = m.M
										v222 = m.ExcPending
										if v222 != 0 {
											return int32(0)
										} else {
											v223 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
											v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1208))
											*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v21
											*(*int32)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = v17
											v230 = int64(base.Ui64(l1) >> (uint(int64(32)) % 64))
											*(*uint32)(unsafe.Add(mBase, uint32(v15)+60)) = uint32(v230)
											*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v224
											*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v223
											*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v15 + int32(256)
											F_report_invalid_record(m, l0, int32(44878), v15+int32(48))
											mBase = m.M
											v241 = m.ExcPending
											if v241 != 0 {
												return int32(0)
											} else {
												v247 = int32(0)
												m.G0 = v15 + int32(320)
												return v247
											}
										}
									}
								}
							}
						} else {
							v124 = int32(0)
							F_report_invalid_record(m, l0, int32(237580), v124)
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return int32(0)
							} else {
								v247 = v124
								m.G0 = v15 + int32(320)
								return v247
							}
						}
					}
				} else {
					v103 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
					if v103 == v100 {
						v114 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
						if v18 != v114 {
							v116 = int32(0)
							F_report_invalid_record(m, l0, int32(237498), v116)
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return int32(0)
							} else {
								v247 = v116
								m.G0 = v15 + int32(320)
								return v247
							}
						} else {
							v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
							if v121 == int32(8192) {
								v163 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
								if l1 != v163 {
									v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
									*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v165
									v168 = base.I64_div_u_s(int64(4294967296), v22)
									v169 = base.I64_div_u_s(v23, v168)
									*(*uint32)(unsafe.Add(mBase, uint32(v15)+132)) = uint32(v169)
									v172 = v23 - v168*v169
									*(*uint32)(unsafe.Add(mBase, uint32(v15)+136)) = uint32(v172)
									v180 = F_pg_snprintf(m, v15+int32(256), int32(64), int32(531349), v15+int32(128))
									mBase = m.M
									v181 = m.ExcPending
									if v181 != 0 {
										return int32(0)
									} else {
										v182 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v15)+116)) = v21
										*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v17
										v185 = int64(32)
										v186 = int64(base.Ui64(l1) >> (uint(v185) % 64))
										*(*uint32)(unsafe.Add(mBase, uint32(v15)+108)) = uint32(v186)
										*(*uint32)(unsafe.Add(mBase, uint32(v15)+100)) = uint32(v182)
										v190 = int64(base.Ui64(v182) >> (uint(v185) % 64))
										*(*uint32)(unsafe.Add(mBase, uint32(v15)+96)) = uint32(v190)
										*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = v15 + int32(256)
										F_report_invalid_record(m, l0, int32(44812), v15+int32(96))
										mBase = m.M
										v199 = m.ExcPending
										if v199 != 0 {
											return int32(0)
										} else {
											v247 = int32(0)
											m.G0 = v15 + int32(320)
											return v247
										}
									}
								} else {
									v201 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1200))
									if base.Ui64(l1) <= base.Ui64(v201) {
										*(*int64)(unsafe.Add(mBase, uint32(l0)+1200)) = l1
										v244 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+1208)) = v244
										v247 = int32(1)
										m.G0 = v15 + int32(320)
										return v247
									} else {
										v203 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
										v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1208))
										if base.Ui32(v204) <= base.Ui32(v203) {
											*(*int64)(unsafe.Add(mBase, uint32(l0)+1200)) = l1
											v244 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+1208)) = v244
											v247 = int32(1)
											m.G0 = v15 + int32(320)
											return v247
										} else {
											v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
											*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v206
											v209 = base.I64_div_u_s(int64(4294967296), v22)
											v210 = base.I64_div_u_s(v23, v209)
											*(*uint32)(unsafe.Add(mBase, uint32(v15)+84)) = uint32(v210)
											v213 = v23 - v209*v210
											*(*uint32)(unsafe.Add(mBase, uint32(v15)+88)) = uint32(v213)
											v221 = F_pg_snprintf(m, v15+int32(256), int32(64), int32(531349), v15+int32(80))
											mBase = m.M
											v222 = m.ExcPending
											if v222 != 0 {
												return int32(0)
											} else {
												v223 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
												v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1208))
												*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v21
												*(*int32)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = v17
												v230 = int64(base.Ui64(l1) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v15)+60)) = uint32(v230)
												*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v224
												*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v223
												*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v15 + int32(256)
												F_report_invalid_record(m, l0, int32(44878), v15+int32(48))
												mBase = m.M
												v241 = m.ExcPending
												if v241 != 0 {
													return int32(0)
												} else {
													v247 = int32(0)
													m.G0 = v15 + int32(320)
													return v247
												}
											}
										}
									}
								}
							} else {
								v124 = int32(0)
								F_report_invalid_record(m, l0, int32(237580), v124)
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return int32(0)
								} else {
									v247 = v124
									m.G0 = v15 + int32(320)
									return v247
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v15)+152)) = v100
						*(*int64)(unsafe.Add(mBase, uint32(v15)+144)) = v103
						F_report_invalid_record(m, l0, int32(39282), v15+int32(144))
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return int32(0)
						} else {
							v247 = int32(0)
							m.G0 = v15 + int32(320)
							return v247
						}
					}
				}
			} else {
				if v21 != 0 {
					v163 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
					if l1 != v163 {
						v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v165
						v168 = base.I64_div_u_s(int64(4294967296), v22)
						v169 = base.I64_div_u_s(v23, v168)
						*(*uint32)(unsafe.Add(mBase, uint32(v15)+132)) = uint32(v169)
						v172 = v23 - v168*v169
						*(*uint32)(unsafe.Add(mBase, uint32(v15)+136)) = uint32(v172)
						v180 = F_pg_snprintf(m, v15+int32(256), int32(64), int32(531349), v15+int32(128))
						mBase = m.M
						v181 = m.ExcPending
						if v181 != 0 {
							return int32(0)
						} else {
							v182 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v15)+116)) = v21
							*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v17
							v185 = int64(32)
							v186 = int64(base.Ui64(l1) >> (uint(v185) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v15)+108)) = uint32(v186)
							*(*uint32)(unsafe.Add(mBase, uint32(v15)+100)) = uint32(v182)
							v190 = int64(base.Ui64(v182) >> (uint(v185) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v15)+96)) = uint32(v190)
							*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = v15 + int32(256)
							F_report_invalid_record(m, l0, int32(44812), v15+int32(96))
							mBase = m.M
							v199 = m.ExcPending
							if v199 != 0 {
								return int32(0)
							} else {
								v247 = int32(0)
								m.G0 = v15 + int32(320)
								return v247
							}
						}
					} else {
						v201 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1200))
						if base.Ui64(l1) <= base.Ui64(v201) {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+1200)) = l1
							v244 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+1208)) = v244
							v247 = int32(1)
							m.G0 = v15 + int32(320)
							return v247
						} else {
							v203 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
							v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1208))
							if base.Ui32(v204) <= base.Ui32(v203) {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+1200)) = l1
								v244 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+1208)) = v244
								v247 = int32(1)
								m.G0 = v15 + int32(320)
								return v247
							} else {
								v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
								*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v206
								v209 = base.I64_div_u_s(int64(4294967296), v22)
								v210 = base.I64_div_u_s(v23, v209)
								*(*uint32)(unsafe.Add(mBase, uint32(v15)+84)) = uint32(v210)
								v213 = v23 - v209*v210
								*(*uint32)(unsafe.Add(mBase, uint32(v15)+88)) = uint32(v213)
								v221 = F_pg_snprintf(m, v15+int32(256), int32(64), int32(531349), v15+int32(80))
								mBase = m.M
								v222 = m.ExcPending
								if v222 != 0 {
									return int32(0)
								} else {
									v223 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
									v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1208))
									*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v21
									*(*int32)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = v17
									v230 = int64(base.Ui64(l1) >> (uint(int64(32)) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v15)+60)) = uint32(v230)
									*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v224
									*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v223
									*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v15 + int32(256)
									F_report_invalid_record(m, l0, int32(44878), v15+int32(48))
									mBase = m.M
									v241 = m.ExcPending
									if v241 != 0 {
										return int32(0)
									} else {
										v247 = int32(0)
										m.G0 = v15 + int32(320)
										return v247
									}
								}
							}
						}
					}
				} else {
					v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v129
					v132 = base.I64_div_u_s(int64(4294967296), v22)
					v133 = base.I64_div_u_s(v23, v132)
					*(*uint32)(unsafe.Add(mBase, uint32(v15)+36)) = uint32(v133)
					v136 = v23 - v132*v133
					*(*uint32)(unsafe.Add(mBase, uint32(v15)+40)) = uint32(v136)
					v144 = F_pg_snprintf(m, v15+int32(256), int32(64), int32(531349), v15+int32(32))
					mBase = m.M
					v145 = m.ExcPending
					if v145 != 0 {
						return int32(0)
					} else {
						v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
						v147 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v147
						*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v17
						v152 = int64(base.Ui64(l1) >> (uint(int64(32)) % 64))
						*(*uint32)(unsafe.Add(mBase, uint32(v15)+8)) = uint32(v152)
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v146
						*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v15 + int32(256)
						F_report_invalid_record(m, l0, int32(44683), v15)
						mBase = m.M
						v160 = m.ExcPending
						if v160 != 0 {
							return int32(0)
						} else {
							v247 = v147
							m.G0 = v15 + int32(320)
							return v247
						}
					}
				}
			}
		}
	}
}
func F_XLogSendLogical(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v35 int64
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v107 int64
	_ = v107
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*uint8)(unsafe.Add(mBase, _consts[744])) = uint8(v1)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[745]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v17 = F_XLogReadRecord(m, v14, v7+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		if v19 == int32(0) {
			if v17 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, _consts[745]))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
				F_LogicalDecodingProcessRecord(m, v23, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, _consts[745]))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
					v31 = *(*int64)(unsafe.Add(mBase, uint32(v30)+40))
					*(*int64)(unsafe.Add(mBase, _consts[746])) = v31
					v35 = *(*int64)(unsafe.Add(mBase, _consts[747]))
					if v35 != int64(0) {
						v39 = *(*int32)(unsafe.Add(mBase, _consts[745]))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
						v41 = *(*int64)(unsafe.Add(mBase, uint32(v40)+40))
						if base.Ui64(v41) < base.Ui64(v35) {
							v74 = v35
							v76 = *(*int32)(unsafe.Add(mBase, _consts[745]))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
							v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+40))
							if base.Ui64(v74) <= base.Ui64(v78) {
								v81 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[744])) = uint8(v81)
								v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
								if v88 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
								}
							} else {
								v84 = int32(*(*uint8)(unsafe.Add(mBase, _consts[744])))
								if v84 == int32(0) {
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
									if v88 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
									}
								}
							}
							v95 = *(*int32)(unsafe.Add(mBase, _consts[716]))
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+76))
							*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(1)
							if v96 != 0 {
								F_s_lock(m, v95+int32(76), int32(515498), int32(3496), int32(326969))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return
								} else {
									v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
									*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
									m.G0 = v7 + int32(16)
									return
								}
							} else {
								v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
								*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
								m.G0 = v7 + int32(16)
								return
							}
						} else {
							v45 = int32(*(*uint8)(unsafe.Add(mBase, _consts[750])))
							if v45 == int32(1) {
								v49 = F_GetXLogReplayRecPtr(m, int32(0))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									v72 = v49
									*(*int64)(unsafe.Add(mBase, _consts[747])) = v72
									v74 = v72
									v76 = *(*int32)(unsafe.Add(mBase, _consts[745]))
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
									v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+40))
									if base.Ui64(v74) <= base.Ui64(v78) {
										v81 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _consts[744])) = uint8(v81)
										v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
										if v88 == int32(0) {
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
										}
									} else {
										v84 = int32(*(*uint8)(unsafe.Add(mBase, _consts[744])))
										if v84 == int32(0) {
										} else {
											v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
											if v88 == int32(0) {
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
											}
										}
									}
									v95 = *(*int32)(unsafe.Add(mBase, _consts[716]))
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+76))
									*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(1)
									if v96 != 0 {
										F_s_lock(m, v95+int32(76), int32(515498), int32(3496), int32(326969))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return
										} else {
											v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
											*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
											*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
											m.G0 = v7 + int32(16)
											return
										}
									} else {
										v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
										*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
										*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
										m.G0 = v7 + int32(16)
										return
									}
								}
							} else {
								v54 = int32(4444592)
								v55 = *(*int32)(unsafe.Add(mBase, _consts[2]))
								v56 = *(*int64)(unsafe.Add(mBase, uint32(v55)+280))
								*(*int64)(unsafe.Add(mBase, uint32(v55)+280)) = v56
								*(*int64)(unsafe.Add(mBase, _consts[189])) = v56
								v61 = *(*int32)(unsafe.Add(mBase, _consts[2]))
								v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)+272))
								*(*int64)(unsafe.Add(mBase, uint32(v61)+272)) = v62
								*(*int64)(unsafe.Add(mBase, _consts[190])) = v62
								v71 = *(*int64)(unsafe.Add(mBase, _consts[189]))
								v72 = v71
								*(*int64)(unsafe.Add(mBase, _consts[747])) = v72
								v74 = v72
								v76 = *(*int32)(unsafe.Add(mBase, _consts[745]))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
								v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+40))
								if base.Ui64(v74) <= base.Ui64(v78) {
									v81 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[744])) = uint8(v81)
									v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
									if v88 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
									}
								} else {
									v84 = int32(*(*uint8)(unsafe.Add(mBase, _consts[744])))
									if v84 == int32(0) {
									} else {
										v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
										if v88 == int32(0) {
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
										}
									}
								}
								v95 = *(*int32)(unsafe.Add(mBase, _consts[716]))
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+76))
								*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(1)
								if v96 != 0 {
									F_s_lock(m, v95+int32(76), int32(515498), int32(3496), int32(326969))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return
									} else {
										v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
										*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
										*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
										m.G0 = v7 + int32(16)
										return
									}
								} else {
									v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
									*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
									m.G0 = v7 + int32(16)
									return
								}
							}
						}
					} else {
						v45 = int32(*(*uint8)(unsafe.Add(mBase, _consts[750])))
						if v45 == int32(1) {
							v49 = F_GetXLogReplayRecPtr(m, int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								v72 = v49
								*(*int64)(unsafe.Add(mBase, _consts[747])) = v72
								v74 = v72
								v76 = *(*int32)(unsafe.Add(mBase, _consts[745]))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
								v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+40))
								if base.Ui64(v74) <= base.Ui64(v78) {
									v81 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[744])) = uint8(v81)
									v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
									if v88 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
									}
								} else {
									v84 = int32(*(*uint8)(unsafe.Add(mBase, _consts[744])))
									if v84 == int32(0) {
									} else {
										v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
										if v88 == int32(0) {
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
										}
									}
								}
								v95 = *(*int32)(unsafe.Add(mBase, _consts[716]))
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+76))
								*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(1)
								if v96 != 0 {
									F_s_lock(m, v95+int32(76), int32(515498), int32(3496), int32(326969))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return
									} else {
										v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
										*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
										*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
										m.G0 = v7 + int32(16)
										return
									}
								} else {
									v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
									*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
									m.G0 = v7 + int32(16)
									return
								}
							}
						} else {
							v54 = int32(4444592)
							v55 = *(*int32)(unsafe.Add(mBase, _consts[2]))
							v56 = *(*int64)(unsafe.Add(mBase, uint32(v55)+280))
							*(*int64)(unsafe.Add(mBase, uint32(v55)+280)) = v56
							*(*int64)(unsafe.Add(mBase, _consts[189])) = v56
							v61 = *(*int32)(unsafe.Add(mBase, _consts[2]))
							v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)+272))
							*(*int64)(unsafe.Add(mBase, uint32(v61)+272)) = v62
							*(*int64)(unsafe.Add(mBase, _consts[190])) = v62
							v71 = *(*int64)(unsafe.Add(mBase, _consts[189]))
							v72 = v71
							*(*int64)(unsafe.Add(mBase, _consts[747])) = v72
							v74 = v72
							v76 = *(*int32)(unsafe.Add(mBase, _consts[745]))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
							v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+40))
							if base.Ui64(v74) <= base.Ui64(v78) {
								v81 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[744])) = uint8(v81)
								v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
								if v88 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
								}
							} else {
								v84 = int32(*(*uint8)(unsafe.Add(mBase, _consts[744])))
								if v84 == int32(0) {
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
									if v88 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
									}
								}
							}
							v95 = *(*int32)(unsafe.Add(mBase, _consts[716]))
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+76))
							*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(1)
							if v96 != 0 {
								F_s_lock(m, v95+int32(76), int32(515498), int32(3496), int32(326969))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return
								} else {
									v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
									*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
									m.G0 = v7 + int32(16)
									return
								}
							} else {
								v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
								*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
								m.G0 = v7 + int32(16)
								return
							}
						}
					}
				}
			} else {
				v35 = *(*int64)(unsafe.Add(mBase, _consts[747]))
				if v35 != int64(0) {
					v39 = *(*int32)(unsafe.Add(mBase, _consts[745]))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
					v41 = *(*int64)(unsafe.Add(mBase, uint32(v40)+40))
					if base.Ui64(v41) < base.Ui64(v35) {
						v74 = v35
						v76 = *(*int32)(unsafe.Add(mBase, _consts[745]))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
						v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+40))
						if base.Ui64(v74) <= base.Ui64(v78) {
							v81 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[744])) = uint8(v81)
							v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
							if v88 == int32(0) {
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
							}
						} else {
							v84 = int32(*(*uint8)(unsafe.Add(mBase, _consts[744])))
							if v84 == int32(0) {
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
								if v88 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
								}
							}
						}
						v95 = *(*int32)(unsafe.Add(mBase, _consts[716]))
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+76))
						*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(1)
						if v96 != 0 {
							F_s_lock(m, v95+int32(76), int32(515498), int32(3496), int32(326969))
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return
							} else {
								v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
								*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
								m.G0 = v7 + int32(16)
								return
							}
						} else {
							v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
							*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
							m.G0 = v7 + int32(16)
							return
						}
					} else {
						v45 = int32(*(*uint8)(unsafe.Add(mBase, _consts[750])))
						if v45 == int32(1) {
							v49 = F_GetXLogReplayRecPtr(m, int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								v72 = v49
								*(*int64)(unsafe.Add(mBase, _consts[747])) = v72
								v74 = v72
								v76 = *(*int32)(unsafe.Add(mBase, _consts[745]))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
								v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+40))
								if base.Ui64(v74) <= base.Ui64(v78) {
									v81 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[744])) = uint8(v81)
									v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
									if v88 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
									}
								} else {
									v84 = int32(*(*uint8)(unsafe.Add(mBase, _consts[744])))
									if v84 == int32(0) {
									} else {
										v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
										if v88 == int32(0) {
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
										}
									}
								}
								v95 = *(*int32)(unsafe.Add(mBase, _consts[716]))
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+76))
								*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(1)
								if v96 != 0 {
									F_s_lock(m, v95+int32(76), int32(515498), int32(3496), int32(326969))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return
									} else {
										v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
										*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
										*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
										m.G0 = v7 + int32(16)
										return
									}
								} else {
									v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
									*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
									m.G0 = v7 + int32(16)
									return
								}
							}
						} else {
							v54 = int32(4444592)
							v55 = *(*int32)(unsafe.Add(mBase, _consts[2]))
							v56 = *(*int64)(unsafe.Add(mBase, uint32(v55)+280))
							*(*int64)(unsafe.Add(mBase, uint32(v55)+280)) = v56
							*(*int64)(unsafe.Add(mBase, _consts[189])) = v56
							v61 = *(*int32)(unsafe.Add(mBase, _consts[2]))
							v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)+272))
							*(*int64)(unsafe.Add(mBase, uint32(v61)+272)) = v62
							*(*int64)(unsafe.Add(mBase, _consts[190])) = v62
							v71 = *(*int64)(unsafe.Add(mBase, _consts[189]))
							v72 = v71
							*(*int64)(unsafe.Add(mBase, _consts[747])) = v72
							v74 = v72
							v76 = *(*int32)(unsafe.Add(mBase, _consts[745]))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
							v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+40))
							if base.Ui64(v74) <= base.Ui64(v78) {
								v81 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[744])) = uint8(v81)
								v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
								if v88 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
								}
							} else {
								v84 = int32(*(*uint8)(unsafe.Add(mBase, _consts[744])))
								if v84 == int32(0) {
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
									if v88 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
									}
								}
							}
							v95 = *(*int32)(unsafe.Add(mBase, _consts[716]))
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+76))
							*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(1)
							if v96 != 0 {
								F_s_lock(m, v95+int32(76), int32(515498), int32(3496), int32(326969))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return
								} else {
									v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
									*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
									m.G0 = v7 + int32(16)
									return
								}
							} else {
								v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
								*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
								m.G0 = v7 + int32(16)
								return
							}
						}
					}
				} else {
					v45 = int32(*(*uint8)(unsafe.Add(mBase, _consts[750])))
					if v45 == int32(1) {
						v49 = F_GetXLogReplayRecPtr(m, int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							v72 = v49
							*(*int64)(unsafe.Add(mBase, _consts[747])) = v72
							v74 = v72
							v76 = *(*int32)(unsafe.Add(mBase, _consts[745]))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
							v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+40))
							if base.Ui64(v74) <= base.Ui64(v78) {
								v81 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[744])) = uint8(v81)
								v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
								if v88 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
								}
							} else {
								v84 = int32(*(*uint8)(unsafe.Add(mBase, _consts[744])))
								if v84 == int32(0) {
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
									if v88 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
									}
								}
							}
							v95 = *(*int32)(unsafe.Add(mBase, _consts[716]))
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+76))
							*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(1)
							if v96 != 0 {
								F_s_lock(m, v95+int32(76), int32(515498), int32(3496), int32(326969))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return
								} else {
									v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
									*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
									m.G0 = v7 + int32(16)
									return
								}
							} else {
								v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
								*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
								m.G0 = v7 + int32(16)
								return
							}
						}
					} else {
						v54 = int32(4444592)
						v55 = *(*int32)(unsafe.Add(mBase, _consts[2]))
						v56 = *(*int64)(unsafe.Add(mBase, uint32(v55)+280))
						*(*int64)(unsafe.Add(mBase, uint32(v55)+280)) = v56
						*(*int64)(unsafe.Add(mBase, _consts[189])) = v56
						v61 = *(*int32)(unsafe.Add(mBase, _consts[2]))
						v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)+272))
						*(*int64)(unsafe.Add(mBase, uint32(v61)+272)) = v62
						*(*int64)(unsafe.Add(mBase, _consts[190])) = v62
						v71 = *(*int64)(unsafe.Add(mBase, _consts[189]))
						v72 = v71
						*(*int64)(unsafe.Add(mBase, _consts[747])) = v72
						v74 = v72
						v76 = *(*int32)(unsafe.Add(mBase, _consts[745]))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
						v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+40))
						if base.Ui64(v74) <= base.Ui64(v78) {
							v81 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[744])) = uint8(v81)
							v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
							if v88 == int32(0) {
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
							}
						} else {
							v84 = int32(*(*uint8)(unsafe.Add(mBase, _consts[744])))
							if v84 == int32(0) {
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, _consts[748]))
								if v88 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[749])) = int32(1)
								}
							}
						}
						v95 = *(*int32)(unsafe.Add(mBase, _consts[716]))
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+76))
						*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(1)
						if v96 != 0 {
							F_s_lock(m, v95+int32(76), int32(515498), int32(3496), int32(326969))
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return
							} else {
								v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
								*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
								m.G0 = v7 + int32(16)
								return
							}
						} else {
							v107 = *(*int64)(unsafe.Add(mBase, _consts[746]))
							*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v107
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v117 = m.ExcPending
			if v117 != 0 {
				return
			} else {
				v118 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v118
				F_errmsg_internal(m, int32(213482), v7)
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return
				} else {
					F_errfinish(m, int32(515498), int32(3445), int32(326969))
					mBase = m.M
					v127 = m.ExcPending
					if v127 != 0 {
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
}
func F_XLogSendPhysical(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int64
	_ = v55
	var v57 int32
	_ = v57
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v119 int64
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int64
	_ = v136
	var v141 int64
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int64
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int64
	_ = v181
	var v185 int64
	_ = v185
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int64
	_ = v199
	var v203 int64
	_ = v203
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int64
	_ = v217
	var v221 int64
	_ = v221
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v243 int64
	_ = v243
	var v245 int32
	_ = v245
	var v249 int64
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int64
	_ = v281
	var v284 int64
	_ = v284
	var v286 int64
	_ = v286
	var v287 int64
	_ = v287
	var v290 int64
	_ = v290
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int64
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int64
	_ = v344
	var v346 int64
	_ = v346
	var v348 int64
	_ = v348
	var v351 int64
	_ = v351
	var v353 int64
	_ = v353
	var v355 int64
	_ = v355
	var v357 int64
	_ = v357
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int64
	_ = v393
	var v395 int64
	_ = v395
	var v397 int64
	_ = v397
	var v400 int64
	_ = v400
	var v402 int64
	_ = v402
	var v404 int64
	_ = v404
	var v406 int64
	_ = v406
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v451 int64
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v466 int64
	_ = v466
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v500 int64
	_ = v500
	var v503 int64
	_ = v503
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v519 int64
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v531 int64
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int64
	_ = v536
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int64
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int64
	_ = v561
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v599 int32
	_ = v599
	var v606 int32
	_ = v606
	var v608 int64
	_ = v608
	var v609 int64
	_ = v609
	var v613 int64
	_ = v613
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int64
	_ = v628
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v649 int64
	_ = v649
	var v650 int64
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int64
	_ = v714
	var v715 int64
	_ = v715
	var v723 int64
	_ = v723
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int64
	_ = v733
	var v735 int64
	_ = v735
	var v737 int64
	_ = v737
	var v740 int64
	_ = v740
	var v742 int64
	_ = v742
	var v744 int64
	_ = v744
	var v746 int64
	_ = v746
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int64
	_ = v777
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v802 int32
	_ = v802
	var v804 int64
	_ = v804
	var v809 int32
	_ = v809
	var v814 int64
	_ = v814
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	v16 = m.G0
	v18 = v16 - int32(128)
	m.G0 = v18
	v21 = *(*int32)(unsafe.Add(mBase, _consts[748]))
	if v21 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, _consts[751])))
	if v46 != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v26 == int32(4) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+76)) = int32(1)
	if v29 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_s_lock(m, v25+int32(76), int32(515498), int32(3869), int32(368433))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+76)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = int32(4)
	goto L1
L7:
	;
	return
L8:
	;
	goto L6
L9:
	;
	m.G0 = v18 + int32(128)
	return
L10:
	;
	v48 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[744])) = uint8(v48)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, _consts[752])))
	if v51 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v145 = m.G0
	v146 = int32(16)
	v147 = v145 - v146
	m.G0 = v147
	F___gettimeofday(m, v147)
	mBase = m.M
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v147)))
	v151 = int64(*(*int32)(unsafe.Add(mBase, uint32(v147)+8)))
	m.G0 = v147 + v146
	goto L45
L14:
	;
	v55 = *(*int64)(unsafe.Add(mBase, _consts[753]))
	v141 = v55
	goto L13
L15:
	;
	goto L16
L16:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, _consts[750])))
	if v57 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v121 = F_readTimeLineHistory(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L7
	} else {
		goto L42
	}
L18:
	;
	v63 = F_GetWalRcvFlushRecPtr(m, int32(0), v18+int32(88))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v102 = int32(4444592)
	v103 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v103)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v103)+280)) = v104
	*(*int64)(unsafe.Add(mBase, _consts[189])) = v104
	v109 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v109)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+272)) = v110
	*(*int64)(unsafe.Add(mBase, _consts[190])) = v110
	goto L40
L21:
	;
	v67 = F_GetXLogReplayRecPtr(m, v18+int32(32))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
	if v73 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v83 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+316))
	v81 = base.B2i32(v79 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v81)
	v83 = v81
	goto L26
L25:
	;
	v83 = int32(0)
	goto L26
L26:
	;
	goto L23
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+308))
	goto L30
L28:
	;
	goto L29
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[754]))
	if v93 != v70 {
		v120 = v70
		goto L17
	} else {
		goto L31
	}
L30:
	;
	v90 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[750])) = uint8(v90)
	v120 = v88
	goto L17
L31:
	;
	if base.Ui64(v67) < base.Ui64(v63) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v96 = v63
	goto L34
L33:
	;
	v96 = v67
	goto L34
L34:
	;
	if v70 == v69 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v98 = v96
	goto L37
L36:
	;
	v98 = v67
	goto L37
L37:
	;
	v141 = v98
	goto L13
L38:
	;
	v141 = v119
	goto L13
L40:
	;
	goto L41
L41:
	;
	v119 = *(*int64)(unsafe.Add(mBase, _consts[189]))
	goto L38
L42:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _consts[754]))
	v127 = F_tliSwitchPoint(m, v125, v121, int32(4459552))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	*(*int64)(unsafe.Add(mBase, _consts[753])) = v127
	F_list_free_deep(m, v121)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	v133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[752])) = uint8(v133)
	v136 = *(*int64)(unsafe.Add(mBase, _consts[753]))
	v141 = v136
	goto L13
L45:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, _consts[708])))
	if v161 != int32(1) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v243 = *(*int64)(unsafe.Add(mBase, _consts[746]))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, _consts[752])))
	if v245 != int32(1) {
		goto L58
	} else {
		goto L59
	}
L47:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _consts[755]))
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v165)))
	if v166 == v141 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v165))) = v141
	v170 = v165 + int32(8)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v165)+uint32(_consts[756])))
	v175 = base.I32_rem_s(v171+int32(1), int32(8192))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v165)+uint32(_consts[757])))
	if v175 == v176 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v180 = v170 + v175<<(uint(int32(4))%32)
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v180)))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+uint32(_consts[758]))) = v181
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v180)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+uint32(_consts[759]))) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v165)+uint32(_consts[757]))) = int32(-1)
	goto L51
L50:
	;
	goto L51
L51:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v165)+uint32(_consts[760])))
	if v192 == v175 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v198 = v170 + v175<<(uint(int32(4))%32)
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v198)))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+uint32(_consts[761]))) = v199
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v198)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+uint32(_consts[762]))) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v165)+uint32(_consts[760]))) = int32(-1)
	goto L54
L53:
	;
	goto L54
L54:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v165)+uint32(_consts[763])))
	if v210 == v175 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v216 = v170 + v175<<(uint(int32(4))%32)
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v216)))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+uint32(_consts[764]))) = v217
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v216)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+uint32(_consts[765]))) = v221
	*(*int32)(unsafe.Add(mBase, uint32(v165)+uint32(_consts[763]))) = int32(-1)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v226 = int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v170+v171<<(uint(v226)%32)))) = v141
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v165)+uint32(_consts[756])))
	*(*int64)(unsafe.Add(mBase, uint32(v165+v230<<(uint(v226)%32))+16)) = v151 + v150*int64(1000000) - int64(946684800000000)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+uint32(_consts[756]))) = v175
	goto L46
L58:
	;
	if base.Ui64(v141) <= base.Ui64(v243) {
		goto L70
	} else {
		goto L71
	}
L59:
	;
	v249 = *(*int64)(unsafe.Add(mBase, _consts[753]))
	if base.Ui64(v243) < base.Ui64(v249) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _consts[766]))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+1168))
	if int32(0) <= v253 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v252)+1168))
	v257 = F_close(m, v256)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v252)+1168)) = int32(-1)
	goto L64
L62:
	;
	goto L63
L63:
	;
	v261 = int32(0)
	v264 = *(*int32)(unsafe.Add(mBase, _consts[467]))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+20))
	m.T0[v265].(func(*base.Module, int32, int32, int32))(m, int32(99), v261, v261)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L7
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	v269 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[744])) = uint8(v269)
	*(*uint8)(unsafe.Add(mBase, _consts[751])) = uint8(v269)
	v276 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	if v276 == int32(0) {
		goto L9
	} else {
		goto L67
	}
L67:
	;
	v281 = *(*int64)(unsafe.Add(mBase, _consts[753]))
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+4)) = uint32(v281)
	v284 = *(*int64)(unsafe.Add(mBase, _consts[746]))
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+12)) = uint32(v284)
	v286 = int64(32)
	v287 = int64(base.Ui64(v281) >> (uint(v286) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v18))) = uint32(v287)
	v290 = int64(base.Ui64(v284) >> (uint(v286) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+8)) = uint32(v290)
	F_errmsg_internal(m, int32(707484), v18)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(515498), int32(3270), int32(326884))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	goto L9
L70:
	;
	v302 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[744])) = uint8(v302)
	goto L9
L71:
	;
	goto L72
L72:
	;
	v306 = v243 + int64(131072)
	v307 = base.B2i32(base.Ui64(v141) <= base.Ui64(v306))
	v310 = v307 & (v245 ^ int32(-1))
	*(*uint8)(unsafe.Add(mBase, _consts[744])) = uint8(v310)
	v312 = int32(4459472)
	v313 = *(*int32)(unsafe.Add(mBase, _consts[767]))
	v314 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v313))) = uint8(v314)
	*(*int32)(unsafe.Add(mBase, _consts[768])) = v314
	*(*int32)(unsafe.Add(mBase, _consts[769])) = v314
	goto L73
L73:
	;
	F_enlargeStringInfo(m, int32(4459472), int32(1))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	v324 = int32(4459476)
	v325 = *(*int32)(unsafe.Add(mBase, _consts[769]))
	v326 = int32(4459472)
	v327 = *(*int32)(unsafe.Add(mBase, _consts[767]))
	v329 = int32(119)
	*(*uint8)(unsafe.Add(mBase, uint32(v325+v327))) = uint8(v329)
	*(*int32)(unsafe.Add(mBase, _consts[769])) = v325 + int32(1)
	F_enlargeStringInfo(m, v326, int32(8))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	v339 = int32(4459476)
	v340 = *(*int32)(unsafe.Add(mBase, _consts[769]))
	v341 = int32(4459472)
	v342 = *(*int32)(unsafe.Add(mBase, _consts[767]))
	v344 = int64(56)
	v346 = int64(65280)
	v348 = int64(40)
	v351 = int64(16711680)
	v353 = int64(24)
	v355 = int64(4278190080)
	v357 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v340+v342))) = v243<<(uint(v344)%64) | v243&v346<<(uint(v348)%64) | (v243&v351<<(uint(v353)%64) | v243&v355<<(uint(v357)%64)) | (int64(base.Ui64(v243)>>(uint(v357)%64))&v355 | int64(base.Ui64(v243)>>(uint(v353)%64))&v351 | (int64(base.Ui64(v243)>>(uint(v348)%64))&v346 | int64(base.Ui64(v243)>>(uint(v344)%64))))
	v381 = int32(8)
	*(*int32)(unsafe.Add(mBase, _consts[769])) = v340 + v381
	F_enlargeStringInfo(m, v341, v381)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	v388 = int32(4459476)
	v389 = *(*int32)(unsafe.Add(mBase, _consts[769]))
	v390 = int32(4459472)
	v391 = *(*int32)(unsafe.Add(mBase, _consts[767]))
	v393 = int64(56)
	v395 = int64(65280)
	v397 = int64(40)
	v400 = int64(16711680)
	v402 = int64(24)
	v404 = int64(4278190080)
	v406 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v389+v391))) = v141<<(uint(v393)%64) | v141&v395<<(uint(v397)%64) | (v141&v400<<(uint(v402)%64) | v141&v404<<(uint(v406)%64)) | (int64(base.Ui64(v141)>>(uint(v406)%64))&v404 | int64(base.Ui64(v141)>>(uint(v402)%64))&v400 | (int64(base.Ui64(v141)>>(uint(v397)%64))&v395 | int64(base.Ui64(v141)>>(uint(v393)%64))))
	v430 = int32(8)
	*(*int32)(unsafe.Add(mBase, _consts[769])) = v389 + v430
	F_enlargeStringInfo(m, v390, v430)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L7
	} else {
		goto L77
	}
L77:
	;
	v437 = int32(4459476)
	v438 = *(*int32)(unsafe.Add(mBase, _consts[769]))
	v439 = int32(4459472)
	v440 = *(*int32)(unsafe.Add(mBase, _consts[767]))
	*(*int64)(unsafe.Add(mBase, uint32(v438+v440))) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[769])) = v438 + int32(8)
	if base.Ui64(v141) <= base.Ui64(v306) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v451 = v141
	goto L80
L79:
	;
	v451 = v306 & int64(-8192)
	goto L80
L80:
	;
	v453 = base.I32_wrap_i64(v451 - v243)
	F_enlargeStringInfo(m, v439, v453)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L7
	} else {
		goto L81
	}
L81:
	;
	v457 = v453
	v466 = v243
	goto L82
L82:
	;
	v471 = int32(4459476)
	v473 = *(*int32)(unsafe.Add(mBase, _consts[767]))
	v475 = *(*int32)(unsafe.Add(mBase, _consts[769]))
	v476 = v473 + v475
	v478 = *(*int32)(unsafe.Add(mBase, _consts[766]))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)+1184))
	v480 = m.G0
	v482 = v480 - int32(16)
	m.G0 = v482
	v485 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v487 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
	if v487 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	v688 = int32(4459476)
	v690 = *(*int32)(unsafe.Add(mBase, _consts[769]))
	v691 = v690 + v629
	*(*int32)(unsafe.Add(mBase, _consts[769])) = v691
	v694 = *(*int32)(unsafe.Add(mBase, _consts[767]))
	v696 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v694+v691))) = uint8(v696)
	v698 = int32(4459504)
	v699 = *(*int32)(unsafe.Add(mBase, _consts[770]))
	*(*uint8)(unsafe.Add(mBase, uint32(v699))) = uint8(v696)
	*(*int32)(unsafe.Add(mBase, _consts[771])) = v696
	*(*int32)(unsafe.Add(mBase, _consts[772])) = v696
	goto L125
L84:
	;
	v624 = *(*int32)(unsafe.Add(mBase, _consts[769]))
	v625 = v599 + v624
	*(*int32)(unsafe.Add(mBase, _consts[769])) = v625
	v628 = v466 + base.I64_extend_i32_u(v599)
	v629 = v457 - v599
	if v629 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L7
	} else {
		goto L107
	}
L86:
	;
	m.G0 = v482 + int32(16)
	goto L84
L87:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v485)+316))
	v493 = base.B2i32(v491 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v493)
	if v491 != int32(2) {
		v599 = int32(0)
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v485)+308))
	if v479 != v498 {
		v599 = int32(0)
		goto L86
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	v500 = *(*int64)(unsafe.Add(mBase, uint32(v485)+264))
	*(*int64)(unsafe.Add(mBase, uint32(v485)+264)) = v500
	v503 = v466 + base.I64_extend_i32_u(v457)
	if base.Ui64(v500) < base.Ui64(v503) {
		goto L85
	} else {
		goto L92
	}
L92:
	;
	if v457 == int32(0) {
		v572 = v476
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v599 = v572 - v476
	goto L86
L94:
	;
	v507 = v457
	v511 = v476
	v519 = v466
	goto L95
L95:
	;
	v523 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)+300))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v523)+304))
	v531 = base.I64_rem_u_s(int64(base.Ui64(v519)>>(uint(int64(13))%64)), base.I64_extend_i32_s(v527+int32(1)))
	v532 = base.I32_wrap_i64(v531)
	v534 = v532 << (uint(int32(3)) % 32)
	v535 = v524 + v534
	v536 = *(*int64)(unsafe.Add(mBase, uint32(v535)))
	*(*int64)(unsafe.Add(mBase, uint32(v535))) = v536
	v541 = base.I32_wrap_i64(v519) & int32(8191)
	v542 = int32(8192) - v541
	v544 = v519 + base.I64_extend_i32_u(v542)
	if v536 != v544 {
		v572 = v511
		goto L93
	} else {
		goto L97
	}
L96:
	;
	v572 = v564
	goto L93
L97:
	;
	v547 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)+296))
	if base.Ui32(v507) < base.Ui32(v542) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v554 = v507
	goto L100
L99:
	;
	v554 = v542
	goto L100
L100:
	;
	if v554 != 0 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v558 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v558)+300))
	v560 = v559 + v534
	v561 = *(*int64)(unsafe.Add(mBase, uint32(v560)))
	*(*int64)(unsafe.Add(mBase, uint32(v560))) = v561
	if v561 != v544 {
		v572 = v511
		goto L93
	} else {
		goto L105
	}
L102:
	;
	v555 = F__emscripten_memcpy_bulkmem(m, v511, v548+v532<<(uint(int32(13))%32)+v541, v554)
	mBase = m.M
	v556 = v555
	goto L104
L103:
	;
	v556 = v511
	goto L104
L104:
	;
	goto L101
L105:
	;
	v564 = v554 + v556
	v567 = v507 - v554
	if v567 != 0 {
		v507 = v567
		v511 = v564
		v519 = v519 + base.I64_extend_i32_u(v554)
		goto L95
	} else {
		goto L106
	}
L106:
	;
	goto L96
L107:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v482)+12)) = uint32(v500)
	v608 = int64(32)
	v609 = int64(base.Ui64(v500) >> (uint(v608) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v482)+8)) = uint32(v609)
	*(*uint32)(unsafe.Add(mBase, uint32(v482)+4)) = uint32(v503)
	v613 = int64(base.Ui64(v503) >> (uint(v608) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v482))) = uint32(v613)
	F_errmsg(m, int32(536351), v482)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L7
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(518546), int32(1773), int32(142854))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	v648 = *(*int32)(unsafe.Add(mBase, _consts[766]))
	v649 = int64(*(*int32)(unsafe.Add(mBase, uint32(v648)+1160)))
	v650 = base.I64_div_u_s(v628, v649)
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v648)+1184))
	F_CheckXLogRemoved(m, v650, v651)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L7
	} else {
		goto L115
	}
L111:
	;
	v633 = *(*int32)(unsafe.Add(mBase, _consts[766]))
	v635 = *(*int32)(unsafe.Add(mBase, _consts[767]))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v633)+1184))
	v640 = F_WALRead(m, v633, v635+v625, v628, v629, v637, v18+int32(88))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	if v640 != 0 {
		goto L110
	} else {
		goto L113
	}
L113:
	;
	F_WALReadRaiseError(m, v18+int32(88))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L7
	} else {
		goto L114
	}
L114:
	;
	goto L110
L115:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, _consts[750])))
	if v655 != int32(1) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	goto L83
L117:
	;
	v659 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v659)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v659)+76)) = int32(1)
	if v660 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	F_s_lock(m, v659+int32(76), int32(515498), int32(3367), int32(326884))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L7
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v670 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v659)+76)) = v670
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v659)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v659)+16)) = uint8(v670)
	if v672 != int32(1) {
		goto L116
	} else {
		goto L122
	}
L121:
	;
	goto L120
L122:
	;
	v678 = *(*int32)(unsafe.Add(mBase, _consts[766]))
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v678)+1168))
	if v679 < int32(0) {
		goto L116
	} else {
		goto L123
	}
L123:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v678)+1168))
	v683 = F_close(m, v682)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v678)+1168)) = int32(-1)
	goto L124
L124:
	;
	v457 = v629
	v466 = v628
	goto L82
L125:
	;
	v709 = m.G0
	v710 = int32(16)
	v711 = v709 - v710
	m.G0 = v711
	F___gettimeofday(m, v711)
	mBase = m.M
	v714 = *(*int64)(unsafe.Add(mBase, uint32(v711)))
	v715 = int64(*(*int32)(unsafe.Add(mBase, uint32(v711)+8)))
	m.G0 = v711 + v710
	v723 = v715 + v714*int64(1000000) - int64(946684800000000)
	goto L126
L126:
	;
	F_enlargeStringInfo(m, int32(4459504), int32(8))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L7
	} else {
		goto L127
	}
L127:
	;
	v728 = int32(4459508)
	v729 = *(*int32)(unsafe.Add(mBase, _consts[772]))
	v730 = int32(4459504)
	v731 = *(*int32)(unsafe.Add(mBase, _consts[770]))
	v733 = int64(56)
	v735 = int64(65280)
	v737 = int64(40)
	v740 = int64(16711680)
	v742 = int64(24)
	v744 = int64(4278190080)
	v746 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v729+v731))) = v723<<(uint(v733)%64) | v723&v735<<(uint(v737)%64) | (v723&v740<<(uint(v742)%64) | v723&v744<<(uint(v746)%64)) | (int64(base.Ui64(v723)>>(uint(v746)%64))&v744 | int64(base.Ui64(v723)>>(uint(v742)%64))&v740 | (int64(base.Ui64(v723)>>(uint(v737)%64))&v735 | int64(base.Ui64(v723)>>(uint(v733)%64))))
	*(*int32)(unsafe.Add(mBase, _consts[772])) = v729 + int32(8)
	v773 = int32(4459472)
	v774 = *(*int32)(unsafe.Add(mBase, _consts[767]))
	v776 = *(*int32)(unsafe.Add(mBase, _consts[770]))
	v777 = *(*int64)(unsafe.Add(mBase, uint32(v776)))
	*(*int64)(unsafe.Add(mBase, uint32(v774)+17)) = v777
	v781 = *(*int32)(unsafe.Add(mBase, _consts[767]))
	v783 = *(*int32)(unsafe.Add(mBase, _consts[769]))
	v785 = *(*int32)(unsafe.Add(mBase, _consts[467]))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v785)+20))
	m.T0[v786].(func(*base.Module, int32, int32, int32))(m, int32(100), v781, v783)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L7
	} else {
		goto L128
	}
L128:
	;
	*(*int64)(unsafe.Add(mBase, _consts[746])) = v451
	v792 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v792)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v792)+76)) = int32(1)
	if v793 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	F_s_lock(m, v792+int32(76), int32(515498), int32(3399), int32(326884))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L7
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v804 = *(*int64)(unsafe.Add(mBase, _consts[746]))
	*(*int32)(unsafe.Add(mBase, uint32(v792)+76)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v792)+8)) = v804
	v809 = int32(*(*uint8)(unsafe.Add(mBase, _consts[773])))
	if v809 != int32(1) {
		goto L9
	} else {
		goto L133
	}
L132:
	;
	goto L131
L133:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+20)) = uint32(v804)
	v814 = int64(base.Ui64(v804) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+16)) = uint32(v814)
	v822 = F_pg_snprintf(m, v18+int32(32), int32(50), int32(536653), v18+int32(16))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L7
	} else {
		goto L134
	}
L134:
	;
	v826 = F_strlen(m, v18+int32(32))
	mBase = m.M
	goto L9
}
func F_XLogSetAsyncXactLSN(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	v6 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+440)) = int32(1)
	if v7 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[2]))
		F_s_lock(m, v11+int32(440), int32(518546), int32(2618), int32(549283))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _consts[2]))
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)+216))
			if base.Ui64(l0) <= base.Ui64(v21) {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+440)) = int32(0)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+440)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(v20)+216)) = l0
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+321)))
				if v28 != 0 {
					v55 = *(*int32)(unsafe.Add(mBase, _consts[152]))
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+60))
					if v56 == int32(-1) {
						return
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
						F_SetLatch(m, v59+v56*int32(640)+int32(20))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v29 = *(*int64)(unsafe.Add(mBase, uint32(v20)+280))
					*(*int64)(unsafe.Add(mBase, uint32(v20)+280)) = v29
					*(*int64)(unsafe.Add(mBase, _consts[189])) = v29
					v34 = *(*int32)(unsafe.Add(mBase, _consts[2]))
					v35 = *(*int64)(unsafe.Add(mBase, uint32(v34)+272))
					*(*int64)(unsafe.Add(mBase, uint32(v34)+272)) = v35
					*(*int64)(unsafe.Add(mBase, _consts[190])) = v35
					v40 = *(*int32)(unsafe.Add(mBase, _consts[206]))
					if v40 == int32(0) {
						v55 = *(*int32)(unsafe.Add(mBase, _consts[152]))
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+60))
						if v56 == int32(-1) {
							return
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
							F_SetLatch(m, v59+v56*int32(640)+int32(20))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v43 = int64(13)
						v46 = *(*int64)(unsafe.Add(mBase, _consts[189]))
						if base.I32_wrap_i64(int64(base.Ui64(l0)>>(uint(v43)%64))-int64(base.Ui64(v46)>>(uint(v43)%64))) < v40 {
							return
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, _consts[152]))
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+60))
							if v56 == int32(-1) {
								return
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
								F_SetLatch(m, v59+v56*int32(640)+int32(20))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[2]))
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)+216))
		if base.Ui64(l0) <= base.Ui64(v21) {
			*(*int32)(unsafe.Add(mBase, uint32(v20)+440)) = int32(0)
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v20)+440)) = int32(0)
			*(*int64)(unsafe.Add(mBase, uint32(v20)+216)) = l0
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+321)))
			if v28 != 0 {
				v55 = *(*int32)(unsafe.Add(mBase, _consts[152]))
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+60))
				if v56 == int32(-1) {
					return
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
					F_SetLatch(m, v59+v56*int32(640)+int32(20))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v29 = *(*int64)(unsafe.Add(mBase, uint32(v20)+280))
				*(*int64)(unsafe.Add(mBase, uint32(v20)+280)) = v29
				*(*int64)(unsafe.Add(mBase, _consts[189])) = v29
				v34 = *(*int32)(unsafe.Add(mBase, _consts[2]))
				v35 = *(*int64)(unsafe.Add(mBase, uint32(v34)+272))
				*(*int64)(unsafe.Add(mBase, uint32(v34)+272)) = v35
				*(*int64)(unsafe.Add(mBase, _consts[190])) = v35
				v40 = *(*int32)(unsafe.Add(mBase, _consts[206]))
				if v40 == int32(0) {
					v55 = *(*int32)(unsafe.Add(mBase, _consts[152]))
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+60))
					if v56 == int32(-1) {
						return
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
						F_SetLatch(m, v59+v56*int32(640)+int32(20))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v43 = int64(13)
					v46 = *(*int64)(unsafe.Add(mBase, _consts[189]))
					if base.I32_wrap_i64(int64(base.Ui64(l0)>>(uint(v43)%64))-int64(base.Ui64(v46)>>(uint(v43)%64))) < v40 {
						return
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, _consts[152]))
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+60))
						if v56 == int32(-1) {
							return
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
							F_SetLatch(m, v59+v56*int32(640)+int32(20))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_XLogWalRcvSendReply(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v45 int64
	_ = v45
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v58 int64
	_ = v58
	var v64 int64
	_ = v64
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int64
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int64
	_ = v158
	var v160 int64
	_ = v160
	var v162 int64
	_ = v162
	var v165 int64
	_ = v165
	var v167 int64
	_ = v167
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int64
	_ = v209
	var v211 int64
	_ = v211
	var v213 int64
	_ = v213
	var v216 int64
	_ = v216
	var v218 int64
	_ = v218
	var v220 int64
	_ = v220
	var v222 int64
	_ = v222
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int64
	_ = v259
	var v260 int64
	_ = v260
	var v268 int64
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int64
	_ = v278
	var v280 int64
	_ = v280
	var v282 int64
	_ = v282
	var v285 int64
	_ = v285
	var v287 int64
	_ = v287
	var v289 int64
	_ = v289
	var v291 int64
	_ = v291
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int64
	_ = v341
	var v346 int32
	_ = v346
	var v349 int64
	_ = v349
	var v352 int64
	_ = v352
	var v354 int64
	_ = v354
	var v355 int64
	_ = v355
	var v358 int64
	_ = v358
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	v2 = l1
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l0 == int32(0) {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[735]))
		if v14 <= int32(0) {
			m.G0 = v9 + int32(32)
			return
		} else {
			v20 = m.G0
			v21 = int32(16)
			v22 = v20 - v21
			m.G0 = v22
			F___gettimeofday(m, v22)
			mBase = m.M
			v25 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
			v26 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22)+8)))
			m.G0 = v22 + v21
			v34 = v26 + v25*int64(1000000) - int64(946684800000000)
			v36 = *(*int64)(unsafe.Add(mBase, _consts[736]))
			if l0 != 0 {
				v38 = *(*int64)(unsafe.Add(mBase, _consts[737]))
				v50 = v38
				*(*int64)(unsafe.Add(mBase, _consts[738])) = v36
				*(*int64)(unsafe.Add(mBase, _consts[739])) = v50
				v58 = int64(*(*int32)(unsafe.Add(mBase, _consts[735])))
				if v58 <= int64(0) {
					v64 = int64(9223372036854775807)
				} else {
					v64 = v58*int64(1000000) + v34
				}
				*(*int64)(unsafe.Add(mBase, _consts[740])) = v64
				v67 = F_GetXLogReplayRecPtr(m, int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					v69 = int32(4459328)
					v70 = *(*int32)(unsafe.Add(mBase, _consts[741]))
					v71 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v71)
					*(*int32)(unsafe.Add(mBase, _consts[742])) = v71
					*(*int32)(unsafe.Add(mBase, _consts[743])) = v71
					F_enlargeStringInfo(m, int32(4459328), int32(1))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						v81 = int32(4459328)
						v82 = *(*int32)(unsafe.Add(mBase, _consts[741]))
						v83 = int32(4459332)
						v84 = *(*int32)(unsafe.Add(mBase, _consts[743]))
						v86 = int32(114)
						*(*uint8)(unsafe.Add(mBase, uint32(v82+v84))) = uint8(v86)
						v90 = *(*int32)(unsafe.Add(mBase, _consts[743]))
						*(*int32)(unsafe.Add(mBase, _consts[743])) = v90 + int32(1)
						v95 = *(*int64)(unsafe.Add(mBase, _consts[738]))
						F_enlargeStringInfo(m, v81, int32(8))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							v100 = int32(4459328)
							v101 = *(*int32)(unsafe.Add(mBase, _consts[741]))
							v102 = int32(4459332)
							v103 = *(*int32)(unsafe.Add(mBase, _consts[743]))
							v105 = int64(56)
							v107 = int64(65280)
							v109 = int64(40)
							v112 = int64(16711680)
							v114 = int64(24)
							v116 = int64(4278190080)
							v118 = int64(8)
							*(*int64)(unsafe.Add(mBase, uint32(v101+v103))) = v95<<(uint(v105)%64) | v95&v107<<(uint(v109)%64) | (v95&v112<<(uint(v114)%64) | v95&v116<<(uint(v118)%64)) | (int64(base.Ui64(v95)>>(uint(v118)%64))&v116 | int64(base.Ui64(v95)>>(uint(v114)%64))&v112 | (int64(base.Ui64(v95)>>(uint(v109)%64))&v107 | int64(base.Ui64(v95)>>(uint(v105)%64))))
							v143 = *(*int32)(unsafe.Add(mBase, _consts[743]))
							v144 = int32(8)
							*(*int32)(unsafe.Add(mBase, _consts[743])) = v143 + v144
							v148 = *(*int64)(unsafe.Add(mBase, _consts[739]))
							F_enlargeStringInfo(m, v100, v144)
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return
							} else {
								v153 = int32(4459328)
								v154 = *(*int32)(unsafe.Add(mBase, _consts[741]))
								v155 = int32(4459332)
								v156 = *(*int32)(unsafe.Add(mBase, _consts[743]))
								v158 = int64(56)
								v160 = int64(65280)
								v162 = int64(40)
								v165 = int64(16711680)
								v167 = int64(24)
								v169 = int64(4278190080)
								v171 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v154+v156))) = v148<<(uint(v158)%64) | v148&v160<<(uint(v162)%64) | (v148&v165<<(uint(v167)%64) | v148&v169<<(uint(v171)%64)) | (int64(base.Ui64(v148)>>(uint(v171)%64))&v169 | int64(base.Ui64(v148)>>(uint(v167)%64))&v165 | (int64(base.Ui64(v148)>>(uint(v162)%64))&v160 | int64(base.Ui64(v148)>>(uint(v158)%64))))
								v196 = *(*int32)(unsafe.Add(mBase, _consts[743]))
								v197 = int32(8)
								*(*int32)(unsafe.Add(mBase, _consts[743])) = v196 + v197
								F_enlargeStringInfo(m, v153, v197)
								mBase = m.M
								v203 = m.ExcPending
								if v203 != 0 {
									return
								} else {
									v205 = *(*int32)(unsafe.Add(mBase, _consts[741]))
									v206 = int32(4459332)
									v207 = *(*int32)(unsafe.Add(mBase, _consts[743]))
									v209 = int64(56)
									v211 = int64(65280)
									v213 = int64(40)
									v216 = int64(16711680)
									v218 = int64(24)
									v220 = int64(4278190080)
									v222 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v205+v207))) = v67<<(uint(v209)%64) | v67&v211<<(uint(v213)%64) | (v67&v216<<(uint(v218)%64) | v67&v220<<(uint(v222)%64)) | (int64(base.Ui64(v67)>>(uint(v222)%64))&v220 | int64(base.Ui64(v67)>>(uint(v218)%64))&v216 | (int64(base.Ui64(v67)>>(uint(v213)%64))&v211 | int64(base.Ui64(v67)>>(uint(v209)%64))))
									v247 = *(*int32)(unsafe.Add(mBase, _consts[743]))
									*(*int32)(unsafe.Add(mBase, _consts[743])) = v247 + int32(8)
									v254 = m.G0
									v255 = int32(16)
									v256 = v254 - v255
									m.G0 = v256
									F___gettimeofday(m, v256)
									mBase = m.M
									v259 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
									v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+8)))
									m.G0 = v256 + v255
									v268 = v260 + v259*int64(1000000) - int64(946684800000000)
									F_enlargeStringInfo(m, int32(4459328), int32(8))
									mBase = m.M
									v272 = m.ExcPending
									if v272 != 0 {
										return
									} else {
										v273 = int32(4459328)
										v274 = *(*int32)(unsafe.Add(mBase, _consts[741]))
										v275 = int32(4459332)
										v276 = *(*int32)(unsafe.Add(mBase, _consts[743]))
										v278 = int64(56)
										v280 = int64(65280)
										v282 = int64(40)
										v285 = int64(16711680)
										v287 = int64(24)
										v289 = int64(4278190080)
										v291 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v274+v276))) = v268<<(uint(v278)%64) | v268&v280<<(uint(v282)%64) | (v268&v285<<(uint(v287)%64) | v268&v289<<(uint(v291)%64)) | (int64(base.Ui64(v268)>>(uint(v291)%64))&v289 | int64(base.Ui64(v268)>>(uint(v287)%64))&v285 | (int64(base.Ui64(v268)>>(uint(v282)%64))&v280 | int64(base.Ui64(v268)>>(uint(v278)%64))))
										v316 = *(*int32)(unsafe.Add(mBase, _consts[743]))
										*(*int32)(unsafe.Add(mBase, _consts[743])) = v316 + int32(8)
										F_enlargeStringInfo(m, v273, int32(1))
										mBase = m.M
										v323 = m.ExcPending
										if v323 != 0 {
											return
										} else {
											v325 = *(*int32)(unsafe.Add(mBase, _consts[741]))
											v326 = int32(4459332)
											v327 = *(*int32)(unsafe.Add(mBase, _consts[743]))
											*(*uint8)(unsafe.Add(mBase, uint32(v325+v327))) = uint8(v2)
											v332 = *(*int32)(unsafe.Add(mBase, _consts[743]))
											*(*int32)(unsafe.Add(mBase, _consts[743])) = v332 + int32(1)
											v338 = F_errstart(m, int32(13), int32(0))
											mBase = m.M
											v339 = m.ExcPending
											if v339 != 0 {
												return
											} else {
												if v338 != 0 {
													v341 = int64(base.Ui64(v67) >> (uint(int64(32)) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)) = uint32(v341)
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v67)
													if v2 != 0 {
														v346 = int32(703410)
													} else {
														v346 = int32(785690)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v346
													v349 = *(*int64)(unsafe.Add(mBase, _consts[738]))
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v349)
													v352 = *(*int64)(unsafe.Add(mBase, _consts[739]))
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v352)
													v354 = int64(32)
													v355 = int64(base.Ui64(v349) >> (uint(v354) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v355)
													v358 = int64(base.Ui64(v352) >> (uint(v354) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v358)
													F_errmsg_internal(m, int32(184919), v9)
													mBase = m.M
													v362 = m.ExcPending
													if v362 != 0 {
														return
													} else {
														F_errfinish(m, int32(515166), int32(1145), int32(19704))
														mBase = m.M
														v367 = m.ExcPending
														if v367 != 0 {
															return
														} else {
															v371 = *(*int32)(unsafe.Add(mBase, _consts[734]))
															v373 = *(*int32)(unsafe.Add(mBase, _consts[741]))
															v375 = *(*int32)(unsafe.Add(mBase, _consts[743]))
															v377 = *(*int32)(unsafe.Add(mBase, _consts[495]))
															v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
															m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
															mBase = m.M
															v380 = m.ExcPending
															if v380 != 0 {
																return
															} else {
																m.G0 = v9 + int32(32)
																return
															}
														}
													}
												} else {
													v371 = *(*int32)(unsafe.Add(mBase, _consts[734]))
													v373 = *(*int32)(unsafe.Add(mBase, _consts[741]))
													v375 = *(*int32)(unsafe.Add(mBase, _consts[743]))
													v377 = *(*int32)(unsafe.Add(mBase, _consts[495]))
													v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
													m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
													mBase = m.M
													v380 = m.ExcPending
													if v380 != 0 {
														return
													} else {
														m.G0 = v9 + int32(32)
														return
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v40 = *(*int64)(unsafe.Add(mBase, _consts[737]))
				v42 = *(*int64)(unsafe.Add(mBase, _consts[738]))
				if v42 != v36 {
					v50 = v40
					*(*int64)(unsafe.Add(mBase, _consts[738])) = v36
					*(*int64)(unsafe.Add(mBase, _consts[739])) = v50
					v58 = int64(*(*int32)(unsafe.Add(mBase, _consts[735])))
					if v58 <= int64(0) {
						v64 = int64(9223372036854775807)
					} else {
						v64 = v58*int64(1000000) + v34
					}
					*(*int64)(unsafe.Add(mBase, _consts[740])) = v64
					v67 = F_GetXLogReplayRecPtr(m, int32(0))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						v69 = int32(4459328)
						v70 = *(*int32)(unsafe.Add(mBase, _consts[741]))
						v71 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v71)
						*(*int32)(unsafe.Add(mBase, _consts[742])) = v71
						*(*int32)(unsafe.Add(mBase, _consts[743])) = v71
						F_enlargeStringInfo(m, int32(4459328), int32(1))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							v81 = int32(4459328)
							v82 = *(*int32)(unsafe.Add(mBase, _consts[741]))
							v83 = int32(4459332)
							v84 = *(*int32)(unsafe.Add(mBase, _consts[743]))
							v86 = int32(114)
							*(*uint8)(unsafe.Add(mBase, uint32(v82+v84))) = uint8(v86)
							v90 = *(*int32)(unsafe.Add(mBase, _consts[743]))
							*(*int32)(unsafe.Add(mBase, _consts[743])) = v90 + int32(1)
							v95 = *(*int64)(unsafe.Add(mBase, _consts[738]))
							F_enlargeStringInfo(m, v81, int32(8))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								v100 = int32(4459328)
								v101 = *(*int32)(unsafe.Add(mBase, _consts[741]))
								v102 = int32(4459332)
								v103 = *(*int32)(unsafe.Add(mBase, _consts[743]))
								v105 = int64(56)
								v107 = int64(65280)
								v109 = int64(40)
								v112 = int64(16711680)
								v114 = int64(24)
								v116 = int64(4278190080)
								v118 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v101+v103))) = v95<<(uint(v105)%64) | v95&v107<<(uint(v109)%64) | (v95&v112<<(uint(v114)%64) | v95&v116<<(uint(v118)%64)) | (int64(base.Ui64(v95)>>(uint(v118)%64))&v116 | int64(base.Ui64(v95)>>(uint(v114)%64))&v112 | (int64(base.Ui64(v95)>>(uint(v109)%64))&v107 | int64(base.Ui64(v95)>>(uint(v105)%64))))
								v143 = *(*int32)(unsafe.Add(mBase, _consts[743]))
								v144 = int32(8)
								*(*int32)(unsafe.Add(mBase, _consts[743])) = v143 + v144
								v148 = *(*int64)(unsafe.Add(mBase, _consts[739]))
								F_enlargeStringInfo(m, v100, v144)
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return
								} else {
									v153 = int32(4459328)
									v154 = *(*int32)(unsafe.Add(mBase, _consts[741]))
									v155 = int32(4459332)
									v156 = *(*int32)(unsafe.Add(mBase, _consts[743]))
									v158 = int64(56)
									v160 = int64(65280)
									v162 = int64(40)
									v165 = int64(16711680)
									v167 = int64(24)
									v169 = int64(4278190080)
									v171 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v154+v156))) = v148<<(uint(v158)%64) | v148&v160<<(uint(v162)%64) | (v148&v165<<(uint(v167)%64) | v148&v169<<(uint(v171)%64)) | (int64(base.Ui64(v148)>>(uint(v171)%64))&v169 | int64(base.Ui64(v148)>>(uint(v167)%64))&v165 | (int64(base.Ui64(v148)>>(uint(v162)%64))&v160 | int64(base.Ui64(v148)>>(uint(v158)%64))))
									v196 = *(*int32)(unsafe.Add(mBase, _consts[743]))
									v197 = int32(8)
									*(*int32)(unsafe.Add(mBase, _consts[743])) = v196 + v197
									F_enlargeStringInfo(m, v153, v197)
									mBase = m.M
									v203 = m.ExcPending
									if v203 != 0 {
										return
									} else {
										v205 = *(*int32)(unsafe.Add(mBase, _consts[741]))
										v206 = int32(4459332)
										v207 = *(*int32)(unsafe.Add(mBase, _consts[743]))
										v209 = int64(56)
										v211 = int64(65280)
										v213 = int64(40)
										v216 = int64(16711680)
										v218 = int64(24)
										v220 = int64(4278190080)
										v222 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v205+v207))) = v67<<(uint(v209)%64) | v67&v211<<(uint(v213)%64) | (v67&v216<<(uint(v218)%64) | v67&v220<<(uint(v222)%64)) | (int64(base.Ui64(v67)>>(uint(v222)%64))&v220 | int64(base.Ui64(v67)>>(uint(v218)%64))&v216 | (int64(base.Ui64(v67)>>(uint(v213)%64))&v211 | int64(base.Ui64(v67)>>(uint(v209)%64))))
										v247 = *(*int32)(unsafe.Add(mBase, _consts[743]))
										*(*int32)(unsafe.Add(mBase, _consts[743])) = v247 + int32(8)
										v254 = m.G0
										v255 = int32(16)
										v256 = v254 - v255
										m.G0 = v256
										F___gettimeofday(m, v256)
										mBase = m.M
										v259 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
										v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+8)))
										m.G0 = v256 + v255
										v268 = v260 + v259*int64(1000000) - int64(946684800000000)
										F_enlargeStringInfo(m, int32(4459328), int32(8))
										mBase = m.M
										v272 = m.ExcPending
										if v272 != 0 {
											return
										} else {
											v273 = int32(4459328)
											v274 = *(*int32)(unsafe.Add(mBase, _consts[741]))
											v275 = int32(4459332)
											v276 = *(*int32)(unsafe.Add(mBase, _consts[743]))
											v278 = int64(56)
											v280 = int64(65280)
											v282 = int64(40)
											v285 = int64(16711680)
											v287 = int64(24)
											v289 = int64(4278190080)
											v291 = int64(8)
											*(*int64)(unsafe.Add(mBase, uint32(v274+v276))) = v268<<(uint(v278)%64) | v268&v280<<(uint(v282)%64) | (v268&v285<<(uint(v287)%64) | v268&v289<<(uint(v291)%64)) | (int64(base.Ui64(v268)>>(uint(v291)%64))&v289 | int64(base.Ui64(v268)>>(uint(v287)%64))&v285 | (int64(base.Ui64(v268)>>(uint(v282)%64))&v280 | int64(base.Ui64(v268)>>(uint(v278)%64))))
											v316 = *(*int32)(unsafe.Add(mBase, _consts[743]))
											*(*int32)(unsafe.Add(mBase, _consts[743])) = v316 + int32(8)
											F_enlargeStringInfo(m, v273, int32(1))
											mBase = m.M
											v323 = m.ExcPending
											if v323 != 0 {
												return
											} else {
												v325 = *(*int32)(unsafe.Add(mBase, _consts[741]))
												v326 = int32(4459332)
												v327 = *(*int32)(unsafe.Add(mBase, _consts[743]))
												*(*uint8)(unsafe.Add(mBase, uint32(v325+v327))) = uint8(v2)
												v332 = *(*int32)(unsafe.Add(mBase, _consts[743]))
												*(*int32)(unsafe.Add(mBase, _consts[743])) = v332 + int32(1)
												v338 = F_errstart(m, int32(13), int32(0))
												mBase = m.M
												v339 = m.ExcPending
												if v339 != 0 {
													return
												} else {
													if v338 != 0 {
														v341 = int64(base.Ui64(v67) >> (uint(int64(32)) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)) = uint32(v341)
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v67)
														if v2 != 0 {
															v346 = int32(703410)
														} else {
															v346 = int32(785690)
														}
														*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v346
														v349 = *(*int64)(unsafe.Add(mBase, _consts[738]))
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v349)
														v352 = *(*int64)(unsafe.Add(mBase, _consts[739]))
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v352)
														v354 = int64(32)
														v355 = int64(base.Ui64(v349) >> (uint(v354) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v355)
														v358 = int64(base.Ui64(v352) >> (uint(v354) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v358)
														F_errmsg_internal(m, int32(184919), v9)
														mBase = m.M
														v362 = m.ExcPending
														if v362 != 0 {
															return
														} else {
															F_errfinish(m, int32(515166), int32(1145), int32(19704))
															mBase = m.M
															v367 = m.ExcPending
															if v367 != 0 {
																return
															} else {
																v371 = *(*int32)(unsafe.Add(mBase, _consts[734]))
																v373 = *(*int32)(unsafe.Add(mBase, _consts[741]))
																v375 = *(*int32)(unsafe.Add(mBase, _consts[743]))
																v377 = *(*int32)(unsafe.Add(mBase, _consts[495]))
																v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
																m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
																mBase = m.M
																v380 = m.ExcPending
																if v380 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(32)
																	return
																}
															}
														}
													} else {
														v371 = *(*int32)(unsafe.Add(mBase, _consts[734]))
														v373 = *(*int32)(unsafe.Add(mBase, _consts[741]))
														v375 = *(*int32)(unsafe.Add(mBase, _consts[743]))
														v377 = *(*int32)(unsafe.Add(mBase, _consts[495]))
														v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
														m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
														mBase = m.M
														v380 = m.ExcPending
														if v380 != 0 {
															return
														} else {
															m.G0 = v9 + int32(32)
															return
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v45 = *(*int64)(unsafe.Add(mBase, _consts[739]))
					if v45 != v40 {
						v50 = v40
						*(*int64)(unsafe.Add(mBase, _consts[738])) = v36
						*(*int64)(unsafe.Add(mBase, _consts[739])) = v50
						v58 = int64(*(*int32)(unsafe.Add(mBase, _consts[735])))
						if v58 <= int64(0) {
							v64 = int64(9223372036854775807)
						} else {
							v64 = v58*int64(1000000) + v34
						}
						*(*int64)(unsafe.Add(mBase, _consts[740])) = v64
						v67 = F_GetXLogReplayRecPtr(m, int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							v69 = int32(4459328)
							v70 = *(*int32)(unsafe.Add(mBase, _consts[741]))
							v71 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v71)
							*(*int32)(unsafe.Add(mBase, _consts[742])) = v71
							*(*int32)(unsafe.Add(mBase, _consts[743])) = v71
							F_enlargeStringInfo(m, int32(4459328), int32(1))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								v81 = int32(4459328)
								v82 = *(*int32)(unsafe.Add(mBase, _consts[741]))
								v83 = int32(4459332)
								v84 = *(*int32)(unsafe.Add(mBase, _consts[743]))
								v86 = int32(114)
								*(*uint8)(unsafe.Add(mBase, uint32(v82+v84))) = uint8(v86)
								v90 = *(*int32)(unsafe.Add(mBase, _consts[743]))
								*(*int32)(unsafe.Add(mBase, _consts[743])) = v90 + int32(1)
								v95 = *(*int64)(unsafe.Add(mBase, _consts[738]))
								F_enlargeStringInfo(m, v81, int32(8))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									v100 = int32(4459328)
									v101 = *(*int32)(unsafe.Add(mBase, _consts[741]))
									v102 = int32(4459332)
									v103 = *(*int32)(unsafe.Add(mBase, _consts[743]))
									v105 = int64(56)
									v107 = int64(65280)
									v109 = int64(40)
									v112 = int64(16711680)
									v114 = int64(24)
									v116 = int64(4278190080)
									v118 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v101+v103))) = v95<<(uint(v105)%64) | v95&v107<<(uint(v109)%64) | (v95&v112<<(uint(v114)%64) | v95&v116<<(uint(v118)%64)) | (int64(base.Ui64(v95)>>(uint(v118)%64))&v116 | int64(base.Ui64(v95)>>(uint(v114)%64))&v112 | (int64(base.Ui64(v95)>>(uint(v109)%64))&v107 | int64(base.Ui64(v95)>>(uint(v105)%64))))
									v143 = *(*int32)(unsafe.Add(mBase, _consts[743]))
									v144 = int32(8)
									*(*int32)(unsafe.Add(mBase, _consts[743])) = v143 + v144
									v148 = *(*int64)(unsafe.Add(mBase, _consts[739]))
									F_enlargeStringInfo(m, v100, v144)
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return
									} else {
										v153 = int32(4459328)
										v154 = *(*int32)(unsafe.Add(mBase, _consts[741]))
										v155 = int32(4459332)
										v156 = *(*int32)(unsafe.Add(mBase, _consts[743]))
										v158 = int64(56)
										v160 = int64(65280)
										v162 = int64(40)
										v165 = int64(16711680)
										v167 = int64(24)
										v169 = int64(4278190080)
										v171 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v154+v156))) = v148<<(uint(v158)%64) | v148&v160<<(uint(v162)%64) | (v148&v165<<(uint(v167)%64) | v148&v169<<(uint(v171)%64)) | (int64(base.Ui64(v148)>>(uint(v171)%64))&v169 | int64(base.Ui64(v148)>>(uint(v167)%64))&v165 | (int64(base.Ui64(v148)>>(uint(v162)%64))&v160 | int64(base.Ui64(v148)>>(uint(v158)%64))))
										v196 = *(*int32)(unsafe.Add(mBase, _consts[743]))
										v197 = int32(8)
										*(*int32)(unsafe.Add(mBase, _consts[743])) = v196 + v197
										F_enlargeStringInfo(m, v153, v197)
										mBase = m.M
										v203 = m.ExcPending
										if v203 != 0 {
											return
										} else {
											v205 = *(*int32)(unsafe.Add(mBase, _consts[741]))
											v206 = int32(4459332)
											v207 = *(*int32)(unsafe.Add(mBase, _consts[743]))
											v209 = int64(56)
											v211 = int64(65280)
											v213 = int64(40)
											v216 = int64(16711680)
											v218 = int64(24)
											v220 = int64(4278190080)
											v222 = int64(8)
											*(*int64)(unsafe.Add(mBase, uint32(v205+v207))) = v67<<(uint(v209)%64) | v67&v211<<(uint(v213)%64) | (v67&v216<<(uint(v218)%64) | v67&v220<<(uint(v222)%64)) | (int64(base.Ui64(v67)>>(uint(v222)%64))&v220 | int64(base.Ui64(v67)>>(uint(v218)%64))&v216 | (int64(base.Ui64(v67)>>(uint(v213)%64))&v211 | int64(base.Ui64(v67)>>(uint(v209)%64))))
											v247 = *(*int32)(unsafe.Add(mBase, _consts[743]))
											*(*int32)(unsafe.Add(mBase, _consts[743])) = v247 + int32(8)
											v254 = m.G0
											v255 = int32(16)
											v256 = v254 - v255
											m.G0 = v256
											F___gettimeofday(m, v256)
											mBase = m.M
											v259 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
											v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+8)))
											m.G0 = v256 + v255
											v268 = v260 + v259*int64(1000000) - int64(946684800000000)
											F_enlargeStringInfo(m, int32(4459328), int32(8))
											mBase = m.M
											v272 = m.ExcPending
											if v272 != 0 {
												return
											} else {
												v273 = int32(4459328)
												v274 = *(*int32)(unsafe.Add(mBase, _consts[741]))
												v275 = int32(4459332)
												v276 = *(*int32)(unsafe.Add(mBase, _consts[743]))
												v278 = int64(56)
												v280 = int64(65280)
												v282 = int64(40)
												v285 = int64(16711680)
												v287 = int64(24)
												v289 = int64(4278190080)
												v291 = int64(8)
												*(*int64)(unsafe.Add(mBase, uint32(v274+v276))) = v268<<(uint(v278)%64) | v268&v280<<(uint(v282)%64) | (v268&v285<<(uint(v287)%64) | v268&v289<<(uint(v291)%64)) | (int64(base.Ui64(v268)>>(uint(v291)%64))&v289 | int64(base.Ui64(v268)>>(uint(v287)%64))&v285 | (int64(base.Ui64(v268)>>(uint(v282)%64))&v280 | int64(base.Ui64(v268)>>(uint(v278)%64))))
												v316 = *(*int32)(unsafe.Add(mBase, _consts[743]))
												*(*int32)(unsafe.Add(mBase, _consts[743])) = v316 + int32(8)
												F_enlargeStringInfo(m, v273, int32(1))
												mBase = m.M
												v323 = m.ExcPending
												if v323 != 0 {
													return
												} else {
													v325 = *(*int32)(unsafe.Add(mBase, _consts[741]))
													v326 = int32(4459332)
													v327 = *(*int32)(unsafe.Add(mBase, _consts[743]))
													*(*uint8)(unsafe.Add(mBase, uint32(v325+v327))) = uint8(v2)
													v332 = *(*int32)(unsafe.Add(mBase, _consts[743]))
													*(*int32)(unsafe.Add(mBase, _consts[743])) = v332 + int32(1)
													v338 = F_errstart(m, int32(13), int32(0))
													mBase = m.M
													v339 = m.ExcPending
													if v339 != 0 {
														return
													} else {
														if v338 != 0 {
															v341 = int64(base.Ui64(v67) >> (uint(int64(32)) % 64))
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)) = uint32(v341)
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v67)
															if v2 != 0 {
																v346 = int32(703410)
															} else {
																v346 = int32(785690)
															}
															*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v346
															v349 = *(*int64)(unsafe.Add(mBase, _consts[738]))
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v349)
															v352 = *(*int64)(unsafe.Add(mBase, _consts[739]))
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v352)
															v354 = int64(32)
															v355 = int64(base.Ui64(v349) >> (uint(v354) % 64))
															*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v355)
															v358 = int64(base.Ui64(v352) >> (uint(v354) % 64))
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v358)
															F_errmsg_internal(m, int32(184919), v9)
															mBase = m.M
															v362 = m.ExcPending
															if v362 != 0 {
																return
															} else {
																F_errfinish(m, int32(515166), int32(1145), int32(19704))
																mBase = m.M
																v367 = m.ExcPending
																if v367 != 0 {
																	return
																} else {
																	v371 = *(*int32)(unsafe.Add(mBase, _consts[734]))
																	v373 = *(*int32)(unsafe.Add(mBase, _consts[741]))
																	v375 = *(*int32)(unsafe.Add(mBase, _consts[743]))
																	v377 = *(*int32)(unsafe.Add(mBase, _consts[495]))
																	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
																	m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
																	mBase = m.M
																	v380 = m.ExcPending
																	if v380 != 0 {
																		return
																	} else {
																		m.G0 = v9 + int32(32)
																		return
																	}
																}
															}
														} else {
															v371 = *(*int32)(unsafe.Add(mBase, _consts[734]))
															v373 = *(*int32)(unsafe.Add(mBase, _consts[741]))
															v375 = *(*int32)(unsafe.Add(mBase, _consts[743]))
															v377 = *(*int32)(unsafe.Add(mBase, _consts[495]))
															v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
															m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
															mBase = m.M
															v380 = m.ExcPending
															if v380 != 0 {
																return
															} else {
																m.G0 = v9 + int32(32)
																return
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v48 = *(*int64)(unsafe.Add(mBase, _consts[740]))
						if v34 < v48 {
							m.G0 = v9 + int32(32)
							return
						} else {
							v50 = v40
							*(*int64)(unsafe.Add(mBase, _consts[738])) = v36
							*(*int64)(unsafe.Add(mBase, _consts[739])) = v50
							v58 = int64(*(*int32)(unsafe.Add(mBase, _consts[735])))
							if v58 <= int64(0) {
								v64 = int64(9223372036854775807)
							} else {
								v64 = v58*int64(1000000) + v34
							}
							*(*int64)(unsafe.Add(mBase, _consts[740])) = v64
							v67 = F_GetXLogReplayRecPtr(m, int32(0))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								v69 = int32(4459328)
								v70 = *(*int32)(unsafe.Add(mBase, _consts[741]))
								v71 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v71)
								*(*int32)(unsafe.Add(mBase, _consts[742])) = v71
								*(*int32)(unsafe.Add(mBase, _consts[743])) = v71
								F_enlargeStringInfo(m, int32(4459328), int32(1))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									v81 = int32(4459328)
									v82 = *(*int32)(unsafe.Add(mBase, _consts[741]))
									v83 = int32(4459332)
									v84 = *(*int32)(unsafe.Add(mBase, _consts[743]))
									v86 = int32(114)
									*(*uint8)(unsafe.Add(mBase, uint32(v82+v84))) = uint8(v86)
									v90 = *(*int32)(unsafe.Add(mBase, _consts[743]))
									*(*int32)(unsafe.Add(mBase, _consts[743])) = v90 + int32(1)
									v95 = *(*int64)(unsafe.Add(mBase, _consts[738]))
									F_enlargeStringInfo(m, v81, int32(8))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										v100 = int32(4459328)
										v101 = *(*int32)(unsafe.Add(mBase, _consts[741]))
										v102 = int32(4459332)
										v103 = *(*int32)(unsafe.Add(mBase, _consts[743]))
										v105 = int64(56)
										v107 = int64(65280)
										v109 = int64(40)
										v112 = int64(16711680)
										v114 = int64(24)
										v116 = int64(4278190080)
										v118 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v101+v103))) = v95<<(uint(v105)%64) | v95&v107<<(uint(v109)%64) | (v95&v112<<(uint(v114)%64) | v95&v116<<(uint(v118)%64)) | (int64(base.Ui64(v95)>>(uint(v118)%64))&v116 | int64(base.Ui64(v95)>>(uint(v114)%64))&v112 | (int64(base.Ui64(v95)>>(uint(v109)%64))&v107 | int64(base.Ui64(v95)>>(uint(v105)%64))))
										v143 = *(*int32)(unsafe.Add(mBase, _consts[743]))
										v144 = int32(8)
										*(*int32)(unsafe.Add(mBase, _consts[743])) = v143 + v144
										v148 = *(*int64)(unsafe.Add(mBase, _consts[739]))
										F_enlargeStringInfo(m, v100, v144)
										mBase = m.M
										v152 = m.ExcPending
										if v152 != 0 {
											return
										} else {
											v153 = int32(4459328)
											v154 = *(*int32)(unsafe.Add(mBase, _consts[741]))
											v155 = int32(4459332)
											v156 = *(*int32)(unsafe.Add(mBase, _consts[743]))
											v158 = int64(56)
											v160 = int64(65280)
											v162 = int64(40)
											v165 = int64(16711680)
											v167 = int64(24)
											v169 = int64(4278190080)
											v171 = int64(8)
											*(*int64)(unsafe.Add(mBase, uint32(v154+v156))) = v148<<(uint(v158)%64) | v148&v160<<(uint(v162)%64) | (v148&v165<<(uint(v167)%64) | v148&v169<<(uint(v171)%64)) | (int64(base.Ui64(v148)>>(uint(v171)%64))&v169 | int64(base.Ui64(v148)>>(uint(v167)%64))&v165 | (int64(base.Ui64(v148)>>(uint(v162)%64))&v160 | int64(base.Ui64(v148)>>(uint(v158)%64))))
											v196 = *(*int32)(unsafe.Add(mBase, _consts[743]))
											v197 = int32(8)
											*(*int32)(unsafe.Add(mBase, _consts[743])) = v196 + v197
											F_enlargeStringInfo(m, v153, v197)
											mBase = m.M
											v203 = m.ExcPending
											if v203 != 0 {
												return
											} else {
												v205 = *(*int32)(unsafe.Add(mBase, _consts[741]))
												v206 = int32(4459332)
												v207 = *(*int32)(unsafe.Add(mBase, _consts[743]))
												v209 = int64(56)
												v211 = int64(65280)
												v213 = int64(40)
												v216 = int64(16711680)
												v218 = int64(24)
												v220 = int64(4278190080)
												v222 = int64(8)
												*(*int64)(unsafe.Add(mBase, uint32(v205+v207))) = v67<<(uint(v209)%64) | v67&v211<<(uint(v213)%64) | (v67&v216<<(uint(v218)%64) | v67&v220<<(uint(v222)%64)) | (int64(base.Ui64(v67)>>(uint(v222)%64))&v220 | int64(base.Ui64(v67)>>(uint(v218)%64))&v216 | (int64(base.Ui64(v67)>>(uint(v213)%64))&v211 | int64(base.Ui64(v67)>>(uint(v209)%64))))
												v247 = *(*int32)(unsafe.Add(mBase, _consts[743]))
												*(*int32)(unsafe.Add(mBase, _consts[743])) = v247 + int32(8)
												v254 = m.G0
												v255 = int32(16)
												v256 = v254 - v255
												m.G0 = v256
												F___gettimeofday(m, v256)
												mBase = m.M
												v259 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
												v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+8)))
												m.G0 = v256 + v255
												v268 = v260 + v259*int64(1000000) - int64(946684800000000)
												F_enlargeStringInfo(m, int32(4459328), int32(8))
												mBase = m.M
												v272 = m.ExcPending
												if v272 != 0 {
													return
												} else {
													v273 = int32(4459328)
													v274 = *(*int32)(unsafe.Add(mBase, _consts[741]))
													v275 = int32(4459332)
													v276 = *(*int32)(unsafe.Add(mBase, _consts[743]))
													v278 = int64(56)
													v280 = int64(65280)
													v282 = int64(40)
													v285 = int64(16711680)
													v287 = int64(24)
													v289 = int64(4278190080)
													v291 = int64(8)
													*(*int64)(unsafe.Add(mBase, uint32(v274+v276))) = v268<<(uint(v278)%64) | v268&v280<<(uint(v282)%64) | (v268&v285<<(uint(v287)%64) | v268&v289<<(uint(v291)%64)) | (int64(base.Ui64(v268)>>(uint(v291)%64))&v289 | int64(base.Ui64(v268)>>(uint(v287)%64))&v285 | (int64(base.Ui64(v268)>>(uint(v282)%64))&v280 | int64(base.Ui64(v268)>>(uint(v278)%64))))
													v316 = *(*int32)(unsafe.Add(mBase, _consts[743]))
													*(*int32)(unsafe.Add(mBase, _consts[743])) = v316 + int32(8)
													F_enlargeStringInfo(m, v273, int32(1))
													mBase = m.M
													v323 = m.ExcPending
													if v323 != 0 {
														return
													} else {
														v325 = *(*int32)(unsafe.Add(mBase, _consts[741]))
														v326 = int32(4459332)
														v327 = *(*int32)(unsafe.Add(mBase, _consts[743]))
														*(*uint8)(unsafe.Add(mBase, uint32(v325+v327))) = uint8(v2)
														v332 = *(*int32)(unsafe.Add(mBase, _consts[743]))
														*(*int32)(unsafe.Add(mBase, _consts[743])) = v332 + int32(1)
														v338 = F_errstart(m, int32(13), int32(0))
														mBase = m.M
														v339 = m.ExcPending
														if v339 != 0 {
															return
														} else {
															if v338 != 0 {
																v341 = int64(base.Ui64(v67) >> (uint(int64(32)) % 64))
																*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)) = uint32(v341)
																*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v67)
																if v2 != 0 {
																	v346 = int32(703410)
																} else {
																	v346 = int32(785690)
																}
																*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v346
																v349 = *(*int64)(unsafe.Add(mBase, _consts[738]))
																*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v349)
																v352 = *(*int64)(unsafe.Add(mBase, _consts[739]))
																*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v352)
																v354 = int64(32)
																v355 = int64(base.Ui64(v349) >> (uint(v354) % 64))
																*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v355)
																v358 = int64(base.Ui64(v352) >> (uint(v354) % 64))
																*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v358)
																F_errmsg_internal(m, int32(184919), v9)
																mBase = m.M
																v362 = m.ExcPending
																if v362 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(515166), int32(1145), int32(19704))
																	mBase = m.M
																	v367 = m.ExcPending
																	if v367 != 0 {
																		return
																	} else {
																		v371 = *(*int32)(unsafe.Add(mBase, _consts[734]))
																		v373 = *(*int32)(unsafe.Add(mBase, _consts[741]))
																		v375 = *(*int32)(unsafe.Add(mBase, _consts[743]))
																		v377 = *(*int32)(unsafe.Add(mBase, _consts[495]))
																		v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
																		m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
																		mBase = m.M
																		v380 = m.ExcPending
																		if v380 != 0 {
																			return
																		} else {
																			m.G0 = v9 + int32(32)
																			return
																		}
																	}
																}
															} else {
																v371 = *(*int32)(unsafe.Add(mBase, _consts[734]))
																v373 = *(*int32)(unsafe.Add(mBase, _consts[741]))
																v375 = *(*int32)(unsafe.Add(mBase, _consts[743]))
																v377 = *(*int32)(unsafe.Add(mBase, _consts[495]))
																v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
																m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
																mBase = m.M
																v380 = m.ExcPending
																if v380 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(32)
																	return
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v20 = m.G0
		v21 = int32(16)
		v22 = v20 - v21
		m.G0 = v22
		F___gettimeofday(m, v22)
		mBase = m.M
		v25 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
		v26 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22)+8)))
		m.G0 = v22 + v21
		v34 = v26 + v25*int64(1000000) - int64(946684800000000)
		v36 = *(*int64)(unsafe.Add(mBase, _consts[736]))
		if l0 != 0 {
			v38 = *(*int64)(unsafe.Add(mBase, _consts[737]))
			v50 = v38
			*(*int64)(unsafe.Add(mBase, _consts[738])) = v36
			*(*int64)(unsafe.Add(mBase, _consts[739])) = v50
			v58 = int64(*(*int32)(unsafe.Add(mBase, _consts[735])))
			if v58 <= int64(0) {
				v64 = int64(9223372036854775807)
			} else {
				v64 = v58*int64(1000000) + v34
			}
			*(*int64)(unsafe.Add(mBase, _consts[740])) = v64
			v67 = F_GetXLogReplayRecPtr(m, int32(0))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return
			} else {
				v69 = int32(4459328)
				v70 = *(*int32)(unsafe.Add(mBase, _consts[741]))
				v71 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v71)
				*(*int32)(unsafe.Add(mBase, _consts[742])) = v71
				*(*int32)(unsafe.Add(mBase, _consts[743])) = v71
				F_enlargeStringInfo(m, int32(4459328), int32(1))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return
				} else {
					v81 = int32(4459328)
					v82 = *(*int32)(unsafe.Add(mBase, _consts[741]))
					v83 = int32(4459332)
					v84 = *(*int32)(unsafe.Add(mBase, _consts[743]))
					v86 = int32(114)
					*(*uint8)(unsafe.Add(mBase, uint32(v82+v84))) = uint8(v86)
					v90 = *(*int32)(unsafe.Add(mBase, _consts[743]))
					*(*int32)(unsafe.Add(mBase, _consts[743])) = v90 + int32(1)
					v95 = *(*int64)(unsafe.Add(mBase, _consts[738]))
					F_enlargeStringInfo(m, v81, int32(8))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return
					} else {
						v100 = int32(4459328)
						v101 = *(*int32)(unsafe.Add(mBase, _consts[741]))
						v102 = int32(4459332)
						v103 = *(*int32)(unsafe.Add(mBase, _consts[743]))
						v105 = int64(56)
						v107 = int64(65280)
						v109 = int64(40)
						v112 = int64(16711680)
						v114 = int64(24)
						v116 = int64(4278190080)
						v118 = int64(8)
						*(*int64)(unsafe.Add(mBase, uint32(v101+v103))) = v95<<(uint(v105)%64) | v95&v107<<(uint(v109)%64) | (v95&v112<<(uint(v114)%64) | v95&v116<<(uint(v118)%64)) | (int64(base.Ui64(v95)>>(uint(v118)%64))&v116 | int64(base.Ui64(v95)>>(uint(v114)%64))&v112 | (int64(base.Ui64(v95)>>(uint(v109)%64))&v107 | int64(base.Ui64(v95)>>(uint(v105)%64))))
						v143 = *(*int32)(unsafe.Add(mBase, _consts[743]))
						v144 = int32(8)
						*(*int32)(unsafe.Add(mBase, _consts[743])) = v143 + v144
						v148 = *(*int64)(unsafe.Add(mBase, _consts[739]))
						F_enlargeStringInfo(m, v100, v144)
						mBase = m.M
						v152 = m.ExcPending
						if v152 != 0 {
							return
						} else {
							v153 = int32(4459328)
							v154 = *(*int32)(unsafe.Add(mBase, _consts[741]))
							v155 = int32(4459332)
							v156 = *(*int32)(unsafe.Add(mBase, _consts[743]))
							v158 = int64(56)
							v160 = int64(65280)
							v162 = int64(40)
							v165 = int64(16711680)
							v167 = int64(24)
							v169 = int64(4278190080)
							v171 = int64(8)
							*(*int64)(unsafe.Add(mBase, uint32(v154+v156))) = v148<<(uint(v158)%64) | v148&v160<<(uint(v162)%64) | (v148&v165<<(uint(v167)%64) | v148&v169<<(uint(v171)%64)) | (int64(base.Ui64(v148)>>(uint(v171)%64))&v169 | int64(base.Ui64(v148)>>(uint(v167)%64))&v165 | (int64(base.Ui64(v148)>>(uint(v162)%64))&v160 | int64(base.Ui64(v148)>>(uint(v158)%64))))
							v196 = *(*int32)(unsafe.Add(mBase, _consts[743]))
							v197 = int32(8)
							*(*int32)(unsafe.Add(mBase, _consts[743])) = v196 + v197
							F_enlargeStringInfo(m, v153, v197)
							mBase = m.M
							v203 = m.ExcPending
							if v203 != 0 {
								return
							} else {
								v205 = *(*int32)(unsafe.Add(mBase, _consts[741]))
								v206 = int32(4459332)
								v207 = *(*int32)(unsafe.Add(mBase, _consts[743]))
								v209 = int64(56)
								v211 = int64(65280)
								v213 = int64(40)
								v216 = int64(16711680)
								v218 = int64(24)
								v220 = int64(4278190080)
								v222 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v205+v207))) = v67<<(uint(v209)%64) | v67&v211<<(uint(v213)%64) | (v67&v216<<(uint(v218)%64) | v67&v220<<(uint(v222)%64)) | (int64(base.Ui64(v67)>>(uint(v222)%64))&v220 | int64(base.Ui64(v67)>>(uint(v218)%64))&v216 | (int64(base.Ui64(v67)>>(uint(v213)%64))&v211 | int64(base.Ui64(v67)>>(uint(v209)%64))))
								v247 = *(*int32)(unsafe.Add(mBase, _consts[743]))
								*(*int32)(unsafe.Add(mBase, _consts[743])) = v247 + int32(8)
								v254 = m.G0
								v255 = int32(16)
								v256 = v254 - v255
								m.G0 = v256
								F___gettimeofday(m, v256)
								mBase = m.M
								v259 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
								v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+8)))
								m.G0 = v256 + v255
								v268 = v260 + v259*int64(1000000) - int64(946684800000000)
								F_enlargeStringInfo(m, int32(4459328), int32(8))
								mBase = m.M
								v272 = m.ExcPending
								if v272 != 0 {
									return
								} else {
									v273 = int32(4459328)
									v274 = *(*int32)(unsafe.Add(mBase, _consts[741]))
									v275 = int32(4459332)
									v276 = *(*int32)(unsafe.Add(mBase, _consts[743]))
									v278 = int64(56)
									v280 = int64(65280)
									v282 = int64(40)
									v285 = int64(16711680)
									v287 = int64(24)
									v289 = int64(4278190080)
									v291 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v274+v276))) = v268<<(uint(v278)%64) | v268&v280<<(uint(v282)%64) | (v268&v285<<(uint(v287)%64) | v268&v289<<(uint(v291)%64)) | (int64(base.Ui64(v268)>>(uint(v291)%64))&v289 | int64(base.Ui64(v268)>>(uint(v287)%64))&v285 | (int64(base.Ui64(v268)>>(uint(v282)%64))&v280 | int64(base.Ui64(v268)>>(uint(v278)%64))))
									v316 = *(*int32)(unsafe.Add(mBase, _consts[743]))
									*(*int32)(unsafe.Add(mBase, _consts[743])) = v316 + int32(8)
									F_enlargeStringInfo(m, v273, int32(1))
									mBase = m.M
									v323 = m.ExcPending
									if v323 != 0 {
										return
									} else {
										v325 = *(*int32)(unsafe.Add(mBase, _consts[741]))
										v326 = int32(4459332)
										v327 = *(*int32)(unsafe.Add(mBase, _consts[743]))
										*(*uint8)(unsafe.Add(mBase, uint32(v325+v327))) = uint8(v2)
										v332 = *(*int32)(unsafe.Add(mBase, _consts[743]))
										*(*int32)(unsafe.Add(mBase, _consts[743])) = v332 + int32(1)
										v338 = F_errstart(m, int32(13), int32(0))
										mBase = m.M
										v339 = m.ExcPending
										if v339 != 0 {
											return
										} else {
											if v338 != 0 {
												v341 = int64(base.Ui64(v67) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)) = uint32(v341)
												*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v67)
												if v2 != 0 {
													v346 = int32(703410)
												} else {
													v346 = int32(785690)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v346
												v349 = *(*int64)(unsafe.Add(mBase, _consts[738]))
												*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v349)
												v352 = *(*int64)(unsafe.Add(mBase, _consts[739]))
												*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v352)
												v354 = int64(32)
												v355 = int64(base.Ui64(v349) >> (uint(v354) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v355)
												v358 = int64(base.Ui64(v352) >> (uint(v354) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v358)
												F_errmsg_internal(m, int32(184919), v9)
												mBase = m.M
												v362 = m.ExcPending
												if v362 != 0 {
													return
												} else {
													F_errfinish(m, int32(515166), int32(1145), int32(19704))
													mBase = m.M
													v367 = m.ExcPending
													if v367 != 0 {
														return
													} else {
														v371 = *(*int32)(unsafe.Add(mBase, _consts[734]))
														v373 = *(*int32)(unsafe.Add(mBase, _consts[741]))
														v375 = *(*int32)(unsafe.Add(mBase, _consts[743]))
														v377 = *(*int32)(unsafe.Add(mBase, _consts[495]))
														v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
														m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
														mBase = m.M
														v380 = m.ExcPending
														if v380 != 0 {
															return
														} else {
															m.G0 = v9 + int32(32)
															return
														}
													}
												}
											} else {
												v371 = *(*int32)(unsafe.Add(mBase, _consts[734]))
												v373 = *(*int32)(unsafe.Add(mBase, _consts[741]))
												v375 = *(*int32)(unsafe.Add(mBase, _consts[743]))
												v377 = *(*int32)(unsafe.Add(mBase, _consts[495]))
												v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
												m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
												mBase = m.M
												v380 = m.ExcPending
												if v380 != 0 {
													return
												} else {
													m.G0 = v9 + int32(32)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v40 = *(*int64)(unsafe.Add(mBase, _consts[737]))
			v42 = *(*int64)(unsafe.Add(mBase, _consts[738]))
			if v42 != v36 {
				v50 = v40
				*(*int64)(unsafe.Add(mBase, _consts[738])) = v36
				*(*int64)(unsafe.Add(mBase, _consts[739])) = v50
				v58 = int64(*(*int32)(unsafe.Add(mBase, _consts[735])))
				if v58 <= int64(0) {
					v64 = int64(9223372036854775807)
				} else {
					v64 = v58*int64(1000000) + v34
				}
				*(*int64)(unsafe.Add(mBase, _consts[740])) = v64
				v67 = F_GetXLogReplayRecPtr(m, int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					v69 = int32(4459328)
					v70 = *(*int32)(unsafe.Add(mBase, _consts[741]))
					v71 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v71)
					*(*int32)(unsafe.Add(mBase, _consts[742])) = v71
					*(*int32)(unsafe.Add(mBase, _consts[743])) = v71
					F_enlargeStringInfo(m, int32(4459328), int32(1))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						v81 = int32(4459328)
						v82 = *(*int32)(unsafe.Add(mBase, _consts[741]))
						v83 = int32(4459332)
						v84 = *(*int32)(unsafe.Add(mBase, _consts[743]))
						v86 = int32(114)
						*(*uint8)(unsafe.Add(mBase, uint32(v82+v84))) = uint8(v86)
						v90 = *(*int32)(unsafe.Add(mBase, _consts[743]))
						*(*int32)(unsafe.Add(mBase, _consts[743])) = v90 + int32(1)
						v95 = *(*int64)(unsafe.Add(mBase, _consts[738]))
						F_enlargeStringInfo(m, v81, int32(8))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							v100 = int32(4459328)
							v101 = *(*int32)(unsafe.Add(mBase, _consts[741]))
							v102 = int32(4459332)
							v103 = *(*int32)(unsafe.Add(mBase, _consts[743]))
							v105 = int64(56)
							v107 = int64(65280)
							v109 = int64(40)
							v112 = int64(16711680)
							v114 = int64(24)
							v116 = int64(4278190080)
							v118 = int64(8)
							*(*int64)(unsafe.Add(mBase, uint32(v101+v103))) = v95<<(uint(v105)%64) | v95&v107<<(uint(v109)%64) | (v95&v112<<(uint(v114)%64) | v95&v116<<(uint(v118)%64)) | (int64(base.Ui64(v95)>>(uint(v118)%64))&v116 | int64(base.Ui64(v95)>>(uint(v114)%64))&v112 | (int64(base.Ui64(v95)>>(uint(v109)%64))&v107 | int64(base.Ui64(v95)>>(uint(v105)%64))))
							v143 = *(*int32)(unsafe.Add(mBase, _consts[743]))
							v144 = int32(8)
							*(*int32)(unsafe.Add(mBase, _consts[743])) = v143 + v144
							v148 = *(*int64)(unsafe.Add(mBase, _consts[739]))
							F_enlargeStringInfo(m, v100, v144)
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return
							} else {
								v153 = int32(4459328)
								v154 = *(*int32)(unsafe.Add(mBase, _consts[741]))
								v155 = int32(4459332)
								v156 = *(*int32)(unsafe.Add(mBase, _consts[743]))
								v158 = int64(56)
								v160 = int64(65280)
								v162 = int64(40)
								v165 = int64(16711680)
								v167 = int64(24)
								v169 = int64(4278190080)
								v171 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v154+v156))) = v148<<(uint(v158)%64) | v148&v160<<(uint(v162)%64) | (v148&v165<<(uint(v167)%64) | v148&v169<<(uint(v171)%64)) | (int64(base.Ui64(v148)>>(uint(v171)%64))&v169 | int64(base.Ui64(v148)>>(uint(v167)%64))&v165 | (int64(base.Ui64(v148)>>(uint(v162)%64))&v160 | int64(base.Ui64(v148)>>(uint(v158)%64))))
								v196 = *(*int32)(unsafe.Add(mBase, _consts[743]))
								v197 = int32(8)
								*(*int32)(unsafe.Add(mBase, _consts[743])) = v196 + v197
								F_enlargeStringInfo(m, v153, v197)
								mBase = m.M
								v203 = m.ExcPending
								if v203 != 0 {
									return
								} else {
									v205 = *(*int32)(unsafe.Add(mBase, _consts[741]))
									v206 = int32(4459332)
									v207 = *(*int32)(unsafe.Add(mBase, _consts[743]))
									v209 = int64(56)
									v211 = int64(65280)
									v213 = int64(40)
									v216 = int64(16711680)
									v218 = int64(24)
									v220 = int64(4278190080)
									v222 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v205+v207))) = v67<<(uint(v209)%64) | v67&v211<<(uint(v213)%64) | (v67&v216<<(uint(v218)%64) | v67&v220<<(uint(v222)%64)) | (int64(base.Ui64(v67)>>(uint(v222)%64))&v220 | int64(base.Ui64(v67)>>(uint(v218)%64))&v216 | (int64(base.Ui64(v67)>>(uint(v213)%64))&v211 | int64(base.Ui64(v67)>>(uint(v209)%64))))
									v247 = *(*int32)(unsafe.Add(mBase, _consts[743]))
									*(*int32)(unsafe.Add(mBase, _consts[743])) = v247 + int32(8)
									v254 = m.G0
									v255 = int32(16)
									v256 = v254 - v255
									m.G0 = v256
									F___gettimeofday(m, v256)
									mBase = m.M
									v259 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
									v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+8)))
									m.G0 = v256 + v255
									v268 = v260 + v259*int64(1000000) - int64(946684800000000)
									F_enlargeStringInfo(m, int32(4459328), int32(8))
									mBase = m.M
									v272 = m.ExcPending
									if v272 != 0 {
										return
									} else {
										v273 = int32(4459328)
										v274 = *(*int32)(unsafe.Add(mBase, _consts[741]))
										v275 = int32(4459332)
										v276 = *(*int32)(unsafe.Add(mBase, _consts[743]))
										v278 = int64(56)
										v280 = int64(65280)
										v282 = int64(40)
										v285 = int64(16711680)
										v287 = int64(24)
										v289 = int64(4278190080)
										v291 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v274+v276))) = v268<<(uint(v278)%64) | v268&v280<<(uint(v282)%64) | (v268&v285<<(uint(v287)%64) | v268&v289<<(uint(v291)%64)) | (int64(base.Ui64(v268)>>(uint(v291)%64))&v289 | int64(base.Ui64(v268)>>(uint(v287)%64))&v285 | (int64(base.Ui64(v268)>>(uint(v282)%64))&v280 | int64(base.Ui64(v268)>>(uint(v278)%64))))
										v316 = *(*int32)(unsafe.Add(mBase, _consts[743]))
										*(*int32)(unsafe.Add(mBase, _consts[743])) = v316 + int32(8)
										F_enlargeStringInfo(m, v273, int32(1))
										mBase = m.M
										v323 = m.ExcPending
										if v323 != 0 {
											return
										} else {
											v325 = *(*int32)(unsafe.Add(mBase, _consts[741]))
											v326 = int32(4459332)
											v327 = *(*int32)(unsafe.Add(mBase, _consts[743]))
											*(*uint8)(unsafe.Add(mBase, uint32(v325+v327))) = uint8(v2)
											v332 = *(*int32)(unsafe.Add(mBase, _consts[743]))
											*(*int32)(unsafe.Add(mBase, _consts[743])) = v332 + int32(1)
											v338 = F_errstart(m, int32(13), int32(0))
											mBase = m.M
											v339 = m.ExcPending
											if v339 != 0 {
												return
											} else {
												if v338 != 0 {
													v341 = int64(base.Ui64(v67) >> (uint(int64(32)) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)) = uint32(v341)
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v67)
													if v2 != 0 {
														v346 = int32(703410)
													} else {
														v346 = int32(785690)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v346
													v349 = *(*int64)(unsafe.Add(mBase, _consts[738]))
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v349)
													v352 = *(*int64)(unsafe.Add(mBase, _consts[739]))
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v352)
													v354 = int64(32)
													v355 = int64(base.Ui64(v349) >> (uint(v354) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v355)
													v358 = int64(base.Ui64(v352) >> (uint(v354) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v358)
													F_errmsg_internal(m, int32(184919), v9)
													mBase = m.M
													v362 = m.ExcPending
													if v362 != 0 {
														return
													} else {
														F_errfinish(m, int32(515166), int32(1145), int32(19704))
														mBase = m.M
														v367 = m.ExcPending
														if v367 != 0 {
															return
														} else {
															v371 = *(*int32)(unsafe.Add(mBase, _consts[734]))
															v373 = *(*int32)(unsafe.Add(mBase, _consts[741]))
															v375 = *(*int32)(unsafe.Add(mBase, _consts[743]))
															v377 = *(*int32)(unsafe.Add(mBase, _consts[495]))
															v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
															m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
															mBase = m.M
															v380 = m.ExcPending
															if v380 != 0 {
																return
															} else {
																m.G0 = v9 + int32(32)
																return
															}
														}
													}
												} else {
													v371 = *(*int32)(unsafe.Add(mBase, _consts[734]))
													v373 = *(*int32)(unsafe.Add(mBase, _consts[741]))
													v375 = *(*int32)(unsafe.Add(mBase, _consts[743]))
													v377 = *(*int32)(unsafe.Add(mBase, _consts[495]))
													v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
													m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
													mBase = m.M
													v380 = m.ExcPending
													if v380 != 0 {
														return
													} else {
														m.G0 = v9 + int32(32)
														return
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v45 = *(*int64)(unsafe.Add(mBase, _consts[739]))
				if v45 != v40 {
					v50 = v40
					*(*int64)(unsafe.Add(mBase, _consts[738])) = v36
					*(*int64)(unsafe.Add(mBase, _consts[739])) = v50
					v58 = int64(*(*int32)(unsafe.Add(mBase, _consts[735])))
					if v58 <= int64(0) {
						v64 = int64(9223372036854775807)
					} else {
						v64 = v58*int64(1000000) + v34
					}
					*(*int64)(unsafe.Add(mBase, _consts[740])) = v64
					v67 = F_GetXLogReplayRecPtr(m, int32(0))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						v69 = int32(4459328)
						v70 = *(*int32)(unsafe.Add(mBase, _consts[741]))
						v71 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v71)
						*(*int32)(unsafe.Add(mBase, _consts[742])) = v71
						*(*int32)(unsafe.Add(mBase, _consts[743])) = v71
						F_enlargeStringInfo(m, int32(4459328), int32(1))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							v81 = int32(4459328)
							v82 = *(*int32)(unsafe.Add(mBase, _consts[741]))
							v83 = int32(4459332)
							v84 = *(*int32)(unsafe.Add(mBase, _consts[743]))
							v86 = int32(114)
							*(*uint8)(unsafe.Add(mBase, uint32(v82+v84))) = uint8(v86)
							v90 = *(*int32)(unsafe.Add(mBase, _consts[743]))
							*(*int32)(unsafe.Add(mBase, _consts[743])) = v90 + int32(1)
							v95 = *(*int64)(unsafe.Add(mBase, _consts[738]))
							F_enlargeStringInfo(m, v81, int32(8))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								v100 = int32(4459328)
								v101 = *(*int32)(unsafe.Add(mBase, _consts[741]))
								v102 = int32(4459332)
								v103 = *(*int32)(unsafe.Add(mBase, _consts[743]))
								v105 = int64(56)
								v107 = int64(65280)
								v109 = int64(40)
								v112 = int64(16711680)
								v114 = int64(24)
								v116 = int64(4278190080)
								v118 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v101+v103))) = v95<<(uint(v105)%64) | v95&v107<<(uint(v109)%64) | (v95&v112<<(uint(v114)%64) | v95&v116<<(uint(v118)%64)) | (int64(base.Ui64(v95)>>(uint(v118)%64))&v116 | int64(base.Ui64(v95)>>(uint(v114)%64))&v112 | (int64(base.Ui64(v95)>>(uint(v109)%64))&v107 | int64(base.Ui64(v95)>>(uint(v105)%64))))
								v143 = *(*int32)(unsafe.Add(mBase, _consts[743]))
								v144 = int32(8)
								*(*int32)(unsafe.Add(mBase, _consts[743])) = v143 + v144
								v148 = *(*int64)(unsafe.Add(mBase, _consts[739]))
								F_enlargeStringInfo(m, v100, v144)
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return
								} else {
									v153 = int32(4459328)
									v154 = *(*int32)(unsafe.Add(mBase, _consts[741]))
									v155 = int32(4459332)
									v156 = *(*int32)(unsafe.Add(mBase, _consts[743]))
									v158 = int64(56)
									v160 = int64(65280)
									v162 = int64(40)
									v165 = int64(16711680)
									v167 = int64(24)
									v169 = int64(4278190080)
									v171 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v154+v156))) = v148<<(uint(v158)%64) | v148&v160<<(uint(v162)%64) | (v148&v165<<(uint(v167)%64) | v148&v169<<(uint(v171)%64)) | (int64(base.Ui64(v148)>>(uint(v171)%64))&v169 | int64(base.Ui64(v148)>>(uint(v167)%64))&v165 | (int64(base.Ui64(v148)>>(uint(v162)%64))&v160 | int64(base.Ui64(v148)>>(uint(v158)%64))))
									v196 = *(*int32)(unsafe.Add(mBase, _consts[743]))
									v197 = int32(8)
									*(*int32)(unsafe.Add(mBase, _consts[743])) = v196 + v197
									F_enlargeStringInfo(m, v153, v197)
									mBase = m.M
									v203 = m.ExcPending
									if v203 != 0 {
										return
									} else {
										v205 = *(*int32)(unsafe.Add(mBase, _consts[741]))
										v206 = int32(4459332)
										v207 = *(*int32)(unsafe.Add(mBase, _consts[743]))
										v209 = int64(56)
										v211 = int64(65280)
										v213 = int64(40)
										v216 = int64(16711680)
										v218 = int64(24)
										v220 = int64(4278190080)
										v222 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v205+v207))) = v67<<(uint(v209)%64) | v67&v211<<(uint(v213)%64) | (v67&v216<<(uint(v218)%64) | v67&v220<<(uint(v222)%64)) | (int64(base.Ui64(v67)>>(uint(v222)%64))&v220 | int64(base.Ui64(v67)>>(uint(v218)%64))&v216 | (int64(base.Ui64(v67)>>(uint(v213)%64))&v211 | int64(base.Ui64(v67)>>(uint(v209)%64))))
										v247 = *(*int32)(unsafe.Add(mBase, _consts[743]))
										*(*int32)(unsafe.Add(mBase, _consts[743])) = v247 + int32(8)
										v254 = m.G0
										v255 = int32(16)
										v256 = v254 - v255
										m.G0 = v256
										F___gettimeofday(m, v256)
										mBase = m.M
										v259 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
										v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+8)))
										m.G0 = v256 + v255
										v268 = v260 + v259*int64(1000000) - int64(946684800000000)
										F_enlargeStringInfo(m, int32(4459328), int32(8))
										mBase = m.M
										v272 = m.ExcPending
										if v272 != 0 {
											return
										} else {
											v273 = int32(4459328)
											v274 = *(*int32)(unsafe.Add(mBase, _consts[741]))
											v275 = int32(4459332)
											v276 = *(*int32)(unsafe.Add(mBase, _consts[743]))
											v278 = int64(56)
											v280 = int64(65280)
											v282 = int64(40)
											v285 = int64(16711680)
											v287 = int64(24)
											v289 = int64(4278190080)
											v291 = int64(8)
											*(*int64)(unsafe.Add(mBase, uint32(v274+v276))) = v268<<(uint(v278)%64) | v268&v280<<(uint(v282)%64) | (v268&v285<<(uint(v287)%64) | v268&v289<<(uint(v291)%64)) | (int64(base.Ui64(v268)>>(uint(v291)%64))&v289 | int64(base.Ui64(v268)>>(uint(v287)%64))&v285 | (int64(base.Ui64(v268)>>(uint(v282)%64))&v280 | int64(base.Ui64(v268)>>(uint(v278)%64))))
											v316 = *(*int32)(unsafe.Add(mBase, _consts[743]))
											*(*int32)(unsafe.Add(mBase, _consts[743])) = v316 + int32(8)
											F_enlargeStringInfo(m, v273, int32(1))
											mBase = m.M
											v323 = m.ExcPending
											if v323 != 0 {
												return
											} else {
												v325 = *(*int32)(unsafe.Add(mBase, _consts[741]))
												v326 = int32(4459332)
												v327 = *(*int32)(unsafe.Add(mBase, _consts[743]))
												*(*uint8)(unsafe.Add(mBase, uint32(v325+v327))) = uint8(v2)
												v332 = *(*int32)(unsafe.Add(mBase, _consts[743]))
												*(*int32)(unsafe.Add(mBase, _consts[743])) = v332 + int32(1)
												v338 = F_errstart(m, int32(13), int32(0))
												mBase = m.M
												v339 = m.ExcPending
												if v339 != 0 {
													return
												} else {
													if v338 != 0 {
														v341 = int64(base.Ui64(v67) >> (uint(int64(32)) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)) = uint32(v341)
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v67)
														if v2 != 0 {
															v346 = int32(703410)
														} else {
															v346 = int32(785690)
														}
														*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v346
														v349 = *(*int64)(unsafe.Add(mBase, _consts[738]))
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v349)
														v352 = *(*int64)(unsafe.Add(mBase, _consts[739]))
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v352)
														v354 = int64(32)
														v355 = int64(base.Ui64(v349) >> (uint(v354) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v355)
														v358 = int64(base.Ui64(v352) >> (uint(v354) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v358)
														F_errmsg_internal(m, int32(184919), v9)
														mBase = m.M
														v362 = m.ExcPending
														if v362 != 0 {
															return
														} else {
															F_errfinish(m, int32(515166), int32(1145), int32(19704))
															mBase = m.M
															v367 = m.ExcPending
															if v367 != 0 {
																return
															} else {
																v371 = *(*int32)(unsafe.Add(mBase, _consts[734]))
																v373 = *(*int32)(unsafe.Add(mBase, _consts[741]))
																v375 = *(*int32)(unsafe.Add(mBase, _consts[743]))
																v377 = *(*int32)(unsafe.Add(mBase, _consts[495]))
																v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
																m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
																mBase = m.M
																v380 = m.ExcPending
																if v380 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(32)
																	return
																}
															}
														}
													} else {
														v371 = *(*int32)(unsafe.Add(mBase, _consts[734]))
														v373 = *(*int32)(unsafe.Add(mBase, _consts[741]))
														v375 = *(*int32)(unsafe.Add(mBase, _consts[743]))
														v377 = *(*int32)(unsafe.Add(mBase, _consts[495]))
														v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
														m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
														mBase = m.M
														v380 = m.ExcPending
														if v380 != 0 {
															return
														} else {
															m.G0 = v9 + int32(32)
															return
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v48 = *(*int64)(unsafe.Add(mBase, _consts[740]))
					if v34 < v48 {
						m.G0 = v9 + int32(32)
						return
					} else {
						v50 = v40
						*(*int64)(unsafe.Add(mBase, _consts[738])) = v36
						*(*int64)(unsafe.Add(mBase, _consts[739])) = v50
						v58 = int64(*(*int32)(unsafe.Add(mBase, _consts[735])))
						if v58 <= int64(0) {
							v64 = int64(9223372036854775807)
						} else {
							v64 = v58*int64(1000000) + v34
						}
						*(*int64)(unsafe.Add(mBase, _consts[740])) = v64
						v67 = F_GetXLogReplayRecPtr(m, int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							v69 = int32(4459328)
							v70 = *(*int32)(unsafe.Add(mBase, _consts[741]))
							v71 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v71)
							*(*int32)(unsafe.Add(mBase, _consts[742])) = v71
							*(*int32)(unsafe.Add(mBase, _consts[743])) = v71
							F_enlargeStringInfo(m, int32(4459328), int32(1))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								v81 = int32(4459328)
								v82 = *(*int32)(unsafe.Add(mBase, _consts[741]))
								v83 = int32(4459332)
								v84 = *(*int32)(unsafe.Add(mBase, _consts[743]))
								v86 = int32(114)
								*(*uint8)(unsafe.Add(mBase, uint32(v82+v84))) = uint8(v86)
								v90 = *(*int32)(unsafe.Add(mBase, _consts[743]))
								*(*int32)(unsafe.Add(mBase, _consts[743])) = v90 + int32(1)
								v95 = *(*int64)(unsafe.Add(mBase, _consts[738]))
								F_enlargeStringInfo(m, v81, int32(8))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									v100 = int32(4459328)
									v101 = *(*int32)(unsafe.Add(mBase, _consts[741]))
									v102 = int32(4459332)
									v103 = *(*int32)(unsafe.Add(mBase, _consts[743]))
									v105 = int64(56)
									v107 = int64(65280)
									v109 = int64(40)
									v112 = int64(16711680)
									v114 = int64(24)
									v116 = int64(4278190080)
									v118 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v101+v103))) = v95<<(uint(v105)%64) | v95&v107<<(uint(v109)%64) | (v95&v112<<(uint(v114)%64) | v95&v116<<(uint(v118)%64)) | (int64(base.Ui64(v95)>>(uint(v118)%64))&v116 | int64(base.Ui64(v95)>>(uint(v114)%64))&v112 | (int64(base.Ui64(v95)>>(uint(v109)%64))&v107 | int64(base.Ui64(v95)>>(uint(v105)%64))))
									v143 = *(*int32)(unsafe.Add(mBase, _consts[743]))
									v144 = int32(8)
									*(*int32)(unsafe.Add(mBase, _consts[743])) = v143 + v144
									v148 = *(*int64)(unsafe.Add(mBase, _consts[739]))
									F_enlargeStringInfo(m, v100, v144)
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return
									} else {
										v153 = int32(4459328)
										v154 = *(*int32)(unsafe.Add(mBase, _consts[741]))
										v155 = int32(4459332)
										v156 = *(*int32)(unsafe.Add(mBase, _consts[743]))
										v158 = int64(56)
										v160 = int64(65280)
										v162 = int64(40)
										v165 = int64(16711680)
										v167 = int64(24)
										v169 = int64(4278190080)
										v171 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v154+v156))) = v148<<(uint(v158)%64) | v148&v160<<(uint(v162)%64) | (v148&v165<<(uint(v167)%64) | v148&v169<<(uint(v171)%64)) | (int64(base.Ui64(v148)>>(uint(v171)%64))&v169 | int64(base.Ui64(v148)>>(uint(v167)%64))&v165 | (int64(base.Ui64(v148)>>(uint(v162)%64))&v160 | int64(base.Ui64(v148)>>(uint(v158)%64))))
										v196 = *(*int32)(unsafe.Add(mBase, _consts[743]))
										v197 = int32(8)
										*(*int32)(unsafe.Add(mBase, _consts[743])) = v196 + v197
										F_enlargeStringInfo(m, v153, v197)
										mBase = m.M
										v203 = m.ExcPending
										if v203 != 0 {
											return
										} else {
											v205 = *(*int32)(unsafe.Add(mBase, _consts[741]))
											v206 = int32(4459332)
											v207 = *(*int32)(unsafe.Add(mBase, _consts[743]))
											v209 = int64(56)
											v211 = int64(65280)
											v213 = int64(40)
											v216 = int64(16711680)
											v218 = int64(24)
											v220 = int64(4278190080)
											v222 = int64(8)
											*(*int64)(unsafe.Add(mBase, uint32(v205+v207))) = v67<<(uint(v209)%64) | v67&v211<<(uint(v213)%64) | (v67&v216<<(uint(v218)%64) | v67&v220<<(uint(v222)%64)) | (int64(base.Ui64(v67)>>(uint(v222)%64))&v220 | int64(base.Ui64(v67)>>(uint(v218)%64))&v216 | (int64(base.Ui64(v67)>>(uint(v213)%64))&v211 | int64(base.Ui64(v67)>>(uint(v209)%64))))
											v247 = *(*int32)(unsafe.Add(mBase, _consts[743]))
											*(*int32)(unsafe.Add(mBase, _consts[743])) = v247 + int32(8)
											v254 = m.G0
											v255 = int32(16)
											v256 = v254 - v255
											m.G0 = v256
											F___gettimeofday(m, v256)
											mBase = m.M
											v259 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
											v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+8)))
											m.G0 = v256 + v255
											v268 = v260 + v259*int64(1000000) - int64(946684800000000)
											F_enlargeStringInfo(m, int32(4459328), int32(8))
											mBase = m.M
											v272 = m.ExcPending
											if v272 != 0 {
												return
											} else {
												v273 = int32(4459328)
												v274 = *(*int32)(unsafe.Add(mBase, _consts[741]))
												v275 = int32(4459332)
												v276 = *(*int32)(unsafe.Add(mBase, _consts[743]))
												v278 = int64(56)
												v280 = int64(65280)
												v282 = int64(40)
												v285 = int64(16711680)
												v287 = int64(24)
												v289 = int64(4278190080)
												v291 = int64(8)
												*(*int64)(unsafe.Add(mBase, uint32(v274+v276))) = v268<<(uint(v278)%64) | v268&v280<<(uint(v282)%64) | (v268&v285<<(uint(v287)%64) | v268&v289<<(uint(v291)%64)) | (int64(base.Ui64(v268)>>(uint(v291)%64))&v289 | int64(base.Ui64(v268)>>(uint(v287)%64))&v285 | (int64(base.Ui64(v268)>>(uint(v282)%64))&v280 | int64(base.Ui64(v268)>>(uint(v278)%64))))
												v316 = *(*int32)(unsafe.Add(mBase, _consts[743]))
												*(*int32)(unsafe.Add(mBase, _consts[743])) = v316 + int32(8)
												F_enlargeStringInfo(m, v273, int32(1))
												mBase = m.M
												v323 = m.ExcPending
												if v323 != 0 {
													return
												} else {
													v325 = *(*int32)(unsafe.Add(mBase, _consts[741]))
													v326 = int32(4459332)
													v327 = *(*int32)(unsafe.Add(mBase, _consts[743]))
													*(*uint8)(unsafe.Add(mBase, uint32(v325+v327))) = uint8(v2)
													v332 = *(*int32)(unsafe.Add(mBase, _consts[743]))
													*(*int32)(unsafe.Add(mBase, _consts[743])) = v332 + int32(1)
													v338 = F_errstart(m, int32(13), int32(0))
													mBase = m.M
													v339 = m.ExcPending
													if v339 != 0 {
														return
													} else {
														if v338 != 0 {
															v341 = int64(base.Ui64(v67) >> (uint(int64(32)) % 64))
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)) = uint32(v341)
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v67)
															if v2 != 0 {
																v346 = int32(703410)
															} else {
																v346 = int32(785690)
															}
															*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v346
															v349 = *(*int64)(unsafe.Add(mBase, _consts[738]))
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v349)
															v352 = *(*int64)(unsafe.Add(mBase, _consts[739]))
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v352)
															v354 = int64(32)
															v355 = int64(base.Ui64(v349) >> (uint(v354) % 64))
															*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v355)
															v358 = int64(base.Ui64(v352) >> (uint(v354) % 64))
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v358)
															F_errmsg_internal(m, int32(184919), v9)
															mBase = m.M
															v362 = m.ExcPending
															if v362 != 0 {
																return
															} else {
																F_errfinish(m, int32(515166), int32(1145), int32(19704))
																mBase = m.M
																v367 = m.ExcPending
																if v367 != 0 {
																	return
																} else {
																	v371 = *(*int32)(unsafe.Add(mBase, _consts[734]))
																	v373 = *(*int32)(unsafe.Add(mBase, _consts[741]))
																	v375 = *(*int32)(unsafe.Add(mBase, _consts[743]))
																	v377 = *(*int32)(unsafe.Add(mBase, _consts[495]))
																	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
																	m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
																	mBase = m.M
																	v380 = m.ExcPending
																	if v380 != 0 {
																		return
																	} else {
																		m.G0 = v9 + int32(32)
																		return
																	}
																}
															}
														} else {
															v371 = *(*int32)(unsafe.Add(mBase, _consts[734]))
															v373 = *(*int32)(unsafe.Add(mBase, _consts[741]))
															v375 = *(*int32)(unsafe.Add(mBase, _consts[743]))
															v377 = *(*int32)(unsafe.Add(mBase, _consts[495]))
															v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
															m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
															mBase = m.M
															v380 = m.ExcPending
															if v380 != 0 {
																return
															} else {
																m.G0 = v9 + int32(32)
																return
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_XLogWrite(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v75 int64
	_ = v75
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int64
	_ = v92
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v117 int32
	_ = v117
	var v126 int64
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int64
	_ = v197
	var v198 int64
	_ = v198
	var v202 int64
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v222 int64
	_ = v222
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v234 int64
	_ = v234
	var v235 int64
	_ = v235
	var v239 int64
	_ = v239
	var v291 int32
	_ = v291
	var v292 int64
	_ = v292
	var v296 int32
	_ = v296
	var v314 int32
	_ = v314
	var v315 int64
	_ = v315
	var v319 int32
	_ = v319
	var v336 int32
	_ = v336
	var v337 int64
	_ = v337
	var v342 int32
	_ = v342
	var v343 int64
	_ = v343
	var v348 int32
	_ = v348
	var v359 int32
	_ = v359
	var v365 int64
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v404 int64
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v412 int64
	_ = v412
	var v415 int32
	_ = v415
	var v419 int64
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v427 int64
	_ = v427
	var v428 int64
	_ = v428
	var v429 int64
	_ = v429
	var v432 int64
	_ = v432
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v451 int64
	_ = v451
	var v453 int32
	_ = v453
	var v456 int64
	_ = v456
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v505 int64
	_ = v505
	var v506 int64
	_ = v506
	var v509 int64
	_ = v509
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v522 int64
	_ = v522
	var v526 int64
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v537 int64
	_ = v537
	var v539 int32
	_ = v539
	var v540 int64
	_ = v540
	var v547 int64
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v562 int64
	_ = v562
	var v564 int32
	_ = v564
	var v566 int64
	_ = v566
	var v570 int64
	_ = v570
	var v574 int32
	_ = v574
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v593 int32
	_ = v593
	var v595 int64
	_ = v595
	var v597 int32
	_ = v597
	var v598 int64
	_ = v598
	var v602 int64
	_ = v602
	var v603 int64
	_ = v603
	var v610 int32
	_ = v610
	var v612 int64
	_ = v612
	var v620 int32
	_ = v620
	var v622 int64
	_ = v622
	var v624 int64
	_ = v624
	var v625 int64
	_ = v625
	var v629 int64
	_ = v629
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	v4 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(96)
	m.G0 = v19
	v21 = int32(4444592)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+280)) = v23
	*(*int64)(unsafe.Add(mBase, _consts[189])) = v23
	v28 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+272)) = v29
	*(*int64)(unsafe.Add(mBase, _consts[190])) = v29
	v36 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+304))
	v41 = base.I64_rem_u_s(int64(base.Ui64(v29)>>(uint(int64(13))%64)), base.I64_extend_i32_s(v37+int32(1)))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v48 = v4
	v49 = v36
	v52 = base.I32_wrap_i64(v41)
	v53 = v4
	v54 = v4
	v60 = v43
	goto L2
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L13
	} else {
		goto L107
	}
L2:
	;
	v62 = *(*int64)(unsafe.Add(mBase, _consts[190]))
	if base.Ui64(v60) <= base.Ui64(v62) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v505 = *(*int64)(unsafe.Add(mBase, _consts[189]))
	v506 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui64(v506) <= base.Ui64(v505) {
		goto L82
	} else {
		goto L83
	}
L4:
	;
	goto L3
L5:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v49)+300))
	v67 = v64 + v52<<(uint(int32(3))%32)
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v67)))
	*(*int64)(unsafe.Add(mBase, uint32(v67))) = v68
	v71 = *(*int64)(unsafe.Add(mBase, _consts[190]))
	if base.Ui64(v68) <= base.Ui64(v71) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, _consts[190])) = v68
	v75 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v79 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v81 = base.I64_div_u_s(v68-int64(1), base.I64_extend_i32_s(v79))
	v83 = *(*int64)(unsafe.Add(mBase, _consts[192]))
	if v81 != v83 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	if int32(0) <= v86 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v113 = v79
	v114 = v68
	goto L9
L9:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	if v117 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	F_XLogFileClose(m)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v98 = v81
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, _consts[194])) = l1
	*(*int64)(unsafe.Add(mBase, _consts[192])) = v98
	v104 = F_XLogFileInit(m, v98, l1)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L13
	} else {
		goto L15
	}
L13:
	;
	return
L14:
	;
	v92 = *(*int64)(unsafe.Add(mBase, _consts[190]))
	v96 = int64(*(*int32)(unsafe.Add(mBase, _consts[191])))
	v97 = base.I64_div_u_s(v92-int64(1), v96)
	v98 = v97
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[193])) = v104
	F_ReserveExternalFD(m)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v112 = *(*int64)(unsafe.Add(mBase, _consts[190]))
	v113 = v110
	v114 = v112
	goto L9
L17:
	;
	*(*int32)(unsafe.Add(mBase, _consts[194])) = l1
	v126 = base.I64_div_u_s(v114-int64(1), base.I64_extend_i32_s(v113))
	*(*int64)(unsafe.Add(mBase, _consts[192])) = v126
	v129 = F_XLogFileOpen(m, v126, l1)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L13
	} else {
		goto L20
	}
L18:
	;
	v138 = v113
	v139 = v114
	goto L19
L19:
	;
	if v48 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[193])) = v129
	F_ReserveExternalFD(m)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v137 = *(*int64)(unsafe.Add(mBase, _consts[190]))
	v138 = v135
	v139 = v137
	goto L19
L22:
	;
	v140 = v54
	goto L24
L23:
	;
	v140 = v52
	goto L24
L24:
	;
	v141 = base.B2i32(base.Ui64(v68) <= base.Ui64(v75))
	if v48 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v148 = v53
	goto L27
L26:
	;
	v148 = (base.I32_wrap_i64(v139) + int32(-8192)) & (v138 - int32(1))
	goto L27
L27:
	;
	v150 = v48 + int32(1)
	v152 = v150 << (uint(int32(13)) % 32)
	v155 = v141 & base.B2i32(base.Ui32(v138) <= base.Ui32(v148+v152))
	v157 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	if base.Ui64(v139) < base.Ui64(v75) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v141 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L29:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v157)+304))
	if v155|base.B2i32(v52 == v159) == int32(0) {
		v461 = v150
		v466 = v148
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v157)+296))
	v172 = v152
	v174 = v164 + v140<<(uint(int32(13))%32)
	v176 = v148
	goto L33
L32:
	;
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, _consts[195])))
	v191 = m.G0
	v193 = v191 - int32(16)
	m.G0 = v193
	if v188 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v398 = int32(0)
	if v155 == v398 {
		v461 = v398
		v466 = v397
		goto L28
	} else {
		goto L67
	}
L35:
	;
	v206 = int32(4155324)
	v207 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v207))) = int32(167772240)
	v211 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	v213 = F_pwrite(m, v211, v174, v172, base.I64_extend_i32_u(v176))
	mBase = m.M
	v215 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = int32(0)
	v221 = int32(1)
	v222 = base.I64_extend_i32_s(v213)
	v226 = m.G0
	v228 = v226 - int32(16)
	m.G0 = v228
	if v202 != int64(0) {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	F___clock_gettime(m, int32(1), v193)
	mBase = m.M
	v197 = int64(*(*int32)(unsafe.Add(mBase, uint32(v193)+8)))
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v193)))
	v202 = v197 + v198*int64(1000000000)
	goto L38
L37:
	;
	v202 = int64(0)
	goto L38
L38:
	;
	m.G0 = v193 + int32(16)
	goto L35
L39:
	;
	if v213 <= int32(0) {
		goto L57
	} else {
		goto L58
	}
L40:
	;
	F___clock_gettime(m, int32(1), v228)
	mBase = m.M
	v234 = int64(*(*int32)(unsafe.Add(mBase, uint32(v228)+8)))
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v228)))
	v239 = v234 + (v235*int64(1000000000) - v202)
	goto L43
L41:
	;
	goto L42
L42:
	;
	v336 = int32(4530656)
	v337 = *(*int64)(unsafe.Add(mBase, _consts[196]))
	*(*int64)(unsafe.Add(mBase, _consts[196])) = v337 + base.I64_extend_i32_u(v221)
	v342 = int32(4529696)
	v343 = *(*int64)(unsafe.Add(mBase, _consts[197]))
	*(*int64)(unsafe.Add(mBase, _consts[197])) = v343 + v222
	F_pgstat_count_backend_io_op(m, int32(2), int32(3), int32(7), v221, v222)
	mBase = m.M
	v348 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[154])) = uint8(v348)
	*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v348)
	m.G0 = v228 + int32(16)
	goto L39
L43:
	;
	v291 = int32(4531616)
	v292 = *(*int64)(unsafe.Add(mBase, _consts[199]))
	*(*int64)(unsafe.Add(mBase, _consts[199])) = v292 + v239
	v296 = *(*int32)(unsafe.Add(mBase, _consts[200]))
	if base.Ui32(int32(16)) < base.Ui32(v296) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	goto L42
L54:
	;
	if int32(1)<<(uint(v296)%32)&int32(115186) == int32(0) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v314 = int32(4528512)
	v315 = *(*int64)(unsafe.Add(mBase, _consts[201]))
	*(*int64)(unsafe.Add(mBase, _consts[201])) = v315 + v239
	v319 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[154])) = uint8(v319)
	*(*uint8)(unsafe.Add(mBase, _consts[202])) = uint8(v319)
	goto L53
L56:
	;
	if v395 != 0 {
		v172 = v395
		v174 = v396
		v176 = v397
		goto L33
	} else {
		goto L66
	}
L57:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v359 == int32(27) {
		v395 = v172
		v396 = v174
		v397 = v176
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v395 = v172 - v213
	v396 = v213 + v174
	v397 = v213 + v176
	goto L56
L60:
	;
	v365 = *(*int64)(unsafe.Add(mBase, _consts[192]))
	v367 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	F_XLogFileName(m, v19+int32(32), l1, v365, v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L13
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = v359
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L13
	} else {
		goto L62
	}
L62:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v19 + int32(32)
	F_errmsg(m, int32(304048), v19)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L13
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(518546), int32(2455), int32(364777))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	goto L34
L67:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	v404 = *(*int64)(unsafe.Add(mBase, _consts[192]))
	F_issue_xlog_fsync(m, v402, v404, l1)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L13
	} else {
		goto L68
	}
L68:
	;
	v408 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[203])) = uint8(v408)
	v412 = *(*int64)(unsafe.Add(mBase, _consts[190]))
	*(*int64)(unsafe.Add(mBase, _consts[189])) = v412
	v415 = *(*int32)(unsafe.Add(mBase, _consts[204]))
	if int32(0) < v415 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v419 = *(*int64)(unsafe.Add(mBase, _consts[192]))
	v420 = m.G0
	v422 = v420 - int32(80)
	m.G0 = v422
	*(*int32)(unsafe.Add(mBase, uint32(v422))) = l1
	v427 = int64(*(*int32)(unsafe.Add(mBase, _consts[191])))
	v428 = base.I64_div_u_s(int64(4294967296), v427)
	v429 = base.I64_div_u_s(v419, v428)
	*(*uint32)(unsafe.Add(mBase, uint32(v422)+4)) = uint32(v429)
	v432 = v419 - v428*v429
	*(*uint32)(unsafe.Add(mBase, uint32(v422)+8)) = uint32(v432)
	v438 = F_pg_snprintf(m, v422+int32(16), int32(64), int32(531349), v422)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L13
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v451 = F___time(m)
	mBase = m.M
	v453 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v453)+248)) = v451
	v456 = *(*int64)(unsafe.Add(mBase, _consts[189]))
	*(*int64)(unsafe.Add(mBase, uint32(v453)+256)) = v456
	v461 = v398
	v466 = v397
	goto L28
L72:
	;
	F_XLogArchiveNotify(m, v422+int32(16))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L13
	} else {
		goto L73
	}
L73:
	;
	m.G0 = v422 + int32(80)
	goto L71
L74:
	;
	*(*int64)(unsafe.Add(mBase, _consts[190])) = v75
	goto L4
L75:
	;
	goto L76
L76:
	;
	v482 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v482)+304))
	if v52 != v483 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v485 = v52 + int32(1)
	goto L79
L78:
	;
	v485 = int32(0)
	goto L79
L79:
	;
	if l2 == int32(0) {
		v48 = v461
		v49 = v482
		v52 = v485
		v53 = v466
		v54 = v140
		v60 = v75
		goto L2
	} else {
		goto L80
	}
L80:
	;
	if v461 != 0 {
		v48 = v461
		v49 = v482
		v52 = v485
		v53 = v466
		v54 = v140
		v60 = v75
		goto L2
	} else {
		goto L81
	}
L81:
	;
	goto L4
L82:
	;
	v581 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v581)+440)) = int32(1)
	if v582 != 0 {
		goto L97
	} else {
		goto L98
	}
L83:
	;
	v509 = *(*int64)(unsafe.Add(mBase, _consts[190]))
	if base.Ui64(v509) <= base.Ui64(v505) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v512 = *(*int32)(unsafe.Add(mBase, _consts[205]))
	switch v512 - int32(2) {
	case 0, 2:
		v570 = v509
		goto L85
	default:
		goto L86
	}
L85:
	;
	*(*int64)(unsafe.Add(mBase, _consts[189])) = v570
	v574 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[203])) = uint8(v574)
	goto L82
L86:
	;
	v516 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v518 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	if int32(0) <= v518 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v562 = *(*int64)(unsafe.Add(mBase, _consts[192]))
	F_issue_xlog_fsync(m, v557, v562, l1)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L13
	} else {
		goto L96
	}
L88:
	;
	v522 = *(*int64)(unsafe.Add(mBase, _consts[192]))
	v526 = base.I64_div_u_s(v509-int64(1), base.I64_extend_i32_s(v516))
	if v522 == v526 {
		v557 = v518
		goto L87
	} else {
		goto L91
	}
L89:
	;
	v539 = v516
	v540 = v509
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, _consts[194])) = l1
	v547 = base.I64_div_u_s(v540-int64(1), base.I64_extend_i32_s(v539))
	*(*int64)(unsafe.Add(mBase, _consts[192])) = v547
	v550 = F_XLogFileOpen(m, v547, l1)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L13
	} else {
		goto L94
	}
L91:
	;
	F_XLogFileClose(m)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L13
	} else {
		goto L92
	}
L92:
	;
	v531 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	if int32(0) <= v531 {
		v557 = v531
		goto L87
	} else {
		goto L93
	}
L93:
	;
	v535 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v537 = *(*int64)(unsafe.Add(mBase, _consts[190]))
	v539 = v535
	v540 = v537
	goto L90
L94:
	;
	*(*int32)(unsafe.Add(mBase, _consts[193])) = v550
	F_ReserveExternalFD(m)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L13
	} else {
		goto L95
	}
L95:
	;
	v556 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	v557 = v556
	goto L87
L96:
	;
	v566 = *(*int64)(unsafe.Add(mBase, _consts[190]))
	v570 = v566
	goto L85
L97:
	;
	v586 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_s_lock(m, v586+int32(440), int32(518546), int32(2568), int32(364777))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L13
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v595 = *(*int64)(unsafe.Add(mBase, _consts[190]))
	v597 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v598 = *(*int64)(unsafe.Add(mBase, uint32(v597)+184))
	if base.Ui64(v598) < base.Ui64(v595) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L99
L101:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v597)+184)) = v595
	goto L103
L102:
	;
	goto L103
L103:
	;
	v602 = *(*int64)(unsafe.Add(mBase, _consts[189]))
	v603 = *(*int64)(unsafe.Add(mBase, uint32(v597)+192))
	if base.Ui64(v603) < base.Ui64(v602) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v597)+192)) = v602
	goto L106
L105:
	;
	goto L106
L106:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v597)+272)) = v595
	*(*int32)(unsafe.Add(mBase, uint32(v597)+440)) = int32(0)
	v610 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v612 = *(*int64)(unsafe.Add(mBase, _consts[189]))
	*(*int64)(unsafe.Add(mBase, uint32(v610)+280)) = v612
	m.G0 = v19 + int32(96)
	return
L107:
	;
	v622 = *(*int64)(unsafe.Add(mBase, _consts[190]))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+20)) = uint32(v622)
	v624 = int64(32)
	v625 = int64(base.Ui64(v622) >> (uint(v624) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+16)) = uint32(v625)
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+28)) = uint32(v68)
	v629 = int64(base.Ui64(v68) >> (uint(v624) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+24)) = uint32(v629)
	F_errmsg_internal(m, int32(536603), v19+int32(16))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L13
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(518546), int32(2354), int32(364777))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L13
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
