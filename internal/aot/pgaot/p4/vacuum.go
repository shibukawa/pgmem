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
	var v90 int32
	_ = v90
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 float64
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
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
				v90 = *(*int32)(unsafe.Add(mBase, uint32(v87<<(uint(int32(2))%32))+uint32(_c_F_VacuumUpdateCosts[11])))
				if int32(0)|base.B2i32(v90 == int32(15)) != 0 {
					v103 = int32(0)
					v107 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[12]))
					if v107 != int32(2) {
						v120 = v103
					} else {
						v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[13])))
						if v111&int32(1) != 0 {
							v120 = v103
						} else {
							v117 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[14]))
							v120 = int32(0) | base.B2i32(v117 <= v80)
						}
					}
				} else {
					if v90 <= v80 {
						v120 = int32(1)
					} else {
						v103 = int32(0)
						v107 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[12]))
						if v107 != int32(2) {
							v120 = v103
						} else {
							v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[13])))
							if v111&int32(1) != 0 {
								v120 = v103
							} else {
								v117 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[14]))
								v120 = int32(0) | base.B2i32(v117 <= v80)
							}
						}
					}
				}
				if v120 == int32(0) {
					m.G0 = v8 + int32(32)
					return
				} else {
					v125 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[15]))
					v129 = F_LWLockAcquire(m, v125+int32(2816), int32(1))
					mBase = m.M
					v130 = m.ExcPending
					if v130 != 0 {
						return
					} else {
						v132 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[0]))
						v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
						v134 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
						v136 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[15]))
						F_LWLockRelease(m, v136+int32(2816))
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return
						} else {
							v143 = F_errstart(m, int32(13), int32(0))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								if v143 == int32(0) {
									m.G0 = v8 + int32(32)
									return
								} else {
									v148 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[0]))
									v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+32))
									v151 = *(*float64)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[4]))
									*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v151
									v156 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[7])))
									if v156 != 0 {
										v157 = int32(_a_F_VacuumUpdateCosts_0)
									} else {
										v157 = int32(_a_F_VacuumUpdateCosts_1)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v157
									if base.F64_gt(v151, float64(0)) != 0 {
										v163 = int32(_a_F_VacuumUpdateCosts_0)
									} else {
										v163 = int32(_a_F_VacuumUpdateCosts_1)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v163
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v134
									*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v133
									if v149 != 0 {
										v169 = int32(_a_F_VacuumUpdateCosts_0)
									} else {
										v169 = int32(_a_F_VacuumUpdateCosts_1)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v169
									v172 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[6]))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v172
									F_errmsg_internal(m, int32(_a_F_VacuumUpdateCosts_2), v8)
									mBase = m.M
									v176 = m.ExcPending
									if v176 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_VacuumUpdateCosts_3), int32(1738), int32(_a_F_VacuumUpdateCosts_4))
										mBase = m.M
										v181 = m.ExcPending
										if v181 != 0 {
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
			v31 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[16]))
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[17]))
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
					v90 = *(*int32)(unsafe.Add(mBase, uint32(v87<<(uint(int32(2))%32))+uint32(_c_F_VacuumUpdateCosts[11])))
					if int32(0)|base.B2i32(v90 == int32(15)) != 0 {
						v103 = int32(0)
						v107 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[12]))
						if v107 != int32(2) {
							v120 = v103
						} else {
							v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[13])))
							if v111&int32(1) != 0 {
								v120 = v103
							} else {
								v117 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[14]))
								v120 = int32(0) | base.B2i32(v117 <= v80)
							}
						}
					} else {
						if v90 <= v80 {
							v120 = int32(1)
						} else {
							v103 = int32(0)
							v107 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[12]))
							if v107 != int32(2) {
								v120 = v103
							} else {
								v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[13])))
								if v111&int32(1) != 0 {
									v120 = v103
								} else {
									v117 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[14]))
									v120 = int32(0) | base.B2i32(v117 <= v80)
								}
							}
						}
					}
					if v120 == int32(0) {
						m.G0 = v8 + int32(32)
						return
					} else {
						v125 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[15]))
						v129 = F_LWLockAcquire(m, v125+int32(2816), int32(1))
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return
						} else {
							v132 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[0]))
							v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
							v134 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
							v136 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[15]))
							F_LWLockRelease(m, v136+int32(2816))
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return
							} else {
								v143 = F_errstart(m, int32(13), int32(0))
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return
								} else {
									if v143 == int32(0) {
										m.G0 = v8 + int32(32)
										return
									} else {
										v148 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[0]))
										v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+32))
										v151 = *(*float64)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[4]))
										*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v151
										v156 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[7])))
										if v156 != 0 {
											v157 = int32(_a_F_VacuumUpdateCosts_0)
										} else {
											v157 = int32(_a_F_VacuumUpdateCosts_1)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v157
										if base.F64_gt(v151, float64(0)) != 0 {
											v163 = int32(_a_F_VacuumUpdateCosts_0)
										} else {
											v163 = int32(_a_F_VacuumUpdateCosts_1)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v163
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v134
										*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v133
										if v149 != 0 {
											v169 = int32(_a_F_VacuumUpdateCosts_0)
										} else {
											v169 = int32(_a_F_VacuumUpdateCosts_1)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v169
										v172 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[6]))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v172
										F_errmsg_internal(m, int32(_a_F_VacuumUpdateCosts_2), v8)
										mBase = m.M
										v176 = m.ExcPending
										if v176 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_VacuumUpdateCosts_3), int32(1738), int32(_a_F_VacuumUpdateCosts_4))
											mBase = m.M
											v181 = m.ExcPending
											if v181 != 0 {
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
				v42 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[18]))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_VacuumUpdateCosts[19])))
				if v43 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v192 = m.ExcPending
					if v192 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_VacuumUpdateCosts_5), int32(0))
						mBase = m.M
						v196 = m.ExcPending
						if v196 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_VacuumUpdateCosts_3), int32(1782), int32(_a_F_VacuumUpdateCosts_6))
							mBase = m.M
							v201 = m.ExcPending
							if v201 != 0 {
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
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v87<<(uint(int32(2))%32))+uint32(_c_F_VacuumUpdateCosts[11])))
						if int32(0)|base.B2i32(v90 == int32(15)) != 0 {
							v103 = int32(0)
							v107 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[12]))
							if v107 != int32(2) {
								v120 = v103
							} else {
								v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[13])))
								if v111&int32(1) != 0 {
									v120 = v103
								} else {
									v117 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[14]))
									v120 = int32(0) | base.B2i32(v117 <= v80)
								}
							}
						} else {
							if v90 <= v80 {
								v120 = int32(1)
							} else {
								v103 = int32(0)
								v107 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[12]))
								if v107 != int32(2) {
									v120 = v103
								} else {
									v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[13])))
									if v111&int32(1) != 0 {
										v120 = v103
									} else {
										v117 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[14]))
										v120 = int32(0) | base.B2i32(v117 <= v80)
									}
								}
							}
						}
						if v120 == int32(0) {
							m.G0 = v8 + int32(32)
							return
						} else {
							v125 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[15]))
							v129 = F_LWLockAcquire(m, v125+int32(2816), int32(1))
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return
							} else {
								v132 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[0]))
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
								v136 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[15]))
								F_LWLockRelease(m, v136+int32(2816))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return
								} else {
									v143 = F_errstart(m, int32(13), int32(0))
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return
									} else {
										if v143 == int32(0) {
											m.G0 = v8 + int32(32)
											return
										} else {
											v148 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[0]))
											v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+32))
											v151 = *(*float64)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[4]))
											*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v151
											v156 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[7])))
											if v156 != 0 {
												v157 = int32(_a_F_VacuumUpdateCosts_0)
											} else {
												v157 = int32(_a_F_VacuumUpdateCosts_1)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v157
											if base.F64_gt(v151, float64(0)) != 0 {
												v163 = int32(_a_F_VacuumUpdateCosts_0)
											} else {
												v163 = int32(_a_F_VacuumUpdateCosts_1)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v163
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v134
											*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v133
											if v149 != 0 {
												v169 = int32(_a_F_VacuumUpdateCosts_0)
											} else {
												v169 = int32(_a_F_VacuumUpdateCosts_1)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v169
											v172 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[6]))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v172
											F_errmsg_internal(m, int32(_a_F_VacuumUpdateCosts_2), v8)
											mBase = m.M
											v176 = m.ExcPending
											if v176 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_VacuumUpdateCosts_3), int32(1738), int32(_a_F_VacuumUpdateCosts_4))
												mBase = m.M
												v181 = m.ExcPending
												if v181 != 0 {
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
		v56 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[17]))
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
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v87<<(uint(int32(2))%32))+uint32(_c_F_VacuumUpdateCosts[11])))
			if int32(0)|base.B2i32(v90 == int32(15)) != 0 {
				v103 = int32(0)
				v107 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[12]))
				if v107 != int32(2) {
					v120 = v103
				} else {
					v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[13])))
					if v111&int32(1) != 0 {
						v120 = v103
					} else {
						v117 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[14]))
						v120 = int32(0) | base.B2i32(v117 <= v80)
					}
				}
			} else {
				if v90 <= v80 {
					v120 = int32(1)
				} else {
					v103 = int32(0)
					v107 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[12]))
					if v107 != int32(2) {
						v120 = v103
					} else {
						v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[13])))
						if v111&int32(1) != 0 {
							v120 = v103
						} else {
							v117 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[14]))
							v120 = int32(0) | base.B2i32(v117 <= v80)
						}
					}
				}
			}
			if v120 == int32(0) {
				m.G0 = v8 + int32(32)
				return
			} else {
				v125 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[15]))
				v129 = F_LWLockAcquire(m, v125+int32(2816), int32(1))
				mBase = m.M
				v130 = m.ExcPending
				if v130 != 0 {
					return
				} else {
					v132 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[0]))
					v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
					v134 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
					v136 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[15]))
					F_LWLockRelease(m, v136+int32(2816))
					mBase = m.M
					v140 = m.ExcPending
					if v140 != 0 {
						return
					} else {
						v143 = F_errstart(m, int32(13), int32(0))
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return
						} else {
							if v143 == int32(0) {
								m.G0 = v8 + int32(32)
								return
							} else {
								v148 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[0]))
								v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+32))
								v151 = *(*float64)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[4]))
								*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v151
								v156 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[7])))
								if v156 != 0 {
									v157 = int32(_a_F_VacuumUpdateCosts_0)
								} else {
									v157 = int32(_a_F_VacuumUpdateCosts_1)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v157
								if base.F64_gt(v151, float64(0)) != 0 {
									v163 = int32(_a_F_VacuumUpdateCosts_0)
								} else {
									v163 = int32(_a_F_VacuumUpdateCosts_1)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v163
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v134
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v133
								if v149 != 0 {
									v169 = int32(_a_F_VacuumUpdateCosts_0)
								} else {
									v169 = int32(_a_F_VacuumUpdateCosts_1)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v169
								v172 = *(*int32)(unsafe.Add(mBase, _c_F_VacuumUpdateCosts[6]))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v172
								F_errmsg_internal(m, int32(_a_F_VacuumUpdateCosts_2), v8)
								mBase = m.M
								v176 = m.ExcPending
								if v176 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_VacuumUpdateCosts_3), int32(1738), int32(_a_F_VacuumUpdateCosts_4))
									mBase = m.M
									v181 = m.ExcPending
									if v181 != 0 {
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
									v56 = int32(745)
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
										v56 = int32(758)
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
								v56 = int32(745)
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
									v56 = int32(758)
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
