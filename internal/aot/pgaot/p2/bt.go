package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetBTPageStatistics(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int64
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if l1 < int32(0) {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_GetBTPageStatistics[0]))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v18+(l1^int32(-1))<<(uint(int32(2))%32))))
		v32 = v24
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_GetBTPageStatistics[1]))
		v32 = v26 + l1<<(uint(int32(13))%32) + int32(-8192)
	}
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)))
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = l0
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v34 - int32(24)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+19)))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v41 << (uint(int32(8)) % 32)
	v45 = v34 + v32
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+12)))
	if v46&int32(4) != 0 {
		if v46&int32(257) == int32(256) {
			v55 = int32(68)
		} else {
			v55 = int32(100)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(l2)+28)) = uint8(v55)
		if v46&int32(256) != 0 {
			v59 = *(*int64)(unsafe.Add(mBase, uint32(v32)+24))
			v60 = int32(0)
			v63 = F_errstart(m, int32(13), v60)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				if v63 == int32(0) {
					v122 = v60
					v124 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v124
					v126 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v126
					v128 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v128
					v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+12)))
					*(*uint16)(unsafe.Add(mBase, uint32(l2)+44)) = uint16(v130)
					v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+14)))
					*(*uint16)(unsafe.Add(mBase, uint32(l2)+46)) = uint16(v132)
					v134 = int32(0)
					v136 = v122 & int32(_a_F_GetBTPageStatistics_0)
					if v136 != 0 {
						v140 = v134
						v141 = int32(1)
						for {
							v153 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(20)+v141<<(uint(int32(2))%32))))
							v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+v153&int32(_a_F_GetBTPageStatistics_1))+6)))
							v160 = int32(_a_F_GetBTPageStatistics_2)
							if v153&v160 != v160 {
								v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v164 + int32(1)
							} else {
								v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v168 + int32(1)
							}
							v172 = v140 + v157&int32(_a_F_GetBTPageStatistics_3)
							if v141 != v136 {
								v140 = v172
								v141 = v141 + int32(1)
								continue
							} else {
								break
							}
							break
						}
						v176 = v172
					} else {
						v176 = v134
					}
					v186 = int32(4)
					v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+14)))
					v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)))
					v189 = v187 - v188
					if v189 <= v186 {
						v192 = v186
					} else {
						v192 = v189
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v192 - int32(4)
					v196 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					v197 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
					v198 = v196 + v197
					if v198 != 0 {
						v199 = base.I32_div_u_s(v176, v198)
						v201 = v199
					} else {
						v201 = int32(0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v201
					m.G0 = v13 + int32(32)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l0
					*(*uint32)(unsafe.Add(mBase, uint32(v13)+24)) = uint32(v59)
					v70 = int64(base.Ui64(v59) >> (uint(int64(32)) % 64))
					*(*uint32)(unsafe.Add(mBase, uint32(v13)+20)) = uint32(v70)
					F_errmsg_internal(m, int32(_a_F_GetBTPageStatistics_4), v13+int32(16))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_GetBTPageStatistics_5), int32(148), int32(_a_F_GetBTPageStatistics_6))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return
						} else {
							v122 = v60
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v124
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v126
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v128
							v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+12)))
							*(*uint16)(unsafe.Add(mBase, uint32(l2)+44)) = uint16(v130)
							v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+14)))
							*(*uint16)(unsafe.Add(mBase, uint32(l2)+46)) = uint16(v132)
							v134 = int32(0)
							v136 = v122 & int32(_a_F_GetBTPageStatistics_0)
							if v136 != 0 {
								v140 = v134
								v141 = int32(1)
								for {
									v153 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(20)+v141<<(uint(int32(2))%32))))
									v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+v153&int32(_a_F_GetBTPageStatistics_1))+6)))
									v160 = int32(_a_F_GetBTPageStatistics_2)
									if v153&v160 != v160 {
										v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v164 + int32(1)
									} else {
										v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v168 + int32(1)
									}
									v172 = v140 + v157&int32(_a_F_GetBTPageStatistics_3)
									if v141 != v136 {
										v140 = v172
										v141 = v141 + int32(1)
										continue
									} else {
										break
									}
									break
								}
								v176 = v172
							} else {
								v176 = v134
							}
							v186 = int32(4)
							v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+14)))
							v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)))
							v189 = v187 - v188
							if v189 <= v186 {
								v192 = v186
							} else {
								v192 = v189
							}
							*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v192 - int32(4)
							v196 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							v197 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
							v198 = v196 + v197
							if v198 != 0 {
								v199 = base.I32_div_u_s(v176, v198)
								v201 = v199
							} else {
								v201 = int32(0)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v201
							m.G0 = v13 + int32(32)
							return
						}
					}
				}
			}
		} else {
			v82 = int32(0)
			v85 = F_errstart(m, int32(13), v82)
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return
			} else {
				if v85 == int32(0) {
					v122 = v82
					v124 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v124
					v126 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v126
					v128 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v128
					v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+12)))
					*(*uint16)(unsafe.Add(mBase, uint32(l2)+44)) = uint16(v130)
					v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+14)))
					*(*uint16)(unsafe.Add(mBase, uint32(l2)+46)) = uint16(v132)
					v134 = int32(0)
					v136 = v122 & int32(_a_F_GetBTPageStatistics_0)
					if v136 != 0 {
						v140 = v134
						v141 = int32(1)
						for {
							v153 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(20)+v141<<(uint(int32(2))%32))))
							v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+v153&int32(_a_F_GetBTPageStatistics_1))+6)))
							v160 = int32(_a_F_GetBTPageStatistics_2)
							if v153&v160 != v160 {
								v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v164 + int32(1)
							} else {
								v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v168 + int32(1)
							}
							v172 = v140 + v157&int32(_a_F_GetBTPageStatistics_3)
							if v141 != v136 {
								v140 = v172
								v141 = v141 + int32(1)
								continue
							} else {
								break
							}
							break
						}
						v176 = v172
					} else {
						v176 = v134
					}
					v186 = int32(4)
					v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+14)))
					v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)))
					v189 = v187 - v188
					if v189 <= v186 {
						v192 = v186
					} else {
						v192 = v189
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v192 - int32(4)
					v196 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					v197 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
					v198 = v196 + v197
					if v198 != 0 {
						v199 = base.I32_div_u_s(v176, v198)
						v201 = v199
					} else {
						v201 = int32(0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v201
					m.G0 = v13 + int32(32)
					return
				} else {
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v89
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
					F_errmsg_internal(m, int32(_a_F_GetBTPageStatistics_7), v13)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_GetBTPageStatistics_5), int32(152), int32(_a_F_GetBTPageStatistics_6))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							v122 = v82
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v124
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v126
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v128
							v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+12)))
							*(*uint16)(unsafe.Add(mBase, uint32(l2)+44)) = uint16(v130)
							v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+14)))
							*(*uint16)(unsafe.Add(mBase, uint32(l2)+46)) = uint16(v132)
							v134 = int32(0)
							v136 = v122 & int32(_a_F_GetBTPageStatistics_0)
							if v136 != 0 {
								v140 = v134
								v141 = int32(1)
								for {
									v153 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(20)+v141<<(uint(int32(2))%32))))
									v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+v153&int32(_a_F_GetBTPageStatistics_1))+6)))
									v160 = int32(_a_F_GetBTPageStatistics_2)
									if v153&v160 != v160 {
										v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v164 + int32(1)
									} else {
										v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v168 + int32(1)
									}
									v172 = v140 + v157&int32(_a_F_GetBTPageStatistics_3)
									if v141 != v136 {
										v140 = v172
										v141 = v141 + int32(1)
										continue
									} else {
										break
									}
									break
								}
								v176 = v172
							} else {
								v176 = v134
							}
							v186 = int32(4)
							v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+14)))
							v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)))
							v189 = v187 - v188
							if v189 <= v186 {
								v192 = v186
							} else {
								v192 = v189
							}
							*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v192 - int32(4)
							v196 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							v197 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
							v198 = v196 + v197
							if v198 != 0 {
								v199 = base.I32_div_u_s(v176, v198)
								v201 = v199
							} else {
								v201 = int32(0)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v201
							m.G0 = v13 + int32(32)
							return
						}
					}
				}
			}
		}
	} else {
		if base.Ui32(int32(25)) <= base.Ui32(v33) {
			v107 = int32(base.Ui32(v33+int32(_a_F_GetBTPageStatistics_8)) >> (uint(int32(2)) % 32))
		} else {
			v107 = int32(0)
		}
		if v46&int32(16) != 0 {
			v110 = int32(101)
			*(*uint8)(unsafe.Add(mBase, uint32(l2)+28)) = uint8(v110)
			v122 = v107
		} else {
			if v46&int32(1) != 0 {
				v114 = int32(108)
				*(*uint8)(unsafe.Add(mBase, uint32(l2)+28)) = uint8(v114)
				v122 = v107
			} else {
				if v46&int32(2) != 0 {
					v118 = int32(114)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+28)) = uint8(v118)
					v122 = v107
				} else {
					v120 = int32(105)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+28)) = uint8(v120)
					v122 = v107
				}
			}
		}
		v124 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
		*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v124
		v126 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v126
		v128 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v128
		v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+12)))
		*(*uint16)(unsafe.Add(mBase, uint32(l2)+44)) = uint16(v130)
		v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+14)))
		*(*uint16)(unsafe.Add(mBase, uint32(l2)+46)) = uint16(v132)
		v134 = int32(0)
		v136 = v122 & int32(_a_F_GetBTPageStatistics_0)
		if v136 != 0 {
			v140 = v134
			v141 = int32(1)
			for {
				v153 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(20)+v141<<(uint(int32(2))%32))))
				v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+v153&int32(_a_F_GetBTPageStatistics_1))+6)))
				v160 = int32(_a_F_GetBTPageStatistics_2)
				if v153&v160 != v160 {
					v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v164 + int32(1)
				} else {
					v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v168 + int32(1)
				}
				v172 = v140 + v157&int32(_a_F_GetBTPageStatistics_3)
				if v141 != v136 {
					v140 = v172
					v141 = v141 + int32(1)
					continue
				} else {
					break
				}
				break
			}
			v176 = v172
		} else {
			v176 = v134
		}
		v186 = int32(4)
		v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+14)))
		v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)))
		v189 = v187 - v188
		if v189 <= v186 {
			v192 = v186
		} else {
			v192 = v189
		}
		*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v192 - int32(4)
		v196 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		v197 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v198 = v196 + v197
		if v198 != 0 {
			v199 = base.I32_div_u_s(v176, v198)
			v201 = v199
		} else {
			v201 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v201
		m.G0 = v13 + int32(32)
		return
	}
}
func F__bt_binsrch_array_skey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int64
	_ = v98
	var v99 int64
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int64
	_ = v235
	var v236 int64
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v273 int32
	_ = v273
	v9 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v16 = v14 - int32(1)
	if l1 == v9 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(0)
	return v273
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v252
	return v258
L3:
	;
	if v132 < v134 {
		goto L52
	} else {
		goto L53
	}
L4:
	;
	v130 = int32(0)
	v132 = v9
	v133 = int32(-1)
	v134 = v16
	goto L3
L5:
	;
	goto L6
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	if l2 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v125 = int32(0)
	v127 = v20 - int32(2)
	if v127 < v125 {
		v252 = v78
		v258 = v125
		goto L2
	} else {
		goto L51
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v118
	return v121
L9:
	;
	v24 = v20 + int32(1)
	if v16 < v24 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v78 = int32(-1)
	v80 = v20 - int32(1)
	if v80 < int32(0) {
		v252 = v78
		v258 = v9
		goto L2
	} else {
		goto L34
	}
L12:
	;
	if v72 <= v16 {
		v130 = v70
		v132 = v72
		v133 = v73
		v134 = v16
		goto L3
	} else {
		goto L33
	}
L13:
	;
	v70 = int32(0)
	v72 = v24
	v73 = int32(-1)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v30 = v28 & int32(1)
	if l4 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v70 = v66
	v72 = v20 + int32(2)
	v73 = v24
	goto L12
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v45+v24<<(uint(int32(3))%32))))
	v50 = F_FunctionCall2Coll(m, l0, v44, l3, v49)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L26
	} else {
		goto L27
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(-1)
	return v24
L19:
	;
	if v30 != 0 {
		v273 = v24
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v30 == int32(0) {
		goto L17
	} else {
		goto L24
	}
L22:
	;
	if v28&int32(33554432) != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v66 = int32(1)
	goto L16
L24:
	;
	if v28&int32(33554432) == int32(0) {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v66 = int32(1)
	goto L16
L26:
	;
	return int32(0)
L27:
	;
	v54 = base.I32_wrap_i64(v50)
	v55 = int32(1)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+3)))
	if v56&v55 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v54 < int32(0) {
		v66 = v55
		goto L16
	} else {
		goto L31
	}
L29:
	;
	v63 = v54
	goto L30
L30:
	;
	if v63 <= int32(0) {
		v118 = v63
		v121 = v24
		goto L8
	} else {
		goto L32
	}
L31:
	;
	v63 = int32(0) - v54
	goto L30
L32:
	;
	v66 = v63
	goto L16
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(1)
	return v16
L34:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v85 = v83 & int32(1)
	if l4 != 0 {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	if v115 < int32(0) {
		v124 = v115
		goto L7
	} else {
		goto L50
	}
L36:
	;
	v115 = int32(0) - v101
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(1)
	return v80
L38:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v94+v80<<(uint(int32(3))%32))))
	v99 = F_FunctionCall2Coll(m, l0, v93, l3, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L26
	} else {
		goto L47
	}
L39:
	;
	v124 = int32(-1)
	goto L7
L40:
	;
	if v85 != 0 {
		v273 = v80
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if v85 == int32(0) {
		goto L38
	} else {
		goto L45
	}
L43:
	;
	if v83&int32(33554432) != 0 {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L37
L45:
	;
	if v83&int32(33554432) != 0 {
		goto L37
	} else {
		goto L46
	}
L46:
	;
	goto L39
L47:
	;
	v101 = base.I32_wrap_i64(v99)
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+3)))
	if v102&int32(1) == int32(0) {
		v115 = v101
		goto L35
	} else {
		goto L48
	}
L48:
	;
	if int32(0) <= v101 {
		goto L36
	} else {
		goto L49
	}
L49:
	;
	goto L37
L50:
	;
	v118 = v115
	v121 = v80
	goto L8
L51:
	;
	v130 = v124
	v132 = v125
	v133 = v80
	v134 = v127
	goto L3
L52:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v140 = v137
	v146 = v132
	v148 = v134
	goto L55
L53:
	;
	v196 = v130
	v203 = v132
	v204 = v133
	goto L54
L54:
	;
	if v203 == v204 {
		goto L85
	} else {
		goto L86
	}
L55:
	;
	v151 = v140 & int32(1)
	v154 = base.I32_div_s(v148-v146, int32(2))
	v155 = v154 + v146
	if l4 != 0 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v196 = v186
	v203 = v193
	v204 = v155
	goto L54
