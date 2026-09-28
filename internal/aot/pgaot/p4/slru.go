package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SlruReportIOError(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	v9 = m.G0
	v11 = v9 - int32(1184)
	m.G0 = v11
	v14 = base.I64_div_s(l1, int64(32))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v16 == int32(1) {
		*(*int64)(unsafe.Add(mBase, uint32(v11)+136)) = v14
		*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v15
		v27 = F_pg_snprintf(m, v11+int32(160), int32(1024), int32(_a_F_SlruReportIOError_0), v11+int32(128))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, _c_F_SlruReportIOError[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_SlruReportIOError[1])) = v41
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_SlruReportIOError[2]))
			if v44 != int32(4) {
				v52 = base.I32_wrap_i64(l1-v14<<(uint(int64(5))%64)) << (uint(int32(13)) % 32)
				switch v44 - int32(1) {
				case 0:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v11 + int32(160)
							F_errmsg(m, int32(_a_F_SlruReportIOError_1), v11+int32(16))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								if l2 != 0 {
									v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v91 = m.T0[v90].(func(*base.Module, int32) int32)(m, l2)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1119), int32(_a_F_SlruReportIOError_3))
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
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1119), int32(_a_F_SlruReportIOError_3))
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
						}
					}
				case 1:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return
					} else {
						if v41 != 0 {
							F_errcode_for_file_access(m)
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v11 + int32(160)
								F_errmsg(m, int32(_a_F_SlruReportIOError_4), v11+int32(48))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return
								} else {
									if l2 != 0 {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v114 = m.T0[v113].(func(*base.Module, int32) int32)(m, l2)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1127), int32(_a_F_SlruReportIOError_3))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1127), int32(_a_F_SlruReportIOError_3))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v11 + int32(160)
							F_errmsg(m, int32(_a_F_SlruReportIOError_5), v11+int32(32))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return
							} else {
								if l2 != 0 {
									v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v131 = m.T0[v130].(func(*base.Module, int32) int32)(m, l2)
									mBase = m.M
									v132 = m.ExcPending
									if v132 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1132), int32(_a_F_SlruReportIOError_3))
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1132), int32(_a_F_SlruReportIOError_3))
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
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
				case 2:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return
					} else {
						if v41 != 0 {
							F_errcode_for_file_access(m)
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v11 + int32(160)
								F_errmsg(m, int32(_a_F_SlruReportIOError_6), v11+int32(80))
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return
								} else {
									if l2 != 0 {
										v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v154 = m.T0[v153].(func(*base.Module, int32) int32)(m, l2)
										mBase = m.M
										v155 = m.ExcPending
										if v155 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1140), int32(_a_F_SlruReportIOError_3))
											mBase = m.M
											v160 = m.ExcPending
											if v160 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1140), int32(_a_F_SlruReportIOError_3))
										mBase = m.M
										v160 = m.ExcPending
										if v160 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v11 + int32(160)
							F_errmsg(m, int32(_a_F_SlruReportIOError_7), v11-int32(-64))
							mBase = m.M
							v169 = m.ExcPending
							if v169 != 0 {
								return
							} else {
								if l2 != 0 {
									v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v171 = m.T0[v170].(func(*base.Module, int32) int32)(m, l2)
									mBase = m.M
									v172 = m.ExcPending
									if v172 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1145), int32(_a_F_SlruReportIOError_3))
										mBase = m.M
										v177 = m.ExcPending
										if v177 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1145), int32(_a_F_SlruReportIOError_3))
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
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
				case 3:
					base.Wasm_trap_unreachable()
					for {
					}
				case 4:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v210 = m.ExcPending
					if v210 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v212 = m.ExcPending
						if v212 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v11 + int32(160)
							F_errmsg(m, int32(_a_F_SlruReportIOError_8), v11+int32(112))
							mBase = m.M
							v220 = m.ExcPending
							if v220 != 0 {
								return
							} else {
								if l2 != 0 {
									v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v222 = m.T0[v221].(func(*base.Module, int32) int32)(m, l2)
									mBase = m.M
									v223 = m.ExcPending
									if v223 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1159), int32(_a_F_SlruReportIOError_3))
										mBase = m.M
										v228 = m.ExcPending
										if v228 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1159), int32(_a_F_SlruReportIOError_3))
									mBase = m.M
									v228 = m.ExcPending
									if v228 != 0 {
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
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v11 + int32(160)
							F_errmsg(m, int32(_a_F_SlruReportIOError_9), v11)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								if l2 != 0 {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v68 = m.T0[v67].(func(*base.Module, int32) int32)(m, l2)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1112), int32(_a_F_SlruReportIOError_3))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1112), int32(_a_F_SlruReportIOError_3))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
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
			} else {
				v181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SlruReportIOError[3])))
				if v181 != 0 {
					v182 = int32(21)
				} else {
					v182 = int32(24)
				}
				v184 = F_errstart(m, v182, int32(0))
				mBase = m.M
				v185 = m.ExcPending
				if v185 != 0 {
					return
				} else {
					if v184 != 0 {
						F_errcode_for_file_access(m)
						mBase = m.M
						v187 = m.ExcPending
						if v187 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v11 + int32(160)
							F_errmsg(m, int32(_a_F_SlruReportIOError_10), v11+int32(96))
							mBase = m.M
							v195 = m.ExcPending
							if v195 != 0 {
								return
							} else {
								if l2 != 0 {
									v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v197 = m.T0[v196].(func(*base.Module, int32) int32)(m, l2)
									mBase = m.M
									v198 = m.ExcPending
									if v198 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1152), int32(_a_F_SlruReportIOError_3))
										mBase = m.M
										v203 = m.ExcPending
										if v203 != 0 {
											return
										} else {
											m.G0 = v11 + int32(1184)
											return
										}
									}
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1152), int32(_a_F_SlruReportIOError_3))
									mBase = m.M
									v203 = m.ExcPending
									if v203 != 0 {
										return
									} else {
										m.G0 = v11 + int32(1184)
										return
									}
								}
							}
						}
					} else {
						m.G0 = v11 + int32(1184)
						return
					}
				}
			}
		}
	} else {
		*(*uint32)(unsafe.Add(mBase, uint32(v11)+148)) = uint32(v14)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+144)) = v15
		v37 = F_pg_snprintf(m, v11+int32(160), int32(1024), int32(_a_F_SlruReportIOError_11), v11+int32(144))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, _c_F_SlruReportIOError[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_SlruReportIOError[1])) = v41
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_SlruReportIOError[2]))
			if v44 != int32(4) {
				v52 = base.I32_wrap_i64(l1-v14<<(uint(int64(5))%64)) << (uint(int32(13)) % 32)
				switch v44 - int32(1) {
				case 0:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v11 + int32(160)
							F_errmsg(m, int32(_a_F_SlruReportIOError_1), v11+int32(16))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								if l2 != 0 {
									v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v91 = m.T0[v90].(func(*base.Module, int32) int32)(m, l2)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1119), int32(_a_F_SlruReportIOError_3))
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
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1119), int32(_a_F_SlruReportIOError_3))
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
						}
					}
				case 1:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return
					} else {
						if v41 != 0 {
							F_errcode_for_file_access(m)
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v11 + int32(160)
								F_errmsg(m, int32(_a_F_SlruReportIOError_4), v11+int32(48))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return
								} else {
									if l2 != 0 {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v114 = m.T0[v113].(func(*base.Module, int32) int32)(m, l2)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1127), int32(_a_F_SlruReportIOError_3))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1127), int32(_a_F_SlruReportIOError_3))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v11 + int32(160)
							F_errmsg(m, int32(_a_F_SlruReportIOError_5), v11+int32(32))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return
							} else {
								if l2 != 0 {
									v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v131 = m.T0[v130].(func(*base.Module, int32) int32)(m, l2)
									mBase = m.M
									v132 = m.ExcPending
									if v132 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1132), int32(_a_F_SlruReportIOError_3))
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1132), int32(_a_F_SlruReportIOError_3))
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
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
				case 2:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return
					} else {
						if v41 != 0 {
							F_errcode_for_file_access(m)
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v11 + int32(160)
								F_errmsg(m, int32(_a_F_SlruReportIOError_6), v11+int32(80))
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return
								} else {
									if l2 != 0 {
										v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v154 = m.T0[v153].(func(*base.Module, int32) int32)(m, l2)
										mBase = m.M
										v155 = m.ExcPending
										if v155 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1140), int32(_a_F_SlruReportIOError_3))
											mBase = m.M
											v160 = m.ExcPending
											if v160 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1140), int32(_a_F_SlruReportIOError_3))
										mBase = m.M
										v160 = m.ExcPending
										if v160 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v11 + int32(160)
							F_errmsg(m, int32(_a_F_SlruReportIOError_7), v11-int32(-64))
							mBase = m.M
							v169 = m.ExcPending
							if v169 != 0 {
								return
							} else {
								if l2 != 0 {
									v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v171 = m.T0[v170].(func(*base.Module, int32) int32)(m, l2)
									mBase = m.M
									v172 = m.ExcPending
									if v172 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1145), int32(_a_F_SlruReportIOError_3))
										mBase = m.M
										v177 = m.ExcPending
										if v177 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1145), int32(_a_F_SlruReportIOError_3))
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
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
				case 3:
					base.Wasm_trap_unreachable()
					for {
					}
				case 4:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v210 = m.ExcPending
					if v210 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v212 = m.ExcPending
						if v212 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v11 + int32(160)
							F_errmsg(m, int32(_a_F_SlruReportIOError_8), v11+int32(112))
							mBase = m.M
							v220 = m.ExcPending
							if v220 != 0 {
								return
							} else {
								if l2 != 0 {
									v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v222 = m.T0[v221].(func(*base.Module, int32) int32)(m, l2)
									mBase = m.M
									v223 = m.ExcPending
									if v223 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1159), int32(_a_F_SlruReportIOError_3))
										mBase = m.M
										v228 = m.ExcPending
										if v228 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1159), int32(_a_F_SlruReportIOError_3))
									mBase = m.M
									v228 = m.ExcPending
									if v228 != 0 {
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
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v11 + int32(160)
							F_errmsg(m, int32(_a_F_SlruReportIOError_9), v11)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								if l2 != 0 {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v68 = m.T0[v67].(func(*base.Module, int32) int32)(m, l2)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1112), int32(_a_F_SlruReportIOError_3))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1112), int32(_a_F_SlruReportIOError_3))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
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
			} else {
				v181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SlruReportIOError[3])))
				if v181 != 0 {
					v182 = int32(21)
				} else {
					v182 = int32(24)
				}
				v184 = F_errstart(m, v182, int32(0))
				mBase = m.M
				v185 = m.ExcPending
				if v185 != 0 {
					return
				} else {
					if v184 != 0 {
						F_errcode_for_file_access(m)
						mBase = m.M
						v187 = m.ExcPending
						if v187 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v11 + int32(160)
							F_errmsg(m, int32(_a_F_SlruReportIOError_10), v11+int32(96))
							mBase = m.M
							v195 = m.ExcPending
							if v195 != 0 {
								return
							} else {
								if l2 != 0 {
									v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v197 = m.T0[v196].(func(*base.Module, int32) int32)(m, l2)
									mBase = m.M
									v198 = m.ExcPending
									if v198 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1152), int32(_a_F_SlruReportIOError_3))
										mBase = m.M
										v203 = m.ExcPending
										if v203 != 0 {
											return
										} else {
											m.G0 = v11 + int32(1184)
											return
										}
									}
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_2), int32(1152), int32(_a_F_SlruReportIOError_3))
									mBase = m.M
									v203 = m.ExcPending
									if v203 != 0 {
										return
									} else {
										m.G0 = v11 + int32(1184)
										return
									}
								}
							}
						}
					} else {
						m.G0 = v11 + int32(1184)
						return
					}
				}
			}
		}
	}
}
func F_SlruSelectLRUPage(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int64
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v125 int64
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v153 int64
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
	var v171 int32
	_ = v171
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	goto L1
L1:
	;
	v34 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)))
	v35 = base.I64_rem_s(l1, v34)
	v36 = base.I32_wrap_i64(v35)
	v38 = v36 << (uint(int32(4)) % 32)
	v40 = v38 + int32(16)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v44 = v38
	goto L4
