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
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
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
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	v3 = int32(0)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	if v9 == int32(1) {
		v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+74)))
		v17 = l0 + v12<<(uint(int32(2))%32) + int32(80)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v19 != int32(-1) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
			v33 = v19
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
			v44 = F_StartReadBuffer(m, v34+int32(4), v17, v33, v37|v38<<(uint(int32(1))%32)&int32(2))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				if v44 == int32(0) {
					v211 = v18
				} else {
					v48 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)) = uint16(v48)
					v50 = int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v50)
					v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)) = uint16(base.B2i32(v50 < v52))
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
					*(*uint16)(unsafe.Add(mBase, uint32(v56))) = uint16(v12)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v33 + v50
					v202 = v3
					v203 = v18
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v202)
					v211 = v203
				}
				return v211
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v27 = m.T0[v26].(func(*base.Module, int32, int32, int32) int32)(m, l0, v24, int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				if v27 == int32(-1) {
					v61 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v61
					v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+76)))
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)) = uint16(v63)
					*(*int32)(unsafe.Add(mBase, uint32(v17))) = v61
					v202 = v3
					v203 = v18
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v202)
					v211 = v203
					return v211
				} else {
					v33 = v27
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
					v44 = F_StartReadBuffer(m, v34+int32(4), v17, v33, v37|v38<<(uint(int32(1))%32)&int32(2))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						if v44 == int32(0) {
							v211 = v18
						} else {
							v48 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)) = uint16(v48)
							v50 = int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v50)
							v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)) = uint16(base.B2i32(v50 < v52))
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							*(*uint16)(unsafe.Add(mBase, uint32(v56))) = uint16(v12)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v33 + v50
							v202 = v3
							v203 = v18
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v202)
							v211 = v203
						}
						return v211
					}
				}
			}
		}
	} else {
		v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
		if v67 == int32(0) {
			v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
			if v70 == int32(0) {
				v211 = v3
				return v211
			} else {
				F_read_stream_look_ahead(m, l0)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
					if v75 == int32(0) {
						v211 = v3
						return v211
					} else {
						v79 = l0 + int32(80)
						v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+74)))
						v83 = v79 + v80<<(uint(int32(2))%32)
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
						if l1 != 0 {
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
							v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v85 + v86*v80
						} else {
						}
						v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
						if v90 <= int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
							v144 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
							if v80 < v144-int32(1) {
								v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
								v149 = int32(2)
								*(*int32)(unsafe.Add(mBase, uint32(v79+v148<<(uint(v149)%32)+v80<<(uint(v149)%32)))) = int32(0)
							} else {
							}
							v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
							v158 = int32(1)
							v159 = v157 - v158
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v159)
							v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)))
							v163 = v161 + v158
							v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
							if v165 != v163&int32(_a_F_read_stream_next_buffer_0) {
								v169 = v163
							} else {
								v169 = int32(0)
							}
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)) = uint16(v169)
							F_read_stream_look_ahead(m, l0)
							mBase = m.M
							v172 = m.ExcPending
							if v172 != 0 {
								return int32(0)
							} else {
								v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
								if v173 != 0 {
									v211 = v84
								} else {
									v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
									if v174 != 0 {
										v211 = v84
									} else {
										v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
										if v175 != int32(1) {
											v211 = v84
										} else {
											v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
											if v178 != int32(1) {
												v211 = v84
											} else {
												v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
												if v181 != 0 {
													v211 = v84
												} else {
													v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
													if v182 != 0 {
														v211 = v84
													} else {
														v183 = int32(1)
														v184 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+74)))
														v185 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
														if v185-v183 <= v184 {
															v202 = v183
															v203 = v84
														} else {
															v189 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
															v190 = int32(2)
															*(*int32)(unsafe.Add(mBase, uint32(v79+v189<<(uint(v190)%32)+v184<<(uint(v190)%32)))) = int32(0)
															v202 = v183
															v203 = v84
														}
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v202)
														v211 = v203
													}
												}
											}
										}
									}
								}
								return v211
							}
						} else {
							v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+68)))
							v97 = v93 + v94*int32(84)
							v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97))))
							if v98 != v80&int32(_a_F_read_stream_next_buffer_0) {
								*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
								v144 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
								if v80 < v144-int32(1) {
									v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
									v149 = int32(2)
									*(*int32)(unsafe.Add(mBase, uint32(v79+v148<<(uint(v149)%32)+v80<<(uint(v149)%32)))) = int32(0)
								} else {
								}
								v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
								v158 = int32(1)
								v159 = v157 - v158
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v159)
								v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)))
								v163 = v161 + v158
								v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
								if v165 != v163&int32(_a_F_read_stream_next_buffer_0) {
									v169 = v163
								} else {
									v169 = int32(0)
								}
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)) = uint16(v169)
								F_read_stream_look_ahead(m, l0)
								mBase = m.M
								v172 = m.ExcPending
								if v172 != 0 {
									return int32(0)
								} else {
									v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
									if v173 != 0 {
										v211 = v84
									} else {
										v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
										if v174 != 0 {
											v211 = v84
										} else {
											v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
											if v175 != int32(1) {
												v211 = v84
											} else {
												v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
												if v178 != int32(1) {
													v211 = v84
												} else {
													v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
													if v181 != 0 {
														v211 = v84
													} else {
														v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
														if v182 != 0 {
															v211 = v84
														} else {
															v183 = int32(1)
															v184 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+74)))
															v185 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
															if v185-v183 <= v184 {
																v202 = v183
																v203 = v84
															} else {
																v189 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
																v190 = int32(2)
																*(*int32)(unsafe.Add(mBase, uint32(v79+v189<<(uint(v190)%32)+v184<<(uint(v190)%32)))) = int32(0)
																v202 = v183
																v203 = v84
															}
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v202)
															v211 = v203
														}
													}
												}
											}
										}
									}
									return v211
								}
							} else {
								F_WaitReadBuffers(m, v97+int32(4))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
									v107 = int32(1)
									v108 = v106 - v107
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v108)
									v110 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+14)))
									v112 = v110 << (uint(v107) % 32)
									v113 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
									if v112 < v113 {
										v115 = v112
									} else {
										v115 = v113
									}
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v115)
									v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)))
									v119 = v117 + int32(1)
									v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
									if v121 != v119&int32(_a_F_read_stream_next_buffer_0) {
										v125 = v119
									} else {
										v125 = int32(0)
									}
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)) = uint16(v125)
									v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
									if v127 != int32(1) {
									} else {
										v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
										v134 = *(*int32)(unsafe.Add(mBase, uint32(v130+v94*int32(84))+28))
										v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										if v134 != v135 {
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(-1)
										}
									}
									*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
									v144 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
									if v80 < v144-int32(1) {
										v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
										v149 = int32(2)
										*(*int32)(unsafe.Add(mBase, uint32(v79+v148<<(uint(v149)%32)+v80<<(uint(v149)%32)))) = int32(0)
									} else {
									}
									v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
									v158 = int32(1)
									v159 = v157 - v158
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v159)
									v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)))
									v163 = v161 + v158
									v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
									if v165 != v163&int32(_a_F_read_stream_next_buffer_0) {
										v169 = v163
									} else {
										v169 = int32(0)
									}
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)) = uint16(v169)
									F_read_stream_look_ahead(m, l0)
									mBase = m.M
									v172 = m.ExcPending
									if v172 != 0 {
										return int32(0)
									} else {
										v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
										if v173 != 0 {
											v211 = v84
										} else {
											v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
											if v174 != 0 {
												v211 = v84
											} else {
												v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
												if v175 != int32(1) {
													v211 = v84
												} else {
													v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
													if v178 != int32(1) {
														v211 = v84
													} else {
														v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
														if v181 != 0 {
															v211 = v84
														} else {
															v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
															if v182 != 0 {
																v211 = v84
															} else {
																v183 = int32(1)
																v184 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+74)))
																v185 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
																if v185-v183 <= v184 {
																	v202 = v183
																	v203 = v84
																} else {
																	v189 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
																	v190 = int32(2)
																	*(*int32)(unsafe.Add(mBase, uint32(v79+v189<<(uint(v190)%32)+v184<<(uint(v190)%32)))) = int32(0)
																	v202 = v183
																	v203 = v84
																}
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v202)
																v211 = v203
															}
														}
													}
												}
											}
										}
										return v211
									}
								}
							}
						}
					}
				}
			}
		} else {
			v79 = l0 + int32(80)
			v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+74)))
			v83 = v79 + v80<<(uint(int32(2))%32)
			v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
			if l1 != 0 {
				v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v85 + v86*v80
			} else {
			}
			v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
			if v90 <= int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
				v144 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
				if v80 < v144-int32(1) {
					v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
					v149 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v79+v148<<(uint(v149)%32)+v80<<(uint(v149)%32)))) = int32(0)
				} else {
				}
				v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				v158 = int32(1)
				v159 = v157 - v158
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v159)
				v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)))
				v163 = v161 + v158
				v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
				if v165 != v163&int32(_a_F_read_stream_next_buffer_0) {
					v169 = v163
				} else {
					v169 = int32(0)
				}
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)) = uint16(v169)
				F_read_stream_look_ahead(m, l0)
				mBase = m.M
				v172 = m.ExcPending
				if v172 != 0 {
					return int32(0)
				} else {
					v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
					if v173 != 0 {
						v211 = v84
					} else {
						v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
						if v174 != 0 {
							v211 = v84
						} else {
							v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
							if v175 != int32(1) {
								v211 = v84
							} else {
								v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
								if v178 != int32(1) {
									v211 = v84
								} else {
									v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
									if v181 != 0 {
										v211 = v84
									} else {
										v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v182 != 0 {
											v211 = v84
										} else {
											v183 = int32(1)
											v184 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+74)))
											v185 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
											if v185-v183 <= v184 {
												v202 = v183
												v203 = v84
											} else {
												v189 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
												v190 = int32(2)
												*(*int32)(unsafe.Add(mBase, uint32(v79+v189<<(uint(v190)%32)+v184<<(uint(v190)%32)))) = int32(0)
												v202 = v183
												v203 = v84
											}
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v202)
											v211 = v203
										}
									}
								}
							}
						}
					}
					return v211
				}
			} else {
				v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+68)))
				v97 = v93 + v94*int32(84)
				v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97))))
				if v98 != v80&int32(_a_F_read_stream_next_buffer_0) {
					*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
					v144 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
					if v80 < v144-int32(1) {
						v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
						v149 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v79+v148<<(uint(v149)%32)+v80<<(uint(v149)%32)))) = int32(0)
					} else {
					}
					v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
					v158 = int32(1)
					v159 = v157 - v158
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v159)
					v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)))
					v163 = v161 + v158
					v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
					if v165 != v163&int32(_a_F_read_stream_next_buffer_0) {
						v169 = v163
					} else {
						v169 = int32(0)
					}
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)) = uint16(v169)
					F_read_stream_look_ahead(m, l0)
					mBase = m.M
					v172 = m.ExcPending
					if v172 != 0 {
						return int32(0)
					} else {
						v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
						if v173 != 0 {
							v211 = v84
						} else {
							v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
							if v174 != 0 {
								v211 = v84
							} else {
								v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
								if v175 != int32(1) {
									v211 = v84
								} else {
									v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
									if v178 != int32(1) {
										v211 = v84
									} else {
										v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
										if v181 != 0 {
											v211 = v84
										} else {
											v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
											if v182 != 0 {
												v211 = v84
											} else {
												v183 = int32(1)
												v184 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+74)))
												v185 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
												if v185-v183 <= v184 {
													v202 = v183
													v203 = v84
												} else {
													v189 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
													v190 = int32(2)
													*(*int32)(unsafe.Add(mBase, uint32(v79+v189<<(uint(v190)%32)+v184<<(uint(v190)%32)))) = int32(0)
													v202 = v183
													v203 = v84
												}
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v202)
												v211 = v203
											}
										}
									}
								}
							}
						}
						return v211
					}
				} else {
					F_WaitReadBuffers(m, v97+int32(4))
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return int32(0)
					} else {
						v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
						v107 = int32(1)
						v108 = v106 - v107
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v108)
						v110 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+14)))
						v112 = v110 << (uint(v107) % 32)
						v113 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
						if v112 < v113 {
							v115 = v112
						} else {
							v115 = v113
						}
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v115)
						v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)))
						v119 = v117 + int32(1)
						v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
						if v121 != v119&int32(_a_F_read_stream_next_buffer_0) {
							v125 = v119
						} else {
							v125 = int32(0)
						}
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)) = uint16(v125)
						v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
						if v127 != int32(1) {
						} else {
							v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							v134 = *(*int32)(unsafe.Add(mBase, uint32(v130+v94*int32(84))+28))
							v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							if v134 != v135 {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(-1)
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
						v144 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
						if v80 < v144-int32(1) {
							v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
							v149 = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v79+v148<<(uint(v149)%32)+v80<<(uint(v149)%32)))) = int32(0)
						} else {
						}
						v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
						v158 = int32(1)
						v159 = v157 - v158
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v159)
						v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)))
						v163 = v161 + v158
						v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
						if v165 != v163&int32(_a_F_read_stream_next_buffer_0) {
							v169 = v163
						} else {
							v169 = int32(0)
						}
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)) = uint16(v169)
						F_read_stream_look_ahead(m, l0)
						mBase = m.M
						v172 = m.ExcPending
						if v172 != 0 {
							return int32(0)
						} else {
							v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
							if v173 != 0 {
								v211 = v84
							} else {
								v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
								if v174 != 0 {
									v211 = v84
								} else {
									v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
									if v175 != int32(1) {
										v211 = v84
									} else {
										v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
										if v178 != int32(1) {
											v211 = v84
										} else {
											v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
											if v181 != 0 {
												v211 = v84
											} else {
												v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
												if v182 != 0 {
													v211 = v84
												} else {
													v183 = int32(1)
													v184 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+74)))
													v185 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
													if v185-v183 <= v184 {
														v202 = v183
														v203 = v84
													} else {
														v189 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
														v190 = int32(2)
														*(*int32)(unsafe.Add(mBase, uint32(v79+v189<<(uint(v190)%32)+v184<<(uint(v190)%32)))) = int32(0)
														v202 = v183
														v203 = v84
													}
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v202)
													v211 = v203
												}
											}
										}
									}
								}
							}
							return v211
						}
					}
				}
			}
		}
	}
}