L57:
	;
	v189 = base.B2i32(int32(0) < v186)
	if int32(0) < v186 {
		goto L78
	} else {
		goto L79
	}
L58:
	;
	if v151 != 0 {
		v273 = v155
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v151 != 0 {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	if v140&int32(33554432) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v160 = int32(-1)
	goto L64
L63:
	;
	v160 = int32(1)
	goto L64
L64:
	;
	v186 = v160
	v187 = v140
	goto L57
L65:
	;
	if v140&int32(33554432) != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v167+v155<<(uint(int32(3))%32))))
	v172 = F_FunctionCall2Coll(m, l0, v166, l3, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L26
	} else {
		goto L71
	}
L68:
	;
	v165 = int32(1)
	goto L70
L69:
	;
	v165 = int32(-1)
	goto L70
L70:
	;
	v186 = v165
	v187 = v140
	goto L57
L71:
	;
	v174 = base.I32_wrap_i64(v172)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v175&int32(16777216) != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v186 = int32(1)
	v187 = v175
	goto L57
L73:
	;
	if v174 < int32(0) {
		goto L72
	} else {
		goto L76
	}
L74:
	;
	v182 = v174
	goto L75
L75:
	;
	if v182 == int32(0) {
		v273 = v155
		goto L1
	} else {
		goto L77
	}
L76:
	;
	v182 = int32(0) - v174
	goto L75
L77:
	;
	v186 = v182
	v187 = v175
	goto L57
L78:
	;
	v190 = v148
	goto L80
L79:
	;
	v190 = v155
	goto L80
L80:
	;
	if int32(0) < v186 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v193 = v155 + int32(1)
	goto L83
L82:
	;
	v193 = v146
	goto L83
L83:
	;
	if v193 < v190 {
		v140 = v187
		v146 = v193
		v148 = v190
		goto L55
	} else {
		goto L84
	}
L84:
	;
	goto L56
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v196
	return v203
L86:
	;
	goto L87
L87:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v212 = v210 & int32(1)
	if l4 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if v212 != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	if v212 != 0 {
		goto L97
	} else {
		goto L98
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(0)
	return v203
L92:
	;
	goto L93
L93:
	;
	if v210&int32(33554432) != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v220 = int32(-1)
	goto L96
L95:
	;
	v220 = int32(1)
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v220
	return v203
L97:
	;
	if v210&int32(33554432) != 0 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v231+v203<<(uint(int32(3))%32))))
	v236 = F_FunctionCall2Coll(m, l0, v230, l3, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L26
	} else {
		goto L103
	}
L100:
	;
	v227 = int32(1)
	goto L102
L101:
	;
	v227 = int32(-1)
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v227
	return v203
L103:
	;
	v238 = base.I32_wrap_i64(v236)
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+3)))
	if v239&int32(1) == int32(0) {
		v252 = v238
		v258 = v203
		goto L2
	} else {
		goto L104
	}
L104:
	;
	v245 = int32(0)
	if v238 < v245 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v249 = int32(1)
	goto L107
L106:
	;
	v249 = v245 - v238
	goto L107
L107:
	;
	v252 = v249
	v258 = v203
	goto L2
}
func F__bt_check_third_page(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	v6 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+6)))
	v18 = (v12&int32(_a_F__bt_check_third_page_0) + int32(7)) & int32(_a_F__bt_check_third_page_1)
	if base.B2i32(base.Ui32(v18) < base.Ui32(int32(2705)))|base.B2i32(l2|base.B2i32(base.Ui32(int32(2712)) < base.Ui32(v18)) == v6) == v6 {
		v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+16)))
		v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3+v29)+12)))
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			if v31&int32(1) == int32(0) {
				v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v126 + int32(4)
				F_errmsg_internal(m, int32(_a_F__bt_check_third_page_2), v10)
				mBase = m.M
				v133 = m.ExcPending
				if v133 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F__bt_check_third_page_3), int32(1145), int32(_a_F__bt_check_third_page_4))
					mBase = m.M
					v138 = m.ExcPending
					if v138 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if l2 != 0 {
						v46 = int32(2704)
					} else {
						v46 = int32(2712)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v46
					if l2 != 0 {
						v50 = int32(4)
					} else {
						v50 = int32(3)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v50
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v18
					*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v43 + int32(4)
					F_errmsg(m, int32(_a_F__bt_check_third_page_5), v10+int32(32))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+6)))
						if v63&int32(_a_F__bt_check_third_page_6) == int32(0) {
							v91 = l4
						} else {
							v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
							if v68&int32(_a_F__bt_check_third_page_6) == int32(0) {
								v73 = int32(0)
								if v68&int32(_a_F__bt_check_third_page_7) == v73 {
									v89 = v73
									v91 = v89
								} else {
									v91 = l4 + v63&int32(_a_F__bt_check_third_page_0) - int32(6)
								}
							} else {
								v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+2)))
								v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
								v89 = v83 + (l4 + v84<<(uint(int32(16))%32))
								v91 = v89
							}
						}
						v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+2)))
						v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91))))
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
						v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v95
						*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v94 + int32(4)
						v100 = int32(16)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v92 | v93<<(uint(v100)%32)
						v107 = F_errdetail(m, int32(_a_F__bt_check_third_page_8), v10+v100)
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return
						} else {
							F_errhint(m, int32(_a_F__bt_check_third_page_9), int32(0))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return
							} else {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								F_errtableconstraint(m, l1, v113+int32(4))
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F__bt_check_third_page_3), int32(1161), int32(_a_F__bt_check_third_page_4))
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
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
			}
		}
	} else {
		m.G0 = v10 + int32(48)
		return
	}
}
func F__bt_delitems_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	return v3 - v4
}
func F__bt_delitems_delete_check(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v341 int32
	_ = v341
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
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
	var v593 int32
	_ = v593
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v687 int32
	_ = v687
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v774 int32
	_ = v774
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v836 int32
	_ = v836
	var v845 int32
	_ = v845
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v860 int64
	_ = v860
	var v861 int32
	_ = v861
	var v862 int64
	_ = v862
	var v863 int32
	_ = v863
	var v864 int64
	_ = v864
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	v5 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(3296)
	m.G0 = v27
	if l1 < v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+188))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+76))
	v49 = m.T0[v48].(func(*base.Module, int32, int32) int32)(m, l2, l3)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[0]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32+(l1^int32(-1))<<(uint(int32(2))%32))))
	v46 = v38
	goto L1
L3:
	;
	goto L4
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[1]))
	v46 = v40 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	return
L6:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[2]))
	if v52 <= int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[2]))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	F_pg_qsort(m, v89, v90, int32(8), int32(214))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L23
	}
L8:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[3])))
	if v56&int32(1) == int32(0) {
		v86 = v5
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+118)))
	if v62 != int32(112) {
		v86 = v5
		goto L7
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	if int32(0) < v52 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	goto L17
L14:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v67 != 0 {
		v86 = v5
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v68 == int32(0) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v86 = v5
	goto L7
L17:
	;
	if base.Ui32(v72) < base.Ui32(int32(_a_F__bt_delitems_delete_check_0)) {
		v86 = int32(1)
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l2)+180))
	if v75 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v86 = int32(0)
	goto L7
L20:
	;
	goto L21
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+119)))
	switch v81 - int32(109) {
	case 0, 5:
		goto L22
	default:
		v86 = int32(0)
		goto L7
	}
L22:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+112)))
	v86 = v84
	goto L7
L23:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v95 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	m.G0 = v27 + int32(3296)
	return
L25:
	;
	if int32(0) < v95 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v109 = int32(0)
	v113 = v5
	v115 = v5
	v116 = v5
	goto L29
L27:
	;
	v414 = v5
	v416 = v5
	goto L28
L28:
	;
	if l1 < int32(0) {
		goto L73
	} else {
		goto L74
	}
L29:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v132 = int32(*(*int16)(unsafe.Add(mBase, uint32(v128+v116<<(uint(int32(3))%32))+6)))
	v135 = v127 + v132*int32(6)
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135))))
	if v136 == v109&int32(_a_F__bt_delitems_delete_check_1) {
		v382 = v109
		v386 = v113
		v388 = v115
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v414 = v386
	v416 = v388
	goto L28
L31:
	;
	v401 = v116 + int32(1)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v401 < v402 {
		v109 = v382
		v113 = v386
		v115 = v388
		v116 = v401
		goto L29
	} else {
		goto L71
	}
L32:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(20)+v136<<(uint(int32(2))%32))))
	v146 = v46 + v143&int32(_a_F__bt_delitems_delete_check_2)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+7)))
	if v147&int32(32) != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v166 = v150 & int32(4095)
	if v166 == int32(0) {
		v362 = v113
		v364 = v115
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146)+4)))
	if v150&int32(_a_F__bt_delitems_delete_check_3) != 0 {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+2)))
	if v154 != int32(1) {
		v382 = v109
		v386 = v113
		v388 = v115
		goto L31
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v159 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(1648)+v115<<(uint(v159)%32)))) = uint16(v136)
	v382 = v109
	v386 = v113
	v388 = v115 + v159
	goto L31
L39:
	;
	v382 = v136
	v386 = v362
	v388 = v364
	goto L31
L40:
	;
	v173 = int32(0)
	v177 = v116
	v183 = v173
	v184 = v173
	goto L41
L41:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v199 <= v177 {
		v305 = v177
		v311 = v183
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v311 == int32(0) {
		v362 = v113
		v364 = v115
		goto L39
	} else {
		goto L66
	}
L43:
	;
	v328 = v184 + int32(1)
	if v328 != v166 {
		v177 = v305
		v183 = v311
		v184 = v328
		goto L41
	} else {
		goto L65
	}
L44:
	;
	v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146)+2)))
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146))))
	v209 = v201 + (v146 + v202<<(uint(int32(16))%32)) + v184*int32(6)
	v213 = v177
	v217 = int32(-1)
	v218 = v199
	goto L45
L45:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v239 = v236 + v213<<(uint(int32(3))%32)
	v240 = int32(*(*int16)(unsafe.Add(mBase, uint32(v239)+6)))
	v243 = v235 + v240*int32(6)
	v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v243))))
	if v244 != v136 {
		v281 = v213
		v282 = v217
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v282 != 0 {
		v305 = v281
		v311 = v183
		goto L43
	} else {
		goto L59
	}
L47:
	;
	goto L46
L48:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+2)))
	if v246 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v252 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v239)+2)))
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v239))))
	v254 = int32(16)
	v256 = v252 | v253<<(uint(v254)%32)
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209)+2)))
	v258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209))))
	v261 = v257 | v258<<(uint(v254)%32)
	if base.Ui32(v256) < base.Ui32(v261) {
		v272 = int32(-1)
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v276 = v217
	v277 = v218
	goto L51
L51:
	;
	v279 = v213 + int32(1)
	if v279 < v277 {
		v213 = v279
		v217 = v276
		v218 = v277
		goto L45
	} else {
		goto L58
	}
L52:
	;
	if int32(0) <= v272 {
		v281 = v213
		v282 = v272
		goto L47
	} else {
		goto L57
	}
L53:
	;
	goto L52
L54:
	;
	if base.Ui32(v261) < base.Ui32(v256) {
		v272 = int32(1)
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v239)+4)))
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209)+4)))
	if base.Ui32(v266) < base.Ui32(v267) {
		v272 = int32(-1)
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v272 = base.B2i32(base.Ui32(v267) < base.Ui32(v266))
	goto L53
L57:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v276 = v272
	v277 = v275
	goto L51
L58:
	;
	v281 = v279
	v282 = v276
	goto L47
L59:
	;
	if v183 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v294 = int32(1)
	v295 = v292 + v294
	*(*uint16)(unsafe.Add(mBase, uint32(v293)+6)) = uint16(v295)
	*(*uint16)(unsafe.Add(mBase, uint32(v293+v292&int32(_a_F__bt_delitems_delete_check_1)<<(uint(v294)%32))+8)) = uint16(v184)
	v305 = v281
	v311 = v293
	goto L43
L61:
	;
	v284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+6)))
	v292 = v284
	v293 = v183
	goto L60
L62:
	;
	goto L63
L63:
	;
	v286 = F_palloc(m, v166<<(uint(int32(1))%32)+int32(8))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	v288 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v286)+6)) = uint16(v288)
	*(*uint16)(unsafe.Add(mBase, uint32(v286)+4)) = uint16(v136)
	*(*int32)(unsafe.Add(mBase, uint32(v286))) = v146
	v292 = int32(0)
	v293 = v286
	goto L60
L65:
	;
	goto L42
L66:
	;
	v332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v311)+6)))
	if v166 == v332 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(1648)+v115<<(uint(int32(1))%32)))) = uint16(v136)
	F_pfree(m, v311)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L5
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v113<<(uint(int32(2))%32)))) = v311
	v362 = v113 + int32(1)
	v364 = v115
	goto L39
L70:
	;
	v362 = v113
	v364 = v115 + int32(1)
	goto L39
L71:
	;
	goto L30
L72:
	;
	v446 = int32(0)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+118)))
	if v449 != int32(112) {
		v462 = v446
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v431 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[0]))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v431+(l1^int32(-1))<<(uint(int32(2))%32))))
	v445 = v437
	goto L72
L74:
	;
	goto L75
L75:
	;
	v439 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[1]))
	v445 = v439 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L72
L76:
	;
	if int32(0) < v414 {
		goto L82
	} else {
		goto L83
	}
L77:
	;
	v454 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[2]))
	if int32(0) < v454 {
		v462 = int32(1)
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v458 != 0 {
		v462 = int32(0)
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v462 = base.B2i32(v459 == int32(0))
	goto L76
L80:
	;
	if int32(0) < v416 {
		goto L120
	} else {
		goto L121
	}
L81:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L5
	} else {
		goto L113
	}
L82:
	;
	v468 = int32(0)
	v472 = v446
	goto L85
L83:
	;
	goto L84
L84:
	;
	v736 = int32(0)
	v737 = int32(_a_F__bt_delitems_delete_check_4)
	v739 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[4])) = v739 + int32(1)
	v784 = v736
	v787 = v736
	goto L80
L85:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v468<<(uint(int32(2))%32))))
	F__bt_update_posting(m, v495)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L5
	} else {
		goto L87
	}
L86:
	;
	v514 = int32(0)
	if v462 != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v498 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v495)+6)))
	v501 = int32(1)
	v504 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v495)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(2480)+v468<<(uint(v501)%32)))) = uint16(v504)
	v510 = v472 + v498<<(uint(v501)%32) + int32(2)
	v512 = v468 + v501
	if v512 != v414 {
		v468 = v512
		v472 = v510
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v517 = int32(0)
	v518 = F_palloc(m, v510)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L5
	} else {
		goto L92
	}
L90:
	;
	v659 = v514
	v662 = v514
	goto L91
L91:
	;
	v679 = int32(_a_F__bt_delitems_delete_check_4)
	v681 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[4])) = v681 + int32(1)
	v687 = v514
	goto L108
L92:
	;
	if v414 != int32(1) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v659 = v510
	v662 = v518
	goto L91
L94:
	;
	v530 = v517
	v531 = v514
	v538 = int32(0)
	goto L97
