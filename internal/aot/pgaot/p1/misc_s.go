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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
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
	var v115 int32
	_ = v115
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
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	v10 = int32(1)
	v11 = l0 + v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v14 = v12 & v10
	if v12 == v10 {
		v17 = int32(4)
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
		if v19&int32(254) == int32(2) {
			v28 = v17
		} else {
			v28 = base.B2i32(v19 == int32(18)) << (uint(v17) % 32)
		}
		if v19 == int32(1) {
			v31 = v17
		} else {
			v31 = v28
		}
		v42 = v31
	} else {
		v32 = int32(1)
		if v14 != 0 {
			v42 = int32(base.Ui32(v12)>>(uint(v32)%32)) - v32
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v42 = int32(base.Ui32(v36)>>(uint(int32(2))%32)) - int32(4)
		}
	}
	v43 = int32(1)
	v44 = l1 + v43
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v45 == v43 {
		v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
		if base.Ui32((v48-int32(1))&int32(255)) < base.Ui32(int32(3)) {
			v159 = F_palloc(m, v42<<(uint(int32(1))%32)+int32(4))
			mBase = m.M
			v160 = m.ExcPending
			if v160 != 0 {
				return int32(0)
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v168 = m.ExcPending
				if v168 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(84410498))
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(316738), int32(0))
						mBase = m.M
						v175 = m.ExcPending
						if v175 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(572763), int32(0))
							mBase = m.M
							v179 = m.ExcPending
							if v179 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(478714), int32(438), int32(357417))
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
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
			v71 = base.B2i32(v48 == int32(18)) << (uint(int32(4)) % 32)
			if v14 != 0 {
				v74 = v11
			} else {
				v74 = l0 + int32(4)
			}
			v79 = F_palloc(m, v42<<(uint(int32(1))%32)+int32(4))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				v84 = v79 + int32(4)
				switch v71 {
				case 0:
					if v42 <= int32(0) {
						v248 = v84
					} else {
						if v42&int32(1) != 0 {
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
							v109 = v42 - v100
						} else {
							v107 = v84
							v108 = v74
							v109 = v42
						}
						if v42 == int32(1) {
							v248 = v107
						} else {
							v112 = v109
							v114 = v107
							v115 = v108
							for {
								v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
								if v121 == int32(92) {
									v124 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v124)
									v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
									v129 = v128
									v130 = v114 + int32(1)
								} else {
									v129 = v121
									v130 = v114
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v129)
								v133 = v115 + int32(1)
								v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
								if v134 != int32(92) {
									v144 = v134
									v145 = v130 + int32(1)
								} else {
									v139 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)) = uint8(v139)
									v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
									v144 = v141
									v145 = v130 + int32(2)
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v145))) = uint8(v144)
								v148 = v145 + int32(1)
								v149 = int32(2)
								if v149 < v112 {
									v112 = v112 - v149
									v114 = v148
									v115 = v115 + v149
									continue
								} else {
									break
								}
								break
							}
							v248 = v148
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v79))) = (v248 - v79) << (uint(int32(2)) % 32)
					return v79
				case 1:
					v185 = int32(1)
					v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
					v189 = v187 & v185
					if v189 != 0 {
						v190 = v185
					} else {
						v190 = int32(4)
					}
					v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v190))))
					if v192 == int32(92) {
						v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
						if v260 == int32(1) {
							v263 = int32(6)
							v265 = int32(18)
							v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
							if v267 == v265 {
								v270 = v265
							} else {
								v270 = int32(2)
							}
							if v267&int32(254) == int32(2) {
								v275 = v263
							} else {
								v275 = v270
							}
							if v267 == int32(1) {
								v278 = v263
							} else {
								v278 = v275
							}
							if v278 != 0 {
								v279 = F__emscripten_memcpy_bulkmem(m, v79, l0, v278)
								mBase = m.M
								v280 = v279
							} else {
								v280 = v79
							}
							return v280
						} else {
							if v260&int32(1) != 0 {
								v285 = int32(base.Ui32(v260) >> (uint(int32(1)) % 32))
								if v285 != 0 {
									v286 = F__emscripten_memcpy_bulkmem(m, v79, l0, v285)
									mBase = m.M
									v287 = v286
								} else {
									v287 = v79
								}
								return v287
							} else {
								v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v291 = int32(base.Ui32(v289) >> (uint(int32(2)) % 32))
								if v291 != 0 {
									v292 = F__emscripten_memcpy_bulkmem(m, v79, l0, v291)
									mBase = m.M
									v293 = v292
								} else {
									v293 = v79
								}
								return v293
							}
						}
					} else {
						if v42 <= int32(0) {
							v248 = v84
						} else {
							if v189 != 0 {
								v199 = v44
							} else {
								v199 = l1 + int32(4)
							}
							v203 = v84
							v204 = v74
							v205 = v42
							v207 = int32(1)
							for {
								v210 = int32(1)
								v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
								v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
								v215 = v207 ^ v210 | base.B2i32(v212 != v213)
								if v215&v210 == int32(0) {
									v220 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v203))) = uint8(v220)
									v239 = v203 + int32(1)
								} else {
									v225 = v203 + int32(1)
									if v212 == int32(92) {
										v228 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v203))) = uint8(v228)
										if v207&int32(1) == int32(0) {
											v239 = v225
										} else {
											v234 = int32(92)
											*(*uint8)(unsafe.Add(mBase, uint32(v203)+1)) = uint8(v234)
											v239 = v203 + int32(2)
										}
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(v203))) = uint8(v212)
										v239 = v225
									}
								}
								v240 = int32(1)
								if v240 < v205 {
									v203 = v239
									v204 = v204 + v240
									v205 = v205 - v240
									v207 = v215
									continue
								} else {
									break
								}
								break
							}
							v248 = v239
						}
						*(*int32)(unsafe.Add(mBase, uint32(v79))) = (v248 - v79) << (uint(int32(2)) % 32)
						return v79
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(84410498))
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(316738), int32(0))
							mBase = m.M
							v175 = m.ExcPending
							if v175 != 0 {
								return int32(0)
							} else {
								F_errhint(m, int32(572763), int32(0))
								mBase = m.M
								v179 = m.ExcPending
								if v179 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(478714), int32(438), int32(357417))
									mBase = m.M
									v184 = m.ExcPending
									if v184 != 0 {
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
		if v45&v59 != 0 {
			v71 = int32(base.Ui32(v45)>>(uint(v59)%32)) - v59
		} else {
			v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v71 = int32(base.Ui32(v65)>>(uint(int32(2))%32)) - int32(4)
		}
		if v14 != 0 {
			v74 = v11
		} else {
			v74 = l0 + int32(4)
		}
		v79 = F_palloc(m, v42<<(uint(int32(1))%32)+int32(4))
		mBase = m.M
		v82 = m.ExcPending
		if v82 != 0 {
			return int32(0)
		} else {
			v84 = v79 + int32(4)
			switch v71 {
			case 0:
				if v42 <= int32(0) {
					v248 = v84
				} else {
					if v42&int32(1) != 0 {
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
						v109 = v42 - v100
					} else {
						v107 = v84
						v108 = v74
						v109 = v42
					}
					if v42 == int32(1) {
						v248 = v107
					} else {
						v112 = v109
						v114 = v107
						v115 = v108
						for {
							v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
							if v121 == int32(92) {
								v124 = int32(92)
								*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v124)
								v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
								v129 = v128
								v130 = v114 + int32(1)
							} else {
								v129 = v121
								v130 = v114
							}
							*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v129)
							v133 = v115 + int32(1)
							v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
							if v134 != int32(92) {
								v144 = v134
								v145 = v130 + int32(1)
							} else {
								v139 = int32(92)
								*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)) = uint8(v139)
								v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
								v144 = v141
								v145 = v130 + int32(2)
							}
							*(*uint8)(unsafe.Add(mBase, uint32(v145))) = uint8(v144)
							v148 = v145 + int32(1)
							v149 = int32(2)
							if v149 < v112 {
								v112 = v112 - v149
								v114 = v148
								v115 = v115 + v149
								continue
							} else {
								break
							}
							break
						}
						v248 = v148
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(v79))) = (v248 - v79) << (uint(int32(2)) % 32)
				return v79
			case 1:
				v185 = int32(1)
				v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				v189 = v187 & v185
				if v189 != 0 {
					v190 = v185
				} else {
					v190 = int32(4)
				}
				v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v190))))
				if v192 == int32(92) {
					v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if v260 == int32(1) {
						v263 = int32(6)
						v265 = int32(18)
						v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
						if v267 == v265 {
							v270 = v265
						} else {
							v270 = int32(2)
						}
						if v267&int32(254) == int32(2) {
							v275 = v263
						} else {
							v275 = v270
						}
						if v267 == int32(1) {
							v278 = v263
						} else {
							v278 = v275
						}
						if v278 != 0 {
							v279 = F__emscripten_memcpy_bulkmem(m, v79, l0, v278)
							mBase = m.M
							v280 = v279
						} else {
							v280 = v79
						}
						return v280
					} else {
						if v260&int32(1) != 0 {
							v285 = int32(base.Ui32(v260) >> (uint(int32(1)) % 32))
							if v285 != 0 {
								v286 = F__emscripten_memcpy_bulkmem(m, v79, l0, v285)
								mBase = m.M
								v287 = v286
							} else {
								v287 = v79
							}
							return v287
						} else {
							v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v291 = int32(base.Ui32(v289) >> (uint(int32(2)) % 32))
							if v291 != 0 {
								v292 = F__emscripten_memcpy_bulkmem(m, v79, l0, v291)
								mBase = m.M
								v293 = v292
							} else {
								v293 = v79
							}
							return v293
						}
					}
				} else {
					if v42 <= int32(0) {
						v248 = v84
					} else {
						if v189 != 0 {
							v199 = v44
						} else {
							v199 = l1 + int32(4)
						}
						v203 = v84
						v204 = v74
						v205 = v42
						v207 = int32(1)
						for {
							v210 = int32(1)
							v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
							v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
							v215 = v207 ^ v210 | base.B2i32(v212 != v213)
							if v215&v210 == int32(0) {
								v220 = int32(92)
								*(*uint8)(unsafe.Add(mBase, uint32(v203))) = uint8(v220)
								v239 = v203 + int32(1)
							} else {
								v225 = v203 + int32(1)
								if v212 == int32(92) {
									v228 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v203))) = uint8(v228)
									if v207&int32(1) == int32(0) {
										v239 = v225
									} else {
										v234 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v203)+1)) = uint8(v234)
										v239 = v203 + int32(2)
									}
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v203))) = uint8(v212)
									v239 = v225
								}
							}
							v240 = int32(1)
							if v240 < v205 {
								v203 = v239
								v204 = v204 + v240
								v205 = v205 - v240
								v207 = v215
								continue
							} else {
								break
							}
							break
						}
						v248 = v239
					}
					*(*int32)(unsafe.Add(mBase, uint32(v79))) = (v248 - v79) << (uint(int32(2)) % 32)
					return v79
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v168 = m.ExcPending
				if v168 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(84410498))
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(316738), int32(0))
						mBase = m.M
						v175 = m.ExcPending
						if v175 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(572763), int32(0))
							mBase = m.M
							v179 = m.ExcPending
							if v179 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(478714), int32(438), int32(357417))
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
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
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
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
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
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
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int64
	_ = v268
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
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
	return v301
L2:
	;
	v220 = v217
	goto L59
L3:
	;
	v217 = int32(1)
	goto L2
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L11
	} else {
		goto L55
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = int64(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v22 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+8))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
	m.T0[v191].(func(*base.Module, int32))(m, v189)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L11
	} else {
		goto L52
	}
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v27 = v23 << (uint(int32(2)) % 32)
	goto L10
L9:
	;
	v27 = int32(0)
	goto L10
L10:
	;
	v28 = F_palloc(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v32 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v125 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+132)) = uint16(v125)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v128 != 0 {
		goto L32
	} else {
		goto L33
	}
L14:
	;
	v103 = int32(4449520)
	v104 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v106
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v85)+20))
	v111 = m.T0[v110].(func(*base.Module, int32, int32, int32) int32)(m, v85, v18, v13+int32(15))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L11
	} else {
		goto L29
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L11
	} else {
		goto L25
	}
L16:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v85 != 0 {
		goto L14
	} else {
		goto L24
	}
L17:
	;
	v35 = int32(0)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v36 <= v35 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v40 = v35
	goto L19
L19:
	;
	v50 = v40 << (uint(int32(2)) % 32)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50+v51)))
	v54 = int32(4449520)
	v55 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v57
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
	v62 = m.T0[v61].(func(*base.Module, int32, int32, int32) int32)(m, v53, v18, v13+int32(15))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L11
	} else {
		goto L21
	}
L20:
	;
	goto L16
L21:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v28+v50))) = v62
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	if v68 == int32(1) {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	v72 = v40 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v72 < v73 {
		v40 = v72
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v124 = v86
	goto L13
L25:
	;
	F_errcode(m, int32(403177602))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(290185), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(477709), int32(244), int32(95005))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v104
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	if v115 == int32(1) {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v120 = F_DirectFunctionCall1Coll(m, int32(747), int32(0), v111)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L11
	} else {
		goto L31
	}
L31:
	;
	v124 = v120
	goto L13
L32:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	v131 = v129
	goto L34
L33:
	;
	v131 = int32(0)
	goto L34
L34:
	;
	m.T0[v127].(func(*base.Module, int32, int32, int32, int32))(m, l0, v28, v131, v124)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v135 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	F_pfree(m, v28)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L11
	} else {
		goto L51
	}
L37:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+8))
	v141 = int32(0)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)))
	if v146 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v160 = int32(0)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+188))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	m.T0[v168].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v135, v160, int32(1), v162, base.B2i32(v134 == v160), v165)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L11
	} else {
		goto L50
	}
L40:
	;
	v147 = int32(68)
	goto L42
L41:
	;
	v147 = int32(4)
	goto L42
L42:
	;
	if v134 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v150 = v147
	goto L45
L44:
	;
	v150 = v147 | int32(128)
	goto L45
L45:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)))
	if v153 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v154 = v150 | int32(256)
	goto L48
L47:
	;
	v154 = v150
	goto L48
L48:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v138)+188))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
	v157 = m.T0[v156].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v138, v140, v141, v141, v141, v154)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v157
	goto L36
L50:
	;
	goto L36
L51:
	;
	v176 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)) = uint8(v176)
	goto L7
L52:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+153)))
	if v195 != 0 {
		v301 = int32(0)
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)))
	if v196&int32(1) != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v217 = int32(0)
	goto L2
L55:
	;
	F_errcode(m, int32(386400386))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(290222), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(477709), int32(256), int32(95005))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L11
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	if v220 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v220 = int32(0)
	goto L59
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L11
	} else {
		goto L86
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L11
	} else {
		goto L83
	}
L64:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	if v231 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	if v252 != 0 {
		goto L75
	} else {
		goto L76
	}
L67:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, _consts[115])))
	if v233&int32(1) == int32(0) {
		goto L63
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+188))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+172))
	v241 = m.T0[v240].(func(*base.Module, int32, int32) int32)(m, v188, l0)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L11
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	if v241 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v245 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+152)) = uint16(v245)
	v301 = int32(0)
	goto L1
L73:
	;
	goto L74
L74:
	;
	v248 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)) = uint8(v248)
	v220 = v248
	goto L59
L75:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, _consts[115])))
	if v254&int32(1) == int32(0) {
		goto L62
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+188))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+176))
	v262 = m.T0[v261].(func(*base.Module, int32, int32, int32) int32)(m, v188, l0, v189)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L11
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	if v262 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v266 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)) = uint8(v266)
	goto L61
L81:
	;
	goto L82
L82:
	;
	v268 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = v268 + int64(1)
	v301 = v189
	goto L1
L83:
	;
	F_errmsg_internal(m, int32(322544), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L11
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(313138), int32(1972), int32(303607))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L11
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errmsg_internal(m, int32(322682), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L11
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(313138), int32(1995), int32(367762))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L11
	} else {
		goto L88
	}
