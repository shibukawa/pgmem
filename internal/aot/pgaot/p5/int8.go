package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int8_accum_inv(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v43 int64
	_ = v43
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v65 int64
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v14 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v111 = m.ExcPending
		if v111 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(336128), int32(0))
			mBase = m.M
			v115 = m.ExcPending
			if v115 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(476611), int32(6151), int32(30425))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v15 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v111 = m.ExcPending
			if v111 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(336128), int32(0))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(476611), int32(6151), int32(30425))
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v18 == int32(0) {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(0)
				v26 = F_palloc(m, int32(12))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v26
					v31 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v31)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v26 + int32(2)
					if v22 < int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(16384)
						v47 = int64(0) - v22
						v50 = v31
						v53 = v26 + int32(12)
						v57 = v47
						for {
							v60 = v53 - int32(2)
							v62 = base.I64_div_u_s(v57, int64(10000))
							v65 = v62*int64(55536) + v57
							*(*uint16)(unsafe.Add(mBase, uint32(v60))) = uint16(v65)
							v68 = v50 + int32(1)
							if base.Ui64(int64(9999)) < base.Ui64(v57) {
								v50 = v68
								v53 = v60
								v57 = v62
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v60
						v72 = v68
						v77 = v50
					} else {
						v43 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v43
						if v22 == v43 {
							v72 = v31
							v77 = int32(0)
						} else {
							v47 = v22
							v50 = v31
							v53 = v26 + int32(12)
							v57 = v47
							for {
								v60 = v53 - int32(2)
								v62 = base.I64_div_u_s(v57, int64(10000))
								v65 = v62*int64(55536) + v57
								*(*uint16)(unsafe.Add(mBase, uint32(v60))) = uint16(v65)
								v68 = v50 + int32(1)
								if base.Ui64(int64(9999)) < base.Ui64(v57) {
									v50 = v68
									v53 = v60
									v57 = v62
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v60
							v72 = v68
							v77 = v50
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v77
					*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v72
					v86 = F_make_result_opt_error(m, v12+int32(8), int32(0))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v26)
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							v90 = F_do_numeric_discard(m, v15, v86)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								if v90 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v124 = m.ExcPending
									if v124 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(18979), int32(0))
										mBase = m.M
										v128 = m.ExcPending
										if v128 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(476611), int32(6157), int32(30425))
											mBase = m.M
											v133 = m.ExcPending
											if v133 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									m.G0 = v12 + int32(32)
									return v15
								}
							}
						}
					}
				}
			} else {
				m.G0 = v12 + int32(32)
				return v15
			}
		}
	}
}
func F_int8_numeric(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v35 int64
	_ = v35
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v56 int64
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(0)
	v18 = F_palloc(m, int32(12))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v18
		v23 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v18))) = uint16(v23)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v18 + int32(2)
		if v14 < int64(0) {
			*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(16384)
			v39 = int64(0) - v14
			v42 = v23
			v45 = v18 + int32(12)
			v48 = v39
			for {
				v51 = v45 - int32(2)
				v53 = base.I64_div_u_s(v48, int64(10000))
				v56 = v53*int64(55536) + v48
				*(*uint16)(unsafe.Add(mBase, uint32(v51))) = uint16(v56)
				v59 = v42 + int32(1)
				if base.Ui64(int64(9999)) < base.Ui64(v48) {
					v42 = v59
					v45 = v51
					v48 = v53
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v51
			v63 = v59
			v67 = v42
		} else {
			v35 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v35
			if v14 == v35 {
				v63 = v23
				v67 = int32(0)
			} else {
				v39 = v14
				v42 = v23
				v45 = v18 + int32(12)
				v48 = v39
				for {
					v51 = v45 - int32(2)
					v53 = base.I64_div_u_s(v48, int64(10000))
					v56 = v53*int64(55536) + v48
					*(*uint16)(unsafe.Add(mBase, uint32(v51))) = uint16(v56)
					v59 = v42 + int32(1)
					if base.Ui64(int64(9999)) < base.Ui64(v48) {
						v42 = v59
						v45 = v51
						v48 = v53
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v51
				v63 = v59
				v67 = v42
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v67
		*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v63
		v76 = F_make_result_opt_error(m, v11+int32(8), int32(0))
		mBase = m.M
		v77 = m.ExcPending
		if v77 != 0 {
			return int32(0)
		} else {
			F_pfree(m, v18)
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				m.G0 = v11 + int32(32)
				return v76
			}
		}
	}
}
func F_int8_sum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v44 int64
	_ = v44
	var v48 int64
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v66 int64
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
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
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v117 int64
	_ = v117
	var v121 int64
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int64
	_ = v131
	var v134 int32
	_ = v134
	var v136 int64
	_ = v136
	var v139 int64
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v14 == int32(1) {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v17 == int32(1) {
			v20 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
			v172 = v2
			m.G0 = v12 + int32(32)
			return v172
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
			*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(0)
			v27 = F_palloc(m, int32(12))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v27
				v32 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v27 + int32(2)
				if v23 < int64(0) {
					*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(16384)
					v48 = int64(0) - v23
					v51 = v32
					v54 = v27 + int32(12)
					v58 = v48
					for {
						v61 = v54 - int32(2)
						v63 = base.I64_div_u_s(v58, int64(10000))
						v66 = v63*int64(55536) + v58
						*(*uint16)(unsafe.Add(mBase, uint32(v61))) = uint16(v66)
						v69 = v51 + int32(1)
						if base.Ui64(int64(9999)) < base.Ui64(v58) {
							v51 = v69
							v54 = v61
							v58 = v63
							continue
						} else {
							break
						}
						break
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v61
					v73 = v69
					v78 = v51
				} else {
					v44 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v44
					if v23 == v44 {
						v73 = v32
						v78 = v2
					} else {
						v48 = v23
						v51 = v32
						v54 = v27 + int32(12)
						v58 = v48
						for {
							v61 = v54 - int32(2)
							v63 = base.I64_div_u_s(v58, int64(10000))
							v66 = v63*int64(55536) + v58
							*(*uint16)(unsafe.Add(mBase, uint32(v61))) = uint16(v66)
							v69 = v51 + int32(1)
							if base.Ui64(int64(9999)) < base.Ui64(v58) {
								v51 = v69
								v54 = v61
								v58 = v63
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v61
						v73 = v69
						v78 = v51
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v78
				*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v73
				v87 = F_make_result_opt_error(m, v12+int32(8), int32(0))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v27)
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						v172 = v87
						m.G0 = v12 + int32(32)
						return v172
					}
				}
			}
		}
	} else {
		v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v92 = F_pg_detoast_datum(m, v91)
		mBase = m.M
		v93 = m.ExcPending
		if v93 != 0 {
			return int32(0)
		} else {
			v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v94 == int32(1) {
				v172 = v92
				m.G0 = v12 + int32(32)
				return v172
			} else {
				v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v98 = *(*int64)(unsafe.Add(mBase, uint32(v97)))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(0)
				v102 = F_palloc(m, int32(12))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v102
					v105 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v102))) = uint16(v105)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v102 + int32(2)
					if v98 < int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(16384)
						v121 = int64(0) - v98
						v124 = v105
						v127 = v102 + int32(12)
						v131 = v121
						for {
							v134 = v127 - int32(2)
							v136 = base.I64_div_u_s(v131, int64(10000))
							v139 = v136*int64(55536) + v131
							*(*uint16)(unsafe.Add(mBase, uint32(v134))) = uint16(v139)
							v142 = v124 + int32(1)
							if base.Ui64(int64(9999)) < base.Ui64(v131) {
								v124 = v142
								v127 = v134
								v131 = v136
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v134
						v146 = v142
						v151 = v124
					} else {
						v117 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v117
						if v98 == v117 {
							v146 = v105
							v151 = v2
						} else {
							v121 = v98
							v124 = v105
							v127 = v102 + int32(12)
							v131 = v121
							for {
								v134 = v127 - int32(2)
								v136 = base.I64_div_u_s(v131, int64(10000))
								v139 = v136*int64(55536) + v131
								*(*uint16)(unsafe.Add(mBase, uint32(v134))) = uint16(v139)
								v142 = v124 + int32(1)
								if base.Ui64(int64(9999)) < base.Ui64(v131) {
									v124 = v142
									v127 = v134
									v131 = v136
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v134
							v146 = v142
							v151 = v124
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v151
					*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v146
					v160 = F_make_result_opt_error(m, v12+int32(8), int32(0))
					mBase = m.M
					v161 = m.ExcPending
					if v161 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v102)
						mBase = m.M
						v163 = m.ExcPending
						if v163 != 0 {
							return int32(0)
						} else {
							v166 = F_DirectFunctionCall2Coll(m, int32(1293), int32(0), v92, v160)
							mBase = m.M
							v167 = m.ExcPending
							if v167 != 0 {
								return int32(0)
							} else {
								v172 = v166
								m.G0 = v12 + int32(32)
								return v172
							}
						}
					}
				}
			}
		}
	}
}
