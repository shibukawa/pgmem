package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dependencies_object_end(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
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
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v159 float64
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
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
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	v6 = m.G0
	v8 = v6 - int32(176)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 == int32(2) {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
		if v13 == int32(0) {
			v16 = int32(23)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v18 = F_errsave_start(m, v17)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				if v18 == int32(0) {
					v287 = v16
					m.G0 = v8 + int32(176)
					return v287
				} else {
					F_errcode(m, int32(33685634))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+144)) = v27
						F_errmsg(m, int32(_a_F_dependencies_object_end_0), v8+int32(144))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+128)) = int32(_a_F_dependencies_object_end_1)
							v39 = F_errdetail(m, int32(_a_F_dependencies_object_end_2), v8+int32(128))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, v17, int32(_a_F_dependencies_object_end_3), int32(161), int32(_a_F_dependencies_object_end_4))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v287 = v16
									m.G0 = v8 + int32(176)
									return v287
								}
							}
						}
					}
				}
			}
		} else {
			v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
			if v46 == int32(0) {
				v49 = int32(23)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v51 = F_errsave_start(m, v50)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					if v51 == int32(0) {
						v287 = v49
						m.G0 = v8 + int32(176)
						return v287
					} else {
						F_errcode(m, int32(33685634))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = v58
							F_errmsg(m, int32(_a_F_dependencies_object_end_0), v8+int32(112))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = int32(_a_F_dependencies_object_end_5)
								v70 = F_errdetail(m, int32(_a_F_dependencies_object_end_2), v8+int32(96))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									F_errsave_finish(m, v50, int32(_a_F_dependencies_object_end_3), int32(171), int32(_a_F_dependencies_object_end_4))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v287 = v49
										m.G0 = v8 + int32(176)
										return v287
									}
								}
							}
						}
					}
				}
			} else {
				v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)))
				if v77 == int32(0) {
					v80 = int32(23)
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v82 = F_errsave_start(m, v81)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						if v82 == int32(0) {
							v287 = v80
							m.G0 = v8 + int32(176)
							return v287
						} else {
							F_errcode(m, int32(33685634))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v89
								F_errmsg(m, int32(_a_F_dependencies_object_end_0), v8+int32(80))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = int32(_a_F_dependencies_object_end_6)
									v101 = F_errdetail(m, int32(_a_F_dependencies_object_end_2), v8-int32(-64))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										F_errsave_finish(m, v81, int32(_a_F_dependencies_object_end_3), int32(181), int32(_a_F_dependencies_object_end_4))
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return int32(0)
										} else {
											v287 = v80
											m.G0 = v8 + int32(176)
											return v287
										}
									}
								}
							}
						}
					}
				} else {
					v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v108 != 0 {
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
						if base.Ui32(int32(-8)) < base.Ui32(v109-int32(8)) {
							v143 = int32(1)
							v144 = v109 + v143
							v149 = F_palloc0(m, v144<<(uint(v143)%32)+int32(10))
							mBase = m.M
							v150 = m.ExcPending
							if v150 != 0 {
								return int32(0)
							} else {
								*(*uint16)(unsafe.Add(mBase, uint32(v149)+8)) = uint16(v144)
								v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
								*(*uint16)(unsafe.Add(mBase, uint32(v149+int32(10)+v109<<(uint(int32(1))%32)))) = uint16(v157)
								v159 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
								*(*float64)(unsafe.Add(mBase, uint32(v149))) = v159
								v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
								v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
								*(*uint16)(unsafe.Add(mBase, uint32(v149)+10)) = uint16(v163)
								v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
								if v165 == v163&int32(_a_F_dependencies_object_end_7) {
									v231 = int32(23)
									v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v233 = F_errsave_start(m, v232)
									mBase = m.M
									v234 = m.ExcPending
									if v234 != 0 {
										return int32(0)
									} else {
										if v233 == int32(0) {
											v287 = v231
											m.G0 = v8 + int32(176)
											return v287
										} else {
											F_errcode(m, int32(33685634))
											mBase = m.M
											v239 = m.ExcPending
											if v239 != 0 {
												return int32(0)
											} else {
												v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v240
												F_errmsg(m, int32(_a_F_dependencies_object_end_0), v8+int32(48))
												mBase = m.M
												v246 = m.ExcPending
												if v246 != 0 {
													return int32(0)
												} else {
													v247 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
													*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(_a_F_dependencies_object_end_1)
													*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v247
													*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(_a_F_dependencies_object_end_5)
													v256 = F_errdetail(m, int32(_a_F_dependencies_object_end_8), v8+int32(32))
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v232, int32(_a_F_dependencies_object_end_3), int32(225), int32(_a_F_dependencies_object_end_4))
														mBase = m.M
														v262 = m.ExcPending
														if v262 != 0 {
															return int32(0)
														} else {
															v287 = v231
															m.G0 = v8 + int32(176)
															return v287
														}
													}
												}
											}
										}
									}
								} else {
									if v109 == int32(1) {
										v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v266 = F_lappend(m, v265, v149)
										mBase = m.M
										v267 = m.ExcPending
										if v267 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v266
											v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											F_list_free(m, v269)
											mBase = m.M
											v271 = m.ExcPending
											if v271 != 0 {
												return int32(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = int64(0)
												v274 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v274)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v274
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)) = uint8(v274)
												*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v274)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
												v287 = v274
												m.G0 = v8 + int32(176)
												return v287
											}
										}
									} else {
										v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
										v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
										*(*uint16)(unsafe.Add(mBase, uint32(v149)+12)) = uint16(v173)
										v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
										if v175 == v173&int32(_a_F_dependencies_object_end_7) {
											v231 = int32(23)
											v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v233 = F_errsave_start(m, v232)
											mBase = m.M
											v234 = m.ExcPending
											if v234 != 0 {
												return int32(0)
											} else {
												if v233 == int32(0) {
													v287 = v231
													m.G0 = v8 + int32(176)
													return v287
												} else {
													F_errcode(m, int32(33685634))
													mBase = m.M
													v239 = m.ExcPending
													if v239 != 0 {
														return int32(0)
													} else {
														v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v240
														F_errmsg(m, int32(_a_F_dependencies_object_end_0), v8+int32(48))
														mBase = m.M
														v246 = m.ExcPending
														if v246 != 0 {
															return int32(0)
														} else {
															v247 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
															*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(_a_F_dependencies_object_end_1)
															*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v247
															*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(_a_F_dependencies_object_end_5)
															v256 = F_errdetail(m, int32(_a_F_dependencies_object_end_8), v8+int32(32))
															mBase = m.M
															v257 = m.ExcPending
															if v257 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, v232, int32(_a_F_dependencies_object_end_3), int32(225), int32(_a_F_dependencies_object_end_4))
																mBase = m.M
																v262 = m.ExcPending
																if v262 != 0 {
																	return int32(0)
																} else {
																	v287 = v231
																	m.G0 = v8 + int32(176)
																	return v287
																}
															}
														}
													}
												}
											}
										} else {
											if v109 == int32(2) {
												v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v266 = F_lappend(m, v265, v149)
												mBase = m.M
												v267 = m.ExcPending
												if v267 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v266
													v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													F_list_free(m, v269)
													mBase = m.M
													v271 = m.ExcPending
													if v271 != 0 {
														return int32(0)
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = int64(0)
														v274 = int32(0)
														*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v274)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v274
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)) = uint8(v274)
														*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v274)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
														v287 = v274
														m.G0 = v8 + int32(176)
														return v287
													}
												}
											} else {
												v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
												v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+8))
												*(*uint16)(unsafe.Add(mBase, uint32(v149)+14)) = uint16(v183)
												v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
												if v185 == v183&int32(_a_F_dependencies_object_end_7) {
													v231 = int32(23)
													v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v233 = F_errsave_start(m, v232)
													mBase = m.M
													v234 = m.ExcPending
													if v234 != 0 {
														return int32(0)
													} else {
														if v233 == int32(0) {
															v287 = v231
															m.G0 = v8 + int32(176)
															return v287
														} else {
															F_errcode(m, int32(33685634))
															mBase = m.M
															v239 = m.ExcPending
															if v239 != 0 {
																return int32(0)
															} else {
																v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v240
																F_errmsg(m, int32(_a_F_dependencies_object_end_0), v8+int32(48))
																mBase = m.M
																v246 = m.ExcPending
																if v246 != 0 {
																	return int32(0)
																} else {
																	v247 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
																	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(_a_F_dependencies_object_end_1)
																	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v247
																	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(_a_F_dependencies_object_end_5)
																	v256 = F_errdetail(m, int32(_a_F_dependencies_object_end_8), v8+int32(32))
																	mBase = m.M
																	v257 = m.ExcPending
																	if v257 != 0 {
																		return int32(0)
																	} else {
																		F_errsave_finish(m, v232, int32(_a_F_dependencies_object_end_3), int32(225), int32(_a_F_dependencies_object_end_4))
																		mBase = m.M
																		v262 = m.ExcPending
																		if v262 != 0 {
																			return int32(0)
																		} else {
																			v287 = v231
																			m.G0 = v8 + int32(176)
																			return v287
																		}
																	}
																}
															}
														}
													}
												} else {
													if v109 == int32(3) {
														v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v266 = F_lappend(m, v265, v149)
														mBase = m.M
														v267 = m.ExcPending
														if v267 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v266
															v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															F_list_free(m, v269)
															mBase = m.M
															v271 = m.ExcPending
															if v271 != 0 {
																return int32(0)
															} else {
																*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = int64(0)
																v274 = int32(0)
																*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v274)
																*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v274
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)) = uint8(v274)
																*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v274)
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
																v287 = v274
																m.G0 = v8 + int32(176)
																return v287
															}
														}
													} else {
														v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+12))
														v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+12))
														*(*uint16)(unsafe.Add(mBase, uint32(v149)+16)) = uint16(v193)
														v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
														if v195 == v193&int32(_a_F_dependencies_object_end_7) {
															v231 = int32(23)
															v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v233 = F_errsave_start(m, v232)
															mBase = m.M
															v234 = m.ExcPending
															if v234 != 0 {
																return int32(0)
															} else {
																if v233 == int32(0) {
																	v287 = v231
																	m.G0 = v8 + int32(176)
																	return v287
																} else {
																	F_errcode(m, int32(33685634))
																	mBase = m.M
																	v239 = m.ExcPending
																	if v239 != 0 {
																		return int32(0)
																	} else {
																		v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																		*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v240
																		F_errmsg(m, int32(_a_F_dependencies_object_end_0), v8+int32(48))
																		mBase = m.M
																		v246 = m.ExcPending
																		if v246 != 0 {
																			return int32(0)
																		} else {
																			v247 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
																			*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(_a_F_dependencies_object_end_1)
																			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v247
																			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(_a_F_dependencies_object_end_5)
																			v256 = F_errdetail(m, int32(_a_F_dependencies_object_end_8), v8+int32(32))
																			mBase = m.M
																			v257 = m.ExcPending
																			if v257 != 0 {
																				return int32(0)
																			} else {
																				F_errsave_finish(m, v232, int32(_a_F_dependencies_object_end_3), int32(225), int32(_a_F_dependencies_object_end_4))
																				mBase = m.M
																				v262 = m.ExcPending
																				if v262 != 0 {
																					return int32(0)
																				} else {
																					v287 = v231
																					m.G0 = v8 + int32(176)
																					return v287
																				}
																			}
																		}
																	}
																}
															}
														} else {
															if v109 == int32(4) {
																v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																v266 = F_lappend(m, v265, v149)
																mBase = m.M
																v267 = m.ExcPending
																if v267 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v266
																	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	F_list_free(m, v269)
																	mBase = m.M
																	v271 = m.ExcPending
																	if v271 != 0 {
																		return int32(0)
																	} else {
																		*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = int64(0)
																		v274 = int32(0)
																		*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v274)
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v274
																		*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)) = uint8(v274)
																		*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v274)
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
																		v287 = v274
																		m.G0 = v8 + int32(176)
																		return v287
																	}
																}
															} else {
																v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
																v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+16))
																*(*uint16)(unsafe.Add(mBase, uint32(v149)+18)) = uint16(v203)
																v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
																if v205 == v203&int32(_a_F_dependencies_object_end_7) {
																	v231 = int32(23)
																	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	v233 = F_errsave_start(m, v232)
																	mBase = m.M
																	v234 = m.ExcPending
																	if v234 != 0 {
																		return int32(0)
																	} else {
																		if v233 == int32(0) {
																			v287 = v231
																			m.G0 = v8 + int32(176)
																			return v287
																		} else {
																			F_errcode(m, int32(33685634))
																			mBase = m.M
																			v239 = m.ExcPending
																			if v239 != 0 {
																				return int32(0)
																			} else {
																				v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																				*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v240
																				F_errmsg(m, int32(_a_F_dependencies_object_end_0), v8+int32(48))
																				mBase = m.M
																				v246 = m.ExcPending
																				if v246 != 0 {
																					return int32(0)
																				} else {
																					v247 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
																					*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(_a_F_dependencies_object_end_1)
																					*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v247
																					*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(_a_F_dependencies_object_end_5)
																					v256 = F_errdetail(m, int32(_a_F_dependencies_object_end_8), v8+int32(32))
																					mBase = m.M
																					v257 = m.ExcPending
																					if v257 != 0 {
																						return int32(0)
																					} else {
																						F_errsave_finish(m, v232, int32(_a_F_dependencies_object_end_3), int32(225), int32(_a_F_dependencies_object_end_4))
																						mBase = m.M
																						v262 = m.ExcPending
																						if v262 != 0 {
																							return int32(0)
																						} else {
																							v287 = v231
																							m.G0 = v8 + int32(176)
																							return v287
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	if v109 == int32(5) {
																		v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		v266 = F_lappend(m, v265, v149)
																		mBase = m.M
																		v267 = m.ExcPending
																		if v267 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v266
																			v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			F_list_free(m, v269)
																			mBase = m.M
																			v271 = m.ExcPending
																			if v271 != 0 {
																				return int32(0)
																			} else {
																				*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = int64(0)
																				v274 = int32(0)
																				*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v274)
																				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v274
																				*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)) = uint8(v274)
																				*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v274)
																				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
																				v287 = v274
																				m.G0 = v8 + int32(176)
																				return v287
																			}
																		}
																	} else {
																		v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
																		v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+20))
																		*(*uint16)(unsafe.Add(mBase, uint32(v149)+20)) = uint16(v213)
																		v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
																		if v215 == v213&int32(_a_F_dependencies_object_end_7) {
																			v231 = int32(23)
																			v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																			v233 = F_errsave_start(m, v232)
																			mBase = m.M
																			v234 = m.ExcPending
																			if v234 != 0 {
																				return int32(0)
																			} else {
																				if v233 == int32(0) {
																					v287 = v231
																					m.G0 = v8 + int32(176)
																					return v287
																				} else {
																					F_errcode(m, int32(33685634))
																					mBase = m.M
																					v239 = m.ExcPending
																					if v239 != 0 {
																						return int32(0)
																					} else {
																						v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																						*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v240
																						F_errmsg(m, int32(_a_F_dependencies_object_end_0), v8+int32(48))
																						mBase = m.M
																						v246 = m.ExcPending
																						if v246 != 0 {
																							return int32(0)
																						} else {
																							v247 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
																							*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(_a_F_dependencies_object_end_1)
																							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v247
																							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(_a_F_dependencies_object_end_5)
																							v256 = F_errdetail(m, int32(_a_F_dependencies_object_end_8), v8+int32(32))
																							mBase = m.M
																							v257 = m.ExcPending
																							if v257 != 0 {
																								return int32(0)
																							} else {
																								F_errsave_finish(m, v232, int32(_a_F_dependencies_object_end_3), int32(225), int32(_a_F_dependencies_object_end_4))
																								mBase = m.M
																								v262 = m.ExcPending
																								if v262 != 0 {
																									return int32(0)
																								} else {
																									v287 = v231
																									m.G0 = v8 + int32(176)
																									return v287
																								}
																							}
																						}
																					}
																				}
																			}
																		} else {
																			if v109 == int32(6) {
																				v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																				v266 = F_lappend(m, v265, v149)
																				mBase = m.M
																				v267 = m.ExcPending
																				if v267 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v266
																					v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																					F_list_free(m, v269)
																					mBase = m.M
																					v271 = m.ExcPending
																					if v271 != 0 {
																						return int32(0)
																					} else {
																						*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = int64(0)
																						v274 = int32(0)
																						*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v274)
																						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v274
																						*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)) = uint8(v274)
																						*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v274)
																						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
																						v287 = v274
																						m.G0 = v8 + int32(176)
																						return v287
																					}
																				}
																			} else {
																				v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																				v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
																				v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+24))
																				*(*uint16)(unsafe.Add(mBase, uint32(v149)+22)) = uint16(v223)
																				v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
																				if v225 != v223&int32(_a_F_dependencies_object_end_7) {
																					v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																					v266 = F_lappend(m, v265, v149)
																					mBase = m.M
																					v267 = m.ExcPending
																					if v267 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v266
																						v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																						F_list_free(m, v269)
																						mBase = m.M
																						v271 = m.ExcPending
																						if v271 != 0 {
																							return int32(0)
																						} else {
																							*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = int64(0)
																							v274 = int32(0)
																							*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v274)
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v274
																							*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)) = uint8(v274)
																							*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v274)
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
																							v287 = v274
																							m.G0 = v8 + int32(176)
																							return v287
																						}
																					}
																				} else {
																					v231 = int32(23)
																					v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																					v233 = F_errsave_start(m, v232)
																					mBase = m.M
																					v234 = m.ExcPending
																					if v234 != 0 {
																						return int32(0)
																					} else {
																						if v233 == int32(0) {
																							v287 = v231
																							m.G0 = v8 + int32(176)
																							return v287
																						} else {
																							F_errcode(m, int32(33685634))
																							mBase = m.M
																							v239 = m.ExcPending
																							if v239 != 0 {
																								return int32(0)
																							} else {
																								v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																								*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v240
																								F_errmsg(m, int32(_a_F_dependencies_object_end_0), v8+int32(48))
																								mBase = m.M
																								v246 = m.ExcPending
																								if v246 != 0 {
																									return int32(0)
																								} else {
																									v247 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
																									*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(_a_F_dependencies_object_end_1)
																									*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v247
																									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(_a_F_dependencies_object_end_5)
																									v256 = F_errdetail(m, int32(_a_F_dependencies_object_end_8), v8+int32(32))
																									mBase = m.M
																									v257 = m.ExcPending
																									if v257 != 0 {
																										return int32(0)
																									} else {
																										F_errsave_finish(m, v232, int32(_a_F_dependencies_object_end_3), int32(225), int32(_a_F_dependencies_object_end_4))
																										mBase = m.M
																										v262 = m.ExcPending
																										if v262 != 0 {
																											return int32(0)
																										} else {
																											v287 = v231
																											m.G0 = v8 + int32(176)
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
							v115 = int32(23)
							v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v117 = F_errsave_start(m, v116)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								if v117 == int32(0) {
									v287 = v115
									m.G0 = v8 + int32(176)
									return v287
								} else {
									F_errcode(m, int32(33685634))
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return int32(0)
									} else {
										v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v124
										F_errmsg(m, int32(_a_F_dependencies_object_end_0), v8+int32(16))
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return int32(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(30064771073)
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_dependencies_object_end_1)
											v136 = F_errdetail(m, int32(_a_F_dependencies_object_end_9), v8)
											mBase = m.M
											v137 = m.ExcPending
											if v137 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v116, int32(_a_F_dependencies_object_end_3), int32(197), int32(_a_F_dependencies_object_end_4))
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
													return int32(0)
												} else {
													v287 = v115
													m.G0 = v8 + int32(176)
													return v287
												}
											}
										}
									}
								}
							}
						}
					} else {
						v115 = int32(23)
						v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v117 = F_errsave_start(m, v116)
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return int32(0)
						} else {
							if v117 == int32(0) {
								v287 = v115
								m.G0 = v8 + int32(176)
								return v287
							} else {
								F_errcode(m, int32(33685634))
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return int32(0)
								} else {
									v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v124
									F_errmsg(m, int32(_a_F_dependencies_object_end_0), v8+int32(16))
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
										return int32(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(30064771073)
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_dependencies_object_end_1)
										v136 = F_errdetail(m, int32(_a_F_dependencies_object_end_9), v8)
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return int32(0)
										} else {
											F_errsave_finish(m, v116, int32(_a_F_dependencies_object_end_3), int32(197), int32(_a_F_dependencies_object_end_4))
											mBase = m.M
											v142 = m.ExcPending
											if v142 != 0 {
												return int32(0)
											} else {
												v287 = v115
												m.G0 = v8 + int32(176)
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
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v296 = m.ExcPending
		if v296 != 0 {
			return int32(0)
		} else {
			v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+164)) = v297
			*(*int32)(unsafe.Add(mBase, uint32(v8)+160)) = int32(_a_F_dependencies_object_end_10)
			F_errmsg_internal(m, int32(_a_F_dependencies_object_end_11), v8+int32(160))
			mBase = m.M
			v305 = m.ExcPending
			if v305 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_dependencies_object_end_3), int32(153), int32(_a_F_dependencies_object_end_4))
				mBase = m.M
				v310 = m.ExcPending
				if v310 != 0 {
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
func F_dependencies_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
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
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	v5 = m.G0
	v7 = v5 - int32(128)
	m.G0 = v7
	v9 = int32(_a_F_dependencies_object_field_start_0)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dependencies_object_field_start[0])))
	if base.B2i32(v12 == int32(0))|base.B2i32(v12 != v15) != 0 {
		v33 = v12
		v34 = v15
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v7 + int32(128)
	return v236
L2:
	;
	if v33-v34 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	goto L2
L4:
	;
	v18 = l1
	v19 = v9
	goto L5
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v23 == int32(0) {
		v33 = v23
		v34 = v22
		goto L3
	} else {
		goto L7
	}
L6:
	;
	v33 = v23
	v34 = v22
	goto L3
L7:
	;
	v26 = int32(1)
	if v23 == v22 {
		v18 = v18 + v26
		v19 = v19 + v26
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v38 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v74 = int32(_a_F_dependencies_object_field_start_1)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dependencies_object_field_start[1])))
	if base.B2i32(v77 == int32(0))|base.B2i32(v77 != v80) != 0 {
		v98 = v77
		v99 = v80
		goto L23
	} else {
		goto L24
	}