L88:
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
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
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
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v3 = int32(0)
	if l1 == v3 {
		v78 = v3
		v84 = int32(0)
	} else {
		v12 = int32(1)
		if l1 == v12 {
			v55 = l0
			v56 = l1
			v57 = v3
			v58 = v3
			v63 = int32(0)
		} else {
			v20 = l0
			v21 = int32(0)
			v22 = v3
			v24 = v3
			for {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
				v29 = int32(32)
				v30 = v28 | v29
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v33 = v31 | v29
				v34 = int32(17)
				v39 = v30 + (v33+v22*v34)*v34
				v40 = int32(257)
				v45 = (v21*v40+v33)*v40 + v30
				v46 = int32(2)
				v47 = v20 + v46
				v49 = v24 + v46
				if v49 != l1&int32(-2) {
					v20 = v47
					v21 = v45
					v22 = v39
					v24 = v49
					continue
				} else {
					break
				}
				break
			}
			v55 = v47
			v56 = v45
			v57 = v39
			v58 = v45 * int32(257)
			v63 = v39 * int32(17)
		}
		if l1&v12 != 0 {
			v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
			v66 = v64 | int32(32)
			v70 = v66 + v58
			v71 = v63 + v66
		} else {
			v70 = v56
			v71 = v57
		}
		v72 = int32(989)
		v73 = base.I32_rem_u_s(v71, v72)
		v75 = base.I32_rem_u_s(v70, v72)
		v78 = v73
		v84 = v75
	}
	v85 = int32(1)
	v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(v78<<(uint(v85)%32))+uint32(_consts[1249]))))
	v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84<<(uint(v85)%32))+uint32(_consts[1249]))))
	return v89 + v94
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
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v78 int64
	_ = v78
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v101 int64
	_ = v101
	var v108 int64
	_ = v108
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v124 int32
	_ = v124
	var v132 int64
	_ = v132
	var v135 int64
	_ = v135
	var v145 int64
	_ = v145
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
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
					v145 = int64(0)
				} else {
					if int32(10) < v24 {
						v45 = v12 + int32(8)
						v50 = base.B2i32(int32(2) < v24)
						if int32(2) < v24 {
							v51 = int32(4800)
						} else {
							v51 = int32(4799)
						}
						v52 = v51 + v31
						v57 = base.I32_div_s(v52, int32(4))
						v60 = base.I32_div_s(v52, int32(-100))
						v63 = base.I32_div_s(v52, int32(400))
						if int32(2) < v24 {
							v67 = int32(1)
						} else {
							v67 = int32(13)
						}
						v72 = base.I32_div_s((v67+v24)*int32(7834), int32(256))
						v78 = base.I64_extend_i32_s(v28 + v52*int32(365) + v57 + v60 + v63 + v72 - int32(32167) - int32(2451545))
						v87 = int64(32)
						v88 = int64(20)
						v90 = int64(base.Ui64(v78) >> (uint(v87) % 64))
						v93 = int64(4294967295)
						v94 = int64(500654080)
						v96 = v78 & v93
						v97 = v94 * v96
						v101 = int64(base.Ui64(v97)>>(uint(v87)%64)) + v94*v90
						v108 = v96*v88 + v101&v93
						*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v78*int64(0) + v78>>(uint(int64(63))%64)*int64(86400000000) + v88*v90 + int64(base.Ui64(v101)>>(uint(v87)%64)) + int64(base.Ui64(v108)>>(uint(v87)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v45))) = v97&v93 | v108<<(uint(v87)%64)
						v119 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
						if v119 != v120>>(uint(int64(63))%64) {
							v145 = int64(0)
						} else {
							v124 = int32(60)
							v132 = base.I64_extend_i32_s((v27*v124+v26)*v124+v25) * int64(1000000)
							v135 = v132 + v120
							if base.B2i32(v132 < int64(0))^base.B2i32(v135 < v120) != 0 {
								v145 = int64(0)
							} else {
								if base.Ui64(int64(9011559254509551615)) < base.Ui64(v135-int64(9223371331200000000)) {
									v145 = v135
								} else {
									v145 = int64(0)
								}
							}
						}
					} else {
						v145 = int64(0)
					}
				}
			} else {
				if v31 < int32(5874898) {
					v45 = v12 + int32(8)
					v50 = base.B2i32(int32(2) < v24)
					if int32(2) < v24 {
						v51 = int32(4800)
					} else {
						v51 = int32(4799)
					}
					v52 = v51 + v31
					v57 = base.I32_div_s(v52, int32(4))
					v60 = base.I32_div_s(v52, int32(-100))
					v63 = base.I32_div_s(v52, int32(400))
					if int32(2) < v24 {
						v67 = int32(1)
					} else {
						v67 = int32(13)
					}
					v72 = base.I32_div_s((v67+v24)*int32(7834), int32(256))
					v78 = base.I64_extend_i32_s(v28 + v52*int32(365) + v57 + v60 + v63 + v72 - int32(32167) - int32(2451545))
					v87 = int64(32)
					v88 = int64(20)
					v90 = int64(base.Ui64(v78) >> (uint(v87) % 64))
					v93 = int64(4294967295)
					v94 = int64(500654080)
					v96 = v78 & v93
					v97 = v94 * v96
					v101 = int64(base.Ui64(v97)>>(uint(v87)%64)) + v94*v90
					v108 = v96*v88 + v101&v93
					*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v78*int64(0) + v78>>(uint(int64(63))%64)*int64(86400000000) + v88*v90 + int64(base.Ui64(v101)>>(uint(v87)%64)) + int64(base.Ui64(v108)>>(uint(v87)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v45))) = v97&v93 | v108<<(uint(v87)%64)
					v119 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
					v120 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
					if v119 != v120>>(uint(int64(63))%64) {
						v145 = int64(0)
					} else {
						v124 = int32(60)
						v132 = base.I64_extend_i32_s((v27*v124+v26)*v124+v25) * int64(1000000)
						v135 = v132 + v120
						if base.B2i32(v132 < int64(0))^base.B2i32(v135 < v120) != 0 {
							v145 = int64(0)
						} else {
							if base.Ui64(int64(9011559254509551615)) < base.Ui64(v135-int64(9223371331200000000)) {
								v145 = v135
							} else {
								v145 = int64(0)
							}
						}
					}
				} else {
					if v31 != int32(5874898) {
						v145 = int64(0)
					} else {
						if int32(5) < v24 {
							v145 = int64(0)
						} else {
							v45 = v12 + int32(8)
							v50 = base.B2i32(int32(2) < v24)
							if int32(2) < v24 {
								v51 = int32(4800)
							} else {
								v51 = int32(4799)
							}
							v52 = v51 + v31
							v57 = base.I32_div_s(v52, int32(4))
							v60 = base.I32_div_s(v52, int32(-100))
							v63 = base.I32_div_s(v52, int32(400))
							if int32(2) < v24 {
								v67 = int32(1)
							} else {
								v67 = int32(13)
							}
							v72 = base.I32_div_s((v67+v24)*int32(7834), int32(256))
							v78 = base.I64_extend_i32_s(v28 + v52*int32(365) + v57 + v60 + v63 + v72 - int32(32167) - int32(2451545))
							v87 = int64(32)
							v88 = int64(20)
							v90 = int64(base.Ui64(v78) >> (uint(v87) % 64))
							v93 = int64(4294967295)
							v94 = int64(500654080)
							v96 = v78 & v93
							v97 = v94 * v96
							v101 = int64(base.Ui64(v97)>>(uint(v87)%64)) + v94*v90
							v108 = v96*v88 + v101&v93
							*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v78*int64(0) + v78>>(uint(int64(63))%64)*int64(86400000000) + v88*v90 + int64(base.Ui64(v101)>>(uint(v87)%64)) + int64(base.Ui64(v108)>>(uint(v87)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v45))) = v97&v93 | v108<<(uint(v87)%64)
							v119 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
							if v119 != v120>>(uint(int64(63))%64) {
								v145 = int64(0)
							} else {
								v124 = int32(60)
								v132 = base.I64_extend_i32_s((v27*v124+v26)*v124+v25) * int64(1000000)
								v135 = v132 + v120
								if base.B2i32(v132 < int64(0))^base.B2i32(v135 < v120) != 0 {
									v145 = int64(0)
								} else {
									if base.Ui64(int64(9011559254509551615)) < base.Ui64(v135-int64(9223371331200000000)) {
										v145 = v135
									} else {
										v145 = int64(0)
									}
								}
							}
						}
					}
				}
			}
			m.G0 = v12 + int32(32)
			return v145
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v154 = m.ExcPending
			if v154 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(281680), int32(0))
				mBase = m.M
				v158 = m.ExcPending
				if v158 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(476730), int32(2176), int32(361269))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
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
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v113 int64
	_ = v113
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_logicalrep_worker_attach(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = int32(913)
		v13 = m.G0
		v15 = v13 - int32(144)
		m.G0 = v15
		switch int32(915) {
		case 0, 2:
			v25 = v11
		default:
			*(*int32)(unsafe.Add(mBase, _consts[507])) = v11
			v25 = int32(4729)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v25
		F_sigemptyset(m, v15+int32(8))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = int32(268435456)
		v37 = v15 + int32(4)
		if v37 != 0 {
			v48 = F___memcpy(m, int32(4614572), v37, int32(140))
			mBase = m.M
		} else {
		}
		m.G0 = v15 + int32(144)
		v53 = int32(295)
		v55 = m.G0
		v57 = v55 - int32(144)
		m.G0 = v57
		switch int32(297) {
		case 0, 2:
			v67 = v53
		default:
			*(*int32)(unsafe.Add(mBase, _consts[510])) = v53
			v67 = int32(4729)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v67
		F_sigemptyset(m, v57+int32(8))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v57)+136)) = int32(268435456)
		v79 = v57 + int32(4)
		if v79 != 0 {
			v90 = F___memcpy(m, int32(4616532), v79, int32(140))
			mBase = m.M
		} else {
		}
		m.G0 = v57 + int32(144)
		F_BackgroundWorkerUnblockSignals(m)
		mBase = m.M
		v95 = m.ExcPending
		if v95 != 0 {
			return
		} else {
			v99 = m.G0
			v100 = int32(16)
			v101 = v99 - v100
			m.G0 = v101
			F___gettimeofday(m, v101)
			mBase = m.M
			v104 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
			v105 = int64(*(*int32)(unsafe.Add(mBase, uint32(v101)+8)))
			m.G0 = v101 + v100
			v113 = v105 + v104*int64(1000000) - int64(946684800000000)
			v115 = *(*int32)(unsafe.Add(mBase, _consts[648]))
			*(*int64)(unsafe.Add(mBase, uint32(v115)+88)) = v113
			*(*int64)(unsafe.Add(mBase, uint32(v115)+104)) = v113
			*(*int64)(unsafe.Add(mBase, uint32(v115)+80)) = v113
			F_load_file(m, int32(206012), int32(0))
			mBase = m.M
			v122 = m.ExcPending
			if v122 != 0 {
				return
			} else {
				F_InitializeLogRepWorker(m)
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return
				} else {
					v127 = F_errstart(m, int32(14), int32(0))
					mBase = m.M
					v128 = m.ExcPending
					if v128 != 0 {
						return
					} else {
						if v127 != 0 {
							v130 = *(*int32)(unsafe.Add(mBase, _consts[662]))
							v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v131
							F_errmsg_internal(m, int32(675157), v6)
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
								return
							} else {
								F_errfinish(m, int32(476327), int32(4805), int32(212168))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return
								} else {
									F_CacheRegisterSyscacheCallback(m, int32(68), int32(985), int32(0))
									mBase = m.M
									v145 = m.ExcPending
									if v145 != 0 {
										return
									} else {
										m.G0 = v6 + int32(16)
										return
									}
								}
							}
						} else {
							F_CacheRegisterSyscacheCallback(m, int32(68), int32(985), int32(0))
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
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
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
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
						v82 = v15
						m.G0 = v7 + int32(16)
						return v82
					} else {
						F_errcode(m, int32(16581))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(121375), int32(0))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(476110), int32(3971), int32(122300))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									v82 = v15
									m.G0 = v7 + int32(16)
									return v82
								}
							}
						}
					}
				} else {
					if v18 == int32(0) {
						v82 = v15
						m.G0 = v7 + int32(16)
						return v82
					} else {
						F_errmsg_internal(m, int32(120853), int32(0))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(476110), int32(3975), int32(122300))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v82 = v15
								m.G0 = v7 + int32(16)
								return v82
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
								v65 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(12))+uint32(_consts[619])))
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v65
								F_errmsg(m, int32(280983), v7)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(476110), int32(3986), int32(122300))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										if l0 != int32(13) {
											v82 = int32(0)
											m.G0 = v7 + int32(16)
											return v82
										} else {
											F_ExitPostmaster(m, int32(1))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
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
									v82 = int32(0)
									m.G0 = v7 + int32(16)
									return v82
								} else {
									F_ExitPostmaster(m, int32(1))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
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
					v82 = v9
					m.G0 = v7 + int32(16)
					return v82
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
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[765])))
	if v3 == int32(0) {
		v7 = int32(1)
		*(*int32)(unsafe.Add(mBase, _consts[48])) = v7
		*(*int32)(unsafe.Add(mBase, _consts[764])) = v7
	} else {
	}
	v13 = *(*int32)(unsafe.Add(mBase, _consts[516]))
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
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = l1
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = base.I64_extend_i32_s(v6 - v7)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l1 == int64(0) {
		v19 = v11
	} else {
		if base.I64_extend_i32_s(v11-v7) <= l1 {
			v19 = v11
		} else {
			v19 = v7 + base.I32_wrap_i64(l1)
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v19
	return
}
func F___strchrnul(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
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
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	v7 = l1 & int32(255)
	if v7 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v145
L2:
	;
	v135 = v130
	goto L37
L3:
	;
	v130 = v122
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
	if l0&int32(3) == int32(0) {
		v86 = l0
		goto L22
	} else {
		goto L23
	}
L7:
	;
	v12 = l0
	goto L10
L8:
	;
	v25 = l0
	goto L9
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v34 = int32(-2139062144)
	if (int32(16843008)-v31|v31)&v34 != v34 {
		v122 = v25
		goto L3
	} else {
		goto L15
	}
L10:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v17 == int32(0) {
		v145 = v12
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v25 = v22
	goto L9
L12:
	;
	if l1&int32(255) == v17 {
		v145 = v12
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v22 = v12 + int32(1)
	if v22&int32(3) != 0 {
		v12 = v22
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	v40 = v25
	v42 = v31
	goto L16
L16:
	;
	v46 = v42 ^ v7*int32(16843009)
	v49 = int32(-2139062144)
	if (int32(16843008)-v46|v46)&v49 != v49 {
		v122 = v40
		goto L3
	} else {
		goto L18
	}
L17:
	;
	v130 = v55
	goto L2
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v55 = v40 + int32(4)
	v59 = int32(-2139062144)
	if (v53|(int32(16843008)-v53))&v59 == v59 {
		v40 = v55
		v42 = v53
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	return v119 + l0
L21:
	;
	v119 = v111 - l0
	goto L20
L22:
	;
	v90 = v86
	goto L31
L23:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v70 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v119 = int32(0)
	goto L20
L25:
	;
	goto L26
L26:
	;
	v75 = l0
	goto L27
L27:
	;
	v79 = v75 + int32(1)
	if v79&int32(3) == int32(0) {
		v86 = v79
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v111 = v79
	goto L21
L29:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v84 != 0 {
		v75 = v79
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v99 = int32(-2139062144)
	if (int32(16843008)-v96|v96)&v99 == v99 {
		v90 = v90 + int32(4)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v105 = v90
	goto L34
L33:
	;
	goto L32
L34:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	if v109 != 0 {
		v105 = v105 + int32(1)
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v111 = v105
	goto L21
L36:
	;
	goto L35
L37:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v137 == int32(0) {
		v145 = v135
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v145 = v135
	goto L1
L39:
	;
	if v137 != l1&int32(255) {
		v135 = v135 + int32(1)
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
}
func F_sanitize_char_2(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(255)
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0 & v7
	if base.Ui32((l0-int32(33))&v7) < base.Ui32(int32(94)) {
		v20 = int32(652269)
	} else {
		v20 = int32(27929)
	}
	v21 = F_pg_snprintf(m, int32(4348272), int32(5), v20, v5)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return
	} else {
		m.G0 = v5 + int32(16)
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
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
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
						v45 = F_get_commutator(m, v14)
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
								v55 = F_scalarineqsel(m, v17, v52, v51, l2, v15, v12+int32(16), v41, v40)
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
						v52 = v14
						v55 = F_scalarineqsel(m, v17, v52, v51, l2, v15, v12+int32(16), v41, v40)
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
	if int32(16384) <= l3 {
		F___multf3(m, v8+int32(32), l1, l2, int64(0), int64(9222809086901354496))
		mBase = m.M
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v8)+40))
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v8)+32))
		if base.Ui32(l3) < base.Ui32(int32(32767)) {
			v62 = v18
			v63 = v17
			v64 = l3 - int32(16383)
		} else {
			F___multf3(m, v8+int32(16), v18, v17, int64(0), int64(9222809086901354496))
			mBase = m.M
			v28 = int32(49149)
			if base.Ui32(v28) <= base.Ui32(l3) {
				v31 = v28
			} else {
				v31 = l3
			}
			v34 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
			v35 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
			v62 = v35
			v63 = v34
			v64 = v31 - int32(32766)
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
				v64 = l3 + int32(16269)
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
				v64 = v57 + int32(32538)
			}
		}
	}
	F___multf3(m, v8, v62, v63, int64(0), base.I64_extend_i32_u(v64+int32(16383))<<(uint(int64(48))%64))
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
				F_errmsg(m, int32(62700), v8)
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
						F_errfinish(m, int32(301744), int32(1232), int32(201314))
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
				F_errmsg(m, int32(664525), v8+int32(16))
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
						F_errfinish(m, int32(301744), int32(1240), int32(201314))
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
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(12915)
			return int32(-1)
		} else {
			v31 = F_pg_cryptohash_init(m, v6)
			mBase = m.M
			if v31 < int32(0) {
				if v6 == int32(0) {
					v54 = int32(12915)
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
					if v46 == int32(1) {
						v49 = int32(292318)
					} else {
						v49 = int32(122318)
					}
					if v46 == int32(2) {
						v52 = int32(12915)
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
						v54 = int32(12915)
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
						if v46 == int32(1) {
							v49 = int32(292318)
						} else {
							v49 = int32(122318)
						}
						if v46 == int32(2) {
							v52 = int32(12915)
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
							v54 = int32(12915)
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
							if v46 == int32(1) {
								v49 = int32(292318)
							} else {
								v49 = int32(122318)
							}
							if v46 == int32(2) {
								v52 = int32(12915)
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
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v87 int64
	_ = v87
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v115 int64
	_ = v115
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v184 int32
	_ = v184
	v8 = int32(0)
	v17 = l1 << (uint(int64(32)) % 64)
	if l0 == l3 {
		v184 = l3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v184
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v17
	return
L2:
	;
	v19 = F_superuser_arg(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	if v19 != 0 {
		v184 = l3
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v22 = int32(0)
	v24 = F_roles_is_member_of(m, l0, int32(1), v22, v22)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = l0
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = int64(0)
	if v24 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v184 = v57
	goto L1
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L3
	} else {
		goto L37
	}
L9:
	;
	return
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v31 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v45 = v8
	v50 = v8
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
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v45<<(uint(int32(2))%32))))
	F_check_acl(m, l2)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	if v17 == int64(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v17 == v115 {
		goto L7
	} else {
		goto L32
	}
L17:
	;
	v115 = int64(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	if l3 == v57 {
		v184 = l3
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v64 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v74 = (v67<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L23
L22:
	;
	v74 = v64
	goto L23
L23:
	;
	v75 = int64(0)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(16))))
	if v76 <= int32(0) {
		v115 = v75
		goto L16
	} else {
		goto L24
	}
L24:
	;
	v81 = int32(0)
	v87 = v75
	goto L25
L25:
	;
	v98 = v74 + l2 + v81<<(uint(int32(4))%32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v57 == v99 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v115 = v105
	goto L16
L27:
	;
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v98)+8))
	v103 = v101&v17 | v87
	if v103 == v17 {
		goto L7
	} else {
		goto L30
	}
L28:
	;
	v105 = v87
	goto L29
L29:
	;
	v107 = v81 + int32(1)
	if v107 != v76 {
		v81 = v107
		v87 = v105
		goto L25
	} else {
		goto L31
	}
L30:
	;
	v105 = v103
	goto L29
L31:
	;
	goto L26
L32:
	;
	if v115 == int64(0) {
		v133 = v50
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v135 = v45 + int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v135 < v136 {
		v45 = v135
		v50 = v133
		goto L12
	} else {
		goto L36
	}
L34:
	;
	v128 = base.I32_wrap_i64(base.I64_popcnt(v115))
	if v128 <= v50 {
		v133 = v50
		goto L33
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v57
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v115
	v133 = v128
	goto L33
L36:
	;
	goto L13
L37:
	;
	F_errmsg_internal(m, int32(513855), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(478277), int32(1491), int32(103388))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[439]))
	if v11 == int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, _consts[182]))
		v17 = F_MemoryContextStrdup(m, v16, l1)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[439])) = v17
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+380))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+296))
			*(*int32)(unsafe.Add(mBase, _consts[440])) = v22
			v25 = int32(*(*uint8)(unsafe.Add(mBase, _consts[441])))
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
						v37 = *(*int32)(unsafe.Add(mBase, _consts[439]))
						v39 = *(*int32)(unsafe.Add(mBase, _consts[440]))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v39<<(uint(int32(2))%32))+uint32(_consts[442])))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+380))
						v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v44
						*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v46
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v37
						F_errmsg(m, int32(641006), v8)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							F_errfinish(m, int32(478556), int32(369), int32(419044))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
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
		v67 = m.ExcPending
		if v67 != 0 {
			return
		} else {
			F_errmsg(m, int32(398466), int32(0))
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
				v74 = *(*int32)(unsafe.Add(mBase, _consts[439]))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v74
				F_errdetail_log(m, int32(688161), v8+int32(16))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return
				} else {
					F_errfinish(m, int32(478556), int32(356), int32(419044))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	if l0 <= int32(0) {
		F_SetConfigOption(m, int32(161732), int32(401374), l1, l2)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			m.G0 = v7 + int32(80)
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
		v19 = F_pg_sprintf(m, v7+int32(16), int32(448014), v7)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			F_SetConfigOption(m, int32(161732), v7+int32(16), l1, l2)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				if l1 == int32(1) {
					F_SetConfigOption(m, int32(133472), int32(293030), int32(1), l2)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						F_SetConfigOption(m, int32(133398), int32(330316), int32(1), l2)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							if l0 == int32(1) {
								m.G0 = v7 + int32(80)
								return
							} else {
								F_SetConfigOption(m, int32(90412), int32(293030), l1, l2)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									if base.Ui32(l0) < base.Ui32(int32(3)) {
										m.G0 = v7 + int32(80)
										return
									} else {
										F_SetConfigOption(m, int32(346549), int32(330316), l1, l2)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return
										} else {
											if l0 == int32(3) {
												m.G0 = v7 + int32(80)
												return
											} else {
												F_SetConfigOption(m, int32(271203), int32(330316), l1, l2)
												mBase = m.M
												v55 = m.ExcPending
												if v55 != 0 {
													return
												} else {
													if base.Ui32(l0) < base.Ui32(int32(5)) {
														m.G0 = v7 + int32(80)
														return
													} else {
														F_SetConfigOption(m, int32(269323), int32(330316), l1, l2)
														mBase = m.M
														v61 = m.ExcPending
														if v61 != 0 {
															return
														} else {
															m.G0 = v7 + int32(80)
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
						m.G0 = v7 + int32(80)
						return
					} else {
						F_SetConfigOption(m, int32(90412), int32(293030), l1, l2)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							if base.Ui32(l0) < base.Ui32(int32(3)) {
								m.G0 = v7 + int32(80)
								return
							} else {
								F_SetConfigOption(m, int32(346549), int32(330316), l1, l2)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									if l0 == int32(3) {
										m.G0 = v7 + int32(80)
										return
									} else {
										F_SetConfigOption(m, int32(271203), int32(330316), l1, l2)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return
										} else {
											if base.Ui32(l0) < base.Ui32(int32(5)) {
												m.G0 = v7 + int32(80)
												return
											} else {
												F_SetConfigOption(m, int32(269323), int32(330316), l1, l2)
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return
												} else {
													m.G0 = v7 + int32(80)
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v99
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v104 != int32(333) {
		goto L30
	} else {
		goto L31
	}
L9:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v99 = v95
	goto L8
L10:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v91 != int32(5) {
		v99 = l1
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
	v99 = v41
	goto L8
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v99 = v32
	goto L8
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
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
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v48+v52<<(uint(int32(2))%32))))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v61 == int32(336) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L14
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+72))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v64 == v65 {
		v99 = v60
		goto L8
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v68 = v52 + int32(1)
	if v45 != v68 {
		v52 = v68
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
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v81
	F_errmsg_internal(m, int32(455444), v10)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(474851), int32(5246), int32(261554))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
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
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v99 = v94
	goto L8
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v115
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v119 - int32(342) {
	case 0, 13:
		v123 = int32(96)
		goto L35
	default:
		v126 = int32(0)
		goto L34
	case 12:
		goto L36
	}
L30:
	;
	v111 = int32(0)
	if v99 == v111 {
		v115 = v111
		goto L29
	} else {
		goto L33
	}
L31:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v107 != int32(3) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+156))
	v115 = v110
	goto L29
L33:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v99)+44))
	v115 = v114
	goto L29
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v126
	m.G0 = v10 + int32(16)
	return
L35:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1+v123)))
	v126 = v125
	goto L34
L36:
	;
	v123 = int32(104)
	goto L35
}
func F_set_plan_references(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
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
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	v3 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v14 = v13
	goto L3
L2:
	;
	v14 = v3
	goto L3
L3:
	;
	F_add_rtes_to_flat_rtable(m, l0, int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v20 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v80 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v23 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v31 = v3
	goto L9
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v31<<(uint(int32(2))%32))))
	v42 = F_palloc(m, int32(36))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	goto L6
L11:
	;
	v45 = v42 + int32(8)
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = v46
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	*(*int64)(unsafe.Add(mBase, uint32(v42))) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+32)) = v50
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v40)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = v52
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v40)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v42)+16)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = base.I32_wrap_i64(v46) + v14
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v59 + v14
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	v63 = F_lappend(m, v62, v42)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v63
	v67 = v31 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v67 < v68 {
		v31 = v67
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)))
	if v128 != 0 {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v83 = int32(0)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v84 <= v83 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v90 = v83
	goto L17
L17:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+v90<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+20)) = int32(0)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+4)) = v104 + v14
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+8)) = v107 + v14
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v111 = F_lappend(m, v110, v101)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L14
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v111
	v115 = v90 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v115 < v116 {
		v90 = v115
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v129 = int32(0)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v131 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v146 = F_set_plan_refs(m, l0, l1, v14)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L32
	}
L24:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v133 = v132
	goto L26
L25:
	;
	v133 = v129
	goto L26
L26:
	;
	v134 = F_palloc0(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = v134
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v137 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v139 = v138
	goto L30
L29:
	;
	v139 = v129
	goto L30
L30:
	;
	v140 = F_palloc0(m, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+364)) = v140
	goto L23
L32:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)))
	if v148 != int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	return v146
L34:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v151 == int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	if v154 <= int32(0) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v160 = int32(0)
	goto L37
L37:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+360))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168+v160))))
	if v170 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L33
L39:
	;
	v183 = v160 + int32(1)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	if v183 < v184 {
		v160 = v183
		goto L37
	} else {
		goto L42
	}
L40:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173+v160))))
	if v175 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v176+v160<<(uint(int32(2))%32)))) = int32(0)
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v101 int32
	_ = v101
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
	return v101
L12:
	;
	v31 = int32(0)
	v33 = v23
	goto L15
L13:
	;
	goto L14
L14:
	;
	v101 = int32(0)
	goto L11
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+124))
	v38 = v35 + v31*int32(36)
	v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+10)))
	v40 = int32(1)
	v41 = v39 - v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v42))))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v41))))
	if v47 == v40 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L14
L17:
	;
	v87 = v31 + int32(1)
	if v87 < v85 {
		v31 = v87
		v33 = v85
		goto L15
	} else {
		goto L37
	}
L18:
	;
	if v44&int32(1) != 0 {
		v85 = v33
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v44&int32(1) != 0 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+9)))
	if v54 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v55 = int32(-1)
	goto L24
L23:
	;
	v55 = int32(1)
	goto L24
L24:
	;
	return v55
L25:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+9)))
	if v61 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v65 = v41 << (uint(int32(2)) % 32)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65+v66)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69+v65)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v73 = m.T0[v72].(func(*base.Module, int32, int32, int32) int32)(m, v68, v71, v38)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L31
	}
L28:
	;
	v62 = int32(1)
	goto L30
L29:
	;
	v62 = int32(-1)
	goto L30
L30:
	;
	return v62
L31:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+8)))
	if v75 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v73 < int32(0) {
		goto L10
	} else {
		goto L35
	}
L33:
	;
	v82 = v73
	goto L34
L34:
	;
	if v82 != 0 {
		v101 = v82
		goto L11
	} else {
		goto L36
	}
L35:
	;
	v82 = int32(0) - v73
	goto L34
L36:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+120))
	v85 = v83
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
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
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
			F_errmsg(m, int32(194196), v5)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errfinish(m, int32(478058), int32(164), int32(313288))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
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
	v3 = *(*float64)(unsafe.Add(mBase, _consts[1434]))
	return base.I32_reinterpret_f32(base.F32_demote_f64(v3))
}
func F_show_unix_socket_permissions(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v7 = *(*int32)(unsafe.Add(mBase, _consts[429]))
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = v7
	v12 = F_pg_snprintf(m, int32(4346769), int32(12), int32(232448), v4)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		m.G0 = v4 + int32(16)
		return int32(4346769)
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
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
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
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
		v256 = v24
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v256
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
		v256 = v24
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
	v53 = v37
	v54 = v42
	v58 = int32(0)
	goto L13
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v61 = int32(1)
	v62 = v60 & v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v65 = v63 & v61
	if v62 != v65 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v256 = int32(0)
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
		v256 = v153
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
	if v83 == v75 {
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
	v250 = v58 + int32(1)
	if v250 != v26 {
		v53 = v53 + v244
		v54 = v54 + v244
		v58 = v250
		goto L13
	} else {
		goto L80
	}
L56:
	;
	v168 = int32(1)
	v170 = int32(4194302)
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
	v202 = int32(2)
	v203 = v188 + v202
	v204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v203))))
	v205 = int32(16383)
	v206 = v204 & v205
	v208 = v189 + v202
	v209 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v208))))
	v211 = v209 & v205
	if v206 != v211 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if base.Ui32(v221) < base.Ui32(v219) {
		goto L77
	} else {
		goto L78
	}
L67:
	;
	if base.Ui32(v211) < base.Ui32(v206) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	v218 = int32(14)
	v219 = int32(base.Ui32(v204) >> (uint(v218) % 32))
	v221 = int32(base.Ui32(v209) >> (uint(v218) % 32))
	if v219 == v221 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v216 = int32(-1)
	goto L72
L71:
	;
	v216 = int32(1)
	goto L72
L72:
	;
	return v216
L73:
	;
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
	v188 = v203
	v189 = v208
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
	v256 = v229
	goto L4
