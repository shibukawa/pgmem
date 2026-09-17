package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	v10 = int32(1)
	v11 = l0 + v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v14 = v12 & v10
	if v12 == v10 {
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
		if v20 == int32(18) {
			v23 = int32(16)
		} else {
			v23 = int32(0)
		}
		if base.Ui32((v20-int32(1))&int32(255)) < base.Ui32(int32(3)) {
			v30 = int32(4)
		} else {
			v30 = v23
		}
		v41 = v30
	} else {
		v31 = int32(1)
		if v14 != 0 {
			v41 = int32(base.Ui32(v12)>>(uint(v31)%32)) - v31
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v41 = int32(base.Ui32(v35)>>(uint(int32(2))%32)) - int32(4)
		}
	}
	v42 = int32(1)
	v43 = l1 + v42
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v44 == v42 {
		v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
		if base.Ui32((v47-int32(1))&int32(255)) < base.Ui32(int32(3)) {
			v157 = F_palloc(m, v41<<(uint(int32(1))%32)+int32(4))
			mBase = m.M
			v158 = m.ExcPending
			if v158 != 0 {
				return int32(0)
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v166 = m.ExcPending
				if v166 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(84410498))
					mBase = m.M
					v169 = m.ExcPending
					if v169 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_SB_do_like_escape_0), int32(0))
						mBase = m.M
						v173 = m.ExcPending
						if v173 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(_a_F_SB_do_like_escape_1), int32(0))
							mBase = m.M
							v177 = m.ExcPending
							if v177 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_SB_do_like_escape_2), int32(438), int32(_a_F_SB_do_like_escape_3))
								mBase = m.M
								v182 = m.ExcPending
								if v182 != 0 {
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
			if v47 == int32(18) {
				v58 = int32(16)
			} else {
				v58 = int32(0)
			}
			v71 = v58
			if v14 != 0 {
				v74 = v11
			} else {
				v74 = l0 + int32(4)
			}
			v79 = F_palloc(m, v41<<(uint(int32(1))%32)+int32(4))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				v84 = v79 + int32(4)
				switch v71 {
				case 0:
					if v41 <= int32(0) {
						v240 = v84
					} else {
						if v41&int32(1) != 0 {
							v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
							if v89 == int32(92) {
								v92 = int32(92)
								*(*uint8)(unsafe.Add(mBase, uint32(v79)+4)) = uint8(v92)
								v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
								v97 = v96
								v98 = v79 + int32(5)
							} else {
								v97 = v89
								v98 = v84
							}
							*(*uint8)(unsafe.Add(mBase, uint32(v98))) = uint8(v97)
							v100 = int32(1)
							v107 = v98 + v100
							v108 = v74 + v100
							v109 = v41 - v100
						} else {
							v107 = v84
							v108 = v74
							v109 = v41
						}
						if v41 == int32(1) {
							v240 = v107
						} else {
							v112 = v109
							v114 = v107
							v116 = v108
							for {
								v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
								if v121 == int32(92) {
									v124 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v124)
									v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
									v129 = v128
									v130 = v114 + int32(1)
								} else {
									v129 = v121
									v130 = v114
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v129)
								v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
								if v132 != int32(92) {
									v142 = v132
									v143 = v130 + int32(1)
								} else {
									v137 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)) = uint8(v137)
									v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
									v142 = v139
									v143 = v130 + int32(2)
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v143))) = uint8(v142)
								v146 = v143 + int32(1)
								v147 = int32(2)
								if v147 < v112 {
									v112 = v112 - v147
									v114 = v146
									v116 = v116 + v147
									continue
								} else {
									break
								}
								break
							}
							v240 = v146
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v79))) = (v240 - v79) << (uint(int32(2)) % 32)
					return v79
				case 1:
					v183 = int32(1)
					v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
					v187 = v185 & v183
					if v187 != 0 {
						v188 = v183
					} else {
						v188 = int32(4)
					}
					v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v188))))
					if v190 == int32(92) {
						v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
						if v251 == int32(1) {
							v255 = int32(18)
							v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
							if v257 == v255 {
								v260 = v255
							} else {
								v260 = int32(2)
							}
							if base.Ui32((v257-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v267 = int32(6)
							} else {
								v267 = v260
							}
							v276 = v267
						} else {
							v268 = int32(1)
							if v251&v268 != 0 {
								v276 = int32(base.Ui32(v251) >> (uint(v268) % 32))
							} else {
								v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v276 = int32(base.Ui32(v272) >> (uint(int32(2)) % 32))
							}
						}
						if v276 == int32(0) {
							return v79
						} else {
							base.MemoryCopy(m, v79, l0, v276)
							return v79
						}
					} else {
						v193 = int32(0)
						if v41 <= v193 {
							v240 = v84
						} else {
							if v187 != 0 {
								v198 = v43
							} else {
								v198 = l1 + int32(4)
							}
							v201 = v84
							v202 = v41
							v203 = v74
							v205 = v193
							for {
								v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
								v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
								v213 = base.B2i32(v205 == int32(0)) & base.B2i32(v210 == v211)
								if v213 != 0 {
									v214 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v201))) = uint8(v214)
									v231 = v201 + int32(1)
								} else {
									v219 = v201 + int32(1)
									if v210 == int32(92) {
										v222 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v201))) = uint8(v222)
										if v205&int32(1) != 0 {
											v231 = v219
										} else {
											v226 = int32(92)
											*(*uint8)(unsafe.Add(mBase, uint32(v201)+1)) = uint8(v226)
											v231 = v201 + int32(2)
										}
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(v201))) = uint8(v210)
										v231 = v219
									}
								}
								v232 = int32(1)
								if v232 < v202 {
									v201 = v231
									v202 = v202 - v232
									v203 = v203 + v232
									v205 = v213
									continue
								} else {
									break
								}
								break
							}
							v240 = v231
						}
						*(*int32)(unsafe.Add(mBase, uint32(v79))) = (v240 - v79) << (uint(int32(2)) % 32)
						return v79
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v166 = m.ExcPending
					if v166 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(84410498))
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_SB_do_like_escape_0), int32(0))
							mBase = m.M
							v173 = m.ExcPending
							if v173 != 0 {
								return int32(0)
							} else {
								F_errhint(m, int32(_a_F_SB_do_like_escape_1), int32(0))
								mBase = m.M
								v177 = m.ExcPending
								if v177 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_SB_do_like_escape_2), int32(438), int32(_a_F_SB_do_like_escape_3))
									mBase = m.M
									v182 = m.ExcPending
									if v182 != 0 {
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
		v59 = int32(1)
		if v44&v59 != 0 {
			v71 = int32(base.Ui32(v44)>>(uint(v59)%32)) - v59
		} else {
			v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v71 = int32(base.Ui32(v65)>>(uint(int32(2))%32)) - int32(4)
		}
		if v14 != 0 {
			v74 = v11
		} else {
			v74 = l0 + int32(4)
		}
		v79 = F_palloc(m, v41<<(uint(int32(1))%32)+int32(4))
		mBase = m.M
		v82 = m.ExcPending
		if v82 != 0 {
			return int32(0)
		} else {
			v84 = v79 + int32(4)
			switch v71 {
			case 0:
				if v41 <= int32(0) {
					v240 = v84
				} else {
					if v41&int32(1) != 0 {
						v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
						if v89 == int32(92) {
							v92 = int32(92)
							*(*uint8)(unsafe.Add(mBase, uint32(v79)+4)) = uint8(v92)
							v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
							v97 = v96
							v98 = v79 + int32(5)
						} else {
							v97 = v89
							v98 = v84
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v98))) = uint8(v97)
						v100 = int32(1)
						v107 = v98 + v100
						v108 = v74 + v100
						v109 = v41 - v100
					} else {
						v107 = v84
						v108 = v74
						v109 = v41
					}
					if v41 == int32(1) {
						v240 = v107
					} else {
						v112 = v109
						v114 = v107
						v116 = v108
						for {
							v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
							if v121 == int32(92) {
								v124 = int32(92)
								*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v124)
								v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
								v129 = v128
								v130 = v114 + int32(1)
							} else {
								v129 = v121
								v130 = v114
							}
							*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v129)
							v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
							if v132 != int32(92) {
								v142 = v132
								v143 = v130 + int32(1)
							} else {
								v137 = int32(92)
								*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)) = uint8(v137)
								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
								v142 = v139
								v143 = v130 + int32(2)
							}
							*(*uint8)(unsafe.Add(mBase, uint32(v143))) = uint8(v142)
							v146 = v143 + int32(1)
							v147 = int32(2)
							if v147 < v112 {
								v112 = v112 - v147
								v114 = v146
								v116 = v116 + v147
								continue
							} else {
								break
							}
							break
						}
						v240 = v146
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(v79))) = (v240 - v79) << (uint(int32(2)) % 32)
				return v79
			case 1:
				v183 = int32(1)
				v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				v187 = v185 & v183
				if v187 != 0 {
					v188 = v183
				} else {
					v188 = int32(4)
				}
				v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v188))))
				if v190 == int32(92) {
					v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if v251 == int32(1) {
						v255 = int32(18)
						v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
						if v257 == v255 {
							v260 = v255
						} else {
							v260 = int32(2)
						}
						if base.Ui32((v257-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v267 = int32(6)
						} else {
							v267 = v260
						}
						v276 = v267
					} else {
						v268 = int32(1)
						if v251&v268 != 0 {
							v276 = int32(base.Ui32(v251) >> (uint(v268) % 32))
						} else {
							v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v276 = int32(base.Ui32(v272) >> (uint(int32(2)) % 32))
						}
					}
					if v276 == int32(0) {
						return v79
					} else {
						base.MemoryCopy(m, v79, l0, v276)
						return v79
					}
				} else {
					v193 = int32(0)
					if v41 <= v193 {
						v240 = v84
					} else {
						if v187 != 0 {
							v198 = v43
						} else {
							v198 = l1 + int32(4)
						}
						v201 = v84
						v202 = v41
						v203 = v74
						v205 = v193
						for {
							v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
							v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
							v213 = base.B2i32(v205 == int32(0)) & base.B2i32(v210 == v211)
							if v213 != 0 {
								v214 = int32(92)
								*(*uint8)(unsafe.Add(mBase, uint32(v201))) = uint8(v214)
								v231 = v201 + int32(1)
							} else {
								v219 = v201 + int32(1)
								if v210 == int32(92) {
									v222 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v201))) = uint8(v222)
									if v205&int32(1) != 0 {
										v231 = v219
									} else {
										v226 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v201)+1)) = uint8(v226)
										v231 = v201 + int32(2)
									}
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v201))) = uint8(v210)
									v231 = v219
								}
							}
							v232 = int32(1)
							if v232 < v202 {
								v201 = v231
								v202 = v202 - v232
								v203 = v203 + v232
								v205 = v213
								continue
							} else {
								break
							}
							break
						}
						v240 = v231
					}
					*(*int32)(unsafe.Add(mBase, uint32(v79))) = (v240 - v79) << (uint(int32(2)) % 32)
					return v79
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v166 = m.ExcPending
				if v166 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(84410498))
					mBase = m.M
					v169 = m.ExcPending
					if v169 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_SB_do_like_escape_0), int32(0))
						mBase = m.M
						v173 = m.ExcPending
						if v173 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(_a_F_SB_do_like_escape_1), int32(0))
							mBase = m.M
							v177 = m.ExcPending
							if v177 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_SB_do_like_escape_2), int32(438), int32(_a_F_SB_do_like_escape_3))
								mBase = m.M
								v182 = m.ExcPending
								if v182 != 0 {
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
	var v59 int32
	_ = v59
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
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
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
	v141 = m.ExcPending
	if v141 != 0 {
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
	v59 = v37
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
	if v73 < v59 {
		v55 = v76
		v58 = v74
		v59 = v59 - v73
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
	if v81 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	F_s_lock(m, v15, int32(_a_F_SIInsertDataEntries_1), int32(422), int32(_a_F_SIInsertDataEntries_2))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v89 = v19 - v37
	v90 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v74
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_c_F_SIInsertDataEntries[2])))
	if v90 < v93 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_c_F_SIInsertDataEntries[3])))
	v101 = int32(0)
	goto L25
L23:
	;
	goto L24
L24:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_SIInsertDataEntries[1]))
	F_LWLockRelease(m, v132+int32(768))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L5
	} else {
		goto L28
	}
L25:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v96+v101<<(uint(int32(2))%32))))
	v116 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13+v110<<(uint(int32(4))%32))+uint32(_c_F_SIInsertDataEntries[4]))) = uint8(v116)
	v119 = v101 + v116
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_c_F_SIInsertDataEntries[2])))
	if v119 < v120 {
		v101 = v119
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
	if int32(0) < v89 {
		v18 = v76
		v19 = v89
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
func F_SampleNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v281 int64
	_ = v281
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)))
	if v15 == v2 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v363
L2:
	;
	v357 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+152)) = uint16(v357)
	v363 = v190
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L12
	} else {
		goto L91
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L12
	} else {
		goto L88
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L12
	} else {
		goto L84
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = int64(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
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
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+8))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
	m.T0[v187].(func(*base.Module, int32))(m, v185)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L12
	} else {
		goto L51
	}
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v27 = v23 << (uint(int32(2)) % 32)
	goto L11
L10:
	;
	v27 = int32(0)
	goto L11
L11:
	;
	v28 = F_palloc(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v32 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v98 != 0 {
		goto L28
	} else {
		goto L29
	}
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v35 <= int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v41 = v2
	goto L17
L17:
	;
	v49 = v41 << (uint(int32(2)) % 32)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49+v50)))
	v53 = int32(_a_F_SampleNext_0)
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_SampleNext[0]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_SampleNext[0])) = v56
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	v61 = m.T0[v60].(func(*base.Module, int32, int32, int32) int32)(m, v52, v18, v13+int32(15))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L12
	} else {
		goto L20
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L12
	} else {
		goto L23
	}
L19:
	;
	goto L18
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SampleNext[0])) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v28+v49))) = v61
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	if v67 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v69 = v41 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v69 < v70 {
		v41 = v69
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L14
L23:
	;
	F_errcode(m, int32(403177602))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	F_errmsg(m, int32(_a_F_SampleNext_1), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_SampleNext_2), int32(244), int32(_a_F_SampleNext_3))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	v122 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+132)) = uint16(v122)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v125 != 0 {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v99 = int32(_a_F_SampleNext_0)
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_SampleNext[0]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_SampleNext[0])) = v102
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	v107 = m.T0[v106].(func(*base.Module, int32, int32, int32) int32)(m, v98, v18, v13+int32(15))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L12
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v121 = v118
	goto L27
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SampleNext[0])) = v100
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	if v111 == int32(1) {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v116 = F_DirectFunctionCall1Coll(m, int32(748), int32(0), v107)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	v121 = v116
	goto L27
L34:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v128 = v126
	goto L36
L35:
	;
	v128 = int32(0)
	goto L36
L36:
	;
	m.T0[v124].(func(*base.Module, int32, int32, int32, int32))(m, l0, v28, v128, v121)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v132 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	F_pfree(m, v28)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L12
	} else {
		goto L50
	}
L39:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	v138 = int32(0)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)))
	if v146 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v157 = int32(0)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+188))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+16))
	m.T0[v165].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v132, v157, int32(1), v159, base.B2i32(v131 == v157), v162)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L12
	} else {
		goto L49
	}
L42:
	;
	v147 = int32(68)
	goto L44
L43:
	;
	v147 = int32(4)
	goto L44
L44:
	;
	if v131 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v150 = v147
	goto L47
L46:
	;
	v150 = v147 | int32(128)
	goto L47
L47:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v135)+188))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
	v154 = m.T0[v153].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v135, v137, v138, v138, v138, v141<<(uint(int32(8))%32)|v150)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v154
	goto L38
L49:
	;
	goto L38
L50:
	;
	v172 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)) = uint8(v172)
	goto L8
L51:
	;
	v190 = int32(0)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+153)))
	if v191 != 0 {
		v363 = v190
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)))
	if v192 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_SampleNext[1]))
	if v196 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_SampleNext[1]))
	if v213 != 0 {
		goto L62
	} else {
		goto L63
	}
L56:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SampleNext[2])))
	if v198&int32(1) == int32(0) {
		goto L4
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+188))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+172))
	v206 = m.T0[v205].(func(*base.Module, int32, int32) int32)(m, v184, l0)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L12
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	if v206 == int32(0) {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v210 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)) = uint8(v210)
	goto L55
L62:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SampleNext[2])))
	if v215&int32(1) == int32(0) {
		goto L3
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+188))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+176))
	v223 = m.T0[v222].(func(*base.Module, int32, int32, int32) int32)(m, v184, l0, v185)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L12
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	if v223 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	goto L70
L68:
	;
	goto L69
L69:
	;
	v281 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = v281 + int64(1)
	v363 = v185
	goto L1
L70:
	;
	v237 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)) = uint8(v237)
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_SampleNext[1]))
	if v240 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L69
L72:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SampleNext[2])))
	if v242&int32(1) == int32(0) {
		goto L4
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+188))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+172))
	v250 = m.T0[v249].(func(*base.Module, int32, int32) int32)(m, v184, l0)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L12
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	if v250 == int32(0) {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	v254 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)) = uint8(v254)
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_SampleNext[1]))
	if v257 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SampleNext[2])))
	if v259&int32(1) == int32(0) {
		goto L3
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+188))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+176))
	v267 = m.T0[v266].(func(*base.Module, int32, int32, int32) int32)(m, v184, l0, v185)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L12
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	if v267 == int32(0) {
		goto L70
	} else {
		goto L83
	}
L83:
	;
	goto L71
L84:
	;
	F_errcode(m, int32(386400386))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L12
	} else {
		goto L85
	}
L85:
	;
	F_errmsg(m, int32(_a_F_SampleNext_4), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L12
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_SampleNext_2), int32(256), int32(_a_F_SampleNext_3))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L12
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
	F_errmsg_internal(m, int32(_a_F_SampleNext_5), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L12
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_SampleNext_6), int32(1972), int32(_a_F_SampleNext_7))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L12
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
	F_errmsg_internal(m, int32(_a_F_SampleNext_8), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L12
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_SampleNext_6), int32(1995), int32(_a_F_SampleNext_9))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L12
	} else {
		goto L93
	}
L93:
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	v3 = int32(0)
	if l1 != 0 {
		if l1 == int32(1) {
			v50 = l0
			v51 = int32(0)
			v52 = v3
			v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
			v60 = v58 | int32(32)
			v68 = v51*int32(257) + v60
			v69 = v60 + v52*int32(17)
		} else {
			v17 = l0
			v18 = int32(0)
			v19 = v3
			v24 = v3
			for {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
				v26 = int32(32)
				v27 = v25 | v26
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				v30 = v28 | v26
				v31 = int32(17)
				v36 = v27 + (v30+v19*v31)*v31
				v37 = int32(257)
				v42 = (v18*v37+v30)*v37 + v27
				v43 = int32(2)
				v44 = v17 + v43
				v46 = v24 + v43
				if v46 != l1&int32(-2) {
					v17 = v44
					v18 = v42
					v19 = v36
					v24 = v46
					continue
				} else {
					break
				}
				break
			}
			if l1&int32(1) == int32(0) {
				v68 = v42
				v69 = v36
			} else {
				v50 = v44
				v51 = v42
				v52 = v36
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
				v60 = v58 | int32(32)
				v68 = v51*int32(257) + v60
				v69 = v60 + v52*int32(17)
			}
		}
		v75 = int32(989)
		v76 = base.I32_rem_u_s(v69, v75)
		v78 = base.I32_rem_u_s(v68, v75)
		v82 = v76
		v88 = v78
	} else {
		v82 = v3
		v88 = int32(0)
	}
	v89 = int32(1)
	v91 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82<<(uint(v89)%32))+uint32(_c_F_ScanKeywords_hash_func[0]))))
	v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88<<(uint(v89)%32))+uint32(_c_F_ScanKeywords_hash_func[0]))))
	return v91 + v94
}
func F_SelectNeighbors(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
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
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 float32
	_ = v146
	var v147 float64
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v166 int32
	_ = v166
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 float32
	_ = v238
	var v239 float64
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v276 int32
	_ = v276
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 float32
	_ = v307
	var v308 float64
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v341 int32
	_ = v341
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 float32
	_ = v372
	var v373 float64
	_ = v373
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v450 int32
	_ = v450
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v547 int32
	_ = v547
	var v557 int32
	_ = v557
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	v8 = l7
	v9 = int32(0)
	v20 = F_list_copy(m, l1)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v20 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	return v591
L4:
	;
	v35 = F_mul_size(m, int32(4), v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L10
	}
L5:
	;
	v27 = int32(0)
	if v27 <= l2 {
		v591 = v27
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if l2 < v30 {
		v33 = v30
		goto L4
	} else {
		goto L9
	}
L8:
	;
	v33 = v27
	goto L4
L9:
	;
	return v20
L10:
	;
	v37 = F_palloc(m, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v8 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if l0 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if v20 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v41 = int32(_a_F_SelectNeighbors_0)
	goto L17
L16:
	;
	v41 = int32(_a_F_SelectNeighbors_1)
	goto L17
L17:
	;
	F_list_sort(m, v20, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v588
	v591 = v569
	goto L3
L20:
	;
	if l6 == int32(0) {
		v591 = v547
		goto L3
	} else {
		goto L145
	}
L21:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v8)
	v47 = int32(0)
	v547 = v47
	v557 = v47
	goto L20
L22:
	;
	goto L23
L23:
	;
	v53 = int32(0)
	v63 = v20
	v64 = v9
	v66 = v9
	v67 = v9
	goto L25
L24:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v8)
	if v501 <= int32(0) {
		v547 = v487
		v557 = v497
		goto L20
	} else {
		goto L133
	}
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v71 <= int32(0) {
		v487 = v53
		v497 = v63
		v501 = v67
		goto L24
	} else {
		goto L27
	}
L26:
	;
	v487 = v483
	v497 = int32(0)
	v501 = v484
	goto L24
L27:
	;
	if v53 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v76 = v74
	goto L30
L29:
	;
	v76 = int32(0)
	goto L30
L30:
	;
	if l2 <= v76 {
		v487 = v53
		v497 = v63
		v501 = v67
		goto L24
	} else {
		goto L31
	}
L31:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v78+v71<<(uint(int32(2))%32)-int32(4))))
	v85 = F_list_delete_last(m, v63)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v24&int32(1) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+8)))
	if v472 == int32(1) {
		goto L128
	} else {
		goto L129
	}
