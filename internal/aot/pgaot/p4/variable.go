package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_map_variable_attnos_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int64
	_ = v120
	var v122 int32
	_ = v122
	var v124 int64
	_ = v124
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 == v3 {
		v146 = v3
		m.G0 = v8 + int32(16)
		return v146
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v12 != int32(67) {
			if v12 != int32(30) {
				if v12 != int32(6) {
					v142 = F_expression_tree_mutator_impl(m, l0, int32(1055), l1)
					mBase = m.M
					v143 = m.ExcPending
					if v143 != 0 {
						return int32(0)
					} else {
						v146 = v142
						m.G0 = v8 + int32(16)
						return v146
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v19 != v20 {
						v142 = F_expression_tree_mutator_impl(m, l0, int32(1055), l1)
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return int32(0)
						} else {
							v146 = v142
							m.G0 = v8 + int32(16)
							return v146
						}
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						if v22 != v23 {
							v142 = F_expression_tree_mutator_impl(m, l0, int32(1055), l1)
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return int32(0)
							} else {
								v146 = v142
								m.G0 = v8 + int32(16)
								return v146
							}
						} else {
							v26 = F_palloc(m, int32(48))
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
								v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v26)+40)) = v31
								v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v33
								v35 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v26)+24)) = v35
								v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = v37
								v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = v39
								v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
								*(*int64)(unsafe.Add(mBase, uint32(v26))) = v41
								if int32(0) < v30 {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
									if v46 < v30 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v30
											F_errmsg_internal(m, int32(_a_F_map_variable_attnos_mutator_0), v8)
											mBase = m.M
											v160 = m.ExcPending
											if v160 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_map_variable_attnos_mutator_1), int32(1604), int32(_a_F_map_variable_attnos_mutator_2))
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
										v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48+v30<<(uint(int32(1))%32)-int32(2)))))
										if v54 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v156 = m.ExcPending
											if v156 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v30
												F_errmsg_internal(m, int32(_a_F_map_variable_attnos_mutator_0), v8)
												mBase = m.M
												v160 = m.ExcPending
												if v160 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_map_variable_attnos_mutator_1), int32(1604), int32(_a_F_map_variable_attnos_mutator_2))
													mBase = m.M
													v165 = m.ExcPending
													if v165 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											*(*uint16)(unsafe.Add(mBase, uint32(v26)+8)) = uint16(v54)
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
											v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											if v58 != v59 {
												v146 = v26
											} else {
												*(*uint16)(unsafe.Add(mBase, uint32(v26)+40)) = uint16(v54)
												v146 = v26
											}
											m.G0 = v8 + int32(16)
											return v146
										}
									}
								} else {
									if v30 != 0 {
										v146 = v26
										m.G0 = v8 + int32(16)
										return v146
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
										v63 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v63)
										v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
										if v65 == int32(0) {
											v146 = v26
											m.G0 = v8 + int32(16)
											return v146
										} else {
											v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v65 == v68 {
												v146 = v26
												m.G0 = v8 + int32(16)
												return v146
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v65
												v72 = F_palloc0(m, int32(20))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v26
													*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(30)
													v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int64)(unsafe.Add(mBase, uint32(v72)+12)) = int64(-4294967294)
													*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v77
													v146 = v72
													m.G0 = v8 + int32(16)
													return v146
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
				v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
				if v82 != int32(6) {
					v142 = F_expression_tree_mutator_impl(m, l0, int32(1055), l1)
					mBase = m.M
					v143 = m.ExcPending
					if v143 != 0 {
						return int32(0)
					} else {
						v146 = v142
						m.G0 = v8 + int32(16)
						return v146
					}
				} else {
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
					v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v85 != v86 {
						v142 = F_expression_tree_mutator_impl(m, l0, int32(1055), l1)
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return int32(0)
						} else {
							v146 = v142
							m.G0 = v8 + int32(16)
							return v146
						}
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v81)+28))
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						if v88 != v89 {
							v142 = F_expression_tree_mutator_impl(m, l0, int32(1055), l1)
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return int32(0)
							} else {
								v146 = v142
								m.G0 = v8 + int32(16)
								return v146
							}
						} else {
							v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+8)))
							if v91 != 0 {
								v142 = F_expression_tree_mutator_impl(m, l0, int32(1055), l1)
								mBase = m.M
								v143 = m.ExcPending
								if v143 != 0 {
									return int32(0)
								} else {
									v146 = v142
									m.G0 = v8 + int32(16)
									return v146
								}
							} else {
								v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
								if v92 == int32(0) {
									v142 = F_expression_tree_mutator_impl(m, l0, int32(1055), l1)
									mBase = m.M
									v143 = m.ExcPending
									if v143 != 0 {
										return int32(0)
									} else {
										v146 = v142
										m.G0 = v8 + int32(16)
										return v146
									}
								} else {
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
									if v92 == v95 {
										v142 = F_expression_tree_mutator_impl(m, l0, int32(1055), l1)
										mBase = m.M
										v143 = m.ExcPending
										if v143 != 0 {
											return int32(0)
										} else {
											v146 = v142
											m.G0 = v8 + int32(16)
											return v146
										}
									} else {
										v98 = F_palloc(m, int32(48))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
											v101 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v101)
											v103 = *(*int64)(unsafe.Add(mBase, uint32(v81)+40))
											*(*int64)(unsafe.Add(mBase, uint32(v98)+40)) = v103
											v105 = *(*int64)(unsafe.Add(mBase, uint32(v81)+32))
											*(*int64)(unsafe.Add(mBase, uint32(v98)+32)) = v105
											v107 = *(*int64)(unsafe.Add(mBase, uint32(v81)+24))
											*(*int64)(unsafe.Add(mBase, uint32(v98)+24)) = v107
											v109 = *(*int64)(unsafe.Add(mBase, uint32(v81)+16))
											*(*int64)(unsafe.Add(mBase, uint32(v98)+16)) = v109
											v111 = *(*int64)(unsafe.Add(mBase, uint32(v81)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v98)+8)) = v111
											v113 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
											*(*int64)(unsafe.Add(mBase, uint32(v98))) = v113
											v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v98)+12)) = v115
											v118 = F_palloc(m, int32(20))
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int32(0)
											} else {
												v120 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
												*(*int64)(unsafe.Add(mBase, uint32(v118))) = v120
												v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v118)+16)) = v122
												v124 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int64)(unsafe.Add(mBase, uint32(v118)+8)) = v124
												*(*int32)(unsafe.Add(mBase, uint32(v118)+4)) = v98
												v146 = v118
												m.G0 = v8 + int32(16)
												return v146
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
			v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v127 + int32(1)
			v133 = F_query_tree_mutator_impl(m, l0, int32(1055), l1, int32(0))
			mBase = m.M
			v134 = m.ExcPending
			if v134 != 0 {
				return int32(0)
			} else {
				v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v135 - int32(1)
				v146 = v133
				m.G0 = v8 + int32(16)
				return v146
			}
		}
	}
}
func F_variable_coerce_param_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	v6 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v14 != 0 {
		v50 = v6
		m.G0 = v11 + int32(48)
		return v50
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		if v15 != int32(705) {
			v50 = v6
			m.G0 = v11 + int32(48)
			return v50
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v18 <= int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(33685636))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v18
						F_errmsg(m, int32(_a_F_variable_coerce_param_hook_0), v11)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							F_parser_errposition(m, l0, v69)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_variable_coerce_param_hook_1), int32(206), int32(_a_F_variable_coerce_param_hook_2))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
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
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				if v23 < v18 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(33685636))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v18
							F_errmsg(m, int32(_a_F_variable_coerce_param_hook_0), v11)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								F_parser_errposition(m, l0, v69)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_variable_coerce_param_hook_1), int32(206), int32(_a_F_variable_coerce_param_hook_2))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
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
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					v31 = v26 + v18<<(uint(int32(2))%32) - int32(4)
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					if v32 == int32(705) {
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = l2
						*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = l2
						v40 = F_get_typcollation(m, l2)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v40
							if l4 < int32(0) {
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								if base.Ui32(v47) <= base.Ui32(l4) {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l4
								}
							}
							v50 = l1
							m.G0 = v11 + int32(48)
							return v50
						}
					} else {
						if l2 != v32 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134348932))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v18
									F_errmsg(m, int32(_a_F_variable_coerce_param_hook_3), v11+int32(32))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
										v91 = F_format_type_be(m, v90)
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int32(0)
										} else {
											v93 = F_format_type_be(m, l2)
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v93
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v91
												F_errdetail(m, int32(_a_F_variable_coerce_param_hook_4), v11+int32(16))
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return int32(0)
												} else {
													v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
													F_parser_errposition(m, l0, v102)
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_variable_coerce_param_hook_1), int32(227), int32(_a_F_variable_coerce_param_hook_2))
														mBase = m.M
														v109 = m.ExcPending
														if v109 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = l2
							v40 = F_get_typcollation(m, l2)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v40
								if l4 < int32(0) {
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									if base.Ui32(v47) <= base.Ui32(l4) {
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l4
									}
								}
								v50 = l1
								m.G0 = v11 + int32(48)
								return v50
							}
						}
					}
				}
			}
		}
	}
}