L80:
	;
	goto L14
}
func F_similar_escape(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v4 == int32(1) {
		v7 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v7)
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v12 = F_pg_detoast_datum_packed(m, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v16 != 0 {
				v20 = int32(0)
				v21 = F_similar_escape_internal(m, v12, v20)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					return v21
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v18 = F_pg_detoast_datum_packed(m, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = v18
					v21 = F_similar_escape_internal(m, v12, v20)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						return v21
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
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
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
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
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v227 int32
	_ = v227
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
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
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
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
	var v359 int32
	_ = v359
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
	v23 = int32(4)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v25&int32(254) == int32(2) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v38 = int32(1)
	if v20 != 0 {
		v48 = int32(base.Ui32(v18)>>(uint(v38)%32)) - v38
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v34 = v23
	goto L7
L6:
	;
	v34 = base.B2i32(v25 == int32(18)) << (uint(v23) % 32)
	goto L7
L7:
	;
	if v25 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = v23
	goto L10
L9:
	;
	v37 = v34
	goto L10
L10:
	;
	v48 = v37
	goto L1
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L31
	} else {
		goto L115
	}
L13:
	;
	v110 = F_palloc(m, v48*int32(3)+int32(27))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L31
	} else {
		goto L34
	}
L14:
	;
	v103 = int32(1)
	v104 = int32(488120)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v53 = int32(4)
	v54 = int32(1)
	v55 = l1 + v54
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v60 = v58 & v54
	if v60 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = v55
	goto L19
L18:
	;
	v61 = l1 + v53
	goto L19
L19:
	;
	if v58 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v94 = F_pg_mbstrlen_with_len(m, v61, v93)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L31
	} else {
		goto L32
	}
L21:
	;
	if v85 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if base.Ui32((v64-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v93 = v53
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
	if v60 != 0 {
		v85 = int32(base.Ui32(v58)>>(uint(v75)%32)) - v75
		goto L21
	} else {
		goto L26
	}
L25:
	;
	v85 = base.B2i32(v64 == int32(18)) << (uint(int32(4)) % 32)
	goto L21
L26:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v85 = int32(base.Ui32(v79)>>(uint(int32(2))%32)) - int32(4)
	goto L21
L27:
	;
	v88 = int32(0)
	v103 = v88
	v104 = v88
	goto L13
L28:
	;
	goto L29
L29:
	;
	if v85 < int32(2) {
		v103 = v85
		v104 = v61
		goto L13
	} else {
		goto L30
	}
L30:
	;
	v93 = v85
	goto L20
L31:
	;
	return int32(0)
L32:
	;
	if int32(2) <= v94 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	v103 = v93
	v104 = v61
	goto L13
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+4)) = int32(977217630)
	v115 = v110 + int32(8)
	if v48 <= int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v371 = int32(9257)
	*(*uint16)(unsafe.Add(mBase, uint32(v359))) = uint16(v371)
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = (v359-v110)<<(uint(int32(2))%32) + int32(8)
	return v110
L36:
	;
	v359 = v115
	goto L35
L37:
	;
	goto L38
L38:
	;
	if v20 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v120 = v17
	goto L41
L40:
	;
	v120 = l0 + int32(4)
	goto L41
L41:
	;
	v124 = int32(0)
	v127 = v120
	v128 = v115
	v131 = v124
	v132 = v48
	v135 = v124
	v136 = v3
	v137 = v3
	goto L42
L42:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v103 < int32(2) {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v359 = v346
	goto L35
L44:
	;
	if int32(0) < v349 {
		v127 = v353
		v128 = v346
		v131 = v348
		v132 = v349
		v135 = v350
		v136 = v351
		v137 = v352
		goto L42
	} else {
		goto L114
	}
L45:
	;
	v346 = v341
	v348 = v342
	v349 = v132 - v142
	v350 = v135
	v351 = v136
	v352 = v137
	v353 = v127 + v142
	goto L44
L46:
	;
	if v142 != 0 {
		goto L111
	} else {
		goto L112
	}
L47:
	;
	if v131&int32(1) != 0 {
		goto L80
	} else {
		goto L81
	}
L48:
	;
	v142 = F_pg_mblen_range(m, v127, v120+v48)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L31
	} else {
		goto L49
	}
L49:
	;
	if v142 < int32(2) {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	if v131&int32(1) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v148 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v148)
	v335 = v128 + int32(1)
	goto L46
L52:
	;
	goto L53
L53:
	;
	if v104 == int32(0) {
		v335 = v128
		goto L46
	} else {
		goto L54
	}
L54:
	;
	if v142 != v103 {
		v335 = v128
		goto L46
	} else {
		goto L55
	}
L55:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v103) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	if v216 != 0 {
		v335 = v128
		goto L46
	} else {
		goto L74
	}
L57:
	;
	v216 = int32(0)
	goto L56
L58:
	;
	v190 = v185
	v191 = v186
	v192 = v187
	goto L68
L59:
	;
	if (v104|v127)&int32(3) != 0 {
		v185 = v104
		v186 = v127
		v187 = v103
		goto L58
	} else {
		goto L62
	}
L60:
	;
	v178 = v104
	v179 = v127
	v180 = v103
	goto L61
L61:
	;
	if v180 == int32(0) {
		goto L57
	} else {
		goto L67
	}
L62:
	;
	v162 = v104
	v163 = v127
	v164 = v103
	goto L63
L63:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	if v167 != v168 {
		v185 = v162
		v186 = v163
		v187 = v164
		goto L58
	} else {
		goto L65
	}
L64:
	;
	v178 = v173
	v179 = v171
	v180 = v175
	goto L61
L65:
	;
	v170 = int32(4)
	v171 = v163 + v170
	v173 = v162 + v170
	v175 = v164 - v170
	if base.Ui32(int32(3)) < base.Ui32(v175) {
		v162 = v173
		v163 = v171
		v164 = v175
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v185 = v178
	v186 = v179
	v187 = v180
	goto L58
L68:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if v195 == v196 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v216 = v195 - v196
	goto L56
L70:
	;
	v198 = int32(1)
	v203 = v192 - v198
	if v203 != 0 {
		v190 = v190 + v198
		v191 = v191 + v198
		v192 = v203
		goto L68
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	goto L69
L73:
	;
	goto L57
L74:
	;
	v341 = v128
	v342 = int32(1)
	goto L45
L75:
	;
	v331 = int32(1)
	v346 = v326
	v348 = v328
	v349 = v132 - v331
	v350 = v327
	v351 = v329
	v352 = v330
	v353 = v127 + v331
	goto L44
L76:
	;
	v326 = v321
	v327 = v135
	v328 = int32(0)
	v329 = v323
	v330 = v324
	goto L75
L77:
	;
	switch v141 - int32(36) {
	case 0, 10, 56, 58:
		goto L105
	case 1:
		goto L108
	default:
		goto L104
	case 4:
		goto L106
	case 55:
		goto L109
	case 59:
		goto L107
	}
L78:
	;
	v276 = v128 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v141)
	if v141 != int32(93) {
		goto L98
	} else {
		goto L99
	}
L79:
	;
	v326 = v128 + int32(2)
	v327 = int32(3)
	v328 = int32(0)
	v329 = v136
	v330 = v137
	goto L75
L80:
	;
	if v141 != int32(34) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	if v104 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L83:
	;
	v255 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v255)
	*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)) = uint8(v141)
	goto L79
L84:
	;
	if int32(0) < v136 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	switch v137 {
	case 0:
		goto L86
	case 1:
		goto L88
	default:
		goto L87
	}
L86:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v128))) = int64(2900174335198198569)
	v321 = v128 + int32(8)
	v323 = v136
	v324 = v137 + int32(1)
	goto L76
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L31
	} else {
		goto L89
	}
L88:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v128))) = int64(4551025073606196009)
	v227 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v128)+8)) = uint8(v227)
	v321 = v128 + int32(9)
	v323 = v136
	v324 = v137 + int32(1)
	goto L76
L89:
	;
	F_errcode(m, int32(318767234))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L31
	} else {
		goto L90
	}
L90:
	;
	F_errmsg(m, int32(123587), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L31
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(476522), int32(951), int32(299436))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L31
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	if v136 <= int32(0) {
		goto L77
	} else {
		goto L96
	}
L94:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v141 != v261 {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v326 = v128
	v327 = v135
	v328 = int32(1)
	v329 = v136
	v330 = v137
	goto L75
L96:
	;
	if v141 != int32(92) {
		goto L78
	} else {
		goto L97
	}
L97:
	;
	v268 = int32(23644)
	*(*uint16)(unsafe.Add(mBase, uint32(v128))) = uint16(v268)
	goto L79
L98:
	;
	v284 = int32(3)
	v285 = int32(0)
	switch v141 - int32(91) {
	case 0:
		goto L102
	default:
		v326 = v276
		v327 = v284
		v328 = v285
		v329 = v136
		v330 = v137
		goto L75
	case 3:
		goto L101
	}
L99:
	;
	if v135 < int32(3) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v321 = v276
	v323 = v136 - int32(1)
	v324 = v137
	goto L76
L101:
	;
	v326 = v276
	v327 = v135 + int32(1)
	v328 = v285
	v329 = v136
	v330 = v137
	goto L75
L102:
	;
	v326 = v276
	v327 = v284
	v328 = v285
	v329 = v136 + int32(1)
	v330 = v137
	goto L75
L103:
	;
	v321 = v128 + int32(1)
	v323 = v136
	v324 = v137
	goto L76
L104:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v141)
	goto L103
L105:
	;
	v313 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v313)
	*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)) = uint8(v141)
	v321 = v128 + int32(2)
	v323 = v136
	v324 = v137
	goto L76
L106:
	;
	v307 = int32(16168)
	*(*uint16)(unsafe.Add(mBase, uint32(v128))) = uint16(v307)
	v309 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v128)+2)) = uint8(v309)
	v321 = v128 + int32(3)
	v323 = v136
	v324 = v137
	goto L76
L107:
	;
	v305 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v305)
	goto L103
L108:
	;
	v301 = int32(10798)
	*(*uint16)(unsafe.Add(mBase, uint32(v128))) = uint16(v301)
	v321 = v128 + int32(2)
	v323 = v136
	v324 = v137
	goto L76
L109:
	;
	v294 = int32(91)
	*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v294)
	v296 = int32(1)
	v326 = v128 + v296
	v327 = v296
	v328 = int32(0)
	v329 = v296
	v330 = v137
	goto L75
L110:
	;
	v341 = v337 + v142
	v342 = int32(0)
	goto L45
L111:
	;
	v336 = F__emscripten_memcpy_bulkmem(m, v335, v127, v142)
	mBase = m.M
	v337 = v336
	goto L113
L112:
	;
	v337 = v335
	goto L113
L113:
	;
	goto L110
L114:
	;
	goto L43
L115:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L31
	} else {
		goto L116
	}
L116:
	;
	F_errmsg(m, int32(316738), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L31
	} else {
		goto L117
	}
L117:
	;
	F_errhint(m, int32(572763), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L31
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(476522), int32(805), int32(299436))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L31
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
	var v90 int32
	_ = v90
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v7 != int32(1) {
		v90 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v90
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+144))
	if v10 != 0 {
		v90 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)))
	if v11 != 0 {
		v90 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v12 != 0 {
		v90 = v3
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+37)))
	if v13 != 0 {
		v90 = v3
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	if v14 != 0 {
		v90 = v3
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+42)))
	if v15 != 0 {
		v90 = v3
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	if v16 != 0 {
		v90 = v3
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	if v17 != 0 {
		v90 = v3
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	if v18 != 0 {
		v90 = v3
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
		v90 = v3
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
		v90 = v3
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
		v90 = int32(1)
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v90 = int32(1)
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
	v90 = v72
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
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
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
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
	v325 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v315))) = uint8(v325)
	return v320 - v14
L4:
	;
	v315 = v13
	v320 = v14
	goto L3
L5:
	;
	goto L6
L6:
	;
	v27 = v13
	v32 = v14
	v34 = v17
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
	v315 = v300
	v320 = v310
	goto L3
L9:
	;
	v310 = v304 + v32
	v311 = v303 + v34
	if int32(0) < v311 {
		v27 = v300
		v32 = v310
		v34 = v311
		goto L7
	} else {
		goto L73
	}
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)) = uint8(v37)
	v45 = int32(137)
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v45)
	v300 = v27 + int32(2)
	v303 = int32(-1)
	v304 = int32(1)
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
	v300 = v27 + int32(3)
	v303 = int32(-2)
	v304 = v295
	goto L9
L14:
	;
	if v34 == int32(1) {
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
		goto L68
	} else {
		goto L69
	}
L17:
	;
	v84 = v69 & int32(255)
	v87 = v84 | v37<<(uint(int32(8))%32)
	if base.Ui32(v87-int32(60736)) <= base.Ui32(int32(767)) {
		goto L28
	} else {
		goto L29
	}
L18:
	;
	if v12 != 0 {
		v315 = v27
		v320 = v32
		goto L3
	} else {
		goto L26
	}
L19:
	;
	if base.B2i32(v37 != int32(128))&base.B2i32(base.Ui32(v37) < base.Ui32(int32(160))) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if base.Ui32(int32(28)) < base.Ui32((v37+int32(32))&int32(255)) {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v69 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32)+1)))
	if v69 < int32(-3) {
		goto L17
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	if base.Ui32((v69+int32(-64))&int32(255)) < base.Ui32(int32(63)) {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	goto L18
L26:
	;
	F_report_invalid_encoding(m, int32(35), v32, v34)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
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
	v95 = v87
	v97 = v37
	v98 = v84
	v99 = int32(0)
	v100 = int32(61167)
	v102 = int32(2184032)
	goto L31
L29:
	;
	v124 = v87
	v126 = v37
	v127 = v84
	goto L30
L30:
	;
	if v124 <= int32(60222) {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	if v95 == v100 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v124 = v112
	v126 = v113
	v127 = v114
	goto L30
L33:
	;
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+2)))
	v112 = v107
	v113 = int32(base.Ui32(v107) >> (uint(int32(8)) % 32))
	v114 = v107 & int32(255)
	goto L35
L34:
	;
	v112 = v95
	v113 = v97
	v114 = v98
	goto L35
L35:
	;
	v116 = v99 + int32(1)
	v118 = v116 << (uint(int32(3)) % 32)
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v118)+uint32(_consts[1308]))))
	if v121 != int32(65535) {
		v95 = v112
		v97 = v113
		v98 = v114
		v99 = v116
		v100 = v121
		v102 = v118 + int32(2184032)
		goto L31
	} else {
		goto L36
	}
L36:
	;
	goto L32
L37:
	;
	v137 = int32(146)
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v137)
	v139 = int32(2)
	v143 = base.B2i32(int32(158) < v127)
	if int32(158) < v127 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v163 = int32(0)
	if base.B2i32(base.B2i32(v124 != int32(60223))&base.B2i32(base.Ui32(v124) < base.Ui32(int32(61504))) == v163)&base.B2i32(base.Ui32(int32(176)) < base.Ui32(v124-int32(64588))) == v163 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v144 = v139
	goto L42
L41:
	;
	v144 = int32(96)
	goto L42
L42:
	;
	v148 = v144 + v127 + base.B2i32(v127 < int32(128))
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+2)) = uint8(v148)
	v156 = v126<<(uint(int32(1))%32)&int32(126) | v143 + int32(159)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)) = uint8(v156)
	v295 = v139
	goto L13
L43:
	;
	v172 = int32(174)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+2)) = uint8(v172)
	v174 = int32(41618)
	*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v174)
	v300 = v27 + int32(3)
	v303 = int32(-2)
	v304 = int32(2)
	goto L9
L44:
	;
	goto L45
L45:
	;
	if base.Ui32(v124-int32(61504)) <= base.Ui32(int32(1279)) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v184 = int32(146)
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v184)
	v186 = int32(2)
	v190 = base.B2i32(int32(158) < v127)
	if int32(158) < v127 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	if base.Ui32(v124-int32(62784)) <= base.Ui32(int32(1279)) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v191 = v186
	goto L51
L50:
	;
	v191 = int32(96)
	goto L51
L51:
	;
	v195 = v191 + v127 + base.B2i32(v127 < int32(128))
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+2)) = uint8(v195)
	v205 = (v126<<(uint(int32(1))%32)+int32(34))&int32(126) | v190 + int32(243)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)) = uint8(v205)
	v295 = v186
	goto L13
L52:
	;
	v211 = int32(148)
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v211)
	v213 = int32(2)
	v217 = base.B2i32(int32(158) < v127)
	if int32(158) < v127 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v235 = int32(64064)
	v236 = int32(-2)
	v237 = int32(2)
	if base.Ui32(v124) < base.Ui32(v235) {
		v300 = v27
		v303 = v236
		v304 = v237
		goto L9
	} else {
		goto L58
	}
L55:
	;
	v218 = v213
	goto L57
L56:
	;
	v218 = int32(96)
	goto L57
L57:
	;
	v222 = v218 + v127 + base.B2i32(v127 < int32(128))
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+2)) = uint8(v222)
	v232 = (v126<<(uint(int32(1))%32)+int32(24))&int32(126) | v217 + int32(243)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)) = uint8(v232)
	v295 = v213
	goto L13
L58:
	;
	v242 = v124
	v243 = v27
	v244 = v235
	v245 = int32(2184032)
	v249 = int32(0)
	goto L59
L59:
	;
	if v244&int32(65535) == v242 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v300 = v274
	v303 = v236
	v304 = v237
	goto L9
L61:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	v257 = int32(128)
	v258 = v256 | v257
	*(*uint8)(unsafe.Add(mBase, uint32(v243)+2)) = uint8(v258)
	v263 = int32(base.Ui32(v256)>>(uint(int32(8))%32)) | v257
	*(*uint8)(unsafe.Add(mBase, uint32(v243)+1)) = uint8(v263)
	if int32(9371647) < v256 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v273 = v242
	v274 = v243
	goto L63
L63:
	;
	v276 = v249 + int32(1)
	v278 = v276 << (uint(int32(3)) % 32)
	v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v278)+uint32(_consts[1309]))))
	if v281 != int32(65535) {
		v242 = v273
		v243 = v274
		v244 = v281
		v245 = v278 + int32(2184032)
		v249 = v276
		goto L59
	} else {
		goto L67
	}
L64:
	;
	v269 = int32(-108)
	goto L66
L65:
	;
	v269 = int32(-110)
	goto L66
L66:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v243))) = uint8(v269)
	v273 = v256
	v274 = v243 + int32(3)
	goto L63
L67:
	;
	goto L60
L68:
	;
	if v12 != 0 {
		v315 = v27
		v320 = v32
		goto L3
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v37)
	v290 = int32(1)
	v300 = v27 + v290
	v303 = int32(-1)
	v304 = v290
	goto L9
L71:
	;
	F_report_invalid_encoding(m, int32(35), v32, v34)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[656]))
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
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v2 = int32(4444172)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v4 + int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9*int32(80))+uint32(_consts[821])))
	m.T0[v14].(func(*base.Module, int32, int32))(m, l0, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(-1)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v20*int32(80))+uint32(_consts[821])))
		m.T0[v25].(func(*base.Module, int32, int32))(m, l0, int32(1))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(-1)
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v31*int32(80))+uint32(_consts[821])))
			m.T0[v36].(func(*base.Module, int32, int32))(m, l0, int32(2))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v42*int32(80))+uint32(_consts[821])))
				m.T0[v47].(func(*base.Module, int32, int32))(m, l0, int32(3))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					v50 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v50
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v50
					v54 = int32(4444172)
					v56 = *(*int32)(unsafe.Add(mBase, _consts[163]))
					*(*int32)(unsafe.Add(mBase, _consts[163])) = v56 - int32(1)
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v3 = int32(4444172)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v5 + int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9*int32(80))+uint32(_consts[332])))
	v15 = m.T0[v14].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(4444172)
		v21 = *(*int32)(unsafe.Add(mBase, _consts[163]))
		*(*int32)(unsafe.Add(mBase, _consts[163])) = v21 - int32(1)
		return v15
	}
}
func F_smgrextend(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v6 = int32(4444172)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v8 + int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12*int32(80))+uint32(_consts[823])))
	m.T0[v17].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v24 = l0 + l1<<(uint(int32(2))%32) + int32(20)
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
		if v28 != l2 {
			v30 = int32(-1)
		} else {
			v30 = l2 + int32(1)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v24))) = v30
		v32 = int32(4444172)
		v34 = *(*int32)(unsafe.Add(mBase, _consts[163]))
		*(*int32)(unsafe.Add(mBase, _consts[163])) = v34 - int32(1)
		return
	}
}
func F_smgrreleaseall(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[822]))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = int32(4444172)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v11 + int32(1)
	F_hash_seq_init(m, v5+int32(12), v8)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v5 + int32(32)
	return
L4:
	;
	return
L5:
	;
	v21 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v23 = v21
	goto L10
L8:
	;
	goto L9
L9:
	;
	v89 = int32(4444172)
	v91 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v91 - int32(1)
	goto L3
L10:
	;
	v25 = int32(4444172)
	v27 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v27 + int32(1)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+36))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v32*int32(80))+uint32(_consts[821])))
	m.T0[v37].(func(*base.Module, int32, int32))(m, v23, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = int32(-1)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v23)+36))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43*int32(80))+uint32(_consts[821])))
	m.T0[v48].(func(*base.Module, int32, int32))(m, v23, int32(1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(-1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v23)+36))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54*int32(80))+uint32(_consts[821])))
	m.T0[v59].(func(*base.Module, int32, int32))(m, v23, int32(2))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = int32(-1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v23)+36))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65*int32(80))+uint32(_consts[821])))
	m.T0[v70].(func(*base.Module, int32, int32))(m, v23, int32(3))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v73 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v73
	v77 = int32(4444172)
	v79 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v79 - int32(1)
	v85 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if v85 != 0 {
		v23 = v85
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v55 int64
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int64
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int64
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v268 int32
	_ = v268
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v371 int32
	_ = v371
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v420 int32
	_ = v420
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v519 int64
	_ = v519
	var v521 int64
	_ = v521
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	v6 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v24 = m.G0
	v26 = v24 - int32(80)
	m.G0 = v26
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v28
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+40)) = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	if v32 == int32(-1) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v26 + int32(80)
	v519 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = v519
	v521 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = v521
	F_CacheInvalidateSmgr(m, v22)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L23
	} else {
		goto L85
	}
L2:
	;
	if v268 <= int32(0) {
		goto L1
	} else {
		goto L46
	}
L3:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	if base.Ui32(int32(63)) <= base.Ui32(v255+int32(31)) {
		goto L1
	} else {
		goto L45
	}
L4:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v268 = v253
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
	v136 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	if v32 != v136 {
		goto L1
	} else {
		goto L26
	}
L8:
	;
	v43 = v6
	v55 = int64(0)
	goto L9
L9:
	;
	v57 = v43 << (uint(int32(2)) % 32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1+v57)))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, _consts[113])))
	if v64 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	if base.I32_wrap_i64(v83) == int32(-1) {
		v268 = v89
		goto L2
	} else {
		goto L19
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57+(v26+int32(28))))) = v75
	if v75 == int32(-1) {
		goto L4
	} else {
		goto L17
	}
L12:
	;
	goto L11
L13:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0+v62<<(uint(int32(2))%32))+20))
	if v70 != int32(-1) {
		v75 = v70
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v75 = int32(-1)
	goto L12
L16:
	;
	goto L15
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l4+v57)))
	v83 = v55 + base.I64_extend_i32_u(v75-v80)
	v85 = v43 + int32(1)
	if v85 != l2 {
		v43 = v85
		v55 = v83
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	v94 = base.I32_div_s(v89, int32(32))
	if base.Ui64(base.I64_extend_i32_s(v94)) <= base.Ui64(v83) {
		v268 = v89
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v103 = int32(0)
	goto L21
L21:
	;
	v117 = v103 << (uint(int32(2)) % 32)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l4+v117)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1+v117)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v122
	v124 = *(*int64)(unsafe.Add(mBase, uint32(v26)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = v124
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(28)+v117)))
	F_FindAndDropRelationBuffers(m, v26, v121, v129, v119)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L1