L95:
	;
	v592 = v517
	v593 = v514
	goto L96
L96:
	;
	v613 = v593 + v518
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v592<<(uint(int32(2))%32))))
	v620 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v619)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v613))) = uint16(v620)
	v623 = v620 << (uint(int32(1)) % 32)
	if v623 == int32(0) {
		goto L93
	} else {
		goto L107
	}
L97:
	;
	v554 = int32(2)
	v556 = v27 + int32(16) + v530<<(uint(v554)%32)
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	v558 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v557)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v531+v518))) = uint16(v558)
	v561 = v531 + v554
	v563 = v558 << (uint(int32(1)) % 32)
	if v563 != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if v414&int32(1) == int32(0) {
		goto L93
	} else {
		goto L106
	}
L99:
	;
	base.MemoryCopy(m, v561+v518, v557+int32(8), v563)
	goto L101
L100:
	;
	goto L101
L101:
	;
	v568 = v561 + v563
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
	v571 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v570)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v518+v568))) = uint16(v571)
	v574 = v568 + int32(2)
	v576 = v571 << (uint(int32(1)) % 32)
	if v576 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	base.MemoryCopy(m, v574+v518, v570+int32(8), v576)
	goto L104
L103:
	;
	goto L104
L104:
	;
	v581 = v574 + v576
	v582 = int32(2)
	v583 = v530 + v582
	v585 = v538 + v582
	if v585 != v414&int32(2147483646) {
		v530 = v583
		v531 = v581
		v538 = v585
		goto L97
	} else {
		goto L105
	}
L105:
	;
	goto L98
L106:
	;
	v592 = v583
	v593 = v581
	goto L96
L107:
	;
	base.MemoryCopy(m, v613+int32(2), v619+int32(8), v623)
	goto L93
L108:
	;
	v714 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(2480)+v687<<(uint(int32(1))%32)))))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v687<<(uint(int32(2))%32))))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v720)))
	v722 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v721)+6)))
	v729 = F_PageIndexTupleOverwrite(m, v445, v714, v721, (v722&int32(_a_F__bt_delitems_delete_check_5)+int32(7))&int32(_a_F__bt_delitems_delete_check_6))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L5
	} else {
		goto L110
	}
L109:
	;
	v784 = v659
	v787 = v662
	goto L80
L110:
	;
	if v729 == int32(0) {
		goto L81
	} else {
		goto L111
	}
L111:
	;
	v734 = v687 + int32(1)
	if v734 != v414 {
		v687 = v734
		goto L108
	} else {
		goto L112
	}
L112:
	;
	goto L109
L113:
	;
	if l1 < int32(0) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v766
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v767 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_delitems_delete_check_7), v27)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L5
	} else {
		goto L118
	}
L115:
	;
	v751 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[5]))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v751+(l1^int32(-1))*int32(56))+16))
	v766 = v757
	goto L114
L116:
	;
	goto L117
L117:
	;
	v759 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[6]))
	v760 = int32(56)
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v759+l1*v760-v760)+16))
	v766 = v765
	goto L114
L118:
	;
	F_errfinish(m, int32(_a_F__bt_delitems_delete_check_8), int32(1349), int32(_a_F__bt_delitems_delete_check_9))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	F_PageIndexMultiDelete(m, v445, v27+int32(1648), v416)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L5
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v810 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v445)+16)))
	v811 = v445 + v810
	v812 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v811)+12)))
	v814 = v812 & int32(_a_F__bt_delitems_delete_check_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v811)+12)) = uint16(v814)
	F_MarkBufferDirty(m, l1)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L5
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	if v462 != 0 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v445))) = base.I64_rotl(v864, int64(32))
	v869 = int32(_a_F__bt_delitems_delete_check_4)
	v871 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[4])) = v871 - int32(1)
	if v787 != 0 {
		goto L146
	} else {
		goto L147
	}
L126:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+2476)) = uint8(v86)
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+2474)) = uint16(v414)
	v820 = int32(0)
	if v820 < v88 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L128
L128:
	;
	v862 = F_XLogGetFakeLSN(m, l0)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L5
	} else {
		goto L145
	}
L129:
	;
	v823 = v49
	goto L131
L130:
	;
	v823 = v820
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+2468)) = v823
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+2472)) = uint16(v416)
	F_XLogBeginInsert(m)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L5
	} else {
		goto L133
	}
L133:
	;
	F_XLogRegisterData(m, v27+int32(2468), int32(9))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L5
	} else {
		goto L134
	}
L134:
	;
	if int32(0) < v416 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	F_XLogRegisterBufData(m, int32(0), v27+int32(1648), v416<<(uint(int32(1))%32))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L5
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	if int32(0) < v414 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	goto L137
L139:
	;
	F_XLogRegisterBufData(m, int32(0), v27+int32(2480), v414<<(uint(int32(1))%32))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L5
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v860 = F_XLogInsert(m, int32(11), int32(112))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L5
	} else {
		goto L144
	}
L142:
	;
	F_XLogRegisterBufData(m, int32(0), v787, v784)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L5
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	v864 = v860
	goto L125
L145:
	;
	v864 = v862
	goto L125
L146:
	;
	F_pfree(m, v787)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L5
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	if v414 <= int32(0) {
		goto L24
	} else {
		goto L150
	}
L149:
	;
	goto L148
L150:
	;
	v881 = int32(0)
	goto L151
L151:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v881<<(uint(int32(2))%32))))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v908)))
	F_pfree(m, v909)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L5
	} else {
		goto L153
	}
L152:
	;
	v918 = int32(0)
	goto L155
L153:
	;
	v913 = v881 + int32(1)
	if v913 != v414 {
		v881 = v913
		goto L151
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v918<<(uint(int32(2))%32))))
	F_pfree(m, v945)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L5
	} else {
		goto L157
	}
L156:
	;
	goto L24
L157:
	;
	v949 = v918 + int32(1)
	if v949 != v414 {
		v918 = v949
		goto L155
	} else {
		goto L158
	}
L158:
	;
	goto L156
}
func F__bt_killitems(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v335 int32
	_ = v335
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v470 int32
	_ = v470
	v2 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v2
	if int32(2) <= v21 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	F_pg_qsort(m, v26, v21, int32(4), int32(245))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v77 = v21
	goto L3
L3:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+40)))
	if v92 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	return
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	v33 = int32(1)
	v38 = v2
	goto L6
L6:
	;
	v51 = int32(2)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v31+v33<<(uint(v51)%32))))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v31+v38<<(uint(v51)%32))))
	if v54 == v58 {
		v68 = v38
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v77 = v68 + int32(1)
	goto L3
L8:
	;
	v70 = v33 + int32(1)
	if v70 != v21 {
		v33 = v70
		v38 = v68
		goto L6
	} else {
		goto L11
	}
L9:
	;
	v61 = v38 + int32(1)
	if v61 == v33 {
		v68 = v33
		goto L8
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31+v61<<(uint(int32(2))%32)))) = v54
	v68 = v61
	goto L8
L11:
	;
	goto L7
L12:
	;
	F_UnlockReleaseBuffer(m, v452)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L4
	} else {
		goto L82
	}
L13:
	;
	if v107 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	F__bt_lockbuf(m, v95, int32(1))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	v101 = F__bt_getbuf(m, v19, v99, int32(1))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	v107 = v95
	goto L13
L18:
	;
	v103 = F_BufferGetLSNAtomic(m, v101)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v20)+72))
	if v103 != v105 {
		v452 = v101
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v107 = v101
	goto L13
L21:
	;
	if v77 <= int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F__bt_killitems[0]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v111+(v107^int32(-1))<<(uint(int32(2))%32))))
	v125 = v117
	goto L21
L23:
	;
	goto L24
L24:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F__bt_killitems[1]))
	v125 = v119 + v107<<(uint(int32(13))%32) + int32(-8192)
	goto L21
L25:
	;
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+40)))
	if v448 != 0 {
		v452 = v107
		goto L12
	} else {
		goto L80
	}
L26:
	;
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125)+16)))
	v131 = v125 + v130
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	if v132 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v133 = int32(2)
	goto L29
L28:
	;
	v133 = int32(1)
	goto L29
L29:
	;
	v137 = v20 + int32(104)
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v138) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v146 = int32(base.Ui32(v138+int32(_a_F__bt_killitems_0)) >> (uint(int32(2)) % 32))
	goto L32
L31:
	;
	v146 = int32(0)
	goto L32
L32:
	;
	v148 = v146 & int32(_a_F__bt_killitems_1)
	v157 = v2
	v164 = v2
	goto L34
L33:
	;
	v422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131)+12)))
	v424 = v422 | int32(64)
	*(*uint16)(unsafe.Add(mBase, uint32(v131)+12)) = uint16(v424)
	v426 = int32(1)
	F_BufferFinishSetHintBits(m, v107, v426, v426)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L4
	} else {
		goto L79
	}
L34:
	;
	v168 = v157 + int32(1)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v169+v157<<(uint(int32(2))%32))))
	v176 = v137 + v173*int32(10)
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176)+6)))
	if base.B2i32(base.Ui32(v177) < base.Ui32(v133))|base.B2i32(base.Ui32(v148) < base.Ui32(v177)) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v164 == int32(0) {
		goto L25
	} else {
		goto L78
	}
L36:
	;
	v190 = v177
	v192 = v176
	goto L39
L37:
	;
	goto L38
L38:
	;
	if v77 != v168 {
		v157 = v168
		goto L34
	} else {
		goto L77
	}
L39:
	;
	v205 = v125 + int32(20) + v190&int32(_a_F__bt_killitems_1)<<(uint(int32(2))%32)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v209 = v125 + v206&int32(_a_F__bt_killitems_2)
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+7)))
	if v210&int32(32) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	goto L38
L41:
	;
	v379 = v190 + int32(1)
	if base.Ui32(v379&int32(_a_F__bt_killitems_1)) <= base.Ui32(v148) {
		v190 = v379
		v192 = v369
		goto L39
	} else {
		goto L76
	}
L42:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v345 = int32(_a_F__bt_killitems_3)
	if v344&v345 == v345 {
		v369 = v335
		goto L41
	} else {
		goto L69
	}
L43:
	;
	if v312 != v222 {
		v369 = v316
		goto L41
	} else {
		goto L68
	}
L44:
	;
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209)+2)))
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209))))
	v290 = int32(16)
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v192)+2)))
	v294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v192))))
	if v288|v289<<(uint(v290)%32) == v293|v294<<(uint(v290)%32) {
		goto L63
	} else {
		goto L64
	}
L45:
	;
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209)+4)))
	if v215&int32(_a_F__bt_killitems_4) == int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v220 = int32(0)
	v222 = v215 & int32(4095)
	if v222 == v220 {
		v312 = v220
		v316 = v192
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v225 = v168
	v230 = v220
	v234 = v192
	goto L48
L48:
	;
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209)+2)))
	v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209))))
	v245 = int32(16)
	v251 = v243 + (v209 + v244<<(uint(v245)%32)) + v230*int32(6)
	v252 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v251)+2)))
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v251))))
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v234)+2)))
	v258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v234))))
	if v252|v253<<(uint(v245)%32) == v257|v258<<(uint(v245)%32) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v335 = v283
	goto L42
L50:
	;
	if v268 == int32(0) {
		v312 = v230
		v316 = v234
		goto L43
	} else {
		goto L56
	}
L51:
	;
	goto L50
L52:
	;
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v251)+4)))
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v234)+4)))
	if v264 == v265 {
		v268 = int32(1)
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v268 = int32(0)
	goto L51
L55:
	;
	goto L54
L56:
	;
	if v225 < v77 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v272+v225<<(uint(int32(2))%32))))
	v282 = v225 + int32(1)
	v283 = v137 + v276*int32(10)
	goto L59
L58:
	;
	v282 = v225
	v283 = v234
	goto L59
L59:
	;
	v285 = v230 + int32(1)
	if v285 != v222 {
		v225 = v282
		v230 = v285
		v234 = v283
		goto L48
	} else {
		goto L60
	}
L60:
	;
	goto L49
L61:
	;
	if v304 == int32(0) {
		v369 = v192
		goto L41
	} else {
		goto L67
	}
L62:
	;
	goto L61
L63:
	;
	v300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209)+4)))
	v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v192)+4)))
	if v300 == v301 {
		v304 = int32(1)
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v304 = int32(0)
	goto L62
L66:
	;
	goto L65
L67:
	;
	v335 = v192
	goto L42
L68:
	;
	v335 = v316
	goto L42
L69:
	;
	if v164 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v354 = v344
	goto L72
L71:
	;
	v349 = F_BufferBeginSetHintBits(m, v107)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L4
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = v354 | int32(_a_F__bt_killitems_3)
	if v77 != v168 {
		v157 = v168
		v164 = int32(1)
		goto L34
	} else {
		goto L75
	}
L73:
	;
	if v349 == int32(0) {
		goto L25
	} else {
		goto L74
	}
L74:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v354 = v353
	goto L72
L75:
	;
	goto L33
L76:
	;
	goto L40
L77:
	;
	goto L35
L78:
	;
	goto L33
L79:
	;
	goto L25
L80:
	;
	F_UnlockBuffer(m, v107)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	return
L82:
	;
	return
}
func F__bt_lockbuf(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	if l1 == int32(0) {
		F_UnlockBuffer(m, l0)
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	} else {
		F_LockBufferInternal(m, l0, l1)
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F__bt_mkscankey(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v118 int32
	_ = v118
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+10)))
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v64 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)) = uint8(v64)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+2)) = uint16(v64)
	if v63 < v23 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
	if v24&int32(32) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	v56 = F_palloc(m, v23*int32(56)+int32(16))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L12
	}
L5:
	;
	v42 = F_palloc(m, v23*int32(56)+int32(16))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+8)))
	v37 = v35
	goto L5
L7:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v29&int32(_a_F__bt_mkscankey_0) != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v37 = v29 & int32(4095)
	goto L5
L9:
	;
	return int32(0)
L10:
	;
	F__bt_metaversion(m, l0, v42, v42+int32(1))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v60 = v42
	v61 = v50
	v63 = v37
	goto L1
L12:
	;
	v58 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v56))) = uint16(v58)
	v60 = v56
	v61 = int32(1)
	v63 = int32(0)
	goto L1
L13:
	;
	v69 = v63
	goto L15
L14:
	;
	v69 = v23
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+12)) = v69
	v71 = int32(0)
	if base.B2i32(l1 == v71)|base.B2i32(v61&int32(1) == v71) != 0 {
		v107 = v71
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v107
	if int32(0) < v23 {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v79&int32(_a_F__bt_mkscankey_0) == int32(0) {
		v107 = l1
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v84&int32(_a_F__bt_mkscankey_0) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v89 = int32(0)
	if v84&int32(_a_F__bt_mkscankey_1) == v89 {
		v107 = v89
		goto L16
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v107 = v99 + (l1 + v100<<(uint(int32(16))%32))
	goto L16
L22:
	;
	v107 = l1 + v79&int32(_a_F__bt_mkscankey_2) - int32(6)
	goto L16
L23:
	;
	v118 = int32(0)
	goto L26
L24:
	;
	goto L25
L25:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+13)))
	if v194 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L26:
	;
	v129 = int32(1)
	v132 = v118 + v129
	v133 = base.I32_extend16_s(v132)
	v135 = F_index_getprocinfo(m, l0, v133, v129)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L9
	} else {
		goto L28
	}
