package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_btree_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v165 int64
	_ = v165
	var v171 int32
	_ = v171
	var v172 int64
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int64
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v195 int32
	_ = v195
	var v200 int64
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	v10 = m.G0
	v12 = v10 - int32(208)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+48)))
	switch int32(base.Ui32(v16)>>(uint(int32(4))%32)) - int32(3) {
	case 0, 1:
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
		v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+6)))
		v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+8)))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v29
		*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v28
		*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v27
		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v26
		F_appendStringInfo(m, l0, int32(484602), v12+int32(16))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			m.G0 = v12 + int32(208)
			return
		}
	default:
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15))))
		*(*int32)(unsafe.Add(mBase, uint32(v12))) = v21
		F_appendStringInfo(m, l0, int32(58792), v12)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			m.G0 = v12 + int32(208)
			return
		}
	case 3:
		v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15))))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v39
		F_appendStringInfo(m, l0, int32(58104), v12+int32(32))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return
		} else {
			m.G0 = v12 + int32(208)
			return
		}
	case 4:
		v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
		v94 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
		v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+6)))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v96
		*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v95
		*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v94
		if v93 != 0 {
			v102 = int32(84)
		} else {
			v102 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v102
		F_appendStringInfo(m, l0, int32(503589), v12-int32(-64))
		mBase = m.M
		v108 = m.ExcPending
		if v108 != 0 {
			return
		} else {
			v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+119)))
			if v110 != int32(1) {
				m.G0 = v12 + int32(208)
				return
			} else {
				v113 = int32(0)
				v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
				v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+72))
				if v117 < v113 {
					v139 = v113
				} else {
					v123 = v116 + int32(76)
					v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
					if v124 != int32(1) {
						v139 = v113
					} else {
						v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+43)))
						if v127 == int32(0) {
							v139 = v113
						} else {
							v137 = *(*int32)(unsafe.Add(mBase, uint32(v123)+44))
							v139 = v137
						}
					}
				}
				v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
				v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+6)))
				F_delvacuum_desc(m, l0, v139, v143, v144)
				mBase = m.M
				v146 = m.ExcPending
				if v146 != 0 {
					return
				} else {
					m.G0 = v12 + int32(208)
					return
				}
			}
		}
	case 5, 6:
		v158 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
		v159 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
		v160 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
		*(*uint32)(unsafe.Add(mBase, uint32(v12)+128)) = uint32(v160)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = v159
		*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = v158
		v165 = int64(base.Ui64(v160) >> (uint(int64(32)) % 64))
		*(*uint32)(unsafe.Add(mBase, uint32(v12)+124)) = uint32(v165)
		F_appendStringInfo(m, l0, int32(747382), v12+int32(112))
		mBase = m.M
		v171 = m.ExcPending
		if v171 != 0 {
			return
		} else {
			v172 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
			v173 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = v173
			*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = v172
			F_appendStringInfo(m, l0, int32(57646), v12+int32(96))
			mBase = m.M
			v180 = m.ExcPending
			if v180 != 0 {
				return
			} else {
				m.G0 = v12 + int32(208)
				return
			}
		}
	case 7:
		v181 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v181
		F_appendStringInfo(m, l0, int32(58577), v12+int32(144))
		mBase = m.M
		v187 = m.ExcPending
		if v187 != 0 {
			return
		} else {
			m.G0 = v12 + int32(208)
			return
		}
	case 8:
		v147 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
		v148 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
		v149 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v149
		*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = v148
		*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v147
		F_appendStringInfo(m, l0, int32(57693), v12+int32(80))
		mBase = m.M
		v157 = m.ExcPending
		if v157 != 0 {
			return
		} else {
			m.G0 = v12 + int32(208)
			return
		}
	case 9:
		v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15))))
		v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+2)))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v47
		*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v46
		F_appendStringInfo(m, l0, int32(59176), v12+int32(48))
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return
		} else {
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
			v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+119)))
			if v56 != int32(1) {
				m.G0 = v12 + int32(208)
				return
			} else {
				v59 = int32(0)
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+72))
				if v63 < v59 {
					v85 = v59
				} else {
					v69 = v62 + int32(76)
					v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
					if v70 != int32(1) {
						v85 = v59
					} else {
						v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+43)))
						if v73 == int32(0) {
							v85 = v59
						} else {
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v69)+44))
							v85 = v83
						}
					}
				}
				v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15))))
				v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+2)))
				F_delvacuum_desc(m, l0, v85, v89, v90)
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return
				} else {
					m.G0 = v12 + int32(208)
					return
				}
			}
		}
	case 10:
		v188 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
		v189 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
		v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+24)))
		v191 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
		*(*uint32)(unsafe.Add(mBase, uint32(v12)+176)) = uint32(v191)
		if v190 != 0 {
			v195 = int32(84)
		} else {
			v195 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v12)+180)) = v195
		*(*int32)(unsafe.Add(mBase, uint32(v12)+168)) = v189
		*(*int64)(unsafe.Add(mBase, uint32(v12)+160)) = v188
		v200 = int64(base.Ui64(v191) >> (uint(int64(32)) % 64))
		*(*uint32)(unsafe.Add(mBase, uint32(v12)+172)) = uint32(v200)
		F_appendStringInfo(m, l0, int32(503437), v12+int32(160))
		mBase = m.M
		v206 = m.ExcPending
		if v206 != 0 {
			return
		} else {
			m.G0 = v12 + int32(208)
			return
		}
	case 11:
		v207 = int32(0)
		v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
		v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+72))
		if v211 < v207 {
			v233 = v207
		} else {
			v217 = v210 + int32(76)
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
			if v218 != int32(1) {
				v233 = v207
			} else {
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+43)))
				if v221 == int32(0) {
					v233 = v207
				} else {
					v231 = *(*int32)(unsafe.Add(mBase, uint32(v217)+44))
					v233 = v231
				}
			}
		}
		v237 = *(*int32)(unsafe.Add(mBase, uint32(v233)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v237
		F_appendStringInfo(m, l0, int32(58172), v12+int32(192))
		mBase = m.M
		v243 = m.ExcPending
		if v243 != 0 {
			return
		} else {
			m.G0 = v12 + int32(208)
			return
		}
	case 12:
		m.G0 = v12 + int32(208)
		return
	}
}
