package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_calc_hist_selectivity_scalar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64 {
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v62 int32
	_ = v62
	var v65 float64
	_ = v65
	var v66 int32
	_ = v66
	var v70 float64
	_ = v70
	v14 = l3 - int32(1)
	v18 = v14
	v20 = int32(-1)
	goto L1
L1:
	;
	v30 = base.I32_div_s(v18+v20+int32(1), int32(2))
	v34 = F_range_cmp_bounds(m, l0, l2+v30<<(uint(int32(4))%32), l1)
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v49 = int32(0)
	if v49 < v44 {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	return float64(0)
L4:
	;
	v38 = int32(0)
	v43 = base.B2i32(v34 < v38) | l4&base.B2i32(v34 == v38)
	if v43 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v44 = v30
	goto L7
L6:
	;
	v44 = v20
	goto L7
L7:
	;
	if v43 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v47 = v18
	goto L10
L9:
	;
	v47 = v30 - int32(1)
	goto L10
L10:
	;
	if v44 < v47 {
		v18 = v47
		v20 = v44
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L2
L12:
	;
	v52 = v44
	goto L14
L13:
	;
	v52 = v49
	goto L14
L14:
	;
	v54 = base.F64_convert_i32_u(v14)
	v55 = base.F64_div(base.F64_convert_i32_u(v52), v54)
	if base.B2i32(v44 < int32(0))|base.B2i32(v14 <= v44) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v70 = v55
	goto L17
L16:
	;
	v62 = l2 + v44<<(uint(int32(4))%32)
	v65 = F_get_position(m, l0, l1, v62, v62+int32(16))
	v66 = m.ExcPending
	if v66 != 0 {
		goto L3
	} else {
		goto L18
	}
L17:
	;
	return v70
L18:
	;
	v70 = base.F64_add(v55, base.F64_div(v65, v54))
	goto L17
}
func F_calc_key_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v13 = F_pgp_load_digest(m, int32(2), v8+int32(28))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v13 {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
			switch v22 - int32(1) {
			case 0, 1, 2:
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
				v55 = v36 + v38 + int32(10)
			default:
				v55 = int32(6)
			case 15:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
				v55 = v26 + (v28 + v30) + int32(12)
			case 16:
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
				v55 = v43 + (v45 + (v47 + v49)) + int32(14)
			}
			v56 = int32(153)
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+25)) = uint8(v56)
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+27)) = uint8(v55)
			v59 = int32(8)
			v61 = int32(base.Ui32(v55) >> (uint(v59) % 32))
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)) = uint8(v61)
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
			v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
			m.T0[v67].(func(*base.Module, int32, int32, int32))(m, v63, v8+int32(25), int32(3))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
				v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
				m.T0[v72].(func(*base.Module, int32, int32, int32))(m, v70, l0, int32(1))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
					m.T0[v79].(func(*base.Module, int32, int32, int32))(m, v75, l0+int32(1), int32(4))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
						m.T0[v84].(func(*base.Module, int32, int32, int32))(m, v82, l0+int32(5), int32(1))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							v87 = int32(12)
							v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
							switch v88 - int32(1) {
							case 0, 1, 2:
								v106 = v87
								v108 = v59
								v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
								v111 = *(*int32)(unsafe.Add(mBase, uint32(l0+v108)))
								v112 = F_pgp_mpi_hash(m, v109, v111)
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int32(0)
								} else {
									v114 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
									v116 = *(*int32)(unsafe.Add(mBase, uint32(l0+v106)))
									v117 = F_pgp_mpi_hash(m, v114, v116)
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
										v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
										m.T0[v123].(func(*base.Module, int32, int32))(m, v122, v8)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int32(0)
										} else {
											v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
											v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
											m.T0[v127].(func(*base.Module, int32))(m, v126)
											mBase = m.M
											v129 = m.ExcPending
											if v129 != 0 {
												return int32(0)
											} else {
												v130 = *(*int64)(unsafe.Add(mBase, uint32(v8)+12))
												*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v130
												v132 = int32(0)
												base.MemoryFill(m, v8, v132, int32(20))
												v136 = v132
												m.G0 = v8 + int32(32)
												return v136
											}
										}
									}
								}
							default:
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
								v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
								m.T0[v123].(func(*base.Module, int32, int32))(m, v122, v8)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
									v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
									m.T0[v127].(func(*base.Module, int32))(m, v126)
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int32(0)
									} else {
										v130 = *(*int64)(unsafe.Add(mBase, uint32(v8)+12))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v130
										v132 = int32(0)
										base.MemoryFill(m, v8, v132, int32(20))
										v136 = v132
										m.G0 = v8 + int32(32)
										return v136
									}
								}
							case 15:
								v98 = v87
								v99 = v88
								v100 = v59
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
								v103 = *(*int32)(unsafe.Add(mBase, uint32(l0+v100)))
								v104 = F_pgp_mpi_hash(m, v101, v103)
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									v106 = v99
									v108 = v98
									v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
									v111 = *(*int32)(unsafe.Add(mBase, uint32(l0+v108)))
									v112 = F_pgp_mpi_hash(m, v109, v111)
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										v114 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
										v116 = *(*int32)(unsafe.Add(mBase, uint32(l0+v106)))
										v117 = F_pgp_mpi_hash(m, v114, v116)
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
											v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
											m.T0[v123].(func(*base.Module, int32, int32))(m, v122, v8)
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return int32(0)
											} else {
												v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
												v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
												m.T0[v127].(func(*base.Module, int32))(m, v126)
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													v130 = *(*int64)(unsafe.Add(mBase, uint32(v8)+12))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v130
													v132 = int32(0)
													base.MemoryFill(m, v8, v132, int32(20))
													v136 = v132
													m.G0 = v8 + int32(32)
													return v136
												}
											}
										}
									}
								}
							case 16:
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
								v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v93 = F_pgp_mpi_hash(m, v91, v92)
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									v98 = int32(16)
									v99 = int32(20)
									v100 = int32(12)
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
									v103 = *(*int32)(unsafe.Add(mBase, uint32(l0+v100)))
									v104 = F_pgp_mpi_hash(m, v101, v103)
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int32(0)
									} else {
										v106 = v99
										v108 = v98
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
										v111 = *(*int32)(unsafe.Add(mBase, uint32(l0+v108)))
										v112 = F_pgp_mpi_hash(m, v109, v111)
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
											v116 = *(*int32)(unsafe.Add(mBase, uint32(l0+v106)))
											v117 = F_pgp_mpi_hash(m, v114, v116)
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return int32(0)
											} else {
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
												v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
												m.T0[v123].(func(*base.Module, int32, int32))(m, v122, v8)
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int32(0)
												} else {
													v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
													v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
													m.T0[v127].(func(*base.Module, int32))(m, v126)
													mBase = m.M
													v129 = m.ExcPending
													if v129 != 0 {
														return int32(0)
													} else {
														v130 = *(*int64)(unsafe.Add(mBase, uint32(v8)+12))
														*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v130
														v132 = int32(0)
														base.MemoryFill(m, v8, v132, int32(20))
														v136 = v132
														m.G0 = v8 + int32(32)
														return v136
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
			v136 = v13
			m.G0 = v8 + int32(32)
			return v136
		}
	}
}
func F_case_index(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	v2 = int32(0)
	if base.Ui32(l0) <= base.Ui32(int32(1415)) {
		v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
		v134 = v7
	} else {
		if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_0)) {
			if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_1)) {
				if base.Ui32(l0-int32(_a_F_case_index_2)) <= base.Ui32(int32(95)) {
					v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[1]))))
					v134 = v20
				} else {
					if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_3)) {
						v134 = v2
					} else {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_4)) {
							v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[2]))))
							v134 = v29
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_5)) {
								v134 = v2
							} else {
								v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[3]))))
								v134 = v36
							}
						}
					}
				}
			} else {
				if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_6)) {
					v134 = v2
				} else {
					if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_7)) {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_8)) {
							v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[4]))))
							v134 = v47
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_9)) {
								v134 = v2
							} else {
								v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[5]))))
								v134 = v54
							}
						}
					} else {
						if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_10)) {
							v134 = v2
						} else {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_11)) {
								v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[6]))))
								v134 = v63
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_12)) {
									v134 = v2
								} else {
									v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[7]))))
									v134 = v70
								}
							}
						}
					}
				}
			}
		} else {
			if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_13)) {
				v134 = v2
			} else {
				if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_14)) {
					if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_15)) {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_16)) {
							v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[8]))))
							v134 = v83
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_17)) {
								v134 = v2
							} else {
								v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[9]))))
								v134 = v90
							}
						}
					} else {
						if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_18)) {
							v134 = v2
						} else {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_19)) {
								v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[10]))))
								v134 = v99
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_20)) {
									v134 = v2
								} else {
									v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[11]))))
									v134 = v106
								}
							}
						}
					}
				} else {
					if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_21)) {
						v134 = v2
					} else {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_22)) {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_23)) {
								v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[12]))))
								v134 = v117
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_24)) {
									v134 = v2
								} else {
									v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[13]))))
									v134 = v124
								}
							}
						} else {
							if base.Ui32(int32(67)) < base.Ui32(l0-int32(_a_F_case_index_25)) {
								v134 = v2
							} else {
								v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[14]))))
								v134 = v133
							}
						}
					}
				}
			}
		}
	}
	return v134
}
func F_catalan_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v5
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 < v8 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v236
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v236
	v240 = v236 - int32(1)
	if v240 <= v8 {
		goto L65
	} else {
		goto L66
	}