L23:
	;
	return
L24:
	;
	v133 = v103 + int32(1)
	if v133 != l2 {
		v103 = v133
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	if l2 <= int32(0) {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v146 = v6
	goto L28
L28:
	;
	v160 = v146 << (uint(int32(2)) % 32)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l4+v160)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1+v160)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v165
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v26)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = v167
	v170 = v26 + int32(16)
	v172 = *(*int32)(unsafe.Add(mBase, _consts[742]))
	if int32(0) < v172 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L1
L30:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v170)+8))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v188 = int32(0)
	v193 = v176
	v197 = v172
	goto L33
L31:
	;
	goto L32
L32:
	;
	v250 = v146 + int32(1)
	if v250 != l2 {
		v146 = v250
		goto L28
	} else {
		goto L44
	}
L33:
	;
	v202 = v193 + v188<<(uint(int32(6))%32)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+24))
	if v203&int32(33554432) == int32(0) {
		v225 = v193
		v226 = v197
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v228 = v188 + int32(1)
	if v228 < v226 {
		v188 = v228
		v193 = v225
		v197 = v226
		goto L33
	} else {
		goto L43
	}
L36:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	if v208 != v179 {
		v225 = v193
		v226 = v197
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
	if v210 != v178 {
		v225 = v193
		v226 = v197
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	if v212 != v177 {
		v225 = v193
		v226 = v197
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v202)+12))
	if v214 != v164 {
		v225 = v193
		v226 = v197
		goto L35
	} else {
		goto L40
	}
L40:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v202)+16))
	if base.Ui32(v216) < base.Ui32(v162) {
		v225 = v193
		v226 = v197
		goto L35
	} else {
		goto L41
	}
L41:
	;
	F_InvalidateLocalBuffer(m, v202, int32(1))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L23
	} else {
		goto L42
	}
L42:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _consts[742]))
	v224 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v225 = v224
	v226 = v222
	goto L35
L43:
	;
	goto L34
L44:
	;
	goto L29
L45:
	;
	v268 = v255
	goto L2
L46:
	;
	v281 = int32(0)
	v291 = v281
	goto L47
L47:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v307 = v304 + v291<<(uint(int32(6))%32)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	if v308 != v309 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L1
L49:
	;
	v493 = v291 + int32(1)
	v495 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	if v493 < v495 {
		v291 = v493
		goto L47
	} else {
		goto L84
	}
L50:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v307)+4))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v26)+44))
	if v311 != v312 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v307)+8))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	if v314 != v315 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = int32(219672)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+72)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = int32(476033)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+56)) = int64(0)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v307)+24))
	v328 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v307)+24)) = v327 | v328
	if v327&v328 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	goto L56
L54:
	;
	v371 = v327
	goto L55
L55:
	;
	v384 = int32(4069244)
	v385 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(56))+8))
	if v387 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L56:
	;
	F_perform_spin_delay(m, v26+int32(56))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L23
	} else {
		goto L58
	}
L57:
	;
	v371 = v356
	goto L55
L58:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v307)+24))
	v357 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v307)+24)) = v356 | v357
	if v356&v357 != 0 {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	if base.B2i32(l2 <= v281) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L61:
	;
	goto L60
L62:
	;
	*(*int32)(unsafe.Add(mBase, _consts[276])) = v402
	goto L61
L63:
	;
	if int32(999) < v385 {
		goto L61
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v385 < int32(11) {
		goto L61
	} else {
		goto L70
	}
L66:
	;
	v392 = int32(900)
	if v392 <= v385 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v395 = v392
	goto L69
L68:
	;
	v395 = v385
	goto L69
L69:
	;
	v402 = v395 + int32(100)
	goto L62
L70:
	;
	v402 = v385 - int32(1)
	goto L62
L71:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v26)+44))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	v420 = int32(0)
	goto L74
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v307)+24)) = v371 & int32(-4194305)
	goto L49
L74:
	;
	if v409 != v410 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L73
L76:
	;
	v449 = v420 + int32(1)
	if v449 != l2 {
		v420 = v449
		goto L74
	} else {
		goto L83
	}
L77:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v307)+4))
	if v431 != v408 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v307)+8))
	if v433 != v407 {
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v307)+12))
	v437 = v420 << (uint(int32(2)) % 32)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l1+v437)))
	if v435 != v439 {
		goto L76
	} else {
		goto L80
	}
L80:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v307)+16))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l4+v437)))
	if base.Ui32(v441) < base.Ui32(v443) {
		goto L76
	} else {
		goto L81
	}
L81:
	;
	F_InvalidateBuffer(m, v307)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L23
	} else {
		goto L82
	}
L82:
	;
	goto L49
L83:
	;
	goto L75
L84:
	;
	goto L48
L85:
	;
	if int32(0) < l2 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v528 = l0 + int32(20)
	v546 = v6
	goto L89
L87:
	;
	goto L88
L88:
	;
	m.G0 = v22 + int32(16)
	return
L89:
	;
	v548 = int32(2)
	v549 = v546 << (uint(v548) % 32)
	v550 = l1 + v549
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
	*(*int32)(unsafe.Add(mBase, uint32(v528+v551<<(uint(v548)%32)))) = int32(-1)
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
	v558 = l3 + v549
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v558)))
	v560 = l4 + v549
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v562*int32(80))+uint32(_consts[824])))
	m.T0[v567].(func(*base.Module, int32, int32, int32, int32))(m, l0, v557, v559, v561)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L23
	} else {
		goto L91
	}
L90:
	;
	goto L88
L91:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v558)))
	if base.Ui32(v574) < base.Ui32(v575) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v577 = v574
	goto L94
L93:
	;
	v577 = v575
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v528+v570<<(uint(int32(2))%32)))) = v577
	v580 = v546 + int32(1)
	if v580 != l2 {
		v546 = v580
		goto L89
	} else {
		goto L95
	}
L95:
	;
	goto L90
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
	var v31 int32
	_ = v31
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
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
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
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
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
					v31 = int32(0)
					for {
						*(*int32)(unsafe.Add(mBase, uint32(v16+v31<<(uint(int32(2))%32)))) = v29
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
						if v44 != 0 {
							v29 = v44
							v31 = v31 + int32(1)
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				F_pg_qsort(m, v16, v10, int32(4), int32(981))
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
						v138 = int32(1)
					} else {
						v67 = int32(1)
						v68 = int32(3)
						if v10 <= v68 {
							v71 = v68
						} else {
							v71 = v10
						}
						if int32(4) <= v10 {
							v82 = int32(0)
							v83 = v67
							for {
								v90 = int32(2)
								v92 = v16 + v83<<(uint(v90)%32)
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
								v94 = int32(4)
								v95 = v92 + v94
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
								*(*int32)(unsafe.Add(mBase, uint32(v93)+16)) = v96
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v92-v94)))
								*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = v100
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
								v104 = v83 + v90
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v16+v104<<(uint(v90)%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v102)+16)) = v108
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
								*(*int32)(unsafe.Add(mBase, uint32(v102)+20)) = v110
								if v82 != v71&int32(2147483646)-int32(4) {
									v82 = v82 + v90
									v83 = v104
									continue
								} else {
									break
								}
								break
							}
							v117 = v104
						} else {
							v117 = v67
						}
						v125 = v71 - int32(1)
						if v71&int32(1) == int32(0) {
							v138 = v125
						} else {
							v130 = v16 + v117<<(uint(int32(2))%32)
							v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = v132
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v130-int32(4))))
							*(*int32)(unsafe.Add(mBase, uint32(v131)+20)) = v136
							v138 = v125
						}
					}
					v149 = v16 + v138<<(uint(int32(2))%32)
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
					*(*int32)(unsafe.Add(mBase, uint32(v150)+16)) = int32(0)
					v155 = *(*int32)(unsafe.Add(mBase, uint32(v149-int32(4))))
					*(*int32)(unsafe.Add(mBase, uint32(v150)+20)) = v155
					F_pfree(m, v16)
					mBase = m.M
					v158 = m.ExcPending
					if v158 != 0 {
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
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
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
	var v1673 int32
	_ = v1673
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1747 int32
	_ = v1747
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1759 int32
	_ = v1759
	var v1762 int32
	_ = v1762
	var v1766 int32
	_ = v1766
	var v1770 int32
	_ = v1770
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1787 int32
	_ = v1787
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1805 int32
	_ = v1805
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1812 int32
	_ = v1812
	var v1815 int32
	_ = v1815
	var v1818 int32
	_ = v1818
	var v1822 int32
	_ = v1822
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1868 int32
	_ = v1868
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1904 int32
	_ = v1904
	var v1906 int32
	_ = v1906
	var v1908 int32
	_ = v1908
	var v1911 int32
	_ = v1911
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1921 int32
	_ = v1921
	var v1924 int32
	_ = v1924
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1940 int32
	_ = v1940
	var v1942 int32
	_ = v1942
	var v1946 int32
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1975 int32
	_ = v1975
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1981 int32
	_ = v1981
	var v1983 int32
	_ = v1983
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2016 int32
	_ = v2016
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2023 int32
	_ = v2023
	var v2026 int32
	_ = v2026
	var v2029 int32
	_ = v2029
	var v2033 int32
	_ = v2033
	var v2036 int32
	_ = v2036
	var v2038 int32
	_ = v2038
	var v2039 int32
	_ = v2039
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2047 int32
	_ = v2047
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2053 int32
	_ = v2053
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2063 int32
	_ = v2063
	var v2066 int32
	_ = v2066
	var v2070 int32
	_ = v2070
	var v2073 int32
	_ = v2073
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2084 int32
	_ = v2084
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2094 int32
	_ = v2094
	var v2098 int32
	_ = v2098
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2104 int32
	_ = v2104
	var v2108 int32
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2114 int32
	_ = v2114
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2123 int32
	_ = v2123
	var v2127 int32
	_ = v2127
	var v2132 int32
	_ = v2132
	var v2135 int32
	_ = v2135
	var v2137 int32
	_ = v2137
	var v2141 int32
	_ = v2141
	var v2142 int32
	_ = v2142
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2153 int32
	_ = v2153
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2162 int32
	_ = v2162
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2180 int32
	_ = v2180
	var v2182 int32
	_ = v2182
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2188 int32
	_ = v2188
	var v2192 int32
	_ = v2192
	var v2198 int32
	_ = v2198
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2212 int32
	_ = v2212
	var v2214 int32
	_ = v2214
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2222 int32
	_ = v2222
	var v2224 int32
	_ = v2224
	var v2226 int32
	_ = v2226
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2241 int32
	_ = v2241
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2271 int32
	_ = v2271
	var v2272 int32
	_ = v2272
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2279 int32
	_ = v2279
	var v2286 int32
	_ = v2286
	var v2288 int32
	_ = v2288
	var v2293 int32
	_ = v2293
	var v2295 int32
	_ = v2295
	var v2302 int32
	_ = v2302
	var v2305 int32
	_ = v2305
	var v2309 int32
	_ = v2309
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2331 int32
	_ = v2331
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2343 int32
	_ = v2343
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
	*(*int32)(unsafe.Add(mBase, uint32(v1152)+8)) = v1150
	goto L1
L3:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1150 = v1148 + v1146
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
	v100 = v85&int32(63) | (v43<<(uint(int32(18))%32)&int32(1835008) | v52<<(uint(int32(12))%32) | v68<<(uint(int32(6))%32))
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
	v100 = v43<<(uint(int32(12))%32)&int32(61440) | v52<<(uint(int32(6))%32) | v68
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
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v105)>>(uint(int32(3))%32)))+uint32(_consts[1307]))))
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
	v219 = v204&int32(63) | (v162<<(uint(int32(18))%32)&int32(1835008) | v171<<(uint(int32(12))%32) | v187<<(uint(int32(6))%32))
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
	v219 = v162<<(uint(int32(12))%32)&int32(61440) | v171<<(uint(int32(6))%32) | v187
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
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v224)>>(uint(int32(3))%32)))+uint32(_consts[1307]))))
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
		v1146 = v367
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
	v338 = v323&int32(63) | (v281<<(uint(int32(18))%32)&int32(1835008) | v290<<(uint(int32(12))%32) | v306<<(uint(int32(6))%32))
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
	v338 = v281<<(uint(int32(12))%32)&int32(61440) | v290<<(uint(int32(6))%32) | v306
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
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v343)>>(uint(int32(3))%32)))+uint32(_consts[1307]))))
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
	v459 = v444&int32(63) | (v402<<(uint(int32(18))%32)&int32(1835008) | v411<<(uint(int32(12))%32) | v427<<(uint(int32(6))%32))
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
	v459 = v402<<(uint(int32(12))%32)&int32(61440) | v411<<(uint(int32(6))%32) | v427
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
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v464)>>(uint(int32(3))%32)))+uint32(_consts[1307]))))
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
		v1146 = v607
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
	v577 = v562&int32(63) | (v520<<(uint(int32(18))%32)&int32(1835008) | v529<<(uint(int32(12))%32) | v545<<(uint(int32(6))%32))
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
	v577 = v520<<(uint(int32(12))%32)&int32(61440) | v529<<(uint(int32(6))%32) | v545
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
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v582)>>(uint(int32(3))%32)))+uint32(_consts[1307]))))
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
	v700 = v685&int32(63) | (v643<<(uint(int32(18))%32)&int32(1835008) | v652<<(uint(int32(12))%32) | v668<<(uint(int32(6))%32))
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
	v700 = v643<<(uint(int32(12))%32)&int32(61440) | v652<<(uint(int32(6))%32) | v668
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
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v705)>>(uint(int32(3))%32)))+uint32(_consts[1307]))))
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
	v818 = v803&int32(63) | (v761<<(uint(int32(18))%32)&int32(1835008) | v770<<(uint(int32(12))%32) | v786<<(uint(int32(6))%32))
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
	v818 = v761<<(uint(int32(12))%32)&int32(61440) | v770<<(uint(int32(6))%32) | v786
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
	v829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v823)>>(uint(int32(3))%32)))+uint32(_consts[1307]))))
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
		v1146 = v966
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
	v937 = v922&int32(63) | (v880<<(uint(int32(18))%32)&int32(1835008) | v889<<(uint(int32(12))%32) | v905<<(uint(int32(6))%32))
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
	v937 = v880<<(uint(int32(12))%32)&int32(61440) | v889<<(uint(int32(6))%32) | v905
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
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v942)>>(uint(int32(3))%32)))+uint32(_consts[1307]))))
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
	v1058 = v1043&int32(63) | (v1001<<(uint(int32(18))%32)&int32(1835008) | v1010<<(uint(int32(12))%32) | v1026<<(uint(int32(6))%32))
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
	v1058 = v1001<<(uint(int32(12))%32)&int32(61440) | v1010<<(uint(int32(6))%32) | v1026
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
	v1069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1063)>>(uint(int32(3))%32)))+uint32(_consts[1307]))))
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
		v1150 = v1143
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
	v1244 = v1229&int32(63) | (v1187<<(uint(int32(18))%32)&int32(1835008) | v1196<<(uint(int32(12))%32) | v1212<<(uint(int32(6))%32))
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
	v1244 = v1187<<(uint(int32(12))%32)&int32(61440) | v1196<<(uint(int32(6))%32) | v1212
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
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1249)>>(uint(int32(3))%32)))+uint32(_consts[1307]))))
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
	v1366 = v1351&int32(63) | (v1309<<(uint(int32(18))%32)&int32(1835008) | v1318<<(uint(int32(12))%32) | v1334<<(uint(int32(6))%32))
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
	v1366 = v1309<<(uint(int32(12))%32)&int32(61440) | v1318<<(uint(int32(6))%32) | v1334
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
	v1377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1371)>>(uint(int32(3))%32)))+uint32(_consts[1307]))))
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
	v1491 = v1476&int32(63) | (v1434<<(uint(int32(18))%32)&int32(1835008) | v1443<<(uint(int32(12))%32) | v1459<<(uint(int32(6))%32))
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
	v1491 = v1434<<(uint(int32(12))%32)&int32(61440) | v1443<<(uint(int32(6))%32) | v1459
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
	v1502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1496)>>(uint(int32(3))%32)))+uint32(_consts[1307]))))
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
	v1613 = v1598&int32(63) | (v1556<<(uint(int32(18))%32)&int32(1835008) | v1565<<(uint(int32(12))%32) | v1581<<(uint(int32(6))%32))
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
	v1613 = v1556<<(uint(int32(12))%32)&int32(61440) | v1565<<(uint(int32(6))%32) | v1581
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
	v1624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1618)>>(uint(int32(3))%32)))+uint32(_consts[1307]))))
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
	return v2343
L364:
	;
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1759
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1759
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1759-int32(2) <= v1762 {
		goto L400
	} else {
		goto L401
	}
L365:
	;
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1658+v1656))))
	if v1660&int32(224) != int32(96) {
		goto L364
	} else {
		goto L366
	}
L366:
	;
	if int32(1)<<(uint(v1660)%32)&int32(557090) == int32(0) {
		goto L364
	} else {
		goto L367
	}
L367:
	;
	v1673 = F_find_among_b(m, l0, int32(4319024), int32(13))
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	return int32(0)
L369:
	;
	if v1673 == int32(0) {
		goto L364
	} else {
		goto L370
	}
L370:
	;
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1679
	v1682 = v1679 - int32(1)
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1682 <= v1683 {
		goto L364
	} else {
		goto L371
	}
L371:
	;
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1685+v1682))))
	switch v1687 - int32(111) {
	case 0, 3:
		goto L372
	default:
		goto L364
	}
L372:
	;
	v1692 = F_find_among_b(m, l0, int32(4319296), int32(11))
	mBase = m.M
	v1693 = m.ExcPending
	if v1693 != 0 {
		goto L368
	} else {
		goto L373
	}
L373:
	;
	if v1692 == int32(0) {
		goto L364
	} else {
		goto L374
	}
L374:
	;
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(v1697)+8))
	if v1696 < v1698 {
		goto L364
	} else {
		goto L375
	}
L375:
	;
	switch v1692 - int32(1) {
	case 0:
		goto L382
	case 1:
		goto L381
	case 2:
		goto L380
	case 3:
		goto L379
	case 4:
		goto L378
	case 5:
		goto L377
	case 6:
		goto L376
	default:
		goto L364
	}
L376:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1696 <= v1741 {
		goto L364
	} else {
		goto L395
	}
L377:
	;
	v1737 = F_slice_del(m, l0)
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L368
	} else {
		goto L393
	}
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1696
	v1733 = F_slice_from_s(m, l0, int32(2), int32(2176977))
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L368
	} else {
		goto L391
	}
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1696
	v1726 = F_slice_from_s(m, l0, int32(2), int32(2176975))
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L368
	} else {
		goto L389
	}
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1696
	v1719 = F_slice_from_s(m, l0, int32(2), int32(2176973))
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L368
	} else {
		goto L387
	}
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1696
	v1712 = F_slice_from_s(m, l0, int32(4), int32(2176969))
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L368
	} else {
		goto L385
	}
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1696
	v1705 = F_slice_from_s(m, l0, int32(5), int32(2176964))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L368
	} else {
		goto L383
	}
L383:
	;
	if int32(0) <= v1705 {
		goto L364
	} else {
		goto L384
	}
L384:
	;
	v2343 = v1705
	goto L363
L385:
	;
	if int32(0) <= v1712 {
		goto L364
	} else {
		goto L386
	}
L386:
	;
	v2343 = v1712
	goto L363
L387:
	;
	if int32(0) <= v1719 {
		goto L364
	} else {
		goto L388
	}
L388:
	;
	v2343 = v1719
	goto L363
L389:
	;
	if int32(0) <= v1726 {
		goto L364
	} else {
		goto L390
	}
L390:
	;
	v2343 = v1726
	goto L363
L391:
	;
	if int32(0) <= v1733 {
		goto L364
	} else {
		goto L392
	}
L392:
	;
	v2343 = v1733
	goto L363
L393:
	;
	if int32(0) <= v1737 {
		goto L364
	} else {
		goto L394
	}
L394:
	;
	v2343 = v1737
	goto L363
L395:
	;
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1743+v1696-int32(1)))))
	if v1747 != int32(117) {
		goto L364
	} else {
		goto L396
	}
L396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1696 - int32(1)
	v1753 = F_slice_del(m, l0)
	mBase = m.M
	v1754 = m.ExcPending
	if v1754 != 0 {
		goto L368
	} else {
		goto L397
	}
L397:
	;
	if v1753 < int32(0) {
		v2343 = v1753
		goto L363
	} else {
		goto L398
	}
L398:
	;
	goto L364
L399:
	;
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2153
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2153
	v2158 = F_find_among_b(m, l0, int32(4322816), int32(8))
	mBase = m.M
	v2159 = m.ExcPending
	if v2159 != 0 {
		goto L368
	} else {
		goto L529
	}
L400:
	;
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2047
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2050 = *(*int32)(unsafe.Add(mBase, uint32(v2049)+8))
	if v2047 < v2050 {
		goto L490
	} else {
		goto L491
	}
L401:
	;
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1766+v1759-int32(1)))))
	if v1770&int32(224) != int32(96) {
		goto L400
	} else {
		goto L402
	}
L402:
	;
	if int32(1)<<(uint(v1770)%32)&int32(835634) == int32(0) {
		goto L400
	} else {
		goto L403
	}
L403:
	;
	v1783 = F_find_among_b(m, l0, int32(4319520), int32(46))
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		goto L368
	} else {
		goto L404
	}
L404:
	;
	if v1783 == int32(0) {
		goto L400
	} else {
		goto L405
	}
L405:
	;
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1787
	switch v1783 - int32(1) {
	case 0:
		goto L414
	case 1:
		goto L413
	case 2:
		goto L412
	case 3:
		goto L411
	case 4:
		goto L410
	case 5:
		goto L409
	case 6:
		goto L408
	case 7:
		goto L407
	case 8:
		goto L406
	default:
		goto L399
	}
L406:
	;
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(v2009)))
	if v1787 < v2010 {
		goto L400
	} else {
		goto L479
	}
L407:
	;
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(v1968)))
	if v1787 < v1969 {
		goto L400
	} else {
		goto L468
	}
L408:
	;
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1933)))
	if v1787 < v1934 {
		goto L400
	} else {
		goto L458
	}
L409:
	;
	v1861 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1862 = *(*int32)(unsafe.Add(mBase, uint32(v1861)+4))
	if v1787 < v1862 {
		goto L400
	} else {
		goto L438
	}
L410:
	;
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v1852)))
	if v1787 < v1853 {
		goto L400
	} else {
		goto L435
	}
L411:
	;
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(v1843)))
	if v1787 < v1844 {
		goto L400
	} else {
		goto L432
	}
L412:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v1834)))
	if v1787 < v1835 {
		goto L400
	} else {
		goto L429
	}
L413:
	;
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1799 = *(*int32)(unsafe.Add(mBase, uint32(v1798)))
	if v1787 < v1799 {
		goto L400
	} else {
		goto L418
	}
L414:
	;
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v1791)))
	if v1787 < v1792 {
		goto L400
	} else {
		goto L415
	}
L415:
	;
	v1794 = F_slice_del(m, l0)
	mBase = m.M
	v1795 = m.ExcPending
	if v1795 != 0 {
		goto L368
	} else {
		goto L416
	}
L416:
	;
	if int32(0) <= v1794 {
		goto L399
	} else {
		goto L417
	}
L417:
	;
	v2343 = v1794
	goto L363
L418:
	;
	v1801 = F_slice_del(m, l0)
	mBase = m.M
	v1802 = m.ExcPending
	if v1802 != 0 {
		goto L368
	} else {
		goto L419
	}
L419:
	;
	if v1801 < int32(0) {
		v2343 = v1801
		goto L363
	} else {
		goto L420
	}
