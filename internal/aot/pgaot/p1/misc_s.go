package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_SB_do_like_escape(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v11 = int32(1)
	v12 = v10 & v11
	if v10 == v11 {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v18 == int32(18) {
			v21 = int32(16)
		} else {
			v21 = int32(0)
		}
		if base.Ui32((v18-int32(1))&int32(255)) < base.Ui32(int32(3)) {
			v28 = int32(4)
		} else {
			v28 = v21
		}
		v39 = v28
	} else {
		v29 = int32(1)
		if v12 != 0 {
			v39 = int32(base.Ui32(v10)>>(uint(v29)%32)) - v29
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v39 = int32(base.Ui32(v33)>>(uint(int32(2))%32)) - int32(4)
		}
	}
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v40 == int32(1) {
		v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		if base.Ui32((v43-int32(1))&int32(255)) < base.Ui32(int32(3)) {
			v154 = F_palloc(m, v39<<(uint(int32(1))%32)+int32(4))
			mBase = m.M
			v155 = m.ExcPending
			if v155 != 0 {
				return int32(0)
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v163 = m.ExcPending
				if v163 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(84410498))
					mBase = m.M
					v166 = m.ExcPending
					if v166 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_SB_do_like_escape_0), int32(0))
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(_a_F_SB_do_like_escape_1), int32(0))
							mBase = m.M
							v174 = m.ExcPending
							if v174 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_SB_do_like_escape_2), int32(453), int32(_a_F_SB_do_like_escape_3))
								mBase = m.M
								v179 = m.ExcPending
								if v179 != 0 {
									return int32(0)
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
		} else {
			if v43 == int32(18) {
				v54 = int32(16)
			} else {
				v54 = int32(0)
			}
			v67 = v54
			if v12 != 0 {
				v70 = int32(1)
			} else {
				v70 = int32(4)
			}
			v71 = l0 + v70
			v76 = F_palloc(m, v39<<(uint(int32(1))%32)+int32(4))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				v81 = v76 + int32(4)
				switch v67 {
				case 0:
					if v39 <= int32(0) {
						v234 = v81
					} else {
						if v39&int32(1) != 0 {
							v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
							if v86 == int32(92) {
								v89 = int32(92)
								*(*uint8)(unsafe.Add(mBase, uint32(v76)+4)) = uint8(v89)
								v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
								v94 = v93
								v95 = v76 + int32(5)
							} else {
								v94 = v86
								v95 = v81
							}
							*(*uint8)(unsafe.Add(mBase, uint32(v95))) = uint8(v94)
							v97 = int32(1)
							v104 = v95 + v97
							v105 = v71 + v97
							v106 = v39 - v97
						} else {
							v104 = v81
							v105 = v71
							v106 = v39
						}
						if v39 == int32(1) {
							v234 = v104
						} else {
							v109 = v106
							v111 = v104
							v113 = v105
							for {
								v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
								if v118 == int32(92) {
									v121 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v111))) = uint8(v121)
									v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
									v126 = v125
									v127 = v111 + int32(1)
								} else {
									v126 = v118
									v127 = v111
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v127))) = uint8(v126)
								v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)))
								if v129 != int32(92) {
									v139 = v129
									v140 = v127 + int32(1)
								} else {
									v134 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)) = uint8(v134)
									v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)))
									v139 = v136
									v140 = v127 + int32(2)
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v139)
								v143 = v140 + int32(1)
								v144 = int32(2)
								if v144 < v109 {
									v109 = v109 - v144
									v111 = v143
									v113 = v113 + v144
									continue
								} else {
									break
								}
								break
							}
							v234 = v143
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v76))) = (v234 - v76) << (uint(int32(2)) % 32)
					return v76
				case 1:
					v180 = int32(1)
					v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
					if v182&v180 != 0 {
						v185 = v180
					} else {
						v185 = int32(4)
					}
					v186 = l1 + v185
					v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
					if v187 == int32(92) {
						v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
						if v245 == int32(1) {
							v249 = int32(18)
							v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							if v251 == v249 {
								v254 = v249
							} else {
								v254 = int32(2)
							}
							if base.Ui32((v251-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v261 = int32(6)
							} else {
								v261 = v254
							}
							v270 = v261
						} else {
							v262 = int32(1)
							if v245&v262 != 0 {
								v270 = int32(base.Ui32(v245) >> (uint(v262) % 32))
							} else {
								v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v270 = int32(base.Ui32(v266) >> (uint(int32(2)) % 32))
							}
						}
						if v270 == int32(0) {
							return v76
						} else {
							base.MemoryCopy(m, v76, l0, v270)
							return v76
						}
					} else {
						v190 = int32(0)
						if v39 <= v190 {
							v234 = v81
						} else {
							v194 = v190
							v195 = v81
							v196 = v39
							v197 = v71
							for {
								v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
								v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
								v207 = base.B2i32(v194 == int32(0)) & base.B2i32(v204 == v205)
								if v207 != 0 {
									v208 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v208)
									v225 = v195 + int32(1)
								} else {
									v213 = v195 + int32(1)
									if v204 == int32(92) {
										v216 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v216)
										if v194&int32(1) != 0 {
											v225 = v213
										} else {
											v220 = int32(92)
											*(*uint8)(unsafe.Add(mBase, uint32(v195)+1)) = uint8(v220)
											v225 = v195 + int32(2)
										}
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v204)
										v225 = v213
									}
								}
								v226 = int32(1)
								if v226 < v196 {
									v194 = v207
									v195 = v225
									v196 = v196 - v226
									v197 = v197 + v226
									continue
								} else {
									break
								}
								break
							}
							v234 = v225
						}
						*(*int32)(unsafe.Add(mBase, uint32(v76))) = (v234 - v76) << (uint(int32(2)) % 32)
						return v76
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(84410498))
						mBase = m.M
						v166 = m.ExcPending
						if v166 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_SB_do_like_escape_0), int32(0))
							mBase = m.M
							v170 = m.ExcPending
							if v170 != 0 {
								return int32(0)
							} else {
								F_errhint(m, int32(_a_F_SB_do_like_escape_1), int32(0))
								mBase = m.M
								v174 = m.ExcPending
								if v174 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_SB_do_like_escape_2), int32(453), int32(_a_F_SB_do_like_escape_3))
									mBase = m.M
									v179 = m.ExcPending
									if v179 != 0 {
										return int32(0)
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
		v55 = int32(1)
		if v40&v55 != 0 {
			v67 = int32(base.Ui32(v40)>>(uint(v55)%32)) - v55
		} else {
			v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v67 = int32(base.Ui32(v61)>>(uint(int32(2))%32)) - int32(4)
		}
		if v12 != 0 {
			v70 = int32(1)
		} else {
			v70 = int32(4)
		}
		v71 = l0 + v70
		v76 = F_palloc(m, v39<<(uint(int32(1))%32)+int32(4))
		mBase = m.M
		v79 = m.ExcPending
		if v79 != 0 {
			return int32(0)
		} else {
			v81 = v76 + int32(4)
			switch v67 {
			case 0:
				if v39 <= int32(0) {
					v234 = v81
				} else {
					if v39&int32(1) != 0 {
						v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
						if v86 == int32(92) {
							v89 = int32(92)
							*(*uint8)(unsafe.Add(mBase, uint32(v76)+4)) = uint8(v89)
							v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
							v94 = v93
							v95 = v76 + int32(5)
						} else {
							v94 = v86
							v95 = v81
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v95))) = uint8(v94)
						v97 = int32(1)
						v104 = v95 + v97
						v105 = v71 + v97
						v106 = v39 - v97
					} else {
						v104 = v81
						v105 = v71
						v106 = v39
					}
					if v39 == int32(1) {
						v234 = v104
					} else {
						v109 = v106
						v111 = v104
						v113 = v105
						for {
							v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
							if v118 == int32(92) {
								v121 = int32(92)
								*(*uint8)(unsafe.Add(mBase, uint32(v111))) = uint8(v121)
								v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
								v126 = v125
								v127 = v111 + int32(1)
							} else {
								v126 = v118
								v127 = v111
							}
							*(*uint8)(unsafe.Add(mBase, uint32(v127))) = uint8(v126)
							v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)))
							if v129 != int32(92) {
								v139 = v129
								v140 = v127 + int32(1)
							} else {
								v134 = int32(92)
								*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)) = uint8(v134)
								v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)))
								v139 = v136
								v140 = v127 + int32(2)
							}
							*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v139)
							v143 = v140 + int32(1)
							v144 = int32(2)
							if v144 < v109 {
								v109 = v109 - v144
								v111 = v143
								v113 = v113 + v144
								continue
							} else {
								break
							}
							break
						}
						v234 = v143
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(v76))) = (v234 - v76) << (uint(int32(2)) % 32)
				return v76
			case 1:
				v180 = int32(1)
				v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				if v182&v180 != 0 {
					v185 = v180
				} else {
					v185 = int32(4)
				}
				v186 = l1 + v185
				v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
				if v187 == int32(92) {
					v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if v245 == int32(1) {
						v249 = int32(18)
						v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
						if v251 == v249 {
							v254 = v249
						} else {
							v254 = int32(2)
						}
						if base.Ui32((v251-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v261 = int32(6)
						} else {
							v261 = v254
						}
						v270 = v261
					} else {
						v262 = int32(1)
						if v245&v262 != 0 {
							v270 = int32(base.Ui32(v245) >> (uint(v262) % 32))
						} else {
							v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v270 = int32(base.Ui32(v266) >> (uint(int32(2)) % 32))
						}
					}
					if v270 == int32(0) {
						return v76
					} else {
						base.MemoryCopy(m, v76, l0, v270)
						return v76
					}
				} else {
					v190 = int32(0)
					if v39 <= v190 {
						v234 = v81
					} else {
						v194 = v190
						v195 = v81
						v196 = v39
						v197 = v71
						for {
							v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
							v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
							v207 = base.B2i32(v194 == int32(0)) & base.B2i32(v204 == v205)
							if v207 != 0 {
								v208 = int32(92)
								*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v208)
								v225 = v195 + int32(1)
							} else {
								v213 = v195 + int32(1)
								if v204 == int32(92) {
									v216 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v216)
									if v194&int32(1) != 0 {
										v225 = v213
									} else {
										v220 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v195)+1)) = uint8(v220)
										v225 = v195 + int32(2)
									}
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v204)
									v225 = v213
								}
							}
							v226 = int32(1)
							if v226 < v196 {
								v194 = v207
								v195 = v225
								v196 = v196 - v226
								v197 = v197 + v226
								continue
							} else {
								break
							}
							break
						}
						v234 = v225
					}
					*(*int32)(unsafe.Add(mBase, uint32(v76))) = (v234 - v76) << (uint(int32(2)) % 32)
					return v76
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v163 = m.ExcPending
				if v163 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(84410498))
					mBase = m.M
					v166 = m.ExcPending
					if v166 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_SB_do_like_escape_0), int32(0))
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(_a_F_SB_do_like_escape_1), int32(0))
							mBase = m.M
							v174 = m.ExcPending
							if v174 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_SB_do_like_escape_2), int32(453), int32(_a_F_SB_do_like_escape_3))
								mBase = m.M
								v179 = m.ExcPending
								if v179 != 0 {
									return int32(0)
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
}
func F_SIInsertDataEntries(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	if l1 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_SIInsertDataEntries[0]))
	v15 = v13 + int32(12)
	v18 = l0
	v19 = l1
	goto L3
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_SIInsertDataEntries[1]))
	v32 = F_LWLockAcquire(m, v28+int32(768), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v34 = int32(64)
	if base.Ui32(v34) <= base.Ui32(v19) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v37 = v34
	goto L9
L8:
	;
	v37 = v19
	goto L9
L9:
	;
	goto L10
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v49 = v47 - v48
	if int32(_a_F_SIInsertDataEntries_0) < v49+v37 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_SICleanupQueue(m, int32(1), v37)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L30
	}
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v53 <= v49 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v55 = v18
	v58 = v47
	v60 = v37
	goto L15
L15:
	;
	v65 = base.I32_rem_s(v58, int32(_a_F_SIInsertDataEntries_0))
	v68 = v13 + int32(16) + v65<<(uint(int32(4))%32)
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v55)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v68)+8)) = v69
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	*(*int64)(unsafe.Add(mBase, uint32(v68))) = v71
	v73 = int32(1)
	v74 = v58 + v73
	v76 = v55 + int32(16)
	if v73 < v60 {
		v55 = v76
		v58 = v74
		v60 = v60 - v73
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v83 = base.AtomicRmwXchg32(m, v15, int32(0), int32(1))
	if v83 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	F_s_lock(m, v15, int32(_a_F_SIInsertDataEntries_1))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v87 = v19 - v37
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v74
	v89 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v13)+12)), uint32(v89))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_c_F_SIInsertDataEntries[2])))
	if v89 < v92 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_c_F_SIInsertDataEntries[3])))
	v100 = int32(0)
	goto L25
L23:
	;
	goto L24
L24:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_SIInsertDataEntries[1]))
	F_LWLockRelease(m, v131+int32(768))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L28
	}
L25:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v95+v100<<(uint(int32(2))%32))))
	v115 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13+v109<<(uint(int32(4))%32))+uint32(_c_F_SIInsertDataEntries[4]))) = uint8(v115)
	v118 = v100 + v115
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_c_F_SIInsertDataEntries[2])))
	if v118 < v119 {
		v100 = v118
		goto L25
	} else {
		goto L27
	}
L26:
	;
	goto L24
L27:
	;
	goto L26
L28:
	;
	if int32(0) < v87 {
		v18 = v76
		v19 = v87
		goto L3
	} else {
		goto L29
	}
L29:
	;
	goto L1
L30:
	;
	goto L10
}
func F_SUBTRANSShmemRequest(m *base.Module, l0 int32) {
	var v13 int32
	_ = v13
	Fn14207(m, l0, int64(21474836480), int32(_a_F_SUBTRANSShmemRequest_0), int32(410), int32(409), int64(408021893180), int32(_a_F_SUBTRANSShmemRequest_1), int32(_a_F_SUBTRANSShmemRequest_2), int32(_a_F_SUBTRANSShmemRequest_3), int32(_a_F_SUBTRANSShmemRequest_4), int32(_a_F_SUBTRANSShmemRequest_5))
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		return
	}
}
func F_SampleNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int64
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
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
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
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
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v271 int64
	_ = v271
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)))
	if v15 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v317
L2:
	;
	v314 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+152)) = uint16(v314)
	v317 = v212
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L11
	} else {
		goto L79
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L11
	} else {
		goto L75
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = int64(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v23 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+8))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+12))
	m.T0[v209].(func(*base.Module, int32))(m, v207)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L11
	} else {
		goto L58
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v26 = v24
	goto L10
L9:
	;
	v26 = int32(0)
	goto L10
L10:
	;
	v27 = F_palloc_mul(m, int32(8), v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v31 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v100 != 0 {
		goto L27
	} else {
		goto L28
	}
L14:
	;
	v34 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v35 <= v34 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v39 = v34
	goto L16
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v39<<(uint(int32(2))%32))))
	v53 = int32(_a_F_SampleNext_0)
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_SampleNext[0]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_SampleNext[0])) = v56
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	v61 = m.T0[v60].(func(*base.Module, int32, int32, int32) int64)(m, v52, v18, v13+int32(15))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L11
	} else {
		goto L19
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L11
	} else {
		goto L22
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SampleNext[0])) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v27+v39<<(uint(int32(3))%32)))) = v61
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	if v69 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v71 = v39 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v71 < v72 {
		v39 = v71
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L13
L22:
	;
	F_errcode(m, int32(403177602))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L11
	} else {
		goto L23
	}
L23:
	;
	F_errmsg(m, int32(_a_F_SampleNext_1), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_SampleNext_2), int32(245), int32(_a_F_SampleNext_3))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	v125 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+132)) = uint16(v125)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v128 != 0 {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	v101 = int32(_a_F_SampleNext_0)
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_SampleNext[0]))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_SampleNext[0])) = v104
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v100)+24))
	v109 = m.T0[v108].(func(*base.Module, int32, int32, int32) int64)(m, v100, v18, v13+int32(15))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L11
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v124 = v121
	goto L26
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SampleNext[0])) = v102
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	if v113 == int32(1) {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v118 = F_DirectFunctionCall1Coll(m, int32(794), int32(0), v109)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L11
	} else {
		goto L32
	}
L32:
	;
	v124 = base.I32_wrap_i64(v118)
	goto L26
L33:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	v131 = v129
	goto L35
L34:
	;
	v131 = int32(0)
	goto L35
L35:
	;
	m.T0[v127].(func(*base.Module, int32, int32, int32, int32))(m, l0, v27, v131, v124)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v135 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	F_pfree(m, v27)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L11
	} else {
		goto L57
	}
L38:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	v143 = F_ScanRelIsReadOnly(m, l0)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L11
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v176 = int32(0)
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+188))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+16))
	m.T0[v184].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v135, v176, int32(1), v178, base.B2i32(v134 == v176), v181)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L11
	} else {
		goto L56
	}
L41:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_SampleNext[1]))
	if v146 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SampleNext[2])))
	if v148&int32(1) == int32(0) {
		goto L3
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v153 = int32(0)
	if v139&int32(1) != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	v162 = int32(68)
	goto L48
L47:
	;
	v162 = int32(4)
	goto L48
L48:
	;
	if v134 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v165 = v162
	goto L51
L50:
	;
	v165 = v162 | int32(128)
	goto L51
L51:
	;
	if v143 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v169 = int32(1024)
	goto L54
L53:
	;
	v169 = int32(0)
	goto L54
L54:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v140)+188))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+8))
	v173 = m.T0[v172].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v140, v142, v153, v153, v153, v138<<(uint(int32(8))%32)|v165|v169)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L11
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v173
	goto L37
L56:
	;
	goto L37
L57:
	;
	v194 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)) = uint8(v194)
	goto L7
L58:
	;
	v212 = int32(0)
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+153)))
	if v213 != 0 {
		v317 = v212
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)))
	if v214 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+188))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+172))
	v220 = m.T0[v219].(func(*base.Module, int32, int32) int32)(m, v206, l0)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L11
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+188))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+176))
	v229 = m.T0[v228].(func(*base.Module, int32, int32, int32) int32)(m, v206, l0, v207)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L11
	} else {
		goto L65
	}
L63:
	;
	if v220 == int32(0) {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	v224 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)) = uint8(v224)
	goto L62
L65:
	;
	if v229 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	goto L69
L67:
	;
	goto L68
L68:
	;
	v271 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = v271 + int64(1)
	v317 = v207
	goto L1
L69:
	;
	v243 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)) = uint8(v243)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+188))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+172))
	v248 = m.T0[v247].(func(*base.Module, int32, int32) int32)(m, v206, l0)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L11
	} else {
		goto L71
	}
L70:
	;
	goto L68
L71:
	;
	if v248 == int32(0) {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v252 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)) = uint8(v252)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+188))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+176))
	v257 = m.T0[v256].(func(*base.Module, int32, int32, int32) int32)(m, v206, l0, v207)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	if v257 == int32(0) {
		goto L69
	} else {
		goto L74
	}
L74:
	;
	goto L70
L75:
	;
	F_errcode(m, int32(386400386))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L11
	} else {
		goto L76
	}
L76:
	;
	F_errmsg(m, int32(_a_F_SampleNext_4), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L11
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_SampleNext_2), int32(257), int32(_a_F_SampleNext_3))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	F_errmsg_internal(m, int32(_a_F_SampleNext_5), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_SampleNext_6), int32(931), int32(_a_F_SampleNext_7))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L11
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ScanKeywords_hash_func(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	v3 = int32(0)
	if l1 == v3 {
		v86 = int32(0)
		v93 = int32(3)
	} else {
		if l1 == int32(1) {
			v56 = l0
			v57 = int32(0)
			v58 = int32(3)
			v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
			v66 = v64 | int32(32)
			v74 = v57*int32(257) + v66
			v75 = v66 + v58*int32(_a_F_ScanKeywords_hash_func_0)
		} else {
			v23 = l0
			v24 = int32(0)
			v25 = int32(3)
			v30 = v3
			for {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v32 = int32(32)
				v33 = v31 | v32
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v36 = v34 | v32
				v37 = int32(_a_F_ScanKeywords_hash_func_0)
				v42 = v33 + (v36+v25*v37)*v37
				v43 = int32(257)
				v48 = (v24*v43+v36)*v43 + v33
				v49 = int32(2)
				v50 = v23 + v49
				v52 = v30 + v49
				if v52 != l1&int32(-2) {
					v23 = v50
					v24 = v48
					v25 = v42
					v30 = v52
					continue
				} else {
					break
				}
				break
			}
			if l1&int32(1) == int32(0) {
				v74 = v48
				v75 = v42
			} else {
				v56 = v50
				v57 = v48
				v58 = v42
				v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
				v66 = v64 | int32(32)
				v74 = v57*int32(257) + v66
				v75 = v66 + v58*int32(_a_F_ScanKeywords_hash_func_0)
			}
		}
		v81 = int32(999)
		v82 = base.I32_rem_u_s(v74, v81)
		v84 = base.I32_rem_u_s(v75, v81)
		v86 = v82
		v93 = v84
	}
	v94 = int32(1)
	v96 = int32(*(*int16)(unsafe.Add(mBase, uint32(v93<<(uint(v94)%32))+uint32(_c_F_ScanKeywords_hash_func[0]))))
	v99 = int32(*(*int16)(unsafe.Add(mBase, uint32(v86<<(uint(v94)%32))+uint32(_c_F_ScanKeywords_hash_func[0]))))
	return v96 + v99
}
func F_SelectNeighbors(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int64
	_ = v147
	var v148 int32
	_ = v148
	var v149 float32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v169 int32
	_ = v169
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v209 int32
	_ = v209
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int64
	_ = v240
	var v241 int32
	_ = v241
	var v242 float32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v281 int32
	_ = v281
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v312 int64
	_ = v312
	var v313 int32
	_ = v313
	var v314 float32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v370 int32
	_ = v370
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v401 int64
	_ = v401
	var v402 int32
	_ = v402
	var v403 float32
	_ = v403
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v447 int32
	_ = v447
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v479 int32
	_ = v479
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v579 int32
	_ = v579
	var v589 int32
	_ = v589
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	v8 = l7
	v9 = int32(0)
	v21 = F_list_copy(m, l1)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v21 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	return v625
L4:
	;
	v36 = F_palloc_mul(m, int32(4), v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L10
	}
L5:
	;
	v28 = int32(0)
	if v28 <= l2 {
		v625 = v28
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if l2 < v31 {
		v34 = v31
		goto L4
	} else {
		goto L9
	}
L8:
	;
	v34 = v28
	goto L4
L9:
	;
	return v21
L10:
	;
	if v8 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if l0 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	if v21 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v40 = int32(_a_F_SelectNeighbors_0)
	goto L16
L15:
	;
	v40 = int32(_a_F_SelectNeighbors_1)
	goto L16
L16:
	;
	F_list_sort(m, v21, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v621)))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v622
	v625 = v602
	goto L3
L19:
	;
	if l6 == int32(0) {
		v625 = v579
		goto L3
	} else {
		goto L144
	}
L20:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v8)
	v46 = int32(0)
	v579 = v46
	v589 = v46
	goto L19
L21:
	;
	goto L22
L22:
	;
	v52 = int32(0)
	v62 = v21
	v63 = v9
	v64 = v9
	v65 = v9
	goto L24
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v8)
	if v530 <= int32(0) {
		v579 = v517
		v589 = v527
		goto L19
	} else {
		goto L132
	}
L24:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v71 <= int32(0) {
		v517 = v52
		v527 = v62
		v530 = v65
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v517 = v513
	v527 = int32(0)
	v530 = v514
	goto L23
L26:
	;
	if v52 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v76 = v74
	goto L29
L28:
	;
	v76 = int32(0)
	goto L29
L29:
	;
	if l2 <= v76 {
		v517 = v52
		v527 = v62
		v530 = v65
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v78+v71<<(uint(int32(2))%32)-int32(4))))
	v85 = F_list_delete_last(m, v62)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v25&int32(1) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+8)))
	if v502 == int32(1) {
		goto L127
	} else {
		goto L128
	}
L33:
	;
	if l0 != 0 {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	goto L35
L35:
	;
	if v64 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L36:
	;
	v100 = int32(1)
	if v52 == int32(0) {
		v169 = v100
		goto L42
	} else {
		goto L43
	}
L37:
	;
	v99 = l0 + v91 - int32(1)
	goto L36
L38:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0+v89)+87))
	if v91 != 0 {
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+88))
	v99 = v94
	goto L36
L41:
	;
	v99 = int32(0)
	goto L36
L42:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+8)) = uint8(v169)
	v494 = v63
	v495 = v64
	goto L32
L43:
	;
	v103 = int32(0)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v104 <= v103 {
		v169 = v100
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v116 = v103
	goto L45
L45:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128+v116<<(uint(int32(2))%32))))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if l0 != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v169 = v154
	goto L42
L47:
	;
	v147 = F_FunctionCall2Coll(m, v133, v134, base.I64_extend_i32_u(v99), base.I64_extend_i32_u(v145))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L53
	}
L48:
	;
	v145 = l0 + v137 - int32(1)
	goto L47
L49:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0+v135)+87))
	if v137 != 0 {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+88))
	v145 = v140
	goto L47
L52:
	;
	v145 = int32(0)
	goto L47
L53:
	;
	v149 = *(*float32)(unsafe.Add(mBase, uint32(v84)+4))
	v152 = base.F32_ge(v149, base.F32_demote_f64(base.F64_reinterpret_i64(v147)))
	v154 = base.B2i32(v152 == int32(0))
	if v152 != 0 {
		v169 = v154
		goto L42
	} else {
		goto L54
	}
L54:
	;
	v156 = v116 + int32(1)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v156 < v157 {
		v116 = v156
		goto L45
	} else {
		goto L55
	}
L55:
	;
	goto L46
L56:
	;
	v479 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+8)) = uint8(v479)
	v494 = int32(1)
	v495 = v64
	goto L32
L57:
	;
	v455 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+8)) = uint8(v455)
	v457 = F_lappend(m, v64, v84)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L125
	}
L58:
	;
	if l5 != v84 {
		v494 = v63
		v495 = v64
		goto L32
	} else {
		goto L102
	}
L59:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v182 <= int32(0) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+8)))
	if v185 == int32(1) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	if l0 != 0 {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	goto L63
L63:
	;
	v252 = int32(0)
	if v63 == v252 {
		v494 = v252
		v495 = v64
		goto L32
	} else {
		goto L81
	}
L64:
	;
	v209 = int32(0)
	goto L70
L65:
	;
	v198 = l0 + v190 - int32(1)
	goto L64
L66:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0+v188)+87))
	if v190 != 0 {
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+88))
	v198 = v193
	goto L64
L69:
	;
	v198 = int32(0)
	goto L64
L70:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v221+v209<<(uint(int32(2))%32))))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if l0 != 0 {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	v250 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+8)) = uint8(v250)
	v494 = v63
	v495 = v64
	goto L32
L72:
	;
	v240 = F_FunctionCall2Coll(m, v226, v227, base.I64_extend_i32_u(v198), base.I64_extend_i32_u(v238))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L78
	}
L73:
	;
	v238 = l0 + v230 - int32(1)
	goto L72
L74:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0+v228)+87))
	if v230 != 0 {
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+88))
	v238 = v233
	goto L72
L77:
	;
	v238 = int32(0)
	goto L72
L78:
	;
	v242 = *(*float32)(unsafe.Add(mBase, uint32(v84)+4))
	if base.F32_ge(v242, base.F32_demote_f64(base.F64_reinterpret_i64(v240))) != 0 {
		goto L56
	} else {
		goto L79
	}
L79:
	;
	v247 = v209 + int32(1)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v247 < v248 {
		v209 = v247
		goto L70
	} else {
		goto L80
	}
L80:
	;
	goto L71
L81:
	;
	if l0 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	if v52 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L83:
	;
	v265 = l0 + v257 - int32(1)
	goto L82
L84:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0+v255)+87))
	if v257 != 0 {
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+88))
	v265 = v260
	goto L82
L87:
	;
	v265 = int32(0)
	goto L82
L88:
	;
	v447 = int32(1)
	goto L57
L89:
	;
	v268 = int32(0)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v269 <= v268 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v281 = v268
	goto L91
L91:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v293+v281<<(uint(int32(2))%32))))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if l0 != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	goto L88
L93:
	;
	v312 = F_FunctionCall2Coll(m, v298, v299, base.I64_extend_i32_u(v265), base.I64_extend_i32_u(v310))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L99
	}