L2:
	;
	if v57 < int32(0) {
		goto L1
	} else {
		goto L17
	}
L3:
	;
	v19 = v8
	goto L5
L4:
	;
	v19 = v5
	goto L5
L5:
	;
	v26 = v8
	goto L7
L6:
	;
	v57 = v37
	goto L2
L7:
	;
	if v26 == v19 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v57 = int32(-1)
	goto L2
L10:
	;
	goto L11
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v26))))
	if int32(252) < v32 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v49 = v26 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49
	v26 = v49
	goto L7
L13:
	;
	v34 = v32 - int32(97)
	if v34 < int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v37 = int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v34)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v41)>>(uint(v34&int32(7))%32))&v37 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L12
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v61 = v60 + v57
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v72 < v61 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v115 < int32(0) {
		goto L1
	} else {
		goto L32
	}
L19:
	;
	v74 = v61
	goto L21
L20:
	;
	v74 = v72
	goto L21
L21:
	;
	v80 = v61
	goto L23
L22:
	;
	v115 = int32(1)
	goto L18
L23:
	;
	if v80 == v74 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v115 = int32(-1)
	goto L18
L26:
	;
	goto L27
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+v80))))
	if int32(252) < v89 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v91 = v89 - int32(97)
	if v91 < int32(0) {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v91)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v97)>>(uint(v91&int32(7))%32))&int32(1) == int32(0) {
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v106 = v80 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106
	v80 = v106
	goto L23
L32:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v119 = v118 + v115
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v119
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v130 < v119 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v170 < int32(0) {
		goto L1
	} else {
		goto L48
	}
L34:
	;
	v132 = v119
	goto L36
L35:
	;
	v132 = v130
	goto L36
L36:
	;
	v139 = v119
	goto L38
L37:
	;
	v170 = v150
	goto L33
L38:
	;
	if v139 == v132 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v170 = int32(-1)
	goto L33
L41:
	;
	goto L42
L42:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143+v139))))
	if int32(252) < v145 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v162 = v139 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v162
	v139 = v162
	goto L38
