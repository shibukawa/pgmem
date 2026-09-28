package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InstrEndLoop(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v16 float64
	_ = v16
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v28 float64
	_ = v28
	var v31 float64
	_ = v31
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+361)))
	if v5 == int32(1) {
		v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		if v8 != int64(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_InstrEndLoop_0), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_InstrEndLoop_1), int32(211), int32(_a_F_InstrEndLoop_2))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v11 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+361)) = uint8(v11)
			v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+368))
			v14 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+368)) = v14
			v16 = *(*float64)(unsafe.Add(mBase, uint32(l0)+384))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+384)) = v14
			v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+376))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+376)) = v14
			v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+392))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+392)) = v19 + v22
			v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v13 + v25
			v28 = *(*float64)(unsafe.Add(mBase, uint32(l0)+400))
			*(*float64)(unsafe.Add(mBase, uint32(l0)+400)) = base.F64_add(v16, v28)
			v31 = *(*float64)(unsafe.Add(mBase, uint32(l0)+416))
			*(*float64)(unsafe.Add(mBase, uint32(l0)+416)) = base.F64_add(v31, float64(1))
			return
		}
	} else {
		return
	}
}
func F_InstrStopCommon(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v22 int64
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v136 int64
	_ = v136
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v147 int32
	_ = v147
	var v150 int64
	_ = v150
	var v152 int64
	_ = v152
	var v153 int64
	_ = v153
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v164 int64
	_ = v164
	var v166 int64
	_ = v166
	var v167 int64
	_ = v167
	var v171 int64
	_ = v171
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v178 int64
	_ = v178
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v8 == int32(1) {
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		if v11 == int64(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v191 = m.ExcPending
			if v191 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_InstrStopCommon_0), int32(0))
				mBase = m.M
				v195 = m.ExcPending
				if v195 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_InstrStopCommon_1), int32(84), int32(_a_F_InstrStopCommon_2))
					mBase = m.M
					v200 = m.ExcPending
					if v200 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			F___clock_gettime(m, int32(1), v6)
			mBase = m.M
			v16 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			v17 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6)+8)))
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
			v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = v16 + (v17 + v18*int64(1000000000) - v22)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			if v28 == int32(1) {
				v32 = l0 + int32(192)
				v34 = l0 + int32(16)
				v35 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
				v37 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[0]))
				v38 = *(*int64)(unsafe.Add(mBase, uint32(v34)))
				*(*int64)(unsafe.Add(mBase, uint32(v32))) = v35 + (v37 - v38)
				v42 = *(*int64)(unsafe.Add(mBase, uint32(v32)+8))
				v44 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[1]))
				v45 = *(*int64)(unsafe.Add(mBase, uint32(v34)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v42 + (v44 - v45)
				v49 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
				v51 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[2]))
				v52 = *(*int64)(unsafe.Add(mBase, uint32(v34)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+16)) = v49 + (v51 - v52)
				v56 = *(*int64)(unsafe.Add(mBase, uint32(v32)+24))
				v58 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[3]))
				v59 = *(*int64)(unsafe.Add(mBase, uint32(v34)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+24)) = v56 + (v58 - v59)
				v63 = *(*int64)(unsafe.Add(mBase, uint32(v32)+32))
				v65 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[4]))
				v66 = *(*int64)(unsafe.Add(mBase, uint32(v34)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+32)) = v63 + (v65 - v66)
				v70 = *(*int64)(unsafe.Add(mBase, uint32(v32)+40))
				v72 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[5]))
				v73 = *(*int64)(unsafe.Add(mBase, uint32(v34)+40))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+40)) = v70 + (v72 - v73)
				v77 = *(*int64)(unsafe.Add(mBase, uint32(v32)+48))
				v79 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[6]))
				v80 = *(*int64)(unsafe.Add(mBase, uint32(v34)+48))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+48)) = v77 + (v79 - v80)
				v84 = *(*int64)(unsafe.Add(mBase, uint32(v32)+56))
				v86 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[7]))
				v87 = *(*int64)(unsafe.Add(mBase, uint32(v34)+56))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+56)) = v84 + (v86 - v87)
				v91 = *(*int64)(unsafe.Add(mBase, uint32(v32)+64))
				v93 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[8]))
				v94 = *(*int64)(unsafe.Add(mBase, uint32(v34)+64))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+64)) = v91 + (v93 - v94)
				v98 = *(*int64)(unsafe.Add(mBase, uint32(v32)+72))
				v100 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[9]))
				v101 = *(*int64)(unsafe.Add(mBase, uint32(v34)+72))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+72)) = v98 + (v100 - v101)
				v105 = *(*int64)(unsafe.Add(mBase, uint32(v32)+80))
				v107 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[10]))
				v108 = *(*int64)(unsafe.Add(mBase, uint32(v34)+80))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+80)) = v105 + (v107 - v108)
				v112 = *(*int64)(unsafe.Add(mBase, uint32(v32)+88))
				v114 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[11]))
				v115 = *(*int64)(unsafe.Add(mBase, uint32(v34)+88))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+88)) = v112 + (v114 - v115)
				v119 = *(*int64)(unsafe.Add(mBase, uint32(v32)+96))
				v121 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[12]))
				v122 = *(*int64)(unsafe.Add(mBase, uint32(v34)+96))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+96)) = v119 + (v121 - v122)
				v126 = *(*int64)(unsafe.Add(mBase, uint32(v32)+104))
				v128 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[13]))
				v129 = *(*int64)(unsafe.Add(mBase, uint32(v34)+104))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+104)) = v126 + (v128 - v129)
				v133 = *(*int64)(unsafe.Add(mBase, uint32(v32)+112))
				v135 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[14]))
				v136 = *(*int64)(unsafe.Add(mBase, uint32(v34)+112))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+112)) = v133 + (v135 - v136)
				v140 = *(*int64)(unsafe.Add(mBase, uint32(v32)+120))
				v142 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[15]))
				v143 = *(*int64)(unsafe.Add(mBase, uint32(v34)+120))
				*(*int64)(unsafe.Add(mBase, uint32(v32)+120)) = v140 + (v142 - v143)
			} else {
			}
			v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
			if v147 == int32(1) {
				v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)+336))
				v152 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[16]))
				v153 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+336)) = v150 + (v152 - v153)
				v157 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
				v159 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[17]))
				v160 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+320)) = v157 + (v159 - v160)
				v164 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
				v166 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[18]))
				v167 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v164 + (v166 - v167)
				v171 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
				v173 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[19]))
				v174 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+344)) = v171 + (v173 - v174)
				v178 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
				v180 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[20]))
				v181 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+352)) = v178 + (v180 - v181)
			} else {
			}
			m.G0 = v6 + int32(16)
			return
		}
	} else {
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v28 == int32(1) {
			v32 = l0 + int32(192)
			v34 = l0 + int32(16)
			v35 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
			v37 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[0]))
			v38 = *(*int64)(unsafe.Add(mBase, uint32(v34)))
			*(*int64)(unsafe.Add(mBase, uint32(v32))) = v35 + (v37 - v38)
			v42 = *(*int64)(unsafe.Add(mBase, uint32(v32)+8))
			v44 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[1]))
			v45 = *(*int64)(unsafe.Add(mBase, uint32(v34)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v42 + (v44 - v45)
			v49 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
			v51 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[2]))
			v52 = *(*int64)(unsafe.Add(mBase, uint32(v34)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+16)) = v49 + (v51 - v52)
			v56 = *(*int64)(unsafe.Add(mBase, uint32(v32)+24))
			v58 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[3]))
			v59 = *(*int64)(unsafe.Add(mBase, uint32(v34)+24))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+24)) = v56 + (v58 - v59)
			v63 = *(*int64)(unsafe.Add(mBase, uint32(v32)+32))
			v65 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[4]))
			v66 = *(*int64)(unsafe.Add(mBase, uint32(v34)+32))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+32)) = v63 + (v65 - v66)
			v70 = *(*int64)(unsafe.Add(mBase, uint32(v32)+40))
			v72 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[5]))
			v73 = *(*int64)(unsafe.Add(mBase, uint32(v34)+40))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+40)) = v70 + (v72 - v73)
			v77 = *(*int64)(unsafe.Add(mBase, uint32(v32)+48))
			v79 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[6]))
			v80 = *(*int64)(unsafe.Add(mBase, uint32(v34)+48))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+48)) = v77 + (v79 - v80)
			v84 = *(*int64)(unsafe.Add(mBase, uint32(v32)+56))
			v86 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[7]))
			v87 = *(*int64)(unsafe.Add(mBase, uint32(v34)+56))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+56)) = v84 + (v86 - v87)
			v91 = *(*int64)(unsafe.Add(mBase, uint32(v32)+64))
			v93 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[8]))
			v94 = *(*int64)(unsafe.Add(mBase, uint32(v34)+64))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+64)) = v91 + (v93 - v94)
			v98 = *(*int64)(unsafe.Add(mBase, uint32(v32)+72))
			v100 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[9]))
			v101 = *(*int64)(unsafe.Add(mBase, uint32(v34)+72))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+72)) = v98 + (v100 - v101)
			v105 = *(*int64)(unsafe.Add(mBase, uint32(v32)+80))
			v107 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[10]))
			v108 = *(*int64)(unsafe.Add(mBase, uint32(v34)+80))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+80)) = v105 + (v107 - v108)
			v112 = *(*int64)(unsafe.Add(mBase, uint32(v32)+88))
			v114 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[11]))
			v115 = *(*int64)(unsafe.Add(mBase, uint32(v34)+88))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+88)) = v112 + (v114 - v115)
			v119 = *(*int64)(unsafe.Add(mBase, uint32(v32)+96))
			v121 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[12]))
			v122 = *(*int64)(unsafe.Add(mBase, uint32(v34)+96))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+96)) = v119 + (v121 - v122)
			v126 = *(*int64)(unsafe.Add(mBase, uint32(v32)+104))
			v128 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[13]))
			v129 = *(*int64)(unsafe.Add(mBase, uint32(v34)+104))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+104)) = v126 + (v128 - v129)
			v133 = *(*int64)(unsafe.Add(mBase, uint32(v32)+112))
			v135 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[14]))
			v136 = *(*int64)(unsafe.Add(mBase, uint32(v34)+112))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+112)) = v133 + (v135 - v136)
			v140 = *(*int64)(unsafe.Add(mBase, uint32(v32)+120))
			v142 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[15]))
			v143 = *(*int64)(unsafe.Add(mBase, uint32(v34)+120))
			*(*int64)(unsafe.Add(mBase, uint32(v32)+120)) = v140 + (v142 - v143)
		} else {
		}
		v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
		if v147 == int32(1) {
			v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)+336))
			v152 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[16]))
			v153 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+336)) = v150 + (v152 - v153)
			v157 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
			v159 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[17]))
			v160 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+320)) = v157 + (v159 - v160)
			v164 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
			v166 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[18]))
			v167 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v164 + (v166 - v167)
			v171 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
			v173 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[19]))
			v174 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+344)) = v171 + (v173 - v174)
			v178 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
			v180 = *(*int64)(unsafe.Add(mBase, _c_F_InstrStopCommon[20]))
			v181 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+352)) = v178 + (v180 - v181)
		} else {
		}
		m.G0 = v6 + int32(16)
		return
	}
}