L420:
	;
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1805
	v1807 = int32(2)
	v1809 = int32(0)
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1805-v1812 < v1807 {
		v1822 = v1809
		goto L422
	} else {
		goto L423
	}
L421:
	;
	if v1822 == int32(0) {
		goto L399
	} else {
		goto L425
	}
L422:
	;
	goto L421
L423:
	;
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1818 = F_memcmp(m, v1815+v1805-v1807, int32(2177059), v1807)
	mBase = m.M
	if v1818 != 0 {
		v1822 = v1809
		goto L422
	} else {
		goto L424
	}
L424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1805 - v1807
	v1822 = int32(1)
	goto L422
L425:
	;
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1825
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(v1827)))
	if v1825 < v1828 {
		goto L399
	} else {
		goto L426
	}
L426:
	;
	v1830 = F_slice_del(m, l0)
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L368
	} else {
		goto L427
	}
L427:
	;
	if int32(0) <= v1830 {
		goto L399
	} else {
		goto L428
	}
L428:
	;
	v2343 = v1830
	goto L363
L429:
	;
	v1839 = F_slice_from_s(m, l0, int32(3), int32(2177061))
	mBase = m.M
	v1840 = m.ExcPending
	if v1840 != 0 {
		goto L368
	} else {
		goto L430
	}
L430:
	;
	if int32(0) <= v1839 {
		goto L399
	} else {
		goto L431
	}
L431:
	;
	v2343 = v1839
	goto L363
L432:
	;
	v1848 = F_slice_from_s(m, l0, int32(1), int32(2177064))
	mBase = m.M
	v1849 = m.ExcPending
	if v1849 != 0 {
		goto L368
	} else {
		goto L433
	}
L433:
	;
	if int32(0) <= v1848 {
		goto L399
	} else {
		goto L434
	}
L434:
	;
	v2343 = v1848
	goto L363
L435:
	;
	v1857 = F_slice_from_s(m, l0, int32(4), int32(2177065))
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L368
	} else {
		goto L436
	}
L436:
	;
	if int32(0) <= v1857 {
		goto L399
	} else {
		goto L437
	}
L437:
	;
	v2343 = v1857
	goto L363
L438:
	;
	v1864 = F_slice_del(m, l0)
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L368
	} else {
		goto L439
	}
L439:
	;
	if v1864 < int32(0) {
		v2343 = v1864
		goto L363
	} else {
		goto L440
	}
L440:
	;
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1868
	v1871 = v1868 - int32(1)
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1871 <= v1872 {
		goto L399
	} else {
		goto L441
	}
L441:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1874+v1871))))
	if v1876&int32(224) != int32(96) {
		goto L399
	} else {
		goto L442
	}
L442:
	;
	if int32(1)<<(uint(v1876)%32)&int32(4718616) == int32(0) {
		goto L399
	} else {
		goto L443
	}
L443:
	;
	v1889 = F_find_among_b(m, l0, int32(4320448), int32(4))
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L368
	} else {
		goto L444
	}
L444:
	;
	if v1889 == int32(0) {
		goto L399
	} else {
		goto L445
	}
L445:
	;
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1893
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v1895)))
	if v1893 < v1896 {
		goto L399
	} else {
		goto L446
	}
L446:
	;
	v1898 = F_slice_del(m, l0)
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L368
	} else {
		goto L447
	}
L447:
	;
	if v1898 < int32(0) {
		v2343 = v1898
		goto L363
	} else {
		goto L448
	}
L448:
	;
	if v1889 != int32(1) {
		goto L399
	} else {
		goto L449
	}
L449:
	;
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1904
	v1906 = int32(2)
	v1908 = int32(0)
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1904-v1911 < v1906 {
		v1921 = v1908
		goto L451
	} else {
		goto L452
	}
L450:
	;
	if v1921 == int32(0) {
		goto L399
	} else {
		goto L454
	}
L451:
	;
	goto L450
L452:
	;
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1917 = F_memcmp(m, v1914+v1904-v1906, int32(2177069), v1906)
	mBase = m.M
	if v1917 != 0 {
		v1921 = v1908
		goto L451
	} else {
		goto L453
	}
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1904 - v1906
	v1921 = int32(1)
	goto L451
L454:
	;
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1924
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(v1926)))
	if v1924 < v1927 {
		goto L399
	} else {
		goto L455
	}
L455:
	;
	v1929 = F_slice_del(m, l0)
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L368
	} else {
		goto L456
	}
L456:
	;
	if int32(0) <= v1929 {
		goto L399
	} else {
		goto L457
	}
L457:
	;
	v2343 = v1929
	goto L363
L458:
	;
	v1936 = F_slice_del(m, l0)
	mBase = m.M
	v1937 = m.ExcPending
	if v1937 != 0 {
		goto L368
	} else {
		goto L459
	}
L459:
	;
	if v1936 < int32(0) {
		v2343 = v1936
		goto L363
	} else {
		goto L460
	}
L460:
	;
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1940
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1940-int32(3) <= v1942 {
		goto L399
	} else {
		goto L461
	}
L461:
	;
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1950 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1946+v1940-int32(1)))))
	if v1950 != int32(101) {
		goto L399
	} else {
		goto L462
	}
L462:
	;
	v1955 = F_find_among_b(m, l0, int32(4320528), int32(3))
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L368
	} else {
		goto L463
	}
L463:
	;
	if v1955 == int32(0) {
		goto L399
	} else {
		goto L464
	}
L464:
	;
	v1959 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1959
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(v1961)))
	if v1959 < v1962 {
		goto L399
	} else {
		goto L465
	}
L465:
	;
	v1964 = F_slice_del(m, l0)
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
		goto L368
	} else {
		goto L466
	}
L466:
	;
	if int32(0) <= v1964 {
		goto L399
	} else {
		goto L467
	}
L467:
	;
	v2343 = v1964
	goto L363
L468:
	;
	v1971 = F_slice_del(m, l0)
	mBase = m.M
	v1972 = m.ExcPending
	if v1972 != 0 {
		goto L368
	} else {
		goto L469
	}
L469:
	;
	if v1971 < int32(0) {
		v2343 = v1971
		goto L363
	} else {
		goto L470
	}
L470:
	;
	v1975 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1975
	v1978 = v1975 - int32(1)
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1978 <= v1979 {
		goto L399
	} else {
		goto L471
	}
L471:
	;
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981+v1978))))
	if v1983&int32(224) != int32(96) {
		goto L399
	} else {
		goto L472
	}
L472:
	;
	if int32(1)<<(uint(v1983)%32)&int32(4198408) == int32(0) {
		goto L399
	} else {
		goto L473
	}
L473:
	;
	v1996 = F_find_among_b(m, l0, int32(4320592), int32(3))
	mBase = m.M
	v1997 = m.ExcPending
	if v1997 != 0 {
		goto L368
	} else {
		goto L474
	}
L474:
	;
	if v1996 == int32(0) {
		goto L399
	} else {
		goto L475
	}
L475:
	;
	v2000 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2000
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v2002)))
	if v2000 < v2003 {
		goto L399
	} else {
		goto L476
	}
L476:
	;
	v2005 = F_slice_del(m, l0)
	mBase = m.M
	v2006 = m.ExcPending
	if v2006 != 0 {
		goto L368
	} else {
		goto L477
	}
L477:
	;
	if int32(0) <= v2005 {
		goto L399
	} else {
		goto L478
	}
L478:
	;
	v2343 = v2005
	goto L363
L479:
	;
	v2012 = F_slice_del(m, l0)
	mBase = m.M
	v2013 = m.ExcPending
	if v2013 != 0 {
		goto L368
	} else {
		goto L480
	}
L480:
	;
	if v2012 < int32(0) {
		v2343 = v2012
		goto L363
	} else {
		goto L481
	}
L481:
	;
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2016
	v2018 = int32(2)
	v2020 = int32(0)
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2016-v2023 < v2018 {
		v2033 = v2020
		goto L483
	} else {
		goto L484
	}
L482:
	;
	if v2033 == int32(0) {
		goto L399
	} else {
		goto L486
	}
L483:
	;
	goto L482
L484:
	;
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2029 = F_memcmp(m, v2026+v2016-v2018, int32(2177071), v2018)
	mBase = m.M
	if v2029 != 0 {
		v2033 = v2020
		goto L483
	} else {
		goto L485
	}
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2016 - v2018
	v2033 = int32(1)
	goto L483
L486:
	;
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2036
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2039 = *(*int32)(unsafe.Add(mBase, uint32(v2038)))
	if v2036 < v2039 {
		goto L399
	} else {
		goto L487
	}
L487:
	;
	v2041 = F_slice_del(m, l0)
	mBase = m.M
	v2042 = m.ExcPending
	if v2042 != 0 {
		goto L368
	} else {
		goto L488
	}
L488:
	;
	if int32(0) <= v2041 {
		goto L399
	} else {
		goto L489
	}
L489:
	;
	v2343 = v2041
	goto L363
L490:
	;
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2098
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2101 = *(*int32)(unsafe.Add(mBase, uint32(v2100)+8))
	if v2098 < v2101 {
		goto L399
	} else {
		goto L510
	}
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2047
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2050
	v2057 = F_find_among_b(m, l0, int32(4320656), int32(12))
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L368
	} else {
		goto L492
	}
L492:
	;
	if v2057 == int32(0) {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2053
	goto L490
L494:
	;
	goto L495
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2053
	v2063 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2063
	if v2063 <= v2053 {
		goto L490
	} else {
		goto L496
	}
L496:
	;
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2066+v2063-int32(1)))))
	if v2070 != int32(117) {
		goto L490
	} else {
		goto L497
	}
L497:
	;
	v2073 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2063 - v2073
	v2077 = F_slice_del(m, l0)
	mBase = m.M
	v2078 = m.ExcPending
	if v2078 != 0 {
		goto L368
	} else {
		goto L499
	}
L498:
	;
	v2089 = int32(0)
	v2090 = base.B2i32(v2084 < v2089)
	if v2090 == v2089 {
		goto L399
	} else {
		goto L506
	}
L499:
	;
	if int32(0) <= v2077 {
		goto L500
	} else {
		goto L501
	}
L500:
	;
	v2084 = v2073
	goto L502
L501:
	;
	v2084 = v2077 >> (uint(int32(31)) % 32) & v2077
	goto L502
L502:
	;
	if v2084 != 0 {
		goto L503
	} else {
		goto L504
	}
L503:
	;
	v2088 = int32(base.Ui32(v2084) >> (uint(int32(31)) % 32))
	goto L505
L504:
	;
	v2088 = int32(4)
	goto L505
L505:
	;
	switch v2088 {
	case 0:
		goto L399
	default:
		goto L498
	case 4:
		goto L490
	}
L506:
	;
	if v2084 < v2089 {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	v2094 = v2084
	goto L509
L508:
	;
	v2094 = int32(1)
	goto L509
L509:
	;
	return v2094
L510:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2098
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2101
	v2108 = F_find_among_b(m, l0, int32(4320896), int32(96))
	mBase = m.M
	v2109 = m.ExcPending
	if v2109 != 0 {
		goto L368
	} else {
		goto L511
	}
L511:
	;
	if v2108 == int32(0) {
		goto L512
	} else {
		goto L513
	}
L512:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2104
	goto L399
L513:
	;
	goto L514
L514:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2104
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2114
	switch v2108 - int32(1) {
	case 0:
		goto L516
	case 1:
		goto L515
	default:
		goto L399
	}
L515:
	;
	v2145 = F_slice_del(m, l0)
	mBase = m.M
	v2146 = m.ExcPending
	if v2146 != 0 {
		goto L368
	} else {
		goto L526
	}
L516:
	;
	if v2114 <= v2104 {
		v2137 = v2114
		goto L517
	} else {
		goto L518
	}
L517:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2137
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2137
	v2141 = F_slice_del(m, l0)
	mBase = m.M
	v2142 = m.ExcPending
	if v2142 != 0 {
		goto L368
	} else {
		goto L524
	}
L518:
	;
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2120 = v2119 + v2114
	v2123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2120-int32(1)))))
	if v2123 != int32(117) {
		v2137 = v2114
		goto L517
	} else {
		goto L519
	}
L519:
	;
	v2127 = v2114 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2127
	if v2127 <= v2104 {
		v2137 = v2114
		goto L517
	} else {
		goto L520
	}
L520:
	;
	v2132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2120-int32(2)))))
	if v2132 == int32(103) {
		goto L521
	} else {
		goto L522
	}
L521:
	;
	v2135 = v2127
	goto L523
L522:
	;
	v2135 = v2114
	goto L523
L523:
	;
	v2137 = v2135
	goto L517
L524:
	;
	if int32(0) <= v2141 {
		goto L399
	} else {
		goto L525
	}
L525:
	;
	v2343 = v2141
	goto L363
L526:
	;
	if v2145 < int32(0) {
		v2343 = v2145
		goto L363
	} else {
		goto L527
	}
L527:
	;
	goto L399
L528:
	;
	v2212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2212
	v2214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2216 = v2212
	v2217 = v2214
	goto L546
L529:
	;
	if v2158 == int32(0) {
		goto L528
	} else {
		goto L530
	}
L530:
	;
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2162
	switch v2158 - int32(1) {
	case 0:
		goto L532
	case 1:
		goto L531
	default:
		goto L528
	}
L531:
	;
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+8))
	if v2162 < v2174 {
		goto L528
	} else {
		goto L536
	}
L532:
	;
	v2166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2167 = *(*int32)(unsafe.Add(mBase, uint32(v2166)+8))
	if v2162 < v2167 {
		goto L528
	} else {
		goto L533
	}
L533:
	;
	v2169 = F_slice_del(m, l0)
	mBase = m.M
	v2170 = m.ExcPending
	if v2170 != 0 {
		goto L368
	} else {
		goto L534
	}
L534:
	;
	if int32(0) <= v2169 {
		goto L528
	} else {
		goto L535
	}
L535:
	;
	v2343 = v2169
	goto L363
L536:
	;
	v2176 = F_slice_del(m, l0)
	mBase = m.M
	v2177 = m.ExcPending
	if v2177 != 0 {
		goto L368
	} else {
		goto L537
	}
L537:
	;
	if v2176 < int32(0) {
		v2343 = v2176
		goto L363
	} else {
		goto L538
	}
L538:
	;
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2180
	v2182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2180 <= v2182 {
		goto L528
	} else {
		goto L539
	}
L539:
	;
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2185 = v2184 + v2180
	v2188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2185-int32(1)))))
	if v2188 != int32(117) {
		goto L528
	} else {
		goto L540
	}
L540:
	;
	v2192 = v2180 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2192
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2192
	if v2192 <= v2182 {
		goto L528
	} else {
		goto L541
	}
L541:
	;
	v2198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2185-int32(2)))))
	if v2198 != int32(103) {
		goto L528
	} else {
		goto L542
	}
L542:
	;
	v2201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2202 = *(*int32)(unsafe.Add(mBase, uint32(v2201)+8))
	if v2180 <= v2202 {
		goto L528
	} else {
		goto L543
	}
L543:
	;
	v2204 = F_slice_del(m, l0)
	mBase = m.M
	v2205 = m.ExcPending
	if v2205 != 0 {
		goto L368
	} else {
		goto L544
	}
L544:
	;
	if v2204 < int32(0) {
		v2343 = v2204
		goto L363
	} else {
		goto L545
	}
L545:
	;
	goto L528
L546:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2216
	v2222 = v2216 + int32(1)
	if v2217 <= v2222 {
		goto L552
	} else {
		goto L553
	}
L547:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2212
	v2343 = int32(1)
	goto L363
L548:
	;
	goto L547
L549:
	;
	v2338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2216 = v2339
	v2217 = v2338
	goto L546
L550:
	;
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L575
L551:
	;
	v2239 = F_find_among(m, l0, int32(4322976), int32(6))
	mBase = m.M
	v2240 = m.ExcPending
	if v2240 != 0 {
		goto L368
	} else {
		goto L556
	}
L552:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2216
	v2276 = v2216
	v2277 = v2217
	goto L550
L553:
	;
	v2224 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2224+v2222))))
	if v2226&int32(224) != int32(160) {
		goto L552
	} else {
		goto L554
	}
L554:
	;
	if int32(1)<<(uint(v2226)%32)&int32(67641858) != 0 {
		goto L551
	} else {
		goto L555
	}
L555:
	;
	goto L552
L556:
	;
	v2241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2241
	switch v2239 - int32(1) {
	case 0:
		goto L562
	case 1:
		goto L561
	case 2:
		goto L560
	case 3:
		goto L559
	case 4:
		goto L558
	case 5:
		goto L557
	default:
		goto L549
	}
L557:
	;
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2276 = v2241
	v2277 = v2275
	goto L550
L558:
	;
	v2271 = F_slice_from_s(m, l0, int32(1), int32(2177832))
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L368
	} else {
		goto L571
	}
L559:
	;
	v2265 = F_slice_from_s(m, l0, int32(1), int32(2177831))
	mBase = m.M
	v2266 = m.ExcPending
	if v2266 != 0 {
		goto L368
	} else {
		goto L569
	}
L560:
	;
	v2259 = F_slice_from_s(m, l0, int32(1), int32(2177830))
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L368
	} else {
		goto L567
	}
L561:
	;
	v2253 = F_slice_from_s(m, l0, int32(1), int32(2177829))
	mBase = m.M
	v2254 = m.ExcPending
	if v2254 != 0 {
		goto L368
	} else {
		goto L565
	}
L562:
	;
	v2247 = F_slice_from_s(m, l0, int32(1), int32(2177828))
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L368
	} else {
		goto L563
	}
L563:
	;
	if int32(0) <= v2247 {
		goto L549
	} else {
		goto L564
	}
L564:
	;
	v2343 = v2247
	goto L363
L565:
	;
	if int32(0) <= v2253 {
		goto L549
	} else {
		goto L566
	}
L566:
	;
	v2343 = v2253
	goto L363
L567:
	;
	if int32(0) <= v2259 {
		goto L549
	} else {
		goto L568
	}
L568:
	;
	v2343 = v2259
	goto L363
L569:
	;
	if int32(0) <= v2265 {
		goto L549
	} else {
		goto L570
	}
L570:
	;
	v2343 = v2265
	goto L363
L571:
	;
	if int32(0) <= v2271 {
		goto L549
	} else {
		goto L572
	}
L572:
	;
	v2343 = v2271
	goto L363
L573:
	;
	if v2331 < int32(0) {
		goto L548
	} else {
		goto L593
	}
L575:
	;
	goto L576
L576:
	;
	goto L577
L577:
	;
	v2286 = v2276
	v2288 = int32(1)
	goto L580
L579:
	;
	v2331 = v2316
	goto L573
L580:
	;
	if v2277 <= v2286 {
		goto L582
	} else {
		goto L583
	}
L581:
	;
	goto L579
L582:
	;
	v2331 = int32(-1)
	goto L573
L583:
	;
	goto L584
L584:
	;
	v2293 = v2286 + int32(1)
	v2295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2279+v2286))))
	if base.Ui32(v2295) < base.Ui32(int32(192)) {
		v2316 = v2293
		goto L585
	} else {
		goto L586
	}
L585:
	;
	v2317 = int32(1)
	if v2317 < v2288 {
		v2286 = v2316
		v2288 = v2288 - v2317
		goto L580
	} else {
		goto L592
	}
L586:
	;
	if v2277 <= v2293 {
		v2316 = v2293
		goto L585
	} else {
		goto L587
	}
L587:
	;
	v2302 = v2293
	goto L588
L588:
	;
	v2305 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2279+v2302))))
	if int32(-65) < v2305 {
		v2316 = v2302
		goto L585
	} else {
		goto L590
	}
L589:
	;
	v2316 = v2277
	goto L585
L590:
	;
	v2309 = v2302 + int32(1)
	if v2309 != v2277 {
		v2302 = v2309
		goto L588
	} else {
		goto L591
	}
L591:
	;
	goto L589
L592:
	;
	goto L581
L593:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2331
	goto L549
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
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v38 int64
	_ = v38
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v51 int64
	_ = v51
	var v56 int64
	_ = v56
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v103 int64
	_ = v103
	var v109 int64
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int64
	_ = v115
	var v126 int64
	_ = v126
	v8 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v13 = (int64(base.Ui64(v8)>>(uint(int64(23))%64)) ^ v8) * int64(2388976653695081527)
	v17 = int64(-8645972361240307355)
	v20 = (v13 ^ int64(base.Ui64(v13)>>(uint(int64(47))%64)) ^ v17) * v17
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v22 == int32(0) {
		v112 = v21
		v115 = v20
	} else {
		v25 = v21
		v28 = v20
		v31 = v22
		for {
			v33 = int64(*(*int8)(unsafe.Add(mBase, uint32(v25)+1)))
			if v33 == int64(0) {
				v90 = int32(1)
				v91 = int64(0)
				v97 = v90
				v98 = v91 | base.I64_extend8_s(base.I64_extend_i32_u(v31))
			} else {
				v38 = int64(*(*int8)(unsafe.Add(mBase, uint32(v25)+2)))
				if v38 == int64(0) {
					v84 = int32(2)
					v85 = int64(0)
					v90 = v84
					v91 = v85 | v33<<(uint(int64(8))%64)
					v97 = v90
					v98 = v91 | base.I64_extend8_s(base.I64_extend_i32_u(v31))
				} else {
					v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+3)))
					if v42 != 0 {
						v44 = int64(*(*int8)(unsafe.Add(mBase, uint32(v25)+4)))
						if v44 == int64(0) {
							v77 = int32(4)
							v78 = int64(0)
							v79 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v25))))
							v97 = v77
							v98 = v78 | v79
						} else {
							v51 = int64(*(*int8)(unsafe.Add(mBase, uint32(v25)+5)))
							if v51 == int64(0) {
								v72 = int32(5)
								v73 = int64(0)
								v77 = v72
								v78 = v44<<(uint(int64(32))%64) | v73
								v79 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v25))))
								v97 = v77
								v98 = v78 | v79
							} else {
								v56 = int64(*(*int8)(unsafe.Add(mBase, uint32(v25)+6)))
								if v56 == int64(0) {
									v66 = int32(6)
									v67 = int64(0)
									v72 = v66
									v73 = v67 | v51<<(uint(int64(40))%64)
									v77 = v72
									v78 = v44<<(uint(int64(32))%64) | v73
									v79 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v25))))
									v97 = v77
									v98 = v78 | v79
								} else {
									v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+7)))
									if v60 != 0 {
										v62 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
										v97 = int32(8)
										v98 = v62
									} else {
										v66 = int32(7)
										v67 = v56 << (uint(int64(48)) % 64)
										v72 = v66
										v73 = v67 | v51<<(uint(int64(40))%64)
										v77 = v72
										v78 = v44<<(uint(int64(32))%64) | v73
										v79 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v25))))
										v97 = v77
										v98 = v78 | v79
									}
								}
							}
						}
					} else {
						v84 = int32(3)
						v85 = v38 << (uint(int64(16)) % 64)
						v90 = v84
						v91 = v85 | v33<<(uint(int64(8))%64)
						v97 = v90
						v98 = v91 | base.I64_extend8_s(base.I64_extend_i32_u(v31))
					}
				}
			}
			v103 = (v98 ^ int64(base.Ui64(v98)>>(uint(int64(23))%64))) * int64(2388976653695081527)
			v109 = (v28 ^ int64(base.Ui64(v103)>>(uint(int64(47))%64)) ^ v103) * int64(-8645972361240307355)
			v110 = v25 + v97
			v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
			if v111 != 0 {
				v25 = v110
				v28 = v109
				v31 = v111
				continue
			} else {
				break
			}
			break
		}
		v112 = v110
		v115 = v109
	}
	v126 = (base.I64_extend_i32_s(v112-v21) + int64(base.Ui64(v115)>>(uint(int64(23))%64)) ^ v115) * int64(2388976653695081527)
	return base.I32_wrap_i64(int64(base.Ui64(v126)>>(uint(int64(47))%64)) ^ v126 - int64(base.Ui64(v126)>>(uint(int64(32))%64)))
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
	v7 = int32(4449520)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l5)+96))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v10
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
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v8
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
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int64
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
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
	v29 = int32(4449520)
	v30 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v34
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
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v30
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
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	if v107 == int32(3) {
		goto L37
	} else {
		goto L38
	}
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+379)) = uint8(v98)
	v102 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+377)) = uint8(v102)
	goto L10