L94:
	;
	v310 = l0 + v302 - int32(1)
	goto L93
L95:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0+v300)+87))
	if v302 != 0 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+88))
	v310 = v305
	goto L93
L98:
	;
	v310 = int32(0)
	goto L93
L99:
	;
	v314 = *(*float32)(unsafe.Add(mBase, uint32(v84)+4))
	if base.F32_ge(v314, base.F32_demote_f64(base.F64_reinterpret_i64(v312))) != 0 {
		goto L56
	} else {
		goto L100
	}
L100:
	;
	v319 = v281 + int32(1)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v319 < v320 {
		v281 = v319
		goto L91
	} else {
		goto L101
	}
L101:
	;
	goto L92
L102:
	;
	if l0 != 0 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	if v52 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L104:
	;
	v354 = l0 + v346 - int32(1)
	goto L103
L105:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0+v344)+87))
	if v346 != 0 {
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)+88))
	v354 = v349
	goto L103
L108:
	;
	v354 = int32(0)
	goto L103
L109:
	;
	v447 = v63
	goto L57
L110:
	;
	v357 = int32(0)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v358 <= v357 {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v370 = v357
	goto L112
L112:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v382+v370<<(uint(int32(2))%32))))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if l0 != 0 {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	v413 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+8)) = uint8(v413)
	v494 = v63
	v495 = v64
	goto L32
L114:
	;
	v401 = F_FunctionCall2Coll(m, v387, v388, base.I64_extend_i32_u(v354), base.I64_extend_i32_u(v399))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L120
	}
L115:
	;
	v399 = l0 + v391 - int32(1)
	goto L114
L116:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v386)))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0+v389)+87))
	if v391 != 0 {
		goto L115
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v386)))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+88))
	v399 = v394
	goto L114
L119:
	;
	v399 = int32(0)
	goto L114
L120:
	;
	v403 = *(*float32)(unsafe.Add(mBase, uint32(v84)+4))
	if base.F32_ge(v403, base.F32_demote_f64(base.F64_reinterpret_i64(v401))) == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v410 = v370 + int32(1)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v411 <= v410 {
		goto L109
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	goto L113
L124:
	;
	v370 = v410
	goto L112
L125:
	;
	v494 = v447
	v495 = v457
	goto L32
L126:
	;
	if v85 != 0 {
		v52 = v513
		v62 = v85
		v63 = v494
		v64 = v495
		v65 = v514
		goto L24
	} else {
		goto L131
	}
L127:
	;
	v505 = F_lappend(m, v52, v84)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36+v65<<(uint(int32(2))%32)))) = v84
	v513 = v52
	v514 = v65 + int32(1)
	goto L126
L130:
	;
	v513 = v505
	v514 = v65
	goto L126
L131:
	;
	goto L25
L132:
	;
	v540 = int32(0)
	v541 = v517
	goto L133
L133:
	;
	if v541 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if l6 == int32(0) {
		v625 = v541
		goto L3
	} else {
		goto L143
	}
L135:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v541)+4))
	v562 = v560
	goto L137
L136:
	;
	v562 = int32(0)
	goto L137
L137:
	;
	if v562 < l2 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v36+v540<<(uint(int32(2))%32))))
	v568 = F_lappend(m, v541, v567)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	goto L134
L141:
	;
	v571 = v540 + int32(1)
	if v571 == v530 {
		v579 = v568
		v589 = v527
		goto L19
	} else {
		goto L142
	}
L142:
	;
	v540 = v571
	v541 = v568
	goto L133
L143:
	;
	v602 = v541
	v621 = v36 + v540<<(uint(int32(2))%32)
	goto L18
L144:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v589)+12))
	v602 = v579
	v621 = v600
	goto L18
}
func F_SetEpochTimestamp(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v31 int32
	_ = v31
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v77 int64
	_ = v77
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v100 int64
	_ = v100
	var v107 int64
	_ = v107
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v123 int32
	_ = v123
	var v131 int64
	_ = v131
	var v134 int64
	_ = v134
	var v144 int64
	_ = v144
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
	v18 = F_pg_gmtime(m, v12+int32(24))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		if v18 != 0 {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
			v24 = v22 + int32(1)
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
			v31 = v29 + int32(1900)
			if v31 <= int32(-4713) {
				if v31 != int32(-4713) {
					v144 = int64(0)
				} else {
					if int32(10) < v24 {
						v49 = base.B2i32(int32(2) < v24)
						if int32(2) < v24 {
							v50 = int32(_a_F_SetEpochTimestamp_0)
						} else {
							v50 = int32(_a_F_SetEpochTimestamp_1)
						}
						v51 = v50 + v31
						v56 = base.I32_div_s(v51, int32(4))
						v59 = base.I32_div_s(v51, int32(-100))
						v62 = base.I32_div_s(v51, int32(400))
						if int32(2) < v24 {
							v66 = int32(1)
						} else {
							v66 = int32(13)
						}
						v71 = base.I32_div_s((v66+v24)*int32(_a_F_SetEpochTimestamp_2), int32(256))
						v77 = base.I64_extend_i32_s(v28 + v51*int32(365) + v56 + v59 + v62 + v71 - int32(_a_F_SetEpochTimestamp_3) - int32(_a_F_SetEpochTimestamp_4))
						v86 = int64(32)
						v87 = int64(20)
						v89 = int64(base.Ui64(v77) >> (uint(v86) % 64))
						v92 = int64(4294967295)
						v93 = int64(500654080)
						v95 = v77 & v92
						v96 = v93 * v95
						v100 = int64(base.Ui64(v96)>>(uint(v86)%64)) + v93*v89
						v107 = v95*v87 + v100&v92
						*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v77*int64(0) + v77>>(uint(int64(63))%64)*int64(86400000000) + v87*v89 + int64(base.Ui64(v100)>>(uint(v86)%64)) + int64(base.Ui64(v107)>>(uint(v86)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v12))) = v96&v92 | v107<<(uint(v86)%64)
						v118 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
						v119 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
						if v118 != v119>>(uint(int64(63))%64) {
							v144 = int64(0)
						} else {
							v123 = int32(60)
							v131 = base.I64_extend_i32_s((v27*v123+v26)*v123+v25) * int64(1000000)
							v134 = v119 + v131
							if base.B2i32(v131 < int64(0))^base.B2i32(v134 < v119) != 0 {
								v144 = int64(0)
							} else {
								if base.Ui64(int64(9011559254509551615)) < base.Ui64(v134-int64(9223371331200000000)) {
									v144 = v134
								} else {
									v144 = int64(0)
								}
							}
						}
					} else {
						v144 = int64(0)
					}
				}
			} else {
				if v31 < int32(_a_F_SetEpochTimestamp_5) {
					v49 = base.B2i32(int32(2) < v24)
					if int32(2) < v24 {
						v50 = int32(_a_F_SetEpochTimestamp_0)
					} else {
						v50 = int32(_a_F_SetEpochTimestamp_1)
					}
					v51 = v50 + v31
					v56 = base.I32_div_s(v51, int32(4))
					v59 = base.I32_div_s(v51, int32(-100))
					v62 = base.I32_div_s(v51, int32(400))
					if int32(2) < v24 {
						v66 = int32(1)
					} else {
						v66 = int32(13)
					}
					v71 = base.I32_div_s((v66+v24)*int32(_a_F_SetEpochTimestamp_2), int32(256))
					v77 = base.I64_extend_i32_s(v28 + v51*int32(365) + v56 + v59 + v62 + v71 - int32(_a_F_SetEpochTimestamp_3) - int32(_a_F_SetEpochTimestamp_4))
					v86 = int64(32)
					v87 = int64(20)
					v89 = int64(base.Ui64(v77) >> (uint(v86) % 64))
					v92 = int64(4294967295)
					v93 = int64(500654080)
					v95 = v77 & v92
					v96 = v93 * v95
					v100 = int64(base.Ui64(v96)>>(uint(v86)%64)) + v93*v89
					v107 = v95*v87 + v100&v92
					*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v77*int64(0) + v77>>(uint(int64(63))%64)*int64(86400000000) + v87*v89 + int64(base.Ui64(v100)>>(uint(v86)%64)) + int64(base.Ui64(v107)>>(uint(v86)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v12))) = v96&v92 | v107<<(uint(v86)%64)
					v118 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
					v119 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
					if v118 != v119>>(uint(int64(63))%64) {
						v144 = int64(0)
					} else {
						v123 = int32(60)
						v131 = base.I64_extend_i32_s((v27*v123+v26)*v123+v25) * int64(1000000)
						v134 = v119 + v131
						if base.B2i32(v131 < int64(0))^base.B2i32(v134 < v119) != 0 {
							v144 = int64(0)
						} else {
							if base.Ui64(int64(9011559254509551615)) < base.Ui64(v134-int64(9223371331200000000)) {
								v144 = v134
							} else {
								v144 = int64(0)
							}
						}
					}
				} else {
					if base.B2i32(v31 != int32(_a_F_SetEpochTimestamp_5))|base.B2i32(int32(5) < v24) != 0 {
						v144 = int64(0)
					} else {
						v49 = base.B2i32(int32(2) < v24)
						if int32(2) < v24 {
							v50 = int32(_a_F_SetEpochTimestamp_0)
						} else {
							v50 = int32(_a_F_SetEpochTimestamp_1)
						}
						v51 = v50 + v31
						v56 = base.I32_div_s(v51, int32(4))
						v59 = base.I32_div_s(v51, int32(-100))
						v62 = base.I32_div_s(v51, int32(400))
						if int32(2) < v24 {
							v66 = int32(1)
						} else {
							v66 = int32(13)
						}
						v71 = base.I32_div_s((v66+v24)*int32(_a_F_SetEpochTimestamp_2), int32(256))
						v77 = base.I64_extend_i32_s(v28 + v51*int32(365) + v56 + v59 + v62 + v71 - int32(_a_F_SetEpochTimestamp_3) - int32(_a_F_SetEpochTimestamp_4))
						v86 = int64(32)
						v87 = int64(20)
						v89 = int64(base.Ui64(v77) >> (uint(v86) % 64))
						v92 = int64(4294967295)
						v93 = int64(500654080)
						v95 = v77 & v92
						v96 = v93 * v95
						v100 = int64(base.Ui64(v96)>>(uint(v86)%64)) + v93*v89
						v107 = v95*v87 + v100&v92
						*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v77*int64(0) + v77>>(uint(int64(63))%64)*int64(86400000000) + v87*v89 + int64(base.Ui64(v100)>>(uint(v86)%64)) + int64(base.Ui64(v107)>>(uint(v86)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v12))) = v96&v92 | v107<<(uint(v86)%64)
						v118 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
						v119 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
						if v118 != v119>>(uint(int64(63))%64) {
							v144 = int64(0)
						} else {
							v123 = int32(60)
							v131 = base.I64_extend_i32_s((v27*v123+v26)*v123+v25) * int64(1000000)
							v134 = v119 + v131
							if base.B2i32(v131 < int64(0))^base.B2i32(v134 < v119) != 0 {
								v144 = int64(0)
							} else {
								if base.Ui64(int64(9011559254509551615)) < base.Ui64(v134-int64(9223371331200000000)) {
									v144 = v134
								} else {
									v144 = int64(0)
								}
							}
						}
					}
				}
			}
			m.G0 = v12 + int32(32)
			return v144
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v153 = m.ExcPending
			if v153 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_SetEpochTimestamp_6), int32(0))
				mBase = m.M
				v157 = m.ExcPending
				if v157 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_SetEpochTimestamp_7), int32(2188), int32(_a_F_SetEpochTimestamp_8))
					mBase = m.M
					v162 = m.ExcPending
					if v162 != 0 {
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
func F_SetupApplyOrSyncWorker(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	F_logicalrep_worker_attach(m, l0)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v9 = m.G0
		v11 = v9 - int32(32)
		m.G0 = v11
		v14 = int32(967)
		switch v14 {
		case 0, 2:
		default:
			*(*int32)(unsafe.Add(mBase, _c_F_SetupApplyOrSyncWorker[0])) = int32(965)
		}
		F_sigemptyset(m, v11+int32(16))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(268435456)
		switch v14 {
		case 0:
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(-2)
		default:
			*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(268435460)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(_a_F_SetupApplyOrSyncWorker_0)
		case 2:
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(0)
		}
		v43 = F___sigaction(m, int32(1), v11+int32(12), int32(0))
		mBase = m.M
		m.G0 = v11 + int32(32)
		F_BackgroundWorkerUnblockSignals(m)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return
		} else {
			v50 = *(*int32)(unsafe.Add(mBase, _c_F_SetupApplyOrSyncWorker[1]))
			v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+16)))
			if v51 != int32(1) {
				v66 = m.G0
				v67 = int32(16)
				v68 = v66 - v67
				m.G0 = v68
				F_gettimeofday(m, v68)
				mBase = m.M
				v71 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
				v72 = int64(*(*int32)(unsafe.Add(mBase, uint32(v68)+8)))
				m.G0 = v68 + v67
				v80 = v72 + v71*int64(1000000) - int64(946684800000000)
				v82 = *(*int32)(unsafe.Add(mBase, _c_F_SetupApplyOrSyncWorker[1]))
				*(*int64)(unsafe.Add(mBase, uint32(v82)+96)) = v80
				*(*int64)(unsafe.Add(mBase, uint32(v82)+112)) = v80
				*(*int64)(unsafe.Add(mBase, uint32(v82)+88)) = v80
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
				if v54 != int32(2) {
					v66 = m.G0
					v67 = int32(16)
					v68 = v66 - v67
					m.G0 = v68
					F_gettimeofday(m, v68)
					mBase = m.M
					v71 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
					v72 = int64(*(*int32)(unsafe.Add(mBase, uint32(v68)+8)))
					m.G0 = v68 + v67
					v80 = v72 + v71*int64(1000000) - int64(946684800000000)
					v82 = *(*int32)(unsafe.Add(mBase, _c_F_SetupApplyOrSyncWorker[1]))
					*(*int64)(unsafe.Add(mBase, uint32(v82)+96)) = v80
					*(*int64)(unsafe.Add(mBase, uint32(v82)+112)) = v80
					*(*int64)(unsafe.Add(mBase, uint32(v82)+88)) = v80
				} else {
					v57 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v50)+88)) = v57
					*(*int64)(unsafe.Add(mBase, uint32(v50)+112)) = v57
					*(*int64)(unsafe.Add(mBase, uint32(v50)+96)) = v57
				}
			}
			F_load_file(m, int32(_a_F_SetupApplyOrSyncWorker_1), int32(0))
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return
			} else {
				F_InitializeLogRepWorker(m)
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return
				} else {
					F_CacheRegisterSyscacheCallback(m, int32(68), int32(1051), int64(0))
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_StartChildProcess(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = F_AssignPostmasterChildSlot(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			v15 = int32(0)
			v18 = F_errstart(m, int32(15), v15)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if l0 == int32(4) {
					if v18 == int32(0) {
						v80 = v15
						m.G0 = v7 + int32(16)
						return v80
					} else {
						F_errcode(m, int32(_a_F_StartChildProcess_0))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_StartChildProcess_1), int32(0))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_StartChildProcess_2), int32(4029), int32(_a_F_StartChildProcess_3))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									v80 = v15
									m.G0 = v7 + int32(16)
									return v80
								}
							}
						}
					}
				} else {
					if v18 == int32(0) {
						v80 = v15
						m.G0 = v7 + int32(16)
						return v80
					} else {
						F_errmsg_internal(m, int32(_a_F_StartChildProcess_4), int32(0))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_StartChildProcess_2), int32(4033), int32(_a_F_StartChildProcess_3))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v80 = v15
								m.G0 = v7 + int32(16)
								return v80
							}
						}
					}
				}
			}
		} else {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v48 = int32(0)
			v51 = F_postmaster_child_launch(m, l0, v47, v48, v48, v48)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				if v51 < int32(0) {
					v55 = F_ReleasePostmasterChildSlot(m, v9)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						v59 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							if v59 != 0 {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(12))+uint32(_c_F_StartChildProcess[0])))
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v63
								F_errmsg(m, int32(_a_F_StartChildProcess_5), v7)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_StartChildProcess_2), int32(4044), int32(_a_F_StartChildProcess_3))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int32(0)
									} else {
										if l0 != int32(13) {
											v80 = int32(0)
											m.G0 = v7 + int32(16)
											return v80
										} else {
											F_ExitPostmaster(m, int32(1))
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								if l0 != int32(13) {
									v80 = int32(0)
									m.G0 = v7 + int32(16)
									return v80
								} else {
									F_ExitPostmaster(m, int32(1))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v51
					v80 = v9
					m.G0 = v7 + int32(16)
					return v80
				}
			}
		}
	}
}
func F_StatementCancelHandler(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StatementCancelHandler[0])))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_StatementCancelHandler[1])) = v8
	*(*int32)(unsafe.Add(mBase, _c_F_StatementCancelHandler[2])) = v8
	goto L3
L2:
	;
	goto L3
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_StatementCancelHandler[3]))
	v15 = int32(0)
	v18 = base.AtomicRmwOr32(m, v15, int32(_a_F_StatementCancelHandler_0), v15)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(1)
	v22 = int32(0)
	v25 = base.AtomicRmwOr32(m, v22, int32(_a_F_StatementCancelHandler_0), v22)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v26 == v22 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v29 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_StatementCancelHandler[4]))
	if v33 == v29 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v35 = m.G0
	v37 = v35 - int32(16)
	m.G0 = v37
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_StatementCancelHandler[5]))
	if v40 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v63 = F_pgmem_kill(m, v29, int32(23))
	mBase = m.M
	goto L5
L12:
	;
	m.G0 = v37 + int32(16)
	goto L4
L13:
	;
	v43 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+15)) = uint8(v43)
	goto L14
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_StatementCancelHandler[6]))
	v51 = F_write(m, v47, v37+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v51 {
		goto L12
	} else {
		goto L16
	}
L15:
	;
	goto L12
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_StatementCancelHandler[7]))
	if v55 == int32(27) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
}
func F_StrategyCtlShmemRequest(m *base.Module, l0 int32) {
	var v6 int32
	_ = v6
	Fn14222(m, l0, int32(_a_F_StrategyCtlShmemRequest_0), int64(20), int32(_a_F_StrategyCtlShmemRequest_1))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F___shlim(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = l1
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = base.I64_extend_i32_s(v5 - v6)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.B2i32(l1 == int64(0))|base.B2i32(base.I64_extend_i32_s(v12-v6) <= l1) != 0 {
		v19 = v12
	} else {
		v19 = v6 + base.I32_wrap_i64(l1)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v19
	return
}
func F___strchrnul(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	v7 = l1 & int32(255)
	if v7 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v88
L2:
	;
	v78 = v73
	goto L19
L3:
	;
	v73 = v65
	goto L2
L4:
	;
	if l0&int32(3) != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v62 = F_strlen(m, l0)
	mBase = m.M
	return v62 + l0
L7:
	;
	v10 = l0
	goto L10
L8:
	;
	v24 = l0
	goto L9
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v33 = int32(-2139062144)
	if (int32(16843008)-v30|v30)&v33 != v33 {
		v65 = v24
		goto L3
	} else {
		goto L14
	}
L10:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if base.B2i32(v15 == int32(0))|base.B2i32(v7 == v15) != 0 {
		v88 = v10
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v24 = v21
	goto L9
L12:
	;
	v21 = v10 + int32(1)
	if v21&int32(3) != 0 {
		v10 = v21
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v39 = v24
	v41 = v30
	goto L15
L15:
	;
	v45 = v41 ^ v7*int32(16843009)
	v48 = int32(-2139062144)
	if (int32(16843008)-v45|v45)&v48 != v48 {
		v65 = v39
		goto L3
	} else {
		goto L17
	}
L16:
	;
	v73 = v54
	goto L2
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v54 = v39 + int32(4)
	v58 = int32(-2139062144)
	if (v52|(int32(16843008)-v52))&v58 == v58 {
		v39 = v54
		v41 = v52
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v80 == int32(0) {
		v88 = v78
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v88 = v78
	goto L1
L21:
	;
	if v80 != l1&int32(255) {
		v78 = v78 + int32(1)
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
}
func F_sampler_random_init_state(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v22 int64
	_ = v22
	var v27 int64
	_ = v27
	var v32 int64
	_ = v32
	v3 = base.I64_extend_i32_u(l0)
	v6 = v3 + int64(4354685564936845354)
	v7 = int64(30)
	v10 = int64(-4658895280553007687)
	v11 = (int64(base.Ui64(v6)>>(uint(v7)%64)) ^ v6) * v10
	v12 = int64(27)
	v15 = int64(-7723592293110705685)
	v16 = (int64(base.Ui64(v11)>>(uint(v12)%64)) ^ v11) * v15
	v17 = int64(31)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = int64(base.Ui64(v16)>>(uint(v17)%64)) ^ v16
	v22 = v3 - int64(7046029254386353131)
	v27 = (int64(base.Ui64(v22)>>(uint(v7)%64)) ^ v22) * v10
	v32 = (int64(base.Ui64(v27)>>(uint(v12)%64)) ^ v27) * v15
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(base.Ui64(v32)>>(uint(v17)%64)) ^ v32
	return
}
func F_sanitize_char_2(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	Fn14374(m, l0, int32(_a_F_sanitize_char_2_0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_scalarineqsel_wrapper(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 float64
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int64
	_ = v86
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = int64(4599676419421066581)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v28 = F_get_restriction_variable(m, v19, v20, v21, v14+int32(16), v14+int32(12), v14+int32(11))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int64(0)
	} else {
		if v28 == int32(0) {
			v86 = v18
			m.G0 = v14 + int32(48)
			return v86
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
			if v35 != int32(7) {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
				if v38 == int32(0) {
					v86 = v18
					m.G0 = v14 + int32(48)
					return v86
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
					m.T0[v41].(func(*base.Module, int32))(m, v38)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int64(0)
					} else {
						v86 = v18
						m.G0 = v14 + int32(48)
						return v86
					}
				}
			} else {
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+32)))
				if v44 == int32(1) {
					v47 = int64(0)
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
					if v48 == int32(0) {
						v86 = v47
						m.G0 = v14 + int32(48)
						return v86
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
						m.T0[v51].(func(*base.Module, int32))(m, v48)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int64(0)
						} else {
							v86 = v47
							m.G0 = v14 + int32(48)
							return v86
						}
					}
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
					v55 = *(*int64)(unsafe.Add(mBase, uint32(v34)+24))
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)))
					if v56 == int32(0) {
						v59 = F_get_commutator(m, v17)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int64(0)
						} else {
							if v59 == int32(0) {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
								if v63 == int32(0) {
									v86 = v18
									m.G0 = v14 + int32(48)
									return v86
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
									m.T0[v66].(func(*base.Module, int32))(m, v63)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int64(0)
									} else {
										v86 = v18
										m.G0 = v14 + int32(48)
										return v86
									}
								}
							} else {
								v71 = l1 ^ int32(1)
								v72 = v59
								v75 = F_scalarineqsel(m, v19, v72, v71, l2, v16, v14+int32(16), v55, v54)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int64(0)
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
									if v77 != 0 {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
										m.T0[v78].(func(*base.Module, int32))(m, v77)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int64(0)
										} else {
											v86 = base.I64_reinterpret_f64(v75)
											m.G0 = v14 + int32(48)
											return v86
										}
									} else {
										v86 = base.I64_reinterpret_f64(v75)
										m.G0 = v14 + int32(48)
										return v86
									}
								}
							}
						}
					} else {
						v71 = l1
						v72 = v17
						v75 = F_scalarineqsel(m, v19, v72, v71, l2, v16, v14+int32(16), v55, v54)
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int64(0)
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
							if v77 != 0 {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
								m.T0[v78].(func(*base.Module, int32))(m, v77)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int64(0)
								} else {
									v86 = base.I64_reinterpret_f64(v75)
									m.G0 = v14 + int32(48)
									return v86
								}
							} else {
								v86 = base.I64_reinterpret_f64(v75)
								m.G0 = v14 + int32(48)
								return v86
							}
						}
					}
				}
			}
		}
	}
}
func F_scalarlesel(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_scalarineqsel_wrapper(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_scalbnl(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	if int32(_a_F_scalbnl_0) <= l3 {
		F___multf3(m, v8+int32(32), l1, l2, int64(0), int64(9222809086901354496))
		mBase = m.M
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v8)+40))
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v8)+32))
		if base.Ui32(l3) < base.Ui32(int32(_a_F_scalbnl_1)) {
			v62 = v18
			v63 = v17
			v64 = l3 - int32(_a_F_scalbnl_2)
		} else {
			F___multf3(m, v8+int32(16), v18, v17, int64(0), int64(9222809086901354496))
			mBase = m.M
			v28 = int32(_a_F_scalbnl_3)
			if base.Ui32(v28) <= base.Ui32(l3) {
				v31 = v28
			} else {
				v31 = l3
			}
			v34 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
			v35 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
			v62 = v35
			v63 = v34
			v64 = v31 - int32(_a_F_scalbnl_4)
		}
	} else {
		if int32(-16383) < l3 {
			v62 = l1
			v63 = l2
			v64 = l3
		} else {
			F___multf3(m, v8-int32(-64), l1, l2, int64(0), int64(32088147345014784))
			mBase = m.M
			v43 = *(*int64)(unsafe.Add(mBase, uint32(v8)+72))
			v44 = *(*int64)(unsafe.Add(mBase, uint32(v8)+64))
			if base.Ui32(int32(-32652)) < base.Ui32(l3) {
				v62 = v44
				v63 = v43
				v64 = l3 + int32(_a_F_scalbnl_5)
			} else {
				F___multf3(m, v8+int32(48), v44, v43, int64(0), int64(32088147345014784))
				mBase = m.M
				v54 = int32(-48920)
				if base.Ui32(l3) <= base.Ui32(v54) {
					v57 = v54
				} else {
					v57 = l3
				}
				v60 = *(*int64)(unsafe.Add(mBase, uint32(v8)+56))
				v61 = *(*int64)(unsafe.Add(mBase, uint32(v8)+48))
				v62 = v61
				v63 = v60
				v64 = v57 + int32(_a_F_scalbnl_6)
			}
		}
	}
	F___multf3(m, v8, v62, v63, int64(0), base.I64_extend_i32_u(v64+int32(_a_F_scalbnl_2))<<(uint(int64(48))%64))
	mBase = m.M
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v72
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v74
	m.G0 = v8 + int32(80)
	return
}
func F_scanner_errposition(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	if int32(0) <= l0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v7 = F_pg_mbstrlen_with_len(m, v6, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v11 = F_errposition(m, v7+int32(1))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		return
	}
}
func F_scanner_isspace(m *base.Module, l0 int32) int32 {
	return base.B2i32(l0 == int32(32)) | base.B2i32(base.Ui32((l0-int32(9))&int32(255)) < base.Ui32(int32(5)))
}
func F_scanner_yyerror(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
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
	var v53 int32
	_ = v53
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = v11 + v13
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		F_errcode(m, int32(16801924))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			if v15 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg(m, int32(_a_F_scanner_yyerror_0), v8)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					F_scanner_errposition(m, v30, l1)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_scanner_yyerror_1), int32(1207), int32(_a_F_scanner_yyerror_2))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v14
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
				F_errmsg(m, int32(_a_F_scanner_yyerror_3), v8+int32(16))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
					F_scanner_errposition(m, v46, l1)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_scanner_yyerror_1), int32(1215), int32(_a_F_scanner_yyerror_2))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
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
func F_scb_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v3 = F_geterrcode(m)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		if v3 == int32(67371461) {
			return
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v7 < int32(0) {
				return
			} else {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v13 = F_pg_mbstrlen_with_len(m, v12, v7)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					v17 = F_errposition(m, v13+int32(1))
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_scram_H(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	v6 = F_pg_cryptohash_create(m, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(_a_F_scram_H_0)
			return int32(-1)
		} else {
			v31 = F_pg_cryptohash_init(m, v6)
			mBase = m.M
			if v31 < int32(0) {
				if v6 == int32(0) {
					v54 = int32(_a_F_scram_H_0)
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
					if v46 == int32(1) {
						v49 = int32(_a_F_scram_H_1)
					} else {
						v49 = int32(_a_F_scram_H_2)
					}
					if v46 == int32(2) {
						v52 = int32(_a_F_scram_H_0)
					} else {
						v52 = v49
					}
					v54 = v52
				}
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = v54
				F_pg_cryptohash_free(m, v6)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					return int32(-1)
				}
			} else {
				v34 = F_pg_cryptohash_update(m, v6, l0, l2)
				mBase = m.M
				if v34 < int32(0) {
					if v6 == int32(0) {
						v54 = int32(_a_F_scram_H_0)
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
						if v46 == int32(1) {
							v49 = int32(_a_F_scram_H_1)
						} else {
							v49 = int32(_a_F_scram_H_2)
						}
						if v46 == int32(2) {
							v52 = int32(_a_F_scram_H_0)
						} else {
							v52 = v49
						}
						v54 = v52
					}
					*(*int32)(unsafe.Add(mBase, uint32(l4))) = v54
					F_pg_cryptohash_free(m, v6)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						return int32(-1)
					}
				} else {
					v37 = F_pg_cryptohash_final(m, v6, l3, l2)
					mBase = m.M
					if int32(0) <= v37 {
						F_pg_cryptohash_free(m, v6)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					} else {
						if v6 == int32(0) {
							v54 = int32(_a_F_scram_H_0)
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
							if v46 == int32(1) {
								v49 = int32(_a_F_scram_H_1)
							} else {
								v49 = int32(_a_F_scram_H_2)
							}
							if v46 == int32(2) {
								v52 = int32(_a_F_scram_H_0)
							} else {
								v52 = v49
							}
							v54 = v52
						}
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v54
						F_pg_cryptohash_free(m, v6)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							return int32(-1)
						}
					}
				}
			}
		}
	}
}
func F_select_best_admin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 != l1 {
		v13 = F_roles_is_member_of(m, l0, int32(1), l1, v7+int32(12))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			v18 = v17
			m.G0 = v7 + int32(16)
			return v18
		}
	} else {
		v18 = int32(0)
		m.G0 = v7 + int32(16)
		return v18
	}
}
func F_select_best_grantor(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int64
	_ = v131
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v18 = l1 << (uint(int64(32)) % 64)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_select_best_grantor[0]))
	if l0 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L8
	} else {
		goto L48
	}