L27:
	;
	goto L25
L28:
	;
	if v118 < v63 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v140 = F_index_getattr_2(m, l1, v132, v21, v18+int32(15))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L9
	} else {
		goto L32
	}
L30:
	;
	v143 = v129
	v144 = int64(0)
	goto L31
L31:
	;
	v147 = v60 + int32(16) + v118*int32(56)
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+v118<<(uint(int32(1))%32)))))
	v155 = int32(0)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156+v118<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v147)+48)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v147)+12)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v147)+8)) = v155
	*(*uint16)(unsafe.Add(mBase, uint32(v147)+6)) = uint16(v155)
	*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v133)
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v143 | v151<<(uint(int32(24))%32)
	v171 = *(*int32)(unsafe.Add(mBase, _c_F__bt_mkscankey[0]))
	F_fmgr_info_copy(m, v147+int32(16), v135, v171)
	mBase = m.M
	goto L33
L32:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)))
	v143 = v142
	v144 = v140
	goto L31
L33:
	;
	if v143&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v175 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)) = uint8(v175)
	goto L36
L35:
	;
	goto L36
L36:
	;
	if v132 != v23 {
		v118 = v132
		goto L26
	} else {
		goto L37
	}
L37:
	;
	goto L27
L38:
	;
	v197 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)) = uint8(v197)
	goto L40
L39:
	;
	goto L40
L40:
	;
	m.G0 = v18 + int32(16)
	return v60
}
func F__bt_parallel_seize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int64
	_ = v113
	var v115 int32
	_ = v115
	var v116 int64
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v139 int64
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v20 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v20
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = int64(-4294967296)
	v26 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+88)) = uint16(v26)
	if l3 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v16 + int32(16)
	return v223
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v34 = v18 + v33
	v38 = v34 + int32(40)
	v40 = v34 + int32(12)
	goto L7
L3:
	;
	v28 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+19)) = uint8(v28)
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+17)) = uint16(v28)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+17)))
	if v32 != 0 {
		v223 = v5
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L2
L7:
	;
	v55 = F_LWLockAcquire(m, v40, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	if v59 != int32(2) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v62 = int32(0)
	switch v59 - int32(1) {
	case 0:
		goto L17
	default:
		goto L16
	case 2:
		goto L18
	case 3:
		v178 = v62
		v179 = v5
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_LWLockRelease(m, v40)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L50
	}
L14:
	;
	F_LWLockRelease(m, v40)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L9
	} else {
		goto L38
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v170
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v174
	v178 = int32(1)
	v179 = v5
	goto L14
L16:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v170 = v169
	goto L15
L17:
	;
	if l3 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v65 != 0 {
		v170 = v65
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v178 = v62
	v179 = int32(1)
	goto L14
L20:
	;
	v161 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+19)) = uint8(v161)
	v163 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+17)) = uint16(v163)
	F_LWLockRelease(m, v40)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L9
	} else {
		goto L36
	}
L21:
	;
	v69 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v38 + v71<<(uint(v69)%32)
	if v71 <= int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v80 = int32(0)
	goto L23
L23:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v96 = v93 + v80<<(uint(int32(5))%32)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v100 = v92 + v97*int32(56)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v101 != int32(-1) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L20
L25:
	;
	v145 = v80 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v145 < v146 {
		v80 = v145
		goto L23
	} else {
		goto L35
	}
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v38+v80<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v109+v107<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v100)+48)) = v113
	goto L25
L27:
	;
	goto L28
L28:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+18)))
	if v115 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v100)+48)) = int64(0)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v126
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v128 + int32(4)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
	if v132&int32(24) != 0 {
		goto L25
	} else {
		goto L33
	}
L30:
	;
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v100)+48))
	if v116 == int64(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	F_pfree(m, base.I32_wrap_i64(v116))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v139 = F_datumRestore(m, v16+int32(12), v16+int32(11))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v100)+48)) = v139
	goto L25
L35:
	;
	goto L24
L36:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	v223 = l3
	goto L1
L38:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	if v179 == int32(0) {
		v223 = v178
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v186 == int32(0) {
		v223 = v178
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+17)))
	if v190 != 0 {
		v223 = v178
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v186)+24))
	v192 = v186 + v191
	v194 = v192 + int32(12)
	v196 = F_LWLockAcquire(m, v194, int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	if v198 != int32(4) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192)+8)) = int32(4)
	F_LWLockRelease(m, v194)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L9
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	F_LWLockRelease(m, v194)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L9
	} else {
		goto L49
	}
L47:
	;
	F_ConditionVariableBroadcast(m, v192+int32(28))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	v223 = v178
	goto L1
L49:
	;
	v223 = v178
	goto L1
L50:
	;
	F_ConditionVariableSleep(m, v34+int32(28), int32(134217735))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	goto L7
}
func F__bt_readnextpage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int64
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v452 int32
	_ = v452
	var v468 int32
	_ = v468
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = l1
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l3 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l1 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v25 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+88)) = uint8(v25)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v27 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+89)) = uint8(v27)
	goto L1
L5:
	;
	m.G0 = v17 + int32(32)
	return v468
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = int64(-4294967296)
	F__bt_parallel_done(m, l0)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L22
	} else {
		goto L124
	}
L7:
	;
	if l3 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v35 = int32(4)
	goto L10
L9:
	;
	v35 = int32(0)
	goto L10
L10:
	;
	v40 = l4
	goto L11
L11:
	;
	v51 = base.B2i32(l3 != int32(1))
	if v51 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+40)))
	if v421 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L13:
	;
	v61 = v40 & int32(1)
	if v61 != 0 {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+89)))
	if v54 == int32(0) {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+88)))
	if v57 == int32(0) {
		goto L6
	} else {
		goto L18
	}
L17:
	;
	goto L13
L18:
	;
	goto L13
L19:
	;
	if v51 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v62 == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v70 = F__bt_parallel_seize(m, l0, v17+int32(28), v17+int32(24), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(0)
L23:
	;
	if v70 != 0 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = int64(-4294967296)
	v468 = int32(0)
	goto L5
L25:
	;
	if v342 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[0]))
	if v80 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v96 = v89
	v98 = v88
	goto L36
L29:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L22
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v85 = F__bt_getbuf(m, v22, v83, int32(1))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L22
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v85
	v342 = v85
	goto L25
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L22
	} else {
		goto L88
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L22
	} else {
		goto L85
	}
L36:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[0]))
	if v105 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	F_UnlockReleaseBuffer(m, v287)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L22
	} else {
		goto L84
	}
L38:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L22
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v110 = F__bt_getbuf(m, v22, v108, int32(1))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L22
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129)+16)))
	v131 = v130 + v129
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v133 = int32(0)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+12)))
	if base.B2i32(v134&int32(4) == v133)&base.B2i32(v132 == v96) == v133 {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	if v110 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[1]))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v115+(v110^int32(-1))<<(uint(int32(2))%32))))
	v129 = v121
	goto L42
L45:
	;
	goto L46
L46:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[2]))
	v129 = v123 + v110<<(uint(int32(13))%32) + int32(-8192)
	goto L42
L47:
	;
	v212 = F__bt_relandgetbuf(m, v22, v147, v96, int32(1))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L22
	} else {
		goto L63
	}
L48:
	;
	v144 = v132
	v147 = v110
	v153 = v133
	goto L51
L49:
	;
	v198 = v110
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v198
	if v198 == int32(0) {
		goto L6
	} else {
		goto L60
	}
L51:
	;
	if base.B2i32(v144 == int32(0))|base.B2i32(v153 == int32(4)) != 0 {
		goto L47
	} else {
		goto L53
	}
L52:
	;
	v198 = v164
	goto L50
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v144
	v164 = F__bt_relandgetbuf(m, v22, v147, v144, int32(1))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L22
	} else {
		goto L55
	}
L54:
	;
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+16)))
	v187 = v183 + v186
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+12)))
	if v189&int32(4)|base.B2i32(v188 != v96) != 0 {
		v144 = v188
		v147 = v164
		v153 = v153 + int32(1)
		goto L51
	} else {
		goto L59
	}
L55:
	;
	if v164 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[1]))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v169+(v164^int32(-1))<<(uint(int32(2))%32))))
	v183 = v175
	goto L54
L57:
	;
	goto L58
L58:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[2]))
	v183 = v177 + v164<<(uint(int32(13))%32) + int32(-8192)
	goto L54
L59:
	;
	goto L52
L60:
	;
	v342 = v198
	goto L25
L61:
	;
	if v297 != 0 {
		goto L80
	} else {
		goto L81
	}
L62:
	;
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v231)+16)))
	v233 = v232 + v231
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+12)))
	if v234&int32(4) != 0 {
		goto L67
	} else {
		goto L68
	}
L63:
	;
	if v212 < int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[1]))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v217+(v212^int32(-1))<<(uint(int32(2))%32))))
	v231 = v223
	goto L62
L65:
	;
	goto L66
L66:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[2]))
	v231 = v225 + v212<<(uint(int32(13))%32) + int32(-8192)
	goto L62
L67:
	;
	v238 = v233
	v241 = v212
	goto L70
L68:
	;
	goto L69
L69:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	if v281 == v98 {
		goto L34
	} else {
		goto L79
	}
L70:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	if v251 == int32(0) {
		goto L35
	} else {
		goto L72
	}
L71:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	v287 = v255
	v289 = v251
	v297 = v280
	goto L61
L72:
	;
	v255 = F__bt_relandgetbuf(m, v22, v241, v251, int32(1))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L22
	} else {
		goto L74
	}
L73:
	;
	v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v274)+16)))
	v276 = v275 + v274
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+12)))
	if v277&int32(4) != 0 {
		v238 = v276
		v241 = v255
		goto L70
	} else {
		goto L78
	}
L74:
	;
	if v255 < int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[1]))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v260+(v255^int32(-1))<<(uint(int32(2))%32))))
	v274 = v266
	goto L73
L76:
	;
	goto L77
L77:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[2]))
	v274 = v268 + v255<<(uint(int32(13))%32) + int32(-8192)
	goto L73
L78:
	;
	goto L71
L79:
	;
	v287 = v212
	v289 = v96
	v297 = v281
	goto L61
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v297
	F_UnlockReleaseBuffer(m, v287)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L22
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	goto L37
L83:
	;
	v96 = v289
	v98 = v297
	goto L36
L84:
	;
	goto L6
L85:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v307 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_readnextpage_0), v17+int32(16))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L22
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F__bt_readnextpage_1), int32(2041), int32(_a_F__bt_readnextpage_2))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L22
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v325 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_readnextpage_3), v17)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L22
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F__bt_readnextpage_1), int32(2059), int32(_a_F__bt_readnextpage_2))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L22
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v369)+16)))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v371
	v373 = v369 + v370
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+12)))
	if v374&int32(20) == int32(0) {
		goto L97
	} else {
		goto L98
	}
L92:
	;
	v355 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[1]))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v355+(v342^int32(-1))<<(uint(int32(2))%32))))
	v369 = v361
	goto L91
L93:
	;
	goto L94
L94:
	;
	v363 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[2]))
	v369 = v363 + v342<<(uint(int32(13))%32) + int32(-8192)
	goto L91
L95:
	;
	goto L12
L96:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	F_UnlockReleaseBuffer(m, v414)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L22
	} else {
		goto L115
	}
L97:
	;
	if v51 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v373+v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v406
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v408 == int32(0) {
		goto L96
	} else {
		goto L113
	}
L100:
	;
	v381 = int32(1)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	if v384 != 0 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	v390 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v369)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v390) {
		goto L108
	} else {
		goto L109
	}
L103:
	;
	v385 = int32(2)
	goto L105
L104:
	;
	v385 = v381
	goto L105
L105:
	;
	v386 = F__bt_readpage(m, l0, v381, v385, v61)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L22
	} else {
		goto L106
	}
L106:
	;
	if v386 != 0 {
		goto L95
	} else {
		goto L107
	}
L107:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v388
	goto L96
L108:
	;
	v398 = int32(base.Ui32(v390+int32(_a_F__bt_readnextpage_4)) >> (uint(int32(2)) % 32))
	goto L110
L109:
	;
	v398 = int32(0)
	goto L110
L110:
	;
	v401 = F__bt_readpage(m, l0, l3, v398&int32(_a_F__bt_readnextpage_5), v61)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L22
	} else {
		goto L111
	}
L111:
	;
	if v401 != 0 {
		goto L95
	} else {
		goto L112
	}
L112:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v403
	goto L96
L113:
	;
	F__bt_parallel_release(m, l0, v406, v371)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L22
	} else {
		goto L114
	}
L114:
	;
	goto L96
L115:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	if v418 != 0 {
		v40 = int32(0)
		goto L11
	} else {
		goto L116
	}
L116:
	;
	goto L6
L117:
	;
	v468 = int32(1)
	goto L5
L118:
	;
	F_UnlockBuffer(m, v420)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L22
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v426 = F_BufferGetLSNAtomic(m, v420)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L22
	} else {
		goto L122
	}
L121:
	;
	goto L117