L24:
	;
	if v56 == int32(0) {
		v98 = v54
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
	if v60&int32(2) != 0 {
		v98 = v54
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
	v76 = int32(4449520)
	v77 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v79
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
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v77
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
	v98 = int32(1)
	goto L23
L36:
	;
	if v28 == int64(-1) {
		goto L15
	} else {
		goto L41
	}
L37:
	;
	v110 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	v118 = v110
	goto L36
L38:
	;
	goto L39
L39:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	F_tuplestore_puttupleslot(m, v111, v56)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L20
	} else {
		goto L40
	}
L40:
	;
	v114 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	v116 = v114 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v116
	v118 = v116
	goto L36
L41:
	;
	if v118 <= v28 {
		goto L15
	} else {
		goto L42
	}
L42:
	;
	goto L16
}
func F_standard_qp_callback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
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
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v264 int32
	_ = v264
	var v274 int32
	_ = v274
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v435 int32
	_ = v435
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
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
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v671 int32
	_ = v671
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v839 int32
	_ = v839
	var v858 int32
	_ = v858
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v917 int32
	_ = v917
	var v924 int32
	_ = v924
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v952 int32
	_ = v952
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1050 int32
	_ = v1050
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1084 int32
	_ = v1084
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1163 int32
	_ = v1163
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1259 int32
	_ = v1259
	var v1281 int32
	_ = v1281
	var v1285 int32
	_ = v1285
	var v1296 int32
	_ = v1296
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1321 int32
	_ = v1321
	var v1331 int32
	_ = v1331
	var v1343 int32
	_ = v1343
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1383 int32
	_ = v1383
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	v3 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v25 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v22 != 0 {
		goto L271
	} else {
		goto L272
	}
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v24)+100))
	if v94 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v31 = v29
	goto L7
L6:
	;
	v31 = int32(0)
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v31
	if v31 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v77 != 0 {
		goto L21
	} else {
		goto L22
	}
L9:
	;
	v77 = int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v41 <= int32(0) {
		v69 = int32(1)
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v77 = v69
	goto L8
L13:
	;
	v44 = int32(0)
	if v44 < v41 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v47 = v41
	goto L16
L15:
	;
	v47 = v44
	goto L16
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v50 = int32(0)
	goto L17
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v48+v50<<(uint(int32(2))%32))))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v60 = int32(0)
	v61 = base.B2i32(v59 != v60)
	if v59 == v60 {
		v69 = v61
		goto L12
	} else {
		goto L19
	}
L18:
	;
	v69 = v61
	goto L12
L19:
	;
	v65 = v50 + int32(1)
	if v65 != v47 {
		v50 = v65
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v78 = int32(0)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+45)))
	v86 = F_make_pathkeys_for_sortclauses_extended(m, l0, v20+int32(12), v23, v78, v82, v20+int32(11), v78)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = int64(0)
	goto L1
L24:
	;
	return
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v86
	if v86 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v90 = v89
	goto L28
L27:
	;
	v90 = v78
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v90
	goto L1
L29:
	;
	v285 = int32(0)
	if v274 == v285 {
		goto L71
	} else {
		goto L72
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = int64(0)
	goto L1
L31:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+336))
	if v97 <= int32(0) {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v100 = int32(0)
	v103 = int32(1)
	v108 = F_make_pathkeys_for_sortclauses_extended(m, l0, l0+int32(256), v23, v103, v100, v20+int32(10), v103)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L24
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v108
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+10)))
	if v111 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = int64(0)
	goto L1
L37:
	;
	goto L38
L38:
	;
	if v108 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v117 = v116
	goto L41
L40:
	;
	v117 = v100
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+336))
	if v119 <= int32(0) {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, _consts[609])))
	if v123 != int32(1) {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v126 == int32(0) {
		v274 = v3
		goto L29
	} else {
		goto L44
	}
L44:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v129 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v274 = v3
	goto L29
L46:
	;
	goto L47
L47:
	;
	v136 = v3
	v138 = v3
	v139 = v129
	goto L48
L48:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v149+v136<<(uint(int32(2))%32))))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+50)))
	if v157 != int32(110) {
		v252 = v138
		v253 = v139
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v274 = v252
	goto L29
L50:
	;
	v264 = v136 + int32(1)
	if v264 < v253 {
		v136 = v264
		v138 = v252
		v139 = v253
		goto L48
	} else {
		goto L68
	}
L51:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156)+40))
	if v160 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v156)+36))
	if v163 == int32(0) {
		v252 = v138
		v253 = v139
		goto L50
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v156)+44))
	if v166 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	v243 = F_bms_add_member(m, v138, v136)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L24
	} else {
		goto L67
	}
L57:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v156)+32))
	if v169 == int32(0) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	if v172 <= int32(0) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v169)+12))
	v185 = int32(0)
	goto L60
L60:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v175+v185<<(uint(int32(2))%32))))
	v200 = v197
	goto L62
L61:
	;
	goto L56
L62:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	if v216 == int32(27) {
		v200 = v215
		goto L62
	} else {
		goto L64
	}
L63:
	;
	if base.Ui32(int32(1)) < base.Ui32(v216-int32(6)) {
		v252 = v138
		v253 = v139
		goto L50
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	v224 = v185 + int32(1)
	if v172 != v224 {
		v185 = v224
		goto L60
	} else {
		goto L66
	}
L66:
	;
	goto L61
L67:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v252 = v243
	v253 = v245
	goto L50
L68:
	;
	goto L49
L69:
	;
	if v858 == int32(0) {
		goto L241
	} else {
		goto L242
	}
L70:
	;
	goto L84
L71:
	;
	v321 = int32(0)
	goto L70
L72:
	;
	goto L73
L73:
	;
	v293 = int32(1)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	if v294 <= v293 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v297 = v293
	goto L76
L75:
	;
	v297 = v294
	goto L76
L76:
	;
	v301 = int32(0)
	v303 = v285
	goto L77
L77:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v274+int32(8)+v301<<(uint(int32(2))%32))))
	if v309 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v321 = v312
	goto L70
L79:
	;
	v312 = v303 + base.I32_popcnt(v309)
	goto L81
L80:
	;
	v312 = v303
	goto L81
L81:
	;
	v314 = v301 + int32(1)
	if v314 != v297 {
		v301 = v314
		v303 = v312
		goto L77
	} else {
		goto L82
	}
L82:
	;
	goto L78
L83:
	;
	if v321 <= int32(0) {
		v858 = v285
		goto L69
	} else {
		goto L96
	}
L84:
	;
	goto L83
L96:
	;
	v366 = v274
	v374 = v285
	v375 = v3
	goto L97
L97:
	;
	v377 = int32(0)
	if v366 == v377 {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	if v766 == int32(0) {
		v858 = v803
		goto L69
	} else {
		goto L238
	}
L99:
	;
	if int32(0) <= v435 {
		goto L110
	} else {
		goto L111
	}
L100:
	;
	v435 = base.I32_ctz(v421) | v422<<(uint(int32(5))%32)
	goto L99
L101:
	;
	v435 = int32(-2)
	goto L99
L102:
	;
	v388 = base.I32_div_s(int32(0), int32(32))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v366)+4))
	if v389 <= v388 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v392 = v366 + int32(8)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v392+v388<<(uint(int32(2))%32))))
	v399 = v396 & int32(-1)
	if v399 != 0 {
		v421 = v399
		v422 = v388
		goto L100
	} else {
		goto L104
	}
L104:
	;
	v401 = v388 + int32(1)
	if v401 == v389 {
		goto L101
	} else {
		goto L105
	}
L105:
	;
	v404 = v401
	goto L106
L106:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v392+v404<<(uint(int32(2))%32))))
	if v411 != 0 {
		v421 = v411
		v422 = v404
		goto L100
	} else {
		goto L108
	}
L107:
	;
	goto L101
L108:
	;
	v413 = v404 + int32(1)
	if v413 != v389 {
		v404 = v413
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v443 = v435
	v444 = v366
	v446 = v377
	v449 = v377
	goto L113
L111:
	;
	v680 = v366
	v682 = v377
	v685 = v377
	goto L112
L112:
	;
	v691 = F_bms_del_members(m, v680, v685)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L24
	} else {
		goto L178
	}
L113:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)+12))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v456+v443<<(uint(int32(2))%32))))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)+4))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)+12))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v463)+40))
	if v464 != 0 {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	v680 = v605
	v682 = v607
	v685 = v610
	goto L112
L115:
	;
	if v605 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L116:
	;
	if v446 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L117:
	;
	v466 = v464
	goto L119
L118:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v463)+36))
	v466 = v465
	goto L119
L119:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v463)+32))
	v468 = F_make_pathkeys_for_sortclauses(m, l0, v466, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L24
	} else {
		goto L120
	}
L120:
	;
	if v468 == int32(0) {
		goto L116
	} else {
		goto L121
	}
L121:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v468)+4))
	if v472 <= int32(0) {
		goto L116
	} else {
		goto L122
	}
L122:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v468)+12))
	v479 = int32(0)
	goto L123
L123:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v475+v479<<(uint(int32(2))%32))))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+41)))
	if v499 != int32(1) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v505 = F_bms_del_member(m, v444, v443)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L24
	} else {
		goto L129
	}
L125:
	;
	v503 = v479 + int32(1)
	if v503 != v472 {
		v479 = v503
		goto L123
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	goto L124
L128:
	;
	goto L116
L129:
	;
	v605 = v505
	v607 = v446
	v610 = v449
	goto L115
L130:
	;
	v597 = F_bms_add_member(m, v449, v443)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L24
	} else {
		goto L165
	}
L131:
	;
	v596 = v594
	goto L130
L132:
	;
	if v108 == int32(0) {
		v594 = v468
		goto L131
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	if v108 != 0 {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v528 = F_list_copy(m, v108)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L24
	} else {
		goto L136
	}
L136:
	;
	v530 = F_append_pathkeys(m, v528, v468)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L24
	} else {
		goto L137
	}
L137:
	;
	v596 = v530
	goto L130
L138:
	;
	v532 = F_list_copy(m, v108)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L24
	} else {
		goto L141
	}
L139:
	;
	v536 = v468
	goto L140
L140:
	;
	if v446 == v536 {
		goto L144
	} else {
		goto L145
	}
L141:
	;
	v534 = F_append_pathkeys(m, v532, v468)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L24
	} else {
		goto L142
	}
L142:
	;
	v536 = v534
	goto L140
L143:
	;
	switch v593 {
	case 0, 1:
		v596 = v446
		goto L130
	case 2:
		v594 = v536
		goto L131
	default:
		v605 = v444
		v607 = v446
		v610 = v449
		goto L115
	}
L144:
	;
	v593 = int32(0)
	goto L143
L145:
	;
	goto L146
L146:
	;
	v544 = int32(0)
	goto L149
L147:
	;
	if v583 != 0 {
		goto L162
	} else {
		goto L163
	}
L148:
	;
	v577 = int32(0)
	v583 = base.B2i32(v557 == v577)
	v585 = base.B2i32(v566 != v577) << (uint(int32(1)) % 32)
	goto L147
L149:
	;
	v547 = int32(0)
	if v446 == v547 {
		v557 = v547
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v593 = int32(3)
	goto L143
L151:
	;
	if v536 != 0 {
		goto L155
	} else {
		goto L156
	}
L152:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v446)+4))
	if v551 <= v544 {
		v557 = int32(0)
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v446)+12))
	v557 = v553 + v544<<(uint(int32(2))%32)
	goto L151
L154:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v536)+12))
	v566 = v563 + v544<<(uint(int32(2))%32)
	if v557 == int32(0) {
		goto L148
	} else {
		goto L159
	}
L155:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	if v544 < v558 {
		goto L154
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v560 = int32(0)
	v583 = base.B2i32(v557 == v560)
	v585 = v560
	goto L147
L158:
	;
	goto L157
L159:
	;
	if v566 == int32(0) {
		goto L148
	} else {
		goto L160
	}
L160:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v557)))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v566)))
	if v573 == v574 {
		v544 = v544 + int32(1)
		goto L149
	} else {
		goto L161
	}
L161:
	;
	goto L150
L162:
	;
	v587 = v585
	goto L164
L163:
	;
	v587 = int32(1)
	goto L164
L164:
	;
	v593 = v587
	goto L143
L165:
	;
	v605 = v444
	v607 = v596
	v610 = v597
	goto L115
L166:
	;
	if int32(0) <= v671 {
		v443 = v671
		v444 = v605
		v446 = v607
		v449 = v610
		goto L113
	} else {
		goto L177
	}
L167:
	;
	v671 = base.I32_ctz(v657) | v658<<(uint(int32(5))%32)
	goto L166
L168:
	;
	v671 = int32(-2)
	goto L166
L169:
	;
	v622 = v443 + int32(1)
	v624 = base.I32_div_s(v622, int32(32))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v605)+4))
	if v625 <= v624 {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v628 = v605 + int32(8)
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v628+v624<<(uint(int32(2))%32))))
	v635 = v632 & (int32(-1) << (uint(v622) % 32))
	if v635 != 0 {
		v657 = v635
		v658 = v624
		goto L167
	} else {
		goto L171
	}
L171:
	;
	v637 = v624 + int32(1)
	if v637 == v625 {
		goto L168
	} else {
		goto L172
	}
L172:
	;
	v640 = v637
	goto L173
L173:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v628+v640<<(uint(int32(2))%32))))
	if v647 != 0 {
		v657 = v647
		v658 = v640
		goto L167
	} else {
		goto L175
	}
L174:
	;
	goto L168
L175:
	;
	v649 = v640 + int32(1)
	if v649 != v625 {
		v640 = v649
		goto L173
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	goto L114
L178:
	;
	v693 = int32(0)
	if v685 == v693 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	v729 = int32(0)
	if v374 == v729 {
		goto L193
	} else {
		goto L194
	}
L180:
	;
	v728 = int32(0)
	goto L179
L181:
	;
	goto L182
L182:
	;
	v700 = int32(1)
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v685)+4))
	if v701 <= v700 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v704 = v700
	goto L185
L184:
	;
	v704 = v701
	goto L185
L185:
	;
	v708 = int32(0)
	v710 = v693
	goto L186
L186:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v685+int32(8)+v708<<(uint(int32(2))%32))))
	if v716 != 0 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v728 = v719
	goto L179
L188:
	;
	v719 = v710 + base.I32_popcnt(v716)
	goto L190
L189:
	;
	v719 = v710
	goto L190
L190:
	;
	v721 = v708 + int32(1)
	if v721 != v704 {
		v708 = v721
		v710 = v719
		goto L186
	} else {
		goto L191
	}
L191:
	;
	goto L187
L192:
	;
	v765 = base.B2i32(v764 < v728)
	if v764 < v728 {
		goto L205
	} else {
		goto L206
	}
L193:
	;
	v764 = int32(0)
	goto L192
L194:
	;
	goto L195
L195:
	;
	v736 = int32(1)
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	if v737 <= v736 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v740 = v736
	goto L198
L197:
	;
	v740 = v737
	goto L198
L198:
	;
	v744 = int32(0)
	v746 = v729
	goto L199
L199:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v374+int32(8)+v744<<(uint(int32(2))%32))))
	if v752 != 0 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v764 = v755
	goto L192
L201:
	;
	v755 = v746 + base.I32_popcnt(v752)
	goto L203
L202:
	;
	v755 = v746
	goto L203
L203:
	;
	v757 = v744 + int32(1)
	if v757 != v740 {
		v744 = v757
		v746 = v755
		goto L199
	} else {
		goto L204
	}
L204:
	;
	goto L200
L205:
	;
	v766 = v682
	goto L207
L206:
	;
	v766 = v375
	goto L207
L207:
	;
	v767 = int32(0)
	if v691 == v767 {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	if v764 < v728 {
		goto L221
	} else {
		goto L222
	}
L209:
	;
	v802 = int32(0)
	goto L208
L210:
	;
	goto L211
L211:
	;
	v774 = int32(1)
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v691)+4))
	if v775 <= v774 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v778 = v774
	goto L214
L213:
	;
	v778 = v775
	goto L214
L214:
	;
	v782 = int32(0)
	v784 = v767
	goto L215
L215:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v691+int32(8)+v782<<(uint(int32(2))%32))))
	if v790 != 0 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v802 = v793
	goto L208
L217:
	;
	v793 = v784 + base.I32_popcnt(v790)
	goto L219
L218:
	;
	v793 = v784
	goto L219
L219:
	;
	v795 = v782 + int32(1)
	if v795 != v778 {
		v782 = v795
		v784 = v793
		goto L215
	} else {
		goto L220
	}
L220:
	;
	goto L216
L221:
	;
	v803 = v685
	goto L223
L222:
	;
	v803 = v374
	goto L223
L223:
	;
	v804 = int32(0)
	if v803 == v804 {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	if v839 < v802 {
		v366 = v691
		v374 = v803
		v375 = v766
		goto L97
	} else {
		goto L237
	}
L225:
	;
	v839 = int32(0)
	goto L224
L226:
	;
	goto L227
L227:
	;
	v811 = int32(1)
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v803)+4))
	if v812 <= v811 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v815 = v811
	goto L230
L229:
	;
	v815 = v812
	goto L230
L230:
	;
	v819 = int32(0)
	v821 = v804
	goto L231
L231:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v803+int32(8)+v819<<(uint(int32(2))%32))))
	if v827 != 0 {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v839 = v830
	goto L224
L233:
	;
	v830 = v821 + base.I32_popcnt(v827)
	goto L235
L234:
	;
	v830 = v821
	goto L235
L235:
	;
	v832 = v819 + int32(1)
	if v832 != v815 {
		v819 = v832
		v821 = v830
		goto L231
	} else {
		goto L236
	}
L236:
	;
	goto L232
L237:
	;
	goto L98
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v766
	v858 = v803
	goto L69
L239:
	;
	if v917 < int32(0) {
		goto L1
	} else {
		goto L250
	}
L240:
	;
	v917 = base.I32_ctz(v903) | v904<<(uint(int32(5))%32)
	goto L239
L241:
	;
	v917 = int32(-2)
	goto L239
L242:
	;
	v870 = base.I32_div_s(int32(0), int32(32))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v858)+4))
	if v871 <= v870 {
		goto L241
	} else {
		goto L243
	}
L243:
	;
	v874 = v858 + int32(8)
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v874+v870<<(uint(int32(2))%32))))
	v881 = v878 & int32(-1)
	if v881 != 0 {
		v903 = v881
		v904 = v870
		goto L240
	} else {
		goto L244
	}
L244:
	;
	v883 = v870 + int32(1)
	if v883 == v871 {
		goto L241
	} else {
		goto L245
	}
L245:
	;
	v886 = v883
	goto L246
L246:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v874+v886<<(uint(int32(2))%32))))
	if v893 != 0 {
		v903 = v893
		v904 = v886
		goto L240
	} else {
		goto L248
	}
L247:
	;
	goto L241
L248:
	;
	v895 = v886 + int32(1)
	if v895 != v871 {
		v886 = v895
		goto L246
	} else {
		goto L249
	}
L249:
	;
	goto L247
L250:
	;
	v924 = v917
	goto L251
L251:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v937)+12))
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v938+v924<<(uint(int32(2))%32))))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v942)+4))
	if v943 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	goto L1
L253:
	;
	if v858 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L254:
	;
	v946 = int32(0)
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v943)+4))
	if v947 <= v946 {
		goto L253
	} else {
		goto L255
	}
L255:
	;
	v952 = v946
	goto L256
L256:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v943)+12))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v967+v952<<(uint(int32(2))%32))))
	v972 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v971)+51)) = uint8(v972)
	v975 = v952 + v972
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v943)+4))
	if v975 < v976 {
		v952 = v975
		goto L256
	} else {
		goto L258
	}
L257:
	;
	goto L253
L258:
	;
	goto L257
L259:
	;
	if int32(0) <= v1050 {
		v924 = v1050
		goto L251
	} else {
		goto L270
	}
L260:
	;
	v1050 = base.I32_ctz(v1036) | v1037<<(uint(int32(5))%32)
	goto L259
L261:
	;
	v1050 = int32(-2)
	goto L259
L262:
	;
	v1001 = v924 + int32(1)
	v1003 = base.I32_div_s(v1001, int32(32))
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v858)+4))
	if v1004 <= v1003 {
		goto L261
	} else {
		goto L263
	}
L263:
	;
	v1007 = v858 + int32(8)
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v1007+v1003<<(uint(int32(2))%32))))
	v1014 = v1011 & (int32(-1) << (uint(v1001) % 32))
	if v1014 != 0 {
		v1036 = v1014
		v1037 = v1003
		goto L260
	} else {
		goto L264
	}
L264:
	;
	v1016 = v1003 + int32(1)
	if v1016 == v1004 {
		goto L261
	} else {
		goto L265
	}
L265:
	;
	v1019 = v1016
	goto L266
L266:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v1007+v1019<<(uint(int32(2))%32))))
	if v1026 != 0 {
		v1036 = v1026
		v1037 = v1019
		goto L260
	} else {
		goto L268
	}
L267:
	;
	goto L261
L268:
	;
	v1028 = v1019 + int32(1)
	if v1028 != v1004 {
		v1019 = v1028
		goto L266
	} else {
		goto L269
	}
L269:
	;
	goto L267
L270:
	;
	goto L252
L271:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1070)))
	v1072 = F_make_pathkeys_for_window(m, l0, v1071, v23)
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L24
	} else {
		goto L274
	}
L272:
	;
	v1075 = int32(0)
	goto L273
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v1075
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v24)+120))
	if v1077 != 0 {
		goto L276
	} else {
		goto L277
	}
L274:
	;
	v1075 = v1072
	goto L273
L275:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v24)+124))
	v1097 = F_make_pathkeys_for_sortclauses(m, l0, v1096, v23)
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L24
	} else {
		goto L284
	}
L276:
	;
	v1078 = F_list_copy(m, v1077)
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L24
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	goto L275
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = v1078
	v1084 = int32(0)
	v1088 = F_make_pathkeys_for_sortclauses_extended(m, l0, l0+int32(260), v23, int32(1), v1084, v20+int32(9), v1084)
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L24
	} else {
		goto L280
	}
L280:
	;
	v1091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+9)))
	if v1091 != 0 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1092 = v1088
	goto L283
L282:
	;
	v1092 = int32(0)
	goto L283
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v1092
	goto L275
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v1097
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v1100 != 0 {
		goto L286
	} else {
		goto L287
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v1371
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v1373 != 0 {
		goto L357
	} else {
		goto L358
	}
L286:
	;
	v1101 = int32(0)
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1100)+32))
	v1104 = F_copyObjectImpl(m, v1103)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L24
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	v1371 = int32(0)
	goto L285
L289:
	;
	if v1104 != 0 {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+12))
	v1107 = v1106
	goto L292
L291:
	;
	v1107 = v1101
	goto L292
L292:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1100)+20))
	if v1108 != 0 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+12))
	v1110 = v1109
	goto L295
L294:
	;
	v1110 = v1101
	goto L295
L295:
	;
	if v23 == int32(0) {
		v1331 = v1104
		goto L296
	} else {
		goto L297
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v1331
	v1343 = int32(0)
	v1348 = F_make_pathkeys_for_sortclauses_extended(m, l0, v20+int32(4), v23, v1343, v1343, v20+int32(3), v1343)
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L24
	} else {
		goto L352
	}