L2:
	;
	m.G0 = v15 + int32(16)
	return
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v131
	goto L2
L4:
	;
	v22 = F_get_rolespec_oid(m, l0, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L6
L6:
	;
	if l3 != v20 {
		goto L30
	} else {
		goto L31
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v22
	v74 = F_aclmask_direct(m, l2, v22, l3, v18)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L8
	} else {
		goto L28
	}
L8:
	;
	return
L9:
	;
	if v20 == v22 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v25 = F_superuser_arg(m, v20)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	if v25 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v28 = int32(0)
	v30 = F_roles_is_member_of(m, v20, int32(1), v28, v28)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v32 = int32(0)
	if v30 == v32 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v70 == int32(0) {
		goto L1
	} else {
		goto L27
	}
L15:
	;
	v70 = int32(0)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v38 <= int32(0) {
		v64 = v32
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v70 = v64
	goto L14
L19:
	;
	v41 = int32(0)
	if v41 < v38 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v44 = v38
	goto L22
L21:
	;
	v44 = v41
	goto L22
L22:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v47 = int32(0)
	goto L23
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v45+v47<<(uint(int32(2))%32))))
	v56 = base.B2i32(v55 == v22)
	if v55 == v22 {
		v64 = v56
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v64 = v56
	goto L18
L25:
	;
	v58 = v47 + int32(1)
	if v58 != v44 {
		v47 = v58
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	goto L7
L28:
	;
	v131 = v74
	goto L3
L29:
	;
	v83 = int32(0)
	v85 = F_roles_is_member_of(m, v20, int32(1), v83, v83)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L8
	} else {
		goto L35
	}
L30:
	;
	v77 = F_superuser_arg(m, v20)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = l3
	v131 = v18
	goto L3
L33:
	;
	if v77 == int32(0) {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v20
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = int64(0)
	if v85 == int32(0) {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v92 <= int32(0) {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v102 = int32(0)
	v106 = int32(0)
	goto L38
L38:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v102<<(uint(int32(2))%32))))
	v113 = F_aclmask_direct(m, l2, v112, l3, v18)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L8
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v112
	v131 = v18
	goto L3
L40:
	;
	if v18 != v113 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if v113 == int64(0) {
		v124 = v106
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	goto L39
L44:
	;
	v126 = v102 + int32(1)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v126 < v127 {
		v102 = v126
		v106 = v124
		goto L38
	} else {
		goto L47
	}
L45:
	;
	v119 = base.I32_wrap_i64(base.I64_popcnt(v113))
	if v119 <= v106 {
		v124 = v106
		goto L44
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v112
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v113
	v124 = v119
	goto L44
L47:
	;
	goto L2
L48:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	v166 = F_GetUserNameFromId(m, v22, int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v166
	F_errmsg(m, int32(_a_F_select_best_grantor_0), v15)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_select_best_grantor_1), int32(_a_F_select_best_grantor_2), int32(_a_F_select_best_grantor_3))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L8
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_serial_errdetail_for_io_error(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14249(m, l0, int32(_a_F_serial_errdetail_for_io_error_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_set_authn_id(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_set_authn_id[0]))
	if v11 == int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_set_authn_id[1]))
		v17 = F_MemoryContextStrdup(m, v16, l1)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_set_authn_id[0])) = v17
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+380))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+296))
			*(*int32)(unsafe.Add(mBase, _c_F_set_authn_id[2])) = v22
			v25 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_set_authn_id[3])))
			if v25&int32(2) == int32(0) {
				m.G0 = v8 + int32(32)
				return
			} else {
				v32 = F_errstart(m, int32(15), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					if v32 == int32(0) {
						m.G0 = v8 + int32(32)
						return
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, _c_F_set_authn_id[0]))
						v39 = *(*int32)(unsafe.Add(mBase, _c_F_set_authn_id[2]))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v39<<(uint(int32(2))%32))+uint32(_c_F_set_authn_id[4])))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+380))
						v44 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v42
						*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v44
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v37
						F_errmsg(m, int32(_a_F_set_authn_id_0), v8)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_set_authn_id_1), int32(366), int32(_a_F_set_authn_id_2))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								m.G0 = v8 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return
		} else {
			F_errmsg(m, int32(_a_F_set_authn_id_3), int32(0))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
				v72 = *(*int32)(unsafe.Add(mBase, _c_F_set_authn_id[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v72
				F_errdetail_log(m, int32(_a_F_set_authn_id_4), v8+int32(16))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_set_authn_id_1), int32(353), int32(_a_F_set_authn_id_2))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
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
func F_set_debug_options(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	if l0 <= int32(0) {
		F_SetConfigOption(m, int32(_a_F_set_debug_options_0), int32(_a_F_set_debug_options_1), l1, l2)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			m.G0 = v8 + int32(80)
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
		v18 = v8 + int32(16)
		v20 = F_pg_sprintf(m, v18, int32(_a_F_set_debug_options_2), v8)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			F_SetConfigOption(m, int32(_a_F_set_debug_options_0), v18, l1, l2)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				if l1 == int32(1) {
					F_SetConfigOption(m, int32(_a_F_set_debug_options_3), int32(_a_F_set_debug_options_4), int32(1), l2)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						F_SetConfigOption(m, int32(_a_F_set_debug_options_5), int32(_a_F_set_debug_options_6), int32(1), l2)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							if l0 == int32(1) {
								m.G0 = v8 + int32(80)
								return
							} else {
								F_SetConfigOption(m, int32(_a_F_set_debug_options_7), int32(_a_F_set_debug_options_4), l1, l2)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									if base.Ui32(l0) < base.Ui32(int32(3)) {
										m.G0 = v8 + int32(80)
										return
									} else {
										F_SetConfigOption(m, int32(_a_F_set_debug_options_8), int32(_a_F_set_debug_options_6), l1, l2)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											F_SetConfigOption(m, int32(_a_F_set_debug_options_9), int32(_a_F_set_debug_options_6), l1, l2)
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return
											} else {
												if l0 == int32(3) {
													m.G0 = v8 + int32(80)
													return
												} else {
													F_SetConfigOption(m, int32(_a_F_set_debug_options_10), int32(_a_F_set_debug_options_6), l1, l2)
													mBase = m.M
													v58 = m.ExcPending
													if v58 != 0 {
														return
													} else {
														if base.Ui32(l0) < base.Ui32(int32(5)) {
															m.G0 = v8 + int32(80)
															return
														} else {
															F_SetConfigOption(m, int32(_a_F_set_debug_options_11), int32(_a_F_set_debug_options_6), l1, l2)
															mBase = m.M
															v64 = m.ExcPending
															if v64 != 0 {
																return
															} else {
																m.G0 = v8 + int32(80)
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
				} else {
					if l0 == int32(1) {
						m.G0 = v8 + int32(80)
						return
					} else {
						F_SetConfigOption(m, int32(_a_F_set_debug_options_7), int32(_a_F_set_debug_options_4), l1, l2)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							if base.Ui32(l0) < base.Ui32(int32(3)) {
								m.G0 = v8 + int32(80)
								return
							} else {
								F_SetConfigOption(m, int32(_a_F_set_debug_options_8), int32(_a_F_set_debug_options_6), l1, l2)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									F_SetConfigOption(m, int32(_a_F_set_debug_options_9), int32(_a_F_set_debug_options_6), l1, l2)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return
									} else {
										if l0 == int32(3) {
											m.G0 = v8 + int32(80)
											return
										} else {
											F_SetConfigOption(m, int32(_a_F_set_debug_options_10), int32(_a_F_set_debug_options_6), l1, l2)
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return
											} else {
												if base.Ui32(l0) < base.Ui32(int32(5)) {
													m.G0 = v8 + int32(80)
													return
												} else {
													F_SetConfigOption(m, int32(_a_F_set_debug_options_11), int32(_a_F_set_debug_options_6), l1, l2)
													mBase = m.M
													v64 = m.ExcPending
													if v64 != 0 {
														return
													} else {
														m.G0 = v8 + int32(80)
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
func F_set_deparse_context_plan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = l2
	F_set_deparse_plan(m, v6, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v12 == int32(337) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v15
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v17
		} else {
		}
		return l0
	}
}
func F_set_deparse_plan(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l1
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v13 - int32(338) {
	case 0:
		goto L4
	case 1:
		goto L3
	default:
		goto L2
	}
L1:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v23
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v22 = l1 + int32(52)
	goto L1
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v22 = v19
	goto L1
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v22 = v17
	goto L1
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
	v27 = v25
	goto L7
L6:
	;
	v27 = int32(0)
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v29 - int32(337) {
	case 0:
		goto L10
	default:
		goto L9
	case 14:
		goto L13
	case 18:
		goto L12
	case 20:
		goto L11
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v97
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v103 != int32(337) {
		goto L30
	} else {
		goto L31
	}
L9:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v97 = v94
	goto L8
L10:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v90 != int32(5) {
		v97 = l1
		goto L8
	} else {
		goto L28
	}
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v42 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34+v35<<(uint(int32(2))%32)-int32(4))))
	v97 = v41
	goto L8
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v97 = v32
	goto L8
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L24
	} else {
		goto L25
	}
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v45 <= int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v52 = int32(0)
	goto L17
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v48+v52<<(uint(int32(2))%32))))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v60 == int32(340) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L14
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+72))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v63 == v64 {
		v97 = v59
		goto L8
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v67 = v52 + int32(1)
	if v45 != v67 {
		v52 = v67
		goto L17
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	goto L18
L24:
	;
	return
L25:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v80
	F_errmsg_internal(m, int32(_a_F_set_deparse_plan_0), v10)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_set_deparse_plan_1), int32(_a_F_set_deparse_plan_2), int32(_a_F_set_deparse_plan_3))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v97 = v93
	goto L8
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v114
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v118 - int32(346) {
	case 0, 13:
		v122 = int32(96)
		goto L35
	default:
		v126 = int32(0)
		goto L34
	case 12:
		goto L36
	}
L30:
	;
	v110 = int32(0)
	if v97 == v110 {
		v114 = v110
		goto L29
	} else {
		goto L33
	}
L31:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v106 != int32(3) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+156))
	v114 = v109
	goto L29
L33:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v97)+44))
	v114 = v113
	goto L29
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v126
	m.G0 = v10 + int32(16)
	return
L35:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l1+v122)))
	v126 = v124
	goto L34
L36:
	;
	v122 = int32(104)
	goto L35
}
func F_set_plan_references(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	v3 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v13 = v12
	goto L3
L2:
	;
	v13 = v3
	goto L3
L3:
	;
	F_add_rtes_to_flat_rtable(m, l0, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v19 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v75 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v22 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v32 = v3
	goto L9
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v32<<(uint(int32(2))%32))))
	v40 = F_palloc(m, int32(36))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	goto L6
L11:
	;
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v40)+8)) = v42
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
	*(*int64)(unsafe.Add(mBase, uint32(v40))) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v46
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v38)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v40)+24)) = v48
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v40)+16)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = base.I32_wrap_i64(v42) + v13
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v55 + v13
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	v59 = F_lappend(m, v58, v40)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v59
	v63 = v32 + int32(1)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v63 < v64 {
		v32 = v63
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v121 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v78 <= int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v88 = int32(0)
	goto L17
L17:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91+v88<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+20)) = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v98 + v13
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v101 + v13
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
	v105 = F_lappend(m, v104, v95)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L14
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v105
	v109 = v88 + int32(1)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v109 < v110 {
		v88 = v109
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v124 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v139 = F_set_plan_refs(m, l0, l1, v13)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L32
	}
L24:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v127 = v125
	goto L26
L25:
	;
	v127 = int32(0)
	goto L26
L26:
	;
	v128 = F_palloc0(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+376)) = v128
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v131 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v134 = v132
	goto L30
L29:
	;
	v134 = int32(0)
	goto L30
L30:
	;
	v135 = F_palloc0(m, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+380)) = v135
	goto L23
L32:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v141 != int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	return v139
L34:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v144 == int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	if v147 <= int32(0) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v152 = int32(0)
	goto L37
L37:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160+v152))))
	if v162 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L33
L39:
	;
	v175 = v152 + int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	if v175 < v176 {
		v152 = v175
		goto L37
	} else {
		goto L42
	}
L40:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+380))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165+v152))))
	if v167 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v168+v152<<(uint(int32(2))%32)))) = int32(0)
	goto L39
L42:
	;
	goto L38
}
func F_shimBoolConsistentFn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v4 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+28)))
	v5 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+76)))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v7 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v8 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+72)))
	v9 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+64)))
	v10 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+68)))
	v11 = F_FunctionCall7Coll(m, v2, v3, v4, v5, v6, v7, v8, v9, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v17 = base.I32_wrap_i64(v11) & int32(255)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)) = uint8(base.B2i32(v17 == int32(2)))
		return base.B2i32(v17 != int32(0))
	}
}
func F_show_debug(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v9 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		if v9 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
			F_errmsg(m, int32(_a_F_show_debug_0), v5)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_show_debug_1), int32(164), int32(_a_F_show_debug_2))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					m.G0 = v5 + int32(16)
					return
				}
			}
		} else {
			m.G0 = v5 + int32(16)
			return
		}
	}
}
func F_show_limit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	v3 = *(*float64)(unsafe.Add(mBase, _c_F_show_limit[0]))
	return base.I64_extend_i32_s(base.I32_reinterpret_f32(base.F32_demote_f64(v3)))
}
func F_show_unix_socket_permissions(m *base.Module) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14375(m, int32(_a_F_show_unix_socket_permissions_0), int32(_a_F_show_unix_socket_permissions_1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_silly_cmp_tsvector(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = int32(2)
	v17 = int32(base.Ui32(v15) >> (uint(v16) % 32))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v20 = int32(base.Ui32(v18) >> (uint(v16) % 32))
	if base.Ui32(v17) < base.Ui32(v20) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-1)
L2:
	;
	goto L3
L3:
	;
	v24 = int32(1)
	if base.Ui32(v20) < base.Ui32(v17) {
		v255 = v24
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v255
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v26 < v27 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(-1)
L7:
	;
	goto L8
L8:
	;
	if v27 < v26 {
		v255 = v24
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v26 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	goto L12
L12:
	;
	v36 = int32(8)
	v37 = l1 + v36
	v38 = int32(2)
	v40 = v37 + v27<<(uint(v38)%32)
	v42 = l0 + v36
	v45 = v42 + v26<<(uint(v38)%32)
	v54 = v37
	v55 = v42
	v59 = int32(0)
	goto L13
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v61 = int32(1)
	v62 = v60 & v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v65 = v63 & v61
	if v62 != v65 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v255 = int32(0)
	goto L4
L15:
	;
	if base.Ui32(v65) < base.Ui32(v62) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v72 = int32(1)
	v74 = int32(2047)
	v75 = int32(base.Ui32(v63)>>(uint(v72)%32)) & v74
	v76 = int32(12)
	v77 = int32(base.Ui32(v63) >> (uint(v76) % 32))
	v79 = int32(base.Ui32(v60) >> (uint(v76) % 32))
	v83 = int32(base.Ui32(v60)>>(uint(v72)%32)) & v74
	if v83 != 0 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v70 = int32(-1)
	goto L20
L19:
	;
	v70 = int32(1)
	goto L20
L20:
	;
	return v70
L21:
	;
	if v62 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L22:
	;
	if v75 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	if v75 == int32(0) {
		goto L21
	} else {
		goto L54
	}
L25:
	;
	return int32(1)
L26:
	;
	goto L27
L27:
	;
	v88 = v79 + v45
	v89 = v77 + v40
	v90 = base.B2i32(base.Ui32(v83) < base.Ui32(v75))
	if base.Ui32(v83) < base.Ui32(v75) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v91 = v83
	goto L30
L29:
	;
	v91 = v75
	goto L30
L30:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v91) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	if v153 != 0 {
		v255 = v153
		goto L4
	} else {
		goto L49
	}
L32:
	;
	v153 = int32(0)
	goto L31
L33:
	;
	v127 = v122
	v128 = v123
	v129 = v124
	goto L43
L34:
	;
	if (v88|v89)&int32(3) != 0 {
		v122 = v88
		v123 = v89
		v124 = v91
		goto L33
	} else {
		goto L37
	}
L35:
	;
	v115 = v88
	v116 = v89
	v117 = v91
	goto L36
L36:
	;
	if v117 == int32(0) {
		goto L32
	} else {
		goto L42
	}
L37:
	;
	v99 = v88
	v100 = v89
	v101 = v91
	goto L38
L38:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v104 != v105 {
		v122 = v99
		v123 = v100
		v124 = v101
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v115 = v110
	v116 = v108
	v117 = v112
	goto L36
L40:
	;
	v107 = int32(4)
	v108 = v100 + v107
	v110 = v99 + v107
	v112 = v101 - v107
	if base.Ui32(int32(3)) < base.Ui32(v112) {
		v99 = v110
		v100 = v108
		v101 = v112
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v122 = v115
	v123 = v116
	v124 = v117
	goto L33
L43:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v132 == v133 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v153 = v132 - v133
	goto L31
L45:
	;
	v135 = int32(1)
	v140 = v129 - v135
	if v140 != 0 {
		v127 = v127 + v135
		v128 = v128 + v135
		v129 = v140
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	goto L32
L49:
	;
	if v75 == v83 {
		goto L21
	} else {
		goto L50
	}
L50:
	;
	if base.Ui32(v83) < base.Ui32(v75) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v157 = int32(-1)
	goto L53
L52:
	;
	v157 = int32(1)
	goto L53
L53:
	;
	return v157
L54:
	;
	return int32(-1)
L55:
	;
	v244 = int32(4)
	v250 = v59 + int32(1)
	if v250 != v26 {
		v54 = v54 + v244
		v55 = v55 + v244
		v59 = v250
		goto L13
	} else {
		goto L80
	}
L56:
	;
	v168 = int32(1)
	v170 = int32(_a_F_silly_cmp_tsvector_0)
	v172 = v40 + (v75+v77+v168)&v170
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172))))
	v179 = v45 + (v83+v79+v168)&v170
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v179))))
	if v173 == v180 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v188 = v179
	v189 = v172
	v190 = int32(0)
	goto L65
L58:
	;
	if v180 != 0 {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if base.Ui32(v173) < base.Ui32(v180) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L55
L62:
	;
	v186 = int32(-1)
	goto L64
L63:
	;
	v186 = int32(1)
	goto L64
L64:
	;
	return v186
L65:
	;
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+2)))
	v203 = int32(_a_F_silly_cmp_tsvector_1)
	v204 = v202 & v203
	v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189)+2)))
	v207 = v205 & v203
	if v204 != v207 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if base.Ui32(v217) < base.Ui32(v215) {
		goto L77
	} else {
		goto L78
	}
L67:
	;
	if base.Ui32(v207) < base.Ui32(v204) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	v214 = int32(14)
	v215 = int32(base.Ui32(v202) >> (uint(v214) % 32))
	v217 = int32(base.Ui32(v205) >> (uint(v214) % 32))
	if v215 == v217 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v212 = int32(-1)
	goto L72
L71:
	;
	v212 = int32(1)
	goto L72
L72:
	;
	return v212
L73:
	;
	v219 = int32(2)
	v224 = v190 + int32(1)
	if v224 == v180 {
		goto L55
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	goto L66
L76:
	;
	v188 = v188 + v219
	v189 = v189 + v219
	v190 = v224
	goto L65
L77:
	;
	v229 = int32(-1)
	goto L79
L78:
	;
	v229 = int32(1)
	goto L79
L79:
	;
	v255 = v229
	goto L4
L80:
	;
	goto L14
}
func F_similar_escape(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v2 == int32(1) {
		v5 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v5)
		return int64(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v10 = F_pg_detoast_datum_packed(m, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
			if v14 != 0 {
				v19 = int32(0)
				v20 = F_similar_escape_internal(m, v10, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v20)
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v17 = F_pg_detoast_datum_packed(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					v19 = v17
					v20 = F_similar_escape_internal(m, v10, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(v20)
					}
				}
			}
		}
	}
}
func F_similar_escape_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
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
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v366 int32
	_ = v366
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	v3 = int32(0)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v17 = int32(1)
	v18 = v16 & v17
	if v16 == v17 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l1 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v24 == int32(18) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v35 = int32(1)
	if v18 != 0 {
		v45 = int32(base.Ui32(v16)>>(uint(v35)%32)) - v35
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v27 = int32(16)
	goto L7
L6:
	;
	v27 = int32(0)
	goto L7
L7:
	;
	if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v34 = int32(4)
	goto L10
L9:
	;
	v34 = v27
	goto L10
L10:
	;
	v45 = v34
	goto L1
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L34
	} else {
		goto L115
	}
L13:
	;
	v106 = F_palloc(m, v45*int32(3)+int32(27))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L34
	} else {
		goto L37
	}
L14:
	;
	v100 = int32(1)
	v101 = int32(_a_F_similar_escape_internal_0)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v50 = int32(4)
	v51 = int32(1)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v55 = v53 & v51
	if v55 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v56 = v51
	goto L19
L18:
	;
	v56 = v50
	goto L19
L19:
	;
	v57 = l1 + v56
	if v53 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v91 = F_pg_mbstrlen_with_len(m, v57, v90)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L34
	} else {
		goto L35
	}
L21:
	;
	if v82 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if base.Ui32((v60-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v90 = v50
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v72 = int32(1)
	if v55 != 0 {
		v82 = int32(base.Ui32(v53)>>(uint(v72)%32)) - v72
		goto L21
	} else {
		goto L29
	}
L25:
	;
	if v60 == int32(18) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v71 = int32(16)
	goto L28
L27:
	;
	v71 = int32(0)
	goto L28
L28:
	;
	v82 = v71
	goto L21
L29:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v82 = int32(base.Ui32(v76)>>(uint(int32(2))%32)) - int32(4)
	goto L21
L30:
	;
	v85 = int32(0)
	v100 = v85
	v101 = v85
	goto L13
L31:
	;
	goto L32
L32:
	;
	if v82 < int32(2) {
		v100 = v82
		v101 = v57
		goto L13
	} else {
		goto L33
	}
L33:
	;
	v90 = v82
	goto L20
L34:
	;
	return int32(0)
L35:
	;
	if int32(2) <= v91 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	v100 = v90
	v101 = v57
	goto L13
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = int32(977217630)
	v111 = v106 + int32(8)
	if v45 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v366 = int32(_a_F_similar_escape_internal_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v351))) = uint16(v366)
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = (v351-v106)<<(uint(int32(2))%32) + int32(8)
	return v106
L39:
	;
	v351 = v111
	goto L38
L40:
	;
	goto L41
L41:
	;
	if v18 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v116 = int32(1)
	goto L44
L43:
	;
	v116 = int32(4)
	goto L44
L44:
	;
	v117 = l0 + v116
	v122 = v117
	v123 = v111
	v126 = v3
	v127 = v45
	v130 = v3
	v132 = v3
	v133 = v3
	goto L45
L45:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v100 < int32(2) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	v351 = v340
	goto L38
L47:
	;
	if int32(0) < v344 {
		v122 = v348
		v123 = v340
		v126 = v343
		v127 = v344
		v130 = v345
		v132 = v346
		v133 = v347
		goto L45
	} else {
		goto L114
	}
L48:
	;
	v340 = v337
	v343 = v336
	v344 = v127 - v137
	v345 = v130
	v346 = v132
	v347 = v133
	v348 = v122 + v137
	goto L47
L49:
	;
	if v137 != 0 {
		goto L111
	} else {
		goto L112
	}
L50:
	;
	if v126 != 0 {
		goto L82
	} else {
		goto L83
	}
L51:
	;
	v137 = F_pg_mblen_range(m, v122, v117+v45)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L34
	} else {
		goto L52
	}
L52:
	;
	if v137 < int32(2) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	if v126 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v141 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v141)
	v331 = v123 + int32(1)
	goto L49
L55:
	;
	goto L56
L56:
	;
	if base.B2i32(v101 == int32(0))|base.B2i32(v137 != v100) != 0 {
		v331 = v123
		goto L49
	} else {
		goto L57
	}
L57:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v100) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	if v210 != 0 {
		v331 = v123
		goto L49
	} else {
		goto L76
	}
L59:
	;
	v210 = int32(0)
	goto L58
L60:
	;
	v184 = v179
	v185 = v180
	v186 = v181
	goto L70
L61:
	;
	if (v101|v122)&int32(3) != 0 {
		v179 = v101
		v180 = v122
		v181 = v100
		goto L60
	} else {
		goto L64
	}
L62:
	;
	v172 = v101
	v173 = v122
	v174 = v100
	goto L63
L63:
	;
	if v174 == int32(0) {
		goto L59
	} else {
		goto L69
	}
L64:
	;
	v156 = v101
	v157 = v122
	v158 = v100
	goto L65
L65:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	if v161 != v162 {
		v179 = v156
		v180 = v157
		v181 = v158
		goto L60
	} else {
		goto L67
	}
L66:
	;
	v172 = v167
	v173 = v165
	v174 = v169
	goto L63
L67:
	;
	v164 = int32(4)
	v165 = v157 + v164
	v167 = v156 + v164
	v169 = v158 - v164
	if base.Ui32(int32(3)) < base.Ui32(v169) {
		v156 = v167
		v157 = v165
		v158 = v169
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v179 = v172
	v180 = v173
	v181 = v174
	goto L60
L70:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if v189 == v190 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v210 = v189 - v190
	goto L58
L72:
	;
	v192 = int32(1)
	v197 = v186 - v192
	if v197 != 0 {
		v184 = v184 + v192
		v185 = v185 + v192
		v186 = v197
		goto L70
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	goto L71
L75:
	;
	goto L59
L76:
	;
	v336 = int32(1)
	v337 = v123
	goto L48
L77:
	;
	v327 = int32(1)
	v340 = v322
	v343 = v324
	v344 = v127 - v327
	v345 = v325
	v346 = v326
	v347 = v323
	v348 = v122 + v327
	goto L47
L78:
	;
	v322 = v317
	v323 = v133
	v324 = int32(0)
	v325 = v319
	v326 = v320
	goto L77
L79:
	;
	v291 = v123 + int32(1)
	switch v136 - int32(36) {
	case 0, 10, 56, 58:
		goto L106
	case 1:
		goto L109
	default:
		goto L105
	case 4:
		goto L107
	case 55:
		goto L110
	case 59:
		goto L108
	}
L80:
	;
	v271 = v123 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v136)
	if base.B2i32(v136 != int32(93))|base.B2i32(v133 < int32(3)) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L81:
	;
	v322 = v123 + int32(2)
	v323 = int32(3)
	v324 = int32(0)
	v325 = v130
	v326 = v132
	goto L77
L82:
	;
	v215 = int32(0)
	if base.B2i32(v136 != int32(34))|base.B2i32(v215 < v130) == v215 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	if v101 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L85:
	;
	switch v132 {
	case 0:
		goto L88
	case 1:
		goto L90
	default:
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)) = uint8(v136)
	v251 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v251)
	goto L81
