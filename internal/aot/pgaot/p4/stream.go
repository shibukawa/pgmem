package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_read_stream_next_buffer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v93 int64
	_ = v93
	var v96 int32
	_ = v96
	var v99 int64
	_ = v99
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int64
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v258 int64
	_ = v258
	var v262 int64
	_ = v262
	var v263 int64
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v322 int32
	_ = v322
	v3 = int32(0)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)))
	if v11 == int32(1) {
		v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+86)))
		v19 = l0 + v14<<(uint(int32(2))%32) + int32(92)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v21 != int32(-1) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(-1)
			v35 = v21
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
			v48 = F_StartReadBuffer(m, v36+int32(4), v19, v35, v39|v40<<(uint(int32(1))%32)&int32(2)|int32(8))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				if v48 == int32(0) {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v52 == int32(0) {
						v322 = v20
						return v322
					} else {
						v55 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
						*(*int64)(unsafe.Add(mBase, uint32(v52))) = v55 + int64(1)
						v59 = *(*int64)(unsafe.Add(mBase, uint32(v52)+8))
						v60 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
						*(*int64)(unsafe.Add(mBase, uint32(v52)+8)) = v59 + v60
						v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
						v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+16)))
						if v63 <= v64 {
							v322 = v20
							return v322
						} else {
							*(*uint16)(unsafe.Add(mBase, uint32(v52)+16)) = uint16(v63)
							return v20
						}
					}
				} else {
					v68 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+80)) = uint16(v68)
					v70 = int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v70)
					v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+82)) = uint16(base.B2i32(v70 < v72))
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					*(*uint16)(unsafe.Add(mBase, uint32(v76))) = uint16(v14)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v35 + v70
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v81 == v68 {
					} else {
						v84 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
						v85 = *(*int64)(unsafe.Add(mBase, uint32(v81)+32))
						v86 = int64(1)
						*(*int64)(unsafe.Add(mBase, uint32(v81)+32)) = v85 + v86
						v89 = *(*int64)(unsafe.Add(mBase, uint32(v81)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v81)+40)) = v89 + v86
						v93 = *(*int64)(unsafe.Add(mBase, uint32(v81)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v81)+48)) = v84 + v93
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v96 == int32(0) {
						} else {
							v99 = *(*int64)(unsafe.Add(mBase, uint32(v96)))
							*(*int64)(unsafe.Add(mBase, uint32(v96))) = v99 + int64(1)
							v103 = *(*int64)(unsafe.Add(mBase, uint32(v96)+8))
							v104 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
							*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v103 + v104
							v107 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
							v108 = int32(*(*int16)(unsafe.Add(mBase, uint32(v96)+16)))
							if v107 <= v108 {
							} else {
								*(*uint16)(unsafe.Add(mBase, uint32(v96)+16)) = uint16(v107)
							}
						}
					}
					v122 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v122)
					return v20
				}
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v29 = m.T0[v28].(func(*base.Module, int32, int32, int32) int32)(m, l0, v26, int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				if v29 == int32(-1) {
					v111 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v111)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v111
					v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+86)) = uint16(v115)
					*(*int32)(unsafe.Add(mBase, uint32(v19))) = v111
					v122 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v122)
					return v20
				} else {
					v35 = v29
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
					v48 = F_StartReadBuffer(m, v36+int32(4), v19, v35, v39|v40<<(uint(int32(1))%32)&int32(2)|int32(8))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						if v48 == int32(0) {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if v52 == int32(0) {
								v322 = v20
								return v322
							} else {
								v55 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
								*(*int64)(unsafe.Add(mBase, uint32(v52))) = v55 + int64(1)
								v59 = *(*int64)(unsafe.Add(mBase, uint32(v52)+8))
								v60 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
								*(*int64)(unsafe.Add(mBase, uint32(v52)+8)) = v59 + v60
								v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
								v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+16)))
								if v63 <= v64 {
									v322 = v20
									return v322
								} else {
									*(*uint16)(unsafe.Add(mBase, uint32(v52)+16)) = uint16(v63)
									return v20
								}
							}
						} else {
							v68 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+80)) = uint16(v68)
							v70 = int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v70)
							v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+82)) = uint16(base.B2i32(v70 < v72))
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
							*(*uint16)(unsafe.Add(mBase, uint32(v76))) = uint16(v14)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v35 + v70
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if v81 == v68 {
							} else {
								v84 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
								v85 = *(*int64)(unsafe.Add(mBase, uint32(v81)+32))
								v86 = int64(1)
								*(*int64)(unsafe.Add(mBase, uint32(v81)+32)) = v85 + v86
								v89 = *(*int64)(unsafe.Add(mBase, uint32(v81)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v81)+40)) = v89 + v86
								v93 = *(*int64)(unsafe.Add(mBase, uint32(v81)+48))
								*(*int64)(unsafe.Add(mBase, uint32(v81)+48)) = v84 + v93
								v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v96 == int32(0) {
								} else {
									v99 = *(*int64)(unsafe.Add(mBase, uint32(v96)))
									*(*int64)(unsafe.Add(mBase, uint32(v96))) = v99 + int64(1)
									v103 = *(*int64)(unsafe.Add(mBase, uint32(v96)+8))
									v104 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
									*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v103 + v104
									v107 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
									v108 = int32(*(*int16)(unsafe.Add(mBase, uint32(v96)+16)))
									if v107 <= v108 {
									} else {
										*(*uint16)(unsafe.Add(mBase, uint32(v96)+16)) = uint16(v107)
									}
								}
							}
							v122 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v122)
							return v20
						}
					}
				}
			}
		}
	} else {
		v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
		if v125 == int32(0) {
			v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
			if v128 == int32(0) {
				v322 = v3
				return v322
			} else {
				F_read_stream_look_ahead(m, l0)
				mBase = m.M
				v132 = m.ExcPending
				if v132 != 0 {
					return int32(0)
				} else {
					v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
					if v133 == int32(0) {
						v322 = v3
						return v322
					} else {
						v137 = l0 + int32(92)
						v138 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+86)))
						v141 = v137 + v138<<(uint(int32(2))%32)
						v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
						if l1 != 0 {
							v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v143 + v144*v138
						} else {
						}
						v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
						if v148 <= int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v141))) = int32(0)
							v242 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
							if v138 < v242-int32(1) {
								v246 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
								v247 = int32(2)
								*(*int32)(unsafe.Add(mBase, uint32(v137+v246<<(uint(v247)%32)+v138<<(uint(v247)%32)))) = int32(0)
							} else {
							}
							v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if v255 == int32(0) {
							} else {
								v258 = *(*int64)(unsafe.Add(mBase, uint32(v255)))
								*(*int64)(unsafe.Add(mBase, uint32(v255))) = v258 + int64(1)
								v262 = *(*int64)(unsafe.Add(mBase, uint32(v255)+8))
								v263 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
								*(*int64)(unsafe.Add(mBase, uint32(v255)+8)) = v262 + v263
								v266 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
								v267 = int32(*(*int16)(unsafe.Add(mBase, uint32(v255)+16)))
								if v266 <= v267 {
								} else {
									*(*uint16)(unsafe.Add(mBase, uint32(v255)+16)) = uint16(v266)
								}
							}
							v271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
							v272 = int32(1)
							v273 = v271 - v272
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v273)
							v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+86)))
							v277 = v275 + v272
							v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
							if v279 != v277&int32(_a_F_read_stream_next_buffer_0) {
								v283 = v277
							} else {
								v283 = int32(0)
							}
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+86)) = uint16(v283)
							F_read_stream_look_ahead(m, l0)
							mBase = m.M
							v286 = m.ExcPending
							if v286 != 0 {
								return int32(0)
							} else {
								v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
								if v287 != 0 {
									v322 = v142
								} else {
									v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
									if v288 != 0 {
										v322 = v142
									} else {
										v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
										if v289 != int32(1) {
											v322 = v142
										} else {
											v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
											if v292 != int32(1) {
												v322 = v142
											} else {
												v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
												if v295 != int32(1) {
													v322 = v142
												} else {
													v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)))
													if v298 != 0 {
														v322 = v142
													} else {
														v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
														if v299 != 0 {
															v322 = v142
														} else {
															v300 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+86)))
															v301 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
															if v300 < v301-int32(1) {
																v305 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
																v306 = int32(2)
																*(*int32)(unsafe.Add(mBase, uint32(v137+v305<<(uint(v306)%32)+v300<<(uint(v306)%32)))) = int32(0)
															} else {
															}
															v314 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v314)
															v322 = v142
														}
													}
												}
											}
										}
									}
								}
								return v322
							}
						} else {
							v151 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+80)))
							v153 = v151 * int32(84)
							v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
							v155 = v153 + v154
							v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155))))
							if v156 != v138&int32(_a_F_read_stream_next_buffer_0) {
								*(*int32)(unsafe.Add(mBase, uint32(v141))) = int32(0)
								v242 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
								if v138 < v242-int32(1) {
									v246 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
									v247 = int32(2)
									*(*int32)(unsafe.Add(mBase, uint32(v137+v246<<(uint(v247)%32)+v138<<(uint(v247)%32)))) = int32(0)
								} else {
								}
								v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v255 == int32(0) {
								} else {
									v258 = *(*int64)(unsafe.Add(mBase, uint32(v255)))
									*(*int64)(unsafe.Add(mBase, uint32(v255))) = v258 + int64(1)
									v262 = *(*int64)(unsafe.Add(mBase, uint32(v255)+8))
									v263 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
									*(*int64)(unsafe.Add(mBase, uint32(v255)+8)) = v262 + v263
									v266 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
									v267 = int32(*(*int16)(unsafe.Add(mBase, uint32(v255)+16)))
									if v266 <= v267 {
									} else {
										*(*uint16)(unsafe.Add(mBase, uint32(v255)+16)) = uint16(v266)
									}
								}
								v271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
								v272 = int32(1)
								v273 = v271 - v272
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v273)
								v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+86)))
								v277 = v275 + v272
								v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
								if v279 != v277&int32(_a_F_read_stream_next_buffer_0) {
									v283 = v277
								} else {
									v283 = int32(0)
								}
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+86)) = uint16(v283)
								F_read_stream_look_ahead(m, l0)
								mBase = m.M
								v286 = m.ExcPending
								if v286 != 0 {
									return int32(0)
								} else {
									v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
									if v287 != 0 {
										v322 = v142
									} else {
										v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
										if v288 != 0 {
											v322 = v142
										} else {
											v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
											if v289 != int32(1) {
												v322 = v142
											} else {
												v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
												if v292 != int32(1) {
													v322 = v142
												} else {
													v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
													if v295 != int32(1) {
														v322 = v142
													} else {
														v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)))
														if v298 != 0 {
															v322 = v142
														} else {
															v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
															if v299 != 0 {
																v322 = v142
															} else {
																v300 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+86)))
																v301 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
																if v300 < v301-int32(1) {
																	v305 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
																	v306 = int32(2)
																	*(*int32)(unsafe.Add(mBase, uint32(v137+v305<<(uint(v306)%32)+v300<<(uint(v306)%32)))) = int32(0)
																} else {
																}
																v314 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v314)
																v322 = v142
															}
														}
													}
												}
											}
										}
									}
									return v322
								}
							} else {
								v162 = F_WaitReadBuffers(m, v155+int32(4))
								mBase = m.M
								v163 = m.ExcPending
								if v163 != 0 {
									return int32(0)
								} else {
									v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
									v165 = int32(1)
									v166 = v164 - v165
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v166)
									v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+80)))
									v170 = v168 + v165
									v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
									if v172 != v170&int32(_a_F_read_stream_next_buffer_0) {
										v176 = v170
									} else {
										v176 = int32(0)
									}
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+80)) = uint16(v176)
									v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
									v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178+v153)+32)))
									if v180&int32(8)|v162 == int32(0) {
									} else {
										v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v186 != 0 {
											v187 = *(*int64)(unsafe.Add(mBase, uint32(v186)+24))
											*(*int64)(unsafe.Add(mBase, uint32(v186)+24)) = v187 + int64(1)
										} else {
										}
										v191 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+16)))
										if v191 <= int32(0) {
										} else {
											v195 = v191 << (uint(int32(1)) % 32)
											v196 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
											if v195&int32(_a_F_read_stream_next_buffer_1) < v196 {
												v200 = v195
											} else {
												v200 = v196
											}
											*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v200)
										}
									}
									v204 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v204)
									v206 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+14)))
									if v206 <= int32(0) {
									} else {
										v209 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
										if v209 <= v206 {
										} else {
											v214 = v206 << (uint(int32(1)) % 32) & int32(_a_F_read_stream_next_buffer_1)
											v216 = v209 & int32(_a_F_read_stream_next_buffer_0)
											if base.Ui32(v214) < base.Ui32(v216) {
												v218 = v214
											} else {
												v218 = v216
											}
											if v218 < v204 {
												v220 = v218
											} else {
												v220 = v204
											}
											*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v220)
										}
									}
									v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
									if v224 != int32(1) {
									} else {
										v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
										v231 = *(*int32)(unsafe.Add(mBase, uint32(v227+v151*int32(84))+28))
										v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v231 != v232 {
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(-1)
										}
									}
									*(*int32)(unsafe.Add(mBase, uint32(v141))) = int32(0)
									v242 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
									if v138 < v242-int32(1) {
										v246 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
										v247 = int32(2)
										*(*int32)(unsafe.Add(mBase, uint32(v137+v246<<(uint(v247)%32)+v138<<(uint(v247)%32)))) = int32(0)
									} else {
									}
									v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v255 == int32(0) {
									} else {
										v258 = *(*int64)(unsafe.Add(mBase, uint32(v255)))
										*(*int64)(unsafe.Add(mBase, uint32(v255))) = v258 + int64(1)
										v262 = *(*int64)(unsafe.Add(mBase, uint32(v255)+8))
										v263 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
										*(*int64)(unsafe.Add(mBase, uint32(v255)+8)) = v262 + v263
										v266 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
										v267 = int32(*(*int16)(unsafe.Add(mBase, uint32(v255)+16)))
										if v266 <= v267 {
										} else {
											*(*uint16)(unsafe.Add(mBase, uint32(v255)+16)) = uint16(v266)
										}
									}
									v271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
									v272 = int32(1)
									v273 = v271 - v272
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v273)
									v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+86)))
									v277 = v275 + v272
									v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
									if v279 != v277&int32(_a_F_read_stream_next_buffer_0) {
										v283 = v277
									} else {
										v283 = int32(0)
									}
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+86)) = uint16(v283)
									F_read_stream_look_ahead(m, l0)
									mBase = m.M
									v286 = m.ExcPending
									if v286 != 0 {
										return int32(0)
									} else {
										v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
										if v287 != 0 {
											v322 = v142
										} else {
											v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
											if v288 != 0 {
												v322 = v142
											} else {
												v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
												if v289 != int32(1) {
													v322 = v142
												} else {
													v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
													if v292 != int32(1) {
														v322 = v142
													} else {
														v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
														if v295 != int32(1) {
															v322 = v142
														} else {
															v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)))
															if v298 != 0 {
																v322 = v142
															} else {
																v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
																if v299 != 0 {
																	v322 = v142
																} else {
																	v300 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+86)))
																	v301 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
																	if v300 < v301-int32(1) {
																		v305 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
																		v306 = int32(2)
																		*(*int32)(unsafe.Add(mBase, uint32(v137+v305<<(uint(v306)%32)+v300<<(uint(v306)%32)))) = int32(0)
																	} else {
																	}
																	v314 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v314)
																	v322 = v142
																}
															}
														}
													}
												}
											}
										}
										return v322
									}
								}
							}
						}
					}
				}
			}
		} else {
			v137 = l0 + int32(92)
			v138 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+86)))
			v141 = v137 + v138<<(uint(int32(2))%32)
			v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
			if l1 != 0 {
				v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
				v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v143 + v144*v138
			} else {
			}
			v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
			if v148 <= int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v141))) = int32(0)
				v242 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
				if v138 < v242-int32(1) {
					v246 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
					v247 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v137+v246<<(uint(v247)%32)+v138<<(uint(v247)%32)))) = int32(0)
				} else {
				}
				v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v255 == int32(0) {
				} else {
					v258 = *(*int64)(unsafe.Add(mBase, uint32(v255)))
					*(*int64)(unsafe.Add(mBase, uint32(v255))) = v258 + int64(1)
					v262 = *(*int64)(unsafe.Add(mBase, uint32(v255)+8))
					v263 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
					*(*int64)(unsafe.Add(mBase, uint32(v255)+8)) = v262 + v263
					v266 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
					v267 = int32(*(*int16)(unsafe.Add(mBase, uint32(v255)+16)))
					if v266 <= v267 {
					} else {
						*(*uint16)(unsafe.Add(mBase, uint32(v255)+16)) = uint16(v266)
					}
				}
				v271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				v272 = int32(1)
				v273 = v271 - v272
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v273)
				v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+86)))
				v277 = v275 + v272
				v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
				if v279 != v277&int32(_a_F_read_stream_next_buffer_0) {
					v283 = v277
				} else {
					v283 = int32(0)
				}
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+86)) = uint16(v283)
				F_read_stream_look_ahead(m, l0)
				mBase = m.M
				v286 = m.ExcPending
				if v286 != 0 {
					return int32(0)
				} else {
					v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
					if v287 != 0 {
						v322 = v142
					} else {
						v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
						if v288 != 0 {
							v322 = v142
						} else {
							v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
							if v289 != int32(1) {
								v322 = v142
							} else {
								v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
								if v292 != int32(1) {
									v322 = v142
								} else {
									v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
									if v295 != int32(1) {
										v322 = v142
									} else {
										v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)))
										if v298 != 0 {
											v322 = v142
										} else {
											v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											if v299 != 0 {
												v322 = v142
											} else {
												v300 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+86)))
												v301 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
												if v300 < v301-int32(1) {
													v305 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
													v306 = int32(2)
													*(*int32)(unsafe.Add(mBase, uint32(v137+v305<<(uint(v306)%32)+v300<<(uint(v306)%32)))) = int32(0)
												} else {
												}
												v314 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v314)
												v322 = v142
											}
										}
									}
								}
							}
						}
					}
					return v322
				}
			} else {
				v151 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+80)))
				v153 = v151 * int32(84)
				v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v155 = v153 + v154
				v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155))))
				if v156 != v138&int32(_a_F_read_stream_next_buffer_0) {
					*(*int32)(unsafe.Add(mBase, uint32(v141))) = int32(0)
					v242 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
					if v138 < v242-int32(1) {
						v246 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
						v247 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v137+v246<<(uint(v247)%32)+v138<<(uint(v247)%32)))) = int32(0)
					} else {
					}
					v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v255 == int32(0) {
					} else {
						v258 = *(*int64)(unsafe.Add(mBase, uint32(v255)))
						*(*int64)(unsafe.Add(mBase, uint32(v255))) = v258 + int64(1)
						v262 = *(*int64)(unsafe.Add(mBase, uint32(v255)+8))
						v263 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
						*(*int64)(unsafe.Add(mBase, uint32(v255)+8)) = v262 + v263
						v266 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
						v267 = int32(*(*int16)(unsafe.Add(mBase, uint32(v255)+16)))
						if v266 <= v267 {
						} else {
							*(*uint16)(unsafe.Add(mBase, uint32(v255)+16)) = uint16(v266)
						}
					}
					v271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
					v272 = int32(1)
					v273 = v271 - v272
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v273)
					v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+86)))
					v277 = v275 + v272
					v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
					if v279 != v277&int32(_a_F_read_stream_next_buffer_0) {
						v283 = v277
					} else {
						v283 = int32(0)
					}
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+86)) = uint16(v283)
					F_read_stream_look_ahead(m, l0)
					mBase = m.M
					v286 = m.ExcPending
					if v286 != 0 {
						return int32(0)
					} else {
						v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
						if v287 != 0 {
							v322 = v142
						} else {
							v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
							if v288 != 0 {
								v322 = v142
							} else {
								v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
								if v289 != int32(1) {
									v322 = v142
								} else {
									v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
									if v292 != int32(1) {
										v322 = v142
									} else {
										v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
										if v295 != int32(1) {
											v322 = v142
										} else {
											v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)))
											if v298 != 0 {
												v322 = v142
											} else {
												v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												if v299 != 0 {
													v322 = v142
												} else {
													v300 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+86)))
													v301 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
													if v300 < v301-int32(1) {
														v305 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
														v306 = int32(2)
														*(*int32)(unsafe.Add(mBase, uint32(v137+v305<<(uint(v306)%32)+v300<<(uint(v306)%32)))) = int32(0)
													} else {
													}
													v314 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v314)
													v322 = v142
												}
											}
										}
									}
								}
							}
						}
						return v322
					}
				} else {
					v162 = F_WaitReadBuffers(m, v155+int32(4))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return int32(0)
					} else {
						v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
						v165 = int32(1)
						v166 = v164 - v165
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v166)
						v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+80)))
						v170 = v168 + v165
						v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
						if v172 != v170&int32(_a_F_read_stream_next_buffer_0) {
							v176 = v170
						} else {
							v176 = int32(0)
						}
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+80)) = uint16(v176)
						v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
						v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178+v153)+32)))
						if v180&int32(8)|v162 == int32(0) {
						} else {
							v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if v186 != 0 {
								v187 = *(*int64)(unsafe.Add(mBase, uint32(v186)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v186)+24)) = v187 + int64(1)
							} else {
							}
							v191 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+16)))
							if v191 <= int32(0) {
							} else {
								v195 = v191 << (uint(int32(1)) % 32)
								v196 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
								if v195&int32(_a_F_read_stream_next_buffer_1) < v196 {
									v200 = v195
								} else {
									v200 = v196
								}
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v200)
							}
						}
						v204 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v204)
						v206 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+14)))
						if v206 <= int32(0) {
						} else {
							v209 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
							if v209 <= v206 {
							} else {
								v214 = v206 << (uint(int32(1)) % 32) & int32(_a_F_read_stream_next_buffer_1)
								v216 = v209 & int32(_a_F_read_stream_next_buffer_0)
								if base.Ui32(v214) < base.Ui32(v216) {
									v218 = v214
								} else {
									v218 = v216
								}
								if v218 < v204 {
									v220 = v218
								} else {
									v220 = v204
								}
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v220)
							}
						}
						v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
						if v224 != int32(1) {
						} else {
							v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
							v231 = *(*int32)(unsafe.Add(mBase, uint32(v227+v151*int32(84))+28))
							v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							if v231 != v232 {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(-1)
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v141))) = int32(0)
						v242 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
						if v138 < v242-int32(1) {
							v246 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
							v247 = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v137+v246<<(uint(v247)%32)+v138<<(uint(v247)%32)))) = int32(0)
						} else {
						}
						v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v255 == int32(0) {
						} else {
							v258 = *(*int64)(unsafe.Add(mBase, uint32(v255)))
							*(*int64)(unsafe.Add(mBase, uint32(v255))) = v258 + int64(1)
							v262 = *(*int64)(unsafe.Add(mBase, uint32(v255)+8))
							v263 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
							*(*int64)(unsafe.Add(mBase, uint32(v255)+8)) = v262 + v263
							v266 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
							v267 = int32(*(*int16)(unsafe.Add(mBase, uint32(v255)+16)))
							if v266 <= v267 {
							} else {
								*(*uint16)(unsafe.Add(mBase, uint32(v255)+16)) = uint16(v266)
							}
						}
						v271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
						v272 = int32(1)
						v273 = v271 - v272
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v273)
						v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+86)))
						v277 = v275 + v272
						v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
						if v279 != v277&int32(_a_F_read_stream_next_buffer_0) {
							v283 = v277
						} else {
							v283 = int32(0)
						}
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+86)) = uint16(v283)
						F_read_stream_look_ahead(m, l0)
						mBase = m.M
						v286 = m.ExcPending
						if v286 != 0 {
							return int32(0)
						} else {
							v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
							if v287 != 0 {
								v322 = v142
							} else {
								v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
								if v288 != 0 {
									v322 = v142
								} else {
									v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
									if v289 != int32(1) {
										v322 = v142
									} else {
										v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
										if v292 != int32(1) {
											v322 = v142
										} else {
											v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
											if v295 != int32(1) {
												v322 = v142
											} else {
												v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)))
												if v298 != 0 {
													v322 = v142
												} else {
													v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
													if v299 != 0 {
														v322 = v142
													} else {
														v300 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+86)))
														v301 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
														if v300 < v301-int32(1) {
															v305 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
															v306 = int32(2)
															*(*int32)(unsafe.Add(mBase, uint32(v137+v305<<(uint(v306)%32)+v300<<(uint(v306)%32)))) = int32(0)
														} else {
														}
														v314 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v314)
														v322 = v142
													}
												}
											}
										}
									}
								}
							}
							return v322
						}
					}
				}
			}
		}
	}
}
