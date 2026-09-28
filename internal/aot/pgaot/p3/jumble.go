package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__jumbleCreateSeqStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v92 int64
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int64
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v198 int64
	_ = v198
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v235 int64
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v282 int64
	_ = v282
	var v284 int32
	_ = v284
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F__jumbleNode(m, l0, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v10 = l1 + int32(12)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v16 == int32(0) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v70 = v19
			} else {
				v20 = int32(4)
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if base.Ui32(v22-int32(1021)) < base.Ui32(v20) {
					v31 = v22
					v33 = v20
					v35 = l0 + int32(28)
					for {
						if base.Ui32(int32(1024)) <= base.Ui32(v31) {
							v40 = F_hash_bytes_extended(m, v21, int32(1024), int64(0))
							mBase = m.M
							*(*int64)(unsafe.Add(mBase, uint32(v21))) = v40
							v43 = int32(8)
						} else {
							v43 = v31
						}
						v45 = int32(1024) - v43
						if base.Ui32(v33) < base.Ui32(v45) {
							v47 = v33
						} else {
							v47 = v45
						}
						if v47 != 0 {
							base.MemoryCopy(m, v43+v21, v35, v47)
						} else {
						}
						v51 = v43 + v47
						v52 = v33 - v47
						if v52 != 0 {
							v31 = v51
							v33 = v52
							v35 = v47 + v35
							continue
						} else {
							break
						}
						break
					}
					v60 = v51
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v22+v21))) = v16
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v60 = v55 + int32(4)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v60
				v70 = v60
			}
			v75 = int32(4)
			v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if base.Ui32(v70-int32(1021)) < base.Ui32(v75) {
				v82 = v10
				v83 = v70
				v85 = v75
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v83) {
						v92 = F_hash_bytes_extended(m, v76, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v76))) = v92
						v95 = int32(8)
					} else {
						v95 = v83
					}
					v97 = int32(1024) - v95
					if base.Ui32(v85) < base.Ui32(v97) {
						v99 = v85
					} else {
						v99 = v97
					}
					if v99 != 0 {
						base.MemoryCopy(m, v95+v76, v82, v99)
					} else {
					}
					v103 = v95 + v99
					v104 = v85 - v99
					if v104 != 0 {
						v82 = v82 + v99
						v83 = v103
						v85 = v104
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103
			} else {
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				*(*int32)(unsafe.Add(mBase, uint32(v70+v76))) = v107
				v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v109 + int32(4)
			}
			v121 = l1 + int32(16)
			v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v127 == int32(0) {
				v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v181 = v130
			} else {
				v131 = int32(4)
				v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if base.Ui32(v133-int32(1021)) < base.Ui32(v131) {
					v142 = v133
					v143 = v131
					v146 = l0 + int32(28)
					for {
						if base.Ui32(int32(1024)) <= base.Ui32(v142) {
							v151 = F_hash_bytes_extended(m, v132, int32(1024), int64(0))
							mBase = m.M
							*(*int64)(unsafe.Add(mBase, uint32(v132))) = v151
							v154 = int32(8)
						} else {
							v154 = v142
						}
						v156 = int32(1024) - v154
						if base.Ui32(v143) < base.Ui32(v156) {
							v158 = v143
						} else {
							v158 = v156
						}
						if v158 != 0 {
							base.MemoryCopy(m, v154+v132, v146, v158)
						} else {
						}
						v162 = v154 + v158
						v163 = v143 - v158
						if v163 != 0 {
							v142 = v162
							v143 = v163
							v146 = v158 + v146
							continue
						} else {
							break
						}
						break
					}
					v171 = v162
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v133+v132))) = v127
					v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v171 = v166 + int32(4)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v171
				v181 = v171
			}
			v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v181 != int32(1024) {
				v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
				*(*uint8)(unsafe.Add(mBase, uint32(v181+v186))) = uint8(v190)
				v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v192 + int32(1)
			} else {
				v198 = F_hash_bytes_extended(m, v186, int32(1024), int64(0))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v186))) = v198
				v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
				*(*uint8)(unsafe.Add(mBase, uint32(v186)+8)) = uint8(v200)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(9)
			}
			v205 = l1 + int32(17)
			v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v211 == int32(0) {
				v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v265 = v214
			} else {
				v215 = int32(4)
				v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if base.Ui32(v217-int32(1021)) < base.Ui32(v215) {
					v226 = v217
					v227 = v215
					v230 = l0 + int32(28)
					for {
						if base.Ui32(int32(1024)) <= base.Ui32(v226) {
							v235 = F_hash_bytes_extended(m, v216, int32(1024), int64(0))
							mBase = m.M
							*(*int64)(unsafe.Add(mBase, uint32(v216))) = v235
							v238 = int32(8)
						} else {
							v238 = v226
						}
						v240 = int32(1024) - v238
						if base.Ui32(v227) < base.Ui32(v240) {
							v242 = v227
						} else {
							v242 = v240
						}
						if v242 != 0 {
							base.MemoryCopy(m, v238+v216, v230, v242)
						} else {
						}
						v246 = v238 + v242
						v247 = v227 - v242
						if v247 != 0 {
							v226 = v246
							v227 = v247
							v230 = v242 + v230
							continue
						} else {
							break
						}
						break
					}
					v255 = v246
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v217+v216))) = v211
					v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v255 = v250 + int32(4)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v255
				v265 = v255
			}
			v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v265 != int32(1024) {
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
				*(*uint8)(unsafe.Add(mBase, uint32(v265+v270))) = uint8(v274)
				v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v276 + int32(1)
			} else {
				v282 = F_hash_bytes_extended(m, v270, int32(1024), int64(0))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v270))) = v282
				v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
				*(*uint8)(unsafe.Add(mBase, uint32(v270)+8)) = uint8(v284)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(9)
			}
			return
		}
	}
}
func F__jumbleCreateUserMappingStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 int64
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v155 int64
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v202 int64
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v7 != 0 {
			v8 = F_strlen(m, v7)
			mBase = m.M
			v10 = v8 + int32(1)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v16 == int32(0) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v73 = v19
			} else {
				v20 = int32(4)
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if base.Ui32(v22-int32(1021)) < base.Ui32(v20) {
					v32 = v22
					v34 = v20
					v36 = l0 + int32(28)
					for {
						if base.Ui32(int32(1024)) <= base.Ui32(v32) {
							v41 = F_hash_bytes_extended(m, v21, int32(1024), int64(0))
							mBase = m.M
							*(*int64)(unsafe.Add(mBase, uint32(v21))) = v41
							v44 = int32(8)
						} else {
							v44 = v32
						}
						v46 = int32(1024) - v44
						if base.Ui32(v34) < base.Ui32(v46) {
							v48 = v34
						} else {
							v48 = v46
						}
						if v48 != 0 {
							base.MemoryCopy(m, v44+v21, v36, v48)
						} else {
						}
						v52 = v44 + v48
						v53 = v34 - v48
						if v53 != 0 {
							v32 = v52
							v34 = v53
							v36 = v48 + v36
							continue
						} else {
							break
						}
						break
					}
					v62 = v52
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v22+v21))) = v16
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v62 = v56 + int32(4)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v62
				v73 = v62
			}
			v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if base.Ui32(int32(1024)-v73) < base.Ui32(v10) {
				v83 = v7
				v84 = v10
				v85 = v73
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v85) {
						v94 = F_hash_bytes_extended(m, v78, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v78))) = v94
						v97 = int32(8)
					} else {
						v97 = v85
					}
					v99 = int32(1024) - v97
					if base.Ui32(v84) < base.Ui32(v99) {
						v101 = v84
					} else {
						v101 = v99
					}
					if v101 != 0 {
						base.MemoryCopy(m, v97+v78, v83, v101)
					} else {
					}
					v105 = v97 + v101
					v106 = v84 - v101
					if v106 != 0 {
						v83 = v83 + v101
						v84 = v106
						v85 = v105
						continue
					} else {
						break
					}
					break
				}
				v114 = v105
			} else {
				if v10 != 0 {
					base.MemoryCopy(m, v73+v78, v7, v10)
				} else {
				}
				v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v114 = v109 + v10
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v114
		} else {
			v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v120 + int32(1)
		}
		v125 = l1 + int32(12)
		v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v131 == int32(0) {
			v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v185 = v134
		} else {
			v135 = int32(4)
			v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if base.Ui32(v137-int32(1021)) < base.Ui32(v135) {
				v146 = v137
				v147 = v135
				v150 = l0 + int32(28)
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v146) {
						v155 = F_hash_bytes_extended(m, v136, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v136))) = v155
						v158 = int32(8)
					} else {
						v158 = v146
					}
					v160 = int32(1024) - v158
					if base.Ui32(v147) < base.Ui32(v160) {
						v162 = v147
					} else {
						v162 = v160
					}
					if v162 != 0 {
						base.MemoryCopy(m, v158+v136, v150, v162)
					} else {
					}
					v166 = v158 + v162
					v167 = v147 - v162
					if v167 != 0 {
						v146 = v166
						v147 = v167
						v150 = v162 + v150
						continue
					} else {
						break
					}
					break
				}
				v175 = v166
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v137+v136))) = v131
				v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v175 = v170 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v175
			v185 = v175
		}
		v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v185 != int32(1024) {
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
			*(*uint8)(unsafe.Add(mBase, uint32(v185+v190))) = uint8(v194)
			v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v196 + int32(1)
		} else {
			v202 = F_hash_bytes_extended(m, v190, int32(1024), int64(0))
			mBase = m.M
			*(*int64)(unsafe.Add(mBase, uint32(v190))) = v202
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
			*(*uint8)(unsafe.Add(mBase, uint32(v190)+8)) = uint8(v204)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(9)
		}
		v208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		F__jumbleNode(m, l0, v208)
		mBase = m.M
		v210 = m.ExcPending
		if v210 != 0 {
			return
		} else {
			return
		}
	}
}
func F__jumbleJsonArrayQueryConstructor(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v90 int64
	_ = v90
	var v92 int32
	_ = v92
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F__jumbleNode(m, l0, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			F__jumbleNode(m, l0, v9)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v13 = l1 + int32(16)
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v19 == int32(0) {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v73 = v22
				} else {
					v23 = int32(4)
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if base.Ui32(v25-int32(1021)) < base.Ui32(v23) {
						v34 = v25
						v35 = v23
						v38 = l0 + int32(28)
						for {
							if base.Ui32(int32(1024)) <= base.Ui32(v34) {
								v43 = F_hash_bytes_extended(m, v24, int32(1024), int64(0))
								mBase = m.M
								*(*int64)(unsafe.Add(mBase, uint32(v24))) = v43
								v46 = int32(8)
							} else {
								v46 = v34
							}
							v48 = int32(1024) - v46
							if base.Ui32(v35) < base.Ui32(v48) {
								v50 = v35
							} else {
								v50 = v48
							}
							if v50 != 0 {
								base.MemoryCopy(m, v46+v24, v38, v50)
							} else {
							}
							v54 = v46 + v50
							v55 = v35 - v50
							if v55 != 0 {
								v34 = v54
								v35 = v55
								v38 = v50 + v38
								continue
							} else {
								break
							}
							break
						}
						v63 = v54
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v25+v24))) = v19
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v63 = v58 + int32(4)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v63
					v73 = v63
				}
				v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v73 != int32(1024) {
					v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
					*(*uint8)(unsafe.Add(mBase, uint32(v73+v78))) = uint8(v82)
					v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v84 + int32(1)
				} else {
					v90 = F_hash_bytes_extended(m, v78, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v78))) = v90
					v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
					*(*uint8)(unsafe.Add(mBase, uint32(v78)+8)) = uint8(v92)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(9)
				}
				return
			}
		}
	}
}
func F__jumbleJsonTablePath(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int64
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v6 != 0 {
			v7 = F_strlen(m, v6)
			mBase = m.M
			v9 = v7 + int32(1)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v15 == int32(0) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v72 = v18
			} else {
				v19 = int32(4)
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if base.Ui32(v21-int32(1021)) < base.Ui32(v19) {
					v31 = v21
					v33 = v19
					v35 = l0 + int32(28)
					for {
						if base.Ui32(int32(1024)) <= base.Ui32(v31) {
							v40 = F_hash_bytes_extended(m, v20, int32(1024), int64(0))
							mBase = m.M
							*(*int64)(unsafe.Add(mBase, uint32(v20))) = v40
							v43 = int32(8)
						} else {
							v43 = v31
						}
						v45 = int32(1024) - v43
						if base.Ui32(v33) < base.Ui32(v45) {
							v47 = v33
						} else {
							v47 = v45
						}
						if v47 != 0 {
							base.MemoryCopy(m, v43+v20, v35, v47)
						} else {
						}
						v51 = v43 + v47
						v52 = v33 - v47
						if v52 != 0 {
							v31 = v51
							v33 = v52
							v35 = v47 + v35
							continue
						} else {
							break
						}
						break
					}
					v61 = v51
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v21+v20))) = v15
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v61 = v55 + int32(4)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61
				v72 = v61
			}
			v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if base.Ui32(int32(1024)-v72) < base.Ui32(v9) {
				v82 = v6
				v83 = v9
				v84 = v72
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v84) {
						v93 = F_hash_bytes_extended(m, v77, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v77))) = v93
						v96 = int32(8)
					} else {
						v96 = v84
					}
					v98 = int32(1024) - v96
					if base.Ui32(v83) < base.Ui32(v98) {
						v100 = v83
					} else {
						v100 = v98
					}
					if v100 != 0 {
						base.MemoryCopy(m, v96+v77, v82, v100)
					} else {
					}
					v104 = v96 + v100
					v105 = v83 - v100
					if v105 != 0 {
						v82 = v82 + v100
						v83 = v105
						v84 = v104
						continue
					} else {
						break
					}
					break
				}
				v113 = v104
			} else {
				if v9 != 0 {
					base.MemoryCopy(m, v72+v77, v6, v9)
				} else {
				}
				v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v113 = v108 + v9
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v113
			return
		} else {
			v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v119 + int32(1)
			return
		}
	}
}
func F__jumblePLAssignStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int64
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int64
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v207 int64
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != 0 {
		v5 = F_strlen(m, v4)
		mBase = m.M
		v7 = v5 + int32(1)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v13 == int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v70 = v16
		} else {
			v17 = int32(4)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if base.Ui32(v19-int32(1021)) < base.Ui32(v17) {
				v29 = v19
				v31 = v17
				v33 = l0 + int32(28)
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v29) {
						v38 = F_hash_bytes_extended(m, v18, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v18))) = v38
						v41 = int32(8)
					} else {
						v41 = v29
					}
					v43 = int32(1024) - v41
					if base.Ui32(v31) < base.Ui32(v43) {
						v45 = v31
					} else {
						v45 = v43
					}
					if v45 != 0 {
						base.MemoryCopy(m, v41+v18, v33, v45)
					} else {
					}
					v49 = v41 + v45
					v50 = v31 - v45
					if v50 != 0 {
						v29 = v49
						v31 = v50
						v33 = v45 + v33
						continue
					} else {
						break
					}
					break
				}
				v59 = v49
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19+v18))) = v13
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v59 = v53 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v59
			v70 = v59
		}
		v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(int32(1024)-v70) < base.Ui32(v7) {
			v80 = v4
			v81 = v7
			v82 = v70
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v82) {
					v91 = F_hash_bytes_extended(m, v75, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v75))) = v91
					v94 = int32(8)
				} else {
					v94 = v82
				}
				v96 = int32(1024) - v94
				if base.Ui32(v81) < base.Ui32(v96) {
					v98 = v81
				} else {
					v98 = v96
				}
				if v98 != 0 {
					base.MemoryCopy(m, v94+v75, v80, v98)
				} else {
				}
				v102 = v94 + v98
				v103 = v81 - v98
				if v103 != 0 {
					v80 = v80 + v98
					v81 = v103
					v82 = v102
					continue
				} else {
					break
				}
				break
			}
			v111 = v102
		} else {
			if v7 != 0 {
				base.MemoryCopy(m, v70+v75, v4, v7)
			} else {
			}
			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v111 = v106 + v7
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v111
	} else {
		v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v117 + int32(1)
	}
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F__jumbleNode(m, l0, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		return
	} else {
		v125 = l1 + int32(12)
		v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v131 == int32(0) {
			v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v185 = v134
		} else {
			v135 = int32(4)
			v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if base.Ui32(v137-int32(1021)) < base.Ui32(v135) {
				v146 = v137
				v148 = v135
				v150 = l0 + int32(28)
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v146) {
						v155 = F_hash_bytes_extended(m, v136, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v136))) = v155
						v158 = int32(8)
					} else {
						v158 = v146
					}
					v160 = int32(1024) - v158
					if base.Ui32(v148) < base.Ui32(v160) {
						v162 = v148
					} else {
						v162 = v160
					}
					if v162 != 0 {
						base.MemoryCopy(m, v158+v136, v150, v162)
					} else {
					}
					v166 = v158 + v162
					v167 = v148 - v162
					if v167 != 0 {
						v146 = v166
						v148 = v167
						v150 = v162 + v150
						continue
					} else {
						break
					}
					break
				}
				v175 = v166
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v137+v136))) = v131
				v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v175 = v170 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v175
			v185 = v175
		}
		v190 = int32(4)
		v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(v185-int32(1021)) < base.Ui32(v190) {
			v197 = v125
			v198 = v185
			v200 = v190
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v198) {
					v207 = F_hash_bytes_extended(m, v191, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v191))) = v207
					v210 = int32(8)
				} else {
					v210 = v198
				}
				v212 = int32(1024) - v210
				if base.Ui32(v200) < base.Ui32(v212) {
					v214 = v200
				} else {
					v214 = v212
				}
				if v214 != 0 {
					base.MemoryCopy(m, v210+v191, v197, v214)
				} else {
				}
				v218 = v210 + v214
				v219 = v200 - v214
				if v219 != 0 {
					v197 = v197 + v214
					v198 = v218
					v200 = v219
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v218
		} else {
			v222 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
			*(*int32)(unsafe.Add(mBase, uint32(v185+v191))) = v222
			v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v224 + int32(4)
		}
		v235 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		F__jumbleNode(m, l0, v235)
		mBase = m.M
		v237 = m.ExcPending
		if v237 != 0 {
			return
		} else {
			return
		}
	}
}
func F__jumbleTableSampleClause(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int64
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int64
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	v4 = l1 + int32(4)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v10 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v64 = v13
	} else {
		v14 = int32(4)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if base.Ui32(v16-int32(1021)) < base.Ui32(v14) {
			v25 = v16
			v27 = v14
			v29 = l0 + int32(28)
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v25) {
					v34 = F_hash_bytes_extended(m, v15, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v15))) = v34
					v37 = int32(8)
				} else {
					v37 = v25
				}
				v39 = int32(1024) - v37
				if base.Ui32(v27) < base.Ui32(v39) {
					v41 = v27
				} else {
					v41 = v39
				}
				if v41 != 0 {
					base.MemoryCopy(m, v37+v15, v29, v41)
				} else {
				}
				v45 = v37 + v41
				v46 = v27 - v41
				if v46 != 0 {
					v25 = v45
					v27 = v46
					v29 = v41 + v29
					continue
				} else {
					break
				}
				break
			}
			v54 = v45
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16+v15))) = v10
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v54 = v49 + int32(4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v54
		v64 = v54
	}
	v69 = int32(4)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v64-int32(1021)) < base.Ui32(v69) {
		v76 = v4
		v77 = v64
		v79 = v69
		for {
			if base.Ui32(int32(1024)) <= base.Ui32(v77) {
				v86 = F_hash_bytes_extended(m, v70, int32(1024), int64(0))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v70))) = v86
				v89 = int32(8)
			} else {
				v89 = v77
			}
			v91 = int32(1024) - v89
			if base.Ui32(v79) < base.Ui32(v91) {
				v93 = v79
			} else {
				v93 = v91
			}
			if v93 != 0 {
				base.MemoryCopy(m, v89+v70, v76, v93)
			} else {
			}
			v97 = v89 + v93
			v98 = v79 - v93
			if v98 != 0 {
				v76 = v76 + v93
				v77 = v97
				v79 = v98
				continue
			} else {
				break
			}
			break
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v97
	} else {
		v101 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		*(*int32)(unsafe.Add(mBase, uint32(v64+v70))) = v101
		v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103 + int32(4)
	}
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F__jumbleNode(m, l0, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		return
	} else {
		v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		F__jumbleNode(m, l0, v117)
		mBase = m.M
		v119 = m.ExcPending
		if v119 != 0 {
			return
		} else {
			return
		}
	}
}
func F__jumbleVariableShowStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v90 int64
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v3 != 0 {
		v4 = F_strlen(m, v3)
		mBase = m.M
		v6 = v4 + int32(1)
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v12 == int32(0) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v69 = v15
		} else {
			v16 = int32(4)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if base.Ui32(v18-int32(1021)) < base.Ui32(v16) {
				v28 = v18
				v30 = v16
				v32 = l0 + int32(28)
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v28) {
						v37 = F_hash_bytes_extended(m, v17, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v17))) = v37
						v40 = int32(8)
					} else {
						v40 = v28
					}
					v42 = int32(1024) - v40
					if base.Ui32(v30) < base.Ui32(v42) {
						v44 = v30
					} else {
						v44 = v42
					}
					if v44 != 0 {
						base.MemoryCopy(m, v40+v17, v32, v44)
					} else {
					}
					v48 = v40 + v44
					v49 = v30 - v44
					if v49 != 0 {
						v28 = v48
						v30 = v49
						v32 = v44 + v32
						continue
					} else {
						break
					}
					break
				}
				v58 = v48
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18+v17))) = v12
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v58 = v52 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v58
			v69 = v58
		}
		v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(int32(1024)-v69) < base.Ui32(v6) {
			v79 = v3
			v80 = v6
			v81 = v69
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v81) {
					v90 = F_hash_bytes_extended(m, v74, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v74))) = v90
					v93 = int32(8)
				} else {
					v93 = v81
				}
				v95 = int32(1024) - v93
				if base.Ui32(v80) < base.Ui32(v95) {
					v97 = v80
				} else {
					v97 = v95
				}
				if v97 != 0 {
					base.MemoryCopy(m, v93+v74, v79, v97)
				} else {
				}
				v101 = v93 + v97
				v102 = v80 - v97
				if v102 != 0 {
					v79 = v79 + v97
					v80 = v102
					v81 = v101
					continue
				} else {
					break
				}
				break
			}
			v110 = v101
		} else {
			if v6 != 0 {
				base.MemoryCopy(m, v69+v74, v3, v6)
			} else {
			}
			v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v110 = v105 + v6
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v110
		return
	} else {
		v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v116 + int32(1)
		return
	}
}