L122:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+72)) = v426
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	F_UnlockReleaseBuffer(m, v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L22
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = int32(0)
	goto L117
L124:
	;
	v468 = int32(0)
	goto L5
}
func F__bt_recsplitloc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	v2 = l1
	v5 = int32(0)
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v9 == v2 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = v11 + (l2 - v12)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
		if v17&int32(1) != 0 {
			v65 = v16
			v66 = v14
			v68 = v15
			v70 = v65
			v71 = v66
			v72 = int32(-8)
			v73 = v68
			v75 = v70
			v76 = v71
			v77 = v72
			v78 = v73
			v79 = int32(1)
		} else {
			v75 = v16
			v76 = v14
			v77 = v5
			v78 = v15
			v79 = v5
		}
	} else {
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
		if base.B2i32(v20 != int32(1))|base.B2i32(base.Ui32(l3) < base.Ui32(int32(65))) == int32(0) {
			v28 = int32(-8)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v2<<(uint(int32(2))%32))+20))
			v36 = v29 + v33&int32(_a_F__bt_recsplitloc_0)
			v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+6)))
			if v37&int32(_a_F__bt_recsplitloc_1) == int32(0) {
				v53 = v28
			} else {
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+5)))
				if v42&int32(32) == int32(0) {
					v53 = v28
				} else {
					v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+2)))
					v53 = v37&int32(_a_F__bt_recsplitloc_2) - v49 - int32(8)
				}
			}
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v70 = l3
			v71 = v54 + (l2 - v55)
			v72 = v53
			v73 = v58
			v75 = v70
			v76 = v71
			v77 = v72
			v78 = v73
			v79 = int32(1)
		} else {
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v62 = v59 + (l2 - v60)
			v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v20 != 0 {
				v65 = l3
				v66 = v62
				v68 = v63
				v70 = v65
				v71 = v66
				v72 = int32(-8)
				v73 = v68
				v75 = v70
				v76 = v71
				v77 = v72
				v78 = v73
				v79 = int32(1)
			} else {
				v75 = l3
				v76 = v62
				v77 = int32(0)
				v78 = v63
				v79 = v5
			}
		}
	}
	v82 = v77 + v78 - (l2 + v75)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v79 != 0 {
		v88 = int32(0)
	} else {
		v88 = v75 + int32(_a_F__bt_recsplitloc_3)
	}
	v89 = v76 - v83 + v88
	if (v82|v89)&int32(_a_F__bt_recsplitloc_4) == int32(0) {
		v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		if base.Ui32(v95) < base.Ui32(v75) {
			v97 = v95
		} else {
			v97 = v75
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v97
		v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v101 = int32(10)
		v104 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v99+v100*v101))) = uint16(v104)
		v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*uint16)(unsafe.Add(mBase, uint32(v106+v107*v101)+2)) = uint16(v82)
		v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*uint16)(unsafe.Add(mBase, uint32(v112+v113*v101)+4)) = uint16(v89)
		v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*uint16)(unsafe.Add(mBase, uint32(v118+v119*v101)+6)) = uint16(v2)
		v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*uint8)(unsafe.Add(mBase, uint32(v124+v125*v101)+8)) = uint8(v104)
		v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v131 + int32(1)
	} else {
	}
	return
}
func F__bt_restore_meta(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = F_XLogInitBufferForRedo(m, l0, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v15 = v9 + int32(12)
		v16 = int32(0)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
		if v18 < l1 {
			v40 = v16
			v43 = v40
		} else {
			v22 = v17 + l1*int32(52)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+76)))
			if v23 != int32(1) {
				v40 = v16
				v43 = v40
			} else {
				v27 = v22 + int32(76)
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+43)))
				if v28 == int32(0) {
					if v15 == int32(0) {
						v40 = v16
						v43 = v40
					} else {
						v33 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v33
						v43 = v33
					}
				} else {
					if v15 != 0 {
						v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+48)))
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v36
					} else {
					}
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
					v40 = v38
					v43 = v40
				}
			}
		}
		if v12 < int32(0) {
			v47 = *(*int32)(unsafe.Add(mBase, _c_F__bt_restore_meta[0]))
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v47+(v12^int32(-1))<<(uint(int32(2))%32))))
			v61 = v53
		} else {
			v55 = *(*int32)(unsafe.Add(mBase, _c_F__bt_restore_meta[1]))
			v61 = v55 + v12<<(uint(int32(13))%32) + int32(-8192)
		}
		F_PageInit(m, v61, int32(_a_F__bt_restore_meta_0), int32(16))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v61)+24)) = int32(_a_F__bt_restore_meta_1)
		v67 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
		*(*int32)(unsafe.Add(mBase, uint32(v61)+28)) = v67
		v69 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v61)+32)) = v69
		v71 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v61)+36)) = v71
		v73 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v61)+40)) = v73
		v75 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v61)+44)) = v75
		v77 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
		*(*int64)(unsafe.Add(mBase, uint32(v61)+56)) = int64(-4616189618054758400)
		*(*int32)(unsafe.Add(mBase, uint32(v61)+48)) = v77
		v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+24)))
		*(*uint8)(unsafe.Add(mBase, uint32(v61)+64)) = uint8(v81)
		v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+16)))
		v85 = int32(8)
		*(*uint16)(unsafe.Add(mBase, uint32(v61+v83)+12)) = uint16(v85)
		*(*int64)(unsafe.Add(mBase, uint32(v61))) = base.I64_rotl(v11, int64(32))
		v90 = int32(72)
		*(*uint16)(unsafe.Add(mBase, uint32(v61)+12)) = uint16(v90)
		F_MarkBufferDirty(m, v12)
		mBase = m.M
		v93 = m.ExcPending
		if v93 != 0 {
			return
		} else {
			F_UnlockReleaseBuffer(m, v12)
			mBase = m.M
			v95 = m.ExcPending
			if v95 != 0 {
				return
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		}
	}
}
func F__bt_restore_page(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(2448)
	m.G0 = v9
	if v4 < l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = l1
	v18 = v4
	goto L4
L2:
	;
	v45 = v4
	goto L3
L3:
	;
	v48 = v45
	goto L8
L4:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(816)+v18<<(uint(int32(2))%32)))) = v15
	v27 = int32(1)
	v35 = (v20&int32(_a_F__bt_restore_page_0) + int32(7)) & int32(_a_F__bt_restore_page_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v9+v18<<(uint(v27)%32)))) = uint16(v35)
	v38 = v18 + v27
	v39 = v15 + v35
	if base.Ui32(v39) < base.Ui32(l1+l2) {
		v15 = v39
		v18 = v38
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v45 = v38
	goto L3
L6:
	;
	goto L5
L7:
	;
	m.G0 = v9 + int32(2448)
	return
L8:
	;
	v54 = v48 - int32(1)
	if v54 < int32(0) {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L11
	} else {
		goto L14
	}
L10:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(816)+v54<<(uint(int32(2))%32))))
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9+v54<<(uint(int32(1))%32)))))
	v71 = F_PageAddItemExtended(m, l0, v62, v66, (v45-v54)&int32(_a_F__bt_restore_page_2), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	if v71 != 0 {
		v48 = v54
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	F_errmsg_internal(m, int32(_a_F__bt_restore_page_3), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F__bt_restore_page_4), int32(75), int32(_a_F__bt_restore_page_5))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__bt_scanbehind_checkkeys(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v16&int32(32) == v4 {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)+192))
		v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+8)))
		v30 = v28
	} else {
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
		if v21&int32(_a_F__bt_scanbehind_checkkeys_0) != 0 {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)+192))
			v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+8)))
			v30 = v28
		} else {
			v30 = v21 & int32(4095)
		}
	}
	v31 = int32(0)
	v35 = F__bt_tuple_before_array_skeys(m, l0, l1, l2, v15, v30, v31, v31, v11+int32(7))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		return int32(0)
	} else {
		if v35 != 0 {
			v89 = v4
			m.G0 = v11 + int32(16)
			return v89
		} else {
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+7)))
			if v39&int32(1) != 0 {
				v89 = v4
				m.G0 = v11 + int32(16)
				return v89
			} else {
				v42 = int32(1)
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+19)))
				if v43 != v42 {
					v89 = v42
					m.G0 = v11 + int32(16)
					return v89
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
					if v48&int32(32) == int32(0) {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v46)+192))
						v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59)+8)))
						v62 = v60
					} else {
						v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
						if v53&int32(_a_F__bt_scanbehind_checkkeys_0) != 0 {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v46)+192))
							v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59)+8)))
							v62 = v60
						} else {
							v62 = v53 & int32(4095)
						}
					}
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v64 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v64
					v75 = F__bt_check_compare(m, l0, v64-l1, l2, v62, v47, v64, v64, v11+int32(15), v11+int32(8))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
						if v77 == int32(0) {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
							v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80+v81*int32(56))+6)))
							if v85 != int32(3) {
								v89 = v64
							} else {
								v89 = int32(1)
							}
						} else {
							v89 = int32(1)
						}
						m.G0 = v11 + int32(16)
						return v89
					}
				}
			}
		}
	}
}
func F__bt_search(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	v15 = F__bt_getroot(m, l0, l1, l4)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v15
	if v15 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v25 = base.B2i32(l4 == int32(3))
	v33 = v15
	v34 = int32(1)
	v35 = int32(0)
	goto L6
L6:
	;
	v41 = F__bt_moveright(m, l0, l1, l2, v33, v25, v35, v34)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	if base.B2i32(l4 != int32(3))|base.B2i32(v34 != int32(1)) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v41
	if v41 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+16)))
	v63 = v62 + v61
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+12)))
	if v64&int32(1) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F__bt_search[0]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47+(v41^int32(-1))<<(uint(int32(2))%32))))
	v61 = v53
	goto L9
L11:
	;
	goto L12
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F__bt_search[1]))
	v61 = v55 + v41<<(uint(int32(13))%32) + int32(-8192)
	goto L9
L13:
	;
	v69 = F__bt_binsrch(m, l0, l2, v41)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L7
L16:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v61+v69<<(uint(int32(2))%32))+20))
	v77 = v61 + v74&int32(_a_F__bt_search_0)
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77))))
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+2)))
	if l5 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v84 = F_palloc(m, int32(12))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v110 = v35
	goto L19
L19:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	if v114 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v86 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = v35
	*(*uint16)(unsafe.Add(mBase, uint32(v84)+4)) = uint16(v69)
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v105
	v110 = v84
	goto L19
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F__bt_search[2]))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90+(v86^int32(-1))*int32(56))+16))
	v105 = v96
	goto L21
L23:
	;
	goto L24
L24:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F__bt_search[3]))
	v99 = int32(56)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98+v86*v99-v99)+16))
	v105 = v104
	goto L21
L25:
	;
	v117 = int32(3)
	goto L27
L26:
	;
	v117 = v34
	goto L27
L27:
	;
	if l4 == int32(3) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v118 = v117
	goto L30
L29:
	;
	v118 = v34
	goto L30