L12:
	;
	v41 = int32(23)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v43 = F_errsave_start(m, v42)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(3)
	v71 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v71)
	v236 = int32(0)
	goto L1
L15:
	;
	return int32(0)
L16:
	;
	if v43 == int32(0) {
		v236 = v41
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v52
	F_errmsg(m, int32(_a_F_dependencies_object_field_start_2), v7+int32(16))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_dependencies_object_field_start_0)
	v62 = F_errdetail(m, int32(_a_F_dependencies_object_field_start_3), v7)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	F_errsave_finish(m, v42, int32(_a_F_dependencies_object_field_start_4), int32(352), int32(_a_F_dependencies_object_field_start_5))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v236 = v41
	goto L1
L22:
	;
	if v98-v99 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	goto L22
L24:
	;
	v83 = l1
	v84 = v74
	goto L25
L25:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	if v88 == int32(0) {
		v98 = v88
		v99 = v87
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v98 = v88
	v99 = v87
	goto L23
L27:
	;
	v91 = int32(1)
	if v88 == v87 {
		v83 = v83 + v91
		v84 = v84 + v91
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	if v103 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v139 = int32(_a_F_dependencies_object_field_start_6)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dependencies_object_field_start[2])))
	if base.B2i32(v142 == int32(0))|base.B2i32(v142 != v145) != 0 {
		v163 = v142
		v164 = v145
		goto L42
	} else {
		goto L43
	}
