package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_map_sql_type_to_xml_name(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	v4 = m.G0
	v6 = v4 - int32(160)
	m.G0 = v6
	F_initStringInfo(m, v6+int32(128))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if l0 <= int32(1041) {
			switch l0 - int32(16) {
			case 0:
				F_appendStringInfoString(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_0))
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return int32(0)
				} else {
					v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
					m.G0 = v6 + int32(160)
					return v229
				}
			case 1, 2, 3, 6:
				v180 = F_SearchSysCache1(m, int32(82), l0)
				mBase = m.M
				v181 = m.ExcPending
				if v181 != 0 {
					return int32(0)
				} else {
					if v180 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v213 = m.ExcPending
						if v213 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
							F_errmsg_internal(m, int32(_a_F_map_sql_type_to_xml_name_1), v6)
							mBase = m.M
							v217 = m.ExcPending
							if v217 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_map_sql_type_to_xml_name_2), int32(3833), int32(_a_F_map_sql_type_to_xml_name_3))
								mBase = m.M
								v222 = m.ExcPending
								if v222 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v184 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
						v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+22)))
						v187 = *(*int32)(unsafe.Add(mBase, _c_F_map_sql_type_to_xml_name[0]))
						v188 = F_get_database_name(m, v187)
						mBase = m.M
						v189 = m.ExcPending
						if v189 != 0 {
							return int32(0)
						} else {
							v191 = *(*int32)(unsafe.Add(mBase, uint32(v184+v185)+68))
							v192 = F_get_namespace_name(m, v191)
							mBase = m.M
							v193 = m.ExcPending
							if v193 != 0 {
								return int32(0)
							} else {
								F_initStringInfo(m, v6+int32(144))
								mBase = m.M
								v197 = m.ExcPending
								if v197 != 0 {
									return int32(0)
								} else {
									v198 = F_map_sql_identifier_to_xml_name(m)
									mBase = m.M
									v199 = m.ExcPending
									if v199 != 0 {
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
			case 4:
				F_appendStringInfoString(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_4))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
					m.G0 = v6 + int32(160)
					return v229
				}
			case 5:
				F_appendStringInfoString(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_5))
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int32(0)
				} else {
					v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
					m.G0 = v6 + int32(160)
					return v229
				}
			case 7:
				F_appendStringInfoString(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_6))
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return int32(0)
				} else {
					v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
					m.G0 = v6 + int32(160)
					return v229
				}
			default:
				switch l0 - int32(700) {
				case 0:
					F_appendStringInfoString(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_7))
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int32(0)
					} else {
						v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
						m.G0 = v6 + int32(160)
						return v229
					}
				case 1:
					F_appendStringInfoString(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_8))
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return int32(0)
					} else {
						v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
						m.G0 = v6 + int32(160)
						return v229
					}
				default:
					if l0 == int32(142) {
						F_appendStringInfoString(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_9))
						mBase = m.M
						v227 = m.ExcPending
						if v227 != 0 {
							return int32(0)
						} else {
							v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
							m.G0 = v6 + int32(160)
							return v229
						}
					} else {
						v180 = F_SearchSysCache1(m, int32(82), l0)
						mBase = m.M
						v181 = m.ExcPending
						if v181 != 0 {
							return int32(0)
						} else {
							if v180 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v213 = m.ExcPending
								if v213 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
									F_errmsg_internal(m, int32(_a_F_map_sql_type_to_xml_name_1), v6)
									mBase = m.M
									v217 = m.ExcPending
									if v217 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_map_sql_type_to_xml_name_2), int32(3833), int32(_a_F_map_sql_type_to_xml_name_3))
										mBase = m.M
										v222 = m.ExcPending
										if v222 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v184 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
								v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+22)))
								v187 = *(*int32)(unsafe.Add(mBase, _c_F_map_sql_type_to_xml_name[0]))
								v188 = F_get_database_name(m, v187)
								mBase = m.M
								v189 = m.ExcPending
								if v189 != 0 {
									return int32(0)
								} else {
									v191 = *(*int32)(unsafe.Add(mBase, uint32(v184+v185)+68))
									v192 = F_get_namespace_name(m, v191)
									mBase = m.M
									v193 = m.ExcPending
									if v193 != 0 {
										return int32(0)
									} else {
										F_initStringInfo(m, v6+int32(144))
										mBase = m.M
										v197 = m.ExcPending
										if v197 != 0 {
											return int32(0)
										} else {
											v198 = F_map_sql_identifier_to_xml_name(m)
											mBase = m.M
											v199 = m.ExcPending
											if v199 != 0 {
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
		} else {
			if l0 <= int32(1113) {
				switch l0 - int32(1042) {
				case 0:
					if l1 != int32(-1) {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l1 - int32(4)
						F_appendStringInfo(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_10), v6+int32(16))
						mBase = m.M
						v209 = m.ExcPending
						if v209 != 0 {
							return int32(0)
						} else {
							v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
							m.G0 = v6 + int32(160)
							return v229
						}
					} else {
						F_appendStringInfoString(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_11))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
							m.G0 = v6 + int32(160)
							return v229
						}
					}
				case 1:
					if l1 == int32(-1) {
						F_appendStringInfoString(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_12))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
							m.G0 = v6 + int32(160)
							return v229
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = l1 - int32(4)
						F_appendStringInfo(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_13), v6+int32(32))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
							m.G0 = v6 + int32(160)
							return v229
						}
					}
				default:
					switch l0 - int32(1082) {
					case 0:
						F_appendStringInfoString(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_14))
						mBase = m.M
						v176 = m.ExcPending
						if v176 != 0 {
							return int32(0)
						} else {
							v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
							m.G0 = v6 + int32(160)
							return v229
						}
					case 1:
						if l1 == int32(-1) {
							F_appendStringInfoString(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_15))
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return int32(0)
							} else {
								v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
								m.G0 = v6 + int32(160)
								return v229
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = l1
							F_appendStringInfo(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_16), v6-int32(-64))
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
								return int32(0)
							} else {
								v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
								m.G0 = v6 + int32(160)
								return v229
							}
						}
					default:
						v180 = F_SearchSysCache1(m, int32(82), l0)
						mBase = m.M
						v181 = m.ExcPending
						if v181 != 0 {
							return int32(0)
						} else {
							if v180 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v213 = m.ExcPending
								if v213 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
									F_errmsg_internal(m, int32(_a_F_map_sql_type_to_xml_name_1), v6)
									mBase = m.M
									v217 = m.ExcPending
									if v217 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_map_sql_type_to_xml_name_2), int32(3833), int32(_a_F_map_sql_type_to_xml_name_3))
										mBase = m.M
										v222 = m.ExcPending
										if v222 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v184 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
								v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+22)))
								v187 = *(*int32)(unsafe.Add(mBase, _c_F_map_sql_type_to_xml_name[0]))
								v188 = F_get_database_name(m, v187)
								mBase = m.M
								v189 = m.ExcPending
								if v189 != 0 {
									return int32(0)
								} else {
									v191 = *(*int32)(unsafe.Add(mBase, uint32(v184+v185)+68))
									v192 = F_get_namespace_name(m, v191)
									mBase = m.M
									v193 = m.ExcPending
									if v193 != 0 {
										return int32(0)
									} else {
										F_initStringInfo(m, v6+int32(144))
										mBase = m.M
										v197 = m.ExcPending
										if v197 != 0 {
											return int32(0)
										} else {
											v198 = F_map_sql_identifier_to_xml_name(m)
											mBase = m.M
											v199 = m.ExcPending
											if v199 != 0 {
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
				if l0 <= int32(1265) {
					if l0 == int32(1114) {
						if l1 == int32(-1) {
							F_appendStringInfoString(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_17))
							mBase = m.M
							v155 = m.ExcPending
							if v155 != 0 {
								return int32(0)
							} else {
								v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
								m.G0 = v6 + int32(160)
								return v229
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+96)) = l1
							F_appendStringInfo(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_18), v6+int32(96))
							mBase = m.M
							v163 = m.ExcPending
							if v163 != 0 {
								return int32(0)
							} else {
								v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
								m.G0 = v6 + int32(160)
								return v229
							}
						}
					} else {
						if l0 != int32(1184) {
							v180 = F_SearchSysCache1(m, int32(82), l0)
							mBase = m.M
							v181 = m.ExcPending
							if v181 != 0 {
								return int32(0)
							} else {
								if v180 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v213 = m.ExcPending
									if v213 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
										F_errmsg_internal(m, int32(_a_F_map_sql_type_to_xml_name_1), v6)
										mBase = m.M
										v217 = m.ExcPending
										if v217 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_map_sql_type_to_xml_name_2), int32(3833), int32(_a_F_map_sql_type_to_xml_name_3))
											mBase = m.M
											v222 = m.ExcPending
											if v222 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v184 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
									v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+22)))
									v187 = *(*int32)(unsafe.Add(mBase, _c_F_map_sql_type_to_xml_name[0]))
									v188 = F_get_database_name(m, v187)
									mBase = m.M
									v189 = m.ExcPending
									if v189 != 0 {
										return int32(0)
									} else {
										v191 = *(*int32)(unsafe.Add(mBase, uint32(v184+v185)+68))
										v192 = F_get_namespace_name(m, v191)
										mBase = m.M
										v193 = m.ExcPending
										if v193 != 0 {
											return int32(0)
										} else {
											F_initStringInfo(m, v6+int32(144))
											mBase = m.M
											v197 = m.ExcPending
											if v197 != 0 {
												return int32(0)
											} else {
												v198 = F_map_sql_identifier_to_xml_name(m)
												mBase = m.M
												v199 = m.ExcPending
												if v199 != 0 {
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
							if l1 != int32(-1) {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+112)) = l1
								F_appendStringInfo(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_19), v6+int32(112))
								mBase = m.M
								v171 = m.ExcPending
								if v171 != 0 {
									return int32(0)
								} else {
									v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
									m.G0 = v6 + int32(160)
									return v229
								}
							} else {
								F_appendStringInfoString(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_20))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int32(0)
								} else {
									v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
									m.G0 = v6 + int32(160)
									return v229
								}
							}
						}
					}
				} else {
					if l0 == int32(1266) {
						if l1 == int32(-1) {
							F_appendStringInfoString(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_21))
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return int32(0)
							} else {
								v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
								m.G0 = v6 + int32(160)
								return v229
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = l1
							F_appendStringInfo(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_22), v6+int32(80))
							mBase = m.M
							v148 = m.ExcPending
							if v148 != 0 {
								return int32(0)
							} else {
								v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
								m.G0 = v6 + int32(160)
								return v229
							}
						}
					} else {
						if l0 != int32(1700) {
							v180 = F_SearchSysCache1(m, int32(82), l0)
							mBase = m.M
							v181 = m.ExcPending
							if v181 != 0 {
								return int32(0)
							} else {
								if v180 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v213 = m.ExcPending
									if v213 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
										F_errmsg_internal(m, int32(_a_F_map_sql_type_to_xml_name_1), v6)
										mBase = m.M
										v217 = m.ExcPending
										if v217 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_map_sql_type_to_xml_name_2), int32(3833), int32(_a_F_map_sql_type_to_xml_name_3))
											mBase = m.M
											v222 = m.ExcPending
											if v222 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v184 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
									v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+22)))
									v187 = *(*int32)(unsafe.Add(mBase, _c_F_map_sql_type_to_xml_name[0]))
									v188 = F_get_database_name(m, v187)
									mBase = m.M
									v189 = m.ExcPending
									if v189 != 0 {
										return int32(0)
									} else {
										v191 = *(*int32)(unsafe.Add(mBase, uint32(v184+v185)+68))
										v192 = F_get_namespace_name(m, v191)
										mBase = m.M
										v193 = m.ExcPending
										if v193 != 0 {
											return int32(0)
										} else {
											F_initStringInfo(m, v6+int32(144))
											mBase = m.M
											v197 = m.ExcPending
											if v197 != 0 {
												return int32(0)
											} else {
												v198 = F_map_sql_identifier_to_xml_name(m)
												mBase = m.M
												v199 = m.ExcPending
												if v199 != 0 {
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
							if l1 != int32(-1) {
								v75 = l1 - int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = v75 & int32(_a_F_map_sql_type_to_xml_name_23)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = int32(base.Ui32(v75) >> (uint(int32(16)) % 32))
								F_appendStringInfo(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_24), v6+int32(48))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
									m.G0 = v6 + int32(160)
									return v229
								}
							} else {
								F_appendStringInfoString(m, v6+int32(128), int32(_a_F_map_sql_type_to_xml_name_25))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									v229 = *(*int32)(unsafe.Add(mBase, uint32(v6)+128))
									m.G0 = v6 + int32(160)
									return v229
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_prepare_sql_fn_parse_info(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
	v17 = F_palloc0(m, int32(20))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = v15 + v14
	v24 = F_pstrdup(m, v21+int32(4))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v24
	v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+104)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v28
	if int32(0) < v28 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L37
	}
L5:
	;
	m.G0 = v12 + int32(16)
	return v17
L6:
	;
	v33 = v28 << (uint(int32(2)) % 32)
	v34 = F_palloc(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(0)
	goto L5
L9:
	;
	if v33 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	base.MemoryCopy(m, v34, v21+int32(136), v33)
	goto L12
L11:
	;
	goto L12
L12:
	;
	v43 = int32(0)
	goto L13
L13:
	;
	v51 = v34 + v43<<(uint(int32(2))%32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v52 <= int32(3830) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v34
	v87 = v12 + int32(15)
	v88 = F_SysCacheGetAttr(m, int32(46), l0, int32(23), v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L27
	}
L15:
	;
	v80 = v43 + int32(1)
	if v80 != v28 {
		v43 = v80
		goto L13
	} else {
		goto L26
	}
L16:
	;
	v73 = F_get_call_expr_argtype(m, l1, v43)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L24
	}
L17:
	;
	switch v52 - int32(2277) {
	case 0, 6:
		goto L16
	case 1, 2, 3, 4, 5:
		goto L15
	default:
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if base.B2i32(base.Ui32(v52-int32(_a_F_prepare_sql_fn_parse_info_0)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v52-int32(_a_F_prepare_sql_fn_parse_info_1)) < base.Ui32(int32(2))) != 0 {
		goto L16
	} else {
		goto L22
	}
L20:
	;
	if base.B2i32(v52 == int32(2776))|base.B2i32(v52 == int32(3500)) != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L15
L22:
	;
	if v52 != int32(3831) {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	goto L16
L24:
	;
	if v73 == int32(0) {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v73
	goto L15
L26:
	;
	goto L14
L27:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v90 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v91 = int32(0)
	goto L30
L29:
	;
	v91 = v88
	goto L30
L30:
	;
	v95 = F_SysCacheGetAttr(m, int32(46), l0, int32(22), v87)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v97 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v98 = int32(0)
	goto L34
L33:
	;
	v98 = v95
	goto L34
L34:
	;
	v100 = v17 + int32(12)
	v101 = F_get_func_input_arg_names(m, v91, v98, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v28 <= v101 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = int32(0)
	goto L5
L37:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v129 = F_format_type_be(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v129
	F_errmsg(m, int32(_a_F_prepare_sql_fn_parse_info_2), v12)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_prepare_sql_fn_parse_info_3), int32(293), int32(_a_F_prepare_sql_fn_parse_info_4))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_read_sql_stmt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v5 = int32(0)
	v13 = F_read_sql_construct(m, int32(59), v5, v5, int32(_a_F_read_sql_stmt_0), v5, v5, int32(1), v5, v5, l0, l1, l2)
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		return v13
	}
}
func F_sql_delete_callback(m *base.Module, l0 int32) {
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
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v4 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v29 != 0 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v7 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = v2
	goto L4
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13+v12<<(uint(int32(2))%32))))
	F_DropCachedPlan(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	v21 = v12 + int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v21 < v22 {
		v12 = v21
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	F_MemoryContextDelete(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(0)
	return
L12:
	;
	goto L11
}
func F_sql_fn_resolve_param_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	v4 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14 <= int32(0) {
		v95 = v4
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v95
L5:
	;
	v21 = v4
	goto L6
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v9+v21<<(uint(int32(2))%32))))
	if v28 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v95 = v4
	goto L4
L8:
	;
	v90 = v21 + int32(1)
	if v90 != v14 {
		v21 = v90
		goto L6
	} else {
		goto L22
	}
L9:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v33 == int32(0))|base.B2i32(v33 != v36) != 0 {
		v54 = v33
		v55 = v36
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v54-v55 != 0 {
		goto L8
	} else {
		goto L17
	}
L11:
	;
	goto L10
L12:
	;
	v39 = v28
	v40 = l1
	goto L13
L13:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	if v44 == int32(0) {
		v54 = v44
		v55 = v43
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v54 = v44
	v55 = v43
	goto L11
L15:
	;
	v47 = int32(1)
	if v44 == v43 {
		v39 = v39 + v47
		v40 = v40 + v47
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v58 = F_palloc0(m, int32(28))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	v63 = v21 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v63
	*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(8)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v67+v63<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v73
	v77 = F_get_typcollation(m, v73)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v77
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v82 = int32(0)
	if base.B2i32(v81 == v82)|base.B2i32(v77 == v82) != 0 {
		v95 = v58
		goto L4
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v81
	return v58
L22:
	;
	goto L7
}
