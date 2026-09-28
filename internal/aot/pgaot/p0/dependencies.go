package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dependencies_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
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
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v227 float64
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	v7 = m.G0
	v9 = v7 - int32(224)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_dependencies_scalar[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+216)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_dependencies_scalar[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+208)) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v17 - int32(4) {
	case 0:
		v22 = F_pg_strtoint16_safe(m, l1, v9+int32(208))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+212)))
			if v26 == int32(1) {
				v29 = int32(23)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v31 = F_errsave_start(m, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					if v31 == int32(0) {
						v287 = v29
						m.G0 = v9 + int32(224)
						return v287
					} else {
						F_errcode(m, int32(33685634))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v38
							F_errmsg(m, int32(_a_F_dependencies_scalar_0), v9+int32(32))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(_a_F_dependencies_scalar_1)
								v50 = F_errdetail(m, int32(_a_F_dependencies_scalar_2), v9+int32(16))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									F_errsave_finish(m, v30, int32(_a_F_dependencies_scalar_3), int32(494), int32(_a_F_dependencies_scalar_4))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										v287 = v29
										m.G0 = v9 + int32(224)
										return v287
									}
								}
							}
						}
					}
				}
			} else {
				if int32(-9) < v22 {
					v60 = v22
				} else {
					v60 = int32(0)
				}
				if v60 == int32(0) {
					v63 = int32(23)
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v65 = F_errsave_start(m, v64)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						if v65 == int32(0) {
							v287 = v63
							m.G0 = v9 + int32(224)
							return v287
						} else {
							F_errcode(m, int32(33685634))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v72
								F_errmsg(m, int32(_a_F_dependencies_scalar_0), v9-int32(-64))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v22
									*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = int32(_a_F_dependencies_scalar_1)
									v85 = F_errdetail(m, int32(_a_F_dependencies_scalar_5), v9+int32(48))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										F_errsave_finish(m, v64, int32(_a_F_dependencies_scalar_3), int32(508), int32(_a_F_dependencies_scalar_4))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											v287 = v63
											m.G0 = v9 + int32(224)
											return v287
										}
									}
								}
							}
						}
					}
				} else {
					v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v92 == int32(0) {
						v144 = F_lappend_int(m, v92, v22)
						mBase = m.M
						v145 = m.ExcPending
						if v145 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v144
							v287 = int32(0)
							m.G0 = v9 + int32(224)
							return v287
						}
					} else {
						v95 = int32(_a_F_dependencies_scalar_6)
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
						v104 = *(*int32)(unsafe.Add(mBase, uint32(v97+v98<<(uint(int32(2))%32)-int32(4))))
						v108 = base.I32_extend16_s(v104)
						if int32(0) < v108 {
							v112 = base.B2i32(base.Ui32(v104&v95) < base.Ui32(v22&v95))
						} else {
							v112 = base.B2i32(v22 < v108)
						}
						if v112 != 0 {
							v144 = F_lappend_int(m, v92, v22)
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v144
								v287 = int32(0)
								m.G0 = v9 + int32(224)
								return v287
							}
						} else {
							v113 = int32(23)
							v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v115 = F_errsave_start(m, v114)
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int32(0)
							} else {
								if v115 == int32(0) {
									v287 = v113
									m.G0 = v9 + int32(224)
									return v287
								} else {
									F_errcode(m, int32(33685634))
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v122
										F_errmsg(m, int32(_a_F_dependencies_scalar_0), v9+int32(96))
										mBase = m.M
										v128 = m.ExcPending
										if v128 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+88)) = v108
											*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v22
											*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = int32(_a_F_dependencies_scalar_1)
											v136 = F_errdetail(m, int32(_a_F_dependencies_scalar_7), v9+int32(80))
											mBase = m.M
											v137 = m.ExcPending
											if v137 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v114, int32(_a_F_dependencies_scalar_3), int32(522), int32(_a_F_dependencies_scalar_4))
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
													return int32(0)
												} else {
													v287 = v113
													m.G0 = v9 + int32(224)
													return v287
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
	case 1:
		v150 = F_pg_strtoint16_safe(m, l1, v9+int32(208))
		mBase = m.M
		v151 = m.ExcPending
		if v151 != 0 {
			return int32(0)
		} else {
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v150)
			v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+212)))
			if v153 == int32(1) {
				v156 = int32(23)
				v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v158 = F_errsave_start(m, v157)
				mBase = m.M
				v159 = m.ExcPending
				if v159 != 0 {
					return int32(0)
				} else {
					if v158 == int32(0) {
						v287 = v156
						m.G0 = v9 + int32(224)
						return v287
					} else {
						F_errcode(m, int32(33685634))
						mBase = m.M
						v164 = m.ExcPending
						if v164 != 0 {
							return int32(0)
						} else {
							v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v165
							F_errmsg(m, int32(_a_F_dependencies_scalar_0), v9+int32(128))
							mBase = m.M
							v171 = m.ExcPending
							if v171 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = int32(_a_F_dependencies_scalar_8)
								v177 = F_errdetail(m, int32(_a_F_dependencies_scalar_2), v9+int32(112))
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									F_errsave_finish(m, v157, int32(_a_F_dependencies_scalar_3), int32(539), int32(_a_F_dependencies_scalar_4))
									mBase = m.M
									v183 = m.ExcPending
									if v183 != 0 {
										return int32(0)
									} else {
										v287 = v156
										m.G0 = v9 + int32(224)
										return v287
									}
								}
							}
						}
					}
				}
			} else {
				if int32(-9) < v150 {
					v187 = v150
				} else {
					v187 = int32(0)
				}
				if v187 == int32(0) {
					v190 = int32(23)
					v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v192 = F_errsave_start(m, v191)
					mBase = m.M
					v193 = m.ExcPending
					if v193 != 0 {
						return int32(0)
					} else {
						if v192 == int32(0) {
							v287 = v190
							m.G0 = v9 + int32(224)
							return v287
						} else {
							F_errcode(m, int32(33685634))
							mBase = m.M
							v198 = m.ExcPending
							if v198 != 0 {
								return int32(0)
							} else {
								v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v199
								F_errmsg(m, int32(_a_F_dependencies_scalar_0), v9+int32(160))
								mBase = m.M
								v205 = m.ExcPending
								if v205 != 0 {
									return int32(0)
								} else {
									v206 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+148)) = v206
									*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = int32(_a_F_dependencies_scalar_8)
									v213 = F_errdetail(m, int32(_a_F_dependencies_scalar_9), v9+int32(144))
									mBase = m.M
									v214 = m.ExcPending
									if v214 != 0 {
										return int32(0)
									} else {
										F_errsave_finish(m, v191, int32(_a_F_dependencies_scalar_3), int32(553), int32(_a_F_dependencies_scalar_4))
										mBase = m.M
										v219 = m.ExcPending
										if v219 != 0 {
											return int32(0)
										} else {
											v287 = v190
											m.G0 = v9 + int32(224)
											return v287
										}
									}
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2)
					v287 = int32(0)
					m.G0 = v9 + int32(224)
					return v287
				}
			}
		}
	case 2:
		v227 = F_float8in_internal(m, l1, int32(0), int32(_a_F_dependencies_scalar_10), l1, v9+int32(208))
		mBase = m.M
		v228 = m.ExcPending
		if v228 != 0 {
			return int32(0)
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v227
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+212)))
			if v230 == int32(1) {
				v233 = int32(23)
				v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v235 = F_errsave_start(m, v234)
				mBase = m.M
				v236 = m.ExcPending
				if v236 != 0 {
					return int32(0)
				} else {
					if v235 == int32(0) {
						v287 = v233
						m.G0 = v9 + int32(224)
						return v287
					} else {
						F_errcode(m, int32(33685634))
						mBase = m.M
						v241 = m.ExcPending
						if v241 != 0 {
							return int32(0)
						} else {
							v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+192)) = v242
							F_errmsg(m, int32(_a_F_dependencies_scalar_0), v9+int32(192))
							mBase = m.M
							v248 = m.ExcPending
							if v248 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = int32(_a_F_dependencies_scalar_11)
								v254 = F_errdetail(m, int32(_a_F_dependencies_scalar_2), v9+int32(176))
								mBase = m.M
								v255 = m.ExcPending
								if v255 != 0 {
									return int32(0)
								} else {
									F_errsave_finish(m, v234, int32(_a_F_dependencies_scalar_3), int32(569), int32(_a_F_dependencies_scalar_4))
									mBase = m.M
									v260 = m.ExcPending
									if v260 != 0 {
										return int32(0)
									} else {
										v287 = v233
										m.G0 = v9 + int32(224)
										return v287
									}
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2)
				v287 = int32(0)
				m.G0 = v9 + int32(224)
				return v287
			}
		}
	default:
		v264 = int32(23)
		v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v266 = F_errsave_start(m, v265)
		mBase = m.M
		v267 = m.ExcPending
		if v267 != 0 {
			return int32(0)
		} else {
			if v266 == int32(0) {
				v287 = v264
				m.G0 = v9 + int32(224)
				return v287
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v272 = m.ExcPending
				if v272 != 0 {
					return int32(0)
				} else {
					v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v273
					F_errmsg(m, int32(_a_F_dependencies_scalar_0), v9)
					mBase = m.M
					v277 = m.ExcPending
					if v277 != 0 {
						return int32(0)
					} else {
						v280 = F_errdetail(m, int32(_a_F_dependencies_scalar_12), int32(0))
						mBase = m.M
						v281 = m.ExcPending
						if v281 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v265, int32(_a_F_dependencies_scalar_3), int32(580), int32(_a_F_dependencies_scalar_4))
							mBase = m.M
							v286 = m.ExcPending
							if v286 != 0 {
								return int32(0)
							} else {
								v287 = v264
								m.G0 = v9 + int32(224)
								return v287
							}
						}
					}
				}
			}
		}
	}
}
