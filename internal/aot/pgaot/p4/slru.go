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
		v27 = F_pg_snprintf(m, v10+int32(288), int32(1024), int32(509867), v10+int32(256))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, _consts[159]))
			*(*int32)(unsafe.Add(mBase, _consts[140])) = v41
			v44 = *(*int32)(unsafe.Add(mBase, _consts[160]))
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
							F_errmsg(m, int32(45698), v10+int32(48))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v10 + int32(288)
								F_errdetail(m, int32(621722), v10+int32(32))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return
								} else {
									F_errfinish(m, int32(492478), int32(1070), int32(212588))
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
								F_errmsg(m, int32(45698), v10+int32(112))
								mBase = m.M
								v216 = m.ExcPending
								if v216 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = v52
									*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v10 + int32(288)
									F_errdetail(m, int32(621674), v10+int32(96))
									mBase = m.M
									v225 = m.ExcPending
									if v225 != 0 {
										return
									} else {
										F_errfinish(m, int32(492478), int32(1078), int32(212588))
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
							F_errmsg(m, int32(45698), v10+int32(80))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v10 + int32(288)
								F_errdetail(m, int32(596659), v10-int32(-64))
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return
								} else {
									F_errfinish(m, int32(492478), int32(1082), int32(212588))
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
								F_errmsg(m, int32(45698), v10+int32(176))
								mBase = m.M
								v238 = m.ExcPending
								if v238 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+164)) = v52
									*(*int32)(unsafe.Add(mBase, uint32(v10)+160)) = v10 + int32(288)
									F_errdetail(m, int32(621627), v10+int32(160))
									mBase = m.M
									v247 = m.ExcPending
									if v247 != 0 {
										return
									} else {
										F_errfinish(m, int32(492478), int32(1090), int32(212588))
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
							F_errmsg(m, int32(45698), v10+int32(144))
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v10)+128)) = v10 + int32(288)
								F_errdetail(m, int32(596595), v10+int32(128))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return
								} else {
									F_errfinish(m, int32(492478), int32(1095), int32(212588))
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
							F_errmsg(m, int32(45698), v10+int32(240))
							mBase = m.M
							v195 = m.ExcPending
							if v195 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+224)) = v10 + int32(288)
								F_errdetail(m, int32(621838), v10+int32(224))
								mBase = m.M
								v203 = m.ExcPending
								if v203 != 0 {
									return
								} else {
									F_errfinish(m, int32(492478), int32(1109), int32(212588))
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
							F_errmsg(m, int32(45698), v10+int32(16))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v10 + int32(288)
								F_errdetail(m, int32(621808), v10)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									F_errfinish(m, int32(492478), int32(1063), int32(212588))
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
				v155 = int32(*(*uint8)(unsafe.Add(mBase, _consts[162])))
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
							F_errmsg(m, int32(45698), v10+int32(208))
							mBase = m.M
							v167 = m.ExcPending
							if v167 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+192)) = v10 + int32(288)
								F_errdetail(m, int32(621869), v10+int32(192))
								mBase = m.M
								v175 = m.ExcPending
								if v175 != 0 {
									return
								} else {
									F_errfinish(m, int32(492478), int32(1102), int32(212588))
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
		v37 = F_pg_snprintf(m, v10+int32(288), int32(1024), int32(510254), v10+int32(272))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, _consts[159]))
			*(*int32)(unsafe.Add(mBase, _consts[140])) = v41
			v44 = *(*int32)(unsafe.Add(mBase, _consts[160]))
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
							F_errmsg(m, int32(45698), v10+int32(48))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v10 + int32(288)
								F_errdetail(m, int32(621722), v10+int32(32))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return
								} else {
									F_errfinish(m, int32(492478), int32(1070), int32(212588))
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
								F_errmsg(m, int32(45698), v10+int32(112))
								mBase = m.M
								v216 = m.ExcPending
								if v216 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = v52
									*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v10 + int32(288)
									F_errdetail(m, int32(621674), v10+int32(96))
									mBase = m.M
									v225 = m.ExcPending
									if v225 != 0 {
										return
									} else {
										F_errfinish(m, int32(492478), int32(1078), int32(212588))
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
							F_errmsg(m, int32(45698), v10+int32(80))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v10 + int32(288)
								F_errdetail(m, int32(596659), v10-int32(-64))
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return
								} else {
									F_errfinish(m, int32(492478), int32(1082), int32(212588))
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
								F_errmsg(m, int32(45698), v10+int32(176))
								mBase = m.M
								v238 = m.ExcPending
								if v238 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+164)) = v52
									*(*int32)(unsafe.Add(mBase, uint32(v10)+160)) = v10 + int32(288)
									F_errdetail(m, int32(621627), v10+int32(160))
									mBase = m.M
									v247 = m.ExcPending
									if v247 != 0 {
										return
									} else {
										F_errfinish(m, int32(492478), int32(1090), int32(212588))
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
							F_errmsg(m, int32(45698), v10+int32(144))
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v10)+128)) = v10 + int32(288)
								F_errdetail(m, int32(596595), v10+int32(128))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return
								} else {
									F_errfinish(m, int32(492478), int32(1095), int32(212588))
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
							F_errmsg(m, int32(45698), v10+int32(240))
							mBase = m.M
							v195 = m.ExcPending
							if v195 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+224)) = v10 + int32(288)
								F_errdetail(m, int32(621838), v10+int32(224))
								mBase = m.M
								v203 = m.ExcPending
								if v203 != 0 {
									return
								} else {
									F_errfinish(m, int32(492478), int32(1109), int32(212588))
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
							F_errmsg(m, int32(45698), v10+int32(16))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v10 + int32(288)
								F_errdetail(m, int32(621808), v10)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									F_errfinish(m, int32(492478), int32(1063), int32(212588))
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
				v155 = int32(*(*uint8)(unsafe.Add(mBase, _consts[162])))
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
							F_errmsg(m, int32(45698), v10+int32(208))
							mBase = m.M
							v167 = m.ExcPending
							if v167 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+192)) = v10 + int32(288)
								F_errdetail(m, int32(621869), v10+int32(192))
								mBase = m.M
								v175 = m.ExcPending
								if v175 != 0 {
									return
								} else {
									F_errfinish(m, int32(492478), int32(1102), int32(212588))
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
	var v18 int32
	_ = v18
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int64
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int64
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
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
	var v121 int32
	_ = v121
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v153 int32
	_ = v153
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
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L1
L1:
	;
	v36 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v37 = base.I64_rem_s(l1, v36)
	v38 = base.I32_wrap_i64(v37)
	v40 = v38 << (uint(int32(4)) % 32)
	v42 = v40 + int32(16)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v46 = v40
	goto L4
L2:
	;
	return v171
L3:
	;
	goto L2
L4:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v43+v46<<(uint(int32(2))%32))))
	if v64 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v77 = v74 + v38<<(uint(int32(2))%32)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v78 + int32(1)
	v82 = int64(0)
	v83 = int32(-1)
	v84 = int32(0)
	v90 = v84
	v93 = v40
	v95 = v83
	v96 = v83
	v99 = v84
	v102 = v82
	v103 = v82
	goto L11
