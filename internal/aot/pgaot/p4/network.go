package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_network_abbrev_abort(m *base.Module, l0 int32, l1 int32) int32 {
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v11 = Fn14332(m, l0, l1, int32(_a_F_network_abbrev_abort_0), int32(531), int32(_a_F_network_abbrev_abort_1), int32(_a_F_network_abbrev_abort_2), int32(524), int32(_a_F_network_abbrev_abort_3), int32(506), int32(_a_F_network_abbrev_abort_4))
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
func F_network_fast_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	v5 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(l1))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v19 = int32(1)
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
			if v21&v19 != 0 {
				v24 = v19
			} else {
				v24 = int32(4)
			}
			v25 = v5 + v24
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
			v27 = int32(1)
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			if v29&v27 != 0 {
				v32 = v27
			} else {
				v32 = int32(4)
			}
			v33 = v10 + v32
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
			if v26 == v34 {
				v36 = int32(2)
				v37 = v25 + v36
				v39 = v33 + v36
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
				if base.Ui32(v40) < base.Ui32(v41) {
					v43 = v25
				} else {
					v43 = v33
				}
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
				v46 = int32(base.Ui32(v44) >> (uint(int32(3)) % 32))
				v47 = F_memcmp(m, v37, v39, v46)
				mBase = m.M
				if v47 != 0 {
					v139 = v47
					v158 = v139
				} else {
					v49 = v44 & int32(7)
					if v49 == int32(0) {
						v130 = v40 - v41
						if v130 != 0 {
							v139 = v130
							v158 = v139
						} else {
							if v26 == int32(2) {
								v135 = int32(4)
							} else {
								v135 = int32(16)
							}
							v136 = F_memcmp(m, v37, v39, v135)
							mBase = m.M
							v158 = v136
						}
					} else {
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v37))))
						v54 = int32(128)
						v55 = v53 & v54
						v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v39))))
						if v55 != v57&v54 {
							v146 = v55
							if v146 != 0 {
								v149 = int32(1)
							} else {
								v149 = int32(-1)
							}
							v158 = v149
						} else {
							if v49 == int32(1) {
								v130 = v40 - v41
								if v130 != 0 {
									v139 = v130
									v158 = v139
								} else {
									if v26 == int32(2) {
										v135 = int32(4)
									} else {
										v135 = int32(16)
									}
									v136 = F_memcmp(m, v37, v39, v135)
									mBase = m.M
									v158 = v136
								}
							} else {
								v63 = int32(1)
								v65 = int32(128)
								v66 = v53 << (uint(v63) % 32) & v65
								if v66 != v57<<(uint(v63)%32)&v65 {
									v146 = v66
									if v146 != 0 {
										v149 = int32(1)
									} else {
										v149 = int32(-1)
									}
									v158 = v149
								} else {
									if base.Ui32(v49) < base.Ui32(int32(3)) {
										v130 = v40 - v41
										if v130 != 0 {
											v139 = v130
											v158 = v139
										} else {
											if v26 == int32(2) {
												v135 = int32(4)
											} else {
												v135 = int32(16)
											}
											v136 = F_memcmp(m, v37, v39, v135)
											mBase = m.M
											v158 = v136
										}
									} else {
										v74 = int32(2)
										v76 = int32(128)
										v77 = v53 << (uint(v74) % 32) & v76
										if v77 != v57<<(uint(v74)%32)&v76 {
											v146 = v77
											if v146 != 0 {
												v149 = int32(1)
											} else {
												v149 = int32(-1)
											}
											v158 = v149
										} else {
											if v49 == int32(3) {
												v130 = v40 - v41
												if v130 != 0 {
													v139 = v130
													v158 = v139
												} else {
													if v26 == int32(2) {
														v135 = int32(4)
													} else {
														v135 = int32(16)
													}
													v136 = F_memcmp(m, v37, v39, v135)
													mBase = m.M
													v158 = v136
												}
											} else {
												v85 = int32(3)
												v87 = int32(128)
												v88 = v53 << (uint(v85) % 32) & v87
												if v88 != v57<<(uint(v85)%32)&v87 {
													v146 = v88
													if v146 != 0 {
														v149 = int32(1)
													} else {
														v149 = int32(-1)
													}
													v158 = v149
												} else {
													if base.Ui32(v49) < base.Ui32(int32(5)) {
														v130 = v40 - v41
														if v130 != 0 {
															v139 = v130
															v158 = v139
														} else {
															if v26 == int32(2) {
																v135 = int32(4)
															} else {
																v135 = int32(16)
															}
															v136 = F_memcmp(m, v37, v39, v135)
															mBase = m.M
															v158 = v136
														}
													} else {
														v96 = int32(4)
														v98 = int32(128)
														v99 = v53 << (uint(v96) % 32) & v98
														if v99 != v57<<(uint(v96)%32)&v98 {
															v146 = v99
															if v146 != 0 {
																v149 = int32(1)
															} else {
																v149 = int32(-1)
															}
															v158 = v149
														} else {
															if v49 == int32(5) {
																v130 = v40 - v41
																if v130 != 0 {
																	v139 = v130
																	v158 = v139
																} else {
																	if v26 == int32(2) {
																		v135 = int32(4)
																	} else {
																		v135 = int32(16)
																	}
																	v136 = F_memcmp(m, v37, v39, v135)
																	mBase = m.M
																	v158 = v136
																}
															} else {
																v107 = int32(5)
																v109 = int32(128)
																v110 = v53 << (uint(v107) % 32) & v109
																if v110 != v57<<(uint(v107)%32)&v109 {
																	v146 = v110
																	if v146 != 0 {
																		v149 = int32(1)
																	} else {
																		v149 = int32(-1)
																	}
																	v158 = v149
																} else {
																	if v49 != int32(7) {
																		v130 = v40 - v41
																		if v130 != 0 {
																			v139 = v130
																			v158 = v139
																		} else {
																			if v26 == int32(2) {
																				v135 = int32(4)
																			} else {
																				v135 = int32(16)
																			}
																			v136 = F_memcmp(m, v37, v39, v135)
																			mBase = m.M
																			v158 = v136
																		}
																	} else {
																		v118 = int32(6)
																		v120 = int32(128)
																		v121 = v53 << (uint(v118) % 32) & v120
																		if v121 != v57<<(uint(v118)%32)&v120 {
																			v146 = v121
																			if v146 != 0 {
																				v149 = int32(1)
																			} else {
																				v149 = int32(-1)
																			}
																			v158 = v149
																		} else {
																			v130 = v40 - v41
																			if v130 != 0 {
																				v139 = v130
																				v158 = v139
																			} else {
																				if v26 == int32(2) {
																					v135 = int32(4)
																				} else {
																					v135 = int32(16)
																				}
																				v136 = F_memcmp(m, v37, v39, v135)
																				mBase = m.M
																				v158 = v136
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
								}
							}
						}
					}
				}
			} else {
				v139 = v26 - v34
				v158 = v139
			}
			return v158
		}
	}
}
func F_network_gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v17 = int32(1)
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
			if v19&v17 != 0 {
				v22 = v17
			} else {
				v22 = int32(4)
			}
			v23 = v3 + v22
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
			v25 = int32(1)
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v27&v25 != 0 {
				v30 = v25
			} else {
				v30 = int32(4)
			}
			v31 = v8 + v30
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
			if v24 == v32 {
				v34 = int32(2)
				v35 = v23 + v34
				v37 = v31 + v34
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
				if base.Ui32(v38) < base.Ui32(v39) {
					v41 = v23
				} else {
					v41 = v31
				}
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
				v44 = int32(base.Ui32(v42) >> (uint(int32(3)) % 32))
				v45 = F_memcmp(m, v35, v37, v44)
				mBase = m.M
				if v45 != 0 {
					v137 = v45
					v156 = v137
				} else {
					v47 = v42 & int32(7)
					if v47 == int32(0) {
						v128 = v38 - v39
						if v128 != 0 {
							v137 = v128
							v156 = v137
						} else {
							if v24 == int32(2) {
								v133 = int32(4)
							} else {
								v133 = int32(16)
							}
							v134 = F_memcmp(m, v35, v37, v133)
							mBase = m.M
							v156 = v134
						}
					} else {
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v35))))
						v52 = int32(128)
						v53 = v51 & v52
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v37))))
						if v53 != v55&v52 {
							v144 = v53
							if v144 != 0 {
								v147 = int32(1)
							} else {
								v147 = int32(-1)
							}
							v156 = v147
						} else {
							if v47 == int32(1) {
								v128 = v38 - v39
								if v128 != 0 {
									v137 = v128
									v156 = v137
								} else {
									if v24 == int32(2) {
										v133 = int32(4)
									} else {
										v133 = int32(16)
									}
									v134 = F_memcmp(m, v35, v37, v133)
									mBase = m.M
									v156 = v134
								}
							} else {
								v61 = int32(1)
								v63 = int32(128)
								v64 = v51 << (uint(v61) % 32) & v63
								if v64 != v55<<(uint(v61)%32)&v63 {
									v144 = v64
									if v144 != 0 {
										v147 = int32(1)
									} else {
										v147 = int32(-1)
									}
									v156 = v147
								} else {
									if base.Ui32(v47) < base.Ui32(int32(3)) {
										v128 = v38 - v39
										if v128 != 0 {
											v137 = v128
											v156 = v137
										} else {
											if v24 == int32(2) {
												v133 = int32(4)
											} else {
												v133 = int32(16)
											}
											v134 = F_memcmp(m, v35, v37, v133)
											mBase = m.M
											v156 = v134
										}
									} else {
										v72 = int32(2)
										v74 = int32(128)
										v75 = v51 << (uint(v72) % 32) & v74
										if v75 != v55<<(uint(v72)%32)&v74 {
											v144 = v75
											if v144 != 0 {
												v147 = int32(1)
											} else {
												v147 = int32(-1)
											}
											v156 = v147
										} else {
											if v47 == int32(3) {
												v128 = v38 - v39
												if v128 != 0 {
													v137 = v128
													v156 = v137
												} else {
													if v24 == int32(2) {
														v133 = int32(4)
													} else {
														v133 = int32(16)
													}
													v134 = F_memcmp(m, v35, v37, v133)
													mBase = m.M
													v156 = v134
												}
											} else {
												v83 = int32(3)
												v85 = int32(128)
												v86 = v51 << (uint(v83) % 32) & v85
												if v86 != v55<<(uint(v83)%32)&v85 {
													v144 = v86
													if v144 != 0 {
														v147 = int32(1)
													} else {
														v147 = int32(-1)
													}
													v156 = v147
												} else {
													if base.Ui32(v47) < base.Ui32(int32(5)) {
														v128 = v38 - v39
														if v128 != 0 {
															v137 = v128
															v156 = v137
														} else {
															if v24 == int32(2) {
																v133 = int32(4)
															} else {
																v133 = int32(16)
															}
															v134 = F_memcmp(m, v35, v37, v133)
															mBase = m.M
															v156 = v134
														}
													} else {
														v94 = int32(4)
														v96 = int32(128)
														v97 = v51 << (uint(v94) % 32) & v96
														if v97 != v55<<(uint(v94)%32)&v96 {
															v144 = v97
															if v144 != 0 {
																v147 = int32(1)
															} else {
																v147 = int32(-1)
															}
															v156 = v147
														} else {
															if v47 == int32(5) {
																v128 = v38 - v39
																if v128 != 0 {
																	v137 = v128
																	v156 = v137
																} else {
																	if v24 == int32(2) {
																		v133 = int32(4)
																	} else {
																		v133 = int32(16)
																	}
																	v134 = F_memcmp(m, v35, v37, v133)
																	mBase = m.M
																	v156 = v134
																}
															} else {
																v105 = int32(5)
																v107 = int32(128)
																v108 = v51 << (uint(v105) % 32) & v107
																if v108 != v55<<(uint(v105)%32)&v107 {
																	v144 = v108
																	if v144 != 0 {
																		v147 = int32(1)
																	} else {
																		v147 = int32(-1)
																	}
																	v156 = v147
																} else {
																	if v47 != int32(7) {
																		v128 = v38 - v39
																		if v128 != 0 {
																			v137 = v128
																			v156 = v137
																		} else {
																			if v24 == int32(2) {
																				v133 = int32(4)
																			} else {
																				v133 = int32(16)
																			}
																			v134 = F_memcmp(m, v35, v37, v133)
																			mBase = m.M
																			v156 = v134
																		}
																	} else {
																		v116 = int32(6)
																		v118 = int32(128)
																		v119 = v51 << (uint(v116) % 32) & v118
																		if v119 != v55<<(uint(v116)%32)&v118 {
																			v144 = v119
																			if v144 != 0 {
																				v147 = int32(1)
																			} else {
																				v147 = int32(-1)
																			}
																			v156 = v147
																		} else {
																			v128 = v38 - v39
																			if v128 != 0 {
																				v137 = v128
																				v156 = v137
																			} else {
																				if v24 == int32(2) {
																					v133 = int32(4)
																				} else {
																					v133 = int32(16)
																				}
																				v134 = F_memcmp(m, v35, v37, v133)
																				mBase = m.M
																				v156 = v134
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
								}
							}
						}
					}
				}
			} else {
				v137 = v24 - v32
				v156 = v137
			}
			return base.I64_extend_i32_u(base.B2i32(int32(0) < v156))
		}
	}
}
func F_network_larger(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v9 = F_pg_detoast_datum_packed(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			v18 = int32(1)
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
			if v20&v18 != 0 {
				v23 = v18
			} else {
				v23 = int32(4)
			}
			v24 = v4 + v23
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
			v26 = int32(1)
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v28&v26 != 0 {
				v31 = v26
			} else {
				v31 = int32(4)
			}
			v32 = v9 + v31
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
			if v25 == v33 {
				v35 = int32(2)
				v36 = v24 + v35
				v38 = v32 + v35
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
				if base.Ui32(v39) < base.Ui32(v40) {
					v42 = v24
				} else {
					v42 = v32
				}
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
				v45 = int32(base.Ui32(v43) >> (uint(int32(3)) % 32))
				v46 = F_memcmp(m, v36, v38, v45)
				mBase = m.M
				if v46 != 0 {
					v138 = v46
					v157 = v138
				} else {
					v48 = v43 & int32(7)
					if v48 == int32(0) {
						v129 = v39 - v40
						if v129 != 0 {
							v138 = v129
							v157 = v138
						} else {
							if v25 == int32(2) {
								v134 = int32(4)
							} else {
								v134 = int32(16)
							}
							v135 = F_memcmp(m, v36, v38, v134)
							mBase = m.M
							v157 = v135
						}
					} else {
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v36))))
						v53 = int32(128)
						v54 = v52 & v53
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v38))))
						if v54 != v56&v53 {
							v145 = v54
							if v145 != 0 {
								v148 = int32(1)
							} else {
								v148 = int32(-1)
							}
							v157 = v148
						} else {
							if v48 == int32(1) {
								v129 = v39 - v40
								if v129 != 0 {
									v138 = v129
									v157 = v138
								} else {
									if v25 == int32(2) {
										v134 = int32(4)
									} else {
										v134 = int32(16)
									}
									v135 = F_memcmp(m, v36, v38, v134)
									mBase = m.M
									v157 = v135
								}
							} else {
								v62 = int32(1)
								v64 = int32(128)
								v65 = v52 << (uint(v62) % 32) & v64
								if v65 != v56<<(uint(v62)%32)&v64 {
									v145 = v65
									if v145 != 0 {
										v148 = int32(1)
									} else {
										v148 = int32(-1)
									}
									v157 = v148
								} else {
									if base.Ui32(v48) < base.Ui32(int32(3)) {
										v129 = v39 - v40
										if v129 != 0 {
											v138 = v129
											v157 = v138
										} else {
											if v25 == int32(2) {
												v134 = int32(4)
											} else {
												v134 = int32(16)
											}
											v135 = F_memcmp(m, v36, v38, v134)
											mBase = m.M
											v157 = v135
										}
									} else {
										v73 = int32(2)
										v75 = int32(128)
										v76 = v52 << (uint(v73) % 32) & v75
										if v76 != v56<<(uint(v73)%32)&v75 {
											v145 = v76
											if v145 != 0 {
												v148 = int32(1)
											} else {
												v148 = int32(-1)
											}
											v157 = v148
										} else {
											if v48 == int32(3) {
												v129 = v39 - v40
												if v129 != 0 {
													v138 = v129
													v157 = v138
												} else {
													if v25 == int32(2) {
														v134 = int32(4)
													} else {
														v134 = int32(16)
													}
													v135 = F_memcmp(m, v36, v38, v134)
													mBase = m.M
													v157 = v135
												}
											} else {
												v84 = int32(3)
												v86 = int32(128)
												v87 = v52 << (uint(v84) % 32) & v86
												if v87 != v56<<(uint(v84)%32)&v86 {
													v145 = v87
													if v145 != 0 {
														v148 = int32(1)
													} else {
														v148 = int32(-1)
													}
													v157 = v148
												} else {
													if base.Ui32(v48) < base.Ui32(int32(5)) {
														v129 = v39 - v40
														if v129 != 0 {
															v138 = v129
															v157 = v138
														} else {
															if v25 == int32(2) {
																v134 = int32(4)
															} else {
																v134 = int32(16)
															}
															v135 = F_memcmp(m, v36, v38, v134)
															mBase = m.M
															v157 = v135
														}
													} else {
														v95 = int32(4)
														v97 = int32(128)
														v98 = v52 << (uint(v95) % 32) & v97
														if v98 != v56<<(uint(v95)%32)&v97 {
															v145 = v98
															if v145 != 0 {
																v148 = int32(1)
															} else {
																v148 = int32(-1)
															}
															v157 = v148
														} else {
															if v48 == int32(5) {
																v129 = v39 - v40
																if v129 != 0 {
																	v138 = v129
																	v157 = v138
																} else {
																	if v25 == int32(2) {
																		v134 = int32(4)
																	} else {
																		v134 = int32(16)
																	}
																	v135 = F_memcmp(m, v36, v38, v134)
																	mBase = m.M
																	v157 = v135
																}
															} else {
																v106 = int32(5)
																v108 = int32(128)
																v109 = v52 << (uint(v106) % 32) & v108
																if v109 != v56<<(uint(v106)%32)&v108 {
																	v145 = v109
																	if v145 != 0 {
																		v148 = int32(1)
																	} else {
																		v148 = int32(-1)
																	}
																	v157 = v148
																} else {
																	if v48 != int32(7) {
																		v129 = v39 - v40
																		if v129 != 0 {
																			v138 = v129
																			v157 = v138
																		} else {
																			if v25 == int32(2) {
																				v134 = int32(4)
																			} else {
																				v134 = int32(16)
																			}
																			v135 = F_memcmp(m, v36, v38, v134)
																			mBase = m.M
																			v157 = v135
																		}
																	} else {
																		v117 = int32(6)
																		v119 = int32(128)
																		v120 = v52 << (uint(v117) % 32) & v119
																		if v120 != v56<<(uint(v117)%32)&v119 {
																			v145 = v120
																			if v145 != 0 {
																				v148 = int32(1)
																			} else {
																				v148 = int32(-1)
																			}
																			v157 = v148
																		} else {
																			v129 = v39 - v40
																			if v129 != 0 {
																				v138 = v129
																				v157 = v138
																			} else {
																				if v25 == int32(2) {
																					v134 = int32(4)
																				} else {
																					v134 = int32(16)
																				}
																				v135 = F_memcmp(m, v36, v38, v134)
																				mBase = m.M
																				v157 = v135
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
								}
							}
						}
					}
				}
			} else {
				v138 = v25 - v33
				v157 = v138
			}
			if int32(0) < v157 {
				v160 = v4
			} else {
				v160 = v9
			}
			return base.I64_extend_i32_u(v160)
		}
	}
}
func F_network_le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v17 = int32(1)
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
			if v19&v17 != 0 {
				v22 = v17
			} else {
				v22 = int32(4)
			}
			v23 = v3 + v22
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
			v25 = int32(1)
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v27&v25 != 0 {
				v30 = v25
			} else {
				v30 = int32(4)
			}
			v31 = v8 + v30
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
			if v24 == v32 {
				v34 = int32(2)
				v35 = v23 + v34
				v37 = v31 + v34
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
				if base.Ui32(v38) < base.Ui32(v39) {
					v41 = v23
				} else {
					v41 = v31
				}
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
				v44 = int32(base.Ui32(v42) >> (uint(int32(3)) % 32))
				v45 = F_memcmp(m, v35, v37, v44)
				mBase = m.M
				if v45 != 0 {
					v137 = v45
					v156 = v137
				} else {
					v47 = v42 & int32(7)
					if v47 == int32(0) {
						v128 = v38 - v39
						if v128 != 0 {
							v137 = v128
							v156 = v137
						} else {
							if v24 == int32(2) {
								v133 = int32(4)
							} else {
								v133 = int32(16)
							}
							v134 = F_memcmp(m, v35, v37, v133)
							mBase = m.M
							v156 = v134
						}
					} else {
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v35))))
						v52 = int32(128)
						v53 = v51 & v52
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v37))))
						if v53 != v55&v52 {
							v144 = v53
							if v144 != 0 {
								v147 = int32(1)
							} else {
								v147 = int32(-1)
							}
							v156 = v147
						} else {
							if v47 == int32(1) {
								v128 = v38 - v39
								if v128 != 0 {
									v137 = v128
									v156 = v137
								} else {
									if v24 == int32(2) {
										v133 = int32(4)
									} else {
										v133 = int32(16)
									}
									v134 = F_memcmp(m, v35, v37, v133)
									mBase = m.M
									v156 = v134
								}
							} else {
								v61 = int32(1)
								v63 = int32(128)
								v64 = v51 << (uint(v61) % 32) & v63
								if v64 != v55<<(uint(v61)%32)&v63 {
									v144 = v64
									if v144 != 0 {
										v147 = int32(1)
									} else {
										v147 = int32(-1)
									}
									v156 = v147
								} else {
									if base.Ui32(v47) < base.Ui32(int32(3)) {
										v128 = v38 - v39
										if v128 != 0 {
											v137 = v128
											v156 = v137
										} else {
											if v24 == int32(2) {
												v133 = int32(4)
											} else {
												v133 = int32(16)
											}
											v134 = F_memcmp(m, v35, v37, v133)
											mBase = m.M
											v156 = v134
										}
									} else {
										v72 = int32(2)
										v74 = int32(128)
										v75 = v51 << (uint(v72) % 32) & v74
										if v75 != v55<<(uint(v72)%32)&v74 {
											v144 = v75
											if v144 != 0 {
												v147 = int32(1)
											} else {
												v147 = int32(-1)
											}
											v156 = v147
										} else {
											if v47 == int32(3) {
												v128 = v38 - v39
												if v128 != 0 {
													v137 = v128
													v156 = v137
												} else {
													if v24 == int32(2) {
														v133 = int32(4)
													} else {
														v133 = int32(16)
													}
													v134 = F_memcmp(m, v35, v37, v133)
													mBase = m.M
													v156 = v134
												}
											} else {
												v83 = int32(3)
												v85 = int32(128)
												v86 = v51 << (uint(v83) % 32) & v85
												if v86 != v55<<(uint(v83)%32)&v85 {
													v144 = v86
													if v144 != 0 {
														v147 = int32(1)
													} else {
														v147 = int32(-1)
													}
													v156 = v147
												} else {
													if base.Ui32(v47) < base.Ui32(int32(5)) {
														v128 = v38 - v39
														if v128 != 0 {
															v137 = v128
															v156 = v137
														} else {
															if v24 == int32(2) {
																v133 = int32(4)
															} else {
																v133 = int32(16)
															}
															v134 = F_memcmp(m, v35, v37, v133)
															mBase = m.M
															v156 = v134
														}
													} else {
														v94 = int32(4)
														v96 = int32(128)
														v97 = v51 << (uint(v94) % 32) & v96
														if v97 != v55<<(uint(v94)%32)&v96 {
															v144 = v97
															if v144 != 0 {
																v147 = int32(1)
															} else {
																v147 = int32(-1)
															}
															v156 = v147
														} else {
															if v47 == int32(5) {
																v128 = v38 - v39
																if v128 != 0 {
																	v137 = v128
																	v156 = v137
																} else {
																	if v24 == int32(2) {
																		v133 = int32(4)
																	} else {
																		v133 = int32(16)
																	}
																	v134 = F_memcmp(m, v35, v37, v133)
																	mBase = m.M
																	v156 = v134
																}
															} else {
																v105 = int32(5)
																v107 = int32(128)
																v108 = v51 << (uint(v105) % 32) & v107
																if v108 != v55<<(uint(v105)%32)&v107 {
																	v144 = v108
																	if v144 != 0 {
																		v147 = int32(1)
																	} else {
																		v147 = int32(-1)
																	}
																	v156 = v147
																} else {
																	if v47 != int32(7) {
																		v128 = v38 - v39
																		if v128 != 0 {
																			v137 = v128
																			v156 = v137
																		} else {
																			if v24 == int32(2) {
																				v133 = int32(4)
																			} else {
																				v133 = int32(16)
																			}
																			v134 = F_memcmp(m, v35, v37, v133)
																			mBase = m.M
																			v156 = v134
																		}
																	} else {
																		v116 = int32(6)
																		v118 = int32(128)
																		v119 = v51 << (uint(v116) % 32) & v118
																		if v119 != v55<<(uint(v116)%32)&v118 {
																			v144 = v119
																			if v144 != 0 {
																				v147 = int32(1)
																			} else {
																				v147 = int32(-1)
																			}
																			v156 = v147
																		} else {
																			v128 = v38 - v39
																			if v128 != 0 {
																				v137 = v128
																				v156 = v137
																			} else {
																				if v24 == int32(2) {
																					v133 = int32(4)
																				} else {
																					v133 = int32(16)
																				}
																				v134 = F_memcmp(m, v35, v37, v133)
																				mBase = m.M
																				v156 = v134
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
								}
							}
						}
					}
				}
			} else {
				v137 = v24 - v32
				v156 = v137
			}
			return base.I64_extend_i32_u(base.B2i32(v156 <= int32(0)))
		}
	}
}