L88:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v123))) = int64(2900174335198198569)
	v317 = v123 + int32(8)
	v319 = v130
	v320 = v132 + int32(1)
	goto L78
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L34
	} else {
		goto L91
	}
L90:
	;
	v220 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+8)) = uint8(v220)
	*(*int64)(unsafe.Add(mBase, uint32(v123))) = int64(4551025073606196009)
	v317 = v123 + int32(9)
	v319 = v130
	v320 = v132 + int32(1)
	goto L78
L91:
	;
	F_errcode(m, int32(318767234))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L34
	} else {
		goto L92
	}
L92:
	;
	F_errmsg(m, int32(_a_F_similar_escape_internal_2), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L34
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_similar_escape_internal_3), int32(951), int32(_a_F_similar_escape_internal_4))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L34
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	if v130 <= int32(0) {
		goto L79
	} else {
		goto L98
	}
L96:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v136 != v256 {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v322 = v123
	v323 = v133
	v324 = int32(1)
	v325 = v130
	v326 = v132
	goto L77
L98:
	;
	if v136 != int32(92) {
		goto L80
	} else {
		goto L99
	}
L99:
	;
	v263 = int32(_a_F_similar_escape_internal_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v123))) = uint16(v263)
	goto L81
L100:
	;
	v317 = v271
	v319 = v130 - int32(1)
	v320 = v132
	goto L78
L101:
	;
	goto L102
L102:
	;
	v282 = int32(3)
	v283 = int32(0)
	switch v136 - int32(91) {
	case 0:
		goto L104
	default:
		v322 = v271
		v323 = v282
		v324 = v283
		v325 = v130
		v326 = v132
		goto L77
	case 3:
		goto L103
	}
L103:
	;
	v322 = v271
	v323 = v133 + int32(1)
	v324 = v283
	v325 = v130
	v326 = v132
	goto L77
L104:
	;
	v322 = v271
	v323 = v282
	v324 = v283
	v325 = v130 + int32(1)
	v326 = v132
	goto L77
L105:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v136)
	v317 = v291
	v319 = v130
	v320 = v132
	goto L78
L106:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)) = uint8(v136)
	v312 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v312)
	v317 = v123 + int32(2)
	v319 = v130
	v320 = v132
	goto L78
L107:
	;
	v305 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+2)) = uint8(v305)
	v307 = int32(_a_F_similar_escape_internal_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v123))) = uint16(v307)
	v317 = v123 + int32(3)
	v319 = v130
	v320 = v132
	goto L78
L108:
	;
	v303 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v303)
	v317 = v291
	v319 = v130
	v320 = v132
	goto L78
L109:
	;
	v299 = int32(_a_F_similar_escape_internal_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v123))) = uint16(v299)
	v317 = v123 + int32(2)
	v319 = v130
	v320 = v132
	goto L78
L110:
	;
	v294 = int32(91)
	*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v294)
	v296 = int32(1)
	v322 = v291
	v323 = v296
	v324 = int32(0)
	v325 = v296
	v326 = v132
	goto L77
L111:
	;
	base.MemoryCopy(m, v331, v122, v137)
	goto L113
L112:
	;
	goto L113
L113:
	;
	v336 = int32(0)
	v337 = v331 + v137
	goto L48
L114:
	;
	goto L46
L115:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L34
	} else {
		goto L116
	}
L116:
	;
	F_errmsg(m, int32(_a_F_similar_escape_internal_8), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L34
	} else {
		goto L117
	}
L117:
	;
	F_errhint(m, int32(_a_F_similar_escape_internal_9), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L34
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_similar_escape_internal_3), int32(805), int32(_a_F_similar_escape_internal_4))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L34
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_simplify_EXISTS_query(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v7 != int32(1) {
		v95 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v95
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+144))
	if v10 != 0 {
		v95 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)))
	if v11 != 0 {
		v95 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v12 != 0 {
		v95 = v3
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+37)))
	if v13 != 0 {
		v95 = v3
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	if v14 != 0 {
		v95 = v3
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+42)))
	if v15 != 0 {
		v95 = v3
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	if v16 != 0 {
		v95 = v3
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	if v17 != 0 {
		v95 = v3
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	if v18 != 0 {
		v95 = v3
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v19 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v20 = F_eval_const_expressions(m, l0, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v37 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+124)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(l1)+116)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+100)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v37
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+40)) = uint8(v37)
	v47 = int32(1)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)))
	if v48 != v47 {
		v95 = v47
		goto L1
	} else {
		goto L22
	}
L15:
	;
	return int32(0)
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+132)) = v20
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v25 != int32(7) {
		v95 = v3
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+32)))
	if v28 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v20)+24))
	if v31 <= int64(0) {
		v95 = v3
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+132)) = int32(0)
	goto L14
L21:
	;
	goto L20
L22:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v51 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v89 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v89)
	v95 = v47
	goto L1
L24:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v54 <= int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v57 = int32(0)
	if v57 < v54 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v61 = v54
	goto L28
L27:
	;
	v61 = v57
	goto L28
L28:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v63 = v57
	goto L29
L29:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63<<(uint(int32(2))%32))))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	if v73 != int32(9) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+120)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = int32(8)
	goto L23
L31:
	;
	v77 = v63 + int32(1)
	if v61 != v77 {
		v63 = v77
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	goto L30
L34:
	;
	goto L23
}
func F_slist_delete(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	v4 = l0
	goto L2
L1:
	;
	return
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if v7 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = v11
	goto L1
L4:
	;
	if v7 != l1 {
		v4 = v7
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L3
}
func F_slotsync_worker_disconnect(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_disconnect[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+64))
	m.T0[v6].(func(*base.Module, int32))(m, base.I32_wrap_i64(l1))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_smgrclose(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	v2 = int32(_a_F_smgrclose_0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_smgrclose[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrclose[0])) = v4 + int32(1)
	F_mdclose(m, l0, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(-1)
		F_mdclose(m, l0, int32(1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(-1)
			F_mdclose(m, l0, int32(2))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
				F_mdclose(m, l0, int32(3))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v26
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v26
					v30 = int32(_a_F_smgrclose_0)
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_smgrclose[0]))
					*(*int32)(unsafe.Add(mBase, _c_F_smgrclose[0])) = v32 - int32(1)
					return
				}
			}
		}
	}
}
func F_smgrexists(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v3 = int32(_a_F_smgrexists_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_smgrexists[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrexists[0])) = v5 + int32(1)
	v9 = F_mdexists(m, l0, l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(_a_F_smgrexists_0)
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_smgrexists[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_smgrexists[0])) = v15 - int32(1)
		return v9
	}
}
func F_smgrextend(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v6 = int32(_a_F_smgrextend_0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_smgrextend[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrextend[0])) = v8 + int32(1)
	F_mdextend(m, l0, l1, l2, l3, l4)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v18 = l0 + l1<<(uint(int32(2))%32) + int32(20)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
		if v22 != l2 {
			v24 = int32(-1)
		} else {
			v24 = l2 + int32(1)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v18))) = v24
		v26 = int32(_a_F_smgrextend_0)
		v28 = *(*int32)(unsafe.Add(mBase, _c_F_smgrextend[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_smgrextend[0])) = v28 - int32(1)
		return
	}
}
func F_smgrreleaseall(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_smgrreleaseall[0]))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = int32(_a_F_smgrreleaseall_0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_smgrreleaseall[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrreleaseall[1])) = v12 + int32(1)
	v17 = v6 + int32(12)
	F_hash_seq_init(m, v17, v9)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v6 + int32(32)
	return
L4:
	;
	return
L5:
	;
	v20 = F_hash_seq_search(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v22 = v20
	goto L10
L8:
	;
	goto L9
L9:
	;
	v66 = int32(_a_F_smgrreleaseall_0)
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_smgrreleaseall[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrreleaseall[1])) = v68 - int32(1)
	goto L3
L10:
	;
	v25 = int32(_a_F_smgrreleaseall_0)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_smgrreleaseall[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrreleaseall[1])) = v27 + int32(1)
	F_mdclose(m, v22, int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = int32(-1)
	F_mdclose(m, v22, int32(1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = int32(-1)
	F_mdclose(m, v22, int32(2))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = int32(-1)
	F_mdclose(m, v22, int32(3))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v49 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v49
	v53 = int32(_a_F_smgrreleaseall_0)
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_smgrreleaseall[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrreleaseall[1])) = v55 - int32(1)
	v61 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if v61 != 0 {
		v22 = v61
		goto L10
	} else {
		goto L17
	}
L17:
	;
	goto L11
}
func F_smgrtruncate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int64
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int64
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int64
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int64
	_ = v187
	var v190 int64
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v309 int32
	_ = v309
	var v327 int32
	_ = v327
	var v338 int32
	_ = v338
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int64
	_ = v366
	var v368 int64
	_ = v368
	var v393 int64
	_ = v393
	var v403 int64
	_ = v403
	var v434 int32
	_ = v434
	var v435 int64
	_ = v435
	var v438 int64
	_ = v438
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v488 int64
	_ = v488
	var v490 int64
	_ = v490
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v532 int32
	_ = v532
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v591 int64
	_ = v591
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v644 int64
	_ = v644
	var v646 int64
	_ = v646
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v669 int32
	_ = v669
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v714 int32
	_ = v714
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v752 int64
	_ = v752
	var v755 int64
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v891 int32
	_ = v891
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	v23 = m.G0
	v25 = v23 - int32(16)
	m.G0 = v25
	v27 = m.G0
	v29 = v27 - int32(80)
	m.G0 = v29
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+48)) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+40)) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	if v35 == int32(-1) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v29 + int32(80)
	v644 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v644
	v646 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v646
	F_CacheInvalidateSmgr(m, v25)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L22
	} else {
		goto L90
	}
L2:
	;
	if v309 <= int32(0) {
		goto L1
	} else {
		goto L45
	}
L3:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[0]))
	if base.Ui32(int32(63)) <= base.Ui32(v298+int32(31)) {
		goto L1
	} else {
		goto L44
	}
L4:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[0]))
	v309 = v296
	goto L2
L5:
	;
	v38 = int32(0)
	if l2 <= v38 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[1]))
	if v35 != v147 {
		goto L1
	} else {
		goto L25
	}
L8:
	;
	v48 = v38
	v61 = int64(0)
	goto L9
L9:
	;
	v64 = v48 << (uint(int32(2)) % 32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1+v64)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrtruncate[2])))
	if v71 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[0]))
	v101 = base.I32_div_s(v99, int32(32))
	if base.B2i32(base.I32_wrap_i64(v90) == int32(-1))|base.B2i32(base.Ui64(base.I64_extend_i32_s(v101)) <= base.Ui64(v90)) != 0 {
		v309 = v99
		goto L2
	} else {
		goto L19
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64+(v29+int32(28))))) = v82
	if v82 == int32(-1) {
		goto L4
	} else {
		goto L17
	}
L12:
	;
	goto L11
L13:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0+v69<<(uint(int32(2))%32))+20))
	if v77 != int32(-1) {
		v82 = v77
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v82 = int32(-1)
	goto L12
L16:
	;
	goto L15
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l4+v64)))
	v90 = v61 + base.I64_extend_i32_u(v82-v87)
	v92 = v48 + int32(1)
	if v92 != l2 {
		v48 = v92
		v61 = v90
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	v112 = int32(0)
	goto L20
L20:
	;
	v128 = v112 << (uint(int32(2)) % 32)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l4+v128)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1+v128)))
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v29)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v135
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(28)+v128)))
	F_FindAndDropRelationBuffers(m, v29, v132, v140, v130)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L1
L22:
	;
	return
L23:
	;
	v144 = v112 + int32(1)
	if v144 != l2 {
		v112 = v144
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v149
	v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v151
	v153 = int32(0)
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[3]))
	if v153 < v155 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v159 = v29 + int32(16)
	v167 = v153
	goto L29
L27:
	;
	goto L28
L28:
	;
	goto L1
L29:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[4]))
	v186 = v183 + v167*int32(56)
	v187 = int64(0)
	v190 = base.AtomicRmwCmpxchg64(m, v186, int32(24), v187, v187)
	if v190&int64(33554432) == v187 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	v269 = v167 + int32(1)
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[3]))
	if v269 < v271 {
		v167 = v269
		goto L29
	} else {
		goto L43
	}
L32:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v195 != v196 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if v198 != v199 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	if base.B2i32(v201 != v202)|base.B2i32(l2 <= int32(0)) != 0 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
	v215 = int32(0)
	goto L36
L36:
	;
	v232 = v215 << (uint(int32(2)) % 32)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l1+v232)))
	if v207 != v234 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L31
L38:
	;
	v244 = v215 + int32(1)
	if v244 != l2 {
		v215 = v244
		goto L36
	} else {
		goto L42
	}
L39:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v186)+16))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l4+v232)))
	if base.Ui32(v236) < base.Ui32(v238) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	F_InvalidateLocalBuffer(m, v186, int32(1))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L22
	} else {
		goto L41
	}
L41:
	;
	goto L31
L42:
	;
	goto L37
L43:
	;
	goto L30
L44:
	;
	v309 = v298
	goto L2
L45:
	;
	v327 = int32(0)
	v338 = v327
	goto L46
L46:
	;
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[5]))
	v356 = v353 + v338*int32(56)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	if v357 != v358 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L1
L48:
	;
	v615 = v338 + int32(1)
	v617 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[0]))
	if v615 < v617 {
		v338 = v615
		goto L46
	} else {
		goto L89
	}
L49:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	if v360 != v361 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v356)+8))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	if v363 != v364 {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v366 = int64(4194304)
	v368 = base.AtomicRmwOr64(m, v356, int32(24), v366)
	if v368&v366 != int64(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v393 = v368
	goto L55
L53:
	;
	goto L54
L54:
	;
	if base.B2i32(l2 <= v327) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+76)) = int32(_a_F_smgrtruncate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = int32(_a_F_smgrtruncate_1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = int32(_a_F_smgrtruncate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	v403 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+56)) = v403
	if v393&int64(4194304) != v403 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L54
L57:
	;
	goto L60
L58:
	;
	goto L59
L59:
	;
	v468 = int32(_a_F_smgrtruncate_3)
	v469 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[6]))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(56))+8))
	if v471 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L60:
	;
	F_perform_spin_delay(m, v29+int32(56))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L22
	} else {
		goto L62
	}
L61:
	;
	goto L59
L62:
	;
	v435 = int64(0)
	v438 = base.AtomicRmwCmpxchg64(m, v356, int32(24), v435, v435)
	if v438&int64(4194304) != v435 {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v488 = int64(4194304)
	v490 = base.AtomicRmwOr64(m, v356, int32(24), v488)
	if v490&v488 != int64(0) {
		v393 = v490
		goto L55
	} else {
		goto L75
	}
L65:
	;
	goto L64
L66:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[6])) = v486
	goto L65
L67:
	;
	if int32(999) < v469 {
		goto L65
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if v469 < int32(11) {
		goto L65
	} else {
		goto L74
	}
L70:
	;
	v476 = int32(900)
	if v476 <= v469 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v479 = v476
	goto L73
L72:
	;
	v479 = v469
	goto L73
L73:
	;
	v486 = v479 + int32(100)
	goto L66
L74:
	;
	v486 = v469 - int32(1)
	goto L66
L75:
	;
	goto L56
L76:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	v532 = int32(0)
	goto L79
L77:
	;
	goto L78
L78:
	;
	v591 = base.AtomicRmwSub64(m, v356, int32(24), int64(4194304))
	goto L48
L79:
	;
	if v522 != v523 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	goto L78
L81:
	;
	v565 = v532 + int32(1)
	if v565 != l2 {
		v532 = v565
		goto L79
	} else {
		goto L88
	}
L82:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	if v547 != v521 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v356)+8))
	if v549 != v520 {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v356)+12))
	v553 = v532 << (uint(int32(2)) % 32)
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l1+v553)))
	if v551 != v555 {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v356)+16))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l4+v553)))
	if base.Ui32(v557) < base.Ui32(v559) {
		goto L81
	} else {
		goto L86
	}
L86:
	;
	F_InvalidateBuffer(m, v356)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L22
	} else {
		goto L87
	}
L87:
	;
	goto L48
L88:
	;
	goto L80
L89:
	;
	goto L47
L90:
	;
	if int32(0) < l2 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v653 = l0 + int32(20)
	v669 = int32(0)
	goto L94
L92:
	;
	goto L93
L93:
	;
	m.G0 = v25 + int32(16)
	return
L94:
	;
	v676 = int32(2)
	v677 = v669 << (uint(v676) % 32)
	v678 = l1 + v677
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v678)))
	*(*int32)(unsafe.Add(mBase, uint32(v653+v679<<(uint(v676)%32)))) = int32(-1)
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v678)))
	v686 = l3 + v677
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v686)))
	v688 = l4 + v677
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v688)))
	v690 = m.G0
	v692 = v690 - int32(112)
	m.G0 = v692
	if base.Ui32(v687) < base.Ui32(v689) {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	goto L93
L96:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v678)))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v688)))
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v686)))
	if base.Ui32(v901) < base.Ui32(v902) {
		goto L152
	} else {
		goto L153
	}
L97:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L22
	} else {
		goto L147
	}
L98:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L22
	} else {
		goto L142
	}
L99:
	;
	m.G0 = v692 + int32(112)
	goto L96
L100:
	;
	v696 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrtruncate[2])))
	if v696 != 0 {
		goto L99
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	if v687 == v689 {
		goto L99
	} else {
		goto L108
	}
L103:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L22
	} else {
		goto L104
	}
L104:
	;
	v702 = v692 + int32(40)
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_GetRelationPath(m, v702, v703, v704, v705, v706, v685)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L22
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v692)+8)) = v687
	*(*int32)(unsafe.Add(mBase, uint32(v692)+4)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v692))) = v702
	F_errmsg(m, int32(_a_F_smgrtruncate_4), v692)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L22
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_smgrtruncate_5), int32(1316), int32(_a_F_smgrtruncate_6))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L22
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	v723 = l0 + v685<<(uint(int32(2))%32)
	v725 = v723 + int32(40)
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v725)))
	if v726 <= int32(0) {
		goto L99
	} else {
		goto L109
	}
L109:
	;
	v730 = v723 + int32(56)
	v752 = base.I64_extend_i32_u(v726)
	goto L110
L110:
	;
	v755 = v752 - int64(1)
	v756 = base.I32_wrap_i64(v755)
	v758 = v756 << (uint(int32(3)) % 32)
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v730)))
	v760 = v758 + v759
	v762 = v756 << (uint(int32(17)) % 32)
	if base.Ui32(v689) < base.Ui32(v762) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L99
L112:
	;
	if base.Ui64(int64(1)) < base.Ui64(v752) {
		v752 = v755
		goto L110
	} else {
		goto L141
	}
L113:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v760)))
	v767 = F_FileTruncate(m, v764, int64(0), int32(167772185))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L22
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if base.Ui32(v762+int32(_a_F_smgrtruncate_7)) <= base.Ui32(v689) {
		goto L99
	} else {
		goto L136
	}
L116:
	;
	if v767 < int32(0) {
		goto L98
	} else {
		goto L117
	}
L117:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v771 == int32(-1) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	F_register_dirty_segment(m, l0, v685, v760)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L22
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v760)))
	F_FileClose(m, v776)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L22
	} else {
		goto L122
	}
L121:
	;
	goto L120
L122:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v725)))
	if v755 == int64(0) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v725))) = v756
	goto L112
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v730))) = v799
	goto L123
L125:
	;
	if v779 <= int32(0) {
		goto L123
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	if v779 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v730)))
	F_pfree(m, v784)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L22
	} else {
		goto L129
	}
L129:
	;
	v799 = int32(0)
	goto L124
L130:
	;
	v791 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[7]))
	v792 = F_MemoryContextAlloc(m, v791, v758)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L22
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	if v755 <= base.I64_extend_i32_s(v779) {
		goto L123
	} else {
		goto L134
	}
L133:
	;
	v799 = v792
	goto L124
L134:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v730)))
	v797 = F_repalloc(m, v796, v758)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L22
	} else {
		goto L135
	}
L135:
	;
	v799 = v797
	goto L124
L136:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v760)))
	v811 = F_FileTruncate(m, v805, base.I64_extend_i32_u(v689-v762)<<(uint(int64(13))%64), int32(167772185))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L22
	} else {
		goto L137
	}
L137:
	;
	if v811 < int32(0) {
		goto L97
	} else {
		goto L138
	}
L138:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v815 != int32(-1) {
		goto L112
	} else {
		goto L139
	}
L139:
	;
	F_register_dirty_segment(m, l0, v685, v760)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L22
	} else {
		goto L140
	}
L140:
	;
	goto L112
L141:
	;
	goto L111
L142:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L22
	} else {
		goto L143
	}
L143:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v760)))
	v856 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[8]))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v856+v854*int32(48))+32))
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v692)+16)) = v860
	F_errmsg(m, int32(_a_F_smgrtruncate_8), v692+int32(16))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L22
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_smgrtruncate_5), int32(1344), int32(_a_F_smgrtruncate_6))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L22
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L22
	} else {
		goto L148
	}
L148:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v760)))
	v880 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[8]))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v880+v878*int32(48))+32))
	goto L149
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v692)+36)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v692)+32)) = v884
	F_errmsg(m, int32(_a_F_smgrtruncate_9), v692+int32(32))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L22
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_smgrtruncate_5), int32(1371), int32(_a_F_smgrtruncate_6))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L22
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
	v904 = v901
	goto L154
L153:
	;
	v904 = v902
	goto L154
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v653+v897<<(uint(int32(2))%32)))) = v904
	v907 = v669 + int32(1)
	if v907 != l2 {
		v669 = v907
		goto L94
	} else {
		goto L155
	}
L155:
	;
	goto L95
}
func F_sortouts(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if int32(2) <= v10 {
		v13 = int32(2)
		v16 = F_palloc_extended(m, v10<<(uint(v13)%32), v13)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 == int32(0) {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = int32(101)
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
				if v24 != 0 {
					v26 = v24
				} else {
					v26 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v26
				return
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				if v28 != 0 {
					v29 = v28
					v33 = int32(0)
					for {
						*(*int32)(unsafe.Add(mBase, uint32(v16+v33<<(uint(int32(2))%32)))) = v29
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
						if v44 != 0 {
							v29 = v44
							v33 = v33 + int32(1)
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				F_pg_qsort(m, v16, v10, int32(4), int32(1047))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v58
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v60
					if v10 == int32(2) {
						v140 = int32(1)
					} else {
						v67 = int32(1)
						v69 = v10 - v67
						if v10 != int32(3) {
							v80 = int32(0)
							v83 = v67
							for {
								v88 = int32(2)
								v90 = v16 + v83<<(uint(v88)%32)
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
								v92 = int32(4)
								v93 = v90 + v92
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
								*(*int32)(unsafe.Add(mBase, uint32(v91)+16)) = v94
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v90-v92)))
								*(*int32)(unsafe.Add(mBase, uint32(v91)+20)) = v98
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
								v102 = v83 + v88
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v16+v102<<(uint(v88)%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v100)+16)) = v106
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
								*(*int32)(unsafe.Add(mBase, uint32(v100)+20)) = v108
								if v80 != v10&int32(2147483646)-int32(4) {
									v80 = v80 + v88
									v83 = v102
									continue
								} else {
									break
								}
								break
							}
							if v10&int32(1) == int32(0) {
								v140 = v69
							} else {
								v119 = v102
								v126 = v16 + v119<<(uint(int32(2))%32)
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v127)+16)) = v128
								v132 = *(*int32)(unsafe.Add(mBase, uint32(v126-int32(4))))
								*(*int32)(unsafe.Add(mBase, uint32(v127)+20)) = v132
								v140 = v69
							}
						} else {
							v119 = v67
							v126 = v16 + v119<<(uint(int32(2))%32)
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v127)+16)) = v128
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v126-int32(4))))
							*(*int32)(unsafe.Add(mBase, uint32(v127)+20)) = v132
							v140 = v69
						}
					}
					v145 = v16 + v140<<(uint(int32(2))%32)
					v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
					*(*int32)(unsafe.Add(mBase, uint32(v146)+16)) = int32(0)
					v151 = *(*int32)(unsafe.Add(mBase, uint32(v145-int32(4))))
					*(*int32)(unsafe.Add(mBase, uint32(v146)+20)) = v151
					F_pfree(m, v16)
					mBase = m.M
					v154 = m.ExcPending
					if v154 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		return
	}
}
func F_spanish_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v24 int32
	_ = v24
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v202 int32
	_ = v202
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v321 int32
	_ = v321
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v365 int32
	_ = v365
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v442 int32
	_ = v442
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v560 int32
	_ = v560
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v586 int32
	_ = v586
	var v594 int32
	_ = v594
	var v605 int32
	_ = v605
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v683 int32
	_ = v683
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v709 int32
	_ = v709
	var v720 int32
	_ = v720
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v801 int32
	_ = v801
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v838 int32
	_ = v838
	var v845 int32
	_ = v845
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v869 int32
	_ = v869
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v920 int32
	_ = v920
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v946 int32
	_ = v946
	var v953 int32
	_ = v953
	var v964 int32
	_ = v964
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1041 int32
	_ = v1041
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1067 int32
	_ = v1067
	var v1079 int32
	_ = v1079
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1119 int32
	_ = v1119
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1175 int32
	_ = v1175
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1226 int32
	_ = v1226
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1246 int32
	_ = v1246
	var v1252 int32
	_ = v1252
	var v1259 int32
	_ = v1259
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1297 int32
	_ = v1297
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1335 int32
	_ = v1335
	var v1348 int32
	_ = v1348
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1368 int32
	_ = v1368
	var v1374 int32
	_ = v1374
	var v1382 int32
	_ = v1382
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1421 int32
	_ = v1421
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1443 int32
	_ = v1443
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1472 int32
	_ = v1472
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1492 int32
	_ = v1492
	var v1498 int32
	_ = v1498
	var v1505 int32
	_ = v1505
	var v1516 int32
	_ = v1516
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1543 int32
	_ = v1543
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1556 int32
	_ = v1556
	var v1559 int32
	_ = v1559
	var v1561 int32
	_ = v1561
	var v1565 int32
	_ = v1565
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1581 int32
	_ = v1581
	var v1594 int32
	_ = v1594
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1614 int32
	_ = v1614
	var v1620 int32
	_ = v1620
	var v1628 int32
	_ = v1628
	var v1639 int32
	_ = v1639
	var v1642 int32
	_ = v1642
	var v1647 int32
	_ = v1647
	var v1651 int32
	_ = v1651
	var v1653 int32
	_ = v1653
	var v1655 int32
	_ = v1655
	var v1670 int32
	_ = v1670
	var v1673 int32
	_ = v1673
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1684 int32
	_ = v1684
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1734 int32
	_ = v1734
	var v1737 int32
	_ = v1737
	var v1739 int32
	_ = v1739
	var v1743 int32
	_ = v1743
	var v1749 int32
	_ = v1749
	var v1754 int32
	_ = v1754
	var v1757 int32
	_ = v1757
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1765 int32
	_ = v1765
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1784 int32
	_ = v1784
	var v1788 int32
	_ = v1788
	var v1790 int32
	_ = v1790
	var v1793 int32
	_ = v1793
	var v1795 int32
	_ = v1795
	var v1798 int32
	_ = v1798
	var v1800 int32
	_ = v1800
	var v1802 int32
	_ = v1802
	var v1805 int32
	_ = v1805
	var v1808 int32
	_ = v1808
	var v1811 int32
	_ = v1811
	var v1815 int32
	_ = v1815
	var v1818 int32
	_ = v1818
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1825 int32
	_ = v1825
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1833 int32
	_ = v1833
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1841 int32
	_ = v1841
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1849 int32
	_ = v1849
	var v1851 int32
	_ = v1851
	var v1854 int32
	_ = v1854
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1885 int32
	_ = v1885
	var v1890 int32
	_ = v1890
	var v1892 int32
	_ = v1892
	var v1894 int32
	_ = v1894
	var v1897 int32
	_ = v1897
	var v1900 int32
	_ = v1900
	var v1903 int32
	_ = v1903
	var v1907 int32
	_ = v1907
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1919 int32
	_ = v1919
	var v1922 int32
	_ = v1922
	var v1924 int32
	_ = v1924
	var v1928 int32
	_ = v1928
	var v1932 int32
	_ = v1932
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1946 int32
	_ = v1946
	var v1949 int32
	_ = v1949
	var v1951 int32
	_ = v1951
	var v1954 int32
	_ = v1954
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1960 int32
	_ = v1960
	var v1962 int32
	_ = v1962
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1981 int32
	_ = v1981
	var v1983 int32
	_ = v1983
	var v1985 int32
	_ = v1985
	var v1988 int32
	_ = v1988
	var v1990 int32
	_ = v1990
	var v1993 int32
	_ = v1993
	var v1995 int32
	_ = v1995
	var v1997 int32
	_ = v1997
	var v2000 int32
	_ = v2000
	var v2003 int32
	_ = v2003
	var v2006 int32
	_ = v2006
	var v2010 int32
	_ = v2010
	var v2013 int32
	_ = v2013
	var v2015 int32
	_ = v2015
	var v2017 int32
	_ = v2017
	var v2022 int32
	_ = v2022
	var v2024 int32
	_ = v2024
	var v2027 int32
	_ = v2027
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2038 int32
	_ = v2038
	var v2041 int32
	_ = v2041
	var v2045 int32
	_ = v2045
	var v2048 int32
	_ = v2048
	var v2052 int32
	_ = v2052
	var v2058 int32
	_ = v2058
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2068 int32
	_ = v2068
	var v2072 int32
	_ = v2072
	var v2074 int32
	_ = v2074
	var v2077 int32
	_ = v2077
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2088 int32
	_ = v2088
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2097 int32
	_ = v2097
	var v2101 int32
	_ = v2101
	var v2106 int32
	_ = v2106
	var v2109 int32
	_ = v2109
	var v2112 int32
	_ = v2112
	var v2115 int32
	_ = v2115
	var v2118 int32
	_ = v2118
	var v2125 int32
	_ = v2125
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2135 int32
	_ = v2135
	var v2139 int32
	_ = v2139
	var v2141 int32
	_ = v2141
	var v2144 int32
	_ = v2144
	var v2146 int32
	_ = v2146
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2157 int32
	_ = v2157
	var v2161 int32
	_ = v2161
	var v2167 int32
	_ = v2167
	var v2170 int32
	_ = v2170
	var v2172 int32
	_ = v2172
	var v2179 int32
	_ = v2179
	var v2181 int32
	_ = v2181
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2189 int32
	_ = v2189
	var v2191 int32
	_ = v2191
	var v2193 int32
	_ = v2193
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2247 int32
	_ = v2247
	var v2254 int32
	_ = v2254
	var v2256 int32
	_ = v2256
	var v2261 int32
	_ = v2261
	var v2263 int32
	_ = v2263
	var v2270 int32
	_ = v2270
	var v2273 int32
	_ = v2273
	var v2277 int32
	_ = v2277
	var v2284 int32
	_ = v2284
	var v2285 int32
	_ = v2285
	var v2299 int32
	_ = v2299
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2312 int32
	_ = v2312
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v6
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L7
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1175 = v10
	goto L263
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1149
	goto L1