L34:
	;
	if l0 != 0 {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	goto L36
L36:
	;
	if v66 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L37:
	;
	v100 = int32(1)
	if v53 == int32(0) {
		v166 = v100
		goto L43
	} else {
		goto L44
	}
L38:
	;
	v99 = l0 + v91 - int32(1)
	goto L37
L39:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0+v89)+87))
	if v91 != 0 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+88))
	v99 = v94
	goto L37
L42:
	;
	v99 = int32(0)
	goto L37
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+8)) = uint8(v166)
	v465 = v64
	v467 = v66
	goto L33
L44:
	;
	v103 = int32(0)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v104 <= v103 {
		v166 = v100
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v115 = v103
	goto L46
L46:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126+v115<<(uint(int32(2))%32))))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if l0 != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v166 = v151
	goto L43
L48:
	;
	v144 = F_FunctionCall2Coll(m, v131, v132, v99, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L54
	}
L49:
	;
	v143 = l0 + v135 - int32(1)
	goto L48
L50:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0+v133)+87))
	if v135 != 0 {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+88))
	v143 = v138
	goto L48
L53:
	;
	v143 = int32(0)
	goto L48
L54:
	;
	v146 = *(*float32)(unsafe.Add(mBase, uint32(v84)+4))
	v147 = *(*float64)(unsafe.Add(mBase, uint32(v144)))
	v149 = base.F32_ge(v146, base.F32_demote_f64(v147))
	v151 = base.B2i32(v149 == int32(0))
	if v149 != 0 {
		v166 = v151
		goto L43
	} else {
		goto L55
	}
L55:
	;
	v153 = v115 + int32(1)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v153 < v154 {
		v115 = v153
		goto L46
	} else {
		goto L56
	}
L56:
	;
	goto L47
L57:
	;
	v450 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+8)) = uint8(v450)
	v465 = int32(1)
	v467 = v66
	goto L33
L58:
	;
	v426 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+8)) = uint8(v426)
	v429 = F_lappend(m, v66, v84)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L126
	}
L59:
	;
	if l5 != v84 {
		v465 = v64
		v467 = v66
		goto L33
	} else {
		goto L102
	}
L60:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v178 <= int32(0) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+8)))
	if v181 == int32(1) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	if l0 != 0 {
		goto L67
	} else {
		goto L68
	}
L63:
	;
	goto L64
L64:
	;
	v248 = int32(0)
	if v64 == v248 {
		v465 = v248
		v467 = v66
		goto L33
	} else {
		goto L82
	}
L65:
	;
	v207 = v198
	goto L71
L66:
	;
	v197 = l0 + v186 - int32(1)
	v198 = int32(0)
	goto L65
L67:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0+v184)+87))
	if v186 != 0 {
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+88))
	v197 = v190
	v198 = int32(0)
	goto L65
L70:
	;
	v187 = int32(0)
	v197 = v187
	v198 = v187
	goto L65
L71:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v218+v207<<(uint(int32(2))%32))))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if l0 != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v246 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+8)) = uint8(v246)
	v465 = v64
	v467 = v66
	goto L33
L73:
	;
	v236 = F_FunctionCall2Coll(m, v223, v224, v197, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L79
	}
L74:
	;
	v235 = l0 + v227 - int32(1)
	goto L73
L75:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0+v225)+87))
	if v227 != 0 {
		goto L74
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+88))
	v235 = v230
	goto L73
L78:
	;
	v235 = int32(0)
	goto L73
L79:
	;
	v238 = *(*float32)(unsafe.Add(mBase, uint32(v84)+4))
	v239 = *(*float64)(unsafe.Add(mBase, uint32(v236)))
	if base.F32_ge(v238, base.F32_demote_f64(v239)) != 0 {
		goto L57
	} else {
		goto L80
	}
L80:
	;
	v243 = v207 + int32(1)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v243 < v244 {
		v207 = v243
		goto L71
	} else {
		goto L81
	}
L81:
	;
	goto L72
L82:
	;
	if l0 != 0 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	if v53 == int32(0) {
		goto L58
	} else {
		goto L89
	}
L84:
	;
	v261 = l0 + v253 - int32(1)
	goto L83
L85:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0+v251)+87))
	if v253 != 0 {
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+88))
	v261 = v256
	goto L83
L88:
	;
	v261 = int32(0)
	goto L83
L89:
	;
	v264 = int32(0)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v265 <= v264 {
		goto L58
	} else {
		goto L90
	}
L90:
	;
	v276 = v264
	goto L91
L91:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v287+v276<<(uint(int32(2))%32))))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if l0 != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	goto L58
L93:
	;
	v305 = F_FunctionCall2Coll(m, v292, v293, v261, v304)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L99
	}
L94:
	;
	v304 = l0 + v296 - int32(1)
	goto L93
L95:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0+v294)+87))
	if v296 != 0 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+88))
	v304 = v299
	goto L93
L98:
	;
	v304 = int32(0)
	goto L93
L99:
	;
	v307 = *(*float32)(unsafe.Add(mBase, uint32(v84)+4))
	v308 = *(*float64)(unsafe.Add(mBase, uint32(v305)))
	if base.F32_ge(v307, base.F32_demote_f64(v308)) != 0 {
		goto L57
	} else {
		goto L100
	}
L100:
	;
	v312 = v276 + int32(1)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v312 < v313 {
		v276 = v312
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
	if v53 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L104:
	;
	v326 = l0 + v318 - int32(1)
	goto L103
L105:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0+v316)+87))
	if v318 != 0 {
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+88))
	v326 = v321
	goto L103
L108:
	;
	v326 = int32(0)
	goto L103
L109:
	;
	v403 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+8)) = uint8(v403)
	v405 = F_lappend(m, v66, v84)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L125
	}
L110:
	;
	v329 = int32(0)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v330 <= v329 {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v341 = v329
	goto L112
L112:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v352+v341<<(uint(int32(2))%32))))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if l0 != 0 {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	v382 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+8)) = uint8(v382)
	v465 = v64
	v467 = v66
	goto L33
L114:
	;
	v370 = F_FunctionCall2Coll(m, v357, v358, v326, v369)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L120
	}
L115:
	;
	v369 = l0 + v361 - int32(1)
	goto L114
L116:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0+v359)+87))
	if v361 != 0 {
		goto L115
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)+88))
	v369 = v364
	goto L114
L119:
	;
	v369 = int32(0)
	goto L114
L120:
	;
	v372 = *(*float32)(unsafe.Add(mBase, uint32(v84)+4))
	v373 = *(*float64)(unsafe.Add(mBase, uint32(v370)))
	if base.F32_ge(v372, base.F32_demote_f64(v373)) == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v379 = v341 + int32(1)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v380 <= v379 {
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
	v341 = v379
	goto L112
L125:
	;
	v465 = v64
	v467 = v405
	goto L33
L126:
	;
	v465 = v426
	v467 = v429
	goto L33
L127:
	;
	if v85 != 0 {
		v53 = v483
		v63 = v85
		v64 = v465
		v66 = v467
		v67 = v484
		goto L25
	} else {
		goto L132
	}
L128:
	;
	v475 = F_lappend(m, v53, v84)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37+v67<<(uint(int32(2))%32)))) = v84
	v483 = v53
	v484 = v67 + int32(1)
	goto L127
L131:
	;
	v483 = v475
	v484 = v67
	goto L127
L132:
	;
	goto L26
L133:
	;
	v509 = int32(0)
	v510 = v487
	goto L134
L134:
	;
	if v510 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	if l6 == int32(0) {
		v591 = v510
		goto L3
	} else {
		goto L144
	}
L136:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v510)+4))
	v530 = v528
	goto L138
L137:
	;
	v530 = int32(0)
	goto L138
L138:
	;
	if v530 < l2 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v37+v509<<(uint(int32(2))%32))))
	v536 = F_lappend(m, v510, v535)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	goto L135
L142:
	;
	v539 = v509 + int32(1)
	if v539 == v501 {
		v547 = v536
		v557 = v497
		goto L20
	} else {
		goto L143
	}
L143:
	;
	v509 = v539
	v510 = v536
	goto L134
L144:
	;
	v569 = v510
	v587 = v37 + v509<<(uint(int32(2))%32)
	goto L19
L145:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v557)+12))
	v569 = v547
	v587 = v567
	goto L19
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
					F_errfinish(m, int32(_a_F_SetEpochTimestamp_7), int32(2176), int32(_a_F_SetEpochTimestamp_8))
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
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v121 int64
	_ = v121
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_logicalrep_worker_attach(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = int32(914)
		v13 = m.G0
		v15 = v13 - int32(32)
		m.G0 = v15
		switch int32(916) {
		case 0, 2:
			v25 = v11
		default:
			*(*int32)(unsafe.Add(mBase, _c_F_SetupApplyOrSyncWorker[0])) = v11
			v25 = int32(_a_F_SetupApplyOrSyncWorker_0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v25
		F_sigemptyset(m, v15+int32(16))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(268435456)
		v37 = v15 + int32(12)
		if v37 != 0 {
			v43 = int32(20)
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
			*(*int32)(unsafe.Add(mBase, _c_F_SetupApplyOrSyncWorker[1])) = v45
			v47 = *(*int64)(unsafe.Add(mBase, uint32(v37)+8))
			*(*int64)(unsafe.Add(mBase, _c_F_SetupApplyOrSyncWorker[2])) = v47
			v49 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
			*(*int64)(unsafe.Add(mBase, _c_F_SetupApplyOrSyncWorker[3])) = v49
		} else {
		}
		m.G0 = v15 + int32(32)
		v57 = int32(295)
		v59 = m.G0
		v61 = v59 - int32(32)
		m.G0 = v61
		switch int32(297) {
		case 0, 2:
			v71 = v57
		default:
			*(*int32)(unsafe.Add(mBase, _c_F_SetupApplyOrSyncWorker[4])) = v57
			v71 = int32(_a_F_SetupApplyOrSyncWorker_0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v61)+12)) = v71
		F_sigemptyset(m, v61+int32(16))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v61)+24)) = int32(268435456)
		v83 = v61 + int32(12)
		if v83 != 0 {
			v90 = int32(300)
			v91 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
			*(*int32)(unsafe.Add(mBase, _c_F_SetupApplyOrSyncWorker[5])) = v91
			v93 = *(*int64)(unsafe.Add(mBase, uint32(v83)+8))
			*(*int64)(unsafe.Add(mBase, _c_F_SetupApplyOrSyncWorker[6])) = v93
			v95 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
			*(*int64)(unsafe.Add(mBase, _c_F_SetupApplyOrSyncWorker[7])) = v95
		} else {
		}
		m.G0 = v61 + int32(32)
		F_BackgroundWorkerUnblockSignals(m)
		mBase = m.M
		v103 = m.ExcPending
		if v103 != 0 {
			return
		} else {
			v107 = m.G0
			v108 = int32(16)
			v109 = v107 - v108
			m.G0 = v109
			F_gettimeofday(m, v109)
			mBase = m.M
			v112 = *(*int64)(unsafe.Add(mBase, uint32(v109)))
			v113 = int64(*(*int32)(unsafe.Add(mBase, uint32(v109)+8)))
			m.G0 = v109 + v108
			v121 = v113 + v112*int64(1000000) - int64(946684800000000)
			v123 = *(*int32)(unsafe.Add(mBase, _c_F_SetupApplyOrSyncWorker[8]))
			*(*int64)(unsafe.Add(mBase, uint32(v123)+88)) = v121
			*(*int64)(unsafe.Add(mBase, uint32(v123)+104)) = v121
			*(*int64)(unsafe.Add(mBase, uint32(v123)+80)) = v121
			F_load_file(m, int32(_a_F_SetupApplyOrSyncWorker_1), int32(0))
			mBase = m.M
			v130 = m.ExcPending
			if v130 != 0 {
				return
			} else {
				F_InitializeLogRepWorker(m)
				mBase = m.M
				v132 = m.ExcPending
				if v132 != 0 {
					return
				} else {
					v135 = F_errstart(m, int32(14), int32(0))
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return
					} else {
						if v135 != 0 {
							v138 = *(*int32)(unsafe.Add(mBase, _c_F_SetupApplyOrSyncWorker[9]))
							v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v139
							F_errmsg_internal(m, int32(_a_F_SetupApplyOrSyncWorker_2), v6)
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_SetupApplyOrSyncWorker_3), int32(_a_F_SetupApplyOrSyncWorker_4), int32(_a_F_SetupApplyOrSyncWorker_5))
								mBase = m.M
								v148 = m.ExcPending
								if v148 != 0 {
									return
								} else {
									F_CacheRegisterSyscacheCallback(m, int32(68), int32(986), int32(0))
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return
									} else {
										m.G0 = v6 + int32(16)
										return
									}
								}
							}
						} else {
							F_CacheRegisterSyscacheCallback(m, int32(68), int32(986), int32(0))
							mBase = m.M
							v153 = m.ExcPending
							if v153 != 0 {
								return
							} else {
								m.G0 = v6 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_StartChildProcess(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_AssignPostmasterChildSlot(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			v16 = int32(0)
			v19 = F_errstart(m, int32(15), v16)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if l0 == int32(4) {
					if v19 == int32(0) {
						v79 = v16
						m.G0 = v8 + int32(16)
						return v79
					} else {
						F_errcode(m, int32(_a_F_StartChildProcess_0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_StartChildProcess_1), int32(0))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_StartChildProcess_2), int32(3971), int32(_a_F_StartChildProcess_3))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									v79 = v16
									m.G0 = v8 + int32(16)
									return v79
								}
							}
						}
					}
				} else {
					if v19 == int32(0) {
						v79 = v16
						m.G0 = v8 + int32(16)
						return v79
					} else {
						F_errmsg_internal(m, int32(_a_F_StartChildProcess_4), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_StartChildProcess_2), int32(3975), int32(_a_F_StartChildProcess_3))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v79 = v16
								m.G0 = v8 + int32(16)
								return v79
							}
						}
					}
				}
			}
		} else {
			v50 = F_postmaster_child_launch(m, l0, int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				if v50 < int32(0) {
					v54 = F_ReleasePostmasterChildSlot(m, v10)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v58 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							if v58 != 0 {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(12))+uint32(_c_F_StartChildProcess[0])))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v62
								F_errmsg(m, int32(_a_F_StartChildProcess_5), v8)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_StartChildProcess_2), int32(3986), int32(_a_F_StartChildProcess_3))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										if l0 != int32(13) {
											v79 = int32(0)
											m.G0 = v8 + int32(16)
											return v79
										} else {
											F_ExitPostmaster(m, int32(1))
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
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
									v79 = int32(0)
									m.G0 = v8 + int32(16)
									return v79
								} else {
									F_ExitPostmaster(m, int32(1))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
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
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v50
					v79 = v10
					m.G0 = v8 + int32(16)
					return v79
				}
			}
		}
	}
}
func F_StatementCancelHandler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StatementCancelHandler[0])))
	if v3 == int32(0) {
		v7 = int32(1)
		*(*int32)(unsafe.Add(mBase, _c_F_StatementCancelHandler[1])) = v7
		*(*int32)(unsafe.Add(mBase, _c_F_StatementCancelHandler[2])) = v7
	} else {
	}
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_StatementCancelHandler[3]))
	F_SetLatch(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
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
	Fn13991(m, l0, int32(_a_F_sanitize_char_2_0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_scalarineqsel_wrapper(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 float64
	_ = v55
	var v56 int32
	_ = v56
	var v65 float64
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 float64
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v26 = F_get_restriction_variable(m, v17, v18, v19, v12+int32(16), v12+int32(12), v12+int32(11))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int32(0)
	} else {
		if v26 == int32(0) {
			v76 = float64(0.3333333333333333)
			v77 = F_Float8GetDatum(m, v76)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				m.G0 = v12 + int32(48)
				return v77
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
			if v33 != int32(7) {
				v65 = float64(0.3333333333333333)
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
				if v66 == int32(0) {
					v76 = v65
					v77 = F_Float8GetDatum(m, v76)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						m.G0 = v12 + int32(48)
						return v77
					}
				} else {
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
					m.T0[v69].(func(*base.Module, int32))(m, v66)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						v76 = v65
						v77 = F_Float8GetDatum(m, v76)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							m.G0 = v12 + int32(48)
							return v77
						}
					}
				}
			} else {
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+24)))
				if v37 == int32(1) {
					v65 = float64(0)
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
					if v66 == int32(0) {
						v76 = v65
						v77 = F_Float8GetDatum(m, v76)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							m.G0 = v12 + int32(48)
							return v77
						}
					} else {
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
						m.T0[v69].(func(*base.Module, int32))(m, v66)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							v76 = v65
							v77 = F_Float8GetDatum(m, v76)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								m.G0 = v12 + int32(48)
								return v77
							}
						}
					}
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
					v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
					if v42 == int32(0) {
						v45 = F_get_commutator(m, v15)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							if v45 == int32(0) {
								v65 = float64(0.3333333333333333)
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
								if v66 == int32(0) {
									v76 = v65
									v77 = F_Float8GetDatum(m, v76)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										m.G0 = v12 + int32(48)
										return v77
									}
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
									m.T0[v69].(func(*base.Module, int32))(m, v66)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										v76 = v65
										v77 = F_Float8GetDatum(m, v76)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											m.G0 = v12 + int32(48)
											return v77
										}
									}
								}
							} else {
								v51 = l1 ^ int32(1)
								v52 = v45
								v55 = F_scalarineqsel(m, v17, v52, v51, l2, v14, v12+int32(16), v41, v40)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									v65 = v55
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
									if v66 == int32(0) {
										v76 = v65
										v77 = F_Float8GetDatum(m, v76)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											m.G0 = v12 + int32(48)
											return v77
										}
									} else {
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
										m.T0[v69].(func(*base.Module, int32))(m, v66)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											v76 = v65
											v77 = F_Float8GetDatum(m, v76)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												m.G0 = v12 + int32(48)
												return v77
											}
										}
									}
								}
							}
						}
					} else {
						v51 = l1
						v52 = v15
						v55 = F_scalarineqsel(m, v17, v52, v51, l2, v14, v12+int32(16), v41, v40)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							v65 = v55
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
							if v66 == int32(0) {
								v76 = v65
								v77 = F_Float8GetDatum(m, v76)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									m.G0 = v12 + int32(48)
									return v77
								}
							} else {
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
								m.T0[v69].(func(*base.Module, int32))(m, v66)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									v76 = v65
									v77 = F_Float8GetDatum(m, v76)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										m.G0 = v12 + int32(48)
										return v77
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
func F_scalarlesel(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_scalarineqsel_wrapper(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
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
						F_errfinish(m, int32(_a_F_scanner_yyerror_1), int32(1232), int32(_a_F_scanner_yyerror_2))
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
						F_errfinish(m, int32(_a_F_scanner_yyerror_1), int32(1240), int32(_a_F_scanner_yyerror_2))
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
	var v8 int32
	_ = v8
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v102 int32
	_ = v102
	var v105 int64
	_ = v105
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v176 int32
	_ = v176
	v8 = int32(0)
	v16 = l1 << (uint(int64(32)) % 64)
	if l0 == l3 {
		v176 = l3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v176
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v16
	return
L2:
	;
	v18 = F_superuser_arg(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	if v18 != 0 {
		v176 = l3
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v21 = int32(0)
	v23 = F_roles_is_member_of(m, l0, int32(1), v21, v21)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = l0
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = int64(0)
	if v23 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v176 = v53
	goto L1
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L3
	} else {
		goto L37
	}
L9:
	;
	return
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v30 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v42 = v8
	v46 = v8
	goto L12
L12:
	;
	if l2 == int32(0) {
		goto L8
	} else {
		goto L14
	}
L13:
	;
	goto L9
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v42<<(uint(int32(2))%32))))
	F_check_acl(m, l2)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	if v16 == int64(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v105 == v16 {
		goto L7
	} else {
		goto L32
	}
L17:
	;
	v105 = int64(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	if l3 == v53 {
		v176 = l3
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v60 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v70 = (v63<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L23
L22:
	;
	v70 = v60
	goto L23
L23:
	;
	v71 = int64(0)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v72 <= int32(0) {
		v105 = v71
		goto L16
	} else {
		goto L24
	}
L24:
	;
	v77 = int32(0)
	v78 = v71
	goto L25
L25:
	;
	v93 = v70 + l2 + v77<<(uint(int32(4))%32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v53 == v94 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v105 = v100
	goto L16
L27:
	;
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v93)+8))
	v98 = v96&v16 | v78
	if v98 == v16 {
		goto L7
	} else {
		goto L30
	}
L28:
	;
	v100 = v78
	goto L29
L29:
	;
	v102 = v77 + int32(1)
	if v102 != v72 {
		v77 = v102
		v78 = v100
		goto L25
	} else {
		goto L31
	}
L30:
	;
	v100 = v98
	goto L29
L31:
	;
	goto L26
L32:
	;
	if v105 == int64(0) {
		v127 = v46
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v129 = v42 + int32(1)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v129 < v130 {
		v42 = v129
		v46 = v127
		goto L12
	} else {
		goto L36
	}
L34:
	;
	v122 = base.I32_wrap_i64(base.I64_popcnt(v105))
	if v122 <= v46 {
		v127 = v46
		goto L33
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v53
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v105
	v127 = v122
	goto L33
L36:
	;
	goto L13
L37:
	;
	F_errmsg_internal(m, int32(_a_F_select_best_grantor_0), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_select_best_grantor_1), int32(1491), int32(_a_F_select_best_grantor_2))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
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
							F_errfinish(m, int32(_a_F_set_authn_id_1), int32(369), int32(_a_F_set_authn_id_2))
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
					F_errfinish(m, int32(_a_F_set_authn_id_1), int32(356), int32(_a_F_set_authn_id_2))
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
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
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
											if l0 == int32(3) {
												m.G0 = v8 + int32(80)
												return
											} else {
												F_SetConfigOption(m, int32(_a_F_set_debug_options_9), int32(_a_F_set_debug_options_6), l1, l2)
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return
												} else {
													if base.Ui32(l0) < base.Ui32(int32(5)) {
														m.G0 = v8 + int32(80)
														return
													} else {
														F_SetConfigOption(m, int32(_a_F_set_debug_options_10), int32(_a_F_set_debug_options_6), l1, l2)
														mBase = m.M
														v60 = m.ExcPending
														if v60 != 0 {
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
									if l0 == int32(3) {
										m.G0 = v8 + int32(80)
										return
									} else {
										F_SetConfigOption(m, int32(_a_F_set_debug_options_9), int32(_a_F_set_debug_options_6), l1, l2)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return
										} else {
											if base.Ui32(l0) < base.Ui32(int32(5)) {
												m.G0 = v8 + int32(80)
												return
											} else {
												F_SetConfigOption(m, int32(_a_F_set_debug_options_10), int32(_a_F_set_debug_options_6), l1, l2)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
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
		if v12 == int32(333) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v15
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
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
	switch v13 - int32(334) {
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
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v22 = v19
	goto L1
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
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
	switch v29 - int32(333) {
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
	if v103 != int32(333) {
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
	if v60 == int32(336) {
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
	switch v118 - int32(342) {
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
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
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v19 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
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
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v59
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
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)))
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
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v105
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = v128
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+364)) = v135
	goto L23
L32:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)))
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
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+360))
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
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
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
func F_setop_compare_slots(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
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
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v99 int32
	_ = v99
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v11 < v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_slot_getsomeattrs_int(m, l0, v10)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v19 < v18 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	F_slot_getsomeattrs_int(m, l1, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+120))
	if int32(0) < v23 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L8
L10:
	;
	return int32(1)
L11:
	;
	return v99
L12:
	;
	v29 = v23
	v31 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v99 = int32(0)
	goto L11
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+124))
	v37 = v34 + v31*int32(36)
	v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37)+10)))
	v39 = int32(1)
	v40 = v38 - v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v41))))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v40))))
	if v46 == v39 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L14