L44:
	;
	v147 = v145 - int32(97)
	if v147 < int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v150 = int32(1)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v147)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v154)>>(uint(v147&int32(7))%32))&v150 != 0 {
		goto L37
	} else {
		goto L46
	}
L46:
	;
	goto L43
L48:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v174 = v173 + v170
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v174
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v185 < v174 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v228 < int32(0) {
		goto L1
	} else {
		goto L63
	}
L50:
	;
	v187 = v174
	goto L52
L51:
	;
	v187 = v185
	goto L52
L52:
	;
	v193 = v174
	goto L54
L53:
	;
	v228 = int32(1)
	goto L49
L54:
	;
	if v193 == v187 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v228 = int32(-1)
	goto L49
L57:
	;
	goto L58
L58:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200+v193))))
	if int32(252) < v202 {
		goto L53
	} else {
		goto L59
	}
L59:
	;
	v204 = v202 - int32(97)
	if v204 < int32(0) {
		goto L53
	} else {
		goto L60
	}
L60:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v204)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v210)>>(uint(v204&int32(7))%32))&int32(1) == int32(0) {
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v219 = v193 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v219
	v193 = v219
	goto L54
L63:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v231 + v228
	goto L1
L64:
	;
	return v439
L65:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v274
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v274
	v280 = F_find_among_b(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_0), int32(200), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L68
	} else {
		goto L75
	}
L66:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242+v240))))
	if base.B2i32(v244&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v244)%32)&int32(_a_F_catalan_ISO_8859_1_stem_1) == int32(0)) != 0 {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v259 = F_find_among_b(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_2), int32(39), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	return int32(0)
L69:
	;
	if v259 == int32(0) {
		goto L65
	} else {
		goto L70
	}
L70:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v265
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v265 < v267 {
		goto L65
	} else {
		goto L71
	}
L71:
	;
	v269 = F_slice_del(m, l0)
	mBase = m.M
	if v269 < int32(0) {
		v439 = v269
		goto L64
	} else {
		goto L72
	}
L72:
	;
	goto L65
L73:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v349
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v349
	v355 = F_find_among_b(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_3), int32(22), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L68
	} else {
		goto L104
	}
L74:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v323
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323
	v329 = F_find_among_b(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_4), int32(283), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L68
	} else {
		goto L95
	}