L2:
	;
	return v171
L3:
	;
	goto L2
L4:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v41+v44<<(uint(int32(2))%32))))
	if v61 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v74 = v71 + v36<<(uint(int32(2))%32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v75 + int32(1)
	v79 = int64(0)
	v80 = int32(-1)
	v81 = int32(0)
	v87 = v81
	v90 = v38
	v91 = v81
	v92 = v80
	v93 = v80
	v99 = v79
	v100 = v79
	goto L11
L6:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v62+v44<<(uint(int32(3))%32))))
	if v66 == l1 {
		v171 = v44
		goto L3
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v69 = v44 + int32(1)
	if v69 < v40 {
		v44 = v69
		goto L4
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	goto L5
L11:
	;
	v102 = v90 << (uint(int32(2)) % 32)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v102+v103)))
	if v105 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v150 < int32(0) {
		goto L38
	} else {
		goto L39
	}
L13:
	;
	v171 = v90
	goto L3
L14:
	;
	goto L15
L15:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v109 = v108 + v102
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v111 = v75 - v110
	if v111 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = v75
	v116 = int32(0)
	goto L18
L17:
	;
	v116 = v111
	goto L18
L18:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v117+v90<<(uint(int32(3))%32))))
	v122 = int64(0)
	v125 = base.AtomicRmwCmpxchg64(m, v17, int32(48), v122, v122)
	if v121 == v125 {
		v148 = v87
		v149 = v91
		v150 = v92
		v151 = v93
		v152 = v99
		v153 = v100
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v155 = v90 + int32(1)
	if v155 < v40 {
		v87 = v148
		v90 = v155
		v91 = v149
		v92 = v150
		v93 = v151
		v99 = v152
		v100 = v153
		goto L11
	} else {
		goto L37
	}
L20:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v127+v102)))
	if v129 == int32(2) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v116 <= v92 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	if v116 <= v93 {
		goto L31
	} else {
		goto L32
	}