L17:
	;
	v85 = v31 + int32(1)
	if v85 < v83 {
		v29 = v83
		v31 = v85
		goto L15
	} else {
		goto L37
	}
L18:
	;
	if v43&int32(1) != 0 {
		v83 = v29
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v43&int32(1) != 0 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+9)))
	if v53 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v54 = int32(-1)
	goto L24
L23:
	;
	v54 = int32(1)
	goto L24
L24:
	;
	return v54
L25:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+9)))
	if v60 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v64 = v40 << (uint(int32(2)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64+v65)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v68+v64)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v72 = m.T0[v71].(func(*base.Module, int32, int32, int32) int32)(m, v67, v70, v37)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L31
	}
L28:
	;
	v61 = int32(1)
	goto L30
L29:
	;
	v61 = int32(-1)
	goto L30
L30:
	;
	return v61
L31:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+8)))
	if v74 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v72 < int32(0) {
		goto L10
	} else {
		goto L35
	}
L33:
	;
	v81 = v72
	goto L34
L34:
	;
	if v81 != 0 {
		v99 = v81
		goto L11
	} else {
		goto L36
	}
L35:
	;
	v81 = int32(0) - v72
	goto L34
L36:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+120))
	v83 = v82
	goto L17
L37:
	;
	goto L16
}
func F_shimBoolConsistentFn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v11 = F_FunctionCall7Coll(m, v2, v3, v4, v5, v6, v7, v8, v9, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v16 = v11 & int32(255)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+87)) = uint8(base.B2i32(v16 == int32(2)))
		return base.B2i32(v16 != int32(0))
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
func F_show_limit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	v3 = *(*float64)(unsafe.Add(mBase, _c_F_show_limit[0]))
	return base.I32_reinterpret_f32(base.F32_demote_f64(v3))
}
func F_show_unix_socket_permissions(m *base.Module) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13992(m, int32(_a_F_show_unix_socket_permissions_0), int32(_a_F_show_unix_socket_permissions_1))
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
func F_similar_escape(m *base.Module, l0 int32) int32 {
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
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v2 == int32(1) {
		v5 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v5)
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v10 = F_pg_detoast_datum_packed(m, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v14 != 0 {
				v19 = int32(0)
				v20 = F_similar_escape_internal(m, v10, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					return v20
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v17 = F_pg_detoast_datum_packed(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = v17
					v20 = F_similar_escape_internal(m, v10, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						return v20
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
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
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
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
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
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v371 int32
	_ = v371
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	v3 = int32(0)
	v16 = int32(1)
	v17 = l0 + v16
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v20 = v18 & v16
	if v18 == v16 {
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
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v26 == int32(18) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v37 = int32(1)
	if v20 != 0 {
		v47 = int32(base.Ui32(v18)>>(uint(v37)%32)) - v37
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v29 = int32(16)
	goto L7
L6:
	;
	v29 = int32(0)
	goto L7
L7:
	;
	if base.Ui32((v26-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v36 = int32(4)
	goto L10
L9:
	;
	v36 = v29
	goto L10
L10:
	;
	v47 = v36
	goto L1
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L34
	} else {
		goto L115
	}
L13:
	;
	v110 = F_palloc(m, v47*int32(3)+int32(27))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L34
	} else {
		goto L37
	}
L14:
	;
	v104 = int32(1)
	v105 = int32(_a_F_similar_escape_internal_0)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v52 = int32(4)
	v53 = int32(1)
	v54 = l1 + v53
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v59 = v57 & v53
	if v59 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v60 = v54
	goto L19
L18:
	;
	v60 = l1 + v52
	goto L19
L19:
	;
	if v57 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v94 = F_pg_mbstrlen_with_len(m, v60, v93)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L34
	} else {
		goto L35
	}
L21:
	;
	if v85 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if base.Ui32((v63-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v93 = v52
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v75 = int32(1)
	if v59 != 0 {
		v85 = int32(base.Ui32(v57)>>(uint(v75)%32)) - v75
		goto L21
	} else {
		goto L29
	}
L25:
	;
	if v63 == int32(18) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v74 = int32(16)
	goto L28
L27:
	;
	v74 = int32(0)
	goto L28
L28:
	;
	v85 = v74
	goto L21
L29:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v85 = int32(base.Ui32(v79)>>(uint(int32(2))%32)) - int32(4)
	goto L21
L30:
	;
	v88 = int32(0)
	v104 = v88
	v105 = v88
	goto L13
L31:
	;
	goto L32
L32:
	;
	if v85 < int32(2) {
		v104 = v85
		v105 = v60
		goto L13
	} else {
		goto L33
	}
L33:
	;
	v93 = v85
	goto L20
L34:
	;
	return int32(0)
L35:
	;
	if int32(2) <= v94 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	v104 = v93
	v105 = v60
	goto L13
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+4)) = int32(977217630)
	v115 = v110 + int32(8)
	if v47 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v371 = int32(_a_F_similar_escape_internal_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v356))) = uint16(v371)
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = (v356-v110)<<(uint(int32(2))%32) + int32(8)
	return v110
L39:
	;
	v356 = v115
	goto L38
L40:
	;
	goto L41
L41:
	;
	if v20 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v120 = v17
	goto L44
L43:
	;
	v120 = l0 + int32(4)
	goto L44
L44:
	;
	v124 = int32(0)
	v127 = v120
	v128 = v115
	v131 = v124
	v132 = v124
	v133 = v47
	v136 = v3
	v138 = v3
	goto L45
L45:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v104 < int32(2) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	v356 = v345
	goto L38
L47:
	;
	if int32(0) < v350 {
		v127 = v353
		v128 = v345
		v131 = v348
		v132 = v349
		v133 = v350
		v136 = v351
		v138 = v352
		goto L45
	} else {
		goto L114
	}
L48:
	;
	v345 = v342
	v348 = v341
	v349 = v132
	v350 = v133 - v142
	v351 = v136
	v352 = v138
	v353 = v127 + v142
	goto L47
L49:
	;
	if v142 != 0 {
		goto L111
	} else {
		goto L112
	}
L50:
	;
	if v131 != 0 {
		goto L82
	} else {
		goto L83
	}
L51:
	;
	v142 = F_pg_mblen_range(m, v127, v120+v47)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L34
	} else {
		goto L52
	}
L52:
	;
	if v142 < int32(2) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	if v131 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v146 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v146)
	v336 = v128 + int32(1)
	goto L49
L55:
	;
	goto L56
L56:
	;
	if base.B2i32(v105 == int32(0))|base.B2i32(v142 != v104) != 0 {
		v336 = v128
		goto L49
	} else {
		goto L57
	}
L57:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v104) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	if v215 != 0 {
		v336 = v128
		goto L49
	} else {
		goto L76
	}
L59:
	;
	v215 = int32(0)
	goto L58
L60:
	;
	v189 = v184
	v190 = v185
	v191 = v186
	goto L70
L61:
	;
	if (v105|v127)&int32(3) != 0 {
		v184 = v105
		v185 = v127
		v186 = v104
		goto L60
	} else {
		goto L64
	}
L62:
	;
	v177 = v105
	v178 = v127
	v179 = v104
	goto L63
L63:
	;
	if v179 == int32(0) {
		goto L59
	} else {
		goto L69
	}
L64:
	;
	v161 = v105
	v162 = v127
	v163 = v104
	goto L65
L65:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	if v166 != v167 {
		v184 = v161
		v185 = v162
		v186 = v163
		goto L60
	} else {
		goto L67
	}
L66:
	;
	v177 = v172
	v178 = v170
	v179 = v174
	goto L63
L67:
	;
	v169 = int32(4)
	v170 = v162 + v169
	v172 = v161 + v169
	v174 = v163 - v169
	if base.Ui32(int32(3)) < base.Ui32(v174) {
		v161 = v172
		v162 = v170
		v163 = v174
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v184 = v177
	v185 = v178
	v186 = v179
	goto L60
L70:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	if v194 == v195 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v215 = v194 - v195
	goto L58
L72:
	;
	v197 = int32(1)
	v202 = v191 - v197
	if v202 != 0 {
		v189 = v189 + v197
		v190 = v190 + v197
		v191 = v202
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
	v341 = int32(1)
	v342 = v128
	goto L48
L77:
	;
	v332 = int32(1)
	v345 = v327
	v348 = v329
	v349 = v328
	v350 = v133 - v332
	v351 = v330
	v352 = v331
	v353 = v127 + v332
	goto L47
L78:
	;
	v327 = v322
	v328 = v132
	v329 = int32(0)
	v330 = v324
	v331 = v325
	goto L77
L79:
	;
	v296 = v128 + int32(1)
	switch v141 - int32(36) {
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
	v276 = v128 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v141)
	if base.B2i32(v141 != int32(93))|base.B2i32(v132 < int32(3)) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L81:
	;
	v327 = v128 + int32(2)
	v328 = int32(3)
	v329 = int32(0)
	v330 = v136
	v331 = v138
	goto L77
L82:
	;
	v220 = int32(0)
	if base.B2i32(v141 != int32(34))|base.B2i32(v220 < v136) == v220 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	if v105 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L85:
	;
	switch v138 {
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
	*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)) = uint8(v141)
	v256 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v256)
	goto L81
L88:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v128))) = int64(2900174335198198569)
	v322 = v128 + int32(8)
	v324 = v136
	v325 = v138 + int32(1)
	goto L78
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L34
	} else {
		goto L91
	}
L90:
	;
	v225 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v128)+8)) = uint8(v225)
	*(*int64)(unsafe.Add(mBase, uint32(v128))) = int64(4551025073606196009)
	v322 = v128 + int32(9)
	v324 = v136
	v325 = v138 + int32(1)
	goto L78
L91:
	;
	F_errcode(m, int32(318767234))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L34
	} else {
		goto L92
	}
L92:
	;
	F_errmsg(m, int32(_a_F_similar_escape_internal_2), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L34
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_similar_escape_internal_3), int32(951), int32(_a_F_similar_escape_internal_4))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
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
	if v136 <= int32(0) {
		goto L79
	} else {
		goto L98
	}
L96:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	if v141 != v261 {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v327 = v128
	v328 = v132
	v329 = int32(1)
	v330 = v136
	v331 = v138
	goto L77
L98:
	;
	if v141 != int32(92) {
		goto L80
	} else {
		goto L99
	}
L99:
	;
	v268 = int32(_a_F_similar_escape_internal_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v128))) = uint16(v268)
	goto L81
L100:
	;
	v322 = v276
	v324 = v136 - int32(1)
	v325 = v138
	goto L78
L101:
	;
	goto L102
L102:
	;
	v287 = int32(3)
	v288 = int32(0)
	switch v141 - int32(91) {
	case 0:
		goto L104
	default:
		v327 = v276
		v328 = v287
		v329 = v288
		v330 = v136
		v331 = v138
		goto L77
	case 3:
		goto L103
	}
L103:
	;
	v327 = v276
	v328 = v132 + int32(1)
	v329 = v288
	v330 = v136
	v331 = v138
	goto L77
L104:
	;
	v327 = v276
	v328 = v287
	v329 = v288
	v330 = v136 + int32(1)
	v331 = v138
	goto L77
L105:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v141)
	v322 = v296
	v324 = v136
	v325 = v138
	goto L78
L106:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)) = uint8(v141)
	v317 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v317)
	v322 = v128 + int32(2)
	v324 = v136
	v325 = v138
	goto L78
L107:
	;
	v310 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v128)+2)) = uint8(v310)
	v312 = int32(_a_F_similar_escape_internal_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v128))) = uint16(v312)
	v322 = v128 + int32(3)
	v324 = v136
	v325 = v138
	goto L78
L108:
	;
	v308 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v308)
	v322 = v296
	v324 = v136
	v325 = v138
	goto L78
L109:
	;
	v304 = int32(_a_F_similar_escape_internal_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v128))) = uint16(v304)
	v322 = v128 + int32(2)
	v324 = v136
	v325 = v138
	goto L78
L110:
	;
	v299 = int32(91)
	*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v299)
	v301 = int32(1)
	v327 = v296
	v328 = v301
	v329 = int32(0)
	v330 = v301
	v331 = v138
	goto L77
L111:
	;
	base.MemoryCopy(m, v336, v127, v142)
	goto L113
L112:
	;
	goto L113
L113:
	;
	v341 = int32(0)
	v342 = v336 + v142
	goto L48
L114:
	;
	goto L46
L115:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L34
	} else {
		goto L116
	}
L116:
	;
	F_errmsg(m, int32(_a_F_similar_escape_internal_8), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L34
	} else {
		goto L117
	}
L117:
	;
	F_errhint(m, int32(_a_F_similar_escape_internal_9), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L34
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_similar_escape_internal_3), int32(805), int32(_a_F_similar_escape_internal_4))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
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
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v91 int32
	_ = v91
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v7 != int32(1) {
		v91 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v91
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+144))
	if v10 != 0 {
		v91 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)))
	if v11 != 0 {
		v91 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v12 != 0 {
		v91 = v3
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+37)))
	if v13 != 0 {
		v91 = v3
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	if v14 != 0 {
		v91 = v3
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+42)))
	if v15 != 0 {
		v91 = v3
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	if v16 != 0 {
		v91 = v3
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	if v17 != 0 {
		v91 = v3
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	if v18 != 0 {
		v91 = v3
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
	v38 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+124)) = v38
	*(*int64)(unsafe.Add(mBase, uint32(l1)+116)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+100)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v38
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+40)) = uint8(v38)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v48 != 0 {
		goto L22
	} else {
		goto L23
	}
L15:
	;
	return int32(0)
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+132)) = v20
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v25 != int32(7) {
		v91 = v3
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+24)))
	if v28 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
	if v32 <= int64(0) {
		v91 = v3
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
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v50 <= int32(0) {
		v91 = int32(1)
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v91 = int32(1)
	goto L1
L25:
	;
	v53 = int32(0)
	if v53 < v50 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v56 = v50
	goto L28
L27:
	;
	v56 = v53
	goto L28
L28:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v59 = int32(0)
	goto L29
L29:
	;
	v67 = v57 + v59<<(uint(int32(2))%32)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	if v69 != int32(9) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v76 = F_list_delete_cell(m, v48, v67)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L15
	} else {
		goto L35
	}
L31:
	;
	v72 = int32(1)
	v74 = v59 + v72
	if v56 != v74 {
		v59 = v74
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
	v91 = v72
	goto L1
L35:
	;
	v78 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v78)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v76
	goto L24
}
func F_sjis_to_mic(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v59 int32
	_ = v59
	var v73 int32
	_ = v73
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v15, v16, v17, int32(35), int32(7))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v17 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v334 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v324))) = uint8(v334)
	return v329 - v14
L4:
	;
	v324 = v13
	v329 = v14
	goto L3
L5:
	;
	goto L6
L6:
	;
	v27 = v13
	v32 = v14
	v33 = v17
	goto L7
L7:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if base.Ui32((v37+int32(95))&int32(255)) <= base.Ui32(int32(62)) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v324 = v309
	v329 = v319
	goto L3
L9:
	;
	v319 = v312 + v32
	v320 = v313 + v33
	if int32(0) < v320 {
		v27 = v309
		v32 = v319
		v33 = v320
		goto L7
	} else {
		goto L72
	}
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)) = uint8(v37)
	v45 = int32(137)
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v45)
	v309 = v27 + int32(2)
	v312 = int32(1)
	v313 = int32(-1)
	goto L9
L11:
	;
	goto L12
L12:
	;
	if base.I32_extend8_s(v37) < int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v309 = v27 + int32(3)
	v312 = v304
	v313 = int32(-2)
	goto L9
L14:
	;
	v59 = int32(0)
	if base.B2i32(base.B2i32(v37 != int32(128))&base.B2i32(base.Ui32(v37) < base.Ui32(int32(160))) == v59)&base.B2i32(base.Ui32(int32(28)) < base.Ui32((v37+int32(32))&int32(255)))|base.B2i32(v33 == int32(1)) == v59 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	if v37 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L17:
	;
	v89 = v73 & int32(255)
	v92 = v89 | v37<<(uint(int32(8))%32)
	if base.Ui32(v92-int32(_a_F_sjis_to_mic_0)) <= base.Ui32(int32(767)) {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v73 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32)+1)))
	if base.B2i32(v73 < int32(-3))|base.B2i32(base.Ui32((v73+int32(-64))&int32(255)) < base.Ui32(int32(63))) != 0 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v12 != 0 {
		v324 = v27
		v329 = v32
		goto L3
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	F_report_invalid_encoding(m, int32(35), v32, v33)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v97 = v92
	v99 = v37
	v100 = v89
	v102 = int32(0)
	goto L27
L25:
	;
	v138 = v92
	v140 = v37
	v141 = v89
	goto L26
L26:
	;
	if v138 <= int32(_a_F_sjis_to_mic_1) {
		goto L36
	} else {
		goto L37
	}
L27:
	;
	v109 = v102 << (uint(int32(3)) % 32)
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_sjis_to_mic[0]))))
	if v112 == v97 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v138 = v131
	v140 = v132
	v141 = v133
	goto L26