L75:
	;
	if v280 == int32(0) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v284
	switch v280 - int32(1) {
	case 0:
		goto L81
	case 1:
		goto L80
	case 2:
		goto L79
	case 3:
		goto L78
	case 4:
		goto L77
	default:
		goto L73
	}
L77:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v284 < v314 {
		goto L74
	} else {
		goto L92
	}
L78:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v284 < v306 {
		goto L74
	} else {
		goto L89
	}
L79:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v284 < v298 {
		goto L74
	} else {
		goto L86
	}
L80:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v284 < v293 {
		goto L74
	} else {
		goto L84
	}
L81:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v284 < v288 {
		goto L74
	} else {
		goto L82
	}
L82:
	;
	v290 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v290 {
		goto L73
	} else {
		goto L83
	}
L83:
	;
	v439 = v290
	goto L64
L84:
	;
	v295 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v295 {
		goto L73
	} else {
		goto L85
	}
L85:
	;
	v439 = v295
	goto L64
L86:
	;
	v302 = F_slice_from_s(m, l0, int32(3), int32(_a_F_catalan_ISO_8859_1_stem_5))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L68
	} else {
		goto L87
	}
L87:
	;
	if int32(0) <= v302 {
		goto L73
	} else {
		goto L88
	}
L88:
	;
	v439 = v302
	goto L64
L89:
	;
	v310 = F_slice_from_s(m, l0, int32(2), int32(_a_F_catalan_ISO_8859_1_stem_6))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L68
	} else {
		goto L90
	}
L90:
	;
	if int32(0) <= v310 {
		goto L73
	} else {
		goto L91
	}
L91:
	;
	v439 = v310
	goto L64
L92:
	;
	v318 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_7))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L68
	} else {
		goto L93
	}
L93:
	;
	if int32(0) <= v318 {
		goto L73
	} else {
		goto L94
	}
L94:
	;
	v439 = v318
	goto L64
L95:
	;
	if v329 == int32(0) {
		goto L73
	} else {
		goto L96
	}
L96:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v333
	switch v329 - int32(1) {
	case 0:
		goto L98
	case 1:
		goto L97
	default:
		goto L73
	}
L97:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v333 < v342 {
		goto L73
	} else {
		goto L101
	}
L98:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v333 < v337 {
		goto L73
	} else {
		goto L99
	}
L99:
	;
	v339 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v339 {
		goto L73
	} else {
		goto L100
	}
L100:
	;
	v439 = v339
	goto L64
L101:
	;
	v344 = F_slice_del(m, l0)
	mBase = m.M
	if v344 < int32(0) {
		v439 = v344
		goto L64
	} else {
		goto L102
	}
L102:
	;
	goto L73
L103:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v377
	v380 = v377
	goto L113
L104:
	;
	if v355 == int32(0) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v359
	switch v355 - int32(1) {
	case 0:
		goto L107
	case 1:
		goto L106
	default:
		goto L103
	}
L106:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v359 < v368 {
		goto L103
	} else {
		goto L110
	}
L107:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v359 < v363 {
		goto L103
	} else {
		goto L108
	}
L108:
	;
	v365 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v365 {
		goto L103
	} else {
		goto L109
	}
L109:
	;
	v439 = v365
	goto L64
L110:
	;
	v372 = F_slice_from_s(m, l0, int32(2), int32(_a_F_catalan_ISO_8859_1_stem_8))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L68
	} else {
		goto L111
	}
L111:
	;
	if v372 < int32(0) {
		v439 = v372
		goto L64
	} else {
		goto L112
	}
L112:
	;
	goto L103
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v380
	v387 = F_find_among(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_9), int32(13), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L68
	} else {
		goto L116
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v377
	v439 = int32(1)
	goto L64
L115:
	;
	goto L114
L116:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v389
	switch v387 - int32(1) {
	case 0:
		goto L124
	case 1:
		goto L123
	case 2:
		goto L122
	case 3:
		goto L121
	case 4:
		goto L120
	case 5:
		goto L119
	case 6:
		goto L118
	default:
		goto L117
	}
L117:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v380 = v435
	goto L113
L118:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v429 <= v389 {
		goto L115
	} else {
		goto L137
	}
L119:
	;
	v425 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_10))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L68
	} else {
		goto L135
	}
L120:
	;
	v419 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_11))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L68
	} else {
		goto L133
	}
L121:
	;
	v413 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_12))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L68
	} else {
		goto L131
	}
L122:
	;
	v407 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_13))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L68
	} else {
		goto L129
	}
L123:
	;
	v401 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_14))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L68
	} else {
		goto L127
	}