L32:
	;
	v106 = int32(23)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v108 = F_errsave_start(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L15
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(5)
	v136 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)) = uint8(v136)
	v236 = int32(0)
	goto L1
L35:
	;
	if v108 == int32(0) {
		v236 = v106
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L15
	} else {
		goto L37
	}
L37:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v115
	F_errmsg(m, int32(_a_F_dependencies_object_field_start_2), v7+int32(48))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L15
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(_a_F_dependencies_object_field_start_1)
	v127 = F_errdetail(m, int32(_a_F_dependencies_object_field_start_3), v7+int32(32))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L15
	} else {
		goto L39
	}
L39:
	;
	F_errsave_finish(m, v107, int32(_a_F_dependencies_object_field_start_4), int32(369), int32(_a_F_dependencies_object_field_start_5))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L15
	} else {
		goto L40
	}
L40:
	;
	v236 = v106
	goto L1
L41:
	;
	if v163-v164 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	goto L41
L43:
	;
	v148 = l1
	v149 = v139
	goto L44
L44:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
	if v153 == int32(0) {
		v163 = v153
		v164 = v152
		goto L42
	} else {
		goto L46
	}
L45:
	;
	v163 = v153
	v164 = v152
	goto L42
L46:
	;
	v156 = int32(1)
	if v153 == v152 {
		v148 = v148 + v156
		v149 = v149 + v156
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)))
	if v168 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v204 = int32(23)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v206 = F_errsave_start(m, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L15
	} else {
		goto L60
	}