L29:
	;
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_sjis_to_mic[1]))))
	v119 = v114
	v120 = int32(base.Ui32(v114) >> (uint(int32(8)) % 32))
	v121 = v114 & int32(255)
	goto L31
L30:
	;
	v119 = v97
	v120 = v99
	v121 = v100
	goto L31
L31:
	;
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_sjis_to_mic[2]))))
	if v124 == v119 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_sjis_to_mic[3]))))
	v131 = v126
	v132 = int32(base.Ui32(v126) >> (uint(int32(8)) % 32))
	v133 = v126 & int32(255)
	goto L34
L33:
	;
	v131 = v119
	v132 = v120
	v133 = v121
	goto L34
L34:
	;
	v135 = v102 + int32(2)
	if v135 != int32(388) {
		v97 = v131
		v99 = v132
		v100 = v133
		v102 = v135
		goto L27
	} else {
		goto L35
	}
L35:
	;
	goto L28
L36:
	;
	v151 = int32(146)
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v151)
	v153 = int32(2)
	v157 = base.B2i32(base.Ui32(int32(158)) < base.Ui32(v141))
	if base.Ui32(int32(158)) < base.Ui32(v141) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v177 = int32(0)
	if base.B2i32(base.B2i32(v138 != int32(_a_F_sjis_to_mic_2))&base.B2i32(base.Ui32(v138) < base.Ui32(int32(_a_F_sjis_to_mic_3))) == v177)&base.B2i32(base.Ui32(int32(176)) < base.Ui32(v138-int32(_a_F_sjis_to_mic_4))) == v177 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v158 = v153
	goto L41
L40:
	;
	v158 = int32(96)
	goto L41
L41:
	;
	v162 = v158 + v141 + base.B2i32(base.Ui32(v141) < base.Ui32(int32(128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+2)) = uint8(v162)
	v170 = v140<<(uint(int32(1))%32)&int32(126) | v157 + int32(159)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)) = uint8(v170)
	v304 = v153
	goto L13
L42:
	;
	v186 = int32(174)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+2)) = uint8(v186)
	v188 = int32(_a_F_sjis_to_mic_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v188)
	v309 = v27 + int32(3)
	v312 = int32(2)
	v313 = int32(-2)
	goto L9
L43:
	;
	goto L44
L44:
	;
	if base.Ui32(v138-int32(_a_F_sjis_to_mic_3)) <= base.Ui32(int32(1279)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v198 = int32(146)
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v198)
	v200 = int32(2)
	v204 = base.B2i32(base.Ui32(int32(158)) < base.Ui32(v141))
	if base.Ui32(int32(158)) < base.Ui32(v141) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	if base.Ui32(v138-int32(_a_F_sjis_to_mic_6)) <= base.Ui32(int32(1279)) {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	v205 = v200
	goto L50
L49:
	;
	v205 = int32(96)
	goto L50
L50:
	;
	v209 = v205 + v141 + base.B2i32(base.Ui32(v141) < base.Ui32(int32(128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+2)) = uint8(v209)
	v219 = (v140<<(uint(int32(1))%32)+int32(34))&int32(126) | v204 + int32(243)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)) = uint8(v219)
	v304 = v200
	goto L13
L51:
	;
	v225 = int32(148)
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v225)
	v227 = int32(2)
	v231 = base.B2i32(base.Ui32(int32(158)) < base.Ui32(v141))
	if base.Ui32(int32(158)) < base.Ui32(v141) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v248 = int32(-2)
	v249 = int32(2)
	if base.Ui32(v138) < base.Ui32(int32(_a_F_sjis_to_mic_7)) {
		v309 = v27
		v312 = v249
		v313 = v248
		goto L9
	} else {
		goto L57
	}
L54:
	;
	v232 = v227
	goto L56
L55:
	;
	v232 = int32(96)
	goto L56
L56:
	;
	v236 = v232 + v141 + base.B2i32(base.Ui32(v141) < base.Ui32(int32(128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+2)) = uint8(v236)
	v246 = (v140<<(uint(int32(1))%32)+int32(24))&int32(126) | v231 + int32(243)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)) = uint8(v246)
	v304 = v227
	goto L13
L57:
	;
	v253 = v138
	v254 = v27
	v255 = int32(0)
	goto L58
L58:
	;
	v265 = v255 << (uint(int32(3)) % 32)
	v268 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v265)+uint32(_c_F_sjis_to_mic[1]))))
	if v268 == v253 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v309 = v288
	v312 = v249
	v313 = v248
	goto L9
L60:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v265)+uint32(_c_F_sjis_to_mic[4])))
	v271 = int32(128)
	v272 = v270 | v271
	*(*uint8)(unsafe.Add(mBase, uint32(v254)+2)) = uint8(v272)
	v277 = int32(base.Ui32(v270)>>(uint(int32(8))%32)) | v271
	*(*uint8)(unsafe.Add(mBase, uint32(v254)+1)) = uint8(v277)
	if int32(_a_F_sjis_to_mic_8) < v270 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v287 = v253
	v288 = v254
	goto L62
L62:
	;
	v290 = v255 + int32(1)
	if v290 != int32(388) {
		v253 = v287
		v254 = v288
		v255 = v290
		goto L58
	} else {
		goto L66
	}
L63:
	;
	v283 = int32(-108)
	goto L65
L64:
	;
	v283 = int32(-110)
	goto L65
L65:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v254))) = uint8(v283)
	v287 = v270
	v288 = v254 + int32(3)
	goto L62
L66:
	;
	goto L59
L67:
	;
	if v12 != 0 {
		v324 = v27
		v329 = v32
		goto L3
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v37)
	v299 = int32(1)
	v309 = v27 + v299
	v312 = v299
	v313 = int32(-1)
	goto L9
L70:
	;
	F_report_invalid_encoding(m, int32(35), v32, v33)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	goto L8
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
func F_slotsync_worker_disconnect(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_disconnect[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+64))
	m.T0[v5].(func(*base.Module, int32))(m, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
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
	var v6 int32
	_ = v6
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
	var v47 int32
	_ = v47
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int64
	_ = v89
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v158 int32
	_ = v158
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int64
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v286 int32
	_ = v286
	var v304 int32
	_ = v304
	var v315 int32
	_ = v315
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v400 int32
	_ = v400
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v450 int32
	_ = v450
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v563 int64
	_ = v563
	var v565 int64
	_ = v565
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v671 int64
	_ = v671
	var v674 int64
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	v6 = int32(0)
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
	v563 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v563
	v565 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v565
	F_CacheInvalidateSmgr(m, v25)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L22
	} else {
		goto L83
	}
L2:
	;
	if v286 <= int32(0) {
		goto L1
	} else {
		goto L44
	}
L3:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[0]))
	if base.Ui32(int32(63)) <= base.Ui32(v275+int32(31)) {
		goto L1
	} else {
		goto L43
	}
L4:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[0]))
	v286 = v273
	goto L2
L5:
	;
	if l2 <= int32(0) {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[1]))
	if base.B2i32(v35 != v146)|base.B2i32(l2 <= int32(0)) != 0 {
		goto L1
	} else {
		goto L25
	}
L8:
	;
	v47 = v6
	v60 = int64(0)
	goto L9
L9:
	;
	v63 = v47 << (uint(int32(2)) % 32)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1+v63)))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrtruncate[2])))
	if v70 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[0]))
	v100 = base.I32_div_s(v98, int32(32))
	if base.B2i32(base.I32_wrap_i64(v89) == int32(-1))|base.B2i32(base.Ui64(base.I64_extend_i32_s(v100)) <= base.Ui64(v89)) != 0 {
		v286 = v98
		goto L2
	} else {
		goto L19
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63+(v29+int32(28))))) = v81
	if v81 == int32(-1) {
		goto L4
	} else {
		goto L17
	}
L12:
	;
	goto L11
L13:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0+v68<<(uint(int32(2))%32))+20))
	if v76 != int32(-1) {
		v81 = v76
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v81 = int32(-1)
	goto L12
L16:
	;
	goto L15
L17:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l4+v63)))
	v89 = v60 + base.I64_extend_i32_u(v81-v86)
	v91 = v47 + int32(1)
	if v91 != l2 {
		v47 = v91
		v60 = v89
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	v111 = int32(0)
	goto L20
L20:
	;
	v127 = v111 << (uint(int32(2)) % 32)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l4+v127)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1+v127)))
	v132 = *(*int64)(unsafe.Add(mBase, uint32(v29)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v132
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v134
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(28)+v127)))
	F_FindAndDropRelationBuffers(m, v29, v131, v139, v129)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
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
	v143 = v111 + int32(1)
	if v143 != l2 {
		v111 = v143
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v158 = v6
	goto L26
L26:
	;
	v174 = v158 << (uint(int32(2)) % 32)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l4+v174)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l1+v174)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v179
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v29)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v181
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[3]))
	if int32(0) < v184 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L1
L28:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[4]))
	v190 = v29 + int32(16)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+8))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v203 = int32(0)
	v204 = v184
	v206 = v188
	goto L31
L29:
	;
	goto L30
L30:
	;
	v270 = v158 + int32(1)
	if v270 != l2 {
		v158 = v270
		goto L26
	} else {
		goto L42
	}
L31:
	;
	v219 = v206 + v203<<(uint(int32(6))%32)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
	if v220&int32(33554432) == int32(0) {
		v242 = v204
		v243 = v206
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L30
L33:
	;
	v245 = v203 + int32(1)
	if v245 < v242 {
		v203 = v245
		v204 = v242
		v206 = v243
		goto L31
	} else {
		goto L41
	}
L34:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	if v225 != v193 {
		v242 = v204
		v243 = v206
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if v227 != v192 {
		v242 = v204
		v243 = v206
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v219)+8))
	if v229 != v191 {
		v242 = v204
		v243 = v206
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
	if v231 != v178 {
		v242 = v204
		v243 = v206
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	if base.Ui32(v233) < base.Ui32(v176) {
		v242 = v204
		v243 = v206
		goto L33
	} else {
		goto L39
	}
L39:
	;
	F_InvalidateLocalBuffer(m, v219, int32(1))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L22
	} else {
		goto L40
	}
L40:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[3]))
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[4]))
	v242 = v239
	v243 = v241
	goto L33
L41:
	;
	goto L32
L42:
	;
	goto L27
L43:
	;
	v286 = v275
	goto L2
L44:
	;
	v304 = int32(0)
	v315 = v304
	goto L45
L45:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[5]))
	v333 = v330 + v315<<(uint(int32(6))%32)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	if v334 != v335 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L1
L47:
	;
	v534 = v315 + int32(1)
	v536 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[0]))
	if v534 < v536 {
		v315 = v534
		goto L45
	} else {
		goto L82
	}
L48:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	if v337 != v338 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v333)+8))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	if v340 != v341 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+76)) = int32(_a_F_smgrtruncate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = int32(_a_F_smgrtruncate_1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = int32(_a_F_smgrtruncate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+56)) = int64(0)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v333)+24))
	v354 = int32(_a_F_smgrtruncate_3)
	*(*int32)(unsafe.Add(mBase, uint32(v333)+24)) = v353 | v354
	if v353&v354 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	goto L54
L52:
	;
	v400 = v353
	goto L53
L53:
	;
	v416 = int32(_a_F_smgrtruncate_4)
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[6]))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(56))+8))
	if v419 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L54:
	;
	F_perform_spin_delay(m, v29+int32(56))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L22
	} else {
		goto L56
	}
L55:
	;
	v400 = v385
	goto L53
L56:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v333)+24))
	v386 = int32(_a_F_smgrtruncate_3)
	*(*int32)(unsafe.Add(mBase, uint32(v333)+24)) = v385 | v386
	if v385&v386 != 0 {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	if base.B2i32(l2 <= v304) == int32(0) {
		goto L69
	} else {
		goto L70
	}
L59:
	;
	goto L58
L60:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[6])) = v434
	goto L59
L61:
	;
	if int32(999) < v417 {
		goto L59
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v417 < int32(11) {
		goto L59
	} else {
		goto L68
	}
L64:
	;
	v424 = int32(900)
	if v424 <= v417 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v427 = v424
	goto L67
L66:
	;
	v427 = v417
	goto L67
L67:
	;
	v434 = v427 + int32(100)
	goto L60
L68:
	;
	v434 = v417 - int32(1)
	goto L60
L69:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	v450 = int32(0)
	goto L72
L70:
	;
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333)+24)) = v400 & int32(-4194305)
	goto L47
L72:
	;
	if v441 != v442 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L71
L74:
	;
	v484 = v450 + int32(1)
	if v484 != l2 {
		v450 = v484
		goto L72
	} else {
		goto L81
	}
L75:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	if v466 != v440 {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v333)+8))
	if v468 != v439 {
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v333)+12))
	v472 = v450 << (uint(int32(2)) % 32)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l1+v472)))
	if v470 != v474 {
		goto L74
	} else {
		goto L78
	}
L78:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v333)+16))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l4+v472)))
	if base.Ui32(v476) < base.Ui32(v478) {
		goto L74
	} else {
		goto L79
	}
L79:
	;
	F_InvalidateBuffer(m, v333)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L22
	} else {
		goto L80
	}
L80:
	;
	goto L47
L81:
	;
	goto L73
L82:
	;
	goto L46
L83:
	;
	if int32(0) < l2 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v572 = l0 + int32(20)
	v590 = v6
	goto L87
L85:
	;
	goto L86
L86:
	;
	m.G0 = v25 + int32(16)
	return
L87:
	;
	v595 = int32(2)
	v596 = v590 << (uint(v595) % 32)
	v597 = l1 + v596
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v597)))
	*(*int32)(unsafe.Add(mBase, uint32(v572+v598<<(uint(v595)%32)))) = int32(-1)
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v597)))
	v605 = l3 + v596
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v605)))
	v607 = l4 + v596
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v607)))
	v609 = m.G0
	v611 = v609 - int32(112)
	m.G0 = v611
	if base.Ui32(v606) < base.Ui32(v608) {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	goto L86
L89:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v597)))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v607)))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v605)))
	if base.Ui32(v820) < base.Ui32(v821) {
		goto L145
	} else {
		goto L146
	}
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L22
	} else {
		goto L140
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L22
	} else {
		goto L135
	}
L92:
	;
	m.G0 = v611 + int32(112)
	goto L89
L93:
	;
	v615 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrtruncate[2])))
	if v615 != 0 {
		goto L92
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if v606 == v608 {
		goto L92
	} else {
		goto L101
	}
L96:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L22
	} else {
		goto L97
	}
L97:
	;
	v621 = v611 + int32(40)
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_GetRelationPath(m, v621, v622, v623, v624, v625, v604)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L22
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v611)+8)) = v606
	*(*int32)(unsafe.Add(mBase, uint32(v611)+4)) = v608
	*(*int32)(unsafe.Add(mBase, uint32(v611))) = v621
	F_errmsg(m, int32(_a_F_smgrtruncate_5), v611)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L22
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_smgrtruncate_6), int32(1305), int32(_a_F_smgrtruncate_7))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L22
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	v642 = l0 + v604<<(uint(int32(2))%32)
	v644 = v642 + int32(40)
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v644)))
	if v645 <= int32(0) {
		goto L92
	} else {
		goto L102
	}
L102:
	;
	v649 = v642 + int32(56)
	v671 = base.I64_extend_i32_u(v645)
	goto L103
L103:
	;
	v674 = v671 - int64(1)
	v675 = base.I32_wrap_i64(v674)
	v677 = v675 << (uint(int32(3)) % 32)
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v649)))
	v679 = v677 + v678
	v681 = v675 << (uint(int32(17)) % 32)
	if base.Ui32(v608) < base.Ui32(v681) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L92
L105:
	;
	if base.Ui64(int64(1)) < base.Ui64(v671) {
		v671 = v674
		goto L103
	} else {
		goto L134
	}
L106:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v679)))
	v686 = F_FileTruncate(m, v683, int64(0), int32(167772183))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L22
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	if base.Ui32(v681+int32(_a_F_smgrtruncate_8)) <= base.Ui32(v608) {
		goto L92
	} else {
		goto L129
	}
L109:
	;
	if v686 < int32(0) {
		goto L91
	} else {
		goto L110
	}
L110:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v690 == int32(-1) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	F_register_dirty_segment(m, l0, v604, v679)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L22
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v679)))
	F_FileClose(m, v695)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L22
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v644)))
	if v674 == int64(0) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v644))) = v675
	goto L105
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v649))) = v718
	goto L116
L118:
	;
	if v698 <= int32(0) {
		goto L116
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	if v698 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v649)))
	F_pfree(m, v703)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L22
	} else {
		goto L122
	}
L122:
	;
	v718 = int32(0)
	goto L117
L123:
	;
	v710 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[7]))
	v711 = F_MemoryContextAlloc(m, v710, v677)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L22
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	if v674 <= base.I64_extend_i32_s(v698) {
		goto L116
	} else {
		goto L127
	}
L126:
	;
	v718 = v711
	goto L117
L127:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v649)))
	v716 = F_repalloc(m, v715, v677)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L22
	} else {
		goto L128
	}
L128:
	;
	v718 = v716
	goto L117
L129:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v679)))
	v730 = F_FileTruncate(m, v724, base.I64_extend_i32_u(v608-v681)<<(uint(int64(13))%64), int32(167772183))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L22
	} else {
		goto L130
	}
L130:
	;
	if v730 < int32(0) {
		goto L90
	} else {
		goto L131
	}
L131:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v734 != int32(-1) {
		goto L105
	} else {
		goto L132
	}
L132:
	;
	F_register_dirty_segment(m, l0, v604, v679)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L22
	} else {
		goto L133
	}
L133:
	;
	goto L105
L134:
	;
	goto L104
L135:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L22
	} else {
		goto L136
	}
L136:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v679)))
	v775 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[8]))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v775+v773*int32(48))+32))
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v611)+16)) = v779
	F_errmsg(m, int32(_a_F_smgrtruncate_9), v611+int32(16))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L22
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_smgrtruncate_6), int32(1333), int32(_a_F_smgrtruncate_7))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L22
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L22
	} else {
		goto L141
	}
L141:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v679)))
	v799 = *(*int32)(unsafe.Add(mBase, _c_F_smgrtruncate[8]))
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v799+v797*int32(48))+32))
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v611)+36)) = v608
	*(*int32)(unsafe.Add(mBase, uint32(v611)+32)) = v803
	F_errmsg(m, int32(_a_F_smgrtruncate_10), v611+int32(32))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L22
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_smgrtruncate_6), int32(1360), int32(_a_F_smgrtruncate_7))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L22
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	v823 = v820
	goto L147
L146:
	;
	v823 = v821
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572+v816<<(uint(int32(2))%32)))) = v823
	v826 = v590 + int32(1)
	if v826 != l2 {
		v590 = v826
		goto L87
	} else {
		goto L148
	}