L124:
	;
	v395 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_15))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L68
	} else {
		goto L125
	}
L125:
	;
	if int32(0) <= v395 {
		goto L117
	} else {
		goto L126
	}
L126:
	;
	v439 = v395
	goto L64
L127:
	;
	if int32(0) <= v401 {
		goto L117
	} else {
		goto L128
	}
L128:
	;
	v439 = v401
	goto L64
L129:
	;
	if int32(0) <= v407 {
		goto L117
	} else {
		goto L130
	}
L130:
	;
	v439 = v407
	goto L64
L131:
	;
	if int32(0) <= v413 {
		goto L117
	} else {
		goto L132
	}
L132:
	;
	v439 = v413
	goto L64
L133:
	;
	if int32(0) <= v419 {
		goto L117
	} else {
		goto L134
	}
L134:
	;
	v439 = v419
	goto L64
L135:
	;
	if int32(0) <= v425 {
		goto L117
	} else {
		goto L136
	}
L136:
	;
	v439 = v425
	goto L64
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v389 + int32(1)
	goto L117
}
func F_catalan_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v82 int32
	_ = v82
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v204 int32
	_ = v204
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v328 int32
	_ = v328
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v450 int32
	_ = v450
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v484 int32
	_ = v484
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v770 int32
	_ = v770
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = v9
	goto L4
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v503
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v503
	v507 = v503 - int32(1)
	if v507 <= v9 {
		goto L105
	} else {
		goto L106
	}
L2:
	;
	if v126 < int32(0) {
		goto L1
	} else {
		goto L27
	}
L3:
	;
	v126 = v98
	goto L2
L4:
	;
	if v6 <= v31 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v126 = int32(-1)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v38 = int32(1)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v23))))
	if base.Ui32(v40) < base.Ui32(int32(192)) {
		v97 = v40
		v98 = v38
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if int32(252) < v97 {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	v44 = v31 + int32(1)
	if v44 == v6 {
		v97 = v40
		v98 = v38
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v23))))
	v49 = v47 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v40) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v23))))
	v65 = v63 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v40) {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v53 = v31 + int32(2)
	if v53 != v6 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v97 = v40<<(uint(int32(6))%32)&int32(1984) | v49
	v98 = int32(2)
	goto L9
L16:
	;
	goto L15
L17:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v69))))
	v97 = v82&int32(63) | (v40<<(uint(int32(18))%32)&int32(_a_F_catalan_UTF_8_stem_0) | v49<<(uint(int32(12))%32) | v65<<(uint(int32(6))%32))
	v98 = int32(4)
	goto L9
L18:
	;
	v69 = v31 + int32(3)
	if v69 != v6 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v97 = v40<<(uint(int32(12))%32)&int32(_a_F_catalan_UTF_8_stem_1) | v49<<(uint(int32(6))%32) | v65
	v98 = int32(3)
	goto L9
L21:
	;
	goto L20
L22:
	;
	v115 = v98 + v31
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v115
	v31 = v115
	goto L4
L23:
	;
	v102 = v97 - int32(97)
	if v102 < int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v102)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_UTF_8_stem[0]))))
	if int32(base.Ui32(v108)>>(uint(v102&int32(7))%32))&int32(1) != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	goto L22
L27:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v130 = v129 + v126
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v130
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v153 = v130
	goto L30
L28:
	;
	if v249 < int32(0) {
		goto L1
	} else {
		goto L52
	}
L29:
	;
	v249 = v220
	goto L28
L30:
	;
	if v144 <= v153 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v249 = int32(-1)
	goto L28
L33:
	;
	goto L34
L34:
	;
	v160 = int32(1)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153+v145))))
	if base.Ui32(v162) < base.Ui32(int32(192)) {
		v219 = v162
		v220 = v160
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if int32(252) < v219 {
		goto L29
	} else {
		goto L48
	}
L36:
	;
	v166 = v153 + int32(1)
	if v166 == v144 {
		v219 = v162
		v220 = v160
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166+v145))))
	v171 = v169 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v162) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+v145))))
	v187 = v185 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v162) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v175 = v153 + int32(2)
	if v175 != v144 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v219 = v162<<(uint(int32(6))%32)&int32(1984) | v171
	v220 = int32(2)
	goto L35
L42:
	;
	goto L41
L43:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+v191))))
	v219 = v204&int32(63) | (v162<<(uint(int32(18))%32)&int32(_a_F_catalan_UTF_8_stem_0) | v171<<(uint(int32(12))%32) | v187<<(uint(int32(6))%32))
	v220 = int32(4)
	goto L35