L297:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v1113 <= int32(0) {
		v1331 = v1104
		goto L296
	} else {
		goto L298
	}
L298:
	;
	v1119 = int32(0)
	v1121 = v1110
	v1122 = v1107
	v1124 = v1113
	goto L299
L299:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1134+v1119<<(uint(int32(2))%32))))
	v1139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1138)+26)))
	if v1139 == int32(0) {
		goto L301
	} else {
		goto L302
	}
L300:
	;
	v1331 = v1104
	goto L296
L301:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1122)))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1121)))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1138)+4))
	v1145 = F_exprType(m, v1144)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L24
	} else {
		goto L304
	}
L302:
	;
	v1314 = v1121
	v1315 = v1122
	v1317 = v1124
	goto L303
L303:
	;
	v1321 = v1119 + int32(1)
	if v1321 < v1317 {
		v1119 = v1321
		v1121 = v1314
		v1122 = v1315
		v1124 = v1317
		goto L299
	} else {
		goto L351
	}
L304:
	;
	if v1143 != v1145 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1331 = int32(0)
	goto L296
L306:
	;
	goto L307
L307:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1100)+20))
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1149)+12))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1149)+4))
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+12))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+4))
	v1154 = int32(0)
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1138)+16))
	if v1163 == v1154 {
		goto L309
	} else {
		goto L310
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1142)+4)) = v1285
	v1296 = v1122 + int32(4)
	if base.Ui32(v1296) < base.Ui32(v1152+v1153<<(uint(int32(2))%32)) {
		goto L345
	} else {
		goto L346
	}
L309:
	;
	if v23 == int32(0) {
		v1281 = int32(1)
		goto L312
	} else {
		goto L313
	}
L310:
	;
	v1285 = v1163
	goto L311
L311:
	;
	goto L308
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1138)+16)) = v1281
	v1285 = v1281
	goto L311
L313:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v1170 <= int32(0) {
		v1281 = int32(1)
		goto L312
	} else {
		goto L314
	}
L314:
	;
	v1173 = int32(0)
	if v1173 < v1170 {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1176 = v1170
	goto L317
L316:
	;
	v1176 = v1173
	goto L317
L317:
	;
	v1178 = v1176 & int32(3)
	v1179 = int32(0)
	if int32(4) <= v1170 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v1188 = v1179
	v1189 = v1154
	v1190 = int32(0)
	goto L321
L319:
	;
	v1223 = v1179
	v1224 = v1154
	goto L320
L320:
	;
	if v1178 != 0 {
		goto L336
	} else {
		goto L337
	}
L321:
	;
	v1199 = v1184 + v1189<<(uint(int32(2))%32)
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1199)+12))
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1200)+16))
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1199)+8))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1202)+16))
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1199)+4))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1204)+16))
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1199)))
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+16))
	if base.Ui32(v1188) < base.Ui32(v1207) {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	v1223 = v1215
	v1224 = v1217
	goto L320
L323:
	;
	v1209 = v1207
	goto L325
L324:
	;
	v1209 = v1188
	goto L325
L325:
	;
	if base.Ui32(v1209) < base.Ui32(v1205) {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v1211 = v1205
	goto L328
L327:
	;
	v1211 = v1209
	goto L328
L328:
	;
	if base.Ui32(v1211) < base.Ui32(v1203) {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1213 = v1203
	goto L331
L330:
	;
	v1213 = v1211
	goto L331
L331:
	;
	if base.Ui32(v1213) < base.Ui32(v1201) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v1215 = v1201
	goto L334
L333:
	;
	v1215 = v1213
	goto L334
L334:
	;
	v1216 = int32(4)
	v1217 = v1189 + v1216
	v1219 = v1190 + v1216
	if v1219 != v1176&int32(2147483644) {
		v1188 = v1215
		v1189 = v1217
		v1190 = v1219
		goto L321
	} else {
		goto L335
	}
L335:
	;
	goto L322
L336:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v1234 = int32(0)
	v1236 = v1223
	v1237 = v1224
	goto L339
L337:
	;
	v1259 = v1223
	goto L338
L338:
	;
	v1281 = v1259 + int32(1)
	goto L312
L339:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1232+v1237<<(uint(int32(2))%32))))
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1248)+16))
	if base.Ui32(v1236) < base.Ui32(v1249) {
		goto L341
	} else {
		goto L342
	}
L340:
	;
	v1259 = v1251
	goto L338
L341:
	;
	v1251 = v1249
	goto L343
L342:
	;
	v1251 = v1236
	goto L343
L343:
	;
	v1252 = int32(1)
	v1255 = v1234 + v1252
	if v1255 != v1178 {
		v1234 = v1255
		v1236 = v1251
		v1237 = v1237 + v1252
		goto L339
	} else {
		goto L344
	}
L344:
	;
	goto L340
L345:
	;
	v1302 = v1296
	goto L347
L346:
	;
	v1302 = int32(0)
	goto L347
L347:
	;
	v1304 = v1121 + int32(4)
	if base.Ui32(v1304) < base.Ui32(v1150+v1151<<(uint(int32(2))%32)) {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v1310 = v1304
	goto L350
L349:
	;
	v1310 = int32(0)
	goto L350
L350:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v1314 = v1310
	v1315 = v1302
	v1317 = v1311
	goto L303
L351:
	;
	goto L300
L352:
	;
	v1351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)))
	if v1351 != 0 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v1352 = v1348
	goto L355
L354:
	;
	v1352 = int32(0)
	goto L355
L355:
	;
	v1371 = v1352
	goto L285
L356:
	;
	m.G0 = v20 + int32(16)
	return
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v1373
	goto L356
L358:
	;
	goto L359
L359:
	;
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v1375 != 0 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v1375
	goto L356
L361:
	;
	goto L362
L362:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v1377 != 0 {
		goto L367
	} else {
		goto L368
	}
L363:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v1389
	goto L356
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v1377
	goto L356
L365:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1389)+4))
	if v1390 <= v1391 {
		goto L363
	} else {
		goto L376
	}
L366:
	;
	if v1371 != 0 {
		goto L373
	} else {
		goto L374
	}
L367:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+4))
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v1379 != 0 {
		v1389 = v1379
		v1390 = v1378
		goto L365
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v1383 != 0 {
		v1389 = v1383
		v1390 = int32(0)
		goto L365
	} else {
		goto L372
	}
L370:
	;
	if int32(0) < v1378 {
		goto L364
	} else {
		goto L371
	}
L371:
	;
	goto L366
L372:
	;
	goto L366
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v1371
	goto L356
L374:
	;
	goto L375
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = int32(0)
	goto L356
L376:
	;
	goto L364
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
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
			v32 = int32(1)
			v33 = v18 - v32
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+20)))
			if v35&v32 == v30 {
				v44 = v17 + v33<<(uint(int32(4))%32) + int32(20)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
				if int32(0) <= v45 {
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
					v50 = v34 + v48 + v45
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+6)))
					if v51 != int32(1) {
						v100 = v50
						m.G0 = v10 + int32(16)
						return v100
					} else {
						v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v44)+4)))
						switch v54&int32(65535) - int32(1) {
						case 0:
							v59 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50))))
							v100 = v59
							m.G0 = v10 + int32(16)
							return v100
						case 1:
							v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50))))
							v100 = v60
							m.G0 = v10 + int32(16)
							return v100
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v54
								F_errmsg_internal(m, int32(464696), v10)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(313097), int32(70), int32(65199))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 3:
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
							v100 = v61
							m.G0 = v10 + int32(16)
							return v100
						}
					}
				} else {
					v75 = F_nocachegetattr(m, v16, v18, v17)
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						v100 = v75
						m.G0 = v10 + int32(16)
						return v100
					}
				}
			} else {
				v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+int32(base.Ui32(v33)>>(uint(int32(3))%32)))+23)))
				if int32(base.Ui32(v80)>>(uint(v33&int32(7))%32))&int32(1) == int32(0) {
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
		v14 = *(*int32)(unsafe.Add(mBase, _consts[353]))
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
		return int32(514012)
	default:
		v13 = int32(525707)
		return v13
	case 8:
		return int32(510589)
	case 11:
		v13 = int32(510723)
		return v13
	case 19:
		return int32(523871)
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
	var v50 int32
	_ = v50
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
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(646240)
					F_errmsg(m, int32(241575), v10)
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						F_errhint(m, int32(538820), int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(478903), int32(1719), int32(208779))
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
								v50 = v28 + int32(1)
								v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
								if v51 != 0 {
									v28 = v50
									v30 = v51
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
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
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
	return v159 - l0
L2:
	;
	v111 = F__emscripten_memset_bulkmem(m, v8, base.I32_extend8_s(int32(0)), int32(32))
	mBase = m.M
	goto L31
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
	v159 = v107
	goto L1
L8:
	;
	v107 = v97
	goto L7
L9:
	;
	v87 = v82
	goto L27
L10:
	;
	v82 = v74
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
	v72 = F_strlen(m, l0)
	mBase = m.M
	v107 = v72 + l0
	goto L7
L14:
	;
	v21 = l0
	goto L17
L15:
	;
	v34 = l0
	goto L16
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v43 = int32(-2139062144)
	if (int32(16843008)-v40|v40)&v43 != v43 {
		v74 = v34
		goto L10
	} else {
		goto L22
	}
L17:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v26 == int32(0) {
		v97 = v21
		goto L8
	} else {
		goto L19
	}
L18:
	;
	v34 = v31
	goto L16
L19:
	;
	if v10&int32(255) == v26 {
		v97 = v21
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v31 = v21 + int32(1)
	if v31&int32(3) != 0 {
		v21 = v31
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	v49 = v34
	v51 = v40
	goto L23
L23:
	;
	v55 = v51 ^ v16*int32(16843009)
	v58 = int32(-2139062144)
	if (int32(16843008)-v55|v55)&v58 != v58 {
		v74 = v49
		goto L10
	} else {
		goto L25
	}
L24:
	;
	v82 = v64
	goto L9
L25:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v64 = v49 + int32(4)
	v68 = int32(-2139062144)
	if (v62|(int32(16843008)-v62))&v68 == v68 {
		v49 = v64
		v51 = v62
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v89 == int32(0) {
		v97 = v87
		goto L8
	} else {
		goto L29
	}
L28:
	;
	v97 = v87
	goto L8
L29:
	;
	if v89 != v10&int32(255) {
		v87 = v87 + int32(1)
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v112 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v114 = l1
	v115 = v112
	goto L35
L33:
	;
	goto L34
L34:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v136 == int32(0) {
		v159 = l0
		goto L1
	} else {
		goto L38
	}
L35:
	;
	v122 = v8 + int32(base.Ui32(v115)>>(uint(int32(3))%32))&int32(28)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v124 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v123 | v124<<(uint(v115)%32)
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
	if v128 != 0 {
		v114 = v114 + v124
		v115 = v128
		goto L35
	} else {
		goto L37
	}
L36:
	;
	goto L34
L37:
	;
	goto L36
L38:
	;
	v140 = l0
	v141 = v136
	goto L39
L39:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(base.Ui32(v141)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v149)>>(uint(v141)%32))&int32(1) != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v159 = v155
	goto L1
L41:
	;
	v159 = v140
	goto L1
L42:
	;
	goto L43
L43:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+1)))
	v155 = v140 + int32(1)
	if v153 != 0 {
		v140 = v155
		v141 = v153
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L40
}
func F_strict_word_similarity_dist_commutator_op(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
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
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
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
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 float32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v19 = int32(1)
			v20 = v15 + v19
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			v23 = v21 & v19
			if v21 == v19 {
				v26 = int32(4)
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				if v28&int32(254) == int32(2) {
					v37 = v26
				} else {
					v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
				}
				if v28 == int32(1) {
					v40 = v26
				} else {
					v40 = v37
				}
				v51 = v40
			} else {
				v41 = int32(1)
				if v23 != 0 {
					v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if v23 != 0 {
				v52 = v20
			} else {
				v52 = v15 + int32(4)
			}
			v53 = int32(1)
			v54 = v10 + v53
			v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v59 = v57 & v53
			if v59 != 0 {
				v60 = v54
			} else {
				v60 = v10 + int32(4)
			}
			if v57 == int32(1) {
				v63 = int32(4)
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				if v65&int32(254) == int32(2) {
					v74 = v63
				} else {
					v74 = base.B2i32(v65 == int32(18)) << (uint(v63) % 32)
				}
				if v65 == int32(1) {
					v77 = v63
				} else {
					v77 = v74
				}
				v88 = v77
			} else {
				v78 = int32(1)
				if v59 != 0 {
					v88 = int32(base.Ui32(v57)>>(uint(v78)%32)) - v78
				} else {
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v88 = int32(base.Ui32(v82)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v90 = F_calc_word_similarity(m, v52, v51, v60, v88, int32(2))
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int32(0)
			} else {
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v92 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v96 != v15 {
							F_pfree(m, v15)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								return base.I32_reinterpret_f32(base.F32_sub(float32(1), v90))
							}
						} else {
							return base.I32_reinterpret_f32(base.F32_sub(float32(1), v90))
						}
					}
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v96 != v15 {
						F_pfree(m, v15)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							return base.I32_reinterpret_f32(base.F32_sub(float32(1), v90))
						}
					} else {
						return base.I32_reinterpret_f32(base.F32_sub(float32(1), v90))
					}
				}
			}
		}
	}
}
func F_strlcpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	if l2 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v111&int32(3) == int32(0) {
		v138 = v111
		goto L34
	} else {
		goto L35
	}
L2:
	;
	v111 = l1
	v112 = l0
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
	v108 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v105))) = uint8(v108)
	v111 = v104
	v112 = v105
	goto L1
L6:
	;
	v89 = v84
	v90 = v85
	v91 = v86
	goto L28
L7:
	;
	if v79 == int32(0) {
		v104 = v77
		v105 = v78
		goto L5
	} else {
		goto L27
	}
L8:
	;
	v77 = l1
	v78 = l0
	v79 = v9
	goto L7
L9:
	;
	goto L10
L10:
	;
	v13 = int32(0)
	if l1&int32(3) == v13 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v46 == int32(0) {
		v104 = v43
		v105 = v44
		goto L5
	} else {
		goto L20
	}
L12:
	;
	v43 = l1
	v44 = l0
	v45 = v9
	v46 = base.B2i32(v9 != v13)
	goto L11
L13:
	;
	if v9 == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v22 = l1
	v23 = l0
	v24 = v9
	goto L15
L15:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v26)
	if v26 == int32(0) {
		v84 = v22
		v85 = v23
		v86 = v24
		goto L6
	} else {
		goto L17
	}
L16:
	;
	v43 = v37
	v44 = v31
	v45 = v33
	v46 = v35
	goto L11
L17:
	;
	v30 = int32(1)
	v31 = v23 + v30
	v33 = v24 - v30
	v34 = int32(0)
	v35 = base.B2i32(v33 != v34)
	v37 = v22 + v30
	if v37&int32(3) == v34 {
		v43 = v37
		v44 = v31
		v45 = v33
		v46 = v35
		goto L11
	} else {
		goto L18
	}
L18:
	;
	if v33 != 0 {
		v22 = v37
		v23 = v31
		v24 = v33
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v49 == int32(0) {
		v77 = v43
		v78 = v44
		v79 = v45
		goto L7
	} else {
		goto L21
	}
L21:
	;
	if base.Ui32(v45) < base.Ui32(int32(4)) {
		v77 = v43
		v78 = v44
		v79 = v45
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v55 = v43
	v56 = v44
	v57 = v45
	goto L23
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v63 = int32(-2139062144)
	if (int32(16843008)-v60|v60)&v63 != v63 {
		v84 = v55
		v85 = v56
		v86 = v57
		goto L6
	} else {
		goto L25
	}
L24:
	;
	v77 = v71
	v78 = v69
	v79 = v73
	goto L7
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v60
	v68 = int32(4)
	v69 = v56 + v68
	v71 = v55 + v68
	v73 = v57 - v68
	if base.Ui32(int32(3)) < base.Ui32(v73) {
		v55 = v71
		v56 = v69
		v57 = v73
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v84 = v77
	v85 = v78
	v86 = v79
	goto L6
L28:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	*(*uint8)(unsafe.Add(mBase, uint32(v90))) = uint8(v93)
	if v93 == int32(0) {
		v104 = v89
		v105 = v90
		goto L5
	} else {
		goto L30
	}
L29:
	;
	v104 = v100
	v105 = v98
	goto L5
L30:
	;
	v97 = int32(1)
	v98 = v90 + v97
	v100 = v89 + v97
	v102 = v91 - v97
	if v102 != 0 {
		v89 = v100
		v90 = v98
		v91 = v102
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	return v171 + (v112 - l0)
L33:
	;
	v171 = v163 - v111
	goto L32
L34:
	;
	v142 = v138
	goto L43
L35:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	if v122 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v171 = int32(0)
	goto L32
L37:
	;
	goto L38
L38:
	;
	v127 = v111
	goto L39
L39:
	;
	v131 = v127 + int32(1)
	if v131&int32(3) == int32(0) {
		v138 = v131
		goto L34
	} else {
		goto L41
	}
L40:
	;
	v163 = v131
	goto L33
L41:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v136 != 0 {
		v127 = v131
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v151 = int32(-2139062144)
	if (int32(16843008)-v148|v148)&v151 == v151 {
		v142 = v142 + int32(4)
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v157 = v142
	goto L46
L45:
	;
	goto L44
L46:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v161 != 0 {
		v157 = v157 + int32(1)
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v163 = v157
	goto L33
L48:
	;
	goto L47
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
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v289 int32
	_ = v289
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v480 int32
	_ = v480
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v512 int32
	_ = v512
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	v9 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10*int32(28))+uint32(_consts[355])))
	goto L2
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L27
	} else {
		goto L201
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
		goto L159
	} else {
		goto L160
	}
L6:
	;
	if l2&int32(3) == int32(0) {
		v44 = l2
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v78 = l3
	goto L8
L8:
	;
	v80 = v78 + int32(1)
	if base.Ui32(int32(536870912)) <= base.Ui32(v80) {
		goto L1
	} else {
		goto L26
	}
L9:
	;
	v78 = v77
	goto L8
L10:
	;
	v77 = v69 - l2
	goto L9
L11:
	;
	v48 = v44
	goto L20
L12:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v28 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v77 = int32(0)
	goto L9
L14:
	;
	goto L15
L15:
	;
	v33 = l2
	goto L16
L16:
	;
	v37 = v33 + int32(1)
	if v37&int32(3) == int32(0) {
		v44 = v37
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v69 = v37
	goto L10
L18:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v42 != 0 {
		v33 = v37
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v57 = int32(-2139062144)
	if (int32(16843008)-v54|v54)&v57 == v57 {
		v48 = v48 + int32(4)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v63 = v48
	goto L23
L22:
	;
	goto L21
L23:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v67 != 0 {
		v63 = v63 + int32(1)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v69 = v63
	goto L10
L25:
	;
	goto L24
L26:
	;
	v85 = F_palloc(m, v80<<(uint(int32(2))%32))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	F_char2wchar(m, v85, v80, l2, v78, l4)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v91 = int32(0)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v93 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v96 = v91
	v97 = v93
	goto L33
L31:
	;
	v115 = v91
	goto L32
L32:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122*int32(28))+uint32(_consts[355])))
	goto L37
L33:
	;
	v105 = F_casemap(m, v97, int32(0))
	mBase = m.M
	goto L35
L34:
	;
	v115 = v108
	goto L32
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85+v96<<(uint(int32(2))%32)))) = v105
	v108 = v96 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v85+v108<<(uint(int32(2))%32))))
	if v112 != 0 {
		v96 = v108
		v97 = v112
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v130 = v127*v115 + int32(1)
	v131 = F_palloc(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L27
	} else {
		goto L38
	}
L38:
	;
	if v130 == int32(0) {
		v519 = v91
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if base.Ui32(v519+int32(1)) <= base.Ui32(l1) {
		goto L150
	} else {
		goto L151
	}
L40:
	;
	if l4 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v137 = int32(0)
	v143 = m.G0
	v144 = int32(16)
	v145 = v143 - v144
	m.G0 = v145
	*(*int32)(unsafe.Add(mBase, uint32(v145)+12)) = v85
	v149 = v145 + int32(12)
	v150 = m.G0
	v152 = v150 - v144
	m.G0 = v152
	if v131 != 0 {
		goto L49
	} else {
		goto L50
	}
L42:
	;
	goto L43
L43:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v317 = *(*int32)(unsafe.Add(mBase, _consts[1118]))
	if v314 != 0 {
		goto L88
	} else {
		goto L89
	}
L44:
	;
	v519 = v304
	goto L39
L45:
	;
	v308 = int32(16)
	m.G0 = v152 + v308
	m.G0 = v145 + v308
	goto L44
L46:
	;
	v304 = v130 - v289
	goto L45
L47:
	;
	if v226 != 0 {
		goto L72
	} else {
		goto L73
	}
L48:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v185 = v130
	v186 = v131
	v190 = v184
	goto L61
L49:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v130) {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v157 == int32(0) {
		v304 = v137
		goto L45
	} else {
		goto L53
	}
L52:
	;
	v226 = v130
	v227 = v131
	goto L47
L53:
	;
	v160 = v157
	v161 = v156
	v163 = v137
	goto L54
L54:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v160) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v304 = v183
	goto L45
L56:
	;
	v172 = int32(-1)
	v175 = F_wcrtomb(m, v152+int32(12), v160)
	mBase = m.M
	if v175 == v172 {
		v304 = v172
		goto L45
	} else {
		goto L59
	}
L57:
	;
	v178 = int32(1)
	goto L58
L58:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	v183 = v163 + v178
	if v180 != 0 {
		v160 = v180
		v161 = v161 + int32(4)
		v163 = v183
		goto L54
	} else {
		goto L60
	}
L59:
	;
	v178 = v175
	goto L58
L60:
	;
	goto L55
L61:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	if base.Ui32(v194-int32(128)) <= base.Ui32(int32(-128)) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v226 = v216
	v227 = v219
	goto L47
L63:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v222 = v220 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = v222
	if base.Ui32(int32(3)) < base.Ui32(v216) {
		v185 = v216
		v186 = v219
		v190 = v222
		goto L61
	} else {
		goto L71
	}
L64:
	;
	if v194 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v186))) = uint8(v194)
	v212 = int32(1)
	v216 = v185 - v212
	v219 = v186 + v212
	goto L63
L67:
	;
	v201 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v186))) = uint8(v201)
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = v201
	v289 = v185
	goto L46
L68:
	;
	goto L69
L69:
	;
	v205 = int32(-1)
	v206 = F_wcrtomb(m, v186, v194)
	mBase = m.M
	if v206 == v205 {
		v304 = v205
		goto L45
	} else {
		goto L70
	}
L70:
	;
	v216 = v185 - v206
	v219 = v186 + v206
	goto L63
L71:
	;
	goto L62
L72:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v236 = v226
	v237 = v227
	v239 = v235
	goto L75
L73:
	;
	goto L74
L74:
	;
	v304 = v130
	goto L45
L75:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	if base.Ui32(v245-int32(128)) <= base.Ui32(int32(-128)) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L74
L77:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v278 = v276 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = v278
	if v272 != 0 {
		v236 = v272
		v237 = v275
		v239 = v278
		goto L75
	} else {
		goto L86
	}
L78:
	;
	if v245 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v237))) = uint8(v245)
	v268 = int32(1)
	v272 = v236 - v268
	v275 = v237 + v268
	goto L77
L81:
	;
	v252 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v237))) = uint8(v252)
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = v252
	v289 = v236
	goto L46
L82:
	;
	goto L83
L83:
	;
	v256 = int32(-1)
	v259 = F_wcrtomb(m, v152+int32(12), v245)
	mBase = m.M
	if v259 == v256 {
		v304 = v256
		goto L45
	} else {
		goto L84
	}