L148:
	;
	goto L88
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
				F_pg_qsort(m, v16, v10, int32(4), int32(982))
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
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v204 int32
	_ = v204
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v323 int32
	_ = v323
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v367 int32
	_ = v367
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v444 int32
	_ = v444
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v562 int32
	_ = v562
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v588 int32
	_ = v588
	var v596 int32
	_ = v596
	var v607 int32
	_ = v607
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v685 int32
	_ = v685
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v722 int32
	_ = v722
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v803 int32
	_ = v803
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v829 int32
	_ = v829
	var v840 int32
	_ = v840
	var v847 int32
	_ = v847
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v871 int32
	_ = v871
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v922 int32
	_ = v922
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v948 int32
	_ = v948
	var v955 int32
	_ = v955
	var v966 int32
	_ = v966
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1043 int32
	_ = v1043
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1069 int32
	_ = v1069
	var v1081 int32
	_ = v1081
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1178 int32
	_ = v1178
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1200 int32
	_ = v1200
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1229 int32
	_ = v1229
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1249 int32
	_ = v1249
	var v1255 int32
	_ = v1255
	var v1262 int32
	_ = v1262
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1300 int32
	_ = v1300
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1322 int32
	_ = v1322
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1338 int32
	_ = v1338
	var v1351 int32
	_ = v1351
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1371 int32
	_ = v1371
	var v1377 int32
	_ = v1377
	var v1385 int32
	_ = v1385
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1425 int32
	_ = v1425
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1447 int32
	_ = v1447
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1476 int32
	_ = v1476
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1496 int32
	_ = v1496
	var v1502 int32
	_ = v1502
	var v1509 int32
	_ = v1509
	var v1520 int32
	_ = v1520
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1547 int32
	_ = v1547
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1560 int32
	_ = v1560
	var v1563 int32
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1569 int32
	_ = v1569
	var v1579 int32
	_ = v1579
	var v1581 int32
	_ = v1581
	var v1585 int32
	_ = v1585
	var v1598 int32
	_ = v1598
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1618 int32
	_ = v1618
	var v1624 int32
	_ = v1624
	var v1632 int32
	_ = v1632
	var v1643 int32
	_ = v1643
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1660 int32
	_ = v1660
	var v1674 int32
	_ = v1674
	var v1677 int32
	_ = v1677
	var v1680 int32
	_ = v1680
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1748 int32
	_ = v1748
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1760 int32
	_ = v1760
	var v1763 int32
	_ = v1763
	var v1767 int32
	_ = v1767
	var v1769 int32
	_ = v1769
	var v1771 int32
	_ = v1771
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1789 int32
	_ = v1789
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1811 int32
	_ = v1811
	var v1814 int32
	_ = v1814
	var v1817 int32
	_ = v1817
	var v1820 int32
	_ = v1820
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1870 int32
	_ = v1870
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1878 int32
	_ = v1878
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1896 int32
	_ = v1896
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1907 int32
	_ = v1907
	var v1909 int32
	_ = v1909
	var v1911 int32
	_ = v1911
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1920 int32
	_ = v1920
	var v1924 int32
	_ = v1924
	var v1927 int32
	_ = v1927
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1943 int32
	_ = v1943
	var v1945 int32
	_ = v1945
	var v1949 int32
	_ = v1949
	var v1953 int32
	_ = v1953
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1962 int32
	_ = v1962
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1978 int32
	_ = v1978
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1986 int32
	_ = v1986
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2004 int32
	_ = v2004
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2020 int32
	_ = v2020
	var v2022 int32
	_ = v2022
	var v2024 int32
	_ = v2024
	var v2027 int32
	_ = v2027
	var v2030 int32
	_ = v2030
	var v2033 int32
	_ = v2033
	var v2037 int32
	_ = v2037
	var v2040 int32
	_ = v2040
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2051 int32
	_ = v2051
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2057 int32
	_ = v2057
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2067 int32
	_ = v2067
	var v2070 int32
	_ = v2070
	var v2074 int32
	_ = v2074
	var v2077 int32
	_ = v2077
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2088 int32
	_ = v2088
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2098 int32
	_ = v2098
	var v2102 int32
	_ = v2102
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2108 int32
	_ = v2108
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2118 int32
	_ = v2118
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2127 int32
	_ = v2127
	var v2131 int32
	_ = v2131
	var v2136 int32
	_ = v2136
	var v2139 int32
	_ = v2139
	var v2142 int32
	_ = v2142
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2149 int32
	_ = v2149
	var v2150 int32
	_ = v2150
	var v2157 int32
	_ = v2157
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2166 int32
	_ = v2166
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2184 int32
	_ = v2184
	var v2186 int32
	_ = v2186
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2192 int32
	_ = v2192
	var v2196 int32
	_ = v2196
	var v2202 int32
	_ = v2202
	var v2205 int32
	_ = v2205
	var v2206 int32
	_ = v2206
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2216 int32
	_ = v2216
	var v2218 int32
	_ = v2218
	var v2220 int32
	_ = v2220
	var v2222 int32
	_ = v2222
	var v2226 int32
	_ = v2226
	var v2228 int32
	_ = v2228
	var v2230 int32
	_ = v2230
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2263 int32
	_ = v2263
	var v2264 int32
	_ = v2264
	var v2269 int32
	_ = v2269
	var v2270 int32
	_ = v2270
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2283 int32
	_ = v2283
	var v2290 int32
	_ = v2290
	var v2292 int32
	_ = v2292
	var v2297 int32
	_ = v2297
	var v2299 int32
	_ = v2299
	var v2306 int32
	_ = v2306
	var v2309 int32
	_ = v2309
	var v2313 int32
	_ = v2313
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2335 int32
	_ = v2335
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2347 int32
	_ = v2347
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L7
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1178 = v12
	goto L263
L2:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1152)+8)) = v1151
	goto L1
L3:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1151 = v1148 + v1147
	goto L2
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L136
L5:
	;
	if v130 != 0 {
		goto L4
	} else {
		goto L29
	}
L6:
	;
	v130 = v123
	goto L5
L7:
	;
	if v25 <= v12 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v123 = int32(0)
	goto L6
L9:
	;
	v130 = int32(-1)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v41 = int32(1)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v26))))
	if base.Ui32(v43) < base.Ui32(int32(192)) {
		v100 = v43
		v101 = v41
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if int32(252) < v100 {
		v123 = v101
		goto L6
	} else {
		goto L25
	}
L13:
	;
	v47 = v12 + int32(1)
	if v47 == v25 {
		v100 = v43
		v101 = v41
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+v26))))
	v52 = v50 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v43) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v26))))
	v68 = v66 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v43) {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v56 = v12 + int32(2)
	if v56 != v25 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v100 = v43<<(uint(int32(6))%32)&int32(1984) | v52
	v101 = int32(2)
	goto L12
L19:
	;
	goto L18
L20:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v72))))
	v100 = v85&int32(63) | (v43<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v52<<(uint(int32(12))%32) | v68<<(uint(int32(6))%32))
	v101 = int32(4)
	goto L12
L21:
	;
	v72 = v12 + int32(3)
	if v72 != v25 {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v100 = v43<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v52<<(uint(int32(6))%32) | v68
	v101 = int32(3)
	goto L12
L24:
	;
	goto L23
L25:
	;
	v105 = v100 - int32(97)
	if v105 < int32(0) {
		v123 = v101
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v105)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v111)>>(uint(v105&int32(7))%32))&int32(1) == int32(0) {
		v123 = v101
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v101 + v12
	goto L28
L28:
	;
	goto L8
L29:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L32
L30:
	;
	if v248 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L31:
	;
	v248 = v241
	goto L30
L32:
	;
	if v144 <= v131 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v241 = int32(0)
	goto L31
L34:
	;
	v248 = int32(-1)
	goto L30
L35:
	;
	goto L36
L36:
	;
	v160 = int32(1)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v145))))
	if base.Ui32(v162) < base.Ui32(int32(192)) {
		v219 = v162
		v220 = v160
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if int32(252) < v219 {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	v166 = v131 + int32(1)
	if v166 == v144 {
		v219 = v162
		v220 = v160
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166+v145))))
	v171 = v169 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v162) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+v145))))
	v187 = v185 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v162) {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	v175 = v131 + int32(2)
	if v175 != v144 {
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v219 = v162<<(uint(int32(6))%32)&int32(1984) | v171
	v220 = int32(2)
	goto L37
L44:
	;
	goto L43
L45:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+v191))))
	v219 = v204&int32(63) | (v162<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v171<<(uint(int32(12))%32) | v187<<(uint(int32(6))%32))
	v220 = int32(4)
	goto L37
L46:
	;
	v191 = v131 + int32(3)
	if v191 != v144 {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v219 = v162<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v171<<(uint(int32(6))%32) | v187
	v220 = int32(3)
	goto L37
L49:
	;
	goto L48
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v220 + v131
	goto L54
L51:
	;
	v224 = v219 - int32(97)
	if v224 < int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v224)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v230)>>(uint(v224&int32(7))%32))&int32(1) != 0 {
		v241 = v220
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
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v272 = v262
	goto L60
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v131
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L86
L58:
	;
	if int32(0) <= v367 {
		v1147 = v367
		goto L3
	} else {
		goto L83
	}
L59:
	;
	v367 = v339
	goto L58
L60:
	;
	if v263 <= v272 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v367 = int32(-1)
	goto L58
L63:
	;
	goto L64
L64:
	;
	v279 = int32(1)
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272+v264))))
	if base.Ui32(v281) < base.Ui32(int32(192)) {
		v338 = v281
		v339 = v279
		goto L65
	} else {
		goto L66
	}
L65:
	;
	if int32(252) < v338 {
		goto L78
	} else {
		goto L79
	}
L66:
	;
	v285 = v272 + int32(1)
	if v285 == v263 {
		v338 = v281
		v339 = v279
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285+v264))))
	v290 = v288 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v281) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294+v264))))
	v306 = v304 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v281) {
		goto L74
	} else {
		goto L75
	}
L69:
	;
	v294 = v272 + int32(2)
	if v294 != v263 {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v338 = v281<<(uint(int32(6))%32)&int32(1984) | v290
	v339 = int32(2)
	goto L65
L72:
	;
	goto L71
L73:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264+v310))))
	v338 = v323&int32(63) | (v281<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v290<<(uint(int32(12))%32) | v306<<(uint(int32(6))%32))
	v339 = int32(4)
	goto L65
L74:
	;
	v310 = v272 + int32(3)
	if v310 != v263 {
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v338 = v281<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v290<<(uint(int32(6))%32) | v306
	v339 = int32(3)
	goto L65
L77:
	;
	goto L76
L78:
	;
	v356 = v339 + v272
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v356
	v272 = v356
	goto L60
L79:
	;
	v343 = v338 - int32(97)
	if v343 < int32(0) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v343)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v349)>>(uint(v343&int32(7))%32))&int32(1) != 0 {
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
	if v489 != 0 {
		goto L4
	} else {
		goto L108
	}
L85:
	;
	v489 = v482
	goto L84
L86:
	;
	if v384 <= v131 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v482 = int32(0)
	goto L85
L88:
	;
	v489 = int32(-1)
	goto L84
L89:
	;
	goto L90
L90:
	;
	v400 = int32(1)
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v385))))
	if base.Ui32(v402) < base.Ui32(int32(192)) {
		v459 = v402
		v460 = v400
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if int32(252) < v459 {
		v482 = v460
		goto L85
	} else {
		goto L104
	}
L92:
	;
	v406 = v131 + int32(1)
	if v406 == v384 {
		v459 = v402
		v460 = v400
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406+v385))))
	v411 = v409 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v402) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v385))))
	v427 = v425 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v402) {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	v415 = v131 + int32(2)
	if v415 != v384 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v459 = v402<<(uint(int32(6))%32)&int32(1984) | v411
	v460 = int32(2)
	goto L91
L98:
	;
	goto L97
L99:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385+v431))))
	v459 = v444&int32(63) | (v402<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v411<<(uint(int32(12))%32) | v427<<(uint(int32(6))%32))
	v460 = int32(4)
	goto L91
L100:
	;
	v431 = v131 + int32(3)
	if v431 != v384 {
		goto L99
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v459 = v402<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v411<<(uint(int32(6))%32) | v427
	v460 = int32(3)
	goto L91
L103:
	;
	goto L102
L104:
	;
	v464 = v459 - int32(97)
	if v464 < int32(0) {
		v482 = v460
		goto L85
	} else {
		goto L105
	}
L105:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v464)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v470)>>(uint(v464&int32(7))%32))&int32(1) == int32(0) {
		v482 = v460
		goto L85
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + v131
	goto L107
L107:
	;
	goto L87
L108:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v511 = v501
	goto L111
L109:
	;
	if int32(0) <= v607 {
		v1147 = v607
		goto L3
	} else {
		goto L133
	}
L110:
	;
	v607 = v578
	goto L109
L111:
	;
	if v502 <= v511 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v607 = int32(-1)
	goto L109
L114:
	;
	goto L115
L115:
	;
	v518 = int32(1)
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511+v503))))
	if base.Ui32(v520) < base.Ui32(int32(192)) {
		v577 = v520
		v578 = v518
		goto L116
	} else {
		goto L117
	}
L116:
	;
	if int32(252) < v577 {
		goto L110
	} else {
		goto L129
	}
L117:
	;
	v524 = v511 + int32(1)
	if v524 == v502 {
		v577 = v520
		v578 = v518
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524+v503))))
	v529 = v527 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v520) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533+v503))))
	v545 = v543 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v520) {
		goto L125
	} else {
		goto L126
	}
L120:
	;
	v533 = v511 + int32(2)
	if v533 != v502 {
		goto L119
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v577 = v520<<(uint(int32(6))%32)&int32(1984) | v529
	v578 = int32(2)
	goto L116
L123:
	;
	goto L122
L124:
	;
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503+v549))))
	v577 = v562&int32(63) | (v520<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v529<<(uint(int32(12))%32) | v545<<(uint(int32(6))%32))
	v578 = int32(4)
	goto L116
L125:
	;
	v549 = v511 + int32(3)
	if v549 != v502 {
		goto L124
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v577 = v520<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v529<<(uint(int32(6))%32) | v545
	v578 = int32(3)
	goto L116
L128:
	;
	goto L127
L129:
	;
	v582 = v577 - int32(97)
	if v582 < int32(0) {
		goto L110
	} else {
		goto L130
	}
L130:
	;
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v582)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v588)>>(uint(v582&int32(7))%32))&int32(1) == int32(0) {
		goto L110
	} else {
		goto L131
	}
L131:
	;
	v596 = v578 + v511
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
	v511 = v596
	goto L111
L133:
	;
	goto L4
L134:
	;
	if v729 != 0 {
		goto L1
	} else {
		goto L159
	}
L135:
	;
	v729 = v722
	goto L134
L136:
	;
	if v625 <= v12 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v722 = int32(0)
	goto L135
L138:
	;
	v729 = int32(-1)
	goto L134
L139:
	;
	goto L140
L140:
	;
	v641 = int32(1)
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v626))))
	if base.Ui32(v643) < base.Ui32(int32(192)) {
		v700 = v643
		v701 = v641
		goto L141
	} else {
		goto L142
	}
L141:
	;
	if int32(252) < v700 {
		goto L154
	} else {
		goto L155
	}
L142:
	;
	v647 = v12 + int32(1)
	if v647 == v625 {
		v700 = v643
		v701 = v641
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v647+v626))))
	v652 = v650 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v643) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656+v626))))
	v668 = v666 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v643) {
		goto L150
	} else {
		goto L151
	}
L145:
	;
	v656 = v12 + int32(2)
	if v656 != v625 {
		goto L144
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v700 = v643<<(uint(int32(6))%32)&int32(1984) | v652
	v701 = int32(2)
	goto L141
L148:
	;
	goto L147
L149:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626+v672))))
	v700 = v685&int32(63) | (v643<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v652<<(uint(int32(12))%32) | v668<<(uint(int32(6))%32))
	v701 = int32(4)
	goto L141
L150:
	;
	v672 = v12 + int32(3)
	if v672 != v625 {
		goto L149
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v700 = v643<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v652<<(uint(int32(6))%32) | v668
	v701 = int32(3)
	goto L141
L153:
	;
	goto L152
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v701 + v12
	goto L158
L155:
	;
	v705 = v700 - int32(97)
	if v705 < int32(0) {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v705)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v711)>>(uint(v705&int32(7))%32))&int32(1) != 0 {
		v722 = v701
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
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L162
L160:
	;
	if v847 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L161:
	;
	v847 = v840
	goto L160
L162:
	;
	if v743 <= v730 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v840 = int32(0)
	goto L161
L164:
	;
	v847 = int32(-1)
	goto L160
L165:
	;
	goto L166
L166:
	;
	v759 = int32(1)
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v730+v744))))
	if base.Ui32(v761) < base.Ui32(int32(192)) {
		v818 = v761
		v819 = v759
		goto L167
	} else {
		goto L168
	}
L167:
	;
	if int32(252) < v818 {
		goto L180
	} else {
		goto L181
	}
L168:
	;
	v765 = v730 + int32(1)
	if v765 == v743 {
		v818 = v761
		v819 = v759
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765+v744))))
	v770 = v768 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v761) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v774+v744))))
	v786 = v784 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v761) {
		goto L176
	} else {
		goto L177
	}
L171:
	;
	v774 = v730 + int32(2)
	if v774 != v743 {
		goto L170
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v818 = v761<<(uint(int32(6))%32)&int32(1984) | v770
	v819 = int32(2)
	goto L167
L174:
	;
	goto L173
L175:
	;
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744+v790))))
	v818 = v803&int32(63) | (v761<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v770<<(uint(int32(12))%32) | v786<<(uint(int32(6))%32))
	v819 = int32(4)
	goto L167
L176:
	;
	v790 = v730 + int32(3)
	if v790 != v743 {
		goto L175
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v818 = v761<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v770<<(uint(int32(6))%32) | v786
	v819 = int32(3)
	goto L167
L179:
	;
	goto L178
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v819 + v730
	goto L184
L181:
	;
	v823 = v818 - int32(97)
	if v823 < int32(0) {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v823)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v829)>>(uint(v823&int32(7))%32))&int32(1) != 0 {
		v840 = v819
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
	v861 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v862 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v863 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v871 = v861
	goto L190
L186:
	;
	goto L187
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v730
	v983 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L216
L188:
	;
	if int32(0) <= v966 {
		v1147 = v966
		goto L3
	} else {
		goto L213
	}
L189:
	;
	v966 = v938
	goto L188
L190:
	;
	if v862 <= v871 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v966 = int32(-1)
	goto L188
L193:
	;
	goto L194
L194:
	;
	v878 = int32(1)
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871+v863))))
	if base.Ui32(v880) < base.Ui32(int32(192)) {
		v937 = v880
		v938 = v878
		goto L195
	} else {
		goto L196
	}
L195:
	;
	if int32(252) < v937 {
		goto L208
	} else {
		goto L209
	}
L196:
	;
	v884 = v871 + int32(1)
	if v884 == v862 {
		v937 = v880
		v938 = v878
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v884+v863))))
	v889 = v887 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v880) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v893+v863))))
	v905 = v903 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v880) {
		goto L204
	} else {
		goto L205
	}
L199:
	;
	v893 = v871 + int32(2)
	if v893 != v862 {
		goto L198
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v937 = v880<<(uint(int32(6))%32)&int32(1984) | v889
	v938 = int32(2)
	goto L195
L202:
	;
	goto L201
L203:
	;
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v863+v909))))
	v937 = v922&int32(63) | (v880<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v889<<(uint(int32(12))%32) | v905<<(uint(int32(6))%32))
	v938 = int32(4)
	goto L195
L204:
	;
	v909 = v871 + int32(3)
	if v909 != v862 {
		goto L203
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	v937 = v880<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v889<<(uint(int32(6))%32) | v905
	v938 = int32(3)
	goto L195
L207:
	;
	goto L206
L208:
	;
	v955 = v938 + v871
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v955
	v871 = v955
	goto L190
L209:
	;
	v942 = v937 - int32(97)
	if v942 < int32(0) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v942)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v948)>>(uint(v942&int32(7))%32))&int32(1) != 0 {
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
	if v1088 != 0 {
		goto L1
	} else {
		goto L238
	}
L215:
	;
	v1088 = v1081
	goto L214
L216:
	;
	if v983 <= v730 {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	v1081 = int32(0)
	goto L215
L218:
	;
	v1088 = int32(-1)
	goto L214
L219:
	;
	goto L220
L220:
	;
	v999 = int32(1)
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v730+v984))))
	if base.Ui32(v1001) < base.Ui32(int32(192)) {
		v1058 = v1001
		v1059 = v999
		goto L221
	} else {
		goto L222
	}
L221:
	;
	if int32(252) < v1058 {
		v1081 = v1059
		goto L215
	} else {
		goto L234
	}
L222:
	;
	v1005 = v730 + int32(1)
	if v1005 == v983 {
		v1058 = v1001
		v1059 = v999
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v1008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1005+v984))))
	v1010 = v1008 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1001) {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	v1024 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014+v984))))
	v1026 = v1024 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1001) {
		goto L230
	} else {
		goto L231
	}
L225:
	;
	v1014 = v730 + int32(2)
	if v1014 != v983 {
		goto L224
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v1058 = v1001<<(uint(int32(6))%32)&int32(1984) | v1010
	v1059 = int32(2)
	goto L221
L228:
	;
	goto L227
L229:
	;
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v984+v1030))))
	v1058 = v1043&int32(63) | (v1001<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v1010<<(uint(int32(12))%32) | v1026<<(uint(int32(6))%32))
	v1059 = int32(4)
	goto L221
L230:
	;
	v1030 = v730 + int32(3)
	if v1030 != v983 {
		goto L229
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v1058 = v1001<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v1010<<(uint(int32(6))%32) | v1026
	v1059 = int32(3)
	goto L221
L233:
	;
	goto L232
L234:
	;
	v1063 = v1058 - int32(97)
	if v1063 < int32(0) {
		v1081 = v1059
		goto L215
	} else {
		goto L235
	}
L235:
	;
	v1069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1063)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v1069)>>(uint(v1063&int32(7))%32))&int32(1) == int32(0) {
		v1081 = v1059
		goto L215
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1059 + v730
	goto L237
L237:
	;
	goto L217
L238:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L241
L239:
	;
	if int32(0) <= v1143 {
		v1151 = v1143
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
	v1098 = v1090
	v1100 = int32(1)
	goto L246
L245:
	;
	v1143 = v1128
	goto L239
L246:
	;
	if v1091 <= v1098 {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	goto L245
L248:
	;
	v1143 = int32(-1)
	goto L239
L249:
	;
	goto L250
L250:
	;
	v1105 = v1098 + int32(1)
	v1107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1089+v1098))))
	if base.Ui32(v1107) < base.Ui32(int32(192)) {
		v1128 = v1105
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1129 = int32(1)
	if v1129 < v1100 {
		v1098 = v1128
		v1100 = v1100 - v1129
		goto L246
	} else {
		goto L258
	}
L252:
	;
	if v1091 <= v1105 {
		v1128 = v1105
		goto L251
	} else {
		goto L253
	}
L253:
	;
	v1114 = v1105
	goto L254
L254:
	;
	v1117 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1089+v1114))))
	if int32(-65) < v1117 {
		v1128 = v1114
		goto L251
	} else {
		goto L256
	}
L255:
	;
	v1128 = v1091
	goto L251