L44:
	;
	v191 = v153 + int32(3)
	if v191 != v144 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v219 = v162<<(uint(int32(12))%32)&int32(_a_F_catalan_UTF_8_stem_1) | v171<<(uint(int32(6))%32) | v187
	v220 = int32(3)
	goto L35
L47:
	;
	goto L46
L48:
	;
	v224 = v219 - int32(97)
	if v224 < int32(0) {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v224)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_UTF_8_stem[0]))))
	if int32(base.Ui32(v230)>>(uint(v224&int32(7))%32))&int32(1) == int32(0) {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	v238 = v220 + v153
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v238
	v153 = v238
	goto L30
L52:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v253 = v252 + v249
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v253
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v277 = v253
	goto L55
L53:
	;
	if v372 < int32(0) {
		goto L1
	} else {
		goto L78
	}
L54:
	;
	v372 = v344
	goto L53
L55:
	;
	if v268 <= v277 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v372 = int32(-1)
	goto L53
L58:
	;
	goto L59
L59:
	;
	v284 = int32(1)
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277+v269))))
	if base.Ui32(v286) < base.Ui32(int32(192)) {
		v343 = v286
		v344 = v284
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if int32(252) < v343 {
		goto L73
	} else {
		goto L74
	}
L61:
	;
	v290 = v277 + int32(1)
	if v290 == v268 {
		v343 = v286
		v344 = v284
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290+v269))))
	v295 = v293 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v286) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299+v269))))
	v311 = v309 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v286) {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	v299 = v277 + int32(2)
	if v299 != v268 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v343 = v286<<(uint(int32(6))%32)&int32(1984) | v295
	v344 = int32(2)
	goto L60
L67:
	;
	goto L66
L68:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269+v315))))
	v343 = v328&int32(63) | (v286<<(uint(int32(18))%32)&int32(_a_F_catalan_UTF_8_stem_0) | v295<<(uint(int32(12))%32) | v311<<(uint(int32(6))%32))
	v344 = int32(4)
	goto L60
L69:
	;
	v315 = v277 + int32(3)
	if v315 != v268 {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v343 = v286<<(uint(int32(12))%32)&int32(_a_F_catalan_UTF_8_stem_1) | v295<<(uint(int32(6))%32) | v311
	v344 = int32(3)
	goto L60
L72:
	;
	goto L71
L73:
	;
	v361 = v344 + v277
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v361
	v277 = v361
	goto L55
L74:
	;
	v348 = v343 - int32(97)
	if v348 < int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v348)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_UTF_8_stem[0]))))
	if int32(base.Ui32(v354)>>(uint(v348&int32(7))%32))&int32(1) != 0 {
		goto L54
	} else {
		goto L76
	}
L76:
	;
	goto L73
L78:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v376 = v375 + v372
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v376
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v399 = v376
	goto L81
L79:
	;
	if v495 < int32(0) {
		goto L1
	} else {
		goto L103
	}
L80:
	;
	v495 = v466
	goto L79
L81:
	;
	if v390 <= v399 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v495 = int32(-1)
	goto L79
L84:
	;
	goto L85
L85:
	;
	v406 = int32(1)
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399+v391))))
	if base.Ui32(v408) < base.Ui32(int32(192)) {
		v465 = v408
		v466 = v406
		goto L86
	} else {
		goto L87
	}
L86:
	;
	if int32(252) < v465 {
		goto L80
	} else {
		goto L99
	}
L87:
	;
	v412 = v399 + int32(1)
	if v412 == v390 {
		v465 = v408
		v466 = v406
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412+v391))))
	v417 = v415 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v408) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421+v391))))
	v433 = v431 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v408) {
		goto L95
	} else {
		goto L96
	}
L90:
	;
	v421 = v399 + int32(2)
	if v421 != v390 {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v465 = v408<<(uint(int32(6))%32)&int32(1984) | v417
	v466 = int32(2)
	goto L86
L93:
	;
	goto L92
L94:
	;
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391+v437))))
	v465 = v450&int32(63) | (v408<<(uint(int32(18))%32)&int32(_a_F_catalan_UTF_8_stem_0) | v417<<(uint(int32(12))%32) | v433<<(uint(int32(6))%32))
	v466 = int32(4)
	goto L86
L95:
	;
	v437 = v399 + int32(3)
	if v437 != v390 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v465 = v408<<(uint(int32(12))%32)&int32(_a_F_catalan_UTF_8_stem_1) | v417<<(uint(int32(6))%32) | v433
	v466 = int32(3)
	goto L86
L98:
	;
	goto L97
L99:
	;
	v470 = v465 - int32(97)
	if v470 < int32(0) {
		goto L80
	} else {
		goto L100
	}