L6:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v65+v46<<(uint(int32(3))%32))))
	if v69 == l1 {
		v171 = v46
		goto L3
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v72 = v46 + int32(1)
	if v72 < v42 {
		v46 = v72
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
	v106 = v93 << (uint(int32(2)) % 32)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106+v107)))
	if v109 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v147 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L13:
	;
	v171 = v93
	goto L3
L14:
	;
	goto L15
L15:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v113 = v112 + v106
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v115 = v78 - v114
	if v115 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v78
	v120 = int32(0)
	goto L18
L17:
	;
	v120 = v115
	goto L18
L18:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v121+v93<<(uint(int32(3))%32))))
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v126
	if v125 == v126 {
		v146 = v90
		v147 = v95
		v148 = v96
		v149 = v99
		v150 = v102
		v151 = v103
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v153 = v93 + int32(1)
	if v153 < v42 {
		v90 = v146
		v93 = v153
		v95 = v147
		v96 = v148
		v99 = v149
		v102 = v150
		v103 = v151
		goto L11
	} else {
		goto L34
	}
L20:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129+v106)))
	if v131 == int32(2) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v120 <= v95 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	if v120 <= v96 {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	if v120 != v95 {
		v146 = v90
		v147 = v95
		v148 = v96
		v149 = v99
		v150 = v102
		v151 = v103
		goto L19
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v146 = v93
	v147 = v120
	v148 = v96
	v149 = v99
	v150 = v125
	v151 = v103
	goto L19
L27:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v137 = m.T0[v136].(func(*base.Module, int64, int64) int32)(m, v125, v102)
	mBase = m.M
	if v137 == int32(0) {
		v146 = v90
		v147 = v95
		v148 = v96
		v149 = v99
		v150 = v102
		v151 = v103
		goto L19
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	if v120 != v96 {
		v146 = v90
		v147 = v95
		v148 = v96
		v149 = v99
		v150 = v102
		v151 = v103
		goto L19
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v146 = v90
	v147 = v95
	v148 = v120
	v149 = v93
	v150 = v102
	v151 = v125
	goto L19
L32:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v143 = m.T0[v142].(func(*base.Module, int64, int64) int32)(m, v125, v103)
	mBase = m.M
	if v143 == int32(0) {
		v146 = v90
		v147 = v95
		v148 = v96
		v149 = v99
		v150 = v102
		v151 = v103
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
	F_SimpleLruWaitIO(m, l0, v149)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161+v146))))
	if v163 != int32(1) {
		v171 = v146
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
	F_SlruInternalWritePage(m, l0, v146, int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L1
}