L256:
	;
	v1121 = v1114 + int32(1)
	if v1121 != v1091 {
		v1114 = v1121
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1652
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1652
	v1656 = v1652 - int32(1)
	if v1656 <= v12 {
		goto L364
	} else {
		goto L365
	}
L261:
	;
	if v1273 < int32(0) {
		goto L260
	} else {
		goto L286
	}
L262:
	;
	v1273 = v1245
	goto L261
L263:
	;
	if v1169 <= v1178 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1273 = int32(-1)
	goto L261
L266:
	;
	goto L267
L267:
	;
	v1185 = int32(1)
	v1187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1178+v1170))))
	if base.Ui32(v1187) < base.Ui32(int32(192)) {
		v1244 = v1187
		v1245 = v1185
		goto L268
	} else {
		goto L269
	}
L268:
	;
	if int32(252) < v1244 {
		goto L281
	} else {
		goto L282
	}
L269:
	;
	v1191 = v1178 + int32(1)
	if v1191 == v1169 {
		v1244 = v1187
		v1245 = v1185
		goto L268
	} else {
		goto L270
	}
L270:
	;
	v1194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1191+v1170))))
	v1196 = v1194 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1187) {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	v1210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1200+v1170))))
	v1212 = v1210 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1187) {
		goto L277
	} else {
		goto L278
	}
L272:
	;
	v1200 = v1178 + int32(2)
	if v1200 != v1169 {
		goto L271
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v1244 = v1187<<(uint(int32(6))%32)&int32(1984) | v1196
	v1245 = int32(2)
	goto L268
L275:
	;
	goto L274
L276:
	;
	v1229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1170+v1216))))
	v1244 = v1229&int32(63) | (v1187<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v1196<<(uint(int32(12))%32) | v1212<<(uint(int32(6))%32))
	v1245 = int32(4)
	goto L268
L277:
	;
	v1216 = v1178 + int32(3)
	if v1216 != v1169 {
		goto L276
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v1244 = v1187<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v1196<<(uint(int32(6))%32) | v1212
	v1245 = int32(3)
	goto L268
L280:
	;
	goto L279
L281:
	;
	v1262 = v1245 + v1178
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1262
	v1178 = v1262
	goto L263
L282:
	;
	v1249 = v1244 - int32(97)
	if v1249 < int32(0) {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1249)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v1255)>>(uint(v1249&int32(7))%32))&int32(1) != 0 {
		goto L262
	} else {
		goto L284
	}
L284:
	;
	goto L281
L286:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1277 = v1276 + v1273
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1277
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1300 = v1277
	goto L289
L287:
	;
	if v1396 < int32(0) {
		goto L260
	} else {
		goto L311
	}
L288:
	;
	v1396 = v1367
	goto L287
L289:
	;
	if v1291 <= v1300 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1396 = int32(-1)
	goto L287
L292:
	;
	goto L293
L293:
	;
	v1307 = int32(1)
	v1309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1300+v1292))))
	if base.Ui32(v1309) < base.Ui32(int32(192)) {
		v1366 = v1309
		v1367 = v1307
		goto L294
	} else {
		goto L295
	}
L294:
	;
	if int32(252) < v1366 {
		goto L288
	} else {
		goto L307
	}
L295:
	;
	v1313 = v1300 + int32(1)
	if v1313 == v1291 {
		v1366 = v1309
		v1367 = v1307
		goto L294
	} else {
		goto L296
	}
L296:
	;
	v1316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1313+v1292))))
	v1318 = v1316 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1309) {
		goto L298
	} else {
		goto L299
	}
L297:
	;
	v1332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322+v1292))))
	v1334 = v1332 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1309) {
		goto L303
	} else {
		goto L304
	}
L298:
	;
	v1322 = v1300 + int32(2)
	if v1322 != v1291 {
		goto L297
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	v1366 = v1309<<(uint(int32(6))%32)&int32(1984) | v1318
	v1367 = int32(2)
	goto L294
L301:
	;
	goto L300
L302:
	;
	v1351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1292+v1338))))
	v1366 = v1351&int32(63) | (v1309<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v1318<<(uint(int32(12))%32) | v1334<<(uint(int32(6))%32))
	v1367 = int32(4)
	goto L294
L303:
	;
	v1338 = v1300 + int32(3)
	if v1338 != v1291 {
		goto L302
	} else {
		goto L306
	}
L304:
	;
	goto L305
L305:
	;
	v1366 = v1309<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v1318<<(uint(int32(6))%32) | v1334
	v1367 = int32(3)
	goto L294
L306:
	;
	goto L305
L307:
	;
	v1371 = v1366 - int32(97)
	if v1371 < int32(0) {
		goto L288
	} else {
		goto L308
	}
L308:
	;
	v1377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1371)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v1377)>>(uint(v1371&int32(7))%32))&int32(1) == int32(0) {
		goto L288
	} else {
		goto L309
	}
L309:
	;
	v1385 = v1367 + v1300
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1385
	v1300 = v1385
	goto L289
L311:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1400 = v1399 + v1396
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1400
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1402)+4)) = v1400
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1425 = v1415
	goto L314
L312:
	;
	if v1520 < int32(0) {
		goto L260
	} else {
		goto L337
	}
L313:
	;
	v1520 = v1492
	goto L312
L314:
	;
	if v1416 <= v1425 {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v1520 = int32(-1)
	goto L312
L317:
	;
	goto L318
L318:
	;
	v1432 = int32(1)
	v1434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1425+v1417))))
	if base.Ui32(v1434) < base.Ui32(int32(192)) {
		v1491 = v1434
		v1492 = v1432
		goto L319
	} else {
		goto L320
	}
L319:
	;
	if int32(252) < v1491 {
		goto L332
	} else {
		goto L333
	}
L320:
	;
	v1438 = v1425 + int32(1)
	if v1438 == v1416 {
		v1491 = v1434
		v1492 = v1432
		goto L319
	} else {
		goto L321
	}
L321:
	;
	v1441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1438+v1417))))
	v1443 = v1441 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1434) {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	v1457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1447+v1417))))
	v1459 = v1457 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1434) {
		goto L328
	} else {
		goto L329
	}
L323:
	;
	v1447 = v1425 + int32(2)
	if v1447 != v1416 {
		goto L322
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	v1491 = v1434<<(uint(int32(6))%32)&int32(1984) | v1443
	v1492 = int32(2)
	goto L319
L326:
	;
	goto L325
L327:
	;
	v1476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1417+v1463))))
	v1491 = v1476&int32(63) | (v1434<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v1443<<(uint(int32(12))%32) | v1459<<(uint(int32(6))%32))
	v1492 = int32(4)
	goto L319
L328:
	;
	v1463 = v1425 + int32(3)
	if v1463 != v1416 {
		goto L327
	} else {
		goto L331
	}
L329:
	;
	goto L330
L330:
	;
	v1491 = v1434<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v1443<<(uint(int32(6))%32) | v1459
	v1492 = int32(3)
	goto L319
L331:
	;
	goto L330
L332:
	;
	v1509 = v1492 + v1425
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1509
	v1425 = v1509
	goto L314
L333:
	;
	v1496 = v1491 - int32(97)
	if v1496 < int32(0) {
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v1502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1496)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v1502)>>(uint(v1496&int32(7))%32))&int32(1) != 0 {
		goto L313
	} else {
		goto L335
	}
L335:
	;
	goto L332
L337:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1524 = v1523 + v1520
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1524
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1547 = v1524
	goto L340
L338:
	;
	if v1643 < int32(0) {
		goto L260
	} else {
		goto L362
	}
L339:
	;
	v1643 = v1614
	goto L338
L340:
	;
	if v1538 <= v1547 {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1643 = int32(-1)
	goto L338
L343:
	;
	goto L344
L344:
	;
	v1554 = int32(1)
	v1556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1547+v1539))))
	if base.Ui32(v1556) < base.Ui32(int32(192)) {
		v1613 = v1556
		v1614 = v1554
		goto L345
	} else {
		goto L346
	}
L345:
	;
	if int32(252) < v1613 {
		goto L339
	} else {
		goto L358
	}
L346:
	;
	v1560 = v1547 + int32(1)
	if v1560 == v1538 {
		v1613 = v1556
		v1614 = v1554
		goto L345
	} else {
		goto L347
	}
L347:
	;
	v1563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1560+v1539))))
	v1565 = v1563 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1556) {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	v1579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1569+v1539))))
	v1581 = v1579 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1556) {
		goto L354
	} else {
		goto L355
	}
L349:
	;
	v1569 = v1547 + int32(2)
	if v1569 != v1538 {
		goto L348
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1613 = v1556<<(uint(int32(6))%32)&int32(1984) | v1565
	v1614 = int32(2)
	goto L345
L352:
	;
	goto L351
L353:
	;
	v1598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1539+v1585))))
	v1613 = v1598&int32(63) | (v1556<<(uint(int32(18))%32)&int32(_a_F_spanish_UTF_8_stem_0) | v1565<<(uint(int32(12))%32) | v1581<<(uint(int32(6))%32))
	v1614 = int32(4)
	goto L345
L354:
	;
	v1585 = v1547 + int32(3)
	if v1585 != v1538 {
		goto L353
	} else {
		goto L357
	}
L355:
	;
	goto L356
L356:
	;
	v1613 = v1556<<(uint(int32(12))%32)&int32(_a_F_spanish_UTF_8_stem_1) | v1565<<(uint(int32(6))%32) | v1581
	v1614 = int32(3)
	goto L345
L357:
	;
	goto L356
L358:
	;
	v1618 = v1613 - int32(97)
	if v1618 < int32(0) {
		goto L339
	} else {
		goto L359
	}
L359:
	;
	v1624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1618)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_UTF_8_stem[0]))))
	if int32(base.Ui32(v1624)>>(uint(v1618&int32(7))%32))&int32(1) == int32(0) {
		goto L339
	} else {
		goto L360
	}
L360:
	;
	v1632 = v1614 + v1547
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1632
	v1547 = v1632
	goto L340
L362:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1646))) = v1647 + v1643
	goto L260
L363:
	;
	return v2347
L364:
	;
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1760
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1760
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1760-int32(2) <= v1763 {
		goto L399
	} else {
		goto L400
	}
L365:
	;
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1658+v1656))))
	if base.B2i32(v1660&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1660)%32)&int32(_a_F_spanish_UTF_8_stem_2) == int32(0)) != 0 {
		goto L364
	} else {
		goto L366
	}
L366:
	;
	v1674 = F_find_among_b(m, l0, int32(_a_F_spanish_UTF_8_stem_3), int32(13))
	mBase = m.M
	v1677 = m.ExcPending
	if v1677 != 0 {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	return int32(0)
L368:
	;
	if v1674 == int32(0) {
		goto L364
	} else {
		goto L369
	}
L369:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1680
	v1683 = v1680 - int32(1)
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1683 <= v1684 {
		goto L364
	} else {
		goto L370
	}
L370:
	;
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1686+v1683))))
	switch v1688 - int32(111) {
	case 0, 3:
		goto L371
	default:
		goto L364
	}
L371:
	;
	v1693 = F_find_among_b(m, l0, int32(_a_F_spanish_UTF_8_stem_4), int32(11))
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L367
	} else {
		goto L372
	}
L372:
	;
	if v1693 == int32(0) {
		goto L364
	} else {
		goto L373
	}
L373:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(v1698)+8))
	if v1697 < v1699 {
		goto L364
	} else {
		goto L374
	}
L374:
	;
	switch v1693 - int32(1) {
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
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1697 <= v1742 {
		goto L364
	} else {
		goto L394
	}
L376:
	;
	v1738 = F_slice_del(m, l0)
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L367
	} else {
		goto L392
	}
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1697
	v1734 = F_slice_from_s(m, l0, int32(2), int32(_a_F_spanish_UTF_8_stem_5))
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		goto L367
	} else {
		goto L390
	}
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1697
	v1727 = F_slice_from_s(m, l0, int32(2), int32(_a_F_spanish_UTF_8_stem_6))
	mBase = m.M
	v1728 = m.ExcPending
	if v1728 != 0 {
		goto L367
	} else {
		goto L388
	}
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1697
	v1720 = F_slice_from_s(m, l0, int32(2), int32(_a_F_spanish_UTF_8_stem_7))
	mBase = m.M
	v1721 = m.ExcPending
	if v1721 != 0 {
		goto L367
	} else {
		goto L386
	}
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1697
	v1713 = F_slice_from_s(m, l0, int32(4), int32(_a_F_spanish_UTF_8_stem_8))
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L367
	} else {
		goto L384
	}
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1697
	v1706 = F_slice_from_s(m, l0, int32(5), int32(_a_F_spanish_UTF_8_stem_9))
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L367
	} else {
		goto L382
	}
L382:
	;
	if int32(0) <= v1706 {
		goto L364
	} else {
		goto L383
	}
L383:
	;
	v2347 = v1706
	goto L363
L384:
	;
	if int32(0) <= v1713 {
		goto L364
	} else {
		goto L385
	}
L385:
	;
	v2347 = v1713
	goto L363
L386:
	;
	if int32(0) <= v1720 {
		goto L364
	} else {
		goto L387
	}
L387:
	;
	v2347 = v1720
	goto L363
L388:
	;
	if int32(0) <= v1727 {
		goto L364
	} else {
		goto L389
	}
L389:
	;
	v2347 = v1727
	goto L363
L390:
	;
	if int32(0) <= v1734 {
		goto L364
	} else {
		goto L391
	}
L391:
	;
	v2347 = v1734
	goto L363
L392:
	;
	if int32(0) <= v1738 {
		goto L364
	} else {
		goto L393
	}
L393:
	;
	v2347 = v1738
	goto L363
L394:
	;
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1744+v1697-int32(1)))))
	if v1748 != int32(117) {
		goto L364
	} else {
		goto L395
	}
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1697 - int32(1)
	v1754 = F_slice_del(m, l0)
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L367
	} else {
		goto L396
	}
L396:
	;
	if v1754 < int32(0) {
		v2347 = v1754
		goto L363
	} else {
		goto L397
	}
L397:
	;
	goto L364
L398:
	;
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2157
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2157
	v2162 = F_find_among_b(m, l0, int32(_a_F_spanish_UTF_8_stem_10), int32(8))
	mBase = m.M
	v2163 = m.ExcPending
	if v2163 != 0 {
		goto L367
	} else {
		goto L525
	}
L399:
	;
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2051
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(v2053)+8))
	if v2051 < v2054 {
		goto L486
	} else {
		goto L487
	}
L400:
	;
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1769 = int32(1)
	v1771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1767+v1760-v1769))))
	if base.B2i32(v1771&int32(224) != int32(96))|base.B2i32(v1769<<(uint(v1771)%32)&int32(_a_F_spanish_UTF_8_stem_11) == int32(0)) != 0 {
		goto L399
	} else {
		goto L401
	}
L401:
	;
	v1785 = F_find_among_b(m, l0, int32(_a_F_spanish_UTF_8_stem_12), int32(46))
	mBase = m.M
	v1786 = m.ExcPending
	if v1786 != 0 {
		goto L367
	} else {
		goto L402
	}
L402:
	;
	if v1785 == int32(0) {
		goto L399
	} else {
		goto L403
	}
L403:
	;
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1789
	switch v1785 - int32(1) {
	case 0:
		goto L412
	case 1:
		goto L411
	case 2:
		goto L410
	case 3:
		goto L409
	case 4:
		goto L408
	case 5:
		goto L407
	case 6:
		goto L406
	case 7:
		goto L405
	case 8:
		goto L404
	default:
		goto L398
	}
L404:
	;
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(v2013)))
	if v1789 < v2014 {
		goto L399
	} else {
		goto L475
	}
L405:
	;
	v1971 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v1971)))
	if v1789 < v1972 {
		goto L399
	} else {
		goto L465
	}
L406:
	;
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v1936)))
	if v1789 < v1937 {
		goto L399
	} else {
		goto L455
	}
L407:
	;
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(v1863)+4))
	if v1789 < v1864 {
		goto L399
	} else {
		goto L436
	}
L408:
	;
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v1854)))
	if v1789 < v1855 {
		goto L399
	} else {
		goto L433
	}
L409:
	;
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1845)))
	if v1789 < v1846 {
		goto L399
	} else {
		goto L430
	}
L410:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v1836)))
	if v1789 < v1837 {
		goto L399
	} else {
		goto L427
	}
L411:
	;
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(v1800)))
	if v1789 < v1801 {
		goto L399
	} else {
		goto L416
	}
L412:
	;
	v1793 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(v1793)))
	if v1789 < v1794 {
		goto L399
	} else {
		goto L413
	}
L413:
	;
	v1796 = F_slice_del(m, l0)
	mBase = m.M
	v1797 = m.ExcPending
	if v1797 != 0 {
		goto L367
	} else {
		goto L414
	}
L414:
	;
	if int32(0) <= v1796 {
		goto L398
	} else {
		goto L415
	}
L415:
	;
	v2347 = v1796
	goto L363
L416:
	;
	v1803 = F_slice_del(m, l0)
	mBase = m.M
	v1804 = m.ExcPending
	if v1804 != 0 {
		goto L367
	} else {
		goto L417
	}
L417:
	;
	if v1803 < int32(0) {
		v2347 = v1803
		goto L363
	} else {
		goto L418
	}
L418:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1807
	v1809 = int32(2)
	v1811 = int32(0)
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1807-v1814 < v1809 {
		v1824 = v1811
		goto L420
	} else {
		goto L421
	}
L419:
	;
	if v1824 == int32(0) {
		goto L398
	} else {
		goto L423
	}
L420:
	;
	goto L419
L421:
	;
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1820 = F_memcmp(m, v1817+v1807-v1809, int32(_a_F_spanish_UTF_8_stem_13), v1809)
	mBase = m.M
	if v1820 != 0 {
		v1824 = v1811
		goto L420
	} else {
		goto L422
	}
L422:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1807 - v1809
	v1824 = int32(1)
	goto L420
L423:
	;
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1827
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1830 = *(*int32)(unsafe.Add(mBase, uint32(v1829)))
	if v1827 < v1830 {
		goto L398
	} else {
		goto L424
	}
L424:
	;
	v1832 = F_slice_del(m, l0)
	mBase = m.M
	v1833 = m.ExcPending
	if v1833 != 0 {
		goto L367
	} else {
		goto L425
	}
L425:
	;
	if int32(0) <= v1832 {
		goto L398
	} else {
		goto L426
	}
L426:
	;
	v2347 = v1832
	goto L363
L427:
	;
	v1841 = F_slice_from_s(m, l0, int32(3), int32(_a_F_spanish_UTF_8_stem_14))
	mBase = m.M
	v1842 = m.ExcPending
	if v1842 != 0 {
		goto L367
	} else {
		goto L428
	}
L428:
	;
	if int32(0) <= v1841 {
		goto L398
	} else {
		goto L429
	}
L429:
	;
	v2347 = v1841
	goto L363
L430:
	;
	v1850 = F_slice_from_s(m, l0, int32(1), int32(_a_F_spanish_UTF_8_stem_15))
	mBase = m.M
	v1851 = m.ExcPending
	if v1851 != 0 {
		goto L367
	} else {
		goto L431
	}
L431:
	;
	if int32(0) <= v1850 {
		goto L398
	} else {
		goto L432
	}
L432:
	;
	v2347 = v1850
	goto L363
L433:
	;
	v1859 = F_slice_from_s(m, l0, int32(4), int32(_a_F_spanish_UTF_8_stem_16))
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L367
	} else {
		goto L434
	}
L434:
	;
	if int32(0) <= v1859 {
		goto L398
	} else {
		goto L435
	}
L435:
	;
	v2347 = v1859
	goto L363
L436:
	;
	v1866 = F_slice_del(m, l0)
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		goto L367
	} else {
		goto L437
	}
L437:
	;
	if v1866 < int32(0) {
		v2347 = v1866
		goto L363
	} else {
		goto L438
	}
L438:
	;
	v1870 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1870
	v1873 = v1870 - int32(1)
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1873 <= v1874 {
		goto L398
	} else {
		goto L439
	}
L439:
	;
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1876+v1873))))
	if base.B2i32(v1878&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1878)%32)&int32(_a_F_spanish_UTF_8_stem_17) == int32(0)) != 0 {
		goto L398
	} else {
		goto L440
	}
L440:
	;
	v1892 = F_find_among_b(m, l0, int32(_a_F_spanish_UTF_8_stem_18), int32(4))
	mBase = m.M
	v1893 = m.ExcPending
	if v1893 != 0 {
		goto L367
	} else {
		goto L441
	}
L441:
	;
	if v1892 == int32(0) {
		goto L398
	} else {
		goto L442
	}
L442:
	;
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1896
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1899 = *(*int32)(unsafe.Add(mBase, uint32(v1898)))
	if v1896 < v1899 {
		goto L398
	} else {
		goto L443
	}
L443:
	;
	v1901 = F_slice_del(m, l0)
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L367
	} else {
		goto L444
	}
L444:
	;
	if v1901 < int32(0) {
		v2347 = v1901
		goto L363
	} else {
		goto L445
	}
L445:
	;
	if v1892 != int32(1) {
		goto L398
	} else {
		goto L446
	}
L446:
	;
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1907
	v1909 = int32(2)
	v1911 = int32(0)
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1907-v1914 < v1909 {
		v1924 = v1911
		goto L448
	} else {
		goto L449
	}