L100:
	;
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v470)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_UTF_8_stem[0]))))
	if int32(base.Ui32(v476)>>(uint(v470&int32(7))%32))&int32(1) == int32(0) {
		goto L80
	} else {
		goto L101
	}
L101:
	;
	v484 = v466 + v399
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v484
	v399 = v484
	goto L81
L103:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v498 + v495
	goto L1
L104:
	;
	return v782
L105:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v541
	v547 = F_find_among_b(m, l0, int32(_a_F_catalan_UTF_8_stem_2), int32(200), int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L108
	} else {
		goto L115
	}
L106:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509+v507))))
	if base.B2i32(v511&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v511)%32)&int32(_a_F_catalan_UTF_8_stem_3) == int32(0)) != 0 {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v526 = F_find_among_b(m, l0, int32(_a_F_catalan_UTF_8_stem_4), int32(39), int32(0))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	return int32(0)
L109:
	;
	if v526 == int32(0) {
		goto L105
	} else {
		goto L110
	}
L110:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v532
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v532 < v534 {
		goto L105
	} else {
		goto L111
	}
L111:
	;
	v536 = F_slice_del(m, l0)
	mBase = m.M
	if v536 < int32(0) {
		v782 = v536
		goto L104
	} else {
		goto L112
	}
L112:
	;
	goto L105
L113:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v616
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v616
	v622 = F_find_among_b(m, l0, int32(_a_F_catalan_UTF_8_stem_5), int32(22), int32(0))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L108
	} else {
		goto L144
	}
L114:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v590
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v590
	v596 = F_find_among_b(m, l0, int32(_a_F_catalan_UTF_8_stem_6), int32(283), int32(0))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L108
	} else {
		goto L135
	}
L115:
	;
	if v547 == int32(0) {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v551
	switch v547 - int32(1) {
	case 0:
		goto L121
	case 1:
		goto L120
	case 2:
		goto L119
	case 3:
		goto L118
	case 4:
		goto L117
	default:
		goto L113
	}
L117:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v551 < v581 {
		goto L114
	} else {
		goto L132
	}
L118:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v551 < v573 {
		goto L114
	} else {
		goto L129
	}
L119:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v551 < v565 {
		goto L114
	} else {
		goto L126
	}
L120:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v551 < v560 {
		goto L114
	} else {
		goto L124
	}
L121:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v551 < v555 {
		goto L114
	} else {
		goto L122
	}
L122:
	;
	v557 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v557 {
		goto L113
	} else {
		goto L123
	}
L123:
	;
	v782 = v557
	goto L104
L124:
	;
	v562 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v562 {
		goto L113
	} else {
		goto L125
	}
L125:
	;
	v782 = v562
	goto L104
L126:
	;
	v569 = F_slice_from_s(m, l0, int32(3), int32(_a_F_catalan_UTF_8_stem_7))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L108
	} else {
		goto L127
	}
L127:
	;
	if int32(0) <= v569 {
		goto L113
	} else {
		goto L128
	}
L128:
	;
	v782 = v569
	goto L104
L129:
	;
	v577 = F_slice_from_s(m, l0, int32(2), int32(_a_F_catalan_UTF_8_stem_8))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L108
	} else {
		goto L130
	}
L130:
	;
	if int32(0) <= v577 {
		goto L113
	} else {
		goto L131
	}
L131:
	;
	v782 = v577
	goto L104
L132:
	;
	v585 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_9))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L108
	} else {
		goto L133
	}
L133:
	;
	if int32(0) <= v585 {
		goto L113
	} else {
		goto L134
	}
L134:
	;
	v782 = v585
	goto L104
L135:
	;
	if v596 == int32(0) {
		goto L113
	} else {
		goto L136
	}
L136:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v600
	switch v596 - int32(1) {
	case 0:
		goto L138
	case 1:
		goto L137
	default:
		goto L113
	}
L137:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v600 < v609 {
		goto L113
	} else {
		goto L141
	}
L138:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v600 < v604 {
		goto L113
	} else {
		goto L139
	}
L139:
	;
	v606 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v606 {
		goto L113
	} else {
		goto L140
	}
L140:
	;
	v782 = v606
	goto L104
L141:
	;
	v611 = F_slice_del(m, l0)
	mBase = m.M
	if v611 < int32(0) {
		v782 = v611
		goto L104
	} else {
		goto L142
	}
L142:
	;
	goto L113
L143:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v644
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v648 = v644
	v651 = v646
	goto L153
