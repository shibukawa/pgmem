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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
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
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
			if v40 != 0 {
				v41 = v37 | int32(2)
			} else {
				v41 = v37
			}
			v42 = F_StartReadBuffer(m, v34+int32(4), v17, v33, v41)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				if v42 == int32(0) {
					v205 = v18
				} else {
					v46 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)) = uint16(v46)
					v49 = int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v49)
					v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)) = uint16(base.B2i32(v49 < v51))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
					*(*uint16)(unsafe.Add(mBase, uint32(v55))) = uint16(v12)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v33 + v49
					v195 = v46
					v197 = v18
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v195)
					v205 = v197
				}
				return v205
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
					v60 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v60
					v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+76)))
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)) = uint16(v62)
					*(*int32)(unsafe.Add(mBase, uint32(v17))) = v60
					v195 = v3
					v197 = v18
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v195)
					v205 = v197
					return v205
				} else {
					v33 = v27
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
					if v40 != 0 {
						v41 = v37 | int32(2)
					} else {
						v41 = v37
					}
					v42 = F_StartReadBuffer(m, v34+int32(4), v17, v33, v41)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						if v42 == int32(0) {
							v205 = v18
						} else {
							v46 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)) = uint16(v46)
							v49 = int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v49)
							v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)) = uint16(base.B2i32(v49 < v51))
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							*(*uint16)(unsafe.Add(mBase, uint32(v55))) = uint16(v12)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v33 + v49
							v195 = v46
							v197 = v18
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v195)
							v205 = v197
						}
						return v205
					}
				}
			}
		}
	} else {
		v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
		if v66 == int32(0) {
			v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
			if v69 == int32(0) {
				v205 = v3
				return v205
			} else {
				F_read_stream_look_ahead(m, l0)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
					if v74 == int32(0) {
						v205 = v3
						return v205
					} else {
						v78 = l0 + int32(80)
						v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+74)))
						v82 = v78 + v79<<(uint(int32(2))%32)
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
						if l1 != 0 {
							v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v84 + v85*v79
						} else {
						}
						v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
						if v89 <= int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(0)
							v143 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
							if v79 < v143-int32(1) {
								v147 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
								*(*int32)(unsafe.Add(mBase, uint32(v78+(v147+v79)<<(uint(int32(2))%32)))) = int32(0)
							} else {
							}
							v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
							v155 = int32(1)
							v156 = v154 - v155
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v156)
							v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)))
							v160 = v158 + v155
							v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
							if v162 != v160&int32(_a_F_read_stream_next_buffer_0) {
								v166 = v160
							} else {
								v166 = int32(0)
							}
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)) = uint16(v166)
							F_read_stream_look_ahead(m, l0)
							mBase = m.M
							v169 = m.ExcPending
							if v169 != 0 {
								return int32(0)
							} else {
								v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
								if v170 != 0 {
									v205 = v83
								} else {
									v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
									if v171 != 0 {
										v205 = v83
									} else {
										v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
										if v172 != int32(1) {
											v205 = v83
										} else {
											v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
											if v175 != int32(1) {
												v205 = v83
											} else {
												v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
												if v178 != 0 {
													v205 = v83
												} else {
													v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
													if v179 != 0 {
														v205 = v83
													} else {
														v180 = int32(1)
														v181 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+74)))
														v182 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
														if v182-v180 <= v181 {
															v195 = v180
															v197 = v83
														} else {
															v186 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
															*(*int32)(unsafe.Add(mBase, uint32(v78+(v186+v181)<<(uint(int32(2))%32)))) = int32(0)
															v195 = v180
															v197 = v83
														}
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v195)
														v205 = v197
													}
												}
											}
										}
									}
								}
								return v205
							}
						} else {
							v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							v93 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+68)))
							v96 = v92 + v93*int32(84)
							v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96))))
							if v97 != v79&int32(_a_F_read_stream_next_buffer_0) {
								*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(0)
								v143 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
								if v79 < v143-int32(1) {
									v147 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
									*(*int32)(unsafe.Add(mBase, uint32(v78+(v147+v79)<<(uint(int32(2))%32)))) = int32(0)
								} else {
								}
								v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
								v155 = int32(1)
								v156 = v154 - v155
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v156)
								v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)))
								v160 = v158 + v155
								v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
								if v162 != v160&int32(_a_F_read_stream_next_buffer_0) {
									v166 = v160
								} else {
									v166 = int32(0)
								}
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)) = uint16(v166)
								F_read_stream_look_ahead(m, l0)
								mBase = m.M
								v169 = m.ExcPending
								if v169 != 0 {
									return int32(0)
								} else {
									v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
									if v170 != 0 {
										v205 = v83
									} else {
										v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
										if v171 != 0 {
											v205 = v83
										} else {
											v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
											if v172 != int32(1) {
												v205 = v83
											} else {
												v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
												if v175 != int32(1) {
													v205 = v83
												} else {
													v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
													if v178 != 0 {
														v205 = v83
													} else {
														v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
														if v179 != 0 {
															v205 = v83
														} else {
															v180 = int32(1)
															v181 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+74)))
															v182 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
															if v182-v180 <= v181 {
																v195 = v180
																v197 = v83
															} else {
																v186 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
																*(*int32)(unsafe.Add(mBase, uint32(v78+(v186+v181)<<(uint(int32(2))%32)))) = int32(0)
																v195 = v180
																v197 = v83
															}
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v195)
															v205 = v197
														}
													}
												}
											}
										}
									}
									return v205
								}
							} else {
								F_WaitReadBuffers(m, v96+int32(4))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
									v106 = int32(1)
									v107 = v105 - v106
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v107)
									v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+14)))
									v111 = v109 << (uint(v106) % 32)
									v112 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
									if v111 < v112 {
										v114 = v111
									} else {
										v114 = v112
									}
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v114)
									v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)))
									v118 = v116 + int32(1)
									v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
									if v120 != v118&int32(_a_F_read_stream_next_buffer_0) {
										v124 = v118
									} else {
										v124 = int32(0)
									}
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)) = uint16(v124)
									v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
									if v126 != int32(1) {
									} else {
										v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
										v133 = *(*int32)(unsafe.Add(mBase, uint32(v129+v93*int32(84))+28))
										v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										if v133 != v134 {
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(-1)
										}
									}
									*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(0)
									v143 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
									if v79 < v143-int32(1) {
										v147 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
										*(*int32)(unsafe.Add(mBase, uint32(v78+(v147+v79)<<(uint(int32(2))%32)))) = int32(0)
									} else {
									}
									v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
									v155 = int32(1)
									v156 = v154 - v155
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v156)
									v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)))
									v160 = v158 + v155
									v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
									if v162 != v160&int32(_a_F_read_stream_next_buffer_0) {
										v166 = v160
									} else {
										v166 = int32(0)
									}
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)) = uint16(v166)
									F_read_stream_look_ahead(m, l0)
									mBase = m.M
									v169 = m.ExcPending
									if v169 != 0 {
										return int32(0)
									} else {
										v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
										if v170 != 0 {
											v205 = v83
										} else {
											v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
											if v171 != 0 {
												v205 = v83
											} else {
												v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
												if v172 != int32(1) {
													v205 = v83
												} else {
													v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
													if v175 != int32(1) {
														v205 = v83
													} else {
														v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
														if v178 != 0 {
															v205 = v83
														} else {
															v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
															if v179 != 0 {
																v205 = v83
															} else {
																v180 = int32(1)
																v181 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+74)))
																v182 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
																if v182-v180 <= v181 {
																	v195 = v180
																	v197 = v83
																} else {
																	v186 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
																	*(*int32)(unsafe.Add(mBase, uint32(v78+(v186+v181)<<(uint(int32(2))%32)))) = int32(0)
																	v195 = v180
																	v197 = v83
																}
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v195)
																v205 = v197
															}
														}
													}
												}
											}
										}
										return v205
									}
								}
							}
						}
					}
				}
			}
		} else {
			v78 = l0 + int32(80)
			v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+74)))
			v82 = v78 + v79<<(uint(int32(2))%32)
			v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
			if l1 != 0 {
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
				v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v84 + v85*v79
			} else {
			}
			v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
			if v89 <= int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(0)
				v143 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
				if v79 < v143-int32(1) {
					v147 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
					*(*int32)(unsafe.Add(mBase, uint32(v78+(v147+v79)<<(uint(int32(2))%32)))) = int32(0)
				} else {
				}
				v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				v155 = int32(1)
				v156 = v154 - v155
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v156)
				v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)))
				v160 = v158 + v155
				v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
				if v162 != v160&int32(_a_F_read_stream_next_buffer_0) {
					v166 = v160
				} else {
					v166 = int32(0)
				}
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)) = uint16(v166)
				F_read_stream_look_ahead(m, l0)
				mBase = m.M
				v169 = m.ExcPending
				if v169 != 0 {
					return int32(0)
				} else {
					v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
					if v170 != 0 {
						v205 = v83
					} else {
						v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
						if v171 != 0 {
							v205 = v83
						} else {
							v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
							if v172 != int32(1) {
								v205 = v83
							} else {
								v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
								if v175 != int32(1) {
									v205 = v83
								} else {
									v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
									if v178 != 0 {
										v205 = v83
									} else {
										v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v179 != 0 {
											v205 = v83
										} else {
											v180 = int32(1)
											v181 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+74)))
											v182 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
											if v182-v180 <= v181 {
												v195 = v180
												v197 = v83
											} else {
												v186 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
												*(*int32)(unsafe.Add(mBase, uint32(v78+(v186+v181)<<(uint(int32(2))%32)))) = int32(0)
												v195 = v180
												v197 = v83
											}
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v195)
											v205 = v197
										}
									}
								}
							}
						}
					}
					return v205
				}
			} else {
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				v93 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+68)))
				v96 = v92 + v93*int32(84)
				v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96))))
				if v97 != v79&int32(_a_F_read_stream_next_buffer_0) {
					*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(0)
					v143 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
					if v79 < v143-int32(1) {
						v147 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
						*(*int32)(unsafe.Add(mBase, uint32(v78+(v147+v79)<<(uint(int32(2))%32)))) = int32(0)
					} else {
					}
					v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
					v155 = int32(1)
					v156 = v154 - v155
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v156)
					v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)))
					v160 = v158 + v155
					v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
					if v162 != v160&int32(_a_F_read_stream_next_buffer_0) {
						v166 = v160
					} else {
						v166 = int32(0)
					}
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)) = uint16(v166)
					F_read_stream_look_ahead(m, l0)
					mBase = m.M
					v169 = m.ExcPending
					if v169 != 0 {
						return int32(0)
					} else {
						v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
						if v170 != 0 {
							v205 = v83
						} else {
							v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
							if v171 != 0 {
								v205 = v83
							} else {
								v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
								if v172 != int32(1) {
									v205 = v83
								} else {
									v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
									if v175 != int32(1) {
										v205 = v83
									} else {
										v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
										if v178 != 0 {
											v205 = v83
										} else {
											v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
											if v179 != 0 {
												v205 = v83
											} else {
												v180 = int32(1)
												v181 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+74)))
												v182 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
												if v182-v180 <= v181 {
													v195 = v180
													v197 = v83
												} else {
													v186 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
													*(*int32)(unsafe.Add(mBase, uint32(v78+(v186+v181)<<(uint(int32(2))%32)))) = int32(0)
													v195 = v180
													v197 = v83
												}
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v195)
												v205 = v197
											}
										}
									}
								}
							}
						}
						return v205
					}
				} else {
					F_WaitReadBuffers(m, v96+int32(4))
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return int32(0)
					} else {
						v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
						v106 = int32(1)
						v107 = v105 - v106
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v107)
						v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+14)))
						v111 = v109 << (uint(v106) % 32)
						v112 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
						if v111 < v112 {
							v114 = v111
						} else {
							v114 = v112
						}
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v114)
						v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)))
						v118 = v116 + int32(1)
						v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
						if v120 != v118&int32(_a_F_read_stream_next_buffer_0) {
							v124 = v118
						} else {
							v124 = int32(0)
						}
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)) = uint16(v124)
						v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
						if v126 != int32(1) {
						} else {
							v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							v133 = *(*int32)(unsafe.Add(mBase, uint32(v129+v93*int32(84))+28))
							v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							if v133 != v134 {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(-1)
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(0)
						v143 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
						if v79 < v143-int32(1) {
							v147 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
							*(*int32)(unsafe.Add(mBase, uint32(v78+(v147+v79)<<(uint(int32(2))%32)))) = int32(0)
						} else {
						}
						v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
						v155 = int32(1)
						v156 = v154 - v155
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v156)
						v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)))
						v160 = v158 + v155
						v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
						if v162 != v160&int32(_a_F_read_stream_next_buffer_0) {
							v166 = v160
						} else {
							v166 = int32(0)
						}
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)) = uint16(v166)
						F_read_stream_look_ahead(m, l0)
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
							return int32(0)
						} else {
							v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
							if v170 != 0 {
								v205 = v83
							} else {
								v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
								if v171 != 0 {
									v205 = v83
								} else {
									v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
									if v172 != int32(1) {
										v205 = v83
									} else {
										v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
										if v175 != int32(1) {
											v205 = v83
										} else {
											v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
											if v178 != 0 {
												v205 = v83
											} else {
												v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
												if v179 != 0 {
													v205 = v83
												} else {
													v180 = int32(1)
													v181 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+74)))
													v182 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
													if v182-v180 <= v181 {
														v195 = v180
														v197 = v83
													} else {
														v186 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
														*(*int32)(unsafe.Add(mBase, uint32(v78+(v186+v181)<<(uint(int32(2))%32)))) = int32(0)
														v195 = v180
														v197 = v83
													}
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v195)
													v205 = v197
												}
											}
										}
									}
								}
							}
							return v205
						}
					}
				}
			}
		}
	}
}