L447:
	;
	if v1924 == int32(0) {
		goto L398
	} else {
		goto L451
	}
L448:
	;
	goto L447
L449:
	;
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1920 = F_memcmp(m, v1917+v1907-v1909, int32(_a_F_spanish_UTF_8_stem_19), v1909)
	mBase = m.M
	if v1920 != 0 {
		v1924 = v1911
		goto L448
	} else {
		goto L450
	}
L450:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1907 - v1909
	v1924 = int32(1)
	goto L448
L451:
	;
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1927
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1930 = *(*int32)(unsafe.Add(mBase, uint32(v1929)))
	if v1927 < v1930 {
		goto L398
	} else {
		goto L452
	}
L452:
	;
	v1932 = F_slice_del(m, l0)
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L367
	} else {
		goto L453
	}
L453:
	;
	if int32(0) <= v1932 {
		goto L398
	} else {
		goto L454
	}
L454:
	;
	v2347 = v1932
	goto L363
L455:
	;
	v1939 = F_slice_del(m, l0)
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L367
	} else {
		goto L456
	}
L456:
	;
	if v1939 < int32(0) {
		v2347 = v1939
		goto L363
	} else {
		goto L457
	}
L457:
	;
	v1943 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1943
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1943-int32(3) <= v1945 {
		goto L398
	} else {
		goto L458
	}
L458:
	;
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1949+v1943-int32(1)))))
	if v1953 != int32(101) {
		goto L398
	} else {
		goto L459
	}
L459:
	;
	v1958 = F_find_among_b(m, l0, int32(_a_F_spanish_UTF_8_stem_20), int32(3))
	mBase = m.M
	v1959 = m.ExcPending
	if v1959 != 0 {
		goto L367
	} else {
		goto L460
	}
L460:
	;
	if v1958 == int32(0) {
		goto L398
	} else {
		goto L461
	}
L461:
	;
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1962
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v1964)))
	if v1962 < v1965 {
		goto L398
	} else {
		goto L462
	}
L462:
	;
	v1967 = F_slice_del(m, l0)
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		goto L367
	} else {
		goto L463
	}
L463:
	;
	if int32(0) <= v1967 {
		goto L398
	} else {
		goto L464
	}
L464:
	;
	v2347 = v1967
	goto L363
L465:
	;
	v1974 = F_slice_del(m, l0)
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L367
	} else {
		goto L466
	}
L466:
	;
	if v1974 < int32(0) {
		v2347 = v1974
		goto L363
	} else {
		goto L467
	}
L467:
	;
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1978
	v1981 = v1978 - int32(1)
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1981 <= v1982 {
		goto L398
	} else {
		goto L468
	}
L468:
	;
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1984+v1981))))
	if base.B2i32(v1986&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1986)%32)&int32(_a_F_spanish_UTF_8_stem_21) == int32(0)) != 0 {
		goto L398
	} else {
		goto L469
	}
L469:
	;
	v2000 = F_find_among_b(m, l0, int32(_a_F_spanish_UTF_8_stem_22), int32(3))
	mBase = m.M
	v2001 = m.ExcPending
	if v2001 != 0 {
		goto L367
	} else {
		goto L470
	}
L470:
	;
	if v2000 == int32(0) {
		goto L398
	} else {
		goto L471
	}
L471:
	;
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2004
	v2006 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2007 = *(*int32)(unsafe.Add(mBase, uint32(v2006)))
	if v2004 < v2007 {
		goto L398
	} else {
		goto L472
	}
L472:
	;
	v2009 = F_slice_del(m, l0)
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L367
	} else {
		goto L473
	}
L473:
	;
	if int32(0) <= v2009 {
		goto L398
	} else {
		goto L474
	}
L474:
	;
	v2347 = v2009
	goto L363
L475:
	;
	v2016 = F_slice_del(m, l0)
	mBase = m.M
	v2017 = m.ExcPending
	if v2017 != 0 {
		goto L367
	} else {
		goto L476
	}
L476:
	;
	if v2016 < int32(0) {
		v2347 = v2016
		goto L363
	} else {
		goto L477
	}
L477:
	;
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2020
	v2022 = int32(2)
	v2024 = int32(0)
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2020-v2027 < v2022 {
		v2037 = v2024
		goto L479
	} else {
		goto L480
	}
L478:
	;
	if v2037 == int32(0) {
		goto L398
	} else {
		goto L482
	}
L479:
	;
	goto L478
L480:
	;
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2033 = F_memcmp(m, v2030+v2020-v2022, int32(_a_F_spanish_UTF_8_stem_23), v2022)
	mBase = m.M
	if v2033 != 0 {
		v2037 = v2024
		goto L479
	} else {
		goto L481
	}
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2020 - v2022
	v2037 = int32(1)
	goto L479
L482:
	;
	v2040 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2040
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v2042)))
	if v2040 < v2043 {
		goto L398
	} else {
		goto L483
	}
L483:
	;
	v2045 = F_slice_del(m, l0)
	mBase = m.M
	v2046 = m.ExcPending
	if v2046 != 0 {
		goto L367
	} else {
		goto L484
	}
L484:
	;
	if int32(0) <= v2045 {
		goto L398
	} else {
		goto L485
	}
L485:
	;
	v2347 = v2045
	goto L363
L486:
	;
	v2102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2102
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2105 = *(*int32)(unsafe.Add(mBase, uint32(v2104)+8))
	if v2102 < v2105 {
		goto L398
	} else {
		goto L506
	}
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2051
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2054
	v2061 = F_find_among_b(m, l0, int32(_a_F_spanish_UTF_8_stem_24), int32(12))
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L367
	} else {
		goto L488
	}
L488:
	;
	if v2061 == int32(0) {
		goto L489
	} else {
		goto L490
	}
L489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2057
	goto L486
L490:
	;
	goto L491
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2057
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2067
	if v2067 <= v2057 {
		goto L486
	} else {
		goto L492
	}
L492:
	;
	v2070 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2070+v2067-int32(1)))))
	if v2074 != int32(117) {
		goto L486
	} else {
		goto L493
	}
L493:
	;
	v2077 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2067 - v2077
	v2081 = F_slice_del(m, l0)
	mBase = m.M
	v2082 = m.ExcPending
	if v2082 != 0 {
		goto L367
	} else {
		goto L495
	}
L494:
	;
	v2093 = int32(0)
	v2094 = base.B2i32(v2088 < v2093)
	if v2094 == v2093 {
		goto L398
	} else {
		goto L502
	}
L495:
	;
	if int32(0) <= v2081 {
		goto L496
	} else {
		goto L497
	}
L496:
	;
	v2088 = v2077
	goto L498
L497:
	;
	v2088 = v2081 >> (uint(int32(31)) % 32) & v2081
	goto L498
L498:
	;
	if v2088 != 0 {
		goto L499
	} else {
		goto L500
	}
L499:
	;
	v2092 = int32(base.Ui32(v2088) >> (uint(int32(31)) % 32))
	goto L501
L500:
	;
	v2092 = int32(4)
	goto L501
L501:
	;
	switch v2092 {
	case 0:
		goto L398
	default:
		goto L494
	case 4:
		goto L486
	}
L502:
	;
	if v2088 < v2093 {
		goto L503
	} else {
		goto L504
	}
L503:
	;
	v2098 = v2088
	goto L505
L504:
	;
	v2098 = int32(1)
	goto L505
L505:
	;
	return v2098
L506:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2102
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2105
	v2112 = F_find_among_b(m, l0, int32(_a_F_spanish_UTF_8_stem_25), int32(96))
	mBase = m.M
	v2113 = m.ExcPending
	if v2113 != 0 {
		goto L367
	} else {
		goto L507
	}
L507:
	;
	if v2112 == int32(0) {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2108
	goto L398
L509:
	;
	goto L510
L510:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2108
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2118
	switch v2112 - int32(1) {
	case 0:
		goto L512
	case 1:
		goto L511
	default:
		goto L398
	}
L511:
	;
	v2149 = F_slice_del(m, l0)
	mBase = m.M
	v2150 = m.ExcPending
	if v2150 != 0 {
		goto L367
	} else {
		goto L522
	}
L512:
	;
	if v2118 <= v2108 {
		v2142 = v2118
		goto L513
	} else {
		goto L514
	}
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2142
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2142
	v2145 = F_slice_del(m, l0)
	mBase = m.M
	v2146 = m.ExcPending
	if v2146 != 0 {
		goto L367
	} else {
		goto L520
	}
L514:
	;
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2124 = v2123 + v2118
	v2127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2124-int32(1)))))
	if v2127 != int32(117) {
		v2142 = v2118
		goto L513
	} else {
		goto L515
	}
L515:
	;
	v2131 = v2118 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2131
	if v2131 <= v2108 {
		v2142 = v2118
		goto L513
	} else {
		goto L516
	}
L516:
	;
	v2136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2124-int32(2)))))
	if v2136 == int32(103) {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v2139 = v2131
	goto L519
L518:
	;
	v2139 = v2118
	goto L519
L519:
	;
	v2142 = v2139
	goto L513
L520:
	;
	if int32(0) <= v2145 {
		goto L398
	} else {
		goto L521
	}
L521:
	;
	v2347 = v2145
	goto L363
L522:
	;
	if v2149 < int32(0) {
		v2347 = v2149
		goto L363
	} else {
		goto L523
	}
L523:
	;
	goto L398
L524:
	;
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2216
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2220 = v2216
	v2222 = v2218
	goto L542
L525:
	;
	if v2162 == int32(0) {
		goto L524
	} else {
		goto L526
	}
L526:
	;
	v2166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2166
	switch v2162 - int32(1) {
	case 0:
		goto L528
	case 1:
		goto L527
	default:
		goto L524
	}
L527:
	;
	v2177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2178 = *(*int32)(unsafe.Add(mBase, uint32(v2177)+8))
	if v2166 < v2178 {
		goto L524
	} else {
		goto L532
	}
L528:
	;
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v2170)+8))
	if v2166 < v2171 {
		goto L524
	} else {
		goto L529
	}
L529:
	;
	v2173 = F_slice_del(m, l0)
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L367
	} else {
		goto L530
	}
L530:
	;
	if int32(0) <= v2173 {
		goto L524
	} else {
		goto L531
	}
L531:
	;
	v2347 = v2173
	goto L363
L532:
	;
	v2180 = F_slice_del(m, l0)
	mBase = m.M
	v2181 = m.ExcPending
	if v2181 != 0 {
		goto L367
	} else {
		goto L533
	}
L533:
	;
	if v2180 < int32(0) {
		v2347 = v2180
		goto L363
	} else {
		goto L534
	}
L534:
	;
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2184
	v2186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2184 <= v2186 {
		goto L524
	} else {
		goto L535
	}
L535:
	;
	v2188 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2189 = v2188 + v2184
	v2192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2189-int32(1)))))
	if v2192 != int32(117) {
		goto L524
	} else {
		goto L536
	}
L536:
	;
	v2196 = v2184 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2196
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2196
	if v2196 <= v2186 {
		goto L524
	} else {
		goto L537
	}
L537:
	;
	v2202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2189-int32(2)))))
	if v2202 != int32(103) {
		goto L524
	} else {
		goto L538
	}
L538:
	;
	v2205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2206 = *(*int32)(unsafe.Add(mBase, uint32(v2205)+8))
	if v2184 <= v2206 {
		goto L524
	} else {
		goto L539
	}
L539:
	;
	v2208 = F_slice_del(m, l0)
	mBase = m.M
	v2209 = m.ExcPending
	if v2209 != 0 {
		goto L367
	} else {
		goto L540
	}
L540:
	;
	if v2208 < int32(0) {
		v2347 = v2208
		goto L363
	} else {
		goto L541
	}
L541:
	;
	goto L524
L542:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2220
	v2226 = v2220 + int32(1)
	if v2222 <= v2226 {
		goto L548
	} else {
		goto L549
	}
L543:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2216
	v2347 = int32(1)
	goto L363
L544:
	;
	goto L543
L545:
	;
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2220 = v2343
	v2222 = v2342
	goto L542
L546:
	;
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L571
L547:
	;
	v2243 = F_find_among(m, l0, int32(_a_F_spanish_UTF_8_stem_26), int32(6))
	mBase = m.M
	v2244 = m.ExcPending
	if v2244 != 0 {
		goto L367
	} else {
		goto L552
	}
L548:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2220
	v2280 = v2220
	v2281 = v2222
	goto L546
L549:
	;
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2228+v2226))))
	if v2230&int32(224) != int32(160) {
		goto L548
	} else {
		goto L550
	}
L550:
	;
	if int32(1)<<(uint(v2230)%32)&int32(67641858) != 0 {
		goto L547
	} else {
		goto L551
	}
L551:
	;
	goto L548
L552:
	;
	v2245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2245
	switch v2243 - int32(1) {
	case 0:
		goto L558
	case 1:
		goto L557
	case 2:
		goto L556
	case 3:
		goto L555
	case 4:
		goto L554
	case 5:
		goto L553
	default:
		goto L545
	}
L553:
	;
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2280 = v2245
	v2281 = v2279
	goto L546
L554:
	;
	v2275 = F_slice_from_s(m, l0, int32(1), int32(_a_F_spanish_UTF_8_stem_27))
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		goto L367
	} else {
		goto L567
	}
L555:
	;
	v2269 = F_slice_from_s(m, l0, int32(1), int32(_a_F_spanish_UTF_8_stem_28))
	mBase = m.M
	v2270 = m.ExcPending
	if v2270 != 0 {
		goto L367
	} else {
		goto L565
	}
L556:
	;
	v2263 = F_slice_from_s(m, l0, int32(1), int32(_a_F_spanish_UTF_8_stem_29))
	mBase = m.M
	v2264 = m.ExcPending
	if v2264 != 0 {
		goto L367
	} else {
		goto L563
	}
L557:
	;
	v2257 = F_slice_from_s(m, l0, int32(1), int32(_a_F_spanish_UTF_8_stem_30))
	mBase = m.M
	v2258 = m.ExcPending
	if v2258 != 0 {
		goto L367
	} else {
		goto L561
	}
L558:
	;
	v2251 = F_slice_from_s(m, l0, int32(1), int32(_a_F_spanish_UTF_8_stem_31))
	mBase = m.M
	v2252 = m.ExcPending
	if v2252 != 0 {
		goto L367
	} else {
		goto L559
	}
L559:
	;
	if int32(0) <= v2251 {
		goto L545
	} else {
		goto L560
	}
L560:
	;
	v2347 = v2251
	goto L363
L561:
	;
	if int32(0) <= v2257 {
		goto L545
	} else {
		goto L562
	}
L562:
	;
	v2347 = v2257
	goto L363
L563:
	;
	if int32(0) <= v2263 {
		goto L545
	} else {
		goto L564
	}
L564:
	;
	v2347 = v2263
	goto L363
L565:
	;
	if int32(0) <= v2269 {
		goto L545
	} else {
		goto L566
	}
L566:
	;
	v2347 = v2269
	goto L363
L567:
	;
	if int32(0) <= v2275 {
		goto L545
	} else {
		goto L568
	}
L568:
	;
	v2347 = v2275
	goto L363
L569:
	;
	if v2335 < int32(0) {
		goto L544
	} else {
		goto L589
	}
L571:
	;
	goto L572
L572:
	;
	goto L573
L573:
	;
	v2290 = v2280
	v2292 = int32(1)
	goto L576
L575:
	;
	v2335 = v2320
	goto L569
L576:
	;
	if v2281 <= v2290 {
		goto L578
	} else {
		goto L579
	}
L577:
	;
	goto L575
L578:
	;
	v2335 = int32(-1)
	goto L569
L579:
	;
	goto L580
L580:
	;
	v2297 = v2290 + int32(1)
	v2299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2283+v2290))))
	if base.Ui32(v2299) < base.Ui32(int32(192)) {
		v2320 = v2297
		goto L581
	} else {
		goto L582
	}
L581:
	;
	v2321 = int32(1)
	if v2321 < v2292 {
		v2290 = v2320
		v2292 = v2292 - v2321
		goto L576
	} else {
		goto L588
	}
L582:
	;
	if v2281 <= v2297 {
		v2320 = v2297
		goto L581
	} else {
		goto L583
	}
L583:
	;
	v2306 = v2297
	goto L584
L584:
	;
	v2309 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2283+v2306))))
	if int32(-65) < v2309 {
		v2320 = v2306
		goto L581
	} else {
		goto L586
	}
L585:
	;
	v2320 = v2281
	goto L581
L586:
	;
	v2313 = v2306 + int32(1)
	if v2313 != v2281 {
		v2306 = v2313
		goto L584
	} else {
		goto L587
	}
L587:
	;
	goto L585
L588:
	;
	goto L577
L589:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2335
	goto L545
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
	F_spgWalk(m, v10, v4, int32(1), int32(257))
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
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
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int64
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
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
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+377)))
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
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	if v106 == int32(3) {
		goto L37
	} else {
		goto L38
	}
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+379)) = uint8(v97)
	v101 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+377)) = uint8(v101)
	goto L10
L24:
	;
	if v56 == int32(0) {
		v97 = v54
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
	if v60&int32(2) != 0 {
		v97 = v54
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
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+384))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
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
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v84 = m.T0[v83].(func(*base.Module, int32, int32, int32) int32)(m, v70, v67, v14+int32(15))
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
	if v84 != 0 {
		goto L22
	} else {
		goto L34
	}
L34:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+384))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+32))
	m.T0[v93].(func(*base.Module, int32, int32))(m, v91, v56)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L20
	} else {
		goto L35
	}
L35:
	;
	v97 = int32(1)
	goto L23
L36:
	;
	if base.B2i32(v28 == int64(-1))|base.B2i32(v117 <= v28) != 0 {
		goto L15
	} else {
		goto L41
	}
L37:
	;
	v109 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	v117 = v109
	goto L36
L38:
	;
	goto L39
L39:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	F_tuplestore_puttupleslot(m, v110, v56)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L20
	} else {
		goto L40
	}
L40:
	;
	v113 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	v115 = v113 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v115
	v117 = v115
	goto L36
L41:
	;
	goto L16
}
func F_std_fetch_func(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12+l1<<(uint(int32(2))%32))))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	if int32(0) < v18 {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+18)))
		if base.Ui32(v22&int32(2047)) < base.Ui32(v18) {
			v26 = F_getmissingattr(m, v17, v18, l2)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v100 = v26
				m.G0 = v10 + int32(16)
				return v100
			}
		} else {
			v30 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v30)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+20)))
			if v33&int32(1) == v30 {
				v38 = int32(4)
				v42 = v17 + v18<<(uint(v38)%32) + v38
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
				if int32(0) <= v43 {
					v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
					v48 = v32 + v46 + v43
					v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+6)))
					if v49 != int32(1) {
						v100 = v48
						m.G0 = v10 + int32(16)
						return v100
					} else {
						v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42)+4)))
						switch v52&int32(_a_F_std_fetch_func_0) - int32(1) {
						case 0:
							v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48))))
							v100 = v57
							m.G0 = v10 + int32(16)
							return v100
						case 1:
							v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48))))
							v100 = v58
							m.G0 = v10 + int32(16)
							return v100
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v52
								F_errmsg_internal(m, int32(_a_F_std_fetch_func_1), v10)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_std_fetch_func_2), int32(70), int32(_a_F_std_fetch_func_3))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 3:
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
							v100 = v59
							m.G0 = v10 + int32(16)
							return v100
						}
					}
				} else {
					v73 = F_nocachegetattr(m, v16, v18, v17)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						v100 = v73
						m.G0 = v10 + int32(16)
						return v100
					}
				}
			} else {
				v75 = int32(1)
				v76 = v18 - v75
				v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(base.Ui32(v76)>>(uint(int32(3))%32)))+23)))
				if int32(base.Ui32(v80)>>(uint(v76&int32(7))%32))&v75 == int32(0) {
					v88 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v88)
					v100 = int32(0)
					m.G0 = v10 + int32(16)
					return v100
				} else {
					v91 = F_nocachegetattr(m, v16, v18, v17)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						v100 = v91
						m.G0 = v10 + int32(16)
						return v100
					}
				}
			}
		}
	} else {
		v93 = F_heap_getsysattr(m, v16, v18, l2)
		mBase = m.M
		v94 = m.ExcPending
		if v94 != 0 {
			return int32(0)
		} else {
			v100 = v93
			m.G0 = v10 + int32(16)
			return v100
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
				v61 = int32(508)
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
						v61 = int32(508)
					} else {
						if v46 != 0 {
							v53 = int32(506)
						} else {
							v53 = int32(507)
						}
						if v44 != 0 {
							v55 = v53
						} else {
							v55 = int32(507)
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
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
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v4 {
		v68 = v4
		m.G0 = v10 + int32(16)
		return v68
	} else {
		if l2 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(34209924))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_str_toupper_0)
					F_errmsg(m, int32(_a_F_str_toupper_1), v10)
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						F_errhint(m, int32(_a_F_str_toupper_2), int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_str_toupper_3), int32(1719), int32(_a_F_str_toupper_4))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
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
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+3)))
				if v20 == int32(1) {
					v23 = F_pnstrdup(m, l0, l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
						if v25 == int32(0) {
							v68 = v23
						} else {
							v28 = v23
							v30 = v25
							for {
								v35 = int32(255)
								v36 = v30 & v35
								if base.Ui32((v36-int32(97))&v35) < base.Ui32(int32(26)) {
									v45 = v36 - int32(32)
								} else {
									v45 = v36
								}
								v47 = v45 & int32(255)
								*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v47)
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
								if v49 != 0 {
									v28 = v28 + int32(1)
									v30 = v49
									continue
								} else {
									break
								}
								break
							}
							v68 = v23
						}
						m.G0 = v10 + int32(16)
						return v68
					}
				} else {
					v53 = l1 + int32(1)
					v54 = F_palloc(m, v53)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v56 = F_pg_strupper(m, v54, v53, l0, l1, v16)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v59 = v56 + int32(1)
							if base.Ui32(v59) <= base.Ui32(v53) {
								v68 = v54
								m.G0 = v10 + int32(16)
								return v68
							} else {
								v61 = F_repalloc(m, v54, v59)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									v63 = F_pg_strupper(m, v61, v59, l0, l1, v16)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										v68 = v61
										m.G0 = v10 + int32(16)
										return v68
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
func F_strict_word_similarity_dist_commutator_op(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14021(m, l0, int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
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
func F_strlower_libc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v243 int32
	_ = v243
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_strlower_libc[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10*int32(28))+uint32(_c_F_strlower_libc[1])))
	goto L2
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L10
	} else {
		goto L120
	}
L2:
	;
	if int32(2) <= v15 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if l3 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	if l3 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L6:
	;
	v21 = F_strlen(m, l2)
	mBase = m.M
	v22 = v21
	goto L8
L7:
	;
	v22 = l3
	goto L8
L8:
	;
	v24 = v22 + int32(1)
	if base.Ui32(int32(536870912)) <= base.Ui32(v24) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v29 = F_palloc(m, v24<<(uint(int32(2))%32))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	F_char2wchar(m, v29, v24, l2, v22, l4)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v35 = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v40 = v35
	v41 = v37
	goto L16
L14:
	;
	v59 = v35
	goto L15
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_strlower_libc[0]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66*int32(28))+uint32(_c_F_strlower_libc[1])))
	goto L20
