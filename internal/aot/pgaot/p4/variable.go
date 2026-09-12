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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v127 int64
	_ = v127
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 == v3 {
		v149 = v3
		m.G0 = v9 + int32(16)
		return v149
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v13 != int32(67) {
			if v13 != int32(30) {
				if v13 != int32(6) {
					v145 = F_expression_tree_mutator_impl(m, l0, int32(1055), l1)
					mBase = m.M
					v146 = m.ExcPending
					if v146 != 0 {
						return int32(0)
					} else {
						v149 = v145
						m.G0 = v9 + int32(16)
						return v149
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v20 != v21 {
						v145 = F_expression_tree_mutator_impl(m, l0, int32(1055), l1)
						mBase = m.M
						v146 = m.ExcPending
						if v146 != 0 {
							return int32(0)
						} else {
							v149 = v145
							m.G0 = v9 + int32(16)
							return v149
						}
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						if v23 != v24 {
							v145 = F_expression_tree_mutator_impl(m, l0, int32(1055), l1)
							mBase = m.M
							v146 = m.ExcPending
							if v146 != 0 {
								return int32(0)
							} else {
								v149 = v145
								m.G0 = v9 + int32(16)
								return v149
							}
						} else {
							v27 = F_palloc(m, int32(48))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int32(0)
							} else {
								v32 = l0 + int32(8)
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32))))
								v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v27)+40)) = v34
								v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v27)+32)) = v36
								v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v27)+24)) = v38
								v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v27)+16)) = v40
								v42 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
								*(*int64)(unsafe.Add(mBase, uint32(v27)+8)) = v42
								v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
								*(*int64)(unsafe.Add(mBase, uint32(v27))) = v44
								if int32(0) < v33 {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
									if v49 < v33 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v160 = m.ExcPending
										if v160 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = v33
											F_errmsg_internal(m, int32(451867), v9)
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(495733), int32(1604), int32(208667))
												mBase = m.M
												v169 = m.ExcPending
												if v169 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
										v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51+v33<<(uint(int32(1))%32)-int32(2)))))
										if v57 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v160 = m.ExcPending
											if v160 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9))) = v33
												F_errmsg_internal(m, int32(451867), v9)
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(495733), int32(1604), int32(208667))
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											*(*uint16)(unsafe.Add(mBase, uint32(v27)+8)) = uint16(v57)
											v61 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
											v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											if v61 != v62 {
												v149 = v27
											} else {
												*(*uint16)(unsafe.Add(mBase, uint32(v27)+40)) = uint16(v57)
												v149 = v27
											}
											m.G0 = v9 + int32(16)
											return v149
										}
									}
								} else {
									if v33 != 0 {
										v149 = v27
										m.G0 = v9 + int32(16)
										return v149
									} else {
										v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
										v66 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v65))) = uint8(v66)
										v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
										if v68 == int32(0) {
											v149 = v27
											m.G0 = v9 + int32(16)
											return v149
										} else {
											v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v68 == v71 {
												v149 = v27
												m.G0 = v9 + int32(16)
												return v149
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v68
												v75 = F_palloc0(m, int32(20))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v27
													*(*int32)(unsafe.Add(mBase, uint32(v75))) = int32(30)
													v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int64)(unsafe.Add(mBase, uint32(v75)+12)) = int64(-4294967294)
													*(*int32)(unsafe.Add(mBase, uint32(v75)+8)) = v80
													v149 = v75
													m.G0 = v9 + int32(16)
													return v149
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
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
				if v85 != int32(6) {
					v145 = F_expression_tree_mutator_impl(m, l0, int32(1055), l1)
					mBase = m.M
					v146 = m.ExcPending
					if v146 != 0 {
						return int32(0)
					} else {
						v149 = v145
						m.G0 = v9 + int32(16)
						return v149
					}
				} else {
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
					v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v88 != v89 {
						v145 = F_expression_tree_mutator_impl(m, l0, int32(1055), l1)
						mBase = m.M
						v146 = m.ExcPending
						if v146 != 0 {
							return int32(0)
						} else {
							v149 = v145
							m.G0 = v9 + int32(16)
							return v149
						}
					} else {
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v84)+28))
						v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						if v91 != v92 {
							v145 = F_expression_tree_mutator_impl(m, l0, int32(1055), l1)
							mBase = m.M
							v146 = m.ExcPending
							if v146 != 0 {
								return int32(0)
							} else {
								v149 = v145
								m.G0 = v9 + int32(16)
								return v149
							}
						} else {
							v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+8)))
							if v94 != 0 {
								v145 = F_expression_tree_mutator_impl(m, l0, int32(1055), l1)
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return int32(0)
								} else {
									v149 = v145
									m.G0 = v9 + int32(16)
									return v149
								}
							} else {
								v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
								if v95 == int32(0) {
									v145 = F_expression_tree_mutator_impl(m, l0, int32(1055), l1)
									mBase = m.M
									v146 = m.ExcPending
									if v146 != 0 {
										return int32(0)
									} else {
										v149 = v145
										m.G0 = v9 + int32(16)
										return v149
									}
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
									if v95 == v98 {
										v145 = F_expression_tree_mutator_impl(m, l0, int32(1055), l1)
										mBase = m.M
										v146 = m.ExcPending
										if v146 != 0 {
											return int32(0)
										} else {
											v149 = v145
											m.G0 = v9 + int32(16)
											return v149
										}
									} else {
										v101 = F_palloc(m, int32(48))
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return int32(0)
										} else {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
											v104 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v104)
											v106 = *(*int64)(unsafe.Add(mBase, uint32(v84)+40))
											*(*int64)(unsafe.Add(mBase, uint32(v101)+40)) = v106
											v108 = *(*int64)(unsafe.Add(mBase, uint32(v84)+32))
											*(*int64)(unsafe.Add(mBase, uint32(v101)+32)) = v108
											v110 = *(*int64)(unsafe.Add(mBase, uint32(v84)+24))
											*(*int64)(unsafe.Add(mBase, uint32(v101)+24)) = v110
											v112 = *(*int64)(unsafe.Add(mBase, uint32(v84)+16))
											*(*int64)(unsafe.Add(mBase, uint32(v101)+16)) = v112
											v114 = *(*int64)(unsafe.Add(mBase, uint32(v84)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v101)+8)) = v114
											v116 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
											*(*int64)(unsafe.Add(mBase, uint32(v101))) = v116
											v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v101)+12)) = v118
											v121 = F_palloc(m, int32(20))
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return int32(0)
											} else {
												v123 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
												*(*int64)(unsafe.Add(mBase, uint32(v121))) = v123
												v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v121)+16)) = v125
												v127 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int64)(unsafe.Add(mBase, uint32(v121)+8)) = v127
												*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v101
												v149 = v121
												m.G0 = v9 + int32(16)
												return v149
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
			v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v130 + int32(1)
			v136 = F_query_tree_mutator_impl(m, l0, int32(1055), l1, int32(0))
			mBase = m.M
			v137 = m.ExcPending
			if v137 != 0 {
				return int32(0)
			} else {
				v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v138 - int32(1)
				v149 = v136
				m.G0 = v9 + int32(16)
				return v149
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
						F_errmsg(m, int32(466762), v11)
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
								F_errfinish(m, int32(497051), int32(206), int32(315073))
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
							F_errmsg(m, int32(466762), v11)
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
									F_errfinish(m, int32(497051), int32(206), int32(315073))
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
									F_errmsg(m, int32(466717), v11+int32(32))
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
												F_errdetail(m, int32(179708), v11+int32(16))
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
														F_errfinish(m, int32(497051), int32(227), int32(315073))
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