L3:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1149 = v1146 + v1145
	goto L2
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L136
L5:
	;
	if v128 != 0 {
		goto L4
	} else {
		goto L29
	}
L6:
	;
	v128 = v121
	goto L5
L7:
	;
	if v6 <= v10 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v121 = int32(0)
	goto L6
L9:
	;
	v128 = int32(-1)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v39 = int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v24))))
	if base.Ui32(v41) < base.Ui32(int32(192)) {
		v98 = v41
		v99 = v39
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if int32(252) < v98 {
		v121 = v99
		goto L6
	} else {
		goto L25
	}
L13:
	;
	v45 = v10 + int32(1)
	if v45 == v6 {
		v98 = v41
		v99 = v39
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v24))))
	v50 = v48 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v41) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v24))))
	v66 = v64 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v41) {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v54 = v10 + int32(2)
	if v54 != v6 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v98 = v41<<(uint(int32(6))%32)&int32(1984) | v50
	v99 = int32(2)
	goto L12
L19:
	;
	goto L18
L20:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v70))))
	v98 = v83&int32(63) | (v41<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v50<<(uint(int32(12))%32) | v66<<(uint(int32(6))%32))
	v99 = int32(4)
	goto L12
L21:
	;
	v70 = v10 + int32(3)
	if v70 != v6 {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v98 = v41<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v50<<(uint(int32(6))%32) | v66
	v99 = int32(3)
	goto L12
L24:
	;
	goto L23
L25:
	;
	v103 = v98 - int32(97)
	if v103 < int32(0) {
		v121 = v99
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v103)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v109)>>(uint(v103&int32(7))%32))&int32(1) == int32(0) {
		v121 = v99
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v99 + v10
	goto L28
L28:
	;
	goto L8
L29:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L32
L30:
	;
	if v246 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L31:
	;
	v246 = v239
	goto L30
L32:
	;
	if v142 <= v129 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v239 = int32(0)
	goto L31
L34:
	;
	v246 = int32(-1)
	goto L30
L35:
	;
	goto L36
L36:
	;
	v158 = int32(1)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+v143))))
	if base.Ui32(v160) < base.Ui32(int32(192)) {
		v217 = v160
		v218 = v158
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if int32(252) < v217 {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	v164 = v129 + int32(1)
	if v164 == v142 {
		v217 = v160
		v218 = v158
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+v143))))
	v169 = v167 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v160) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173+v143))))
	v185 = v183 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v160) {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	v173 = v129 + int32(2)
	if v173 != v142 {
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v217 = v160<<(uint(int32(6))%32)&int32(1984) | v169
	v218 = int32(2)
	goto L37
L44:
	;
	goto L43
L45:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143+v189))))
	v217 = v202&int32(63) | (v160<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v169<<(uint(int32(12))%32) | v185<<(uint(int32(6))%32))
	v218 = int32(4)
	goto L37
L46:
	;
	v189 = v129 + int32(3)
	if v189 != v142 {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v217 = v160<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v169<<(uint(int32(6))%32) | v185
	v218 = int32(3)
	goto L37
L49:
	;
	goto L48
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v218 + v129
	goto L54
L51:
	;
	v222 = v217 - int32(97)
	if v222 < int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v222)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v228)>>(uint(v222&int32(7))%32))&int32(1) != 0 {
		v239 = v218
		goto L31
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	goto L33
L55:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v270 = v260
	goto L60
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v129
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L86
L58:
	;
	if int32(0) <= v365 {
		v1145 = v365
		goto L3
	} else {
		goto L83
	}
L59:
	;
	v365 = v337
	goto L58
L60:
	;
	if v261 <= v270 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v365 = int32(-1)
	goto L58
L63:
	;
	goto L64
L64:
	;
	v277 = int32(1)
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270+v262))))
	if base.Ui32(v279) < base.Ui32(int32(192)) {
		v336 = v279
		v337 = v277
		goto L65
	} else {
		goto L66
	}
L65:
	;
	if int32(252) < v336 {
		goto L78
	} else {
		goto L79
	}
L66:
	;
	v283 = v270 + int32(1)
	if v283 == v261 {
		v336 = v279
		v337 = v277
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283+v262))))
	v288 = v286 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v279) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292+v262))))
	v304 = v302 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v279) {
		goto L74
	} else {
		goto L75
	}
L69:
	;
	v292 = v270 + int32(2)
	if v292 != v261 {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v336 = v279<<(uint(int32(6))%32)&int32(1984) | v288
	v337 = int32(2)
	goto L65
L72:
	;
	goto L71
L73:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262+v308))))
	v336 = v321&int32(63) | (v279<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v288<<(uint(int32(12))%32) | v304<<(uint(int32(6))%32))
	v337 = int32(4)
	goto L65
L74:
	;
	v308 = v270 + int32(3)
	if v308 != v261 {
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v336 = v279<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v288<<(uint(int32(6))%32) | v304
	v337 = int32(3)
	goto L65
L77:
	;
	goto L76
L78:
	;
	v354 = v337 + v270
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v354
	v270 = v354
	goto L60
L79:
	;
	v341 = v336 - int32(97)
	if v341 < int32(0) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v341)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v347)>>(uint(v341&int32(7))%32))&int32(1) != 0 {
		goto L59
	} else {
		goto L81
	}
L81:
	;
	goto L78
L83:
	;
	goto L57
L84:
	;
	if v487 != 0 {
		goto L4
	} else {
		goto L108
	}
L85:
	;
	v487 = v480
	goto L84
L86:
	;
	if v382 <= v129 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v480 = int32(0)
	goto L85
L88:
	;
	v487 = int32(-1)
	goto L84
L89:
	;
	goto L90
L90:
	;
	v398 = int32(1)
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+v383))))
	if base.Ui32(v400) < base.Ui32(int32(192)) {
		v457 = v400
		v458 = v398
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if int32(252) < v457 {
		v480 = v458
		goto L85
	} else {
		goto L104
	}
L92:
	;
	v404 = v129 + int32(1)
	if v404 == v382 {
		v457 = v400
		v458 = v398
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v383))))
	v409 = v407 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v400) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413+v383))))
	v425 = v423 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v400) {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	v413 = v129 + int32(2)
	if v413 != v382 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v457 = v400<<(uint(int32(6))%32)&int32(1984) | v409
	v458 = int32(2)
	goto L91
L98:
	;
	goto L97
L99:
	;
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383+v429))))
	v457 = v442&int32(63) | (v400<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v409<<(uint(int32(12))%32) | v425<<(uint(int32(6))%32))
	v458 = int32(4)
	goto L91
L100:
	;
	v429 = v129 + int32(3)
	if v429 != v382 {
		goto L99
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v457 = v400<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v409<<(uint(int32(6))%32) | v425
	v458 = int32(3)
	goto L91
L103:
	;
	goto L102
L104:
	;
	v462 = v457 - int32(97)
	if v462 < int32(0) {
		v480 = v458
		goto L85
	} else {
		goto L105
	}
L105:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v462)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v468)>>(uint(v462&int32(7))%32))&int32(1) == int32(0) {
		v480 = v458
		goto L85
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v458 + v129
	goto L107
L107:
	;
	goto L87
L108:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v509 = v499
	goto L111
L109:
	;
	if int32(0) <= v605 {
		v1145 = v605
		goto L3
	} else {
		goto L133
	}
L110:
	;
	v605 = v576
	goto L109
L111:
	;
	if v500 <= v509 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v605 = int32(-1)
	goto L109
L114:
	;
	goto L115
L115:
	;
	v516 = int32(1)
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509+v501))))
	if base.Ui32(v518) < base.Ui32(int32(192)) {
		v575 = v518
		v576 = v516
		goto L116
	} else {
		goto L117
	}
L116:
	;
	if int32(252) < v575 {
		goto L110
	} else {
		goto L129
	}
L117:
	;
	v522 = v509 + int32(1)
	if v522 == v500 {
		v575 = v518
		v576 = v516
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522+v501))))
	v527 = v525 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v518) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531+v501))))
	v543 = v541 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v518) {
		goto L125
	} else {
		goto L126
	}
L120:
	;
	v531 = v509 + int32(2)
	if v531 != v500 {
		goto L119
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v575 = v518<<(uint(int32(6))%32)&int32(1984) | v527
	v576 = int32(2)
	goto L116
L123:
	;
	goto L122
L124:
	;
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501+v547))))
	v575 = v560&int32(63) | (v518<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v527<<(uint(int32(12))%32) | v543<<(uint(int32(6))%32))
	v576 = int32(4)
	goto L116
L125:
	;
	v547 = v509 + int32(3)
	if v547 != v500 {
		goto L124
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v575 = v518<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v527<<(uint(int32(6))%32) | v543
	v576 = int32(3)
	goto L116
L128:
	;
	goto L127
L129:
	;
	v580 = v575 - int32(97)
	if v580 < int32(0) {
		goto L110
	} else {
		goto L130
	}
L130:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v580)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v586)>>(uint(v580&int32(7))%32))&int32(1) == int32(0) {
		goto L110
	} else {
		goto L131
	}
L131:
	;
	v594 = v576 + v509
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v594
	v509 = v594
	goto L111
L133:
	;
	goto L4
L134:
	;
	if v727 != 0 {
		goto L1
	} else {
		goto L159
	}
L135:
	;
	v727 = v720
	goto L134
L136:
	;
	if v623 <= v10 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v720 = int32(0)
	goto L135
L138:
	;
	v727 = int32(-1)
	goto L134
L139:
	;
	goto L140
L140:
	;
	v639 = int32(1)
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v624))))
	if base.Ui32(v641) < base.Ui32(int32(192)) {
		v698 = v641
		v699 = v639
		goto L141
	} else {
		goto L142
	}
L141:
	;
	if int32(252) < v698 {
		goto L154
	} else {
		goto L155
	}
L142:
	;
	v645 = v10 + int32(1)
	if v645 == v623 {
		v698 = v641
		v699 = v639
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645+v624))))
	v650 = v648 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v641) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654+v624))))
	v666 = v664 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v641) {
		goto L150
	} else {
		goto L151
	}
L145:
	;
	v654 = v10 + int32(2)
	if v654 != v623 {
		goto L144
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v698 = v641<<(uint(int32(6))%32)&int32(1984) | v650
	v699 = int32(2)
	goto L141
L148:
	;
	goto L147
L149:
	;
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624+v670))))
	v698 = v683&int32(63) | (v641<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v650<<(uint(int32(12))%32) | v666<<(uint(int32(6))%32))
	v699 = int32(4)
	goto L141
L150:
	;
	v670 = v10 + int32(3)
	if v670 != v623 {
		goto L149
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v698 = v641<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v650<<(uint(int32(6))%32) | v666
	v699 = int32(3)
	goto L141
L153:
	;
	goto L152
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v699 + v10
	goto L158
L155:
	;
	v703 = v698 - int32(97)
	if v703 < int32(0) {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v703)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v709)>>(uint(v703&int32(7))%32))&int32(1) != 0 {
		v720 = v699
		goto L135
	} else {
		goto L157
	}
L157:
	;
	goto L154
L158:
	;
	goto L137
L159:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L162
L160:
	;
	if v845 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L161:
	;
	v845 = v838
	goto L160
L162:
	;
	if v741 <= v728 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v838 = int32(0)
	goto L161
L164:
	;
	v845 = int32(-1)
	goto L160
L165:
	;
	goto L166
L166:
	;
	v757 = int32(1)
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v728+v742))))
	if base.Ui32(v759) < base.Ui32(int32(192)) {
		v816 = v759
		v817 = v757
		goto L167
	} else {
		goto L168
	}
L167:
	;
	if int32(252) < v816 {
		goto L180
	} else {
		goto L181
	}
L168:
	;
	v763 = v728 + int32(1)
	if v763 == v741 {
		v816 = v759
		v817 = v757
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v763+v742))))
	v768 = v766 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v759) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v772+v742))))
	v784 = v782 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v759) {
		goto L176
	} else {
		goto L177
	}
L171:
	;
	v772 = v728 + int32(2)
	if v772 != v741 {
		goto L170
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v816 = v759<<(uint(int32(6))%32)&int32(1984) | v768
	v817 = int32(2)
	goto L167
L174:
	;
	goto L173
L175:
	;
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v742+v788))))
	v816 = v801&int32(63) | (v759<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v768<<(uint(int32(12))%32) | v784<<(uint(int32(6))%32))
	v817 = int32(4)
	goto L167
L176:
	;
	v788 = v728 + int32(3)
	if v788 != v741 {
		goto L175
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v816 = v759<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v768<<(uint(int32(6))%32) | v784
	v817 = int32(3)
	goto L167
L179:
	;
	goto L178
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v817 + v728
	goto L184
L181:
	;
	v821 = v816 - int32(97)
	if v821 < int32(0) {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v821)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v827)>>(uint(v821&int32(7))%32))&int32(1) != 0 {
		v838 = v817
		goto L161
	} else {
		goto L183
	}
L183:
	;
	goto L180
L184:
	;
	goto L163
L185:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v869 = v859
	goto L190
L186:
	;
	goto L187
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v728
	v981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L216
L188:
	;
	if int32(0) <= v964 {
		v1145 = v964
		goto L3
	} else {
		goto L213
	}
L189:
	;
	v964 = v936
	goto L188
L190:
	;
	if v860 <= v869 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v964 = int32(-1)
	goto L188
L193:
	;
	goto L194
L194:
	;
	v876 = int32(1)
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v869+v861))))
	if base.Ui32(v878) < base.Ui32(int32(192)) {
		v935 = v878
		v936 = v876
		goto L195
	} else {
		goto L196
	}
L195:
	;
	if int32(252) < v935 {
		goto L208
	} else {
		goto L209
	}
L196:
	;
	v882 = v869 + int32(1)
	if v882 == v860 {
		v935 = v878
		v936 = v876
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v882+v861))))
	v887 = v885 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v878) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v891+v861))))
	v903 = v901 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v878) {
		goto L204
	} else {
		goto L205
	}
L199:
	;
	v891 = v869 + int32(2)
	if v891 != v860 {
		goto L198
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v935 = v878<<(uint(int32(6))%32)&int32(1984) | v887
	v936 = int32(2)
	goto L195
L202:
	;
	goto L201
L203:
	;
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861+v907))))
	v935 = v920&int32(63) | (v878<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v887<<(uint(int32(12))%32) | v903<<(uint(int32(6))%32))
	v936 = int32(4)
	goto L195
L204:
	;
	v907 = v869 + int32(3)
	if v907 != v860 {
		goto L203
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	v935 = v878<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v887<<(uint(int32(6))%32) | v903
	v936 = int32(3)
	goto L195
L207:
	;
	goto L206
L208:
	;
	v953 = v936 + v869
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v953
	v869 = v953
	goto L190
L209:
	;
	v940 = v935 - int32(97)
	if v940 < int32(0) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v940)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v946)>>(uint(v940&int32(7))%32))&int32(1) != 0 {
		goto L189
	} else {
		goto L211
	}
L211:
	;
	goto L208
L213:
	;
	goto L187
L214:
	;
	if v1086 != 0 {
		goto L1
	} else {
		goto L238
	}
L215:
	;
	v1086 = v1079
	goto L214
L216:
	;
	if v981 <= v728 {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	v1079 = int32(0)
	goto L215
L218:
	;
	v1086 = int32(-1)
	goto L214
L219:
	;
	goto L220
L220:
	;
	v997 = int32(1)
	v999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v728+v982))))
	if base.Ui32(v999) < base.Ui32(int32(192)) {
		v1056 = v999
		v1057 = v997
		goto L221
	} else {
		goto L222
	}
L221:
	;
	if int32(252) < v1056 {
		v1079 = v1057
		goto L215
	} else {
		goto L234
	}
L222:
	;
	v1003 = v728 + int32(1)
	if v1003 == v981 {
		v1056 = v999
		v1057 = v997
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1003+v982))))
	v1008 = v1006 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v999) {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	v1022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1012+v982))))
	v1024 = v1022 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v999) {
		goto L230
	} else {
		goto L231
	}
L225:
	;
	v1012 = v728 + int32(2)
	if v1012 != v981 {
		goto L224
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v1056 = v999<<(uint(int32(6))%32)&int32(1984) | v1008
	v1057 = int32(2)
	goto L221
L228:
	;
	goto L227
L229:
	;
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v982+v1028))))
	v1056 = v1041&int32(63) | (v999<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v1008<<(uint(int32(12))%32) | v1024<<(uint(int32(6))%32))
	v1057 = int32(4)
	goto L221
L230:
	;
	v1028 = v728 + int32(3)
	if v1028 != v981 {
		goto L229
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v1056 = v999<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v1008<<(uint(int32(6))%32) | v1024
	v1057 = int32(3)
	goto L221
L233:
	;
	goto L232
L234:
	;
	v1061 = v1056 - int32(97)
	if v1061 < int32(0) {
		v1079 = v1057
		goto L215
	} else {
		goto L235
	}
L235:
	;
	v1067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1061)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v1067)>>(uint(v1061&int32(7))%32))&int32(1) == int32(0) {
		v1079 = v1057
		goto L215
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1057 + v728
	goto L237
L237:
	;
	goto L217
L238:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L241
L239:
	;
	if int32(0) <= v1141 {
		v1149 = v1141
		goto L2
	} else {
		goto L259
	}
L241:
	;
	goto L242
L242:
	;
	goto L243
L243:
	;
	v1096 = v1088
	v1098 = int32(1)
	goto L246
L245:
	;
	v1141 = v1126
	goto L239
L246:
	;
	if v1089 <= v1096 {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	goto L245
L248:
	;
	v1141 = int32(-1)
	goto L239
L249:
	;
	goto L250
L250:
	;
	v1103 = v1096 + int32(1)
	v1105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1087+v1096))))
	if base.Ui32(v1105) < base.Ui32(int32(192)) {
		v1126 = v1103
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1127 = int32(1)
	if v1127 < v1098 {
		v1096 = v1126
		v1098 = v1098 - v1127
		goto L246
	} else {
		goto L258
	}
L252:
	;
	if v1089 <= v1103 {
		v1126 = v1103
		goto L251
	} else {
		goto L253
	}
L253:
	;
	v1112 = v1103
	goto L254
L254:
	;
	v1115 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1087+v1112))))
	if int32(-65) < v1115 {
		v1126 = v1112
		goto L251
	} else {
		goto L256
	}
L255:
	;
	v1126 = v1089
	goto L251
L256:
	;
	v1119 = v1112 + int32(1)
	if v1119 != v1089 {
		v1112 = v1119
		goto L254
	} else {
		goto L257
	}
L257:
	;
	goto L255
L258:
	;
	goto L247
L259:
	;
	goto L1
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v10
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1647
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1647
	v1651 = v1647 - int32(1)
	if v1651 <= v10 {
		goto L364
	} else {
		goto L365
	}
L261:
	;
	if v1270 < int32(0) {
		goto L260
	} else {
		goto L286
	}
L262:
	;
	v1270 = v1242
	goto L261
L263:
	;
	if v1166 <= v1175 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1270 = int32(-1)
	goto L261
L266:
	;
	goto L267
L267:
	;
	v1182 = int32(1)
	v1184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175+v1167))))
	if base.Ui32(v1184) < base.Ui32(int32(192)) {
		v1241 = v1184
		v1242 = v1182
		goto L268
	} else {
		goto L269
	}
L268:
	;
	if int32(252) < v1241 {
		goto L281
	} else {
		goto L282
	}
L269:
	;
	v1188 = v1175 + int32(1)
	if v1188 == v1166 {
		v1241 = v1184
		v1242 = v1182
		goto L268
	} else {
		goto L270
	}
L270:
	;
	v1191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1188+v1167))))
	v1193 = v1191 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1184) {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	v1207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1197+v1167))))
	v1209 = v1207 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1184) {
		goto L277
	} else {
		goto L278
	}
L272:
	;
	v1197 = v1175 + int32(2)
	if v1197 != v1166 {
		goto L271
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v1241 = v1184<<(uint(int32(6))%32)&int32(1984) | v1193
	v1242 = int32(2)
	goto L268
L275:
	;
	goto L274
L276:
	;
	v1226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1167+v1213))))
	v1241 = v1226&int32(63) | (v1184<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v1193<<(uint(int32(12))%32) | v1209<<(uint(int32(6))%32))
	v1242 = int32(4)
	goto L268
L277:
	;
	v1213 = v1175 + int32(3)
	if v1213 != v1166 {
		goto L276
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v1241 = v1184<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v1193<<(uint(int32(6))%32) | v1209
	v1242 = int32(3)
	goto L268
L280:
	;
	goto L279
L281:
	;
	v1259 = v1242 + v1175
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1259
	v1175 = v1259
	goto L263
L282:
	;
	v1246 = v1241 - int32(97)
	if v1246 < int32(0) {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1246)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v1252)>>(uint(v1246&int32(7))%32))&int32(1) != 0 {
		goto L262
	} else {
		goto L284
	}
L284:
	;
	goto L281
L286:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1274 = v1273 + v1270
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1274
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1297 = v1274
	goto L289
L287:
	;
	if v1393 < int32(0) {
		goto L260
	} else {
		goto L311
	}
L288:
	;
	v1393 = v1364
	goto L287
L289:
	;
	if v1288 <= v1297 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1393 = int32(-1)
	goto L287
L292:
	;
	goto L293
L293:
	;
	v1304 = int32(1)
	v1306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1297+v1289))))
	if base.Ui32(v1306) < base.Ui32(int32(192)) {
		v1363 = v1306
		v1364 = v1304
		goto L294
	} else {
		goto L295
	}
L294:
	;
	if int32(252) < v1363 {
		goto L288
	} else {
		goto L307
	}
L295:
	;
	v1310 = v1297 + int32(1)
	if v1310 == v1288 {
		v1363 = v1306
		v1364 = v1304
		goto L294
	} else {
		goto L296
	}
L296:
	;
	v1313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1310+v1289))))
	v1315 = v1313 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1306) {
		goto L298
	} else {
		goto L299
	}
L297:
	;
	v1329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1319+v1289))))
	v1331 = v1329 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1306) {
		goto L303
	} else {
		goto L304
	}
L298:
	;
	v1319 = v1297 + int32(2)
	if v1319 != v1288 {
		goto L297
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	v1363 = v1306<<(uint(int32(6))%32)&int32(1984) | v1315
	v1364 = int32(2)
	goto L294
L301:
	;
	goto L300
L302:
	;
	v1348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1289+v1335))))
	v1363 = v1348&int32(63) | (v1306<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v1315<<(uint(int32(12))%32) | v1331<<(uint(int32(6))%32))
	v1364 = int32(4)
	goto L294
L303:
	;
	v1335 = v1297 + int32(3)
	if v1335 != v1288 {
		goto L302
	} else {
		goto L306
	}
L304:
	;
	goto L305
L305:
	;
	v1363 = v1306<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v1315<<(uint(int32(6))%32) | v1331
	v1364 = int32(3)
	goto L294
L306:
	;
	goto L305
L307:
	;
	v1368 = v1363 - int32(97)
	if v1368 < int32(0) {
		goto L288
	} else {
		goto L308
	}
L308:
	;
	v1374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1368)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v1374)>>(uint(v1368&int32(7))%32))&int32(1) == int32(0) {
		goto L288
	} else {
		goto L309
	}
L309:
	;
	v1382 = v1364 + v1297
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1382
	v1297 = v1382
	goto L289
L311:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1397 = v1396 + v1393
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1397
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1421 = v1397
	goto L314
L312:
	;
	if v1516 < int32(0) {
		goto L260
	} else {
		goto L337
	}
L313:
	;
	v1516 = v1488
	goto L312
L314:
	;
	if v1412 <= v1421 {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v1516 = int32(-1)
	goto L312
L317:
	;
	goto L318
L318:
	;
	v1428 = int32(1)
	v1430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1421+v1413))))
	if base.Ui32(v1430) < base.Ui32(int32(192)) {
		v1487 = v1430
		v1488 = v1428
		goto L319
	} else {
		goto L320
	}
L319:
	;
	if int32(252) < v1487 {
		goto L332
	} else {
		goto L333
	}
L320:
	;
	v1434 = v1421 + int32(1)
	if v1434 == v1412 {
		v1487 = v1430
		v1488 = v1428
		goto L319
	} else {
		goto L321
	}
L321:
	;
	v1437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1434+v1413))))
	v1439 = v1437 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1430) {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	v1453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1443+v1413))))
	v1455 = v1453 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1430) {
		goto L328
	} else {
		goto L329
	}
L323:
	;
	v1443 = v1421 + int32(2)
	if v1443 != v1412 {
		goto L322
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	v1487 = v1430<<(uint(int32(6))%32)&int32(1984) | v1439
	v1488 = int32(2)
	goto L319
L326:
	;
	goto L325
L327:
	;
	v1472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1413+v1459))))
	v1487 = v1472&int32(63) | (v1430<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v1439<<(uint(int32(12))%32) | v1455<<(uint(int32(6))%32))
	v1488 = int32(4)
	goto L319
L328:
	;
	v1459 = v1421 + int32(3)
	if v1459 != v1412 {
		goto L327
	} else {
		goto L331
	}
L329:
	;
	goto L330
L330:
	;
	v1487 = v1430<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v1439<<(uint(int32(6))%32) | v1455
	v1488 = int32(3)
	goto L319