L16:
	;
	v49 = F_casemap(m, v41, int32(0))
	mBase = m.M
	goto L18
L17:
	;
	v59 = v52
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29+v40<<(uint(int32(2))%32)))) = v49
	v52 = v40 + int32(1)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v29+v52<<(uint(int32(2))%32))))
	if v56 != 0 {
		v40 = v52
		v41 = v56
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v74 = v71*v59 + int32(1)
	v75 = F_palloc(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	if v74 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_strlower_libc[2]))
	if v77 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v281 = v35
	goto L24
L24:
	;
	if base.Ui32(v281+int32(1)) <= base.Ui32(l1) {
		goto L88
	} else {
		goto L89
	}
L25:
	;
	v91 = int32(0)
	v97 = m.G0
	v98 = int32(16)
	v99 = v97 - v98
	m.G0 = v99
	*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = v29
	v103 = v99 + int32(12)
	v104 = m.G0
	v106 = v104 - v98
	m.G0 = v106
	if v75 != 0 {
		goto L40
	} else {
		goto L41
	}
L26:
	;
	if v77 == int32(-1) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v80 == int32(_a_F_strlower_libc_0) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v85 = int32(_a_F_strlower_libc_0)
	goto L31
L30:
	;
	v85 = v77
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strlower_libc[2])) = v85
	goto L28
L32:
	;
	v90 = int32(-1)
	goto L34
L33:
	;
	v90 = v80
	goto L34
L34:
	;
	goto L25
L35:
	;
	if v90 != 0 {
		goto L79
	} else {
		goto L80
	}
L36:
	;
	v262 = int32(16)
	m.G0 = v106 + v262
	m.G0 = v99 + v262
	goto L35
L37:
	;
	v258 = v74 - v243
	goto L36
L38:
	;
	if v180 != 0 {
		goto L63
	} else {
		goto L64
	}
L39:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v139 = v74
	v140 = v75
	v144 = v138
	goto L52
L40:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v74) {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if v111 == int32(0) {
		v258 = v91
		goto L36
	} else {
		goto L44
	}
L43:
	;
	v180 = v74
	v181 = v75
	goto L38
L44:
	;
	v114 = v111
	v115 = v110
	v117 = v91
	goto L45
L45:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v114) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v258 = v137
	goto L36
L47:
	;
	v126 = int32(-1)
	v129 = F_wcrtomb(m, v106+int32(12), v114)
	mBase = m.M
	if v129 == v126 {
		v258 = v126
		goto L36
	} else {
		goto L50
	}
L48:
	;
	v132 = int32(1)
	goto L49
L49:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	v137 = v132 + v117
	if v134 != 0 {
		v114 = v134
		v115 = v115 + int32(4)
		v117 = v137
		goto L45
	} else {
		goto L51
	}
L50:
	;
	v132 = v129
	goto L49
L51:
	;
	goto L46
L52:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	if base.Ui32(v148-int32(128)) <= base.Ui32(int32(-128)) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v180 = v170
	v181 = v173
	goto L38
L54:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v176 = v174 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v176
	if base.Ui32(int32(3)) < base.Ui32(v170) {
		v139 = v170
		v140 = v173
		v144 = v176
		goto L52
	} else {
		goto L62
	}
L55:
	;
	if v148 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v148)
	v166 = int32(1)
	v170 = v139 - v166
	v173 = v140 + v166
	goto L54
L58:
	;
	v155 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v155)
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v155
	v243 = v139
	goto L37
L59:
	;
	goto L60
L60:
	;
	v159 = int32(-1)
	v160 = F_wcrtomb(m, v140, v148)
	mBase = m.M
	if v160 == v159 {
		v258 = v159
		goto L36
	} else {
		goto L61
	}
L61:
	;
	v170 = v139 - v160
	v173 = v140 + v160
	goto L54
L62:
	;
	goto L53
L63:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v190 = v180
	v191 = v181
	v193 = v189
	goto L66
L64:
	;
	goto L65
L65:
	;
	v258 = v74
	goto L36
L66:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	if base.Ui32(v199-int32(128)) <= base.Ui32(int32(-128)) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L65
L68:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v232 = v230 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v232
	if v226 != 0 {
		v190 = v226
		v191 = v229
		v193 = v232
		goto L66
	} else {
		goto L77
	}
L69:
	;
	if v199 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v199)
	v222 = int32(1)
	v226 = v190 - v222
	v229 = v191 + v222
	goto L68
L72:
	;
	v206 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v206)
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v206
	v243 = v190
	goto L37
L73:
	;
	goto L74
L74:
	;
	v210 = int32(-1)
	v213 = F_wcrtomb(m, v106+int32(12), v199)
	mBase = m.M
	if v213 == v210 {
		v258 = v210
		goto L36
	} else {
		goto L75
	}
L75:
	;
	if base.Ui32(v190) < base.Ui32(v213) {
		v243 = v190
		goto L37
	} else {
		goto L76
	}
L76:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v218 = F_wcrtomb(m, v191, v217)
	mBase = m.M
	v226 = v190 - v213
	v229 = v191 + v213
	goto L68
L77:
	;
	goto L67
L78:
	;
	v281 = v258
	goto L24
L79:
	;
	if v90 == int32(-1) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	goto L85
L82:
	;
	v275 = int32(_a_F_strlower_libc_0)
	goto L84
L83:
	;
	v275 = v90
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strlower_libc[2])) = v275
	goto L81
L85:
	;
	goto L87
L87:
	;
	goto L78
L88:
	;
	if v281 != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	F_pfree(m, v29)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L10
	} else {
		goto L94
	}
L91:
	;
	base.MemoryCopy(m, l0, v75, v281)
	goto L93
L92:
	;
	goto L93
L93:
	;
	v287 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v281))) = uint8(v287)
	goto L90
L94:
	;
	F_pfree(m, v75)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L10
	} else {
		goto L95
	}
L95:
	;
	return v281
L96:
	;
	v296 = F_strlen(m, l2)
	mBase = m.M
	v297 = v296
	goto L98
L97:
	;
	v297 = l3
	goto L98
L98:
	;
	if base.Ui32(l1) < base.Ui32(v297+int32(1)) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	return v297
L100:
	;
	if v297 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	base.MemoryCopy(m, l0, l2, v297)
	goto L103
L102:
	;
	goto L103
L103:
	;
	v304 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v297))) = uint8(v304)
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v306 == v304 {
		goto L99
	} else {
		goto L104
	}
L104:
	;
	v309 = l0
	v311 = v306
	goto L105
L105:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v316 == int32(1) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	goto L99
L107:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v309))) = uint8(v339)
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+1)))
	if v341 != 0 {
		v309 = v309 + int32(1)
		v311 = v341
		goto L105
	} else {
		goto L119
	}
L108:
	;
	v319 = int32(255)
	v320 = v311 & v319
	if base.Ui32((v320-int32(65))&v319) < base.Ui32(int32(26)) {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	goto L110
L110:
	;
	v331 = v311 & int32(255)
	if base.Ui32(v331-int32(65)) < base.Ui32(int32(26)) {
		goto L116
	} else {
		goto L117
	}
L111:
	;
	v339 = v329
	goto L107
L112:
	;
	v329 = v320 | int32(32)
	goto L114
L113:
	;
	v329 = v320
	goto L114
L114:
	;
	goto L111
L115:
	;
	v339 = v338
	goto L107
L116:
	;
	v338 = v331 | int32(32)
	goto L118
L117:
	;
	v338 = v331
	goto L118
L118:
	;
	goto L115
L119:
	;
	goto L106
L120:
	;
	F_errcode(m, int32(_a_F_strlower_libc_1))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L10
	} else {
		goto L121
	}
L121:
	;
	F_errmsg(m, int32(_a_F_strlower_libc_2), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L10
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_strlower_libc_3), int32(207), int32(_a_F_strlower_libc_4))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L10
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v9 = m.G0
	v11 = v9 - int32(1024)
	m.G0 = v11
	v14 = l3 + int32(1)
	if v14 == int32(0) {
		v19 = F_strlen(m, l2)
		mBase = m.M
		if base.Ui32(v19) < base.Ui32(l1) {
			v21 = F_strcpy(m, l0, l2)
			mBase = m.M
		} else {
		}
		v41 = v19
		m.G0 = v11 + int32(1024)
		return v41
	} else {
		if base.Ui32(int32(1025)) <= base.Ui32(v14) {
			v24 = F_palloc(m, v14)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = v24
				if l3 != 0 {
					base.MemoryCopy(m, v28, l2, l3)
				} else {
				}
				v31 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3+v28))) = uint8(v31)
				v35 = F_strlen(m, v28)
				mBase = m.M
				if base.Ui32(v35) < base.Ui32(l1) {
					v37 = F_strcpy(m, l0, v28)
					mBase = m.M
				} else {
				}
				if v28 == v11 {
					v41 = v35
					m.G0 = v11 + int32(1024)
					return v41
				} else {
					F_pfree(m, v28)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						v41 = v35
						m.G0 = v11 + int32(1024)
						return v41
					}
				}
			}
		} else {
			v28 = v11
			if l3 != 0 {
				base.MemoryCopy(m, v28, l2, l3)
			} else {
			}
			v31 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3+v28))) = uint8(v31)
			v35 = F_strlen(m, v28)
			mBase = m.M
			if base.Ui32(v35) < base.Ui32(l1) {
				v37 = F_strcpy(m, l0, v28)
				mBase = m.M
			} else {
			}
			if v28 == v11 {
				v41 = v35
				m.G0 = v11 + int32(1024)
				return v41
			} else {
				F_pfree(m, v28)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v41 = v35
					m.G0 = v11 + int32(1024)
					return v41
				}
			}
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
func F_subarray(m *base.Module, l0 int32) int32 {
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
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v13 == int32(3) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v17 = v16
		} else {
			v17 = int32(0)
		}
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
		if v19 != 0 {
			v20 = F_array_contains_nulls(m, v9)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				if v20 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_subarray_0), int32(0))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_subarray_1), int32(287), int32(_a_F_subarray_2))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
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
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					v24 = v9 + int32(16)
					v25 = F_ArrayGetNItemsSafe(m, v22, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if v25 == int32(0) {
							v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v109 != v9 {
								F_pfree(m, v9)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									v114 = F_new_intArrayType(m, int32(0))
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										return v114
									}
								}
							} else {
								v114 = F_new_intArrayType(m, int32(0))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int32(0)
								} else {
									return v114
								}
							}
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							v30 = F_ArrayGetNItemsSafe(m, v29, v24)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
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
									v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v109 != v9 {
										F_pfree(m, v9)
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return int32(0)
										} else {
											v114 = F_new_intArrayType(m, int32(0))
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												return v114
											}
										}
									} else {
										v114 = F_new_intArrayType(m, int32(0))
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return int32(0)
										} else {
											return v114
										}
									}
								} else {
									v52 = v46 - v50
									v53 = F_new_intArrayType(m, v52)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
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
										v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										if v85 != v9 {
											F_pfree(m, v9)
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return int32(0)
											} else {
												return v53
											}
										} else {
											return v53
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
				return int32(0)
			} else {
				if v25 == int32(0) {
					v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v109 != v9 {
						F_pfree(m, v9)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							v114 = F_new_intArrayType(m, int32(0))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return int32(0)
							} else {
								return v114
							}
						}
					} else {
						v114 = F_new_intArrayType(m, int32(0))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int32(0)
						} else {
							return v114
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					v30 = F_ArrayGetNItemsSafe(m, v29, v24)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
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
							v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v109 != v9 {
								F_pfree(m, v9)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									v114 = F_new_intArrayType(m, int32(0))
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										return v114
									}
								}
							} else {
								v114 = F_new_intArrayType(m, int32(0))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int32(0)
								} else {
									return v114
								}
							}
						} else {
							v52 = v46 - v50
							v53 = F_new_intArrayType(m, v52)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
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
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v85 != v9 {
									F_pfree(m, v9)
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return int32(0)
									} else {
										return v53
									}
								} else {
									return v53
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_subpath(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v13 == int32(3) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v17 = v16
		} else {
			v17 = int32(0)
		}
		if v12 < int32(0) {
			v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+4)))
			v21 = v12 + v20
			v26 = v21 + v21>>(uint(int32(31))%32)&v20
		} else {
			v26 = v12
		}
		if v17 < int32(0) {
			v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+4)))
			v38 = v17 + v30
		} else {
			if v13 != int32(3) {
				v36 = int32(_a_F_subpath_0)
			} else {
				v36 = v26
			}
			if v17 != 0 {
				v37 = v26 + v17
			} else {
				v37 = v36
			}
			v38 = v37
		}
		v39 = F_inner_subltree(m, v8, v26, v38)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v41 != v8 {
				F_pfree(m, v8)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					return v39
				}
			} else {
				return v39
			}
		}
	}
}
func F_substitute_actual_srf_parameters_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == int32(0) {
		v54 = int32(0)
		m.G0 = v7 + int32(16)
		return v54
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v12 != int32(8) {
			if v12 != int32(67) {
				v52 = F_expression_tree_mutator_impl(m, l0, int32(875), l1)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v54 = v52
					m.G0 = v7 + int32(16)
					return v54
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v17 + int32(1)
				v23 = F_query_tree_mutator_impl(m, l0, int32(875), l1, int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v27 - int32(1)
					v54 = v23
					m.G0 = v7 + int32(16)
					return v54
				}
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v31 != 0 {
				v52 = F_expression_tree_mutator_impl(m, l0, int32(875), l1)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v54 = v52
					m.G0 = v7 + int32(16)
					return v54
				}
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v32 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v64
						F_errmsg_internal(m, int32(_a_F_substitute_actual_srf_parameters_mutator_0), v7)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_substitute_actual_srf_parameters_mutator_1), int32(_a_F_substitute_actual_srf_parameters_mutator_2), int32(_a_F_substitute_actual_srf_parameters_mutator_3))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v35 < v32 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v64
							F_errmsg_internal(m, int32(_a_F_substitute_actual_srf_parameters_mutator_0), v7)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_substitute_actual_srf_parameters_mutator_1), int32(_a_F_substitute_actual_srf_parameters_mutator_2), int32(_a_F_substitute_actual_srf_parameters_mutator_3))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v38+v32<<(uint(int32(2))%32)-int32(4))))
						v45 = F_copyObjectImpl(m, v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							F_IncrementVarSublevelsUp(m, v45, v47, int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v54 = v45
								m.G0 = v7 + int32(16)
								return v54
							}
						}
					}
				}
			}
		}
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_superuser_arg[0]))
	if base.B2i32(v5 == v2)|base.B2i32(l0 != v5) == v2 {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[1])))
		v55 = v13
		return v55 & int32(1)
	} else {
		if l0 == int32(10) {
			v16 = int32(1)
			v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[2])))
			if v18&v16 == int32(0) {
				v55 = v16
				return v55 & int32(1)
			} else {
				v26 = F_SearchSysCache1(m, int32(11), l0)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					if v26 != 0 {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v31)+68)))
						F_ReleaseCatCache(m, v26)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v36 = v33
							v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[3])))
							if v38 == int32(0) {
								F_CacheRegisterSyscacheCallback(m, int32(11), int32(1768), int32(0))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v47 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[3])) = uint8(v47)
									*(*int32)(unsafe.Add(mBase, _c_F_superuser_arg[0])) = l0
									v53 = v36 & int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[1])) = uint8(v53)
									v55 = v36
									return v55 & int32(1)
								}
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_superuser_arg[0])) = l0
								v53 = v36 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[1])) = uint8(v53)
								v55 = v36
								return v55 & int32(1)
							}
						}
					} else {
						v36 = int32(0)
						v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[3])))
						if v38 == int32(0) {
							F_CacheRegisterSyscacheCallback(m, int32(11), int32(1768), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								v47 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[3])) = uint8(v47)
								*(*int32)(unsafe.Add(mBase, _c_F_superuser_arg[0])) = l0
								v53 = v36 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[1])) = uint8(v53)
								v55 = v36
								return v55 & int32(1)
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_superuser_arg[0])) = l0
							v53 = v36 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[1])) = uint8(v53)
							v55 = v36
							return v55 & int32(1)
						}
					}
				}
			}
		} else {
			v26 = F_SearchSysCache1(m, int32(11), l0)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				if v26 != 0 {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v31)+68)))
					F_ReleaseCatCache(m, v26)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = v33
						v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[3])))
						if v38 == int32(0) {
							F_CacheRegisterSyscacheCallback(m, int32(11), int32(1768), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								v47 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[3])) = uint8(v47)
								*(*int32)(unsafe.Add(mBase, _c_F_superuser_arg[0])) = l0
								v53 = v36 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[1])) = uint8(v53)
								v55 = v36
								return v55 & int32(1)
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_superuser_arg[0])) = l0
							v53 = v36 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[1])) = uint8(v53)
							v55 = v36
							return v55 & int32(1)
						}
					}
				} else {
					v36 = int32(0)
					v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[3])))
					if v38 == int32(0) {
						F_CacheRegisterSyscacheCallback(m, int32(11), int32(1768), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v47 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[3])) = uint8(v47)
							*(*int32)(unsafe.Add(mBase, _c_F_superuser_arg[0])) = l0
							v53 = v36 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[1])) = uint8(v53)
							v55 = v36
							return v55 & int32(1)
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_superuser_arg[0])) = l0
						v53 = v36 & int32(1)
						*(*uint8)(unsafe.Add(mBase, _c_F_superuser_arg[1])) = uint8(v53)
						v55 = v36
						return v55 & int32(1)
					}
				}
			}
		}
	}
}
func F_suppress_redundant_updates_trigger(m *base.Module, l0 int32) int32 {
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
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
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
	v171 = m.ExcPending
	if v171 != 0 {
		goto L36
	} else {
		goto L49
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L36
	} else {
		goto L45
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L36
	} else {
		goto L41
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L36
	} else {
		goto L37
	}
L5:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v8 != int32(442) {
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
	return v115
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
	return int32(0)
L37:
	;
	F_errcode(m, int32(16908867))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(_a_F_suppress_redundant_updates_trigger_0), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_suppress_redundant_updates_trigger_1), int32(41), int32(_a_F_suppress_redundant_updates_trigger_2))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
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
	v142 = m.ExcPending
	if v142 != 0 {
		goto L36
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(_a_F_suppress_redundant_updates_trigger_3), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L36
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_suppress_redundant_updates_trigger_1), int32(47), int32(_a_F_suppress_redundant_updates_trigger_2))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
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
	v158 = m.ExcPending
	if v158 != 0 {
		goto L36
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(_a_F_suppress_redundant_updates_trigger_4), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L36
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_suppress_redundant_updates_trigger_1), int32(53), int32(_a_F_suppress_redundant_updates_trigger_2))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
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
	v174 = m.ExcPending
	if v174 != 0 {
		goto L36
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_suppress_redundant_updates_trigger_5), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L36
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_suppress_redundant_updates_trigger_1), int32(59), int32(_a_F_suppress_redundant_updates_trigger_2))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
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
