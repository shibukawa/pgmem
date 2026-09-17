package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_fsm_readbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int64
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v14 <= int32(0) {
		v64 = v13
	} else {
		v18 = v14 & int32(7)
		if base.Ui32(int32(8)) <= base.Ui32(v14) {
			v25 = int32(0)
			v27 = v13
			for {
				v33 = v27 * int32(-1799145247)
				v35 = v25 + int32(8)
				if v35 != v14&int32(2147483640) {
					v25 = v35
					v27 = v33
					continue
				} else {
					break
				}
				break
			}
			if v18 == int32(0) {
				v64 = v33
			} else {
				v42 = v33
				v49 = int32(0)
				v51 = v42
				for {
					v57 = v51 * int32(4069)
					v59 = v49 + int32(1)
					if v59 != v18 {
						v49 = v59
						v51 = v57
						continue
					} else {
						break
					}
					break
				}
				v64 = v57
			}
		} else {
			v42 = v13
			v49 = int32(0)
			v51 = v42
			for {
				v57 = v51 * int32(4069)
				v59 = v49 + int32(1)
				if v59 != v18 {
					v49 = v59
					v51 = v57
					continue
				} else {
					break
				}
				break
			}
			v64 = v57
		}
	}
	v70 = base.I32_div_u_s(v64, int32(4069))
	v73 = base.I32_div_u_s(v64, int32(16556761))
	v76 = v64 + v70 + v73 + int32(3)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v79 == int32(0) {
		v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v83
		v85 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v85
		v89 = F_smgropen(m, v11+int32(24), v82)
		mBase = m.M
		v92 = m.ExcPending
		if v92 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v89
			v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)+72))
			if v95 != 0 {
				v103 = v95
			} else {
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v89)+76))
				v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)+80))
				*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v97
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v89)+76))
				*(*int32)(unsafe.Add(mBase, uint32(v97))) = v99
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v89)+72))
				v103 = v101
			}
			*(*int32)(unsafe.Add(mBase, uint32(v89)+72)) = v103 + int32(1)
			v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v108 = v107
			v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+24))
			v112 = v76 + (v14 ^ int32(-1))
			if base.B2i32(v109 != int32(-1))&base.B2i32(base.Ui32(v112) < base.Ui32(v109)) != 0 {
				v150 = F_ReadBufferExtended(m, l0, int32(1), v112, int32(3), int32(0))
				mBase = m.M
				v151 = m.ExcPending
				if v151 != 0 {
					return int32(0)
				} else {
					v153 = v150
					if v153 < int32(0) {
						v159 = (v153 ^ int32(-1)) << (uint(int32(2)) % 32)
						v161 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
						v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v161)))
						v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
						if v164 != 0 {
							v253 = v153
							m.G0 = v11 + int32(48)
							return v253
						} else {
							F_LockBuffer(m, v153, int32(2))
							mBase = m.M
							v167 = m.ExcPending
							if v167 != 0 {
								return int32(0)
							} else {
								v169 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
								v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v159)))
								v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+14)))
								if v172 == int32(0) {
									v195 = v171
									v196 = int32(_a_F_fsm_readbuf_0)
									v197 = int32(0)
									if v197|(v195&int32(3)|int32(1)) == v197 {
										v214 = v195 + v196
										v216 = v195 + int32(4)
										if base.Ui32(v216) < base.Ui32(v214) {
											v218 = v214
										} else {
											v218 = v216
										}
										v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
										if v223 == int32(0) {
										} else {
											base.MemoryFill(m, v195, int32(0), v223)
										}
									} else {
										base.MemoryFill(m, v195, int32(0), v196)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
									v237 = int32(_a_F_fsm_readbuf_2)
									*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
									v243 = int32(_a_F_fsm_readbuf_0)
									*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
									*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
								} else {
								}
								F_LockBuffer(m, v153, int32(0))
								mBase = m.M
								v250 = m.ExcPending
								if v250 != 0 {
									return int32(0)
								} else {
									v253 = v153
									m.G0 = v11 + int32(48)
									return v253
								}
							}
						}
					} else {
						v176 = v153 << (uint(int32(13)) % 32)
						v178 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
						v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+v178-int32(_a_F_fsm_readbuf_3)))))
						if v182 != 0 {
							v253 = v153
							m.G0 = v11 + int32(48)
							return v253
						} else {
							F_LockBuffer(m, v153, int32(2))
							mBase = m.M
							v185 = m.ExcPending
							if v185 != 0 {
								return int32(0)
							} else {
								v187 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
								v188 = v187 + v176
								v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188-int32(_a_F_fsm_readbuf_3)))))
								if v191 != 0 {
								} else {
									v195 = v188 + int32(-8192)
									v196 = int32(_a_F_fsm_readbuf_0)
									v197 = int32(0)
									if v197|(v195&int32(3)|int32(1)) == v197 {
										v214 = v195 + v196
										v216 = v195 + int32(4)
										if base.Ui32(v216) < base.Ui32(v214) {
											v218 = v214
										} else {
											v218 = v216
										}
										v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
										if v223 == int32(0) {
										} else {
											base.MemoryFill(m, v195, int32(0), v223)
										}
									} else {
										base.MemoryFill(m, v195, int32(0), v196)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
									v237 = int32(_a_F_fsm_readbuf_2)
									*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
									v243 = int32(_a_F_fsm_readbuf_0)
									*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
									*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
								}
								F_LockBuffer(m, v153, int32(0))
								mBase = m.M
								v250 = m.ExcPending
								if v250 != 0 {
									return int32(0)
								} else {
									v253 = v153
									m.G0 = v11 + int32(48)
									return v253
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v108)+24)) = int32(-1)
				v118 = F_smgrexists(m, v108, int32(1))
				mBase = m.M
				v119 = m.ExcPending
				if v119 != 0 {
					return int32(0)
				} else {
					if v118 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v108)+24)) = int32(0)
						v129 = int32(0)
						if l2 == v129 {
							v253 = v129
							m.G0 = v11 + int32(48)
							return v253
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = l0
							v135 = *(*int64)(unsafe.Add(mBase, uint32(v11)+36))
							*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v135
							v137 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v137
							v145 = F_ExtendBufferedRelTo(m, v11+int32(8), int32(1), int32(20), v76-v14, int32(3))
							mBase = m.M
							v146 = m.ExcPending
							if v146 != 0 {
								return int32(0)
							} else {
								v153 = v145
								if v153 < int32(0) {
									v159 = (v153 ^ int32(-1)) << (uint(int32(2)) % 32)
									v161 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
									v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v161)))
									v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
									if v164 != 0 {
										v253 = v153
										m.G0 = v11 + int32(48)
										return v253
									} else {
										F_LockBuffer(m, v153, int32(2))
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return int32(0)
										} else {
											v169 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
											v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v159)))
											v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+14)))
											if v172 == int32(0) {
												v195 = v171
												v196 = int32(_a_F_fsm_readbuf_0)
												v197 = int32(0)
												if v197|(v195&int32(3)|int32(1)) == v197 {
													v214 = v195 + v196
													v216 = v195 + int32(4)
													if base.Ui32(v216) < base.Ui32(v214) {
														v218 = v214
													} else {
														v218 = v216
													}
													v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
													if v223 == int32(0) {
													} else {
														base.MemoryFill(m, v195, int32(0), v223)
													}
												} else {
													base.MemoryFill(m, v195, int32(0), v196)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
												v237 = int32(_a_F_fsm_readbuf_2)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
												v243 = int32(_a_F_fsm_readbuf_0)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
											} else {
											}
											F_LockBuffer(m, v153, int32(0))
											mBase = m.M
											v250 = m.ExcPending
											if v250 != 0 {
												return int32(0)
											} else {
												v253 = v153
												m.G0 = v11 + int32(48)
												return v253
											}
										}
									}
								} else {
									v176 = v153 << (uint(int32(13)) % 32)
									v178 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
									v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+v178-int32(_a_F_fsm_readbuf_3)))))
									if v182 != 0 {
										v253 = v153
										m.G0 = v11 + int32(48)
										return v253
									} else {
										F_LockBuffer(m, v153, int32(2))
										mBase = m.M
										v185 = m.ExcPending
										if v185 != 0 {
											return int32(0)
										} else {
											v187 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
											v188 = v187 + v176
											v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188-int32(_a_F_fsm_readbuf_3)))))
											if v191 != 0 {
											} else {
												v195 = v188 + int32(-8192)
												v196 = int32(_a_F_fsm_readbuf_0)
												v197 = int32(0)
												if v197|(v195&int32(3)|int32(1)) == v197 {
													v214 = v195 + v196
													v216 = v195 + int32(4)
													if base.Ui32(v216) < base.Ui32(v214) {
														v218 = v214
													} else {
														v218 = v216
													}
													v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
													if v223 == int32(0) {
													} else {
														base.MemoryFill(m, v195, int32(0), v223)
													}
												} else {
													base.MemoryFill(m, v195, int32(0), v196)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
												v237 = int32(_a_F_fsm_readbuf_2)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
												v243 = int32(_a_F_fsm_readbuf_0)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
											}
											F_LockBuffer(m, v153, int32(0))
											mBase = m.M
											v250 = m.ExcPending
											if v250 != 0 {
												return int32(0)
											} else {
												v253 = v153
												m.G0 = v11 + int32(48)
												return v253
											}
										}
									}
								}
							}
						}
					} else {
						v125 = F_smgrnblocks(m, v108, int32(1))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v108)+24))
							if base.Ui32(v112) < base.Ui32(v127) {
								v150 = F_ReadBufferExtended(m, l0, int32(1), v112, int32(3), int32(0))
								mBase = m.M
								v151 = m.ExcPending
								if v151 != 0 {
									return int32(0)
								} else {
									v153 = v150
									if v153 < int32(0) {
										v159 = (v153 ^ int32(-1)) << (uint(int32(2)) % 32)
										v161 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
										v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v161)))
										v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
										if v164 != 0 {
											v253 = v153
											m.G0 = v11 + int32(48)
											return v253
										} else {
											F_LockBuffer(m, v153, int32(2))
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int32(0)
											} else {
												v169 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
												v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v159)))
												v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+14)))
												if v172 == int32(0) {
													v195 = v171
													v196 = int32(_a_F_fsm_readbuf_0)
													v197 = int32(0)
													if v197|(v195&int32(3)|int32(1)) == v197 {
														v214 = v195 + v196
														v216 = v195 + int32(4)
														if base.Ui32(v216) < base.Ui32(v214) {
															v218 = v214
														} else {
															v218 = v216
														}
														v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
														if v223 == int32(0) {
														} else {
															base.MemoryFill(m, v195, int32(0), v223)
														}
													} else {
														base.MemoryFill(m, v195, int32(0), v196)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
													v237 = int32(_a_F_fsm_readbuf_2)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
													v243 = int32(_a_F_fsm_readbuf_0)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
												} else {
												}
												F_LockBuffer(m, v153, int32(0))
												mBase = m.M
												v250 = m.ExcPending
												if v250 != 0 {
													return int32(0)
												} else {
													v253 = v153
													m.G0 = v11 + int32(48)
													return v253
												}
											}
										}
									} else {
										v176 = v153 << (uint(int32(13)) % 32)
										v178 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
										v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+v178-int32(_a_F_fsm_readbuf_3)))))
										if v182 != 0 {
											v253 = v153
											m.G0 = v11 + int32(48)
											return v253
										} else {
											F_LockBuffer(m, v153, int32(2))
											mBase = m.M
											v185 = m.ExcPending
											if v185 != 0 {
												return int32(0)
											} else {
												v187 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
												v188 = v187 + v176
												v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188-int32(_a_F_fsm_readbuf_3)))))
												if v191 != 0 {
												} else {
													v195 = v188 + int32(-8192)
													v196 = int32(_a_F_fsm_readbuf_0)
													v197 = int32(0)
													if v197|(v195&int32(3)|int32(1)) == v197 {
														v214 = v195 + v196
														v216 = v195 + int32(4)
														if base.Ui32(v216) < base.Ui32(v214) {
															v218 = v214
														} else {
															v218 = v216
														}
														v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
														if v223 == int32(0) {
														} else {
															base.MemoryFill(m, v195, int32(0), v223)
														}
													} else {
														base.MemoryFill(m, v195, int32(0), v196)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
													v237 = int32(_a_F_fsm_readbuf_2)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
													v243 = int32(_a_F_fsm_readbuf_0)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
												}
												F_LockBuffer(m, v153, int32(0))
												mBase = m.M
												v250 = m.ExcPending
												if v250 != 0 {
													return int32(0)
												} else {
													v253 = v153
													m.G0 = v11 + int32(48)
													return v253
												}
											}
										}
									}
								}
							} else {
								v129 = int32(0)
								if l2 == v129 {
									v253 = v129
									m.G0 = v11 + int32(48)
									return v253
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = int64(0)
									*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = l0
									v135 = *(*int64)(unsafe.Add(mBase, uint32(v11)+36))
									*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v135
									v137 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v137
									v145 = F_ExtendBufferedRelTo(m, v11+int32(8), int32(1), int32(20), v76-v14, int32(3))
									mBase = m.M
									v146 = m.ExcPending
									if v146 != 0 {
										return int32(0)
									} else {
										v153 = v145
										if v153 < int32(0) {
											v159 = (v153 ^ int32(-1)) << (uint(int32(2)) % 32)
											v161 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
											v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v161)))
											v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
											if v164 != 0 {
												v253 = v153
												m.G0 = v11 + int32(48)
												return v253
											} else {
												F_LockBuffer(m, v153, int32(2))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return int32(0)
												} else {
													v169 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
													v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v159)))
													v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+14)))
													if v172 == int32(0) {
														v195 = v171
														v196 = int32(_a_F_fsm_readbuf_0)
														v197 = int32(0)
														if v197|(v195&int32(3)|int32(1)) == v197 {
															v214 = v195 + v196
															v216 = v195 + int32(4)
															if base.Ui32(v216) < base.Ui32(v214) {
																v218 = v214
															} else {
																v218 = v216
															}
															v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
															if v223 == int32(0) {
															} else {
																base.MemoryFill(m, v195, int32(0), v223)
															}
														} else {
															base.MemoryFill(m, v195, int32(0), v196)
														}
														*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
														v237 = int32(_a_F_fsm_readbuf_2)
														*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
														v243 = int32(_a_F_fsm_readbuf_0)
														*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
														*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
													} else {
													}
													F_LockBuffer(m, v153, int32(0))
													mBase = m.M
													v250 = m.ExcPending
													if v250 != 0 {
														return int32(0)
													} else {
														v253 = v153
														m.G0 = v11 + int32(48)
														return v253
													}
												}
											}
										} else {
											v176 = v153 << (uint(int32(13)) % 32)
											v178 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
											v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+v178-int32(_a_F_fsm_readbuf_3)))))
											if v182 != 0 {
												v253 = v153
												m.G0 = v11 + int32(48)
												return v253
											} else {
												F_LockBuffer(m, v153, int32(2))
												mBase = m.M
												v185 = m.ExcPending
												if v185 != 0 {
													return int32(0)
												} else {
													v187 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
													v188 = v187 + v176
													v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188-int32(_a_F_fsm_readbuf_3)))))
													if v191 != 0 {
													} else {
														v195 = v188 + int32(-8192)
														v196 = int32(_a_F_fsm_readbuf_0)
														v197 = int32(0)
														if v197|(v195&int32(3)|int32(1)) == v197 {
															v214 = v195 + v196
															v216 = v195 + int32(4)
															if base.Ui32(v216) < base.Ui32(v214) {
																v218 = v214
															} else {
																v218 = v216
															}
															v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
															if v223 == int32(0) {
															} else {
																base.MemoryFill(m, v195, int32(0), v223)
															}
														} else {
															base.MemoryFill(m, v195, int32(0), v196)
														}
														*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
														v237 = int32(_a_F_fsm_readbuf_2)
														*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
														v243 = int32(_a_F_fsm_readbuf_0)
														*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
														*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
													}
													F_LockBuffer(m, v153, int32(0))
													mBase = m.M
													v250 = m.ExcPending
													if v250 != 0 {
														return int32(0)
													} else {
														v253 = v153
														m.G0 = v11 + int32(48)
														return v253
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
		v108 = v79
		v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+24))
		v112 = v76 + (v14 ^ int32(-1))
		if base.B2i32(v109 != int32(-1))&base.B2i32(base.Ui32(v112) < base.Ui32(v109)) != 0 {
			v150 = F_ReadBufferExtended(m, l0, int32(1), v112, int32(3), int32(0))
			mBase = m.M
			v151 = m.ExcPending
			if v151 != 0 {
				return int32(0)
			} else {
				v153 = v150
				if v153 < int32(0) {
					v159 = (v153 ^ int32(-1)) << (uint(int32(2)) % 32)
					v161 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
					v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v161)))
					v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
					if v164 != 0 {
						v253 = v153
						m.G0 = v11 + int32(48)
						return v253
					} else {
						F_LockBuffer(m, v153, int32(2))
						mBase = m.M
						v167 = m.ExcPending
						if v167 != 0 {
							return int32(0)
						} else {
							v169 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
							v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v159)))
							v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+14)))
							if v172 == int32(0) {
								v195 = v171
								v196 = int32(_a_F_fsm_readbuf_0)
								v197 = int32(0)
								if v197|(v195&int32(3)|int32(1)) == v197 {
									v214 = v195 + v196
									v216 = v195 + int32(4)
									if base.Ui32(v216) < base.Ui32(v214) {
										v218 = v214
									} else {
										v218 = v216
									}
									v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
									if v223 == int32(0) {
									} else {
										base.MemoryFill(m, v195, int32(0), v223)
									}
								} else {
									base.MemoryFill(m, v195, int32(0), v196)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
								v237 = int32(_a_F_fsm_readbuf_2)
								*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
								v243 = int32(_a_F_fsm_readbuf_0)
								*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
								*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
							} else {
							}
							F_LockBuffer(m, v153, int32(0))
							mBase = m.M
							v250 = m.ExcPending
							if v250 != 0 {
								return int32(0)
							} else {
								v253 = v153
								m.G0 = v11 + int32(48)
								return v253
							}
						}
					}
				} else {
					v176 = v153 << (uint(int32(13)) % 32)
					v178 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
					v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+v178-int32(_a_F_fsm_readbuf_3)))))
					if v182 != 0 {
						v253 = v153
						m.G0 = v11 + int32(48)
						return v253
					} else {
						F_LockBuffer(m, v153, int32(2))
						mBase = m.M
						v185 = m.ExcPending
						if v185 != 0 {
							return int32(0)
						} else {
							v187 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
							v188 = v187 + v176
							v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188-int32(_a_F_fsm_readbuf_3)))))
							if v191 != 0 {
							} else {
								v195 = v188 + int32(-8192)
								v196 = int32(_a_F_fsm_readbuf_0)
								v197 = int32(0)
								if v197|(v195&int32(3)|int32(1)) == v197 {
									v214 = v195 + v196
									v216 = v195 + int32(4)
									if base.Ui32(v216) < base.Ui32(v214) {
										v218 = v214
									} else {
										v218 = v216
									}
									v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
									if v223 == int32(0) {
									} else {
										base.MemoryFill(m, v195, int32(0), v223)
									}
								} else {
									base.MemoryFill(m, v195, int32(0), v196)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
								v237 = int32(_a_F_fsm_readbuf_2)
								*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
								v243 = int32(_a_F_fsm_readbuf_0)
								*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
								*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
							}
							F_LockBuffer(m, v153, int32(0))
							mBase = m.M
							v250 = m.ExcPending
							if v250 != 0 {
								return int32(0)
							} else {
								v253 = v153
								m.G0 = v11 + int32(48)
								return v253
							}
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v108)+24)) = int32(-1)
			v118 = F_smgrexists(m, v108, int32(1))
			mBase = m.M
			v119 = m.ExcPending
			if v119 != 0 {
				return int32(0)
			} else {
				if v118 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v108)+24)) = int32(0)
					v129 = int32(0)
					if l2 == v129 {
						v253 = v129
						m.G0 = v11 + int32(48)
						return v253
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = l0
						v135 = *(*int64)(unsafe.Add(mBase, uint32(v11)+36))
						*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v135
						v137 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v137
						v145 = F_ExtendBufferedRelTo(m, v11+int32(8), int32(1), int32(20), v76-v14, int32(3))
						mBase = m.M
						v146 = m.ExcPending
						if v146 != 0 {
							return int32(0)
						} else {
							v153 = v145
							if v153 < int32(0) {
								v159 = (v153 ^ int32(-1)) << (uint(int32(2)) % 32)
								v161 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
								v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v161)))
								v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
								if v164 != 0 {
									v253 = v153
									m.G0 = v11 + int32(48)
									return v253
								} else {
									F_LockBuffer(m, v153, int32(2))
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return int32(0)
									} else {
										v169 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
										v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v159)))
										v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+14)))
										if v172 == int32(0) {
											v195 = v171
											v196 = int32(_a_F_fsm_readbuf_0)
											v197 = int32(0)
											if v197|(v195&int32(3)|int32(1)) == v197 {
												v214 = v195 + v196
												v216 = v195 + int32(4)
												if base.Ui32(v216) < base.Ui32(v214) {
													v218 = v214
												} else {
													v218 = v216
												}
												v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
												if v223 == int32(0) {
												} else {
													base.MemoryFill(m, v195, int32(0), v223)
												}
											} else {
												base.MemoryFill(m, v195, int32(0), v196)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
											v237 = int32(_a_F_fsm_readbuf_2)
											*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
											v243 = int32(_a_F_fsm_readbuf_0)
											*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
											*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
										} else {
										}
										F_LockBuffer(m, v153, int32(0))
										mBase = m.M
										v250 = m.ExcPending
										if v250 != 0 {
											return int32(0)
										} else {
											v253 = v153
											m.G0 = v11 + int32(48)
											return v253
										}
									}
								}
							} else {
								v176 = v153 << (uint(int32(13)) % 32)
								v178 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
								v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+v178-int32(_a_F_fsm_readbuf_3)))))
								if v182 != 0 {
									v253 = v153
									m.G0 = v11 + int32(48)
									return v253
								} else {
									F_LockBuffer(m, v153, int32(2))
									mBase = m.M
									v185 = m.ExcPending
									if v185 != 0 {
										return int32(0)
									} else {
										v187 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
										v188 = v187 + v176
										v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188-int32(_a_F_fsm_readbuf_3)))))
										if v191 != 0 {
										} else {
											v195 = v188 + int32(-8192)
											v196 = int32(_a_F_fsm_readbuf_0)
											v197 = int32(0)
											if v197|(v195&int32(3)|int32(1)) == v197 {
												v214 = v195 + v196
												v216 = v195 + int32(4)
												if base.Ui32(v216) < base.Ui32(v214) {
													v218 = v214
												} else {
													v218 = v216
												}
												v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
												if v223 == int32(0) {
												} else {
													base.MemoryFill(m, v195, int32(0), v223)
												}
											} else {
												base.MemoryFill(m, v195, int32(0), v196)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
											v237 = int32(_a_F_fsm_readbuf_2)
											*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
											v243 = int32(_a_F_fsm_readbuf_0)
											*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
											*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
										}
										F_LockBuffer(m, v153, int32(0))
										mBase = m.M
										v250 = m.ExcPending
										if v250 != 0 {
											return int32(0)
										} else {
											v253 = v153
											m.G0 = v11 + int32(48)
											return v253
										}
									}
								}
							}
						}
					}
				} else {
					v125 = F_smgrnblocks(m, v108, int32(1))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						v127 = *(*int32)(unsafe.Add(mBase, uint32(v108)+24))
						if base.Ui32(v112) < base.Ui32(v127) {
							v150 = F_ReadBufferExtended(m, l0, int32(1), v112, int32(3), int32(0))
							mBase = m.M
							v151 = m.ExcPending
							if v151 != 0 {
								return int32(0)
							} else {
								v153 = v150
								if v153 < int32(0) {
									v159 = (v153 ^ int32(-1)) << (uint(int32(2)) % 32)
									v161 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
									v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v161)))
									v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
									if v164 != 0 {
										v253 = v153
										m.G0 = v11 + int32(48)
										return v253
									} else {
										F_LockBuffer(m, v153, int32(2))
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return int32(0)
										} else {
											v169 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
											v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v159)))
											v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+14)))
											if v172 == int32(0) {
												v195 = v171
												v196 = int32(_a_F_fsm_readbuf_0)
												v197 = int32(0)
												if v197|(v195&int32(3)|int32(1)) == v197 {
													v214 = v195 + v196
													v216 = v195 + int32(4)
													if base.Ui32(v216) < base.Ui32(v214) {
														v218 = v214
													} else {
														v218 = v216
													}
													v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
													if v223 == int32(0) {
													} else {
														base.MemoryFill(m, v195, int32(0), v223)
													}
												} else {
													base.MemoryFill(m, v195, int32(0), v196)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
												v237 = int32(_a_F_fsm_readbuf_2)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
												v243 = int32(_a_F_fsm_readbuf_0)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
											} else {
											}
											F_LockBuffer(m, v153, int32(0))
											mBase = m.M
											v250 = m.ExcPending
											if v250 != 0 {
												return int32(0)
											} else {
												v253 = v153
												m.G0 = v11 + int32(48)
												return v253
											}
										}
									}
								} else {
									v176 = v153 << (uint(int32(13)) % 32)
									v178 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
									v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+v178-int32(_a_F_fsm_readbuf_3)))))
									if v182 != 0 {
										v253 = v153
										m.G0 = v11 + int32(48)
										return v253
									} else {
										F_LockBuffer(m, v153, int32(2))
										mBase = m.M
										v185 = m.ExcPending
										if v185 != 0 {
											return int32(0)
										} else {
											v187 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
											v188 = v187 + v176
											v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188-int32(_a_F_fsm_readbuf_3)))))
											if v191 != 0 {
											} else {
												v195 = v188 + int32(-8192)
												v196 = int32(_a_F_fsm_readbuf_0)
												v197 = int32(0)
												if v197|(v195&int32(3)|int32(1)) == v197 {
													v214 = v195 + v196
													v216 = v195 + int32(4)
													if base.Ui32(v216) < base.Ui32(v214) {
														v218 = v214
													} else {
														v218 = v216
													}
													v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
													if v223 == int32(0) {
													} else {
														base.MemoryFill(m, v195, int32(0), v223)
													}
												} else {
													base.MemoryFill(m, v195, int32(0), v196)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
												v237 = int32(_a_F_fsm_readbuf_2)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
												v243 = int32(_a_F_fsm_readbuf_0)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
											}
											F_LockBuffer(m, v153, int32(0))
											mBase = m.M
											v250 = m.ExcPending
											if v250 != 0 {
												return int32(0)
											} else {
												v253 = v153
												m.G0 = v11 + int32(48)
												return v253
											}
										}
									}
								}
							}
						} else {
							v129 = int32(0)
							if l2 == v129 {
								v253 = v129
								m.G0 = v11 + int32(48)
								return v253
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = l0
								v135 = *(*int64)(unsafe.Add(mBase, uint32(v11)+36))
								*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v135
								v137 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v137
								v145 = F_ExtendBufferedRelTo(m, v11+int32(8), int32(1), int32(20), v76-v14, int32(3))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return int32(0)
								} else {
									v153 = v145
									if v153 < int32(0) {
										v159 = (v153 ^ int32(-1)) << (uint(int32(2)) % 32)
										v161 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
										v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v161)))
										v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
										if v164 != 0 {
											v253 = v153
											m.G0 = v11 + int32(48)
											return v253
										} else {
											F_LockBuffer(m, v153, int32(2))
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int32(0)
											} else {
												v169 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
												v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v159)))
												v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+14)))
												if v172 == int32(0) {
													v195 = v171
													v196 = int32(_a_F_fsm_readbuf_0)
													v197 = int32(0)
													if v197|(v195&int32(3)|int32(1)) == v197 {
														v214 = v195 + v196
														v216 = v195 + int32(4)
														if base.Ui32(v216) < base.Ui32(v214) {
															v218 = v214
														} else {
															v218 = v216
														}
														v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
														if v223 == int32(0) {
														} else {
															base.MemoryFill(m, v195, int32(0), v223)
														}
													} else {
														base.MemoryFill(m, v195, int32(0), v196)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
													v237 = int32(_a_F_fsm_readbuf_2)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
													v243 = int32(_a_F_fsm_readbuf_0)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
												} else {
												}
												F_LockBuffer(m, v153, int32(0))
												mBase = m.M
												v250 = m.ExcPending
												if v250 != 0 {
													return int32(0)
												} else {
													v253 = v153
													m.G0 = v11 + int32(48)
													return v253
												}
											}
										}
									} else {
										v176 = v153 << (uint(int32(13)) % 32)
										v178 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
										v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+v178-int32(_a_F_fsm_readbuf_3)))))
										if v182 != 0 {
											v253 = v153
											m.G0 = v11 + int32(48)
											return v253
										} else {
											F_LockBuffer(m, v153, int32(2))
											mBase = m.M
											v185 = m.ExcPending
											if v185 != 0 {
												return int32(0)
											} else {
												v187 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
												v188 = v187 + v176
												v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188-int32(_a_F_fsm_readbuf_3)))))
												if v191 != 0 {
												} else {
													v195 = v188 + int32(-8192)
													v196 = int32(_a_F_fsm_readbuf_0)
													v197 = int32(0)
													if v197|(v195&int32(3)|int32(1)) == v197 {
														v214 = v195 + v196
														v216 = v195 + int32(4)
														if base.Ui32(v216) < base.Ui32(v214) {
															v218 = v214
														} else {
															v218 = v216
														}
														v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
														if v223 == int32(0) {
														} else {
															base.MemoryFill(m, v195, int32(0), v223)
														}
													} else {
														base.MemoryFill(m, v195, int32(0), v196)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
													v237 = int32(_a_F_fsm_readbuf_2)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
													v243 = int32(_a_F_fsm_readbuf_0)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
												}
												F_LockBuffer(m, v153, int32(0))
												mBase = m.M
												v250 = m.ExcPending
												if v250 != 0 {
													return int32(0)
												} else {
													v253 = v153
													m.G0 = v11 + int32(48)
													return v253
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