L331:
	;
	goto L330
L332:
	;
	v1505 = v1488 + v1421
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1505
	v1421 = v1505
	goto L314
L333:
	;
	v1492 = v1487 - int32(97)
	if v1492 < int32(0) {
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v1498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1492)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v1498)>>(uint(v1492&int32(7))%32))&int32(1) != 0 {
		goto L313
	} else {
		goto L335
	}
L335:
	;
	goto L332
L337:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1520 = v1519 + v1516
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1520
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1543 = v1520
	goto L340
L338:
	;
	if v1639 < int32(0) {
		goto L260
	} else {
		goto L362
	}
L339:
	;
	v1639 = v1610
	goto L338
L340:
	;
	if v1534 <= v1543 {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1639 = int32(-1)
	goto L338
L343:
	;
	goto L344
L344:
	;
	v1550 = int32(1)
	v1552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1543+v1535))))
	if base.Ui32(v1552) < base.Ui32(int32(192)) {
		v1609 = v1552
		v1610 = v1550
		goto L345
	} else {
		goto L346
	}
L345:
	;
	if int32(252) < v1609 {
		goto L339
	} else {
		goto L358
	}
L346:
	;
	v1556 = v1543 + int32(1)
	if v1556 == v1534 {
		v1609 = v1552
		v1610 = v1550
		goto L345
	} else {
		goto L347
	}
L347:
	;
	v1559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1556+v1535))))
	v1561 = v1559 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1552) {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	v1575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1565+v1535))))
	v1577 = v1575 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1552) {
		goto L354
	} else {
		goto L355
	}
L349:
	;
	v1565 = v1543 + int32(2)
	if v1565 != v1534 {
		goto L348
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1609 = v1552<<(uint(int32(6))%32)&int32(1984) | v1561
	v1610 = int32(2)
	goto L345
L352:
	;
	goto L351
L353:
	;
	v1594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1535+v1581))))
	v1609 = v1594&int32(63) | (v1552<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v1561<<(uint(int32(12))%32) | v1577<<(uint(int32(6))%32))
	v1610 = int32(4)
	goto L345
L354:
	;
	v1581 = v1543 + int32(3)
	if v1581 != v1534 {
		goto L353
	} else {
		goto L357
	}
L355:
	;
	goto L356
L356:
	;
	v1609 = v1552<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v1561<<(uint(int32(6))%32) | v1577
	v1610 = int32(3)
	goto L345
L357:
	;
	goto L356
L358:
	;
	v1614 = v1609 - int32(97)
	if v1614 < int32(0) {
		goto L339
	} else {
		goto L359
	}
L359:
	;
	v1620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1614)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v1620)>>(uint(v1614&int32(7))%32))&int32(1) == int32(0) {
		goto L339
	} else {
		goto L360
	}
L360:
	;
	v1628 = v1610 + v1543
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1628
	v1543 = v1628
	goto L340
L362:
	;
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v1642 + v1639
	goto L260
L363:
	;
	return v2312
L364:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1754
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1754
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1754-int32(2) <= v1757 {
		goto L397
	} else {
		goto L398
	}
L365:
	;
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1653+v1651))))
	if base.B2i32(v1655&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1655)%32)&int32(_a_F_spanish_UTF_8_stem_2) == int32(0)) != 0 {
		goto L364
	} else {
		goto L366
	}
L366:
	;
	v1670 = F_find_among_b(m, l0, int32(_a_F_spanish_UTF_8_stem_3), int32(13), int32(0))
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	return int32(0)
L368:
	;
	if v1670 == int32(0) {
		goto L364
	} else {
		goto L369
	}
L369:
	;
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1676
	v1679 = v1676 - int32(1)
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1679 <= v1680 {
		goto L364
	} else {
		goto L370
	}
L370:
	;
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1682+v1679))))
	switch v1684 - int32(111) {
	case 0, 3:
		goto L371
	default:
		goto L364
	}
L371:
	;
	v1690 = F_find_among_b(m, l0, int32(_a_F_spanish_UTF_8_stem_4), int32(11), int32(0))
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L367
	} else {
		goto L372
	}
L372:
	;
	if v1690 == int32(0) {
		goto L364
	} else {
		goto L373
	}
L373:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1694 < v1695 {
		goto L364
	} else {
		goto L374
	}
L374:
	;
	switch v1690 - int32(1) {
	case 0:
		goto L381
	case 1:
		goto L380
	case 2:
		goto L379
	case 3:
		goto L378
	case 4:
		goto L377
	case 5:
		goto L376
	case 6:
		goto L375
	default:
		goto L364
	}
L375:
	;
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1694 <= v1737 {
		goto L364
	} else {
		goto L393
	}
L376:
	;
	v1734 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1734 {
		goto L364
	} else {
		goto L392
	}
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1694
	v1730 = F_slice_from_s(m, l0, int32(2), int32(_a_F_spanish_UTF_8_stem_5))
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L367
	} else {
		goto L390
	}
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1694
	v1723 = F_slice_from_s(m, l0, int32(2), int32(_a_F_spanish_UTF_8_stem_6))
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L367
	} else {
		goto L388
	}
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1694
	v1716 = F_slice_from_s(m, l0, int32(2), int32(_a_F_spanish_UTF_8_stem_7))
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L367
	} else {
		goto L386
	}
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1694
	v1709 = F_slice_from_s(m, l0, int32(4), int32(_a_F_spanish_UTF_8_stem_8))
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L367
	} else {
		goto L384
	}
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1694
	v1702 = F_slice_from_s(m, l0, int32(5), int32(_a_F_spanish_UTF_8_stem_9))
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L367
	} else {
		goto L382
	}
L382:
	;
	if int32(0) <= v1702 {
		goto L364
	} else {
		goto L383
	}
L383:
	;
	v2312 = v1702
	goto L363
L384:
	;
	if int32(0) <= v1709 {
		goto L364
	} else {
		goto L385
	}
L385:
	;
	v2312 = v1709
	goto L363
L386:
	;
	if int32(0) <= v1716 {
		goto L364
	} else {
		goto L387
	}
L387:
	;
	v2312 = v1716
	goto L363
L388:
	;
	if int32(0) <= v1723 {
		goto L364
	} else {
		goto L389
	}
L389:
	;
	v2312 = v1723
	goto L363
L390:
	;
	if int32(0) <= v1730 {
		goto L364
	} else {
		goto L391
	}
L391:
	;
	v2312 = v1730
	goto L363
L392:
	;
	v2312 = v1734
	goto L363
L393:
	;
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1739+v1694-int32(1)))))
	if v1743 != int32(117) {
		goto L364
	} else {
		goto L394
	}
L394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1694 - int32(1)
	v1749 = F_slice_del(m, l0)
	mBase = m.M
	if v1749 < int32(0) {
		v2312 = v1749
		goto L363
	} else {
		goto L395
	}
L395:
	;
	goto L364
L396:
	;
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2125
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2125
	v2131 = F_find_among_b(m, l0, int32(_a_F_spanish_UTF_8_stem_10), int32(8), int32(0))
	mBase = m.M
	v2132 = m.ExcPending
	if v2132 != 0 {
		goto L367
	} else {
		goto L508
	}
L397:
	;
	v2022 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2022
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2022 < v2024 {
		goto L472
	} else {
		goto L473
	}
L398:
	;
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1763 = int32(1)
	v1765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1761+v1754-v1763))))
	if base.B2i32(v1765&int32(224) != int32(96))|base.B2i32(v1763<<(uint(v1765)%32)&int32(_a_F_spanish_UTF_8_stem_11) == int32(0)) != 0 {
		goto L397
	} else {
		goto L399
	}
L399:
	;
	v1780 = F_find_among_b(m, l0, int32(_a_F_spanish_UTF_8_stem_12), int32(48), int32(0))
	mBase = m.M
	v1781 = m.ExcPending
	if v1781 != 0 {
		goto L367
	} else {
		goto L400
	}
L400:
	;
	if v1780 == int32(0) {
		goto L397
	} else {
		goto L401
	}
L401:
	;
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1784
	switch v1780 - int32(1) {
	case 0:
		goto L410
	case 1:
		goto L409
	case 2:
		goto L408
	case 3:
		goto L407
	case 4:
		goto L406
	case 5:
		goto L405
	case 6:
		goto L404
	case 7:
		goto L403
	case 8:
		goto L402
	default:
		goto L396
	}
L402:
	;
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1784 < v1988 {
		goto L397
	} else {
		goto L463
	}
L403:
	;
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1784 < v1949 {
		goto L397
	} else {
		goto L455
	}
L404:
	;
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1784 < v1917 {
		goto L397
	} else {
		goto L447
	}
L405:
	;
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1784 < v1849 {
		goto L397
	} else {
		goto L431
	}
L406:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1784 < v1841 {
		goto L397
	} else {
		goto L428
	}
L407:
	;
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1784 < v1833 {
		goto L397
	} else {
		goto L425
	}
L408:
	;
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1784 < v1825 {
		goto L397
	} else {
		goto L422
	}
L409:
	;
	v1793 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1784 < v1793 {
		goto L397
	} else {
		goto L413
	}
L410:
	;
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1784 < v1788 {
		goto L397
	} else {
		goto L411
	}
L411:
	;
	v1790 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1790 {
		goto L396
	} else {
		goto L412
	}
L412:
	;
	v2312 = v1790
	goto L363
L413:
	;
	v1795 = F_slice_del(m, l0)
	mBase = m.M
	if v1795 < int32(0) {
		v2312 = v1795
		goto L363
	} else {
		goto L414
	}
L414:
	;
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1798
	v1800 = int32(2)
	v1802 = int32(0)
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1798-v1805 < v1800 {
		v1815 = v1802
		goto L416
	} else {
		goto L417
	}
L415:
	;
	if v1815 == int32(0) {
		goto L396
	} else {
		goto L419
	}
L416:
	;
	goto L415
L417:
	;
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1811 = F_memcmp(m, v1808+v1798-v1800, int32(_a_F_spanish_UTF_8_stem_13), v1800)
	mBase = m.M
	if v1811 != 0 {
		v1815 = v1802
		goto L416
	} else {
		goto L418
	}
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1798 - v1800
	v1815 = int32(1)
	goto L416
L419:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1818
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1818 < v1820 {
		goto L396
	} else {
		goto L420
	}
L420:
	;
	v1822 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1822 {
		goto L396
	} else {
		goto L421
	}
L421:
	;
	v2312 = v1822
	goto L363
L422:
	;
	v1829 = F_slice_from_s(m, l0, int32(3), int32(_a_F_spanish_UTF_8_stem_14))
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L367
	} else {
		goto L423
	}
L423:
	;
	if int32(0) <= v1829 {
		goto L396
	} else {
		goto L424
	}
L424:
	;
	v2312 = v1829
	goto L363
L425:
	;
	v1837 = F_slice_from_s(m, l0, int32(1), int32(_a_F_spanish_UTF_8_stem_15))
	mBase = m.M
	v1838 = m.ExcPending
	if v1838 != 0 {
		goto L367
	} else {
		goto L426
	}
L426:
	;
	if int32(0) <= v1837 {
		goto L396
	} else {
		goto L427
	}
L427:
	;
	v2312 = v1837
	goto L363
L428:
	;
	v1845 = F_slice_from_s(m, l0, int32(4), int32(_a_F_spanish_UTF_8_stem_16))
	mBase = m.M
	v1846 = m.ExcPending
	if v1846 != 0 {
		goto L367
	} else {
		goto L429
	}
L429:
	;
	if int32(0) <= v1845 {
		goto L396
	} else {
		goto L430
	}
L430:
	;
	v2312 = v1845
	goto L363
L431:
	;
	v1851 = F_slice_del(m, l0)
	mBase = m.M
	if v1851 < int32(0) {
		v2312 = v1851
		goto L363
	} else {
		goto L432
	}
L432:
	;
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1854
	v1857 = v1854 - int32(1)
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1857 <= v1858 {
		goto L396
	} else {
		goto L433
	}
L433:
	;
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1860+v1857))))
	if base.B2i32(v1862&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1862)%32)&int32(_a_F_spanish_UTF_8_stem_17) == int32(0)) != 0 {
		goto L396
	} else {
		goto L434
	}
L434:
	;
	v1877 = F_find_among_b(m, l0, int32(_a_F_spanish_UTF_8_stem_18), int32(4), int32(0))
	mBase = m.M
	v1878 = m.ExcPending
	if v1878 != 0 {
		goto L367
	} else {
		goto L435
	}
L435:
	;
	if v1877 == int32(0) {
		goto L396
	} else {
		goto L436
	}
L436:
	;
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1881
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1881 < v1883 {
		goto L396
	} else {
		goto L437
	}
L437:
	;
	v1885 = F_slice_del(m, l0)
	mBase = m.M
	if v1885 < int32(0) {
		v2312 = v1885
		goto L363
	} else {
		goto L438
	}
L438:
	;
	if v1877 != int32(1) {
		goto L396
	} else {
		goto L439
	}
L439:
	;
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1890
	v1892 = int32(2)
	v1894 = int32(0)
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1890-v1897 < v1892 {
		v1907 = v1894
		goto L441
	} else {
		goto L442
	}
L440:
	;
	if v1907 == int32(0) {
		goto L396
	} else {
		goto L444
	}
L441:
	;
	goto L440
L442:
	;
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1903 = F_memcmp(m, v1900+v1890-v1892, int32(_a_F_spanish_UTF_8_stem_19), v1892)
	mBase = m.M
	if v1903 != 0 {
		v1907 = v1894
		goto L441
	} else {
		goto L443
	}
L443:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1890 - v1892
	v1907 = int32(1)
	goto L441
L444:
	;
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1910
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1910 < v1912 {
		goto L396
	} else {
		goto L445
	}
L445:
	;
	v1914 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1914 {
		goto L396
	} else {
		goto L446
	}
L446:
	;
	v2312 = v1914
	goto L363
L447:
	;
	v1919 = F_slice_del(m, l0)
	mBase = m.M
	if v1919 < int32(0) {
		v2312 = v1919
		goto L363
	} else {
		goto L448
	}
L448:
	;
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1922
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1922-int32(3) <= v1924 {
		goto L396
	} else {
		goto L449
	}
L449:
	;
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1928+v1922-int32(1)))))
	if v1932 != int32(101) {
		goto L396
	} else {
		goto L450
	}
L450:
	;
	v1938 = F_find_among_b(m, l0, int32(_a_F_spanish_UTF_8_stem_20), int32(3), int32(0))
	mBase = m.M
	v1939 = m.ExcPending
	if v1939 != 0 {
		goto L367
	} else {
		goto L451
	}
L451:
	;
	if v1938 == int32(0) {
		goto L396
	} else {
		goto L452
	}
L452:
	;
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1942
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1942 < v1944 {
		goto L396
	} else {
		goto L453
	}
L453:
	;
	v1946 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1946 {
		goto L396
	} else {
		goto L454
	}
L454:
	;
	v2312 = v1946
	goto L363
L455:
	;
	v1951 = F_slice_del(m, l0)
	mBase = m.M
	if v1951 < int32(0) {
		v2312 = v1951
		goto L363
	} else {
		goto L456
	}
L456:
	;
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1954
	v1957 = v1954 - int32(1)
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1957 <= v1958 {
		goto L396
	} else {
		goto L457
	}
L457:
	;
	v1960 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1960+v1957))))
	if base.B2i32(v1962&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1962)%32)&int32(_a_F_spanish_UTF_8_stem_21) == int32(0)) != 0 {
		goto L396
	} else {
		goto L458
	}
L458:
	;
	v1977 = F_find_among_b(m, l0, int32(_a_F_spanish_UTF_8_stem_22), int32(3), int32(0))
	mBase = m.M
	v1978 = m.ExcPending
	if v1978 != 0 {
		goto L367
	} else {
		goto L459
	}
L459:
	;
	if v1977 == int32(0) {
		goto L396
	} else {
		goto L460
	}
L460:
	;
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1981
	v1983 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1981 < v1983 {
		goto L396
	} else {
		goto L461
	}
L461:
	;
	v1985 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1985 {
		goto L396
	} else {
		goto L462
	}
L462:
	;
	v2312 = v1985
	goto L363
L463:
	;
	v1990 = F_slice_del(m, l0)
	mBase = m.M
	if v1990 < int32(0) {
		v2312 = v1990
		goto L363
	} else {
		goto L464
	}
L464:
	;
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1993
	v1995 = int32(2)
	v1997 = int32(0)
	v2000 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1993-v2000 < v1995 {
		v2010 = v1997
		goto L466
	} else {
		goto L467
	}
L465:
	;
	if v2010 == int32(0) {
		goto L396
	} else {
		goto L469
	}
L466:
	;
	goto L465
L467:
	;
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2006 = F_memcmp(m, v2003+v1993-v1995, int32(_a_F_spanish_UTF_8_stem_23), v1995)
	mBase = m.M
	if v2006 != 0 {
		v2010 = v1997
		goto L466
	} else {
		goto L468
	}
L468:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1993 - v1995
	v2010 = int32(1)
	goto L466
L469:
	;
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2013
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2013 < v2015 {
		goto L396
	} else {
		goto L470
	}
L470:
	;
	v2017 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2017 {
		goto L396
	} else {
		goto L471
	}
L471:
	;
	v2312 = v2017
	goto L363
L472:
	;
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2072
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2072 < v2074 {
		goto L396
	} else {
		goto L491
	}
L473:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2022
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2024
	v2032 = F_find_among_b(m, l0, int32(_a_F_spanish_UTF_8_stem_24), int32(12), int32(0))
	mBase = m.M
	v2033 = m.ExcPending
	if v2033 != 0 {
		goto L367
	} else {
		goto L474
	}
L474:
	;
	if v2032 == int32(0) {
		goto L475
	} else {
		goto L476
	}
L475:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2027
	goto L472
L476:
	;
	goto L477
L477:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2027
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2038
	if v2038 <= v2027 {
		goto L472
	} else {
		goto L478
	}
L478:
	;
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2041+v2038-int32(1)))))
	if v2045 != int32(117) {
		goto L472
	} else {
		goto L479
	}
L479:
	;
	v2048 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2038 - v2048
	v2052 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2052 {
		goto L481
	} else {
		goto L482
	}
L480:
	;
	v2063 = int32(0)
	v2064 = base.B2i32(v2058 < v2063)
	if v2064 == v2063 {
		goto L396
	} else {
		goto L487
	}
L481:
	;
	v2058 = v2048
	goto L483
L482:
	;
	v2058 = v2052 >> (uint(int32(31)) % 32) & v2052
	goto L483
L483:
	;
	if v2058 != 0 {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v2062 = int32(base.Ui32(v2058) >> (uint(int32(31)) % 32))
	goto L486
L485:
	;
	v2062 = int32(5)
	goto L486
L486:
	;
	switch v2062 {
	case 0:
		goto L396
	default:
		goto L480
	case 5:
		goto L472
	}
L487:
	;
	if v2058 < v2063 {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	v2068 = v2058
	goto L490
L489:
	;
	v2068 = int32(1)
	goto L490
L490:
	;
	return v2068
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2072
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2074
	v2082 = F_find_among_b(m, l0, int32(_a_F_spanish_UTF_8_stem_25), int32(96), int32(0))
	mBase = m.M
	v2083 = m.ExcPending
	if v2083 != 0 {
		goto L367
	} else {
		goto L492
	}
L492:
	;
	if v2082 == int32(0) {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2077
	goto L396
L494:
	;
	goto L495
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2077
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2088
	switch v2082 - int32(1) {
	case 0:
		goto L497
	case 1:
		goto L496
	default:
		goto L396
	}
L496:
	;
	v2118 = F_slice_del(m, l0)
	mBase = m.M
	if v2118 < int32(0) {
		v2312 = v2118
		goto L363
	} else {
		goto L506
	}
L497:
	;
	if v2088 <= v2077 {
		v2112 = v2088
		goto L498
	} else {
		goto L499
	}
L498:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2112
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2112
	v2115 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2115 {
		goto L396
	} else {
		goto L505
	}
L499:
	;
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2094 = v2093 + v2088
	v2097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2094-int32(1)))))
	if v2097 != int32(117) {
		v2112 = v2088
		goto L498
	} else {
		goto L500
	}
L500:
	;
	v2101 = v2088 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2101
	if v2101 <= v2077 {
		v2112 = v2088
		goto L498
	} else {
		goto L501
	}
L501:
	;
	v2106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2094-int32(2)))))
	if v2106 == int32(103) {
		goto L502
	} else {
		goto L503
	}
L502:
	;
	v2109 = v2101
	goto L504
L503:
	;
	v2109 = v2088
	goto L504
L504:
	;
	v2112 = v2109
	goto L498
L505:
	;
	v2312 = v2115
	goto L363
L506:
	;
	goto L396
L507:
	;
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2179
	v2181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2184 = v2179
	v2185 = v2181
	goto L522
L508:
	;
	if v2131 == int32(0) {
		goto L507
	} else {
		goto L509
	}
L509:
	;
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2135
	switch v2131 - int32(1) {
	case 0:
		goto L511
	case 1:
		goto L510
	default:
		goto L507
	}
L510:
	;
	v2144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2135 < v2144 {
		goto L507
	} else {
		goto L514
	}
L511:
	;
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2135 < v2139 {
		goto L507
	} else {
		goto L512
	}
L512:
	;
	v2141 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2141 {
		goto L507
	} else {
		goto L513
	}
L513:
	;
	v2312 = v2141
	goto L363
L514:
	;
	v2146 = F_slice_del(m, l0)
	mBase = m.M
	if v2146 < int32(0) {
		v2312 = v2146
		goto L363
	} else {
		goto L515
	}
L515:
	;
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2149
	v2151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2149 <= v2151 {
		goto L507
	} else {
		goto L516
	}
L516:
	;
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2154 = v2153 + v2149
	v2157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2154-int32(1)))))
	if v2157 != int32(117) {
		goto L507
	} else {
		goto L517
	}
L517:
	;
	v2161 = v2149 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2161
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2161
	if v2161 <= v2151 {
		goto L507
	} else {
		goto L518
	}
L518:
	;
	v2167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2154-int32(2)))))
	if v2167 != int32(103) {
		goto L507
	} else {
		goto L519
	}
L519:
	;
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2149 <= v2170 {
		goto L507
	} else {
		goto L520
	}
L520:
	;
	v2172 = F_slice_del(m, l0)
	mBase = m.M
	if v2172 < int32(0) {
		v2312 = v2172
		goto L363
	} else {
		goto L521
	}
L521:
	;
	goto L507
L522:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2184
	v2189 = v2184 + int32(1)
	if v2185 <= v2189 {
		goto L528
	} else {
		goto L529
	}
L523:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2179
	v2312 = int32(1)
	goto L363
L524:
	;
	goto L523
L525:
	;
	v2306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2184 = v2307
	v2185 = v2306
	goto L522
L526:
	;
	v2247 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L551
L527:
	;
	v2207 = F_find_among(m, l0, int32(_a_F_spanish_UTF_8_stem_26), int32(6), int32(0))
	mBase = m.M
	v2208 = m.ExcPending
	if v2208 != 0 {
		goto L367
	} else {
		goto L532
	}
L528:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2184
	v2244 = v2184
	v2245 = v2185
	goto L526
L529:
	;
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2191+v2189))))
	if v2193&int32(224) != int32(160) {
		goto L528
	} else {
		goto L530
	}
L530:
	;
	if int32(1)<<(uint(v2193)%32)&int32(67641858) != 0 {
		goto L527
	} else {
		goto L531
	}
L531:
	;
	goto L528
L532:
	;
	v2209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2209
	switch v2207 - int32(1) {
	case 0:
		goto L538
	case 1:
		goto L537
	case 2:
		goto L536
	case 3:
		goto L535
	case 4:
		goto L534
	case 5:
		goto L533
	default:
		goto L525
	}
L533:
	;
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2244 = v2209
	v2245 = v2243
	goto L526
L534:
	;
	v2239 = F_slice_from_s(m, l0, int32(1), int32(_a_F_spanish_UTF_8_stem_27))
	mBase = m.M
	v2240 = m.ExcPending
	if v2240 != 0 {
		goto L367
	} else {
		goto L547
	}
L535:
	;
	v2233 = F_slice_from_s(m, l0, int32(1), int32(_a_F_spanish_UTF_8_stem_28))
	mBase = m.M
	v2234 = m.ExcPending
	if v2234 != 0 {
		goto L367
	} else {
		goto L545
	}
L536:
	;
	v2227 = F_slice_from_s(m, l0, int32(1), int32(_a_F_spanish_UTF_8_stem_29))
	mBase = m.M
	v2228 = m.ExcPending
	if v2228 != 0 {
		goto L367
	} else {
		goto L543
	}
L537:
	;
	v2221 = F_slice_from_s(m, l0, int32(1), int32(_a_F_spanish_UTF_8_stem_30))
	mBase = m.M
	v2222 = m.ExcPending
	if v2222 != 0 {
		goto L367
	} else {
		goto L541
	}
L538:
	;
	v2215 = F_slice_from_s(m, l0, int32(1), int32(_a_F_spanish_UTF_8_stem_31))
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L367
	} else {
		goto L539
	}
L539:
	;
	if int32(0) <= v2215 {
		goto L525
	} else {
		goto L540
	}
L540:
	;
	v2312 = v2215
	goto L363
L541:
	;
	if int32(0) <= v2221 {
		goto L525
	} else {
		goto L542
	}
L542:
	;
	v2312 = v2221
	goto L363
L543:
	;
	if int32(0) <= v2227 {
		goto L525
	} else {
		goto L544
	}
L544:
	;
	v2312 = v2227
	goto L363
L545:
	;
	if int32(0) <= v2233 {
		goto L525
	} else {
		goto L546
	}
L546:
	;
	v2312 = v2233
	goto L363
L547:
	;
	if int32(0) <= v2239 {
		goto L525
	} else {
		goto L548
	}
L548:
	;
	v2312 = v2239
	goto L363
L549:
	;
	if v2299 < int32(0) {
		goto L524
	} else {
		goto L569
	}
L551:
	;
	goto L552
L552:
	;
	goto L553
L553:
	;
	v2254 = v2244
	v2256 = int32(1)
	goto L556
L555:
	;
	v2299 = v2284
	goto L549
L556:
	;
	if v2245 <= v2254 {
		goto L558
	} else {
		goto L559
	}
L557:
	;
	goto L555
L558:
	;
	v2299 = int32(-1)
	goto L549
L559:
	;
	goto L560
L560:
	;
	v2261 = v2254 + int32(1)
	v2263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2247+v2254))))
	if base.Ui32(v2263) < base.Ui32(int32(192)) {
		v2284 = v2261
		goto L561
	} else {
		goto L562
	}
L561:
	;
	v2285 = int32(1)
	if v2285 < v2256 {
		v2254 = v2284
		v2256 = v2256 - v2285
		goto L556
	} else {
		goto L568
	}
L562:
	;
	if v2245 <= v2261 {
		v2284 = v2261
		goto L561
	} else {
		goto L563
	}
L563:
	;
	v2270 = v2261
	goto L564
L564:
	;
	v2273 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2247+v2270))))
	if int32(-65) < v2273 {
		v2284 = v2270
		goto L561
	} else {
		goto L566
	}
L565:
	;
	v2284 = v2245
	goto L561
L566:
	;
	v2277 = v2270 + int32(1)
	if v2277 != v2245 {
		v2270 = v2277
		goto L564
	} else {
		goto L567
	}
L567:
	;
	goto L565
L568:
	;
	goto L557
