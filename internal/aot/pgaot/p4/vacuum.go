package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_VacuumUpdateCosts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 float64
	_ = v13
	var v17 float64
	_ = v17
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 float64
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 float64
	_ = v59
	var v64 float64
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 float64
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[0]))
	if v11 != 0 {
		v13 = *(*float64)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[1]))
		if base.F64_ge(v13, float64(0)) != 0 {
			v22 = v13
		} else {
			v17 = *(*float64)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[2]))
			if base.F64_ge(v17, float64(0)) != 0 {
				v22 = v17
			} else {
				v21 = *(*float64)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[3]))
				v22 = v21
			}
		}
		*(*float64)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[4])) = v22
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[5]))
		if int32(0) < v26 {
			v57 = v26
			v59 = v22
			*(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[6])) = v57
			v64 = v59
			v66 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[7])))
			if v66 != 0 {
			} else {
				if base.F64_gt(v64, float64(0)) != 0 {
					v70 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[8])) = uint8(v70)
				} else {
					v73 = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[9])) = v73
					*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[8])) = uint8(v73)
				}
			}
			if v11 == int32(0) {
				m.G0 = v8 + int32(32)
				return
			} else {
				v80 = int32(13)
				v87 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[10]))
				if int32(0)|base.B2i32(v87 == int32(15)) != 0 {
					v100 = int32(0)
					v104 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[11]))
					if v104 != int32(2) {
						v117 = v100
					} else {
						v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[12])))
						if v108&int32(1) != 0 {
							v117 = v100
						} else {
							v114 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[13]))
							v117 = int32(0) | base.B2i32(v114 <= v80)
						}
					}
				} else {
					if v87 <= v80 {
						v117 = int32(1)
					} else {
						v100 = int32(0)
						v104 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[11]))
						if v104 != int32(2) {
							v117 = v100
						} else {
							v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[12])))
							if v108&int32(1) != 0 {
								v117 = v100
							} else {
								v114 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[13]))
								v117 = int32(0) | base.B2i32(v114 <= v80)
							}
						}
					}
				}
				if v117 == int32(0) {
					m.G0 = v8 + int32(32)
					return
				} else {
					v122 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[14]))
					v126 = F_LWLockAcquire(m, v122+int32(2816), int32(1))
					mBase = m.M
					v127 = m.ExcPending
					if v127 != 0 {
						return
					} else {
						v129 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[0]))
						v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
						v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
						v133 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[14]))
						F_LWLockRelease(m, v133+int32(2816))
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return
						} else {
							v140 = F_errstart(m, int32(13), int32(0))
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								if v140 == int32(0) {
									m.G0 = v8 + int32(32)
									return
								} else {
									v145 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[0]))
									v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+32))
									v148 = *(*float64)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[4]))
									*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v148
									v153 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[7])))
									if v153 != 0 {
										v154 = int32(_a_F_VacuumUpdateCosts_0)
									} else {
										v154 = int32(_a_F_VacuumUpdateCosts_1)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v154
									if base.F64_gt(v148, float64(0)) != 0 {
										v160 = int32(_a_F_VacuumUpdateCosts_0)
									} else {
										v160 = int32(_a_F_VacuumUpdateCosts_1)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v160
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v131
									*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v130
									if v146 != 0 {
										v166 = int32(_a_F_VacuumUpdateCosts_0)
									} else {
										v166 = int32(_a_F_VacuumUpdateCosts_1)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v166
									v169 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[6]))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v169
									F_errmsg_internal(m, int32(_a_F_VacuumUpdateCosts_2), v8)
									mBase = m.M
									v173 = m.ExcPending
									if v173 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_VacuumUpdateCosts_3), int32(1710), int32(_a_F_VacuumUpdateCosts_4))
										mBase = m.M
										v178 = m.ExcPending
										if v178 != 0 {
											return
										} else {
											m.G0 = v8 + int32(32)
											return
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[15]))
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[16]))
			if int32(0) < v31 {
				v36 = v31
			} else {
				v36 = v33
			}
			*(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[6])) = v36
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
			if v38 == int32(0) {
				v64 = v22
				v66 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[7])))
				if v66 != 0 {
				} else {
					if base.F64_gt(v64, float64(0)) != 0 {
						v70 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[8])) = uint8(v70)
					} else {
						v73 = int32(0)
						*(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[9])) = v73
						*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[8])) = uint8(v73)
					}
				}
				if v11 == int32(0) {
					m.G0 = v8 + int32(32)
					return
				} else {
					v80 = int32(13)
					v87 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[10]))
					if int32(0)|base.B2i32(v87 == int32(15)) != 0 {
						v100 = int32(0)
						v104 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[11]))
						if v104 != int32(2) {
							v117 = v100
						} else {
							v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[12])))
							if v108&int32(1) != 0 {
								v117 = v100
							} else {
								v114 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[13]))
								v117 = int32(0) | base.B2i32(v114 <= v80)
							}
						}
					} else {
						if v87 <= v80 {
							v117 = int32(1)
						} else {
							v100 = int32(0)
							v104 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[11]))
							if v104 != int32(2) {
								v117 = v100
							} else {
								v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[12])))
								if v108&int32(1) != 0 {
									v117 = v100
								} else {
									v114 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[13]))
									v117 = int32(0) | base.B2i32(v114 <= v80)
								}
							}
						}
					}
					if v117 == int32(0) {
						m.G0 = v8 + int32(32)
						return
					} else {
						v122 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[14]))
						v126 = F_LWLockAcquire(m, v122+int32(2816), int32(1))
						mBase = m.M
						v127 = m.ExcPending
						if v127 != 0 {
							return
						} else {
							v129 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[0]))
							v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
							v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
							v133 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[14]))
							F_LWLockRelease(m, v133+int32(2816))
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return
							} else {
								v140 = F_errstart(m, int32(13), int32(0))
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return
								} else {
									if v140 == int32(0) {
										m.G0 = v8 + int32(32)
										return
									} else {
										v145 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[0]))
										v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+32))
										v148 = *(*float64)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[4]))
										*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v148
										v153 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[7])))
										if v153 != 0 {
											v154 = int32(_a_F_VacuumUpdateCosts_0)
										} else {
											v154 = int32(_a_F_VacuumUpdateCosts_1)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v154
										if base.F64_gt(v148, float64(0)) != 0 {
											v160 = int32(_a_F_VacuumUpdateCosts_0)
										} else {
											v160 = int32(_a_F_VacuumUpdateCosts_1)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v160
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v131
										*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v130
										if v146 != 0 {
											v166 = int32(_a_F_VacuumUpdateCosts_0)
										} else {
											v166 = int32(_a_F_VacuumUpdateCosts_1)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v166
										v169 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[6]))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v169
										F_errmsg_internal(m, int32(_a_F_VacuumUpdateCosts_2), v8)
										mBase = m.M
										v173 = m.ExcPending
										if v173 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_VacuumUpdateCosts_3), int32(1710), int32(_a_F_VacuumUpdateCosts_4))
											mBase = m.M
											v178 = m.ExcPending
											if v178 != 0 {
												return
											} else {
												m.G0 = v8 + int32(32)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[17]))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_VacuumUpdateCosts[18])))
				if v43 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v189 = m.ExcPending
					if v189 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_VacuumUpdateCosts_5), int32(0))
						mBase = m.M
						v193 = m.ExcPending
						if v193 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_VacuumUpdateCosts_3), int32(1754), int32(_a_F_VacuumUpdateCosts_6))
							mBase = m.M
							v198 = m.ExcPending
							if v198 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v46 = int32(1)
					v47 = base.I32_div_s(v36, v43)
					if v47 <= v46 {
						v50 = v46
					} else {
						v50 = v47
					}
					v57 = v50
					v59 = v22
					*(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[6])) = v57
					v64 = v59
					v66 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[7])))
					if v66 != 0 {
					} else {
						if base.F64_gt(v64, float64(0)) != 0 {
							v70 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[8])) = uint8(v70)
						} else {
							v73 = int32(0)
							*(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[9])) = v73
							*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[8])) = uint8(v73)
						}
					}
					if v11 == int32(0) {
						m.G0 = v8 + int32(32)
						return
					} else {
						v80 = int32(13)
						v87 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[10]))
						if int32(0)|base.B2i32(v87 == int32(15)) != 0 {
							v100 = int32(0)
							v104 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[11]))
							if v104 != int32(2) {
								v117 = v100
							} else {
								v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[12])))
								if v108&int32(1) != 0 {
									v117 = v100
								} else {
									v114 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[13]))
									v117 = int32(0) | base.B2i32(v114 <= v80)
								}
							}
						} else {
							if v87 <= v80 {
								v117 = int32(1)
							} else {
								v100 = int32(0)
								v104 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[11]))
								if v104 != int32(2) {
									v117 = v100
								} else {
									v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[12])))
									if v108&int32(1) != 0 {
										v117 = v100
									} else {
										v114 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[13]))
										v117 = int32(0) | base.B2i32(v114 <= v80)
									}
								}
							}
						}
						if v117 == int32(0) {
							m.G0 = v8 + int32(32)
							return
						} else {
							v122 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[14]))
							v126 = F_LWLockAcquire(m, v122+int32(2816), int32(1))
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
								return
							} else {
								v129 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[0]))
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
								v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
								v133 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[14]))
								F_LWLockRelease(m, v133+int32(2816))
								mBase = m.M
								v137 = m.ExcPending
								if v137 != 0 {
									return
								} else {
									v140 = F_errstart(m, int32(13), int32(0))
									mBase = m.M
									v141 = m.ExcPending
									if v141 != 0 {
										return
									} else {
										if v140 == int32(0) {
											m.G0 = v8 + int32(32)
											return
										} else {
											v145 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[0]))
											v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+32))
											v148 = *(*float64)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[4]))
											*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v148
											v153 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[7])))
											if v153 != 0 {
												v154 = int32(_a_F_VacuumUpdateCosts_0)
											} else {
												v154 = int32(_a_F_VacuumUpdateCosts_1)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v154
											if base.F64_gt(v148, float64(0)) != 0 {
												v160 = int32(_a_F_VacuumUpdateCosts_0)
											} else {
												v160 = int32(_a_F_VacuumUpdateCosts_1)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v160
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v131
											*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v130
											if v146 != 0 {
												v166 = int32(_a_F_VacuumUpdateCosts_0)
											} else {
												v166 = int32(_a_F_VacuumUpdateCosts_1)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v166
											v169 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[6]))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v169
											F_errmsg_internal(m, int32(_a_F_VacuumUpdateCosts_2), v8)
											mBase = m.M
											v173 = m.ExcPending
											if v173 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_VacuumUpdateCosts_3), int32(1710), int32(_a_F_VacuumUpdateCosts_4))
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return
												} else {
													m.G0 = v8 + int32(32)
													return
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
		v53 = *(*float64)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[3]))
		*(*float64)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[4])) = v53
		v56 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[16]))
		v57 = v56
		v59 = v53
		*(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[6])) = v57
		v64 = v59
		v66 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[7])))
		if v66 != 0 {
		} else {
			if base.F64_gt(v64, float64(0)) != 0 {
				v70 = int32(1)
				*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[8])) = uint8(v70)
			} else {
				v73 = int32(0)
				*(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[9])) = v73
				*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[8])) = uint8(v73)
			}
		}
		if v11 == int32(0) {
			m.G0 = v8 + int32(32)
			return
		} else {
			v80 = int32(13)
			v87 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[10]))
			if int32(0)|base.B2i32(v87 == int32(15)) != 0 {
				v100 = int32(0)
				v104 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[11]))
				if v104 != int32(2) {
					v117 = v100
				} else {
					v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[12])))
					if v108&int32(1) != 0 {
						v117 = v100
					} else {
						v114 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[13]))
						v117 = int32(0) | base.B2i32(v114 <= v80)
					}
				}
			} else {
				if v87 <= v80 {
					v117 = int32(1)
				} else {
					v100 = int32(0)
					v104 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[11]))
					if v104 != int32(2) {
						v117 = v100
					} else {
						v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[12])))
						if v108&int32(1) != 0 {
							v117 = v100
						} else {
							v114 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[13]))
							v117 = int32(0) | base.B2i32(v114 <= v80)
						}
					}
				}
			}
			if v117 == int32(0) {
				m.G0 = v8 + int32(32)
				return
			} else {
				v122 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[14]))
				v126 = F_LWLockAcquire(m, v122+int32(2816), int32(1))
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return
				} else {
					v129 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[0]))
					v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
					v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
					v133 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[14]))
					F_LWLockRelease(m, v133+int32(2816))
					mBase = m.M
					v137 = m.ExcPending
					if v137 != 0 {
						return
					} else {
						v140 = F_errstart(m, int32(13), int32(0))
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return
						} else {
							if v140 == int32(0) {
								m.G0 = v8 + int32(32)
								return
							} else {
								v145 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[0]))
								v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+32))
								v148 = *(*float64)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[4]))
								*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v148
								v153 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[7])))
								if v153 != 0 {
									v154 = int32(_a_F_VacuumUpdateCosts_0)
								} else {
									v154 = int32(_a_F_VacuumUpdateCosts_1)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v154
								if base.F64_gt(v148, float64(0)) != 0 {
									v160 = int32(_a_F_VacuumUpdateCosts_0)
								} else {
									v160 = int32(_a_F_VacuumUpdateCosts_1)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v160
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v131
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v130
								if v146 != 0 {
									v166 = int32(_a_F_VacuumUpdateCosts_0)
								} else {
									v166 = int32(_a_F_VacuumUpdateCosts_1)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v166
								v169 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[6]))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v169
								F_errmsg_internal(m, int32(_a_F_VacuumUpdateCosts_2), v8)
								mBase = m.M
								v173 = m.ExcPending
								if v173 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_VacuumUpdateCosts_3), int32(1710), int32(_a_F_VacuumUpdateCosts_4))
									mBase = m.M
									v178 = m.ExcPending
									if v178 != 0 {
										return
									} else {
										m.G0 = v8 + int32(32)
										return
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
func F_vacuum_is_permitted_for_relation(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_is_permitted_for_relation[0]))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_is_permitted_for_relation[1]))
	v15 = F_object_ownercheck(m, int32(1262), v12, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 != 0 {
			v19 = int32(1)
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+117)))
			if v20 != v19 {
				v69 = v19
				m.G0 = v8 + int32(16)
				return v69
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_is_permitted_for_relation[1]))
				v27 = F_pg_class_aclcheck(m, l0, v25, int64(16384))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if v27 == int32(0) {
						v69 = int32(1)
						m.G0 = v8 + int32(16)
						return v69
					} else {
						if l2&int32(1) != 0 {
							v34 = int32(0)
							v37 = F_errstart(m, int32(19), v34)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								if v37 == int32(0) {
									v69 = v34
									m.G0 = v8 + int32(16)
									return v69
								} else {
									v56 = int32(760)
									v57 = int32(_a_F_vacuum_is_permitted_for_relation_0)
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1 + int32(4)
									F_errmsg(m, v57, v8)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_vacuum_is_permitted_for_relation_1), v56, int32(_a_F_vacuum_is_permitted_for_relation_2))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											v69 = int32(0)
											m.G0 = v8 + int32(16)
											return v69
										}
									}
								}
							}
						} else {
							if l2&int32(2) == int32(0) {
								v69 = int32(0)
								m.G0 = v8 + int32(16)
								return v69
							} else {
								v47 = int32(0)
								v50 = F_errstart(m, int32(19), v47)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									if v50 == int32(0) {
										v69 = v47
										m.G0 = v8 + int32(16)
										return v69
									} else {
										v56 = int32(773)
										v57 = int32(_a_F_vacuum_is_permitted_for_relation_3)
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1 + int32(4)
										F_errmsg(m, v57, v8)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_vacuum_is_permitted_for_relation_1), v56, int32(_a_F_vacuum_is_permitted_for_relation_2))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v69 = int32(0)
												m.G0 = v8 + int32(16)
												return v69
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
			v25 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_is_permitted_for_relation[1]))
			v27 = F_pg_class_aclcheck(m, l0, v25, int64(16384))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v27 == int32(0) {
					v69 = int32(1)
					m.G0 = v8 + int32(16)
					return v69
				} else {
					if l2&int32(1) != 0 {
						v34 = int32(0)
						v37 = F_errstart(m, int32(19), v34)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							if v37 == int32(0) {
								v69 = v34
								m.G0 = v8 + int32(16)
								return v69
							} else {
								v56 = int32(760)
								v57 = int32(_a_F_vacuum_is_permitted_for_relation_0)
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1 + int32(4)
								F_errmsg(m, v57, v8)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_vacuum_is_permitted_for_relation_1), v56, int32(_a_F_vacuum_is_permitted_for_relation_2))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										v69 = int32(0)
										m.G0 = v8 + int32(16)
										return v69
									}
								}
							}
						}
					} else {
						if l2&int32(2) == int32(0) {
							v69 = int32(0)
							m.G0 = v8 + int32(16)
							return v69
						} else {
							v47 = int32(0)
							v50 = F_errstart(m, int32(19), v47)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								if v50 == int32(0) {
									v69 = v47
									m.G0 = v8 + int32(16)
									return v69
								} else {
									v56 = int32(773)
									v57 = int32(_a_F_vacuum_is_permitted_for_relation_3)
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1 + int32(4)
									F_errmsg(m, v57, v8)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_vacuum_is_permitted_for_relation_1), v56, int32(_a_F_vacuum_is_permitted_for_relation_2))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											v69 = int32(0)
											m.G0 = v8 + int32(16)
											return v69
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
