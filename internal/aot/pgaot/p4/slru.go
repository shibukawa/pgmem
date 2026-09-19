package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SlruReportIOError(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int64
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
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	v8 = m.G0
	v10 = v8 - int32(1312)
	m.G0 = v10
	v13 = l0 + int32(16)
	v15 = base.I64_div_s(l1, int64(32))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v16 == int32(1) {
		*(*int64)(unsafe.Add(mBase, uint32(v10)+264)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v10)+256)) = v13
		v27 = F_pg_snprintf(m, v10+int32(288), int32(1024), int32(_a_F_SlruReportIOError_0), v10+int32(256))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, _c_F_SlruReportIOError[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_SlruReportIOError[1])) = v41
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_SlruReportIOError[2]))
			if v44 != int32(4) {
				v52 = base.I32_wrap_i64(l1-v15<<(uint(int64(5))%64)) << (uint(int32(13)) % 32)
				switch v44 - int32(1) {
				case 0:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l2
							F_errmsg(m, int32(_a_F_SlruReportIOError_1), v10+int32(48))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v10 + int32(288)
								F_errdetail(m, int32(_a_F_SlruReportIOError_2), v10+int32(32))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_3), int32(1070), int32(_a_F_SlruReportIOError_4))
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
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
					v107 = m.ExcPending
					if v107 != 0 {
						return
					} else {
						if v41 != 0 {
							F_errcode_for_file_access(m)
							mBase = m.M
							v210 = m.ExcPending
							if v210 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = l2
								F_errmsg(m, int32(_a_F_SlruReportIOError_1), v10+int32(112))
								mBase = m.M
								v216 = m.ExcPending
								if v216 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = v52
									*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v10 + int32(288)
									F_errdetail(m, int32(_a_F_SlruReportIOError_5), v10+int32(96))
									mBase = m.M
									v225 = m.ExcPending
									if v225 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_3), int32(1078), int32(_a_F_SlruReportIOError_4))
										mBase = m.M
										v230 = m.ExcPending
										if v230 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = l2
							F_errmsg(m, int32(_a_F_SlruReportIOError_1), v10+int32(80))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v10 + int32(288)
								F_errdetail(m, int32(_a_F_SlruReportIOError_6), v10-int32(-64))
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_3), int32(1082), int32(_a_F_SlruReportIOError_4))
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
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
					v131 = m.ExcPending
					if v131 != 0 {
						return
					} else {
						if v41 != 0 {
							F_errcode_for_file_access(m)
							mBase = m.M
							v232 = m.ExcPending
							if v232 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+176)) = l2
								F_errmsg(m, int32(_a_F_SlruReportIOError_1), v10+int32(176))
								mBase = m.M
								v238 = m.ExcPending
								if v238 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+164)) = v52
									*(*int32)(unsafe.Add(mBase, uint32(v10)+160)) = v10 + int32(288)
									F_errdetail(m, int32(_a_F_SlruReportIOError_7), v10+int32(160))
									mBase = m.M
									v247 = m.ExcPending
									if v247 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_3), int32(1090), int32(_a_F_SlruReportIOError_4))
										mBase = m.M
										v252 = m.ExcPending
										if v252 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(v10)+144)) = l2
							F_errmsg(m, int32(_a_F_SlruReportIOError_1), v10+int32(144))
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v10)+128)) = v10 + int32(288)
								F_errdetail(m, int32(_a_F_SlruReportIOError_8), v10+int32(128))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_3), int32(1095), int32(_a_F_SlruReportIOError_4))
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
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
					v187 = m.ExcPending
					if v187 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v189 = m.ExcPending
						if v189 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+240)) = l2
							F_errmsg(m, int32(_a_F_SlruReportIOError_1), v10+int32(240))
							mBase = m.M
							v195 = m.ExcPending
							if v195 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+224)) = v10 + int32(288)
								F_errdetail(m, int32(_a_F_SlruReportIOError_9), v10+int32(224))
								mBase = m.M
								v203 = m.ExcPending
								if v203 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_3), int32(1109), int32(_a_F_SlruReportIOError_4))
									mBase = m.M
									v208 = m.ExcPending
									if v208 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l2
							F_errmsg(m, int32(_a_F_SlruReportIOError_1), v10+int32(16))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v10 + int32(288)
								F_errdetail(m, int32(_a_F_SlruReportIOError_10), v10)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_3), int32(1063), int32(_a_F_SlruReportIOError_4))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
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
				v155 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SlruReportIOError[3])))
				if v155 != 0 {
					v156 = int32(21)
				} else {
					v156 = int32(23)
				}
				v158 = F_errstart(m, v156, int32(0))
				mBase = m.M
				v159 = m.ExcPending
				if v159 != 0 {
					return
				} else {
					if v158 != 0 {
						F_errcode_for_file_access(m)
						mBase = m.M
						v161 = m.ExcPending
						if v161 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+208)) = l2
							F_errmsg(m, int32(_a_F_SlruReportIOError_1), v10+int32(208))
							mBase = m.M
							v167 = m.ExcPending
							if v167 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+192)) = v10 + int32(288)
								F_errdetail(m, int32(_a_F_SlruReportIOError_11), v10+int32(192))
								mBase = m.M
								v175 = m.ExcPending
								if v175 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_3), int32(1102), int32(_a_F_SlruReportIOError_4))
									mBase = m.M
									v180 = m.ExcPending
									if v180 != 0 {
										return
									} else {
										m.G0 = v10 + int32(1312)
										return
									}
								}
							}
						}
					} else {
						m.G0 = v10 + int32(1312)
						return
					}
				}
			}
		}
	} else {
		*(*uint32)(unsafe.Add(mBase, uint32(v10)+276)) = uint32(v15)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+272)) = v13
		v37 = F_pg_snprintf(m, v10+int32(288), int32(1024), int32(_a_F_SlruReportIOError_12), v10+int32(272))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, _c_F_SlruReportIOError[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_SlruReportIOError[1])) = v41
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_SlruReportIOError[2]))
			if v44 != int32(4) {
				v52 = base.I32_wrap_i64(l1-v15<<(uint(int64(5))%64)) << (uint(int32(13)) % 32)
				switch v44 - int32(1) {
				case 0:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l2
							F_errmsg(m, int32(_a_F_SlruReportIOError_1), v10+int32(48))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v10 + int32(288)
								F_errdetail(m, int32(_a_F_SlruReportIOError_2), v10+int32(32))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_3), int32(1070), int32(_a_F_SlruReportIOError_4))
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
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
					v107 = m.ExcPending
					if v107 != 0 {
						return
					} else {
						if v41 != 0 {
							F_errcode_for_file_access(m)
							mBase = m.M
							v210 = m.ExcPending
							if v210 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = l2
								F_errmsg(m, int32(_a_F_SlruReportIOError_1), v10+int32(112))
								mBase = m.M
								v216 = m.ExcPending
								if v216 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = v52
									*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v10 + int32(288)
									F_errdetail(m, int32(_a_F_SlruReportIOError_5), v10+int32(96))
									mBase = m.M
									v225 = m.ExcPending
									if v225 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_3), int32(1078), int32(_a_F_SlruReportIOError_4))
										mBase = m.M
										v230 = m.ExcPending
										if v230 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = l2
							F_errmsg(m, int32(_a_F_SlruReportIOError_1), v10+int32(80))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v10 + int32(288)
								F_errdetail(m, int32(_a_F_SlruReportIOError_6), v10-int32(-64))
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_3), int32(1082), int32(_a_F_SlruReportIOError_4))
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
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
					v131 = m.ExcPending
					if v131 != 0 {
						return
					} else {
						if v41 != 0 {
							F_errcode_for_file_access(m)
							mBase = m.M
							v232 = m.ExcPending
							if v232 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+176)) = l2
								F_errmsg(m, int32(_a_F_SlruReportIOError_1), v10+int32(176))
								mBase = m.M
								v238 = m.ExcPending
								if v238 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+164)) = v52
									*(*int32)(unsafe.Add(mBase, uint32(v10)+160)) = v10 + int32(288)
									F_errdetail(m, int32(_a_F_SlruReportIOError_7), v10+int32(160))
									mBase = m.M
									v247 = m.ExcPending
									if v247 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SlruReportIOError_3), int32(1090), int32(_a_F_SlruReportIOError_4))
										mBase = m.M
										v252 = m.ExcPending
										if v252 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(v10)+144)) = l2
							F_errmsg(m, int32(_a_F_SlruReportIOError_1), v10+int32(144))
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v10)+128)) = v10 + int32(288)
								F_errdetail(m, int32(_a_F_SlruReportIOError_8), v10+int32(128))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_3), int32(1095), int32(_a_F_SlruReportIOError_4))
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
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
					v187 = m.ExcPending
					if v187 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v189 = m.ExcPending
						if v189 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+240)) = l2
							F_errmsg(m, int32(_a_F_SlruReportIOError_1), v10+int32(240))
							mBase = m.M
							v195 = m.ExcPending
							if v195 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+224)) = v10 + int32(288)
								F_errdetail(m, int32(_a_F_SlruReportIOError_9), v10+int32(224))
								mBase = m.M
								v203 = m.ExcPending
								if v203 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_3), int32(1109), int32(_a_F_SlruReportIOError_4))
									mBase = m.M
									v208 = m.ExcPending
									if v208 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l2
							F_errmsg(m, int32(_a_F_SlruReportIOError_1), v10+int32(16))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v10 + int32(288)
								F_errdetail(m, int32(_a_F_SlruReportIOError_10), v10)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_3), int32(1063), int32(_a_F_SlruReportIOError_4))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
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
				v155 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SlruReportIOError[3])))
				if v155 != 0 {
					v156 = int32(21)
				} else {
					v156 = int32(23)
				}
				v158 = F_errstart(m, v156, int32(0))
				mBase = m.M
				v159 = m.ExcPending
				if v159 != 0 {
					return
				} else {
					if v158 != 0 {
						F_errcode_for_file_access(m)
						mBase = m.M
						v161 = m.ExcPending
						if v161 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+208)) = l2
							F_errmsg(m, int32(_a_F_SlruReportIOError_1), v10+int32(208))
							mBase = m.M
							v167 = m.ExcPending
							if v167 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+192)) = v10 + int32(288)
								F_errdetail(m, int32(_a_F_SlruReportIOError_11), v10+int32(192))
								mBase = m.M
								v175 = m.ExcPending
								if v175 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_SlruReportIOError_3), int32(1102), int32(_a_F_SlruReportIOError_4))
									mBase = m.M
									v180 = m.ExcPending
									if v180 != 0 {
										return
									} else {
										m.G0 = v10 + int32(1312)
										return
									}
								}
							}
						}
					} else {
						m.G0 = v10 + int32(1312)
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
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L1
L1:
	;
	v34 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v35 = base.I64_rem_s(l1, v34)
	v36 = base.I32_wrap_i64(v35)
	v38 = v36 << (uint(int32(4)) % 32)
	v40 = v38 + int32(16)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v44 = v38
	goto L4
L2:
	;
	return v169
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
		v169 = v44
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
	if v146 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L13:
	;
	v169 = v90
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
		v144 = v87
		v145 = v91
		v146 = v92
		v147 = v93
		v148 = v99
		v149 = v100
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v151 = v90 + int32(1)
	if v151 < v40 {
		v87 = v144
		v90 = v151
		v91 = v145
		v92 = v146
		v93 = v147
		v99 = v148
		v100 = v149
		goto L11
	} else {
		goto L34
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
		goto L29
	} else {
		goto L30
	}
L24:
	;
	if v116 != v92 {
		v144 = v87
		v145 = v91
		v146 = v92
		v147 = v93
		v148 = v99
		v149 = v100
		goto L19
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v144 = v90
	v145 = v91
	v146 = v116
	v147 = v93
	v148 = v121
	v149 = v100
	goto L19
L27:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v135 = m.T0[v134].(func(*base.Module, int64, int64) int32)(m, v121, v99)
	mBase = m.M
	if v135 == int32(0) {
		v144 = v87
		v145 = v91
		v146 = v92
		v147 = v93
		v148 = v99
		v149 = v100
		goto L19
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	if v116 != v93 {
		v144 = v87
		v145 = v91
		v146 = v92
		v147 = v93
		v148 = v99
		v149 = v100
		goto L19
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v144 = v87
	v145 = v90
	v146 = v92
	v147 = v116
	v148 = v99
	v149 = v121
	goto L19
L32:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v141 = m.T0[v140].(func(*base.Module, int64, int64) int32)(m, v121, v100)
	mBase = m.M
	if v141 == int32(0) {
		v144 = v87
		v145 = v91
		v146 = v92
		v147 = v93
		v148 = v99
		v149 = v100
		goto L19
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L12
L35:
	;
	F_SimpleLruWaitIO(m, l0, v145)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159+v144))))
	if v161 != int32(1) {
		v169 = v144
		goto L3
	} else {
		goto L40
	}
L38:
	;
	return int32(0)
L39:
	;
	goto L1
L40:
	;
	F_SlruInternalWritePage(m, l0, v144, int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L1
}