L569:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2299
	goto L525
}
func F_spcachekey_hash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v13 int64
	_ = v13
	var v17 int64
	_ = v17
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v34 int64
	_ = v34
	var v41 int64
	_ = v41
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v52 int64
	_ = v52
	var v57 int64
	_ = v57
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v104 int64
	_ = v104
	var v110 int64
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v127 int64
	_ = v127
	v8 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v13 = (int64(base.Ui64(v8)>>(uint(int64(23))%64)) ^ v8) * int64(2388976653695081527)
	v17 = int64(-8645972361240307355)
	v20 = (v13 ^ int64(base.Ui64(v13)>>(uint(int64(47))%64)) ^ v17) * v17
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v22 == int32(0) {
		v113 = v21
		v115 = v20
	} else {
		v25 = v21
		v27 = v20
		v30 = v22
		for {
			v34 = int64(*(*int8)(unsafe.Add(mBase, uint32(v25)+1)))
			if v34 == int64(0) {
				v92 = int32(1)
				v93 = int64(0)
				v98 = v92
				v99 = base.I64_extend8_s(base.I64_extend_i32_u(v30)) | v93
			} else {
				v41 = int64(*(*int8)(unsafe.Add(mBase, uint32(v25)+2)))
				if v41 == int64(0) {
					v88 = int32(2)
					v89 = int64(0)
					v92 = v88
					v93 = v34<<(uint(int64(8))%64) | v89
					v98 = v92
					v99 = base.I64_extend8_s(base.I64_extend_i32_u(v30)) | v93
				} else {
					v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+3)))
					if v46 != 0 {
						v47 = int64(*(*int8)(unsafe.Add(mBase, uint32(v25)+4)))
						if v47 == int64(0) {
							v81 = int32(4)
							v82 = int64(0)
							v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v25))))
							v98 = v81
							v99 = v82 | v83
						} else {
							v52 = int64(*(*int8)(unsafe.Add(mBase, uint32(v25)+5)))
							if v52 == int64(0) {
								v74 = int32(5)
								v75 = int64(0)
								v81 = v74
								v82 = v75 | v47<<(uint(int64(32))%64)
								v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v25))))
								v98 = v81
								v99 = v82 | v83
							} else {
								v57 = int64(*(*int8)(unsafe.Add(mBase, uint32(v25)+6)))
								if v57 == int64(0) {
									v68 = int64(0)
									v69 = int32(6)
									v74 = v69
									v75 = v68 | v52<<(uint(int64(40))%64)
									v81 = v74
									v82 = v75 | v47<<(uint(int64(32))%64)
									v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v25))))
									v98 = v81
									v99 = v82 | v83
								} else {
									v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+7)))
									if v62 != 0 {
										v64 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
										v98 = int32(8)
										v99 = v64
									} else {
										v68 = v57 << (uint(int64(48)) % 64)
										v69 = int32(7)
										v74 = v69
										v75 = v68 | v52<<(uint(int64(40))%64)
										v81 = v74
										v82 = v75 | v47<<(uint(int64(32))%64)
										v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v25))))
										v98 = v81
										v99 = v82 | v83
									}
								}
							}
						}
					} else {
						v88 = int32(3)
						v89 = v41 << (uint(int64(16)) % 64)
						v92 = v88
						v93 = v34<<(uint(int64(8))%64) | v89
						v98 = v92
						v99 = base.I64_extend8_s(base.I64_extend_i32_u(v30)) | v93
					}
				}
			}
			v104 = (v99 ^ int64(base.Ui64(v99)>>(uint(int64(23))%64))) * int64(2388976653695081527)
			v110 = (v27 ^ int64(base.Ui64(v104)>>(uint(int64(47))%64)) ^ v104) * int64(-8645972361240307355)
			v111 = v25 + v98
			v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
			if v112 != 0 {
				v25 = v111
				v27 = v110
				v30 = v112
				continue
			} else {
				break
			}
			break
		}
		v113 = v111
		v115 = v110
	}
	v127 = (base.I64_extend_i32_s(v113-v21) + int64(base.Ui64(v115)>>(uint(int64(23))%64)) ^ v115) * int64(2388976653695081527)
	return base.I32_wrap_i64(int64(base.Ui64(v127)>>(uint(int64(47))%64)) ^ v127 - int64(base.Ui64(v127)>>(uint(int64(32))%64)))
}
func F_spgbulkdelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	if l1 == int32(0) {
		v13 = F_palloc0(m, int32(40))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = v13
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v17
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
			F_spgvacuumscan(m, v8)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(112)
				return v17
			}
		}
	} else {
		v17 = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
		F_spgvacuumscan(m, v8)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(112)
			return v17
		}
	}
}
func F_spggetbitmap(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v4)+200)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+196)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(v4)+208)) = uint8(v3)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_spgWalk(m, v10, v4, int32(1), int32(265))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v4)+200))
		return v17
	}
}
func F_spgistBuildCallback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v35 int64
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	v7 = int32(_a_F_spgistBuildCallback_0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_spgistBuildCallback[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	*(*int32)(unsafe.Add(mBase, _c_F_spgistBuildCallback[0])) = v10
	v12 = F_spgdoinsert(m, l0, l5, l1, l2, l3)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v12 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	goto L6
L4:
	;
	goto L5
L5:
	;
	v35 = *(*int64)(unsafe.Add(mBase, uint32(l5)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+88)) = v35 + int64(1)
	*(*int32)(unsafe.Add(mBase, _c_F_spgistBuildCallback[0])) = v8
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	F_MemoryContextReset(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L11
	}
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	F_MemoryContextReset(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v25 = F_spgdoinsert(m, l0, l5, l1, l2, l3)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v25 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	return
}
func F_spool_tuples(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int64
	_ = v38
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v119 int64
	_ = v119
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v16 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return
L2:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+385)))
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	if v22 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v28 = int64(-1)
	goto L6
L6:
	;
	v29 = int32(_a_F_spool_tuples_0)
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_spool_tuples[0]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_spool_tuples[0])) = v34
	if v28 != int64(-1) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v27 = int64(-1)
	goto L9
L8:
	;
	v27 = l1
	goto L9
L9:
	;
	v28 = v27
	goto L6
L10:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spool_tuples[0])) = v30
	goto L1
L11:
	;
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	if v28 < v38 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L15
L14:
	;
	goto L13
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v31)+52))
	if v51 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L10
L17:
	;
	F_ExecReScan(m, v31)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v54 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v31)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L20
	} else {
		goto L24
	}
L20:
	;
	return
L21:
	;
	goto L19
L22:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	if v108 == int32(3) {
		goto L37
	} else {
		goto L38
	}
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+387)) = uint8(v99)
	v103 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+385)) = uint8(v103)
	goto L10
L24:
	;
	if v56 == int32(0) {
		v99 = v54
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
	if v60&int32(2) != 0 {
		v99 = v54
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
	if v63 <= int32(0) {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+380))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+12)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v67)+8)) = v66
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v70 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	F_MemoryContextReset(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L20
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v76 = int32(_a_F_spool_tuples_0)
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_spool_tuples[0]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_spool_tuples[0])) = v79
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v70)+24))
	v84 = m.T0[v83].(func(*base.Module, int32, int32, int32) int64)(m, v70, v67, v14+int32(15))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L20
	} else {
		goto L32
	}
L31:
	;
	goto L22
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spool_tuples[0])) = v77
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	F_MemoryContextReset(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L20
	} else {
		goto L33
	}
L33:
	;
	if v84 != int64(0) {
		goto L22
	} else {
		goto L34
	}
L34:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+32))
	m.T0[v95].(func(*base.Module, int32, int32))(m, v93, v56)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L20
	} else {
		goto L35
	}
L35:
	;
	v99 = int32(1)
	goto L23
L36:
	;
	if base.B2i32(v28 == int64(-1))|base.B2i32(v119 <= v28) != 0 {
		goto L15
	} else {
		goto L41
	}
L37:
	;
	v111 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	v119 = v111
	goto L36
L38:
	;
	goto L39
L39:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	F_tuplestore_puttupleslot(m, v112, v56)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L20
	} else {
		goto L40
	}
L40:
	;
	v115 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	v117 = v115 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v117
	v119 = v117
	goto L36
L41:
	;
	goto L16
}
func F_statapprox_heap_read_stream_next(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if base.Ui32(v5) < base.Ui32(v6) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	return int32(-1)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v10 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_statapprox_heap_read_stream_next[0]))
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v24 = F_visibilitymap_get_status(m, v23, v10, l1+int32(20))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L11
	}
L9:
	;
	return int32(0)
L10:
	;
	goto L8
L11:
	;
	if v24&int32(1) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v30 + int32(1)
	return v10
L13:
	;
	goto L14
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v36 = F_GetRecordedFreeSpace(m, v35, v10)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v39 + base.I64_extend_i32_u(int32(_a_F_statapprox_heap_read_stream_next_0)-v36)
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v38)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+64)) = v45 + base.I64_extend_i32_u(v36)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if base.Ui32(v49) < base.Ui32(v50) {
		v10 = v49
		goto L4
	} else {
		goto L16
	}
L16:
	;
	goto L5
}
func F_statatt_build_stavalues(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = *(*int64)(unsafe.Add(mBase, _c_F_statatt_build_stavalues[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_statatt_build_stavalues[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v16
	v18 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+29)) = uint8(v18)
	v21 = F_text_to_cstring(m, base.I32_wrap_i64(l2))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int64(0)
	} else {
		v29 = F_InputFunctionCallSafe(m, l1, v21, l3, l4, v10+int32(24), v10+int32(40))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int64(0)
		} else {
			F_pfree(m, v21)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int64(0)
			} else {
				if v29 == int32(0) {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(19)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
					F_ThrowErrorData(m, v38)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int64(0)
					} else {
						v41 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v41)
						v98 = int64(0)
						m.G0 = v10 + int32(48)
						return v98
					}
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
					v45 = F_pg_detoast_datum(m, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int64(0)
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
						if v47 != int32(1) {
							v52 = F_errstart(m, int32(19), int32(0))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int64(0)
							} else {
								if v52 != 0 {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
										F_errmsg(m, int32(_a_F_statatt_build_stavalues_0), v10+int32(16))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_statatt_build_stavalues_1), int32(592), int32(_a_F_statatt_build_stavalues_2))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int64(0)
											} else {
												v68 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v68)
												v98 = int64(0)
												m.G0 = v10 + int32(48)
												return v98
											}
										}
									}
								} else {
									v68 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v68)
									v98 = int64(0)
									m.G0 = v10 + int32(48)
									return v98
								}
							}
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
							v72 = F_pg_detoast_datum(m, v71)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int64(0)
							} else {
								v74 = F_array_contains_nulls(m, v72)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int64(0)
								} else {
									if v74 != 0 {
										v78 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int64(0)
										} else {
											if v78 != 0 {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
													F_errmsg(m, int32(_a_F_statatt_build_stavalues_3), v10)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_statatt_build_stavalues_1), int32(601), int32(_a_F_statatt_build_stavalues_2))
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int64(0)
														} else {
															v92 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v92)
															v98 = int64(0)
															m.G0 = v10 + int32(48)
															return v98
														}
													}
												}
											} else {
												v92 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v92)
												v98 = int64(0)
												m.G0 = v10 + int32(48)
												return v98
											}
										}
									} else {
										v95 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v95)
										v97 = *(*int64)(unsafe.Add(mBase, uint32(v10)+40))
										v98 = v97
										m.G0 = v10 + int32(48)
										return v98
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
func F_statatt_get_elem_type(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	if l0 == int32(3614) {
		v6 = int32(25)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v6
		v16 = v6
		v18 = F_lookup_type_cache(m, v16, int32(1))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
			if v20 == int32(0) {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v20
				return int32(1)
			}
		}
	} else {
		v9 = F_get_base_element_type(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
			if v9 != 0 {
				v16 = v9
				v18 = F_lookup_type_cache(m, v16, int32(1))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
					if v20 == int32(0) {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v20
						return int32(1)
					}
				}
			} else {
				return int32(0)
			}
		}
	}
}
func F_statatt_set_slot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int64, l7 int32, l8 int64, l9 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int64
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	v11 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	v24 = v22 & int64(65535)
	v25 = int32(_a_F_statatt_set_slot_0)
	if l3&v25 == base.I32_wrap_i64(v22)&v25 {
		if v24 != int64(0) {
			v35 = int32(-1)
		} else {
			v35 = int32(0)
		}
		v110 = v35
		v111 = v11
		v113 = v11
	} else {
		v36 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)))
		if v24 != int64(0) {
			if base.I32_wrap_i64(v36) != l3&int32(_a_F_statatt_set_slot_0) {
				if v36 != int64(0) {
					v62 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)))
					if v62 == int64(0) {
						v65 = int32(2)
					} else {
						v65 = int32(-1)
					}
					v66 = v65
					v67 = v62
				} else {
					v58 = int32(1)
					v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
					v66 = v58
					v67 = v59
				}
				v68 = int32(_a_F_statatt_set_slot_0)
				v69 = l3 & v68
				if v69 == base.I32_wrap_i64(v67)&v68 {
					v110 = v66
					v111 = int32(2)
					v113 = v11
				} else {
					v75 = int32(3)
					v78 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
					if v78&int64(65535) == int64(0) {
						v83 = v75
					} else {
						v83 = int32(-1)
					}
					if v66 < int32(0) {
						v86 = v83
					} else {
						v86 = v66
					}
					if v69 == base.I32_wrap_i64(v78)&int32(_a_F_statatt_set_slot_0) {
						v110 = v86
						v111 = v75
						v113 = v11
					} else {
						v93 = int32(_a_F_statatt_set_slot_0)
						v95 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
						v99 = base.B2i32(l3&v93 != base.I32_wrap_i64(v95)&v93)
						if l3&v93 != base.I32_wrap_i64(v95)&v93 {
							v100 = int32(5)
						} else {
							v100 = int32(4)
						}
						if v95&int64(65535) == int64(0) {
							v106 = int32(4)
						} else {
							v106 = v86
						}
						if v86 < int32(0) {
							v109 = v106
						} else {
							v109 = v86
						}
						v110 = v109
						v111 = v100
						v113 = v99
					}
				}
			} else {
				v43 = int32(1)
				if v36 == int64(0) {
					v48 = v43
				} else {
					v48 = int32(-1)
				}
				v110 = v48
				v111 = v43
				v113 = v11
			}
		} else {
			if base.I32_wrap_i64(v36) != l3&int32(_a_F_statatt_set_slot_0) {
				v58 = int32(0)
				v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
				v66 = v58
				v67 = v59
				v68 = int32(_a_F_statatt_set_slot_0)
				v69 = l3 & v68
				if v69 == base.I32_wrap_i64(v67)&v68 {
					v110 = v66
					v111 = int32(2)
					v113 = v11
				} else {
					v75 = int32(3)
					v78 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
					if v78&int64(65535) == int64(0) {
						v83 = v75
					} else {
						v83 = int32(-1)
					}
					if v66 < int32(0) {
						v86 = v83
					} else {
						v86 = v66
					}
					if v69 == base.I32_wrap_i64(v78)&int32(_a_F_statatt_set_slot_0) {
						v110 = v86
						v111 = v75
						v113 = v11
					} else {
						v93 = int32(_a_F_statatt_set_slot_0)
						v95 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
						v99 = base.B2i32(l3&v93 != base.I32_wrap_i64(v95)&v93)
						if l3&v93 != base.I32_wrap_i64(v95)&v93 {
							v100 = int32(5)
						} else {
							v100 = int32(4)
						}
						if v95&int64(65535) == int64(0) {
							v106 = int32(4)
						} else {
							v106 = v86
						}
						if v86 < int32(0) {
							v109 = v106
						} else {
							v109 = v86
						}
						v110 = v109
						v111 = v100
						v113 = v99
					}
				}
			} else {
				v110 = v11
				v111 = int32(1)
				v113 = v11
			}
		}
	}
	if int32(0) <= v110 {
		v117 = v110
	} else {
		v117 = v111
	}
	if v113 != 0 {
		v118 = v117
	} else {
		v118 = v111
	}
	if v118 < int32(5) {
		v122 = v118 << (uint(int32(16)) % 32)
		v124 = v122 + int32(_a_F_statatt_set_slot_1)
		v127 = l0 + int32(base.Ui32(v124)>>(uint(int32(13))%32))
		v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127))))
		if v128 != l3&int32(_a_F_statatt_set_slot_0) {
			*(*int64)(unsafe.Add(mBase, uint32(v127))) = base.I64_extend_i32_s(l3)
			v137 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v124)>>(uint(int32(16))%32))))) = uint8(v137)
		} else {
		}
		v140 = v122 + int32(_a_F_statatt_set_slot_2)
		v143 = l0 + int32(base.Ui32(v140)>>(uint(int32(13))%32))
		v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
		if v144 != l4 {
			*(*int64)(unsafe.Add(mBase, uint32(v143))) = base.I64_extend_i32_u(l4)
			v151 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v140)>>(uint(int32(16))%32))))) = uint8(v151)
		} else {
		}
		v154 = v122 - int32(-1048576)
		v157 = l0 + int32(base.Ui32(v154)>>(uint(int32(13))%32))
		v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
		if v158 != l5 {
			*(*int64)(unsafe.Add(mBase, uint32(v157))) = base.I64_extend_i32_u(l5)
			v165 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v154)>>(uint(int32(16))%32))))) = uint8(v165)
		} else {
		}
		if l7 == int32(0) {
			v170 = v118 + int32(21)
			*(*int64)(unsafe.Add(mBase, uint32(l0+v170<<(uint(int32(3))%32)))) = l6
			v176 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1+v170))) = uint8(v176)
			v179 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2+v170))) = uint8(v179)
		} else {
		}
		if l9 == int32(0) {
			v185 = v118 + int32(26)
			*(*int64)(unsafe.Add(mBase, uint32(l0+v185<<(uint(int32(3))%32)))) = l8
			v191 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1+v185))) = uint8(v191)
			v194 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2+v185))) = uint8(v194)
		} else {
		}
		m.G0 = v20 + int32(16)
		return
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v203 = m.ExcPending
		if v203 != 0 {
			return
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v206 = m.ExcPending
			if v206 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = v118 + int32(1)
				F_errmsg(m, int32(_a_F_statatt_set_slot_3), v20)
				mBase = m.M
				v212 = m.ExcPending
				if v212 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_statatt_set_slot_4), int32(661), int32(_a_F_statatt_set_slot_5))
					mBase = m.M
					v217 = m.ExcPending
					if v217 != 0 {
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
func F_std_fetch_func(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v94 int64
	_ = v94
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v103 int64
	_ = v103
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13+l1<<(uint(int32(2))%32))))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	if int32(0) < v19 {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
		v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+18)))
		if base.Ui32(v23&int32(2047)) < base.Ui32(v19) {
			v27 = F_getmissingattr(m, v18, v19, l2)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				v103 = v27
				m.G0 = v11 + int32(16)
				return v103
			}
		} else {
			v31 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v31)
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+20)))
			if v34&int32(1) == v31 {
				v43 = v18 + v19<<(uint(int32(3))%32) + int32(20)
				v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43))))
				if int32(0) <= v44 {
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
					v49 = v33 + v47 + v44
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+4)))
					if v50 == int32(1) {
						v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+2)))
						if base.I32_popcnt(v53) != int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v53
								F_errmsg_internal(m, int32(_a_F_std_fetch_func_0), v11)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_std_fetch_func_1), int32(123), int32(_a_F_std_fetch_func_2))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							switch base.I32_ctz(v53) {
							case 0:
								v58 = int64(*(*int8)(unsafe.Add(mBase, uint32(v49))))
								v103 = v58
								m.G0 = v11 + int32(16)
								return v103
							case 1:
								v59 = int64(*(*int16)(unsafe.Add(mBase, uint32(v49))))
								v103 = v59
								m.G0 = v11 + int32(16)
								return v103
							case 2:
								v60 = int64(*(*int32)(unsafe.Add(mBase, uint32(v49))))
								v103 = v60
								m.G0 = v11 + int32(16)
								return v103
							case 3:
								v61 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
								v103 = v61
								m.G0 = v11 + int32(16)
								return v103
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v53
									F_errmsg_internal(m, int32(_a_F_std_fetch_func_0), v11)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_std_fetch_func_1), int32(123), int32(_a_F_std_fetch_func_2))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
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
					} else {
						v103 = base.I64_extend_i32_u(v49)
						m.G0 = v11 + int32(16)
						return v103
					}
				} else {
					v76 = F_nocachegetattr(m, v17, v19, v18)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int64(0)
					} else {
						v103 = v76
						m.G0 = v11 + int32(16)
						return v103
					}
				}
			} else {
				v78 = int32(1)
				v79 = v19 - v78
				v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+int32(base.Ui32(v79)>>(uint(int32(3))%32)))+23)))
				if int32(base.Ui32(v83)>>(uint(v79&int32(7))%32))&v78 == int32(0) {
					v91 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v91)
					v103 = int64(0)
					m.G0 = v11 + int32(16)
					return v103
				} else {
					v94 = F_nocachegetattr(m, v17, v19, v18)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int64(0)
					} else {
						v103 = v94
						m.G0 = v11 + int32(16)
						return v103
					}
				}
			}
		}
	} else {
		v96 = F_heap_getsysattr(m, v17, v19, l2)
		mBase = m.M
		v97 = m.ExcPending
		if v97 != 0 {
			return int64(0)
		} else {
			v103 = v96
			m.G0 = v11 + int32(16)
			return v103
		}
	}
}
func F_std_typanalyze(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10 < int32(0) {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_std_typanalyze[0]))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v14
	} else {
	}
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = int32(0)
	F_get_sort_group_operators(m, v16, v17, v17, v17, v8+int32(12), v8+int32(8), v17, v17)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int32(0)
	} else {
		v31 = F_palloc(m, int32(12))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v31))) = v33
			if v33 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = int32(0)
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v39
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v31
				v61 = int32(542)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v61
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v63 * int32(300)
				m.G0 = v8 + int32(16)
				return int32(1)
			} else {
				v42 = F_get_opcode(m, v33)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v42
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v46
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v31
					if v44 == int32(0) {
						v61 = int32(542)
					} else {
						if v46 != 0 {
							v53 = int32(540)
						} else {
							v53 = int32(541)
						}
						if v44 != 0 {
							v55 = v53
						} else {
							v55 = int32(541)
						}
						v61 = v55
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v61
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v63 * int32(300)
					m.G0 = v8 + int32(16)
					return int32(1)
				}
			}
		}
	}
}
func F_storage_name(m *base.Module, l0 int32) int32 {
	var v13 int32
	_ = v13
	switch l0 - int32(101) {
	case 0:
		return int32(_a_F_storage_name_0)
	default:
		v13 = int32(_a_F_storage_name_1)
		return v13
	case 8:
		return int32(_a_F_storage_name_2)
	case 11:
		v13 = int32(_a_F_storage_name_3)
		return v13
	case 19:
		return int32(_a_F_storage_name_4)
	}
}
func F_store_match(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	if base.Ui32(int32(2)) <= base.Ui32(l2) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v8 < v9 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v25 = v8
			v26 = v11
			v27 = int32(3)
			*(*int32)(unsafe.Add(mBase, uint32(v26+v25<<(uint(v27)%32)))) = l1
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v37 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v31+v32<<(uint(v27)%32))+4)) = l1 + l2 - v37
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v40 + v37
			return v37
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v9 << (uint(int32(1)) % 32)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v18 = F_emscripten_builtin_realloc(m, v15, v9<<(uint(int32(4))%32))
			mBase = m.M
			if v18 == int32(0) {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v18
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v25 = v24
				v26 = v18
				v27 = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v26+v25<<(uint(v27)%32)))) = l1
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v37 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v31+v32<<(uint(v27)%32))+4)) = l1 + l2 - v37
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v40 + v37
				return v37
			}
		}
	} else {
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v46 < v47 {
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v63 = v46
			v64 = v49
			v65 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v63 + v65
			*(*int32)(unsafe.Add(mBase, uint32(v64+v63<<(uint(int32(2))%32)))) = l1
			return v65
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v47 << (uint(int32(1)) % 32)
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v56 = F_emscripten_builtin_realloc(m, v53, v47<<(uint(int32(3))%32))
			mBase = m.M
			if v56 == int32(0) {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v56
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v63 = v62
				v64 = v56
				v65 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v63 + v65
				*(*int32)(unsafe.Add(mBase, uint32(v64+v63<<(uint(int32(2))%32)))) = l1
				return v65
			}
		}
	}
}
func F_str_toupper(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v4 {
		v64 = v4
		m.G0 = v10 + int32(16)
		return v64
	} else {
		if l2 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(34209924))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_str_toupper_0)
					F_errmsg(m, int32(_a_F_str_toupper_1), v10)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						F_errhint(m, int32(_a_F_str_toupper_2), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_str_toupper_3), int32(1703), int32(_a_F_str_toupper_4))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			v16 = F_pg_newlocale_from_collation(m, l2)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+2)))
				if v20 == int32(1) {
					v23 = F_pnstrdup(m, l0, l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
						if v25 == int32(0) {
							v64 = v23
						} else {
							v28 = v25
							v30 = v23
							for {
								if base.Ui32((v28-int32(97))&int32(255)) < base.Ui32(int32(26)) {
									v43 = v28 - int32(32)
								} else {
									v43 = v28
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v43)
								v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
								if v45 != 0 {
									v28 = v45
									v30 = v30 + int32(1)
									continue
								} else {
									break
								}
								break
							}
							v64 = v23
						}
						m.G0 = v10 + int32(16)
						return v64
					}
				} else {
					v49 = l1 + int32(1)
					v50 = F_palloc(m, v49)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v52 = F_pg_strupper(m, v50, v49, l0, l1, v16)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							v55 = v52 + int32(1)
							if base.Ui32(v55) <= base.Ui32(v49) {
								v64 = v50
								m.G0 = v10 + int32(16)
								return v64
							} else {
								v57 = F_repalloc(m, v50, v55)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									v59 = F_pg_strupper(m, v57, v55, l0, l1, v16)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										v64 = v57
										m.G0 = v10 + int32(16)
										return v64
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
func F_strcspn(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int64
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return v260 - l0
L2:
	;
	v107 = int32(0)
	v108 = int32(32)
	goto L32
L3:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v11 != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = v10 & int32(255)
	if v16 != 0 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	goto L5
L7:
	;
	v260 = v106
	goto L1
L8:
	;
	v106 = v96
	goto L7
L9:
	;
	v86 = v81
	goto L26
L10:
	;
	v81 = v73
	goto L9
L11:
	;
	if l0&int32(3) != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v71 = F_strlen(m, l0)
	mBase = m.M
	v106 = v71 + l0
	goto L7
L14:
	;
	v19 = l0
	goto L17
L15:
	;
	v33 = l0
	goto L16
L16:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v42 = int32(-2139062144)
	if (int32(16843008)-v39|v39)&v42 != v42 {
		v73 = v33
		goto L10
	} else {
		goto L21
	}
L17:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if base.B2i32(v24 == int32(0))|base.B2i32(v16 == v24) != 0 {
		v96 = v19
		goto L8
	} else {
		goto L19
	}
L18:
	;
	v33 = v30
	goto L16
L19:
	;
	v30 = v19 + int32(1)
	if v30&int32(3) != 0 {
		v19 = v30
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v48 = v33
	v50 = v39
	goto L22
L22:
	;
	v54 = v50 ^ v16*int32(16843009)
	v57 = int32(-2139062144)
	if (int32(16843008)-v54|v54)&v57 != v57 {
		v73 = v48
		goto L10
	} else {
		goto L24
	}
L23:
	;
	v81 = v63
	goto L9
L24:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v63 = v48 + int32(4)
	v67 = int32(-2139062144)
	if (v61|(int32(16843008)-v61))&v67 == v67 {
		v48 = v63
		v50 = v61
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v88 == int32(0) {
		v96 = v86
		goto L8
	} else {
		goto L28
	}
L27:
	;
	v96 = v86
	goto L8
L28:
	;
	if v88 != v10&int32(255) {
		v86 = v86 + int32(1)
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v215 != 0 {
		goto L42
	} else {
		goto L43
	}
L31:
	;
	goto L30
L32:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v107)
	v115 = v8 + v108
	*(*uint8)(unsafe.Add(mBase, uint32(v115-int32(1)))) = uint8(v107)
	goto L33
L33:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+2)) = uint8(v107)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)) = uint8(v107)
	*(*uint8)(unsafe.Add(mBase, uint32(v115-int32(3)))) = uint8(v107)
	*(*uint8)(unsafe.Add(mBase, uint32(v115-int32(2)))) = uint8(v107)
	goto L34
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+3)) = uint8(v107)
	*(*uint8)(unsafe.Add(mBase, uint32(v115-int32(4)))) = uint8(v107)
	goto L35
L35:
	;
	v137 = int32(0)
	v140 = (v137 - v8) & int32(3)
	v141 = v8 + v140
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v137
	v149 = (v108 - v140) & int32(-4)
	v150 = v141 + v149
	*(*int32)(unsafe.Add(mBase, uint32(v150-int32(4)))) = v137
	if base.Ui32(v149) < base.Ui32(int32(9)) {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141)+8)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v141)+4)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v150-int32(8)))) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v150-int32(12)))) = v137
	if base.Ui32(v149) < base.Ui32(int32(25)) {
		goto L31
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141)+24)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v141)+20)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v141)+16)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v141)+12)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v150-int32(16)))) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v150-int32(20)))) = v137
	v176 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v150-v176))) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v150-int32(28)))) = v137
	v185 = v141&int32(4) | v176
	v186 = v149 - v185
	if base.Ui32(v186) < base.Ui32(int32(32)) {
		goto L31
	} else {
		goto L38
	}