L144:
	;
	if v622 == int32(0) {
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v626
	switch v622 - int32(1) {
	case 0:
		goto L147
	case 1:
		goto L146
	default:
		goto L143
	}
L146:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v626 < v635 {
		goto L143
	} else {
		goto L150
	}
L147:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v626 < v630 {
		goto L143
	} else {
		goto L148
	}
L148:
	;
	v632 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v632 {
		goto L143
	} else {
		goto L149
	}
L149:
	;
	v782 = v632
	goto L104
L150:
	;
	v639 = F_slice_from_s(m, l0, int32(2), int32(_a_F_catalan_UTF_8_stem_10))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L108
	} else {
		goto L151
	}
L151:
	;
	if v639 < int32(0) {
		v782 = v639
		goto L104
	} else {
		goto L152
	}
L152:
	;
	goto L143
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v648
	v654 = v648 + int32(1)
	if v651 <= v654 {
		goto L159
	} else {
		goto L160
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v644
	v782 = int32(1)
	goto L104
L155:
	;
	goto L154
L156:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v648 = v778
	v651 = v777
	goto L153
L157:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L185
L158:
	;
	v672 = F_find_among(m, l0, int32(_a_F_catalan_UTF_8_stem_11), int32(13), int32(0))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L108
	} else {
		goto L163
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v648
	v715 = v648
	v717 = v651
	goto L157
L160:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656+v654))))
	if v658&int32(224) != int32(160) {
		goto L159
	} else {
		goto L161
	}
L161:
	;
	if int32(1)<<(uint(v658)%32)&int32(344765187) != 0 {
		goto L158
	} else {
		goto L162
	}
L162:
	;
	goto L159
L163:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v674
	switch v672 - int32(1) {
	case 0:
		goto L170
	case 1:
		goto L169
	case 2:
		goto L168
	case 3:
		goto L167
	case 4:
		goto L166
	case 5:
		goto L165
	case 6:
		goto L164
	default:
		goto L156
	}
L164:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v715 = v674
	v717 = v714
	goto L157
L165:
	;
	v710 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_12))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L108
	} else {
		goto L181
	}
L166:
	;
	v704 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_13))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L108
	} else {
		goto L179
	}
L167:
	;
	v698 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_14))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L108
	} else {
		goto L177
	}
L168:
	;
	v692 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_15))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L108
	} else {
		goto L175
	}
L169:
	;
	v686 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_16))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L108
	} else {
		goto L173
	}
L170:
	;
	v680 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_17))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L108
	} else {
		goto L171
	}
L171:
	;
	if int32(0) <= v680 {
		goto L156
	} else {
		goto L172
	}
L172:
	;
	v782 = v680
	goto L104
L173:
	;
	if int32(0) <= v686 {
		goto L156
	} else {
		goto L174
	}
L174:
	;
	v782 = v686
	goto L104
L175:
	;
	if int32(0) <= v692 {
		goto L156
	} else {
		goto L176
	}
L176:
	;
	v782 = v692
	goto L104
L177:
	;
	if int32(0) <= v698 {
		goto L156
	} else {
		goto L178
	}
L178:
	;
	v782 = v698
	goto L104
L179:
	;
	if int32(0) <= v704 {
		goto L156
	} else {
		goto L180
	}
L180:
	;
	v782 = v704
	goto L104
L181:
	;
	if int32(0) <= v710 {
		goto L156
	} else {
		goto L182
	}
L182:
	;
	v782 = v710
	goto L104
L183:
	;
	if v770 < int32(0) {
		goto L155
	} else {
		goto L203
	}
L185:
	;
	goto L186
L186:
	;
	goto L187
L187:
	;
	v725 = v715
	v727 = int32(1)
	goto L190
L189:
	;
	v770 = v755
	goto L183
L190:
	;
	if v717 <= v725 {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	goto L189
L192:
	;
	v770 = int32(-1)
	goto L183
L193:
	;
	goto L194
L194:
	;
	v732 = v725 + int32(1)
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v718+v725))))
	if base.Ui32(v734) < base.Ui32(int32(192)) {
		v755 = v732
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v756 = int32(1)
	if v756 < v727 {
		v725 = v755
		v727 = v727 - v756
		goto L190
	} else {
		goto L202
	}
L196:
	;
	if v717 <= v732 {
		v755 = v732
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v741 = v732
	goto L198
L198:
	;
	v744 = int32(*(*int8)(unsafe.Add(mBase, uint32(v718+v741))))
	if int32(-65) < v744 {
		v755 = v741
		goto L195
	} else {
		goto L200
	}
L199:
	;
	v755 = v717
	goto L195
L200:
	;
	v748 = v741 + int32(1)
	if v748 != v717 {
		v741 = v748
		goto L198
	} else {
		goto L201
	}
L201:
	;
	goto L199
L202:
	;
	goto L191
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v770
	goto L156
}