L51:
	;
	v171 = int32(23)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v173 = F_errsave_start(m, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L15
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(6)
	v201 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)) = uint8(v201)
	v236 = int32(0)
	goto L1
L54:
	;
	if v173 == int32(0) {
		v236 = v171
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L15
	} else {
		goto L56
	}
L56:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+80)) = v180
	F_errmsg(m, int32(_a_F_dependencies_object_field_start_2), v7+int32(80))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L15
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = int32(_a_F_dependencies_object_field_start_6)
	v192 = F_errdetail(m, int32(_a_F_dependencies_object_field_start_3), v7-int32(-64))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L15
	} else {
		goto L58
	}
L58:
	;
	F_errsave_finish(m, v172, int32(_a_F_dependencies_object_field_start_4), int32(386), int32(_a_F_dependencies_object_field_start_5))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L15
	} else {
		goto L59
	}
L59:
	;
	v236 = v171
	goto L1
L60:
	;
	if v206 == int32(0) {
		v236 = v204
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L15
	} else {
		goto L62
	}
L62:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+112)) = v213
	F_errmsg(m, int32(_a_F_dependencies_object_field_start_2), v7+int32(112))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L15
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+104)) = int32(_a_F_dependencies_object_field_start_6)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+100)) = int32(_a_F_dependencies_object_field_start_1)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = int32(_a_F_dependencies_object_field_start_0)
	v229 = F_errdetail(m, int32(_a_F_dependencies_object_field_start_7), v7+int32(96))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L15
	} else {
		goto L64
	}
L64:
	;
	F_errsave_finish(m, v205, int32(_a_F_dependencies_object_field_start_4), int32(401), int32(_a_F_dependencies_object_field_start_5))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L15
	} else {
		goto L65
	}
L65:
	;
	v236 = v204
	goto L1
}