L84:
	;
	if base.Ui32(v236) < base.Ui32(v259) {
		v289 = v236
		goto L46
	} else {
		goto L85
	}
L85:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	v264 = F_wcrtomb(m, v237, v263)
	mBase = m.M
	v272 = v236 - v259
	v275 = v237 + v259
	goto L77
L86:
	;
	goto L76
L87:
	;
	v328 = int32(0)
	v334 = m.G0
	v335 = int32(16)
	v336 = v334 - v335
	m.G0 = v336
	*(*int32)(unsafe.Add(mBase, uint32(v336)+12)) = v85
	v340 = v336 + int32(12)
	v341 = m.G0
	v343 = v341 - v335
	m.G0 = v343
	if v131 != 0 {
		goto L102
	} else {
		goto L103
	}
L88:
	;
	if v314 == int32(-1) {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	if v317 == int32(4613600) {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	v322 = int32(4613600)
	goto L93
L92:
	;
	v322 = v314
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1118])) = v322
	goto L90
L94:
	;
	v327 = int32(-1)
	goto L96
L95:
	;
	v327 = v317
	goto L96
L96:
	;
	goto L87
L97:
	;
	if v327 != 0 {
		goto L141
	} else {
		goto L142
	}
L98:
	;
	v499 = int32(16)
	m.G0 = v343 + v499
	m.G0 = v336 + v499
	goto L97
L99:
	;
	v495 = v130 - v480
	goto L98
L100:
	;
	if v417 != 0 {
		goto L125
	} else {
		goto L126
	}
L101:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v376 = v130
	v377 = v131
	v381 = v375
	goto L114
L102:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v130) {
		goto L101
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	if v348 == int32(0) {
		v495 = v328
		goto L98
	} else {
		goto L106
	}
L105:
	;
	v417 = v130
	v418 = v131
	goto L100
L106:
	;
	v351 = v348
	v352 = v347
	v354 = v328
	goto L107
L107:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v351) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v495 = v374
	goto L98
L109:
	;
	v363 = int32(-1)
	v366 = F_wcrtomb(m, v343+int32(12), v351)
	mBase = m.M
	if v366 == v363 {
		v495 = v363
		goto L98
	} else {
		goto L112
	}
L110:
	;
	v369 = int32(1)
	goto L111
L111:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	v374 = v354 + v369
	if v371 != 0 {
		v351 = v371
		v352 = v352 + int32(4)
		v354 = v374
		goto L107
	} else {
		goto L113
	}
L112:
	;
	v369 = v366
	goto L111
L113:
	;
	goto L108
L114:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
	if base.Ui32(v385-int32(128)) <= base.Ui32(int32(-128)) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v417 = v407
	v418 = v410
	goto L100
L116:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v413 = v411 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v340))) = v413
	if base.Ui32(int32(3)) < base.Ui32(v407) {
		v376 = v407
		v377 = v410
		v381 = v413
		goto L114
	} else {
		goto L124
	}
L117:
	;
	if v385 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	goto L119
L119:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v377))) = uint8(v385)
	v403 = int32(1)
	v407 = v376 - v403
	v410 = v377 + v403
	goto L116
L120:
	;
	v392 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v377))) = uint8(v392)
	*(*int32)(unsafe.Add(mBase, uint32(v340))) = v392
	v480 = v376
	goto L99
L121:
	;
	goto L122
L122:
	;
	v396 = int32(-1)
	v397 = F_wcrtomb(m, v377, v385)
	mBase = m.M
	if v397 == v396 {
		v495 = v396
		goto L98
	} else {
		goto L123
	}
L123:
	;
	v407 = v376 - v397
	v410 = v377 + v397
	goto L116
L124:
	;
	goto L115
L125:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v427 = v417
	v428 = v418
	v430 = v426
	goto L128
L126:
	;
	goto L127
L127:
	;
	v495 = v130
	goto L98
L128:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	if base.Ui32(v436-int32(128)) <= base.Ui32(int32(-128)) {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	goto L127
L130:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v469 = v467 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v340))) = v469
	if v463 != 0 {
		v427 = v463
		v428 = v466
		v430 = v469
		goto L128
	} else {
		goto L139
	}
L131:
	;
	if v436 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L133
L133:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v428))) = uint8(v436)
	v459 = int32(1)
	v463 = v427 - v459
	v466 = v428 + v459
	goto L130
L134:
	;
	v443 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v428))) = uint8(v443)
	*(*int32)(unsafe.Add(mBase, uint32(v340))) = v443
	v480 = v427
	goto L99
L135:
	;
	goto L136
L136:
	;
	v447 = int32(-1)
	v450 = F_wcrtomb(m, v343+int32(12), v436)
	mBase = m.M
	if v450 == v447 {
		v495 = v447
		goto L98
	} else {
		goto L137
	}
L137:
	;
	if base.Ui32(v427) < base.Ui32(v450) {
		v480 = v427
		goto L99
	} else {
		goto L138
	}
L138:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	v455 = F_wcrtomb(m, v428, v454)
	mBase = m.M
	v463 = v427 - v450
	v466 = v428 + v450
	goto L130
L139:
	;
	goto L129
L140:
	;
	v519 = v495
	goto L39
L141:
	;
	if v327 == int32(-1) {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	goto L143
L143:
	;
	goto L147
L144:
	;
	v512 = int32(4613600)
	goto L146
L145:
	;
	v512 = v327
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1118])) = v512
	goto L143
L147:
	;
	goto L149
L149:
	;
	goto L140
L150:
	;
	if v519 != 0 {
		goto L154
	} else {
		goto L155
	}
L151:
	;
	goto L152
L152:
	;
	F_pfree(m, v85)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L27
	} else {
		goto L157
	}
L153:
	;
	v526 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v524+v519))) = uint8(v526)
	goto L152
L154:
	;
	v523 = F__emscripten_memcpy_bulkmem(m, l0, v131, v519)
	mBase = m.M
	v524 = v523
	goto L156
L155:
	;
	v524 = l0
	goto L156
L156:
	;
	goto L153
L157:
	;
	F_pfree(m, v131)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L27
	} else {
		goto L158
	}
L158:
	;
	return v519
L159:
	;
	if l2&int32(3) == int32(0) {
		v558 = l2
		goto L164
	} else {
		goto L165
	}
L160:
	;
	v592 = l3
	goto L161
L161:
	;
	if base.Ui32(l1) < base.Ui32(v592+int32(1)) {
		goto L179
	} else {
		goto L180
	}
L162:
	;
	v592 = v591
	goto L161
L163:
	;
	v591 = v583 - l2
	goto L162
L164:
	;
	v562 = v558
	goto L173
L165:
	;
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v542 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v591 = int32(0)
	goto L162
L167:
	;
	goto L168
L168:
	;
	v547 = l2
	goto L169
L169:
	;
	v551 = v547 + int32(1)
	if v551&int32(3) == int32(0) {
		v558 = v551
		goto L164
	} else {
		goto L171
	}
L170:
	;
	v583 = v551
	goto L163
L171:
	;
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551))))
	if v556 != 0 {
		v547 = v551
		goto L169
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v562)))
	v571 = int32(-2139062144)
	if (int32(16843008)-v568|v568)&v571 == v571 {
		v562 = v562 + int32(4)
		goto L173
	} else {
		goto L175
	}
L174:
	;
	v577 = v562
	goto L176
L175:
	;
	goto L174
L176:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577))))
	if v581 != 0 {
		v577 = v577 + int32(1)
		goto L176
	} else {
		goto L178
	}
L177:
	;
	v583 = v577
	goto L163
L178:
	;
	goto L177
L179:
	;
	return v592
L180:
	;
	if v592 != 0 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v600 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v598+v592))) = uint8(v600)
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598))))
	if v602 == v600 {
		goto L179
	} else {
		goto L185
	}
L182:
	;
	v597 = F__emscripten_memcpy_bulkmem(m, l0, l2, v592)
	mBase = m.M
	v598 = v597
	goto L184
L183:
	;
	v598 = l0
	goto L184
L184:
	;
	goto L181
L185:
	;
	v605 = l0
	v607 = v602
	goto L186
L186:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v612 == int32(1) {
		goto L189
	} else {
		goto L190
	}
L187:
	;
	goto L179
L188:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v605))) = uint8(v635)
	v638 = v605 + int32(1)
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638))))
	if v639 != 0 {
		v605 = v638
		v607 = v639
		goto L186
	} else {
		goto L200
	}
L189:
	;
	v615 = int32(255)
	v616 = v607 & v615
	if base.Ui32((v616-int32(65))&v615) < base.Ui32(int32(26)) {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	goto L191
L191:
	;
	v627 = v607 & int32(255)
	if base.Ui32(v627-int32(65)) < base.Ui32(int32(26)) {
		goto L197
	} else {
		goto L198
	}
L192:
	;
	v635 = v625
	goto L188
L193:
	;
	v625 = v616 | int32(32)
	goto L195
L194:
	;
	v625 = v616
	goto L195
L195:
	;
	goto L192
L196:
	;
	v635 = v634
	goto L188
L197:
	;
	v634 = v627 | int32(32)
	goto L199
L198:
	;
	v634 = v627
	goto L199
L199:
	;
	goto L196
L200:
	;
	goto L187
L201:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L27
	} else {
		goto L202
	}
L202:
	;
	F_errmsg(m, int32(12915), int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L27
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(480742), int32(207), int32(483492))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L27
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_strnxfrm_libc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v10 = m.G0
	v12 = v10 - int32(1024)
	m.G0 = v12
	v15 = l3 + int32(1)
	if v15 == int32(0) {
		v19 = F_strlen(m, l2)
		mBase = m.M
		if base.Ui32(v19) < base.Ui32(l1) {
			v21 = F_strcpy(m, l0, l2)
			mBase = m.M
		} else {
		}
		v42 = v19
		m.G0 = v12 + int32(1024)
		return v42
	} else {
		if base.Ui32(int32(1025)) <= base.Ui32(v15) {
			v24 = F_palloc(m, v15)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = v24
				if l3 != 0 {
					v29 = F__emscripten_memcpy_bulkmem(m, v28, l2, l3)
					mBase = m.M
					v30 = v29
				} else {
					v30 = v28
				}
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v30+l3))) = uint8(v32)
				v35 = F_strlen(m, v30)
				mBase = m.M
				if base.Ui32(v35) < base.Ui32(l1) {
					v37 = F_strcpy(m, l0, v30)
					mBase = m.M
				} else {
				}
				if v30 == v12 {
					v42 = v35
					m.G0 = v12 + int32(1024)
					return v42
				} else {
					F_pfree(m, v30)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						v42 = v35
						m.G0 = v12 + int32(1024)
						return v42
					}
				}
			}
		} else {
			v28 = v12
			if l3 != 0 {
				v29 = F__emscripten_memcpy_bulkmem(m, v28, l2, l3)
				mBase = m.M
				v30 = v29
			} else {
				v30 = v28
			}
			v32 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v30+l3))) = uint8(v32)
			v35 = F_strlen(m, v30)
			mBase = m.M
			if base.Ui32(v35) < base.Ui32(l1) {
				v37 = F_strcpy(m, l0, v30)
				mBase = m.M
			} else {
			}
			if v30 == v12 {
				v42 = v35
				m.G0 = v12 + int32(1024)
				return v42
			} else {
				F_pfree(m, v30)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v42 = v35
					m.G0 = v12 + int32(1024)
					return v42
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
	var v120 int64
	_ = v120
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
	var v236 int32
	_ = v236
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v245 int64
	_ = v245
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
	return v245
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
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(28)
	v245 = int64(0)
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
	v120 = int64(0)
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
	v157 = int64(base.Ui64(v120) >> (uint(v156) % 64))
	v159 = int64(base.Ui64(v106) >> (uint(v156) % 64))
	v162 = int64(4294967295)
	v163 = v120 & v162
	v165 = v106 & v162
	v166 = v163 * v165
	v170 = int64(base.Ui64(v166)>>(uint(v156)%64)) + v163*v159
	v177 = v165*v157 + v170&v162
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v106*v150 + v150*v120 + v157*v159 + int64(base.Ui64(v170)>>(uint(v156)%64)) + int64(base.Ui64(v177)>>(uint(v156)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v166&v162 | v177<<(uint(v156)%64)
	goto L41
L41:
	;
	v188 = int32(1)
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	if v189 != int64(0) {
		v201 = v188
		v202 = v117
		v203 = v120
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v110 = v201
	v112 = v112 + int32(1)
	v117 = v202
	v120 = v203
	goto L31
L43:
	;
	v192 = v106 * v120
	v195 = base.I64_extend_i32_u(v146) & int64(255)
	if base.Ui64(v195^int64(-1)) < base.Ui64(v192) {
		v201 = v188
		v202 = v117
		v203 = v120
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
	v239 = base.I64_extend_i32_s(v236)
	v245 = v238 ^ v239 - v239
	goto L1
L52:
	;
	if base.I32_wrap_i64(v224) != 0 {
		goto L60
	} else {
		goto L61
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(68)
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
	if base.Ui64(v120) < base.Ui64(l3) {
		v236 = v73
		v238 = v120
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
	v224 = v216
	v225 = l3
	goto L52
L59:
	;
	v223 = v73
	v224 = l3 & int64(1)
	v225 = v120
	goto L52
L60:
	;
	if base.Ui64(v225) <= base.Ui64(l3) {
		v236 = v223
		v238 = v225
		goto L51
	} else {
		goto L63
	}
L61:
	;
	if v223 != 0 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(68)
	v245 = l3 - int64(1)
	goto L1
L63:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(68)
	v245 = l3
	goto L1
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
				v52 = F_expression_tree_mutator_impl(m, l0, int32(874), l1)
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
				v23 = F_query_tree_mutator_impl(m, l0, int32(874), l1, int32(0))
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
				v52 = F_expression_tree_mutator_impl(m, l0, int32(874), l1)
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
						F_errmsg_internal(m, int32(468923), v7)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(475017), int32(5399), int32(199947))
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
							F_errmsg_internal(m, int32(468923), v7)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(475017), int32(5399), int32(199947))
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	v4 = m.G0
	v6 = v4 - int32(1040)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l1
	v14 = F_pg_snprintf(m, v6+int32(16), int32(1024), int32(118790), v6)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[657]))
		if v17 == int32(0) {
			v21 = *(*int32)(unsafe.Add(mBase, _consts[658]))
			if v21 != 0 {
				F_pfree(m, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v25 = int64(0)
					*(*int64)(unsafe.Add(mBase, _consts[659])) = v25
					*(*int64)(unsafe.Add(mBase, _consts[657])) = v25
					v31 = *(*int32)(unsafe.Add(mBase, _consts[648]))
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
				*(*int64)(unsafe.Add(mBase, _consts[659])) = v25
				*(*int64)(unsafe.Add(mBase, _consts[657])) = v25
				v31 = *(*int32)(unsafe.Add(mBase, _consts[648]))
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
			v39 = *(*int32)(unsafe.Add(mBase, _consts[648]))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+60))
			v45 = F_BufFileOpenFileSet(m, v40, v6+int32(16), int32(2), int32(1))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				if v45 == int32(0) {
					v50 = *(*int32)(unsafe.Add(mBase, _consts[648]))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+60))
					v54 = F_BufFileCreateFileSet(m, v51, v6+int32(16))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						v56 = v54
						v57 = int32(4359708)
						v58 = *(*int32)(unsafe.Add(mBase, _consts[657]))
						F_BufFileWrite(m, v56, v57, int32(4))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, _consts[658]))
							F_BufFileWrite(m, v56, v64, v58<<(uint(int32(4))%32))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								F_BufFileClose(m, v56)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									v72 = *(*int32)(unsafe.Add(mBase, _consts[658]))
									if v72 != 0 {
										F_pfree(m, v72)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											v76 = int64(0)
											*(*int64)(unsafe.Add(mBase, _consts[659])) = v76
											*(*int64)(unsafe.Add(mBase, _consts[657])) = v76
											m.G0 = v6 + int32(1040)
											return
										}
									} else {
										v76 = int64(0)
										*(*int64)(unsafe.Add(mBase, _consts[659])) = v76
										*(*int64)(unsafe.Add(mBase, _consts[657])) = v76
										m.G0 = v6 + int32(1040)
										return
									}
								}
							}
						}
					}
				} else {
					v56 = v45
					v57 = int32(4359708)
					v58 = *(*int32)(unsafe.Add(mBase, _consts[657]))
					F_BufFileWrite(m, v56, v57, int32(4))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, _consts[658]))
						F_BufFileWrite(m, v56, v64, v58<<(uint(int32(4))%32))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							F_BufFileClose(m, v56)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								v72 = *(*int32)(unsafe.Add(mBase, _consts[658]))
								if v72 != 0 {
									F_pfree(m, v72)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										v76 = int64(0)
										*(*int64)(unsafe.Add(mBase, _consts[659])) = v76
										*(*int64)(unsafe.Add(mBase, _consts[657])) = v76
										m.G0 = v6 + int32(1040)
										return
									}
								} else {
									v76 = int64(0)
									*(*int64)(unsafe.Add(mBase, _consts[659])) = v76
									*(*int64)(unsafe.Add(mBase, _consts[657])) = v76
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
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1219]))
	if v5 == int32(0) {
		if l0 == int32(10) {
			v13 = int32(1)
			v15 = int32(*(*uint8)(unsafe.Add(mBase, _consts[131])))
			if v15&v13 == int32(0) {
				v52 = v13
				return v52 & int32(1)
			} else {
				v23 = F_SearchSysCache1(m, int32(11), l0)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					if v23 != 0 {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
						v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
						v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v28)+68)))
						F_ReleaseCatCache(m, v23)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = v30
							v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1220])))
							if v35 == int32(0) {
								F_CacheRegisterSyscacheCallback(m, int32(11), int32(1783), int32(0))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v44 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[1220])) = uint8(v44)
									*(*int32)(unsafe.Add(mBase, _consts[1219])) = l0
									v50 = v33 & int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[1221])) = uint8(v50)
									v52 = v33
									return v52 & int32(1)
								}
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[1219])) = l0
								v50 = v33 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[1221])) = uint8(v50)
								v52 = v33
								return v52 & int32(1)
							}
						}
					} else {
						v33 = int32(0)
						v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1220])))
						if v35 == int32(0) {
							F_CacheRegisterSyscacheCallback(m, int32(11), int32(1783), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v44 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[1220])) = uint8(v44)
								*(*int32)(unsafe.Add(mBase, _consts[1219])) = l0
								v50 = v33 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[1221])) = uint8(v50)
								v52 = v33
								return v52 & int32(1)
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[1219])) = l0
							v50 = v33 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[1221])) = uint8(v50)
							v52 = v33
							return v52 & int32(1)
						}
					}
				}
			}
		} else {
			v23 = F_SearchSysCache1(m, int32(11), l0)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v23 != 0 {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v28)+68)))
					F_ReleaseCatCache(m, v23)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = v30
						v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1220])))
						if v35 == int32(0) {
							F_CacheRegisterSyscacheCallback(m, int32(11), int32(1783), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v44 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[1220])) = uint8(v44)
								*(*int32)(unsafe.Add(mBase, _consts[1219])) = l0
								v50 = v33 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[1221])) = uint8(v50)
								v52 = v33
								return v52 & int32(1)
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[1219])) = l0
							v50 = v33 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[1221])) = uint8(v50)
							v52 = v33
							return v52 & int32(1)
						}
					}
				} else {
					v33 = int32(0)
					v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1220])))
					if v35 == int32(0) {
						F_CacheRegisterSyscacheCallback(m, int32(11), int32(1783), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v44 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[1220])) = uint8(v44)
							*(*int32)(unsafe.Add(mBase, _consts[1219])) = l0
							v50 = v33 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[1221])) = uint8(v50)
							v52 = v33
							return v52 & int32(1)
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[1219])) = l0
						v50 = v33 & int32(1)
						*(*uint8)(unsafe.Add(mBase, _consts[1221])) = uint8(v50)
						v52 = v33
						return v52 & int32(1)
					}
				}
			}
		}
	} else {
		if l0 != v5 {
			if l0 == int32(10) {
				v13 = int32(1)
				v15 = int32(*(*uint8)(unsafe.Add(mBase, _consts[131])))
				if v15&v13 == int32(0) {
					v52 = v13
					return v52 & int32(1)
				} else {
					v23 = F_SearchSysCache1(m, int32(11), l0)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if v23 != 0 {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
							v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
							v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v28)+68)))
							F_ReleaseCatCache(m, v23)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = v30
								v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1220])))
								if v35 == int32(0) {
									F_CacheRegisterSyscacheCallback(m, int32(11), int32(1783), int32(0))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										v44 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _consts[1220])) = uint8(v44)
										*(*int32)(unsafe.Add(mBase, _consts[1219])) = l0
										v50 = v33 & int32(1)
										*(*uint8)(unsafe.Add(mBase, _consts[1221])) = uint8(v50)
										v52 = v33
										return v52 & int32(1)
									}
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[1219])) = l0
									v50 = v33 & int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[1221])) = uint8(v50)
									v52 = v33
									return v52 & int32(1)
								}
							}
						} else {
							v33 = int32(0)
							v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1220])))
							if v35 == int32(0) {
								F_CacheRegisterSyscacheCallback(m, int32(11), int32(1783), int32(0))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v44 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[1220])) = uint8(v44)
									*(*int32)(unsafe.Add(mBase, _consts[1219])) = l0
									v50 = v33 & int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[1221])) = uint8(v50)
									v52 = v33
									return v52 & int32(1)
								}
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[1219])) = l0
								v50 = v33 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[1221])) = uint8(v50)
								v52 = v33
								return v52 & int32(1)
							}
						}
					}
				}
			} else {
				v23 = F_SearchSysCache1(m, int32(11), l0)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					if v23 != 0 {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
						v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
						v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v28)+68)))
						F_ReleaseCatCache(m, v23)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = v30
							v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1220])))
							if v35 == int32(0) {
								F_CacheRegisterSyscacheCallback(m, int32(11), int32(1783), int32(0))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v44 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[1220])) = uint8(v44)
									*(*int32)(unsafe.Add(mBase, _consts[1219])) = l0
									v50 = v33 & int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[1221])) = uint8(v50)
									v52 = v33
									return v52 & int32(1)
								}
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[1219])) = l0
								v50 = v33 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[1221])) = uint8(v50)
								v52 = v33
								return v52 & int32(1)
							}
						}
					} else {
						v33 = int32(0)
						v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1220])))
						if v35 == int32(0) {
							F_CacheRegisterSyscacheCallback(m, int32(11), int32(1783), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v44 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[1220])) = uint8(v44)
								*(*int32)(unsafe.Add(mBase, _consts[1219])) = l0
								v50 = v33 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[1221])) = uint8(v50)
								v52 = v33
								return v52 & int32(1)
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[1219])) = l0
							v50 = v33 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[1221])) = uint8(v50)
							v52 = v33
							return v52 & int32(1)
						}
					}
				}
			}
		} else {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1221])))
			v52 = v10
			return v52 & int32(1)
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
	F_errmsg(m, int32(214793), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(475774), int32(41), int32(214579))
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
	F_errmsg(m, int32(342092), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L36
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(475774), int32(47), int32(214579))
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
	F_errmsg(m, int32(342153), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L36
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(475774), int32(53), int32(214579))
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
	F_errmsg(m, int32(28502), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L36
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(475774), int32(59), int32(214579))
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
func F_syncrep_yyerror(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v13 int32
	_ = v13
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
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
		if v13 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v12
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
			v19 = F_psprintf(m, int32(664525), v7+int32(16))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v25 = v19
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v25
				m.G0 = v7 + int32(32)
				return
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l2
			v23 = F_psprintf(m, int32(62700), v7)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v25 = v23
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v25
				m.G0 = v7 + int32(32)
				return
			}
		}
	} else {
		m.G0 = v7 + int32(32)
		return
	}
}