L38:
	;
	v191 = base.I64_extend_i32_u(v137) * int64(4294967297)
	v194 = v185 + v141
	v195 = v186
	goto L39
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v194)+24)) = v191
	*(*int64)(unsafe.Add(mBase, uint32(v194)+16)) = v191
	*(*int64)(unsafe.Add(mBase, uint32(v194)+8)) = v191
	*(*int64)(unsafe.Add(mBase, uint32(v194))) = v191
	v203 = int32(32)
	v206 = v195 - v203
	if base.Ui32(int32(31)) < base.Ui32(v206) {
		v194 = v194 + v203
		v195 = v206
		goto L39
	} else {
		goto L41
	}
L40:
	;
	goto L31
L41:
	;
	goto L40
L42:
	;
	v217 = l1
	v218 = v215
	goto L45
L43:
	;
	goto L44
L44:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v239 == int32(0) {
		v260 = l0
		goto L1
	} else {
		goto L48
	}
L45:
	;
	v225 = v8 + int32(base.Ui32(v218)>>(uint(int32(3))%32))&int32(28)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v227 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v225))) = v226 | v227<<(uint(v218)%32)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+1)))
	if v231 != 0 {
		v217 = v217 + v227
		v218 = v231
		goto L45
	} else {
		goto L47
	}
L46:
	;
	goto L44
L47:
	;
	goto L46
L48:
	;
	v243 = l0
	v244 = v239
	goto L49
L49:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(base.Ui32(v244)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v252)>>(uint(v244)%32))&int32(1) != 0 {
		v260 = v243
		goto L1
	} else {
		goto L51
	}
L50:
	;
	v260 = v258
	goto L1
L51:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+1)))
	v258 = v243 + int32(1)
	if v256 != 0 {
		v243 = v258
		v244 = v256
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
}
func F_strict_word_similarity_dist_commutator_op(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14408(m, l0, int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_strlcpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	if l2 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v119 = F_strlen(m, v115)
	mBase = m.M
	return v119 + (v116 - l0)
L2:
	;
	v115 = l1
	v116 = l0
	goto L1
L3:
	;
	goto L4
L4:
	;
	v9 = l2 - int32(1)
	if (l0^l1)&int32(3) != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v112)
	v115 = v108
	v116 = v109
	goto L1
L6:
	;
	v93 = v88
	v94 = v89
	v95 = v90
	goto L27
L7:
	;
	if v83 == int32(0) {
		v108 = v81
		v109 = v82
		goto L5
	} else {
		goto L26
	}
L8:
	;
	v81 = l1
	v82 = l0
	v83 = v9
	goto L7
L9:
	;
	goto L10
L10:
	;
	v13 = int32(0)
	if base.B2i32(l1&int32(3) == v13)|base.B2i32(v9 == v13) == v13 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v49 == int32(0) {
		v108 = v46
		v109 = v47
		goto L5
	} else {
		goto L20
	}
L12:
	;
	v25 = l1
	v26 = l0
	v27 = v9
	goto L15
L13:
	;
	goto L14
L14:
	;
	v46 = l1
	v47 = l0
	v48 = v9
	v49 = base.B2i32(v9 != v13)
	goto L11
L15:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v29)
	if v29 == int32(0) {
		v88 = v25
		v89 = v26
		v90 = v27
		goto L6
	} else {
		goto L17
	}
L16:
	;
	v46 = v40
	v47 = v34
	v48 = v36
	v49 = v38
	goto L11
L17:
	;
	v33 = int32(1)
	v34 = v26 + v33
	v36 = v27 - v33
	v37 = int32(0)
	v38 = base.B2i32(v36 != v37)
	v40 = v25 + v33
	if v40&int32(3) == v37 {
		v46 = v40
		v47 = v34
		v48 = v36
		v49 = v38
		goto L11
	} else {
		goto L18
	}
L18:
	;
	if v36 != 0 {
		v25 = v40
		v26 = v34
		v27 = v36
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if base.B2i32(v52 == int32(0))|base.B2i32(base.Ui32(v48) < base.Ui32(int32(4))) != 0 {
		v81 = v46
		v82 = v47
		v83 = v48
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v59 = v46
	v60 = v47
	v61 = v48
	goto L22
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v67 = int32(-2139062144)
	if (int32(16843008)-v64|v64)&v67 != v67 {
		v88 = v59
		v89 = v60
		v90 = v61
		goto L6
	} else {
		goto L24
	}
L23:
	;
	v81 = v75
	v82 = v73
	v83 = v77
	goto L7
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v64
	v72 = int32(4)
	v73 = v60 + v72
	v75 = v59 + v72
	v77 = v61 - v72
	if base.Ui32(int32(3)) < base.Ui32(v77) {
		v59 = v75
		v60 = v73
		v61 = v77
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v88 = v81
	v89 = v82
	v90 = v83
	goto L6
L27:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	*(*uint8)(unsafe.Add(mBase, uint32(v94))) = uint8(v97)
	if v97 == int32(0) {
		v108 = v93
		v109 = v94
		goto L5
	} else {
		goto L29
	}
L28:
	;
	v108 = v104
	v109 = v102
	goto L5
L29:
	;
	v101 = int32(1)
	v102 = v94 + v101
	v104 = v93 + v101
	v106 = v95 - v101
	if v106 != 0 {
		v93 = v104
		v94 = v102
		v95 = v106
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
}
func F_strnxfrm_libc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	v9 = m.G0
	v11 = v9 - int32(1024)
	m.G0 = v11
	v14 = l3 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v14) {
		v17 = F_palloc(m, v14)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = v17
			if l3 != 0 {
				base.MemoryCopy(m, v21, l2, l3)
			} else {
			}
			v24 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3+v21))) = uint8(v24)
			v28 = F_strlen(m, v21)
			mBase = m.M
			if base.Ui32(v28) < base.Ui32(l1) {
				v30 = F_strcpy(m, l0, v21)
				mBase = m.M
			} else {
			}
			if v21 != v11 {
				F_pfree(m, v21)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					m.G0 = v11 + int32(1024)
					return v28
				}
			} else {
				m.G0 = v11 + int32(1024)
				return v28
			}
		}
	} else {
		v21 = v11
		if l3 != 0 {
			base.MemoryCopy(m, v21, l2, l3)
		} else {
		}
		v24 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l3+v21))) = uint8(v24)
		v28 = F_strlen(m, v21)
		mBase = m.M
		if base.Ui32(v28) < base.Ui32(l1) {
			v30 = F_strcpy(m, l0, v21)
			mBase = m.M
		} else {
		}
		if v21 != v11 {
			F_pfree(m, v21)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v11 + int32(1024)
				return v28
			}
		} else {
			m.G0 = v11 + int32(1024)
			return v28
		}
	}
}
func F_strtox_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int64
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v146 int32
	_ = v146
	var v150 int64
	_ = v150
	var v156 int64
	_ = v156
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v170 int64
	_ = v170
	var v177 int64
	_ = v177
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v192 int64
	_ = v192
	var v195 int64
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int64
	_ = v203
	var v209 int32
	_ = v209
	var v216 int64
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int64
	_ = v224
	var v225 int64
	_ = v225
	var v239 int32
	_ = v239
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v248 int64
	_ = v248
	v5 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	if l2 <= int32(36) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v18 + int32(16)
	return v248
L2:
	;
	v81 = int32(16)
	if l2|v81 != v81 {
		goto L19
	} else {
		goto L20
	}
L3:
	;
	v31 = l0
	v32 = v22
	goto L9
L4:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v22 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strtox_2[0])) = int32(28)
	v248 = int64(0)
	goto L1
L7:
	;
	v70 = l0
	v73 = v5
	goto L2
L8:
	;
	v56 = v32 & int32(255)
	switch v56 - int32(43) {
	case 0, 2:
		goto L14
	default:
		v70 = v31
		v73 = v5
		goto L2
	}
L9:
	;
	v42 = base.I32_extend8_s(v32)
	goto L11
L10:
	;
	v70 = v54
	v73 = v5
	goto L2
L11:
	;
	if base.B2i32(v42 == int32(32))|base.B2i32(base.Ui32(v42-int32(9)) < base.Ui32(int32(5))) == int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v54 = v31 + int32(1)
	if v52 != 0 {
		v31 = v54
		v32 = v52
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	if v56 == int32(45) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v63 = int32(-1)
	goto L17
L16:
	;
	v63 = int32(0)
	goto L17
L17:
	;
	v70 = v31 + int32(1)
	v73 = v63
	goto L2
L18:
	;
	v106 = base.I64_extend_i32_u(v105)
	v110 = int32(0)
	v112 = v103
	v117 = v104
	v119 = int64(0)
	goto L31
L19:
	;
	if l2 != 0 {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v85 != int32(48) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v88 = int32(1)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v89&int32(223) == int32(88) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v103 = v70 + int32(2)
	v104 = v88
	v105 = int32(16)
	goto L18
L23:
	;
	goto L24
L24:
	;
	if l2 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v100 = l2
	goto L27
L26:
	;
	v100 = int32(8)
	goto L27
L27:
	;
	v103 = v70 + int32(1)
	v104 = v88
	v105 = v100
	goto L18
L28:
	;
	v102 = l2
	goto L30
L29:
	;
	v102 = int32(10)
	goto L30
L30:
	;
	v103 = v70
	v104 = v5
	v105 = v102
	goto L18
L31:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	v125 = v123 - int32(48)
	if base.Ui32(v125&int32(255)) < base.Ui32(int32(10)) {
		v146 = v125
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if l1 != 0 {
		goto L45
	} else {
		goto L46
	}
L33:
	;
	goto L32
L34:
	;
	if v105 <= v146&int32(255) {
		goto L33
	} else {
		goto L40
	}
L35:
	;
	if base.Ui32((v123-int32(97))&int32(255)) <= base.Ui32(int32(25)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v146 = v123 - int32(87)
	goto L34
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(int32(25)) < base.Ui32((v123-int32(65))&int32(255)) {
		goto L33
	} else {
		goto L39
	}
L39:
	;
	v146 = v123 - int32(55)
	goto L34
L40:
	;
	v150 = int64(0)
	v156 = int64(32)
	v157 = int64(base.Ui64(v119) >> (uint(v156) % 64))
	v159 = int64(base.Ui64(v106) >> (uint(v156) % 64))
	v162 = int64(4294967295)
	v163 = v119 & v162
	v165 = v106 & v162
	v166 = v163 * v165
	v170 = int64(base.Ui64(v166)>>(uint(v156)%64)) + v163*v159
	v177 = v165*v157 + v170&v162
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v106*v150 + v150*v119 + v157*v159 + int64(base.Ui64(v170)>>(uint(v156)%64)) + int64(base.Ui64(v177)>>(uint(v156)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v166&v162 | v177<<(uint(v156)%64)
	goto L41
L41:
	;
	v188 = int32(1)
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	if v189 != int64(0) {
		v201 = v188
		v202 = v117
		v203 = v119
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v110 = v201
	v112 = v112 + int32(1)
	v117 = v202
	v119 = v203
	goto L31
L43:
	;
	v192 = v119 * v106
	v195 = base.I64_extend_i32_u(v146) & int64(255)
	if base.Ui64(v195^int64(-1)) < base.Ui64(v192) {
		v201 = v188
		v202 = v117
		v203 = v119
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v201 = v110
	v202 = int32(1)
	v203 = v192 + v195
	goto L42
L45:
	;
	if v117 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	if v110 != 0 {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v209 = v112
	goto L50
L49:
	;
	v209 = l0
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v209
	goto L47
L51:
	;
	v242 = base.I64_extend_i32_s(v239)
	v248 = v240 ^ v242 - v242
	goto L1
L52:
	;
	if base.I32_wrap_i64(v225)|v223 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strtox_2[0])) = int32(68)
	v216 = l3 & int64(1)
	if v216 == int64(0) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	if base.Ui64(v119) < base.Ui64(l3) {
		v239 = v73
		v240 = v119
		goto L51
	} else {
		goto L59
	}
L56:
	;
	v219 = v73
	goto L58
L57:
	;
	v219 = int32(0)
	goto L58
L58:
	;
	v223 = v219
	v224 = l3
	v225 = v216
	goto L52
L59:
	;
	v223 = v73
	v224 = v119
	v225 = l3 & int64(1)
	goto L52
L60:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strtox_2[0])) = int32(68)
	v248 = l3 - int64(1)
	goto L1
L61:
	;
	goto L62
L62:
	;
	if base.Ui64(v224) <= base.Ui64(l3) {
		v239 = v223
		v240 = v224
		goto L51
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strtox_2[0])) = int32(68)
	v248 = l3
	goto L1
}
func F_subarray(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v13 == int32(3) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v17 = v16
		} else {
			v17 = int32(0)
		}
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
		if v19 != 0 {
			v20 = F_array_contains_nulls(m, v9)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				if v20 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_subarray_0), int32(0))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_subarray_1), int32(288), int32(_a_F_subarray_2))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					v24 = v9 + int32(16)
					v25 = F_ArrayGetNItemsSafe(m, v22, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						if v25 == int32(0) {
							v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v110 != v9 {
								F_pfree(m, v9)
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int64(0)
								} else {
									v115 = F_new_intArrayType(m, int32(0))
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return int64(0)
									} else {
										return base.I64_extend_i32_u(v115)
									}
								}
							} else {
								v115 = F_new_intArrayType(m, int32(0))
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int64(0)
								} else {
									return base.I64_extend_i32_u(v115)
								}
							}
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							v30 = F_ArrayGetNItemsSafe(m, v29, v24)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int64(0)
							} else {
								v35 = v18 - base.B2i32(int32(0) < v18)
								v39 = v35>>(uint(int32(31))%32)&v30 + v35
								if v17 != 0 {
									v41 = v39 + v17
								} else {
									v41 = v30
								}
								if v17 < int32(0) {
									v44 = v30 + v17
								} else {
									v44 = v41
								}
								if v44 < v30 {
									v46 = v44
								} else {
									v46 = v30
								}
								v47 = int32(0)
								if v47 < v39 {
									v50 = v39
								} else {
									v50 = v47
								}
								if v46 <= v50 {
									v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									if v110 != v9 {
										F_pfree(m, v9)
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int64(0)
										} else {
											v115 = F_new_intArrayType(m, int32(0))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return int64(0)
											} else {
												return base.I64_extend_i32_u(v115)
											}
										}
									} else {
										v115 = F_new_intArrayType(m, int32(0))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int64(0)
										} else {
											return base.I64_extend_i32_u(v115)
										}
									}
								} else {
									v52 = v46 - v50
									v53 = F_new_intArrayType(m, v52)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int64(0)
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
										if v55 == int32(0) {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
											v65 = (v58<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										} else {
											v65 = v55
										}
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
										if v66 == int32(0) {
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
											v76 = (v69<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										} else {
											v76 = v66
										}
										v78 = v52 << (uint(int32(2)) % 32)
										if v78 != 0 {
											base.MemoryCopy(m, v53+v65, v9+v76+v50<<(uint(int32(2))%32), v78)
										} else {
										}
										v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										if v85 != v9 {
											F_pfree(m, v9)
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return int64(0)
											} else {
												return base.I64_extend_i32_u(v53)
											}
										} else {
											return base.I64_extend_i32_u(v53)
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v24 = v9 + int32(16)
			v25 = F_ArrayGetNItemsSafe(m, v22, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int64(0)
			} else {
				if v25 == int32(0) {
					v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v110 != v9 {
						F_pfree(m, v9)
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int64(0)
						} else {
							v115 = F_new_intArrayType(m, int32(0))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v115)
							}
						}
					} else {
						v115 = F_new_intArrayType(m, int32(0))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v115)
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					v30 = F_ArrayGetNItemsSafe(m, v29, v24)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						v35 = v18 - base.B2i32(int32(0) < v18)
						v39 = v35>>(uint(int32(31))%32)&v30 + v35
						if v17 != 0 {
							v41 = v39 + v17
						} else {
							v41 = v30
						}
						if v17 < int32(0) {
							v44 = v30 + v17
						} else {
							v44 = v41
						}
						if v44 < v30 {
							v46 = v44
						} else {
							v46 = v30
						}
						v47 = int32(0)
						if v47 < v39 {
							v50 = v39
						} else {
							v50 = v47
						}
						if v46 <= v50 {
							v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v110 != v9 {
								F_pfree(m, v9)
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int64(0)
								} else {
									v115 = F_new_intArrayType(m, int32(0))
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return int64(0)
									} else {
										return base.I64_extend_i32_u(v115)
									}
								}
							} else {
								v115 = F_new_intArrayType(m, int32(0))
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int64(0)
								} else {
									return base.I64_extend_i32_u(v115)
								}
							}
						} else {
							v52 = v46 - v50
							v53 = F_new_intArrayType(m, v52)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int64(0)
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
								if v55 == int32(0) {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
									v65 = (v58<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								} else {
									v65 = v55
								}
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								if v66 == int32(0) {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
									v76 = (v69<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								} else {
									v76 = v66
								}
								v78 = v52 << (uint(int32(2)) % 32)
								if v78 != 0 {
									base.MemoryCopy(m, v53+v65, v9+v76+v50<<(uint(int32(2))%32), v78)
								} else {
								}
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								if v85 != v9 {
									F_pfree(m, v9)
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return int64(0)
									} else {
										return base.I64_extend_i32_u(v53)
									}
								} else {
									return base.I64_extend_i32_u(v53)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_subpath(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v12 == int32(3) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v16 = v15
		} else {
			v16 = int32(0)
		}
		if v11 < int32(0) {
			v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+4)))
			v21 = v19 + v11
		} else {
			v21 = v11
		}
		if v16 < int32(0) {
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+4)))
			v33 = v16 + v24
		} else {
			if v12 != int32(3) {
				v29 = int32(_a_F_subpath_0)
			} else {
				v29 = v21
			}
			if v16 == int32(0) {
				v33 = v29
			} else {
				v33 = v21 + v16
			}
		}
		v34 = F_inner_subltree(m, v7, v21, v33)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int64(0)
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v36 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v34)
				}
			} else {
				return base.I64_extend_i32_u(v34)
			}
		}
	}
}
func F_subtrans_errdetail_for_io_error(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14249(m, l0, int32(_a_F_subtrans_errdetail_for_io_error_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_subxact_info_write(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	v4 = m.G0
	v6 = v4 - int32(1040)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l1
	v14 = F_pg_snprintf(m, v6+int32(16), int32(1024), int32(_a_F_subxact_info_write_0), v6)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_write[0]))
		if v17 == int32(0) {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_write[1]))
			if v21 != 0 {
				F_pfree(m, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v25 = int64(0)
					*(*int64)(unsafe.Add(mBase, _c_F_subxact_info_write[2])) = v25
					*(*int64)(unsafe.Add(mBase, _c_F_subxact_info_write[0])) = v25
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_write[3]))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+60))
					F_BufFileDeleteFileSet(m, v32, v6+int32(16), int32(1))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						m.G0 = v6 + int32(1040)
						return
					}
				}
			} else {
				v25 = int64(0)
				*(*int64)(unsafe.Add(mBase, _c_F_subxact_info_write[2])) = v25
				*(*int64)(unsafe.Add(mBase, _c_F_subxact_info_write[0])) = v25
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_write[3]))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+60))
				F_BufFileDeleteFileSet(m, v32, v6+int32(16), int32(1))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					m.G0 = v6 + int32(1040)
					return
				}
			}
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_write[3]))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+60))
			v42 = v6 + int32(16)
			v45 = F_BufFileOpenFileSet(m, v40, v42, int32(2), int32(1))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				if v45 == int32(0) {
					v50 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_write[3]))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+60))
					v52 = F_BufFileCreateFileSet(m, v51, v42)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						v54 = v52
						v55 = int32(_a_F_subxact_info_write_1)
						v56 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_write[0]))
						F_BufFileWrite(m, v54, v55, int32(4))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_write[1]))
							F_BufFileWrite(m, v54, v62, v56<<(uint(int32(4))%32))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								F_BufFileClose(m, v54)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_write[1]))
									if v70 != 0 {
										F_pfree(m, v70)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return
										} else {
											v74 = int64(0)
											*(*int64)(unsafe.Add(mBase, _c_F_subxact_info_write[2])) = v74
											*(*int64)(unsafe.Add(mBase, _c_F_subxact_info_write[0])) = v74
											m.G0 = v6 + int32(1040)
											return
										}
									} else {
										v74 = int64(0)
										*(*int64)(unsafe.Add(mBase, _c_F_subxact_info_write[2])) = v74
										*(*int64)(unsafe.Add(mBase, _c_F_subxact_info_write[0])) = v74
										m.G0 = v6 + int32(1040)
										return
									}
								}
							}
						}
					}
				} else {
					v54 = v45
					v55 = int32(_a_F_subxact_info_write_1)
					v56 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_write[0]))
					F_BufFileWrite(m, v54, v55, int32(4))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_write[1]))
						F_BufFileWrite(m, v54, v62, v56<<(uint(int32(4))%32))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							F_BufFileClose(m, v54)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_write[1]))
								if v70 != 0 {
									F_pfree(m, v70)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										v74 = int64(0)
										*(*int64)(unsafe.Add(mBase, _c_F_subxact_info_write[2])) = v74
										*(*int64)(unsafe.Add(mBase, _c_F_subxact_info_write[0])) = v74
										m.G0 = v6 + int32(1040)
										return
									}
								} else {
									v74 = int64(0)
									*(*int64)(unsafe.Add(mBase, _c_F_subxact_info_write[2])) = v74
									*(*int64)(unsafe.Add(mBase, _c_F_subxact_info_write[0])) = v74
									m.G0 = v6 + int32(1040)
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
func F_superuser_arg(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_superuser_arg[0]))
	if base.B2i32(v5 == v2)|base.B2i32(l0 != v5) == v2 {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[1])))
		v56 = v13
		return v56 & int32(1)
	} else {
		if l0 == int32(10) {
			v16 = int32(1)
			v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[2])))
			if v18&v16 == int32(0) {
				v56 = v16
				return v56 & int32(1)
			} else {
				v27 = F_SearchSysCache1(m, int32(11), base.I64_extend_i32_u(l0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					if v27 != 0 {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
						v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v32)+68)))
						F_ReleaseCatCache(m, v27)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v37 = v34
							v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[3])))
							if v39 == int32(0) {
								F_CacheRegisterSyscacheCallback(m, int32(11), int32(1991), int64(0))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									v48 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[3])) = uint8(v48)
									*(*int32)(unsafe.Add(mBase, _c_F_superuser_arg[0])) = l0
									v54 = v37 & int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[1])) = uint8(v54)
									v56 = v37
									return v56 & int32(1)
								}
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_superuser_arg[0])) = l0
								v54 = v37 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[1])) = uint8(v54)
								v56 = v37
								return v56 & int32(1)
							}
						}
					} else {
						v37 = int32(0)
						v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[3])))
						if v39 == int32(0) {
							F_CacheRegisterSyscacheCallback(m, int32(11), int32(1991), int64(0))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v48 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[3])) = uint8(v48)
								*(*int32)(unsafe.Add(mBase, _c_F_superuser_arg[0])) = l0
								v54 = v37 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[1])) = uint8(v54)
								v56 = v37
								return v56 & int32(1)
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_superuser_arg[0])) = l0
							v54 = v37 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[1])) = uint8(v54)
							v56 = v37
							return v56 & int32(1)
						}
					}
				}
			}
		} else {
			v27 = F_SearchSysCache1(m, int32(11), base.I64_extend_i32_u(l0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				if v27 != 0 {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v32)+68)))
					F_ReleaseCatCache(m, v27)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v37 = v34
						v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[3])))
						if v39 == int32(0) {
							F_CacheRegisterSyscacheCallback(m, int32(11), int32(1991), int64(0))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v48 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[3])) = uint8(v48)
								*(*int32)(unsafe.Add(mBase, _c_F_superuser_arg[0])) = l0
								v54 = v37 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[1])) = uint8(v54)
								v56 = v37
								return v56 & int32(1)
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_superuser_arg[0])) = l0
							v54 = v37 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[1])) = uint8(v54)
							v56 = v37
							return v56 & int32(1)
						}
					}
				} else {
					v37 = int32(0)
					v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[3])))
					if v39 == int32(0) {
						F_CacheRegisterSyscacheCallback(m, int32(11), int32(1991), int64(0))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v48 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[3])) = uint8(v48)
							*(*int32)(unsafe.Add(mBase, _c_F_superuser_arg[0])) = l0
							v54 = v37 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[1])) = uint8(v54)
							v56 = v37
							return v56 & int32(1)
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_superuser_arg[0])) = l0
						v54 = v37 & int32(1)
						*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[1])) = uint8(v54)
						v56 = v37
						return v56 & int32(1)
					}
				}
			}
		}
	}
}
func F_suppress_redundant_updates_trigger(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L36
	} else {
		goto L49
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L36
	} else {
		goto L45
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L36
	} else {
		goto L41
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L36
	} else {
		goto L37
	}
L5:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v8 != int32(448) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v11&int32(3) != int32(2) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if v11&int32(24) != int32(8) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	if v11&int32(4) == int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v25 != v27 {
		v115 = v24
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return base.I64_extend_i32_u(v115)
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
	if v30 != v32 {
		v115 = v24
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+18)))
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+18)))
	if (v34^v35)&int32(2047) != 0 {
		v115 = v24
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+20)))
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+20)))
	if (v39^v40)&int32(15) != 0 {
		v115 = v24
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v45 = int32(23)
	v46 = v29 + v45
	v48 = v31 + v45
	v50 = v25 - v45
	if base.Ui32(int32(4)) <= base.Ui32(v50) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	if v112 != 0 {
		goto L33
	} else {
		goto L34
	}
L16:
	;
	v112 = int32(0)
	goto L15
L17:
	;
	v86 = v81
	v87 = v82
	v88 = v83
	goto L27
L18:
	;
	if (v46|v48)&int32(3) != 0 {
		v81 = v46
		v82 = v48
		v83 = v50
		goto L17
	} else {
		goto L21
	}
L19:
	;
	v74 = v46
	v75 = v48
	v76 = v50
	goto L20
L20:
	;
	if v76 == int32(0) {
		goto L16
	} else {
		goto L26
	}
L21:
	;
	v58 = v46
	v59 = v48
	v60 = v50
	goto L22
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v63 != v64 {
		v81 = v58
		v82 = v59
		v83 = v60
		goto L17
	} else {
		goto L24
	}
L23:
	;
	v74 = v69
	v75 = v67
	v76 = v71
	goto L20
L24:
	;
	v66 = int32(4)
	v67 = v59 + v66
	v69 = v58 + v66
	v71 = v60 - v66
	if base.Ui32(int32(3)) < base.Ui32(v71) {
		v58 = v69
		v59 = v67
		v60 = v71
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v81 = v74
	v82 = v75
	v83 = v76
	goto L17
L27:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v91 == v92 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v112 = v91 - v92
	goto L15
L29:
	;
	v94 = int32(1)
	v99 = v88 - v94
	if v99 != 0 {
		v86 = v86 + v94
		v87 = v87 + v94
		v88 = v99
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	goto L16
L33:
	;
	v113 = v24
	goto L35
L34:
	;
	v113 = int32(0)
	goto L35
L35:
	;
	v115 = v113
	goto L10
L36:
	;
	return int64(0)
L37:
	;
	F_errcode(m, int32(16908867))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(_a_F_suppress_redundant_updates_trigger_0), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_suppress_redundant_updates_trigger_1), int32(41), int32(_a_F_suppress_redundant_updates_trigger_2))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_errcode(m, int32(16908867))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L36
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(_a_F_suppress_redundant_updates_trigger_3), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L36
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_suppress_redundant_updates_trigger_1), int32(47), int32(_a_F_suppress_redundant_updates_trigger_2))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errcode(m, int32(16908867))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L36
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(_a_F_suppress_redundant_updates_trigger_4), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L36
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_suppress_redundant_updates_trigger_1), int32(53), int32(_a_F_suppress_redundant_updates_trigger_2))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L36
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errcode(m, int32(16908867))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L36
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_suppress_redundant_updates_trigger_5), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L36
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_suppress_redundant_updates_trigger_1), int32(59), int32(_a_F_suppress_redundant_updates_trigger_2))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L36
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