L24:
	;
	if v116 != v92 {
		v148 = v87
		v149 = v91
		v150 = v92
		v151 = v93
		v152 = v99
		v153 = v100
		goto L19
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v148 = v90
	v149 = v91
	v150 = v116
	v151 = v93
	v152 = v121
	v153 = v100
	goto L19
L27:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v135 = m.T0[v134].(func(*base.Module, int64, int64) int32)(m, v121, v99)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	return int32(0)
L29:
	;
	if v135 == int32(0) {
		v148 = v87
		v149 = v91
		v150 = v92
		v151 = v93
		v152 = v99
		v153 = v100
		goto L19
	} else {
		goto L30
	}
L30:
	;
	goto L26
L31:
	;
	if v116 != v93 {
		v148 = v87
		v149 = v91
		v150 = v92
		v151 = v93
		v152 = v99
		v153 = v100
		goto L19
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v148 = v87
	v149 = v90
	v150 = v92
	v151 = v116
	v152 = v99
	v153 = v121
	goto L19
L34:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v144 = m.T0[v143].(func(*base.Module, int64, int64) int32)(m, v121, v100)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L28
	} else {
		goto L35
	}
L35:
	;
	if v144 == int32(0) {
		v148 = v87
		v149 = v91
		v150 = v92
		v151 = v93
		v152 = v99
		v153 = v100
		goto L19
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	goto L12
L38:
	;
	F_SimpleLruWaitIO(m, l0, v149)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L28
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161+v148))))
	if v163 != int32(1) {
		v171 = v148
		goto L3
	} else {
		goto L42
	}
L41:
	;
	goto L1
L42:
	;
	F_SlruInternalWritePage(m, l0, v148, int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L28
	} else {
		goto L43
	}
L43:
	;
	goto L1
}