L30:
	;
	v119 = F__bt_relandgetbuf(m, l0, v112, v78<<(uint(int32(16))%32)|v81, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v119
	v33 = v119
	v34 = v118
	v35 = v110
	goto L6
L32:
	;
	F_UnlockBuffer(m, v41)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	return v35
L35:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F__bt_lockbuf(m, v131, int32(3))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v138 = F__bt_moveright(m, l0, l1, l2, v135, int32(1), v35, int32(3))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v138
	goto L34
}
func F__bt_setup_array_cmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	v19 = (v15 - int32(1)) << (uint(int32(2)) % 32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+212))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19+v21)))
	if v23 == l2 {
		v26 = F_index_getprocinfo(m, v20, v15, int32(1))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v28
			v30 = *(*int64)(unsafe.Add(mBase, uint32(v26)+16))
			*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v30
			v32 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
			*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v32
			v34 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
			*(*int64)(unsafe.Add(mBase, uint32(l3))) = v34
			if l4 == int32(0) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = l3
			}
			m.G0 = v13 - int32(-64)
			return
		}
	} else {
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v20)+208))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v40+v19)))
		v44 = F_get_opfamily_proc(m, v42, v23, l2, int32(1))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return
		} else {
			if v44 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return
				} else {
					v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v81 + int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v80
					*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v23
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(1)
					F_errmsg_internal(m, int32(_a_F__bt_setup_array_cmp_0), v13)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F__bt_setup_array_cmp_1), int32(2693), int32(_a_F__bt_setup_array_cmp_2))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
				F_fmgr_info_cxt(m, v44, l3, v48)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					if l4 == int32(0) {
						m.G0 = v13 - int32(-64)
						return
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v20)+208))
						v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v53+v54<<(uint(int32(2))%32)-int32(4))))
						v62 = F_get_opfamily_proc(m, v60, l2, l2, int32(1))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							if v62 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									v102 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v103 + int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v102
									*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(1)
									F_errmsg_internal(m, int32(_a_F__bt_setup_array_cmp_0), v11+int32(-32))
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F__bt_setup_array_cmp_1), int32(2715), int32(_a_F__bt_setup_array_cmp_2))
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
								F_fmgr_info_cxt(m, v62, v66, v67)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									m.G0 = v13 - int32(-64)
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
func F__bt_upgrademetapage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	v2 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)) = uint8(v2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(3)
	v10 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v10)
	return
}
func F_bt_index_block_validate(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+119)))
	if v10 != int32(105) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			F_errcode(m, int32(151027844))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(_a_F_bt_index_block_validate_0)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v36 + int32(4)
				F_errmsg(m, int32(_a_F_bt_index_block_validate_1), v7)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_bt_index_block_validate_2), int32(232), int32(_a_F_bt_index_block_validate_3))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+84))
		if v13 != int32(403) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(_a_F_bt_index_block_validate_0)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v36 + int32(4)
					F_errmsg(m, int32(_a_F_bt_index_block_validate_1), v7)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_bt_index_block_validate_2), int32(232), int32(_a_F_bt_index_block_validate_3))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+118)))
			if v16 == int32(116) {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
				if v19 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_bt_index_block_validate_4), int32(0))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_bt_index_block_validate_2), int32(242), int32(_a_F_bt_index_block_validate_3))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					if l1 == int64(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_bt_index_block_validate_5), int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_bt_index_block_validate_2), int32(247), int32(_a_F_bt_index_block_validate_3))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						F_check_relation_block_range(m, l0, l1)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			} else {
				if l1 == int64(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_bt_index_block_validate_5), int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_bt_index_block_validate_2), int32(247), int32(_a_F_bt_index_block_validate_3))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					F_check_relation_block_range(m, l0, l1)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_bt_multi_page_stats(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int64
	_ = v217
	var v218 int32
	_ = v218
	var v219 int64
	_ = v219
	var v220 int64
	_ = v220
	var v223 int64
	_ = v223
	var v227 int64
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v246 int64
	_ = v246
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	v10 = m.G0
	v12 = v10 - int32(272)
	m.G0 = v12
	v14 = F_superuser(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		if v14 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
			if v19 == int32(0) {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v23 = F_pg_detoast_datum_packed(m, v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
					v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
					v27 = F_init_MultiFuncCall(m, l0)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int64(0)
					} else {
						v29 = F_textToQualifiedNameList(m, v23)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int64(0)
						} else {
							v31 = F_makeRangeVarFromNameList(m, v29)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int64(0)
							} else {
								v34 = F_relation_openrv(m, v31, int32(1))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int64(0)
								} else {
									F_bt_index_block_validate(m, v34, v26)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int64(0)
									} else {
										if int64(2) <= v25 {
											F_check_relation_block_range(m, v34, v25+v26-int64(1))
											mBase = m.M
											v44 = m.ExcPending
											if v44 != 0 {
												return int64(0)
											} else {
												v45 = int32(_a_F_bt_multi_page_stats_0)
												v46 = *(*int32)(unsafe.Add(mBase, _c_F_bt_multi_page_stats[0]))
												v48 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
												*(*int32)(unsafe.Add(mBase, _c_F_bt_multi_page_stats[0])) = v48
												v51 = F_palloc(m, int32(32))
												mBase = m.M
												v52 = m.ExcPending
												if v52 != 0 {
													return int64(0)
												} else {
													v53 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
													v55 = int64(base.Ui64(v25) >> (uint(int64(63)) % 64))
													*(*uint8)(unsafe.Add(mBase, uint32(v51)+24)) = uint8(v55)
													*(*int64)(unsafe.Add(mBase, uint32(v51)+16)) = v25
													*(*int64)(unsafe.Add(mBase, uint32(v51)+8)) = v26
													*(*int32)(unsafe.Add(mBase, uint32(v51))) = v53
													*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v51
													*(*int32)(unsafe.Add(mBase, _c_F_bt_multi_page_stats[0])) = v46
													F_relation_close(m, v34, int32(0))
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return int64(0)
													} else {
														v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
														v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
														v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
														v78 = F_relation_open(m, v76, int32(0))
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int64(0)
														} else {
															v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+24)))
															if v80 == int32(0) {
																v83 = *(*int64)(unsafe.Add(mBase, uint32(v75)+16))
																v91 = v83
																if int64(0) < v91 {
																	v94 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
																	v95 = F_ReadBuffer(m, v78, v94)
																	mBase = m.M
																	v96 = m.ExcPending
																	if v96 != 0 {
																		return int64(0)
																	} else {
																		F_LockBufferInternal(m, v95, int32(1))
																		mBase = m.M
																		v99 = m.ExcPending
																		if v99 != 0 {
																			return int64(0)
																		} else {
																			*(*int64)(unsafe.Add(mBase, uint32(v12)+208)) = int64(-1)
																			v102 = int32(0)
																			*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)) = uint16(v102)
																			*(*int64)(unsafe.Add(mBase, uint32(v12)+196)) = int64(0)
																			v106 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
																			F_GetBTPageStatistics(m, v106, v95, v12+int32(176))
																			mBase = m.M
																			v110 = m.ExcPending
																			if v110 != 0 {
																				return int64(0)
																			} else {
																				F_UnlockReleaseBuffer(m, v95)
																				mBase = m.M
																				v112 = m.ExcPending
																				if v112 != 0 {
																					return int64(0)
																				} else {
																					F_relation_close(m, v78, int32(0))
																					mBase = m.M
																					v115 = m.ExcPending
																					if v115 != 0 {
																						return int64(0)
																					} else {
																						v119 = F_get_call_result_type(m, l0, int32(0), v12+int32(172))
																						mBase = m.M
																						v120 = m.ExcPending
																						if v120 != 0 {
																							return int64(0)
																						} else {
																							if v119 != int32(1) {
																								F_errstart_cold(m, int32(21), int32(0))
																								mBase = m.M
																								v270 = m.ExcPending
																								if v270 != 0 {
																									return int64(0)
																								} else {
																									F_errmsg_internal(m, int32(_a_F_bt_multi_page_stats_1), int32(0))
																									mBase = m.M
																									v274 = m.ExcPending
																									if v274 != 0 {
																										return int64(0)
																									} else {
																										F_errfinish(m, int32(_a_F_bt_multi_page_stats_2), int32(438), int32(_a_F_bt_multi_page_stats_3))
																										mBase = m.M
																										v279 = m.ExcPending
																										if v279 != 0 {
																											return int64(0)
																										} else {
																											base.Wasm_trap_unreachable()
																											for {
																											}
																										}
																									}
																								}
																							} else {
																								v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+176))
																								*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v123
																								v128 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(160))
																								mBase = m.M
																								v129 = m.ExcPending
																								if v129 != 0 {
																									return int64(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v12)+224)) = v128
																									v131 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+204)))
																									*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v131
																									v136 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_5), v12+int32(144))
																									mBase = m.M
																									v137 = m.ExcPending
																									if v137 != 0 {
																										return int64(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v12)+228)) = v136
																										v139 = *(*int32)(unsafe.Add(mBase, uint32(v12)+180))
																										*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v139
																										v144 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(128))
																										mBase = m.M
																										v145 = m.ExcPending
																										if v145 != 0 {
																											return int64(0)
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(v12)+232)) = v144
																											v147 = *(*int32)(unsafe.Add(mBase, uint32(v12)+184))
																											*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v147
																											v152 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(112))
																											mBase = m.M
																											v153 = m.ExcPending
																											if v153 != 0 {
																												return int64(0)
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(v12)+236)) = v152
																												v155 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
																												*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v155
																												v160 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(96))
																												mBase = m.M
																												v161 = m.ExcPending
																												if v161 != 0 {
																													return int64(0)
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(v12)+240)) = v160
																													v163 = *(*int32)(unsafe.Add(mBase, uint32(v12)+188))
																													*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v163
																													v168 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(80))
																													mBase = m.M
																													v169 = m.ExcPending
																													if v169 != 0 {
																														return int64(0)
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(v12)+244)) = v168
																														v171 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
																														*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v171
																														v176 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12-int32(-64))
																														mBase = m.M
																														v177 = m.ExcPending
																														if v177 != 0 {
																															return int64(0)
																														} else {
																															*(*int32)(unsafe.Add(mBase, uint32(v12)+248)) = v176
																															v179 = *(*int32)(unsafe.Add(mBase, uint32(v12)+208))
																															*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v179
																															v184 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(48))
																															mBase = m.M
																															v185 = m.ExcPending
																															if v185 != 0 {
																																return int64(0)
																															} else {
																																*(*int32)(unsafe.Add(mBase, uint32(v12)+252)) = v184
																																v187 = *(*int32)(unsafe.Add(mBase, uint32(v12)+212))
																																*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v187
																																v192 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(32))
																																mBase = m.M
																																v193 = m.ExcPending
																																if v193 != 0 {
																																	return int64(0)
																																} else {
																																	*(*int32)(unsafe.Add(mBase, uint32(v12)+256)) = v192
																																	v195 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
																																	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v195
																																	v200 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(16))
																																	mBase = m.M
																																	v201 = m.ExcPending
																																	if v201 != 0 {
																																		return int64(0)
																																	} else {
																																		*(*int32)(unsafe.Add(mBase, uint32(v12)+260)) = v200
																																		v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)))
																																		*(*int32)(unsafe.Add(mBase, uint32(v12))) = v203
																																		v206 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_6), v12)
																																		mBase = m.M
																																		v207 = m.ExcPending
																																		if v207 != 0 {
																																			return int64(0)
																																		} else {
																																			*(*int32)(unsafe.Add(mBase, uint32(v12)+264)) = v206
																																			v209 = *(*int32)(unsafe.Add(mBase, uint32(v12)+172))
																																			v210 = F_TupleDescGetAttInMetadata(m, v209)
																																			mBase = m.M
																																			v211 = m.ExcPending
																																			if v211 != 0 {
																																				return int64(0)
																																			} else {
																																				v214 = F_BuildTupleFromCStrings(m, v210, v12+int32(224))
																																				mBase = m.M
																																				v215 = m.ExcPending
																																				if v215 != 0 {
																																					return int64(0)
																																				} else {
																																					v216 = *(*int32)(unsafe.Add(mBase, uint32(v214)+16))
																																					v217 = F_HeapTupleHeaderGetDatum(m, v216)
																																					mBase = m.M
																																					v218 = m.ExcPending
																																					if v218 != 0 {
																																						return int64(0)
																																					} else {
																																						v219 = *(*int64)(unsafe.Add(mBase, uint32(v75)+8))
																																						v220 = int64(1)
																																						*(*int64)(unsafe.Add(mBase, uint32(v75)+8)) = v219 + v220
																																						v223 = *(*int64)(unsafe.Add(mBase, uint32(v75)+16))
																																						*(*int64)(unsafe.Add(mBase, uint32(v75)+16)) = v223 - v220
																																						v227 = *(*int64)(unsafe.Add(mBase, uint32(v74)))
																																						*(*int64)(unsafe.Add(mBase, uint32(v74))) = v227 + v220
																																						v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																																						*(*int32)(unsafe.Add(mBase, uint32(v231)+20)) = int32(1)
																																						v246 = v217
																																						m.G0 = v12 + int32(272)
																																						return v246
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
																					}
																				}
																			}
																		}
																	}
																} else {
																	F_relation_close(m, v78, int32(1))
																	mBase = m.M
																	v236 = m.ExcPending
																	if v236 != 0 {
																		return int64(0)
																	} else {
																		F_end_MultiFuncCall(m, l0)
																		mBase = m.M
																		v238 = m.ExcPending
																		if v238 != 0 {
																			return int64(0)
																		} else {
																			v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																			*(*int32)(unsafe.Add(mBase, uint32(v239)+20)) = int32(2)
																			v242 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v242)
																			v246 = int64(0)
																			m.G0 = v12 + int32(272)
																			return v246
																		}
																	}
																}
															} else {
																v85 = F_RelationGetNumberOfBlocksInFork(m, v78, int32(0))
																mBase = m.M
																v86 = m.ExcPending
																if v86 != 0 {
																	return int64(0)
																} else {
																	v88 = *(*int64)(unsafe.Add(mBase, uint32(v75)+8))
																	v89 = base.I64_extend_i32_u(v85) - v88
																	*(*int64)(unsafe.Add(mBase, uint32(v75)+16)) = v89
																	v91 = v89
																	if int64(0) < v91 {
																		v94 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
																		v95 = F_ReadBuffer(m, v78, v94)
																		mBase = m.M
																		v96 = m.ExcPending
																		if v96 != 0 {
																			return int64(0)
																		} else {
																			F_LockBufferInternal(m, v95, int32(1))
																			mBase = m.M
																			v99 = m.ExcPending
																			if v99 != 0 {
																				return int64(0)
																			} else {
																				*(*int64)(unsafe.Add(mBase, uint32(v12)+208)) = int64(-1)
																				v102 = int32(0)
																				*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)) = uint16(v102)
																				*(*int64)(unsafe.Add(mBase, uint32(v12)+196)) = int64(0)
																				v106 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
																				F_GetBTPageStatistics(m, v106, v95, v12+int32(176))
																				mBase = m.M
																				v110 = m.ExcPending
																				if v110 != 0 {
																					return int64(0)
																				} else {
																					F_UnlockReleaseBuffer(m, v95)
																					mBase = m.M
																					v112 = m.ExcPending
																					if v112 != 0 {
																						return int64(0)
																					} else {
																						F_relation_close(m, v78, int32(0))
																						mBase = m.M
																						v115 = m.ExcPending
																						if v115 != 0 {
																							return int64(0)
																						} else {
																							v119 = F_get_call_result_type(m, l0, int32(0), v12+int32(172))
																							mBase = m.M
																							v120 = m.ExcPending
																							if v120 != 0 {
																								return int64(0)
																							} else {
																								if v119 != int32(1) {
																									F_errstart_cold(m, int32(21), int32(0))
																									mBase = m.M
																									v270 = m.ExcPending
																									if v270 != 0 {
																										return int64(0)
																									} else {
																										F_errmsg_internal(m, int32(_a_F_bt_multi_page_stats_1), int32(0))
																										mBase = m.M
																										v274 = m.ExcPending
																										if v274 != 0 {
																											return int64(0)
																										} else {
																											F_errfinish(m, int32(_a_F_bt_multi_page_stats_2), int32(438), int32(_a_F_bt_multi_page_stats_3))
																											mBase = m.M
																											v279 = m.ExcPending
																											if v279 != 0 {
																												return int64(0)
																											} else {
																												base.Wasm_trap_unreachable()
																												for {
																												}
																											}
																										}
																									}
																								} else {
																									v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+176))
																									*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v123
																									v128 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(160))
																									mBase = m.M
																									v129 = m.ExcPending
																									if v129 != 0 {
																										return int64(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v12)+224)) = v128
																										v131 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+204)))
																										*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v131
																										v136 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_5), v12+int32(144))
																										mBase = m.M
																										v137 = m.ExcPending
																										if v137 != 0 {
																											return int64(0)
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(v12)+228)) = v136
																											v139 = *(*int32)(unsafe.Add(mBase, uint32(v12)+180))
																											*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v139
																											v144 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(128))
																											mBase = m.M
																											v145 = m.ExcPending
																											if v145 != 0 {
																												return int64(0)
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(v12)+232)) = v144
																												v147 = *(*int32)(unsafe.Add(mBase, uint32(v12)+184))
																												*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v147
																												v152 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(112))
																												mBase = m.M
																												v153 = m.ExcPending
																												if v153 != 0 {
																													return int64(0)
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(v12)+236)) = v152
																													v155 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
																													*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v155
																													v160 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(96))
																													mBase = m.M
																													v161 = m.ExcPending
																													if v161 != 0 {
																														return int64(0)
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(v12)+240)) = v160
																														v163 = *(*int32)(unsafe.Add(mBase, uint32(v12)+188))
																														*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v163
																														v168 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(80))
																														mBase = m.M
																														v169 = m.ExcPending
																														if v169 != 0 {
																															return int64(0)
																														} else {
																															*(*int32)(unsafe.Add(mBase, uint32(v12)+244)) = v168
																															v171 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
																															*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v171
																															v176 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12-int32(-64))
																															mBase = m.M
																															v177 = m.ExcPending
																															if v177 != 0 {
																																return int64(0)
																															} else {
																																*(*int32)(unsafe.Add(mBase, uint32(v12)+248)) = v176
																																v179 = *(*int32)(unsafe.Add(mBase, uint32(v12)+208))
																																*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v179
																																v184 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(48))
																																mBase = m.M
																																v185 = m.ExcPending
																																if v185 != 0 {
																																	return int64(0)
																																} else {
																																	*(*int32)(unsafe.Add(mBase, uint32(v12)+252)) = v184
																																	v187 = *(*int32)(unsafe.Add(mBase, uint32(v12)+212))
																																	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v187
																																	v192 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(32))
																																	mBase = m.M
																																	v193 = m.ExcPending
																																	if v193 != 0 {
																																		return int64(0)
																																	} else {
																																		*(*int32)(unsafe.Add(mBase, uint32(v12)+256)) = v192
																																		v195 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
																																		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v195
																																		v200 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(16))
																																		mBase = m.M
																																		v201 = m.ExcPending
																																		if v201 != 0 {
																																			return int64(0)
																																		} else {
																																			*(*int32)(unsafe.Add(mBase, uint32(v12)+260)) = v200
																																			v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)))
																																			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v203
																																			v206 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_6), v12)
																																			mBase = m.M
																																			v207 = m.ExcPending
																																			if v207 != 0 {
																																				return int64(0)
																																			} else {
																																				*(*int32)(unsafe.Add(mBase, uint32(v12)+264)) = v206
																																				v209 = *(*int32)(unsafe.Add(mBase, uint32(v12)+172))
																																				v210 = F_TupleDescGetAttInMetadata(m, v209)
																																				mBase = m.M
																																				v211 = m.ExcPending
																																				if v211 != 0 {
																																					return int64(0)
																																				} else {
																																					v214 = F_BuildTupleFromCStrings(m, v210, v12+int32(224))
																																					mBase = m.M
																																					v215 = m.ExcPending
																																					if v215 != 0 {
																																						return int64(0)
																																					} else {
																																						v216 = *(*int32)(unsafe.Add(mBase, uint32(v214)+16))
																																						v217 = F_HeapTupleHeaderGetDatum(m, v216)
																																						mBase = m.M
																																						v218 = m.ExcPending
																																						if v218 != 0 {
																																							return int64(0)
																																						} else {
																																							v219 = *(*int64)(unsafe.Add(mBase, uint32(v75)+8))
																																							v220 = int64(1)
																																							*(*int64)(unsafe.Add(mBase, uint32(v75)+8)) = v219 + v220
																																							v223 = *(*int64)(unsafe.Add(mBase, uint32(v75)+16))
																																							*(*int64)(unsafe.Add(mBase, uint32(v75)+16)) = v223 - v220
																																							v227 = *(*int64)(unsafe.Add(mBase, uint32(v74)))
																																							*(*int64)(unsafe.Add(mBase, uint32(v74))) = v227 + v220
																																							v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																																							*(*int32)(unsafe.Add(mBase, uint32(v231)+20)) = int32(1)
																																							v246 = v217
																																							m.G0 = v12 + int32(272)
																																							return v246
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
																						}
																					}
																				}
																			}
																		}
																	} else {
																		F_relation_close(m, v78, int32(1))
																		mBase = m.M
																		v236 = m.ExcPending
																		if v236 != 0 {
																			return int64(0)
																		} else {
																			F_end_MultiFuncCall(m, l0)
																			mBase = m.M
																			v238 = m.ExcPending
																			if v238 != 0 {
																				return int64(0)
																			} else {
																				v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																				*(*int32)(unsafe.Add(mBase, uint32(v239)+20)) = int32(2)
																				v242 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v242)
																				v246 = int64(0)
																				m.G0 = v12 + int32(272)
																				return v246
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
											v45 = int32(_a_F_bt_multi_page_stats_0)
											v46 = *(*int32)(unsafe.Add(mBase, _c_F_bt_multi_page_stats[0]))
											v48 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
											*(*int32)(unsafe.Add(mBase, _c_F_bt_multi_page_stats[0])) = v48
											v51 = F_palloc(m, int32(32))
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return int64(0)
											} else {
												v53 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
												v55 = int64(base.Ui64(v25) >> (uint(int64(63)) % 64))
												*(*uint8)(unsafe.Add(mBase, uint32(v51)+24)) = uint8(v55)
												*(*int64)(unsafe.Add(mBase, uint32(v51)+16)) = v25
												*(*int64)(unsafe.Add(mBase, uint32(v51)+8)) = v26
												*(*int32)(unsafe.Add(mBase, uint32(v51))) = v53
												*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v51
												*(*int32)(unsafe.Add(mBase, _c_F_bt_multi_page_stats[0])) = v46
												F_relation_close(m, v34, int32(0))
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int64(0)
												} else {
													v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
													v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
													v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
													v78 = F_relation_open(m, v76, int32(0))
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int64(0)
													} else {
														v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+24)))
														if v80 == int32(0) {
															v83 = *(*int64)(unsafe.Add(mBase, uint32(v75)+16))
															v91 = v83
															if int64(0) < v91 {
																v94 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
																v95 = F_ReadBuffer(m, v78, v94)
																mBase = m.M
																v96 = m.ExcPending
																if v96 != 0 {
																	return int64(0)
																} else {
																	F_LockBufferInternal(m, v95, int32(1))
																	mBase = m.M
																	v99 = m.ExcPending
																	if v99 != 0 {
																		return int64(0)
																	} else {
																		*(*int64)(unsafe.Add(mBase, uint32(v12)+208)) = int64(-1)
																		v102 = int32(0)
																		*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)) = uint16(v102)
																		*(*int64)(unsafe.Add(mBase, uint32(v12)+196)) = int64(0)
																		v106 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
																		F_GetBTPageStatistics(m, v106, v95, v12+int32(176))
																		mBase = m.M
																		v110 = m.ExcPending
																		if v110 != 0 {
																			return int64(0)
																		} else {
																			F_UnlockReleaseBuffer(m, v95)
																			mBase = m.M
																			v112 = m.ExcPending
																			if v112 != 0 {
																				return int64(0)
																			} else {
																				F_relation_close(m, v78, int32(0))
																				mBase = m.M
																				v115 = m.ExcPending
																				if v115 != 0 {
																					return int64(0)
																				} else {
																					v119 = F_get_call_result_type(m, l0, int32(0), v12+int32(172))
																					mBase = m.M
																					v120 = m.ExcPending
																					if v120 != 0 {
																						return int64(0)
																					} else {
																						if v119 != int32(1) {
																							F_errstart_cold(m, int32(21), int32(0))
																							mBase = m.M
																							v270 = m.ExcPending
																							if v270 != 0 {
																								return int64(0)
																							} else {
																								F_errmsg_internal(m, int32(_a_F_bt_multi_page_stats_1), int32(0))
																								mBase = m.M
																								v274 = m.ExcPending
																								if v274 != 0 {
																									return int64(0)
																								} else {
																									F_errfinish(m, int32(_a_F_bt_multi_page_stats_2), int32(438), int32(_a_F_bt_multi_page_stats_3))
																									mBase = m.M
																									v279 = m.ExcPending
																									if v279 != 0 {
																										return int64(0)
																									} else {
																										base.Wasm_trap_unreachable()
																										for {
																										}
																									}
																								}
																							}
																						} else {
																							v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+176))
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v123
																							v128 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(160))
																							mBase = m.M
																							v129 = m.ExcPending
																							if v129 != 0 {
																								return int64(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v12)+224)) = v128
																								v131 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+204)))
																								*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v131
																								v136 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_5), v12+int32(144))
																								mBase = m.M
																								v137 = m.ExcPending
																								if v137 != 0 {
																									return int64(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v12)+228)) = v136
																									v139 = *(*int32)(unsafe.Add(mBase, uint32(v12)+180))
																									*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v139
																									v144 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(128))
																									mBase = m.M
																									v145 = m.ExcPending
																									if v145 != 0 {
																										return int64(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v12)+232)) = v144
																										v147 = *(*int32)(unsafe.Add(mBase, uint32(v12)+184))
																										*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v147
																										v152 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(112))
																										mBase = m.M
																										v153 = m.ExcPending
																										if v153 != 0 {
																											return int64(0)
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(v12)+236)) = v152
																											v155 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
																											*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v155
																											v160 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(96))
																											mBase = m.M
																											v161 = m.ExcPending
																											if v161 != 0 {
																												return int64(0)
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(v12)+240)) = v160
																												v163 = *(*int32)(unsafe.Add(mBase, uint32(v12)+188))
																												*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v163
																												v168 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(80))
																												mBase = m.M
																												v169 = m.ExcPending
																												if v169 != 0 {
																													return int64(0)
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(v12)+244)) = v168
																													v171 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
																													*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v171
																													v176 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12-int32(-64))
																													mBase = m.M
																													v177 = m.ExcPending
																													if v177 != 0 {
																														return int64(0)
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(v12)+248)) = v176
																														v179 = *(*int32)(unsafe.Add(mBase, uint32(v12)+208))
																														*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v179
																														v184 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(48))
																														mBase = m.M
																														v185 = m.ExcPending
																														if v185 != 0 {
																															return int64(0)
																														} else {
																															*(*int32)(unsafe.Add(mBase, uint32(v12)+252)) = v184
																															v187 = *(*int32)(unsafe.Add(mBase, uint32(v12)+212))
																															*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v187
																															v192 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(32))
																															mBase = m.M
																															v193 = m.ExcPending
																															if v193 != 0 {
																																return int64(0)
																															} else {
																																*(*int32)(unsafe.Add(mBase, uint32(v12)+256)) = v192
																																v195 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
																																*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v195
																																v200 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(16))
																																mBase = m.M
																																v201 = m.ExcPending
																																if v201 != 0 {
																																	return int64(0)
																																} else {
																																	*(*int32)(unsafe.Add(mBase, uint32(v12)+260)) = v200
																																	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)))
																																	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v203
																																	v206 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_6), v12)
																																	mBase = m.M
																																	v207 = m.ExcPending
																																	if v207 != 0 {
																																		return int64(0)
																																	} else {
																																		*(*int32)(unsafe.Add(mBase, uint32(v12)+264)) = v206
																																		v209 = *(*int32)(unsafe.Add(mBase, uint32(v12)+172))
																																		v210 = F_TupleDescGetAttInMetadata(m, v209)
																																		mBase = m.M
																																		v211 = m.ExcPending
																																		if v211 != 0 {
																																			return int64(0)
																																		} else {
																																			v214 = F_BuildTupleFromCStrings(m, v210, v12+int32(224))
																																			mBase = m.M
																																			v215 = m.ExcPending
																																			if v215 != 0 {
																																				return int64(0)
																																			} else {
																																				v216 = *(*int32)(unsafe.Add(mBase, uint32(v214)+16))
																																				v217 = F_HeapTupleHeaderGetDatum(m, v216)
																																				mBase = m.M
																																				v218 = m.ExcPending
																																				if v218 != 0 {
																																					return int64(0)
																																				} else {
																																					v219 = *(*int64)(unsafe.Add(mBase, uint32(v75)+8))
																																					v220 = int64(1)
																																					*(*int64)(unsafe.Add(mBase, uint32(v75)+8)) = v219 + v220
																																					v223 = *(*int64)(unsafe.Add(mBase, uint32(v75)+16))
																																					*(*int64)(unsafe.Add(mBase, uint32(v75)+16)) = v223 - v220
																																					v227 = *(*int64)(unsafe.Add(mBase, uint32(v74)))
																																					*(*int64)(unsafe.Add(mBase, uint32(v74))) = v227 + v220
																																					v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																																					*(*int32)(unsafe.Add(mBase, uint32(v231)+20)) = int32(1)
																																					v246 = v217
																																					m.G0 = v12 + int32(272)
																																					return v246
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
																				}
																			}
																		}
																	}
																}
															} else {
																F_relation_close(m, v78, int32(1))
																mBase = m.M
																v236 = m.ExcPending
																if v236 != 0 {
																	return int64(0)
																} else {
																	F_end_MultiFuncCall(m, l0)
																	mBase = m.M
																	v238 = m.ExcPending
																	if v238 != 0 {
																		return int64(0)
																	} else {
																		v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		*(*int32)(unsafe.Add(mBase, uint32(v239)+20)) = int32(2)
																		v242 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v242)
																		v246 = int64(0)
																		m.G0 = v12 + int32(272)
																		return v246
																	}
																}
															}
														} else {
															v85 = F_RelationGetNumberOfBlocksInFork(m, v78, int32(0))
															mBase = m.M
															v86 = m.ExcPending
															if v86 != 0 {
																return int64(0)
															} else {
																v88 = *(*int64)(unsafe.Add(mBase, uint32(v75)+8))
																v89 = base.I64_extend_i32_u(v85) - v88
																*(*int64)(unsafe.Add(mBase, uint32(v75)+16)) = v89
																v91 = v89
																if int64(0) < v91 {
																	v94 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
																	v95 = F_ReadBuffer(m, v78, v94)
																	mBase = m.M
																	v96 = m.ExcPending
																	if v96 != 0 {
																		return int64(0)
																	} else {
																		F_LockBufferInternal(m, v95, int32(1))
																		mBase = m.M
																		v99 = m.ExcPending
																		if v99 != 0 {
																			return int64(0)
																		} else {
																			*(*int64)(unsafe.Add(mBase, uint32(v12)+208)) = int64(-1)
																			v102 = int32(0)
																			*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)) = uint16(v102)
																			*(*int64)(unsafe.Add(mBase, uint32(v12)+196)) = int64(0)
																			v106 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
																			F_GetBTPageStatistics(m, v106, v95, v12+int32(176))
																			mBase = m.M
																			v110 = m.ExcPending
																			if v110 != 0 {
																				return int64(0)
																			} else {
																				F_UnlockReleaseBuffer(m, v95)
																				mBase = m.M
																				v112 = m.ExcPending
																				if v112 != 0 {
																					return int64(0)
																				} else {
																					F_relation_close(m, v78, int32(0))
																					mBase = m.M
																					v115 = m.ExcPending
																					if v115 != 0 {
																						return int64(0)
																					} else {
																						v119 = F_get_call_result_type(m, l0, int32(0), v12+int32(172))
																						mBase = m.M
																						v120 = m.ExcPending
																						if v120 != 0 {
																							return int64(0)
																						} else {
																							if v119 != int32(1) {
																								F_errstart_cold(m, int32(21), int32(0))
																								mBase = m.M
																								v270 = m.ExcPending
																								if v270 != 0 {
																									return int64(0)
																								} else {
																									F_errmsg_internal(m, int32(_a_F_bt_multi_page_stats_1), int32(0))
																									mBase = m.M
																									v274 = m.ExcPending
																									if v274 != 0 {
																										return int64(0)
																									} else {
																										F_errfinish(m, int32(_a_F_bt_multi_page_stats_2), int32(438), int32(_a_F_bt_multi_page_stats_3))
																										mBase = m.M
																										v279 = m.ExcPending
																										if v279 != 0 {
																											return int64(0)
																										} else {
																											base.Wasm_trap_unreachable()
																											for {
																											}
																										}
																									}
																								}
																							} else {
																								v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+176))
																								*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v123
																								v128 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(160))
																								mBase = m.M
																								v129 = m.ExcPending
																								if v129 != 0 {
																									return int64(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v12)+224)) = v128
																									v131 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+204)))
																									*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v131
																									v136 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_5), v12+int32(144))
																									mBase = m.M
																									v137 = m.ExcPending
																									if v137 != 0 {
																										return int64(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v12)+228)) = v136
																										v139 = *(*int32)(unsafe.Add(mBase, uint32(v12)+180))
																										*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v139
																										v144 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(128))
																										mBase = m.M
																										v145 = m.ExcPending
																										if v145 != 0 {
																											return int64(0)
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(v12)+232)) = v144
																											v147 = *(*int32)(unsafe.Add(mBase, uint32(v12)+184))
																											*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v147
																											v152 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(112))
																											mBase = m.M
																											v153 = m.ExcPending
																											if v153 != 0 {
																												return int64(0)
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(v12)+236)) = v152
																												v155 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
																												*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v155
																												v160 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(96))
																												mBase = m.M
																												v161 = m.ExcPending
																												if v161 != 0 {
																													return int64(0)
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(v12)+240)) = v160
																													v163 = *(*int32)(unsafe.Add(mBase, uint32(v12)+188))
																													*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v163
																													v168 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(80))
																													mBase = m.M
																													v169 = m.ExcPending
																													if v169 != 0 {
																														return int64(0)
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(v12)+244)) = v168
																														v171 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
																														*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v171
																														v176 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12-int32(-64))
																														mBase = m.M
																														v177 = m.ExcPending
																														if v177 != 0 {
																															return int64(0)
																														} else {
																															*(*int32)(unsafe.Add(mBase, uint32(v12)+248)) = v176
																															v179 = *(*int32)(unsafe.Add(mBase, uint32(v12)+208))
																															*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v179
																															v184 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(48))
																															mBase = m.M
																															v185 = m.ExcPending
																															if v185 != 0 {
																																return int64(0)
																															} else {
																																*(*int32)(unsafe.Add(mBase, uint32(v12)+252)) = v184
																																v187 = *(*int32)(unsafe.Add(mBase, uint32(v12)+212))
																																*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v187
																																v192 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(32))
																																mBase = m.M
																																v193 = m.ExcPending
																																if v193 != 0 {
																																	return int64(0)
																																} else {
																																	*(*int32)(unsafe.Add(mBase, uint32(v12)+256)) = v192
																																	v195 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
																																	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v195
																																	v200 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(16))
																																	mBase = m.M
																																	v201 = m.ExcPending
																																	if v201 != 0 {
																																		return int64(0)
																																	} else {
																																		*(*int32)(unsafe.Add(mBase, uint32(v12)+260)) = v200
																																		v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)))
																																		*(*int32)(unsafe.Add(mBase, uint32(v12))) = v203
																																		v206 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_6), v12)
																																		mBase = m.M
																																		v207 = m.ExcPending
																																		if v207 != 0 {
																																			return int64(0)
																																		} else {
																																			*(*int32)(unsafe.Add(mBase, uint32(v12)+264)) = v206
																																			v209 = *(*int32)(unsafe.Add(mBase, uint32(v12)+172))
																																			v210 = F_TupleDescGetAttInMetadata(m, v209)
																																			mBase = m.M
																																			v211 = m.ExcPending
																																			if v211 != 0 {
																																				return int64(0)
																																			} else {
																																				v214 = F_BuildTupleFromCStrings(m, v210, v12+int32(224))
																																				mBase = m.M
																																				v215 = m.ExcPending
																																				if v215 != 0 {
																																					return int64(0)
																																				} else {
																																					v216 = *(*int32)(unsafe.Add(mBase, uint32(v214)+16))
																																					v217 = F_HeapTupleHeaderGetDatum(m, v216)
																																					mBase = m.M
																																					v218 = m.ExcPending
																																					if v218 != 0 {
																																						return int64(0)
																																					} else {
																																						v219 = *(*int64)(unsafe.Add(mBase, uint32(v75)+8))
																																						v220 = int64(1)
																																						*(*int64)(unsafe.Add(mBase, uint32(v75)+8)) = v219 + v220
																																						v223 = *(*int64)(unsafe.Add(mBase, uint32(v75)+16))
																																						*(*int64)(unsafe.Add(mBase, uint32(v75)+16)) = v223 - v220
																																						v227 = *(*int64)(unsafe.Add(mBase, uint32(v74)))
																																						*(*int64)(unsafe.Add(mBase, uint32(v74))) = v227 + v220
																																						v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																																						*(*int32)(unsafe.Add(mBase, uint32(v231)+20)) = int32(1)
																																						v246 = v217
																																						m.G0 = v12 + int32(272)
																																						return v246
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
																					}
																				}
																			}
																		}
																	}
																} else {
																	F_relation_close(m, v78, int32(1))
																	mBase = m.M
																	v236 = m.ExcPending
																	if v236 != 0 {
																		return int64(0)
																	} else {
																		F_end_MultiFuncCall(m, l0)
																		mBase = m.M
																		v238 = m.ExcPending
																		if v238 != 0 {
																			return int64(0)
																		} else {
																			v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																			*(*int32)(unsafe.Add(mBase, uint32(v239)+20)) = int32(2)
																			v242 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v242)
																			v246 = int64(0)
																			m.G0 = v12 + int32(272)
																			return v246
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
				v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
				v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
				v78 = F_relation_open(m, v76, int32(0))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int64(0)
				} else {
					v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+24)))
					if v80 == int32(0) {
						v83 = *(*int64)(unsafe.Add(mBase, uint32(v75)+16))
						v91 = v83
						if int64(0) < v91 {
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
							v95 = F_ReadBuffer(m, v78, v94)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int64(0)
							} else {
								F_LockBufferInternal(m, v95, int32(1))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int64(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v12)+208)) = int64(-1)
									v102 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)) = uint16(v102)
									*(*int64)(unsafe.Add(mBase, uint32(v12)+196)) = int64(0)
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
									F_GetBTPageStatistics(m, v106, v95, v12+int32(176))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int64(0)
									} else {
										F_UnlockReleaseBuffer(m, v95)
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return int64(0)
										} else {
											F_relation_close(m, v78, int32(0))
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int64(0)
											} else {
												v119 = F_get_call_result_type(m, l0, int32(0), v12+int32(172))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return int64(0)
												} else {
													if v119 != int32(1) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v270 = m.ExcPending
														if v270 != 0 {
															return int64(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_bt_multi_page_stats_1), int32(0))
															mBase = m.M
															v274 = m.ExcPending
															if v274 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_bt_multi_page_stats_2), int32(438), int32(_a_F_bt_multi_page_stats_3))
																mBase = m.M
																v279 = m.ExcPending
																if v279 != 0 {
																	return int64(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+176))
														*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v123
														v128 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(160))
														mBase = m.M
														v129 = m.ExcPending
														if v129 != 0 {
															return int64(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v12)+224)) = v128
															v131 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+204)))
															*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v131
															v136 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_5), v12+int32(144))
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return int64(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v12)+228)) = v136
																v139 = *(*int32)(unsafe.Add(mBase, uint32(v12)+180))
																*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v139
																v144 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(128))
																mBase = m.M
																v145 = m.ExcPending
																if v145 != 0 {
																	return int64(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v12)+232)) = v144
																	v147 = *(*int32)(unsafe.Add(mBase, uint32(v12)+184))
																	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v147
																	v152 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(112))
																	mBase = m.M
																	v153 = m.ExcPending
																	if v153 != 0 {
																		return int64(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+236)) = v152
																		v155 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v155
																		v160 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(96))
																		mBase = m.M
																		v161 = m.ExcPending
																		if v161 != 0 {
																			return int64(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v12)+240)) = v160
																			v163 = *(*int32)(unsafe.Add(mBase, uint32(v12)+188))
																			*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v163
																			v168 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(80))
																			mBase = m.M
																			v169 = m.ExcPending
																			if v169 != 0 {
																				return int64(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+244)) = v168
																				v171 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v171
																				v176 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12-int32(-64))
																				mBase = m.M
																				v177 = m.ExcPending
																				if v177 != 0 {
																					return int64(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+248)) = v176
																					v179 = *(*int32)(unsafe.Add(mBase, uint32(v12)+208))
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v179
																					v184 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(48))
																					mBase = m.M
																					v185 = m.ExcPending
																					if v185 != 0 {
																						return int64(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v12)+252)) = v184
																						v187 = *(*int32)(unsafe.Add(mBase, uint32(v12)+212))
																						*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v187
																						v192 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(32))
																						mBase = m.M
																						v193 = m.ExcPending
																						if v193 != 0 {
																							return int64(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+256)) = v192
																							v195 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v195
																							v200 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(16))
																							mBase = m.M
																							v201 = m.ExcPending
																							if v201 != 0 {
																								return int64(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v12)+260)) = v200
																								v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)))
																								*(*int32)(unsafe.Add(mBase, uint32(v12))) = v203
																								v206 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_6), v12)
																								mBase = m.M
																								v207 = m.ExcPending
																								if v207 != 0 {
																									return int64(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v12)+264)) = v206
																									v209 = *(*int32)(unsafe.Add(mBase, uint32(v12)+172))
																									v210 = F_TupleDescGetAttInMetadata(m, v209)
																									mBase = m.M
																									v211 = m.ExcPending
																									if v211 != 0 {
																										return int64(0)
																									} else {
																										v214 = F_BuildTupleFromCStrings(m, v210, v12+int32(224))
																										mBase = m.M
																										v215 = m.ExcPending
																										if v215 != 0 {
																											return int64(0)
																										} else {
																											v216 = *(*int32)(unsafe.Add(mBase, uint32(v214)+16))
																											v217 = F_HeapTupleHeaderGetDatum(m, v216)
																											mBase = m.M
																											v218 = m.ExcPending
																											if v218 != 0 {
																												return int64(0)
																											} else {
																												v219 = *(*int64)(unsafe.Add(mBase, uint32(v75)+8))
																												v220 = int64(1)
																												*(*int64)(unsafe.Add(mBase, uint32(v75)+8)) = v219 + v220
																												v223 = *(*int64)(unsafe.Add(mBase, uint32(v75)+16))
																												*(*int64)(unsafe.Add(mBase, uint32(v75)+16)) = v223 - v220
																												v227 = *(*int64)(unsafe.Add(mBase, uint32(v74)))
																												*(*int64)(unsafe.Add(mBase, uint32(v74))) = v227 + v220
																												v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																												*(*int32)(unsafe.Add(mBase, uint32(v231)+20)) = int32(1)
																												v246 = v217
																												m.G0 = v12 + int32(272)
																												return v246
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
											}
										}
									}
								}
							}
						} else {
							F_relation_close(m, v78, int32(1))
							mBase = m.M
							v236 = m.ExcPending
							if v236 != 0 {
								return int64(0)
							} else {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v238 = m.ExcPending
								if v238 != 0 {
									return int64(0)
								} else {
									v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v239)+20)) = int32(2)
									v242 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v242)
									v246 = int64(0)
									m.G0 = v12 + int32(272)
									return v246
								}
							}
						}
					} else {
						v85 = F_RelationGetNumberOfBlocksInFork(m, v78, int32(0))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int64(0)
						} else {
							v88 = *(*int64)(unsafe.Add(mBase, uint32(v75)+8))
							v89 = base.I64_extend_i32_u(v85) - v88
							*(*int64)(unsafe.Add(mBase, uint32(v75)+16)) = v89
							v91 = v89
							if int64(0) < v91 {
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
								v95 = F_ReadBuffer(m, v78, v94)
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int64(0)
								} else {
									F_LockBufferInternal(m, v95, int32(1))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v12)+208)) = int64(-1)
										v102 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)) = uint16(v102)
										*(*int64)(unsafe.Add(mBase, uint32(v12)+196)) = int64(0)
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
										F_GetBTPageStatistics(m, v106, v95, v12+int32(176))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int64(0)
										} else {
											F_UnlockReleaseBuffer(m, v95)
											mBase = m.M
											v112 = m.ExcPending
											if v112 != 0 {
												return int64(0)
											} else {
												F_relation_close(m, v78, int32(0))
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int64(0)
												} else {
													v119 = F_get_call_result_type(m, l0, int32(0), v12+int32(172))
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return int64(0)
													} else {
														if v119 != int32(1) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v270 = m.ExcPending
															if v270 != 0 {
																return int64(0)
															} else {
																F_errmsg_internal(m, int32(_a_F_bt_multi_page_stats_1), int32(0))
																mBase = m.M
																v274 = m.ExcPending
																if v274 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_bt_multi_page_stats_2), int32(438), int32(_a_F_bt_multi_page_stats_3))
																	mBase = m.M
																	v279 = m.ExcPending
																	if v279 != 0 {
																		return int64(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+176))
															*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v123
															v128 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(160))
															mBase = m.M
															v129 = m.ExcPending
															if v129 != 0 {
																return int64(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v12)+224)) = v128
																v131 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+204)))
																*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v131
																v136 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_5), v12+int32(144))
																mBase = m.M
																v137 = m.ExcPending
																if v137 != 0 {
																	return int64(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v12)+228)) = v136
																	v139 = *(*int32)(unsafe.Add(mBase, uint32(v12)+180))
																	*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v139
																	v144 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(128))
																	mBase = m.M
																	v145 = m.ExcPending
																	if v145 != 0 {
																		return int64(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+232)) = v144
																		v147 = *(*int32)(unsafe.Add(mBase, uint32(v12)+184))
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v147
																		v152 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(112))
																		mBase = m.M
																		v153 = m.ExcPending
																		if v153 != 0 {
																			return int64(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v12)+236)) = v152
																			v155 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
																			*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v155
																			v160 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(96))
																			mBase = m.M
																			v161 = m.ExcPending
																			if v161 != 0 {
																				return int64(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+240)) = v160
																				v163 = *(*int32)(unsafe.Add(mBase, uint32(v12)+188))
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v163
																				v168 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(80))
																				mBase = m.M
																				v169 = m.ExcPending
																				if v169 != 0 {
																					return int64(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+244)) = v168
																					v171 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v171
																					v176 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12-int32(-64))
																					mBase = m.M
																					v177 = m.ExcPending
																					if v177 != 0 {
																						return int64(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v12)+248)) = v176
																						v179 = *(*int32)(unsafe.Add(mBase, uint32(v12)+208))
																						*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v179
																						v184 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(48))
																						mBase = m.M
																						v185 = m.ExcPending
																						if v185 != 0 {
																							return int64(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+252)) = v184
																							v187 = *(*int32)(unsafe.Add(mBase, uint32(v12)+212))
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v187
																							v192 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(32))
																							mBase = m.M
																							v193 = m.ExcPending
																							if v193 != 0 {
																								return int64(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v12)+256)) = v192
																								v195 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
																								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v195
																								v200 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(16))
																								mBase = m.M
																								v201 = m.ExcPending
																								if v201 != 0 {
																									return int64(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v12)+260)) = v200
																									v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)))
																									*(*int32)(unsafe.Add(mBase, uint32(v12))) = v203
																									v206 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_6), v12)
																									mBase = m.M
																									v207 = m.ExcPending
																									if v207 != 0 {
																										return int64(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v12)+264)) = v206
																										v209 = *(*int32)(unsafe.Add(mBase, uint32(v12)+172))
																										v210 = F_TupleDescGetAttInMetadata(m, v209)
																										mBase = m.M
																										v211 = m.ExcPending
																										if v211 != 0 {
																											return int64(0)
																										} else {
																											v214 = F_BuildTupleFromCStrings(m, v210, v12+int32(224))
																											mBase = m.M
																											v215 = m.ExcPending
																											if v215 != 0 {
																												return int64(0)
																											} else {
																												v216 = *(*int32)(unsafe.Add(mBase, uint32(v214)+16))
																												v217 = F_HeapTupleHeaderGetDatum(m, v216)
																												mBase = m.M
																												v218 = m.ExcPending
																												if v218 != 0 {
																													return int64(0)
																												} else {
																													v219 = *(*int64)(unsafe.Add(mBase, uint32(v75)+8))
																													v220 = int64(1)
																													*(*int64)(unsafe.Add(mBase, uint32(v75)+8)) = v219 + v220
																													v223 = *(*int64)(unsafe.Add(mBase, uint32(v75)+16))
																													*(*int64)(unsafe.Add(mBase, uint32(v75)+16)) = v223 - v220
																													v227 = *(*int64)(unsafe.Add(mBase, uint32(v74)))
																													*(*int64)(unsafe.Add(mBase, uint32(v74))) = v227 + v220
																													v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																													*(*int32)(unsafe.Add(mBase, uint32(v231)+20)) = int32(1)
																													v246 = v217
																													m.G0 = v12 + int32(272)
																													return v246
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
												}
											}
										}
									}
								}
							} else {
								F_relation_close(m, v78, int32(1))
								mBase = m.M
								v236 = m.ExcPending
								if v236 != 0 {
									return int64(0)
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v238 = m.ExcPending
									if v238 != 0 {
										return int64(0)
									} else {
										v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v239)+20)) = int32(2)
										v242 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v242)
										v246 = int64(0)
										m.G0 = v12 + int32(272)
										return v246
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v254 = m.ExcPending
			if v254 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v257 = m.ExcPending
				if v257 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_bt_multi_page_stats_7), int32(0))
					mBase = m.M
					v261 = m.ExcPending
					if v261 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_bt_multi_page_stats_2), int32(354), int32(_a_F_bt_multi_page_stats_3))
						mBase = m.M
						v266 = m.ExcPending
						if v266 != 0 {
							return int64(0)
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
}
